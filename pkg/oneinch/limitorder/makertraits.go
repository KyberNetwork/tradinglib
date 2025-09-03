package limitorder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var ErrEpochManagerNotAllowed = errors.New("epoch manager allowed only when partialFills and multipleFills enabled")

// source: https://github.com/1inch/limit-order-protocol/blob/master/contracts/libraries/MakerTraitsLib.sol

// MakerTraits defines the maker's preferences for an order in a single uint256
// High bits are used for flags:
// 255 bit NO_PARTIAL_FILLS_FLAG          - if set, the order does not allow partial fills
// 254 bit ALLOW_MULTIPLE_FILLS_FLAG      - if set, the order permits multiple fills
// 253 bit                                - unused
// 252 bit PRE_INTERACTION_CALL_FLAG      - if set, the order requires pre-interaction call
// 251 bit POST_INTERACTION_CALL_FLAG     - if set, the order requires post-interaction call
// 250 bit NEED_CHECK_EPOCH_MANAGER_FLAG  - if set, the order requires to check the epoch manager
// 249 bit HAS_EXTENSION_FLAG             - if set, the order has extension(s)
// 248 bit MAKER_USE_PERMIT2_FLAG         - if set, the order uses permit2
// 247 bit MAKER_UNWRAP_WETH_FLAG         - if set, the order requires to unwrap WETH
//
// Low 200 bits are used for allowed sender, expiration, nonceOrEpoch, and series:
// uint80 last 10 bytes of allowed sender address (0 if any)
// uint40 expiration timestamp (0 if none)
// uint40 nonce or epoch
// uint40 series
type MakerTraits struct {
	value *big.Int
}

const (
	// Bit masks for low 200 bits
	allowedSenderStart = uint(0)
	allowedSenderEnd   = uint(80)
	expirationStart    = uint(80)
	expirationEnd      = uint(120)
	nonceOrEpochStart  = uint(120)
	nonceOrEpochEnd    = uint(160)
	seriesStart        = uint(160)
	seriesEnd          = uint(200)

	// Flag bit positions
	noPartialFillsFlag        = uint(255)
	allowMultipleFillsFlag    = uint(254)
	preInteractionCallFlag    = uint(252)
	postInteractionCallFlag   = uint(251)
	needCheckEpochManagerFlag = uint(250)
	hasExtensionFlag          = uint(249)
	makerUsePermit2Flag       = uint(248)
	makerUnwrapWethFlag       = uint(247)
)

// NewMakerTraits creates a new MakerTraits instance with the given value
func NewMakerTraits(val string) *MakerTraits {
	value := new(big.Int)
	if val != "" {
		if len(val) >= 2 && val[0:2] == "0x" {
			value.SetString(val[2:], 16)
		} else {
			value.SetString(val, 10)
		}
	}
	return &MakerTraits{value: value}
}

type MakerTraitsOption struct {
	AllowPartialFills     bool     `json:"allow_partial_fills"`
	AllowMultipleFills    bool     `json:"allow_multiple_fills"`
	NeedCheckEpochManager bool     `json:"need_check_epoch_manager"`
	UseBitInvalidator     bool     `json:"use_bit_invalidator"`
	NeedPreInteraction    bool     `json:"need_pre_interaction"`
	NeedPostInteraction   bool     `json:"need_post_interaction"`
	UnwrapWeth            bool     `json:"unwrap_weth"`
	HasExtension          bool     `json:"has_extension"`
	IsPrivate             bool     `json:"is_private"`
	Expiration            *big.Int `json:"expiration"`
	NonceOrEpoch          *big.Int `json:"nonce_or_epoch"`
	Series                *big.Int `json:"series"`
	AllowedSender         []byte   `json:"allowed_sender"`
}

type makerTraitsOptionJson struct {
	AllowPartialFills     string `json:"allow_partial_fills"`
	AllowMultipleFills    string `json:"allow_multiple_fills"`
	NeedCheckEpochManager string `json:"need_check_epoch_manager"`
	UseBitInvalidator     string `json:"use_bit_invalidator"`
	NeedPreInteraction    string `json:"need_pre_interaction"`
	NeedPostInteraction   string `json:"need_post_interaction"`
	UnwrapWeth            string `json:"unwrap_weth"`
	HasExtension          string `json:"has_extension"`
	IsPrivate             string `json:"is_private"`
	Expiration            string `json:"expiration"`
	NonceOrEpoch          string `json:"nonce_or_epoch"`
	Series                string `json:"series"`
	AllowedSender         string `json:"allowed_sender"`
}

func (mt *MakerTraits) Decode() MakerTraitsOption {
	return MakerTraitsOption{
		AllowPartialFills:     mt.AllowPartialFills(),
		AllowMultipleFills:    mt.AllowMultipleFills(),
		NeedCheckEpochManager: mt.NeedCheckEpochManager(),
		UseBitInvalidator:     mt.UseBitInvalidator(),
		NeedPreInteraction:    mt.NeedPreInteraction(),
		NeedPostInteraction:   mt.NeedPostInteraction(),
		UnwrapWeth:            mt.UnwrapWeth(),
		HasExtension:          mt.HasExtension(),
		IsPrivate:             mt.IsPrivate(),
		Expiration:            mt.Expiration(),
		NonceOrEpoch:          mt.NonceOrEpoch(),
		Series:                mt.Series(),
		AllowedSender:         mt.AllowedSender(),
	}
}

// DefaultMakerTraits returns a MakerTraits instance with default values
func DefaultMakerTraits() *MakerTraits {
	return NewMakerTraits("")
}

// getBit gets value a bit at the specified position
func (mt *MakerTraits) getBit(pos uint) uint {
	return mt.value.Bit(int(pos))
}

// setBit sets value a bit at the specified position
func (mt *MakerTraits) setBit(pos uint, val uint) *MakerTraits {
	mt.value.SetBit(mt.value, int(pos), val)
	return mt
}

// AllowedSender returns the last 10 bytes of allowed sender address
func (mt *MakerTraits) AllowedSender() []byte {
	val := getMask(mt.value, allowedSenderStart, allowedSenderEnd)
	result := make([]byte, 10)
	val.FillBytes(result) // Fill only last 10 bytes
	return result
}

func (mt *MakerTraits) IsAllowedSender(sender common.Address) bool {
	allowedSender := mt.AllowedSender()
	if bytes.Equal(allowedSender, make([]byte, 10)) {
		return true // If no specific sender is set, allow any sender
	}
	return bytes.Equal(sender.Bytes()[10:], allowedSender)
}

// IsPrivate returns true if the order has a specific allowed sender
func (mt *MakerTraits) IsPrivate() bool {
	return getMask(mt.value, allowedSenderStart, allowedSenderEnd).Sign() != 0
}

// WithAllowedSender sets the allowed sender for the order
func (mt *MakerTraits) WithAllowedSender(sender common.Address) *MakerTraits {
	if sender == (common.Address{}) {
		return mt.WithAnySender()
	}
	// Take last 10 bytes of the address
	val := new(big.Int).SetBytes(sender.Bytes()[10:])
	setMask(mt.value, newBitMask(allowedSenderStart, allowedSenderEnd), val)
	return mt
}

// WithAnySender removes sender check
func (mt *MakerTraits) WithAnySender() *MakerTraits {
	setMask(mt.value, newBitMask(allowedSenderStart, allowedSenderEnd), big.NewInt(0))
	return mt
}

// Expiration returns the expiration timestamp in seconds, nil if no expiration
func (mt *MakerTraits) Expiration() *big.Int {
	val := getMask(mt.value, expirationStart, expirationEnd)
	if val.Sign() == 0 {
		return nil
	}
	return val
}

// WithExpiration sets the expiration timestamp
func (mt *MakerTraits) WithExpiration(expiration *big.Int) *MakerTraits {
	if expiration == nil {
		expiration = big.NewInt(0)
	}
	setMask(mt.value, newBitMask(expirationStart, expirationEnd), expiration)
	return mt
}

// NonceOrEpoch returns the nonce or epoch value
func (mt *MakerTraits) NonceOrEpoch() *big.Int {
	return getMask(mt.value, nonceOrEpochStart, nonceOrEpochEnd)
}

// WithNonce sets the nonce value
func (mt *MakerTraits) WithNonce(nonce *big.Int) *MakerTraits {
	setMask(mt.value, newBitMask(nonceOrEpochStart, nonceOrEpochEnd), nonce)
	return mt
}

// WithEpoch sets the epoch and series values
func (mt *MakerTraits) WithEpoch(series, epoch *big.Int) (*MakerTraits, error) {
	mt.setSeries(series)
	err := mt.enableEpochManagerCheck()
	if err != nil {
		return nil, err
	}
	return mt.WithNonce(epoch), nil
}

// Series returns the current series value
func (mt *MakerTraits) Series() *big.Int {
	return getMask(mt.value, seriesStart, seriesEnd)
}

// HasExtension returns true if order has an extension
func (mt *MakerTraits) HasExtension() bool {
	return mt.getBit(hasExtensionFlag) == 1
}

// WithExtension marks that order has an extension
func (mt *MakerTraits) WithExtension() *MakerTraits {
	return mt.setBit(hasExtensionFlag, 1)
}

// AllowPartialFills returns true if partial fills are allowed
func (mt *MakerTraits) AllowPartialFills() bool {
	return mt.getBit(noPartialFillsFlag) == 0
}

// SetAllowPartialFills allows partial fills for the order
func (mt *MakerTraits) SetAllowPartialFills(allow bool) *MakerTraits {
	return mt.setBit(noPartialFillsFlag, boolToBit(!allow))
}

// AllowMultipleFills returns true if multiple fills are allowed
func (mt *MakerTraits) AllowMultipleFills() bool {
	return mt.getBit(allowMultipleFillsFlag) == 1
}

// SetAllowMultipleFills allows multiple fills for the order
func (mt *MakerTraits) SetAllowMultipleFills(allow bool) *MakerTraits {
	return mt.setBit(allowMultipleFillsFlag, boolToBit(allow))
}

// NeedPreInteraction returns true if maker has pre-interaction
func (mt *MakerTraits) NeedPreInteraction() bool {
	return mt.getBit(preInteractionCallFlag) == 1
}

func (mt *MakerTraits) SetNeedPreInteraction(need bool) *MakerTraits {
	return mt.setBit(preInteractionCallFlag, boolToBit(need))
}

// NeedPostInteraction returns true if maker has post-interaction
func (mt *MakerTraits) NeedPostInteraction() bool {
	return mt.getBit(postInteractionCallFlag) == 1
}

func (mt *MakerTraits) SetNeedPostInteraction(need bool) *MakerTraits {
	return mt.setBit(postInteractionCallFlag, boolToBit(need))
}

// NeedCheckEpochManager returns true if epoch manager is enabled
func (mt *MakerTraits) NeedCheckEpochManager() bool {
	return mt.getBit(needCheckEpochManagerFlag) == 1
}

// IsPermit2 returns true if permit2 is enabled
func (mt *MakerTraits) IsPermit2() bool {
	return mt.getBit(makerUsePermit2Flag) == 1
}

func (mt *MakerTraits) SetUsePermit2(permit bool) *MakerTraits {
	return mt.setBit(makerUsePermit2Flag, boolToBit(permit))
}

// UnwrapWeth returns true if WETH unwrap is enabled
func (mt *MakerTraits) UnwrapWeth() bool {
	return mt.getBit(makerUnwrapWethFlag) == 1
}

func (mt *MakerTraits) SetUnwrapWeth(unwrap bool) *MakerTraits {
	return mt.setBit(makerUnwrapWethFlag, boolToBit(unwrap))
}

// Build returns the final traits value
func (mt *MakerTraits) Build() *big.Int {
	return new(big.Int).Set(mt.value)
}

// UseBitInvalidator returns true if bit invalidator mode is used
func (mt *MakerTraits) UseBitInvalidator() bool {
	return !mt.AllowPartialFills() || !mt.AllowMultipleFills()
}

// IsExpired checks if the order has expired
func (mt *MakerTraits) IsExpired(currentTime int64) bool {
	expiration := mt.Expiration()
	return expiration != nil && expiration.Cmp(big.NewInt(currentTime)) < 0
}

// enableEpochManagerCheck enables epoch manager check
func (mt *MakerTraits) enableEpochManagerCheck() error {
	if mt.UseBitInvalidator() {
		return ErrEpochManagerNotAllowed
	}
	mt.setBit(needCheckEpochManagerFlag, 1)
	return nil
}

// setSeries sets the series value
func (mt *MakerTraits) setSeries(series *big.Int) {
	setMask(mt.value, newBitMask(seriesStart, seriesEnd), series)
}

func (o *MakerTraitsOption) Marshal() ([]byte, error) {
	dto := makerTraitsOptionJson{
		AllowPartialFills:     fmt.Sprintf("%t", o.AllowPartialFills),
		AllowMultipleFills:    fmt.Sprintf("%t", o.AllowMultipleFills),
		NeedCheckEpochManager: fmt.Sprintf("%t", o.NeedCheckEpochManager),
		UseBitInvalidator:     fmt.Sprintf("%t", o.UseBitInvalidator),
		NeedPreInteraction:    fmt.Sprintf("%t", o.NeedPreInteraction),
		NeedPostInteraction:   fmt.Sprintf("%t", o.NeedPostInteraction),
		UnwrapWeth:            fmt.Sprintf("%t", o.UnwrapWeth),
		HasExtension:          fmt.Sprintf("%t", o.HasExtension),
		IsPrivate:             fmt.Sprintf("%t", o.IsPrivate),
	}

	if o.Expiration != nil {
		dto.Expiration = o.Expiration.String()
	}
	if o.NonceOrEpoch != nil {
		dto.NonceOrEpoch = o.NonceOrEpoch.String()
	}
	if o.Series != nil {
		dto.Series = o.Series.String()
	}
	if len(o.AllowedSender) > 0 {
		dto.AllowedSender = hexutil.Encode(o.AllowedSender)
	}

	return json.Marshal(dto)
}

func (o *MakerTraitsOption) Unmarshal(data []byte) error {
	var dto makerTraitsOptionJson
	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	parseBool := func(s string) bool {
		return s == "true" || s == "1"
	}
	parseBig := func(s string) (*big.Int, error) {
		if s == "" {
			return nil, fmt.Errorf("the number is empty")
		}
		bi, ok := new(big.Int).SetString(s, 10)
		if !ok {
			return nil, fmt.Errorf("invalid big.Int: %s", s)
		}
		return bi, nil
	}

	o.AllowPartialFills = parseBool(dto.AllowPartialFills)
	o.AllowMultipleFills = parseBool(dto.AllowMultipleFills)
	o.NeedCheckEpochManager = parseBool(dto.NeedCheckEpochManager)
	o.UseBitInvalidator = parseBool(dto.UseBitInvalidator)
	o.NeedPreInteraction = parseBool(dto.NeedPreInteraction)
	o.NeedPostInteraction = parseBool(dto.NeedPostInteraction)
	o.UnwrapWeth = parseBool(dto.UnwrapWeth)
	o.HasExtension = parseBool(dto.HasExtension)
	o.IsPrivate = parseBool(dto.IsPrivate)

	var err error
	if o.Expiration, err = parseBig(dto.Expiration); err != nil {
		return err
	}
	if o.NonceOrEpoch, err = parseBig(dto.NonceOrEpoch); err != nil {
		return err
	}
	if o.Series, err = parseBig(dto.Series); err != nil {
		return err
	}

	if dto.AllowedSender != "" {
		b, err := hexutil.Decode(dto.AllowedSender)
		if err != nil {
			return fmt.Errorf("invalid allowed_sender: %w", err)
		}
		o.AllowedSender = b
	}

	return nil
}
