package eth

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const signatureLength = 65

func RecoverSignerAddress(hexEncodedHash string, hexEncodedSignature string) (common.Address, error) {
	hash, err := hexutil.Decode(hexEncodedHash)
	if err != nil {
		return common.Address{}, fmt.Errorf("decode hash error: %w", err)
	}

	signature, err := hexutil.Decode(hexEncodedSignature)
	if err != nil {
		return common.Address{}, fmt.Errorf("decode signature error: %w", err)
	}

	// Signature in Ethereum consists of R, S, V; the V is at last byte, R and S are the rest.
	// The standard Ethereum signature is 65 bytes (R: 32 bytes, S: 32 bytes, V: 1 byte).
	// Ensure the signature is 65 bytes and split it into R, S, and V components.
	if len(signature) != signatureLength {
		return common.Address{}, fmt.Errorf("invalid signature length, expect: %d, got: %v", signatureLength, len(signature))
	}

	// Ethereum uses a 'recovery id' for V, but go-ethereum expects V to be 27 or 28.
	signature[64] -= 27

	// Recover the public key from the signature
	pubKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return common.Address{}, fmt.Errorf("signature to public key error: %w", err)
	}

	// Derive the Ethereum address from the public key
	address := crypto.PubkeyToAddress(*pubKey)

	return address, nil
}

func GetFrom(tx *types.Transaction) (common.Address, error) {
	var chainID *big.Int
	// NOTE: handle case "Legacy Homestead Transaction", the tx chainID returns 0 instead of nil.
	if tx.ChainId().Sign() != 0 {
		chainID = tx.ChainId()
	}
	signer := types.LatestSignerForChainID(chainID)
	from, err := types.Sender(signer, tx)
	return from, err
}

func DecodeSignature(sig []byte) (r, s, v *big.Int, err error) {
	if len(sig) != crypto.SignatureLength {
		return nil, nil, nil, errors.New("invalid signature length")
	}

	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64]})

	return r, s, v, nil
}
