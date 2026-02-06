package basefee

import (
	"errors"
	"math/big"

	baseChainEIP1559 "github.com/KyberNetwork/tradinglib/pkg/basefee/op-geth/base-eip1559"
	baseChainParams "github.com/KyberNetwork/tradinglib/pkg/basefee/op-geth/base-params"
	bscChainEIP1559 "github.com/KyberNetwork/tradinglib/pkg/basefee/op-geth/bsc-eip1559"
	bscChainParams "github.com/KyberNetwork/tradinglib/pkg/basefee/op-geth/bsc-params"
	polChainEIP1559 "github.com/KyberNetwork/tradinglib/pkg/basefee/op-geth/pol-eip1559"
	polChainParams "github.com/KyberNetwork/tradinglib/pkg/basefee/op-geth/pol-params"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var errUnknownChainID = errors.New("unknown chain ID")

const (
	EthChainID     uint64 = 1
	BaseChainID    uint64 = 8453
	BscChainID     uint64 = 56
	PolygonChainID uint64 = 137
)

func CalcNextBaseFee(chainId uint64, header *types.Header) (*big.Int, error) {
	switch chainId {
	case EthChainID: // Ethereum mainnet
		return eip1559.CalcBaseFee(params.MainnetChainConfig, header), nil
	case BaseChainID: // Base mainnet
		return baseChainEIP1559.CalcBaseFee(baseChainParams.OptimismTestConfig, header, header.Time), nil
	case BscChainID: // BSC mainnet
		return bscChainEIP1559.CalcBaseFee(bscChainParams.MainnetChainConfig, header), nil
	case PolygonChainID:
		return polChainEIP1559.CalcBaseFee(polChainParams.BorMainnetChainConfig, header), nil
	default:
		return nil, errUnknownChainID
	}
}
