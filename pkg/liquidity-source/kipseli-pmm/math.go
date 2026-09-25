package kipselipmm

import (
	"fmt"
)

// removeFirstLevelsByAmount removes the first price levels until the total quantity is at least amount.
func removeFirstLevelsByAmount(priceLevels []PriceLevel, amount float64) []PriceLevel {
	result := make([]PriceLevel, 0)
	remainingAmount := amount
	for i, level := range priceLevels {
		if remainingAmount <= 0 {
			result = append(result, priceLevels[i:]...)
			break
		}

		if level.Quantity <= remainingAmount {
			remainingAmount -= level.Quantity
			continue
		}

		// level.Quantity > remainingAmount
		result = append(result, PriceLevel{
			Rate:     level.Rate,
			Quantity: level.Quantity - remainingAmount,
		})
		remainingAmount = 0
	}
	return result
}

func calcQuoteAmount(priceLevels []PriceLevel, baseAmount float64) (float64, error) {
	if len(priceLevels) == 0 {
		return 0, ErrEmptyPriceLevels
	}

	totalIn := 0.0
	for _, level := range priceLevels {
		totalIn += level.Quantity
	}
	if totalIn < baseAmount {
		return 0, fmt.Errorf("%w: available %f, required %f", ErrInsufficientLiquidity, totalIn, baseAmount)
	}

	amountOut := float64(0)
	for _, level := range priceLevels {
		currentLevelAmount := min(level.Quantity, baseAmount)
		amountOut += currentLevelAmount * level.Rate
		baseAmount -= currentLevelAmount
		if baseAmount <= 0 {
			break
		}
	}

	return amountOut, nil
}

func calcBaseAmount(priceLevels []PriceLevel, quoteAmount float64) (float64, error) {
	if len(priceLevels) == 0 {
		return 0, ErrEmptyPriceLevels
	}

	totalOut := 0.0
	for _, level := range priceLevels {
		totalOut += level.Quantity * level.Rate
	}
	if totalOut < quoteAmount {
		return 0, fmt.Errorf("%w: available %f, required %f", ErrInsufficientLiquidity, totalOut, quoteAmount)
	}

	amountIn := float64(0)
	for _, level := range priceLevels {
		levelAmountOut := level.Quantity * level.Rate
		currentLevelAmount := min(levelAmountOut, quoteAmount)
		amountIn += currentLevelAmount / level.Rate
		quoteAmount -= currentLevelAmount
		if quoteAmount <= 0 {
			break
		}
	}
	return amountIn, nil
}

func pairKey(base, quote string) string {
	return base + "/" + quote
}
