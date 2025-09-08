package auctiondetail

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/encode"
	"github.com/ethereum/go-ethereum/common"
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
	StartTime       int64              `json:"startTime"`
	Duration        int64              `json:"duration"`
	InitialRateBump int64              `json:"initialRateBump"`
	Points          []AuctionPoint     `json:"points"`
	GasCost         AuctionGasCostInfo `json:"gasCost"`
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
	Delay       int64 `json:"delay"`
	Coefficient int64 `json:"coefficient"`
}

type AuctionGasCostInfo struct {
	GasBumpEstimate  int64 `json:"gasBumpEstimate"`
	GasPriceEstimate int64 `json:"gasPriceEstimate"`
}

// DecodeAuctionDetails decodes auctioncalculator details from hex string
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
func DecodeAuctionDetails(iter *decode.BytesIterator) (AuctionDetails, error) {
	_, err := iter.NextUint160() // Skip the address of the extension
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("skip address of extension: %w", err)
	}
	gasBumpEstimate, err := iter.NextUint24()
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("next gas bump estimate: %w", err)
	}
	gasPriceEstimate, err := iter.NextUint32()
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("next gas price estimate: %w", err)
	}
	startTime, err := iter.NextUint32()
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("next start time: %w", err)
	}
	duration, err := iter.NextUint24()
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("next duration: %w", err)
	}
	initialRateBump, err := iter.NextUint24()
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("next initial rate bump: %w", err)
	}

	points, err := decodeAuctionPoints(iter)
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("decode auctioncalculator points: %w", err)
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

func decodeAuctionPoints(iter *decode.BytesIterator) ([]AuctionPoint, error) {
	pointsLength, err := iter.NextUint8()
	if err != nil {
		return nil, fmt.Errorf("get points length: %w", err)
	}

	points := make([]AuctionPoint, 0, pointsLength)
	for range pointsLength {
		coefficient, err := iter.NextUint24()
		if err != nil {
			return nil, fmt.Errorf("next coefficient: %w", err)
		}
		delay, err := iter.NextUint16()
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
	buf.Write(common.HexToAddress("0x").Bytes())
	buf.Write(encode.EncodeInt64ToBytes(a.GasCost.GasBumpEstimate, 3))
	buf.Write(encode.EncodeInt64ToBytes(a.GasCost.GasPriceEstimate, 4))
	buf.Write(encode.EncodeInt64ToBytes(a.StartTime, 4))
	buf.Write(encode.EncodeInt64ToBytes(a.Duration, 3))
	buf.Write(encode.EncodeInt64ToBytes(a.InitialRateBump, 3))
	buf.WriteByte(byte(len(a.Points)))
	for _, point := range a.Points {
		buf.Write(encode.EncodeInt64ToBytes(point.Coefficient, 3))
		buf.Write(encode.EncodeInt64ToBytes(point.Delay, 2))
	}

	return buf.Bytes()
}
