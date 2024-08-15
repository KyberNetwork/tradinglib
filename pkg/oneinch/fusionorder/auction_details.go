package fusionorder

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/utils"
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
func DecodeAuctionDetails(data string) (AuctionDetails, error) {
	if !utils.IsHexString(data) {
		return AuctionDetails{}, errors.New("invalid auction details data")
	}

	hexData, err := hex.DecodeString(utils.Trim0x(data))
	if err != nil {
		return AuctionDetails{}, fmt.Errorf("decode auction details: %w", err)
	}

	gasBumpEstimate := new(big.Int).SetBytes(hexData[:3])
	gasPriceEstimate := new(big.Int).SetBytes(hexData[3:7])
	startTime := new(big.Int).SetBytes(hexData[7:11])
	duration := new(big.Int).SetBytes(hexData[11:14])
	initialRateBump := new(big.Int).SetBytes(hexData[14:17])

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

func (a AuctionDetails) Encode() string {
	buf := new(bytes.Buffer)
	buf.Write(utils.PadOrTrim(a.GasCost.GasBumpEstimate.Bytes(), 3))
	buf.Write(utils.PadOrTrim(a.GasCost.GasPriceEstimate.Bytes(), 4))
	buf.Write(utils.PadOrTrim(a.StartTime.Bytes(), 4))
	buf.Write(utils.PadOrTrim(a.Duration.Bytes(), 3))
	buf.Write(utils.PadOrTrim(a.InitialRateBump.Bytes(), 3))
	for _, point := range a.Points {
		buf.Write(utils.PadOrTrim(big.NewInt(point.Delay).Bytes(), 3))
		buf.Write(utils.PadOrTrim(big.NewInt(point.Coefficient).Bytes(), 2))
	}

	return utils.Add0x(hex.EncodeToString(buf.Bytes()))
}
