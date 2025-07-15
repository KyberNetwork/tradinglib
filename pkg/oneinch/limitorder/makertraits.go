package limitorder

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

type MakerTraitsOption struct {
	NoPartialFills      bool           `json:"no_partial_fills"`
	MultipleFills       bool           `json:"multiple_fills"`
	PreInteractionCall  bool           `json:"pre_interaction_call"`
	PostInteractionCall bool           `json:"post_interaction_call"`
	CheckEpochManager   bool           `json:"check_epoch_manager"`
	HasExtension        bool           `json:"has_extension"`
	UsePermit2          bool           `json:"use_permit2"`
	UnwrapWETH          bool           `json:"unwrap_weth"`
	AllowedSender       common.Address `json:"allowed_sender"`
	Expiration          int64          `json:"expiration"`
	NonceOrEpoch        int64          `json:"nonce_or_epoch"`
	Series              int64          `json:"series"`
}

const (
	noPartialFillsMask  = 7
	multipleFillsMask   = 6
	preInteractionMask  = 4
	postInteractionMask = 3
	checkEpochMask      = 2
	hasExtensionMask    = 1
	usePermit2Mask      = 0

	unwrapWETHMask = 7
)

var fiveBytes = [5]byte{}

func BuildMakerTraits(values MakerTraitsOption) *big.Int {
	buff := new(bytes.Buffer)
	{
		var lb uint8
		if values.NoPartialFills {
			lb = lb | (0xff & 1 << noPartialFillsMask)
		}
		if values.MultipleFills {
			lb = lb | (1 << multipleFillsMask)
		}
		if values.PreInteractionCall {
			lb = lb | (0xff & 1 << preInteractionMask)
		}
		if values.PostInteractionCall {
			lb = lb | (0xff & 1 << postInteractionMask)
		}
		if values.CheckEpochManager {
			lb = lb | (0xff & 1 << checkEpochMask)
		}
		if values.HasExtension {
			lb = lb | (0xff & 1 << hasExtensionMask)
		}
		if values.UsePermit2 {
			lb = lb | (0xff & 1 << usePermit2Mask)
		}
		buff.WriteByte(lb)
	}
	{
		var unwrapETH uint8
		if values.UnwrapWETH {
			unwrapETH = 1 << unwrapWETHMask
		}
		_ = buff.WriteByte(unwrapETH)
	}
	_, _ = buff.Write(fiveBytes[:])
	_, _ = buff.Write(math.PaddedBigBytes(big.NewInt(values.Series), 5))
	_, _ = buff.Write(math.PaddedBigBytes(big.NewInt(values.NonceOrEpoch), 5))
	_, _ = buff.Write(math.PaddedBigBytes(big.NewInt(values.Expiration), 5))
	_, _ = buff.Write(values.AllowedSender[10:])

	return big.NewInt(0).SetBytes(buff.Bytes())
}

func DecodeMarkerTraits(makerTraits *big.Int) (MakerTraitsOption, error) {
	data := math.PaddedBigBytes(makerTraits, common.HashLength)
	res := MakerTraitsOption{}
	buff := bytes.NewBuffer(data)
	{
		b, _ := buff.ReadByte()
		if b>>noPartialFillsMask != 0 {
			res.NoPartialFills = true
		}
		if (b>>multipleFillsMask)&1 != 0 {
			res.MultipleFills = true
		}
		if (b>>preInteractionMask)&1 != 0 {
			res.PreInteractionCall = true
		}
		if (b>>postInteractionMask)&1 != 0 {
			res.PostInteractionCall = true
		}
		if (b>>checkEpochMask)&1 != 0 {
			res.CheckEpochManager = true
		}
		if (b>>hasExtensionMask)&1 != 0 {
			res.HasExtension = true
		}
		if (b & 1) != 0 {
			res.UsePermit2 = true
		}
	}
	{
		b, _ := buff.ReadByte()
		if b>>unwrapWETHMask != 0 {
			res.UnwrapWETH = true
		}
	}
	var temp [5]byte
	if _, err := buff.Read(temp[:]); err != nil {
		return MakerTraitsOption{}, err
	}
	_, err := buff.Read(temp[:])
	if err != nil {
		return MakerTraitsOption{}, err
	}
	series := big.NewInt(0).SetBytes(temp[:])
	res.Series = series.Int64()
	_, err = buff.Read(temp[:])
	if err != nil {
		return MakerTraitsOption{}, err
	}
	nonce := big.NewInt(0).SetBytes(temp[:])
	res.NonceOrEpoch = nonce.Int64()
	_, err = buff.Read(temp[:])
	if err != nil {
		return MakerTraitsOption{}, err
	}
	expire := big.NewInt(0).SetBytes(temp[:])
	res.Expiration = expire.Int64()
	var addr common.Address
	_, err = buff.Read(addr[10:])
	if err != nil {
		return MakerTraitsOption{}, err
	}
	res.AllowedSender = addr
	return res, nil
}
