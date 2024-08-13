package fusionorder

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
)

const (
	MaxUint24 = 1<<24 - 1
)

var (
	ErrGasBumpEstimateTooLarge  = errors.New("gas bump estimate is too large")
	ErrGasPriceEstimateTooLarge = errors.New("gas price estimate is too large")
	ErrStartTimeTooLarge        = errors.New("start time is too large")
	ErrDurationTooLarge         = errors.New("duration is too large")
	ErrInitialRateBumpTooLarge  = errors.New("initial rate bump is too large")
)

type AuctionDetails struct {
	StartTime       *big.Int
	Duration        *big.Int
	InitialRateBump *big.Int
	Points          []AuctionPoint
	GasCost         AuctionGasCostInfo
}

func NewAuctionDetails(
	startTime *big.Int,
	initialRateBump *big.Int,
	duration *big.Int,
	points []AuctionPoint,
	gasCost AuctionGasCostInfo,
) (AuctionDetails, error) {
	if gasCost.GasPriceEstimate == nil {
		gasCost.GasPriceEstimate = big.NewInt(0)
	}
	if gasCost.GasBumpEstimate == nil {
		gasCost.GasBumpEstimate = big.NewInt(0)
	}

	if gasCost.GasBumpEstimate.Cmp(big.NewInt(MaxUint24)) > 0 { // gasCost.GasBumpEstimate > MaxUint24
		return AuctionDetails{}, ErrGasBumpEstimateTooLarge
	}
	if gasCost.GasPriceEstimate.Cmp(big.NewInt(math.MaxUint32)) > 0 { // gasCost.GasPriceEstimate > MaxUint32
		return AuctionDetails{}, ErrGasPriceEstimateTooLarge
	}
	if startTime.Cmp(big.NewInt(math.MaxUint32)) > 0 { // startTime > MaxUint32
		return AuctionDetails{}, ErrStartTimeTooLarge
	}
	if duration.Cmp(big.NewInt(MaxUint24)) > 0 { // duration > MaxUint24
		return AuctionDetails{}, ErrDurationTooLarge
	}
	if initialRateBump.Cmp(big.NewInt(MaxUint24)) > 0 { // initialRateBump > MaxUint24
		return AuctionDetails{}, ErrInitialRateBumpTooLarge
	}

	return AuctionDetails{
		StartTime:       startTime,
		Duration:        duration,
		InitialRateBump: initialRateBump,
		Points:          points,
		GasCost:         gasCost,
	}, nil
}

type AuctionPoint struct {
	Delay       int64
	Coefficient int64
}

type AuctionGasCostInfo struct {
	GasBumpEstimate  *big.Int
	GasPriceEstimate *big.Int
}
