package mev

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

type GasBundleEstimator struct {
	ethClient *ethclient.Client
}

func NewGasBundleEstimator(ethClient *ethclient.Client) GasBundleEstimator {
	return GasBundleEstimator{
		ethClient: ethClient,
	}
}

func (g GasBundleEstimator) EstimateBundleGas(
	_ context.Context,
	messages []ethereum.CallMsg,
	overrides *map[common.Address]gethclient.OverrideAccount,
) ([]uint64, error) {
	bundles := make([]interface{}, 0, len(messages))
	for _, msg := range messages {
		bundles = append(bundles, ToCallArg(msg))
	}

	var gasEstimateCost []hexutil.Uint64

	err := g.ethClient.Client().Call(
		&gasEstimateCost, ETHEstimateGasBundleMethod,
		map[string]interface{}{
			"transactions": bundles,
		}, "latest", overrides,
	)
	if err != nil {
		return nil, err
	}
	result := make([]uint64, 0, len(gasEstimateCost))

	for _, gasEstimate := range gasEstimateCost {
		result = append(result, uint64(gasEstimate))
	}
	return result, nil
}
