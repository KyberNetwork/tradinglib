package fusionorder

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
)

const (
	MaxUint24 = 1<<24 - 1
)

var (
	ErrGasBumpEstimateInvalid  = errors.New("gas bump estimate is invalid")
	ErrGasPriceEstimateInvalid = errors.New("gas price estimate is invalid")
	ErrStartTimeInvalid        = errors.New("start time is invalid")
	ErrDurationInvalid         = errors.New("duration is invalid")
	ErrInitialRateBumpInvalid  = errors.New("initial rate bump is invalid")
)

type AuctionDetails struct {
	StartTime       int64
	Duration        int64
	InitialRateBump int64
	Points          []AuctionPoint
	GasCost         AuctionGasCostInfo
}

func NewAuctionDetails(
	startTime int64,
	initialRateBump int64,
	duration int64,
	points []AuctionPoint,
	gasCost AuctionGasCostInfo,
) (AuctionDetails, error) {
	if gasCost.GasBumpEstimate > MaxUint24 || gasCost.GasBumpEstimate < 0 {
		return AuctionDetails{}, ErrGasBumpEstimateInvalid
	}
	if gasCost.GasPriceEstimate > math.MaxUint32 || gasCost.GasPriceEstimate < 0 {
		return AuctionDetails{}, ErrGasPriceEstimateInvalid
	}
	if startTime > math.MaxUint32 || startTime <= 0 {
		return AuctionDetails{}, ErrStartTimeInvalid
	}
	if duration > MaxUint24 || duration <= 0 {
		return AuctionDetails{}, ErrDurationInvalid
	}
	if initialRateBump > MaxUint24 || initialRateBump < 0 {
		return AuctionDetails{}, ErrInitialRateBumpInvalid
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
	GasBumpEstimate  int64
	GasPriceEstimate int64
}

// DecodeAuctionDetails decodes auction details from hex string
// ```
//
//	struct AuctionDetails {
//	     bytes3 gasBumpEstimate;
//	     bytes4 gasPriceEstimate;
//	     bytes4 auctionStartTime;
//	     bytes3 auctionDuration;
//	     bytes3 initialRateBump;
//	     (bytes3,bytes2)[N] pointsAndTimeDeltas;
//	 }
//
// ```
// Logic is copied from
// https://etherscan.io/address/0xfb2809a5314473e1165f6b58018e20ed8f07b840#code#F23#L140
// nolint: gomnd
func DecodeAuctionDetails(data []byte) (AuctionDetails, error) {
	bi := decode.NewBytesIterator(data)

	gasBumpEstimate, err := bi.NextUint24()
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("next gas bump estimate: %w", err)
	}
	gasPriceEstimate, err := bi.NextUint32()
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("next gas price estimate: %w", err)
	}
	startTime, err := bi.NextUint32()
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("next start time: %w", err)
	}
	duration, err := bi.NextUint24()
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("next duration: %w", err)
	}
	initialRateBump, err := bi.NextUint24()
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("next initial rate bump: %w", err)
	}

	points, err := decodeAuctionPoints(bi.RemainingData())
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("decode auction points: %w", err)
	}

	return NewAuctionDetails(
		int64(startTime),
		int64(initialRateBump),
		int64(duration),
		points,
		AuctionGasCostInfo{
			GasBumpEstimate:  int64(gasBumpEstimate),
			GasPriceEstimate: int64(gasPriceEstimate),
		},
	)
}

func decodeAuctionPoints(data []byte) ([]AuctionPoint, error) {
	bi := decode.NewBytesIterator(data)
	pointsLength, err := bi.NextUint8()
	if err != nil {
		return nil, fmt.Errorf("get points length: %w", err)
	}

	points := make([]AuctionPoint, 0, pointsLength)
	for range pointsLength {
		coefficient, err := bi.NextUint24()
		if err != nil {
			return nil, fmt.Errorf("next coefficient: %w", err)
		}
		delay, err := bi.NextUint16()
		if err != nil {
			return nil, fmt.Errorf("next delay: %w", err)
		}
		points = append(points, AuctionPoint{
			Coefficient: int64(coefficient),
			Delay:       int64(delay),
		})
	}
	return points, nil
}

// Encode encodes AuctionDetails to bytes
// nolint: gomnd
func (a AuctionDetails) Encode() []byte {
	buf := new(bytes.Buffer)
	buf.Write(encodeInt64ToBytes(a.GasCost.GasBumpEstimate, 3))
	buf.Write(encodeInt64ToBytes(a.GasCost.GasPriceEstimate, 4))
	buf.Write(encodeInt64ToBytes(a.StartTime, 4))
	buf.Write(encodeInt64ToBytes(a.Duration, 3))
	buf.Write(encodeInt64ToBytes(a.InitialRateBump, 3))
	buf.WriteByte(byte(len(a.Points)))
	for _, point := range a.Points {
		buf.Write(encodeInt64ToBytes(point.Coefficient, 3))
		buf.Write(encodeInt64ToBytes(point.Delay, 2))
	}

	return buf.Bytes()
}
