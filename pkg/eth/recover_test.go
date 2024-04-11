package eth_test

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/require"
)

func TestRecover(t *testing.T) {
	hashHex := "0xef6777697377372751e5daa1772d3b2b03f6f50b85dd3d7a79b4f12381897ec8"
	signatureHex := "0xf89e6165aa337b8ebf74effa1a2b0980cc54e3780af077337b12438f30c901ab69ad655518007bd15073c919780ac61d3a4d5918c8966ffdb07256776ca6223f1c" // nolint: lll

	addr, err := eth.RecoverSignerAddress(hashHex, signatureHex)
	require.NoError(t, err)

	t.Log(addr.String())
}

// nolint: funlen
func TestGetFrom(t *testing.T) {
	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	// Derive the public key and address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Test data: chain ID and transaction types
	tests := []struct {
		name       string
		chainID    *big.Int
		signer     types.Signer
		txType     byte
		shouldPass bool
	}{
		{
			"Legacy Homestead Transaction",
			big.NewInt(1),
			types.HomesteadSigner{},
			types.LegacyTxType,
			true,
		},
		{
			"EIP-155 Replay Protected Transaction",
			big.NewInt(1),
			types.NewEIP155Signer(big.NewInt(1)),
			types.LegacyTxType, true,
		},
		{
			"EIP-2930 Access List Transaction",
			big.NewInt(1),
			types.NewEIP2930Signer(big.NewInt(1)),
			types.AccessListTxType,
			true,
		},
		{
			"EIP-1559 Dynamic Fee Transaction",
			big.NewInt(1),
			types.NewLondonSigner(big.NewInt(1)),
			types.DynamicFeeTxType,
			true,
		},
		{
			"EIP-4844 Blob Transaction",
			big.NewInt(1),
			types.NewCancunSigner(big.NewInt(1)),
			types.BlobTxType,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new transaction of the specified type
			tx := createTestTransaction(tt.txType, tt.chainID)

			// Sign the transaction
			signedTx, err := types.SignTx(tx, tt.signer, privateKey)
			if err != nil {
				t.Fatal(err)
			}

			// Get the signer address from the signed transaction
			signerAddress, err := eth.GetFrom(signedTx)
			if tt.shouldPass {
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, address, signerAddress, "The recovered address should match the original address")
			} else {
				assert.Error(t, err, "Expected an error")
			}
		})
	}
}

func createTestTransaction(txType byte, chainID *big.Int) *types.Transaction {
	switch txType {
	case types.LegacyTxType:
		return types.NewTransaction(0,
			common.HexToAddress("0x0000000000000000000000000000000000000000"),
			big.NewInt(0), 21000, big.NewInt(1), nil)
	case types.AccessListTxType:
		return types.NewTx(&types.AccessListTx{
			ChainID:    chainID,
			Nonce:      0,
			To:         &common.Address{},
			Value:      big.NewInt(0),
			Gas:        21000,
			GasPrice:   big.NewInt(1),
			Data:       nil,
			AccessList: types.AccessList{},
		})
	case types.DynamicFeeTxType:
		return types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     0,
			To:        &common.Address{},
			Value:     big.NewInt(0),
			Gas:       21000,
			GasFeeCap: big.NewInt(1),
			GasTipCap: big.NewInt(1),
			Data:      nil,
		})
	case types.BlobTxType:
		uint256ChainID, _ := uint256.FromBig(chainID)
		return types.NewTx(&types.BlobTx{
			ChainID: uint256ChainID,
			Nonce:   0,
			To:      common.Address{},
			Value:   uint256.NewInt(0),
			Gas:     21000,
			// Additional fields for blob transactions
		})
	default:
		panic("Unsupported transaction type")
	}
}
