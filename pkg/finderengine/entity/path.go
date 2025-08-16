package entity

import (
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/finderengine/utils"
)

type Path struct {
	ID             string
	AmountIn       *big.Int
	AmountOut      *big.Int
	AmountOutPrice float64
	GasUsed        int64
	GasFeePrice    float64
	L1GasFeePrice  float64
	TokenOrders    []string
	HopOrders      []*Hop
}

func NewPath(amountIn *big.Int) *Path {
	return &Path{
		AmountIn:    new(big.Int).Set(amountIn),
		TokenOrders: []string{},
		HopOrders:   []*Hop{},
	}
}

func (p *Path) AddToken(token string) *Path {
	p.TokenOrders = append(p.TokenOrders, token)
	return p
}

func (p *Path) AddHop(hop *Hop) *Path {
	p.HopOrders = append(p.HopOrders, hop)
	return p
}

func (p *Path) SetAmountOutAndPrice(
	amountOut *big.Int,
	decimals uint8,
	price float64,
) *Path {
	p.AmountOut.Set(amountOut)
	p.AmountOutPrice = utils.CalcAmountPrice(amountOut, decimals, price)

	return p
}

func (p *Path) SetGasUsedAndPrice(
	gasUsed int64,
	gasPrice *big.Int,
	gasTokenDecimals uint8,
	gasTokenPrice float64,
	l1GasFeePrice float64,
) *Path {
	p.GasUsed = gasUsed

	var gasFee big.Int
	gasFee.SetInt64(gasUsed)
	gasFee.Mul(&gasFee, gasPrice)

	p.GasFeePrice = utils.CalcAmountPrice(&gasFee, gasTokenDecimals, gasTokenPrice)

	p.L1GasFeePrice = l1GasFeePrice

	return p
}

func (p *Path) Clone() *Path {
	return &Path{
		ID:             p.ID,
		AmountIn:       new(big.Int).Set(p.AmountIn),
		AmountOut:      new(big.Int).Set(p.AmountOut),
		AmountOutPrice: p.AmountOutPrice,
		GasUsed:        p.GasUsed,
		GasFeePrice:    p.GasFeePrice,
		L1GasFeePrice:  p.L1GasFeePrice,
		HopOrders:      append([]*Hop{}, p.HopOrders...),
		TokenOrders:    append([]string{}, p.TokenOrders...),
	}
}

func (p *Path) Cmp(y *Path, gasIncluded bool) int {
	priceAvailable := p.AmountOutPrice != 0 || y.AmountOutPrice != 0

	if gasIncluded && priceAvailable {
		xValue := p.AmountOutPrice - p.GasFeePrice - p.L1GasFeePrice
		yValue := y.AmountOutPrice - y.GasFeePrice - y.L1GasFeePrice

		if utils.AlmostEqual(xValue, yValue) {
			return p.AmountOut.Cmp(y.AmountOut)
		}

		if xValue < yValue {
			return -1
		} else {
			return 1
		}
	}

	return p.AmountOut.Cmp(y.AmountOut)
}
