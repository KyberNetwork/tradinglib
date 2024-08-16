package fusionorder

import (
	"bytes"
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
	if gasCost.GasBumpEstimate > MaxUint24 {
		return AuctionDetails{}, ErrGasBumpEstimateTooLarge
	}
	if gasCost.GasPriceEstimate > math.MaxUint32 {
		return AuctionDetails{}, ErrGasPriceEstimateTooLarge
	}
	if startTime > math.MaxUint32 {
		return AuctionDetails{}, ErrStartTimeTooLarge
	}
	if duration > MaxUint24 {
		return AuctionDetails{}, ErrDurationTooLarge
	}
	if initialRateBump > MaxUint24 {
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
func DecodeAuctionDetails(hexData []byte) (AuctionDetails, error) {
	gasBumpEstimate := new(big.Int).SetBytes(hexData[:3]).Int64()
	gasPriceEstimate := new(big.Int).SetBytes(hexData[3:7]).Int64()
	startTime := new(big.Int).SetBytes(hexData[7:11]).Int64()
	duration := new(big.Int).SetBytes(hexData[11:14]).Int64()
	initialRateBump := new(big.Int).SetBytes(hexData[14:17]).Int64()

	points := decodeAuctionPoints(hexData[17:])

	return NewAuctionDetails(
		startTime,
		initialRateBump,
		duration,
		points,
		AuctionGasCostInfo{
			GasBumpEstimate:  gasBumpEstimate,
			GasPriceEstimate: gasPriceEstimate,
		})

}

func decodeAuctionPoints(data []byte) []AuctionPoint {
	points := make([]AuctionPoint, 0)
	for len(data) > 0 {
		coefficient := new(big.Int).SetBytes(data[:3]).Int64()
		delay := new(big.Int).SetBytes(data[3:5]).Int64()
		points = append(points, AuctionPoint{
			Coefficient: coefficient,
			Delay:       delay,
		})
		data = data[5:]
	}
	return points
}

func (a AuctionDetails) Encode() []byte {
	buf := new(bytes.Buffer)
	buf.Write(encodeInt64ToBytes(a.GasCost.GasBumpEstimate, 3))
	buf.Write(encodeInt64ToBytes(a.GasCost.GasPriceEstimate, 4))
	buf.Write(encodeInt64ToBytes(a.StartTime, 4))
	buf.Write(encodeInt64ToBytes(a.Duration, 3))
	buf.Write(encodeInt64ToBytes(a.InitialRateBump, 3))
	for _, point := range a.Points {
		buf.Write(encodeInt64ToBytes(point.Coefficient, 3))
		buf.Write(encodeInt64ToBytes(point.Delay, 2))
	}

	return buf.Bytes()
}
