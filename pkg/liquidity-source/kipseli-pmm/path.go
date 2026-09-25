package kipselipmm

import "fmt"

type PairState struct {
	SoldAmount   float64
	BoughtAmount float64
}

type Path struct {
	TokenIn  Token
	TokenOut Token

	// private state, do not set directly
	numOfPairs          int
	isSells             []bool // isSell=true means tokenIn is quote, tokenOut is base
	pairs               []Pair
	pairStates          []PairState
	minTradeBaseAmounts []float64
}

func (p Path) AddPair(isSell bool, pair Pair, state PairState, minTradeBaseAmount float64) Path {
	p.numOfPairs += 1
	p.isSells = append(p.isSells, isSell)
	p.pairs = append(p.pairs, pair)
	p.pairStates = append(p.pairStates, state)
	// if there is no min trade base amount specified, use 1% of first level quantity
	// nolint: nestif
	if minTradeBaseAmount <= 0 {
		if isSell {
			if len(pair.Sell) > 0 {
				minTradeBaseAmount = pair.Sell[0].Quantity * 0.01
			}
		} else {
			if len(pair.Buy) > 0 {
				minTradeBaseAmount = pair.Buy[0].Quantity * 0.01
			}
		}
	}

	p.minTradeBaseAmounts = append(p.minTradeBaseAmounts, minTradeBaseAmount)
	return p
}

func (p Path) CalcAmountOut(amountIn float64) (float64, []SwapLog, error) {
	if p.numOfPairs == 0 {
		return 0, nil, ErrNoPathFound
	}

	amountOut := amountIn
	swapLogs := make([]SwapLog, 0, p.numOfPairs)
	var err error
	for i, pair := range p.pairs {
		isSell := p.isSells[i]
		pairState := p.pairStates[i]
		minTradeBaseAmount := p.minTradeBaseAmounts[i]
		// check min trade base amount
		if tokenInIsBase(isSell) && amountIn < minTradeBaseAmount {
			return 0, nil, fmt.Errorf("pair %d %s/%s: %w: %f < %f",
				i, pair.Base, pair.Quote, ErrAmountInTooSmall, amountIn, minTradeBaseAmount)
		}

		var amountOutThisStep float64
		amountOutThisStep, err = calcAmountOut(pair, pairState, isSell, amountOut)
		if err != nil {
			return 0, nil, fmt.Errorf("pair %d %s/%s: %w", i, pair.Base, pair.Quote, err)
		}

		if tokenOutIsBase(isSell) && amountOutThisStep < minTradeBaseAmount {
			return 0, nil, fmt.Errorf("pair %d %s/%s: %w: %f < %f",
				i, pair.Base, pair.Quote, ErrAmountOutTooSmall, amountOutThisStep, minTradeBaseAmount)
		}

		var soldBaseAmount, boughtBaseAmount float64
		if isSell {
			soldBaseAmount = amountOutThisStep
		} else {
			boughtBaseAmount = amountIn
		}
		swapLog := SwapLog{
			Base:             pair.Base,
			Quote:            pair.Quote,
			AmountIn:         amountIn,
			AmountOut:        amountOutThisStep,
			IsSell:           isSell,
			SoldBaseAmount:   soldBaseAmount,
			BoughtBaseAmount: boughtBaseAmount,
		}
		swapLogs = append(swapLogs, swapLog)
		amountIn = amountOutThisStep  // for next step
		amountOut = amountOutThisStep // final result
	}

	return amountOut, swapLogs, nil
}

func calcAmountOut(pair Pair, pairState PairState, isSell bool, amountIn float64) (float64, error) {
	if isSell {
		// selling base to get quote, remove base, receive quote
		// => amountIn will be quote token, amountOut will be base token
		priceLevels := removeFirstLevelsByAmount(pair.Sell, pairState.SoldAmount)
		return calcBaseAmount(priceLevels, amountIn)
	} else {
		// selling quote to get base, remove quote, receive base
		// => amountIn will be base token, amountOut will be quote token
		priceLevels := removeFirstLevelsByAmount(pair.Buy, pairState.BoughtAmount)
		return calcQuoteAmount(priceLevels, amountIn)
	}
}

func (p Path) CalcAmountIn(amountOut float64) (float64, []SwapLog, error) {
	if p.numOfPairs == 0 {
		return 0, nil, ErrNoPathFound
	}

	amountIn := amountOut
	swapLogs := make([]SwapLog, 0, p.numOfPairs)
	var err error
	// reverse order
	for i := len(p.pairs) - 1; i >= 0; i-- {
		pair := p.pairs[i]
		isSell := p.isSells[i]
		pairState := p.pairStates[i]
		minTradeBaseAmount := p.minTradeBaseAmounts[i]
		// check min trade base amount
		if tokenOutIsBase(isSell) && amountOut < minTradeBaseAmount {
			return 0, nil, fmt.Errorf("pair %d %s/%s: %w: %f < %f",
				i, pair.Base, pair.Quote, ErrAmountOutTooSmall, amountOut, minTradeBaseAmount)
		}

		var amountInThisStep float64
		amountInThisStep, err = calcAmountIn(pair, pairState, isSell, amountIn)
		if err != nil {
			return 0, nil, fmt.Errorf("pair %d %s/%s: %w", i, pair.Base, pair.Quote, err)
		}

		if tokenInIsBase(isSell) && amountInThisStep < minTradeBaseAmount {
			return 0, nil, fmt.Errorf("pair %d %s/%s: %w: %f < %f",
				i, pair.Base, pair.Quote, ErrAmountInTooSmall, amountInThisStep, minTradeBaseAmount)
		}

		var soldBaseAmount, boughtBaseAmount float64
		if isSell {
			soldBaseAmount = amountOut
		} else {
			boughtBaseAmount = amountInThisStep
		}
		swapLog := SwapLog{
			Base:             pair.Base,
			Quote:            pair.Quote,
			AmountIn:         amountInThisStep,
			AmountOut:        amountOut,
			IsSell:           isSell,
			SoldBaseAmount:   soldBaseAmount,
			BoughtBaseAmount: boughtBaseAmount,
		}
		swapLogs = append([]SwapLog{swapLog}, swapLogs...) // prepend to keep order
		amountOut = amountInThisStep                       // for next step
		amountIn = amountInThisStep                        // final result
	}

	return amountIn, swapLogs, nil
}

func calcAmountIn(pair Pair, pairState PairState, isSell bool, amountOut float64) (float64, error) {
	if isSell {
		// selling base to get quote, remove base, receive quote
		// => amountOut will be base token, amountIn will be quote token
		priceLevels := removeFirstLevelsByAmount(pair.Sell, pairState.SoldAmount)
		return calcQuoteAmount(priceLevels, amountOut)
	} else {
		// selling quote to get base, remove quote, receive base
		// => amountOut will be quote token, amountIn will be base token
		priceLevels := removeFirstLevelsByAmount(pair.Buy, pairState.BoughtAmount)
		return calcBaseAmount(priceLevels, amountOut)
	}
}

func tokenInIsBase(isSell bool) bool {
	return !isSell
}

func tokenOutIsBase(isSell bool) bool {
	return isSell
}
