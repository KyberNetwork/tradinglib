//nolint:testpackage
package limitorder

import (
	"encoding/json"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

var (
	// UINT_160_MAX represents the maximum value for a uint160
	UINT_160_MAX = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 160), big.NewInt(1))
	// UINT_40_MAX represents the maximum value for a uint40
	UINT_40_MAX = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 40), big.NewInt(1))
)

func TestMakerTraits_NewWithHexValue(t *testing.T) {
	hexValue := "0x4000000000000000000000000000000000006777c0f900000000000000000000"
	traits := NewMakerTraits(hexValue)

	// Test that the value was correctly parsed
	assert.Equal(t, hexValue[2:], traits.Build().Text(16))

	// Test specific bits and values that should be set based on this hex value
	// The hex value 0x4000... has the following properties:
	// - Bit 254 (ALLOW_MULTIPLE_FILLS_FLAG) is set to 1
	assert.True(t, traits.AllowMultipleFills())

	expectedExpiration := big.NewInt(1735901433)
	assert.Equal(t, expectedExpiration, traits.Expiration())

	// Test that other flags are not set
	assert.True(t, traits.AllowPartialFills())
	assert.False(t, traits.NeedPreInteraction())
	assert.False(t, traits.NeedPostInteraction())
	assert.False(t, traits.NeedCheckEpochManager())
	assert.False(t, traits.HasExtension())
	assert.False(t, traits.IsPermit2())
	assert.False(t, traits.UnwrapWeth())
}

func TestMakerTraits_AllowedSender(t *testing.T) {
	traits := DefaultMakerTraits()

	// Create an address with value 1337
	sender := common.BigToAddress(big.NewInt(1337))

	traits.WithAllowedSender(sender)
	senderHalf := traits.AllowedSender()

	// Compare the last 10 bytes
	assert.Equal(t, sender.Bytes()[10:], senderHalf)
}

func TestMakerTraits_Nonce(t *testing.T) {
	traits := DefaultMakerTraits()

	// Test normal nonce (1 << 10)
	nonce := new(big.Int).Lsh(big.NewInt(1), 10)
	traits.WithNonce(nonce)
	assert.True(t, nonce.Cmp(traits.NonceOrEpoch()) == 0)

	// Test too large nonce (1 << 50)
	bigNonce := new(big.Int).Lsh(big.NewInt(1), 50)
	traits.WithNonce(bigNonce)
	// In Go, we handle overflow by masking, so we should get the masked value
	masked := new(big.Int).And(bigNonce, new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 40), big.NewInt(1)))
	assert.True(t, masked.Cmp(traits.NonceOrEpoch()) == 0)
}

func TestMakerTraits_Expiration(t *testing.T) {
	traits := DefaultMakerTraits()
	expiration := big.NewInt(1000000)

	traits.WithExpiration(expiration)
	assert.Equal(t, expiration, traits.Expiration())
}

func TestMakerTraits_Epoch(t *testing.T) {
	traits := DefaultMakerTraits()
	series := big.NewInt(100)
	epoch := big.NewInt(1)

	traits, err := traits.SetAllowPartialFills(true).SetAllowMultipleFills(true).WithEpoch(series, epoch)
	assert.NoError(t, err)
	assert.Equal(t, series, traits.Series())
	assert.Equal(t, epoch, traits.NonceOrEpoch())
	assert.True(t, traits.NeedCheckEpochManager())
}

func TestMakerTraits_Extension(t *testing.T) {
	traits := DefaultMakerTraits()
	assert.False(t, traits.HasExtension())

	traits.WithExtension()
	assert.True(t, traits.HasExtension())
}

func TestMakerTraits_PartialFills(t *testing.T) {
	traits := DefaultMakerTraits()
	assert.True(t, traits.AllowPartialFills())

	traits.SetAllowPartialFills(false)
	assert.False(t, traits.AllowPartialFills())

	traits.SetAllowPartialFills(true)
	assert.True(t, traits.AllowPartialFills())
}

func TestMakerTraits_MultipleFills(t *testing.T) {
	traits := DefaultMakerTraits()
	assert.False(t, traits.AllowMultipleFills())

	traits.SetAllowMultipleFills(true)
	assert.True(t, traits.AllowMultipleFills())

	traits.SetAllowMultipleFills(false)
	assert.False(t, traits.AllowMultipleFills())
}

func TestMakerTraits_PreInteraction(t *testing.T) {
	traits := DefaultMakerTraits()
	assert.False(t, traits.NeedPreInteraction())

	traits.SetNeedPreInteraction(true)
	assert.True(t, traits.NeedPreInteraction())

	traits.SetNeedPreInteraction(false)
	assert.False(t, traits.NeedPreInteraction())
}

func TestMakerTraits_PostInteraction(t *testing.T) {
	traits := DefaultMakerTraits()
	assert.False(t, traits.NeedPostInteraction())

	traits.SetNeedPostInteraction(true)
	assert.True(t, traits.NeedPostInteraction())

	traits.SetNeedPostInteraction(false)
	assert.False(t, traits.NeedPostInteraction())
}

func TestMakerTraits_Permit2(t *testing.T) {
	traits := DefaultMakerTraits()
	assert.False(t, traits.IsPermit2())

	traits.SetUsePermit2(true)
	assert.True(t, traits.IsPermit2())

	traits.SetUsePermit2(false)
	assert.False(t, traits.IsPermit2())
}

func TestMakerTraits_NativeUnwrap(t *testing.T) {
	traits := DefaultMakerTraits()
	assert.False(t, traits.UnwrapWeth())

	traits.SetUnwrapWeth(true)
	assert.True(t, traits.UnwrapWeth())

	traits.SetUnwrapWeth(false)
	assert.False(t, traits.UnwrapWeth())
}

func TestMakerTraits_IsExpired(t *testing.T) {
	traits := DefaultMakerTraits()

	// Test with no expiration set
	assert.False(t, traits.IsExpired(1704279454)) // current timestamp

	// Test with future expiration
	futureTime := int64(1704279454 + 3600) // current time + 1 hour
	traits.WithExpiration(big.NewInt(futureTime))
	assert.False(t, traits.IsExpired(1704279454))

	// Test with past expiration
	pastTime := int64(1704279454 - 3600) // current time - 1 hour
	traits.WithExpiration(big.NewInt(pastTime))
	assert.True(t, traits.IsExpired(1704279454))

	// Test at exact expiration time
	exactTime := int64(1704279454)
	traits.WithExpiration(big.NewInt(exactTime))
	assert.False(t, traits.IsExpired(exactTime))  // should not be expired at exact time
	assert.True(t, traits.IsExpired(exactTime+1)) // should be expired one second later
}

func TestMakerTraits_All(t *testing.T) {
	traits, err := DefaultMakerTraits().
		WithAllowedSender(common.BigToAddress(UINT_160_MAX)).
		SetAllowPartialFills(true).
		SetAllowMultipleFills(true).
		WithEpoch(UINT_40_MAX, UINT_40_MAX)
	assert.NoError(t, err)
	traits.WithExpiration(UINT_40_MAX).
		WithExtension().
		SetUnwrapWeth(true).
		SetUsePermit2(true).
		SetUsePermit2(true).
		SetNeedPreInteraction(true).
		SetNeedPostInteraction(true)

	expected := "5f800000000000ffffffffffffffffffffffffffffffffffffffffffffffffff"
	assert.Equal(t, expected, traits.Build().Text(16))
}

func TestIsAllowedSender(t *testing.T) {
	traits := DefaultMakerTraits()
	addr1 := common.HexToAddress("0x1")
	addr2 := common.HexToAddress("0x2")

	assert.True(t, traits.IsAllowedSender(addr1))
	traits.WithAllowedSender(addr2)
	assert.False(t, traits.IsAllowedSender(addr1))
	assert.True(t, traits.IsAllowedSender(addr2))
}

func TestMakerTraitsOptionMarshalUnmarshal(t *testing.T) {
	exp := big.NewInt(1756372966)
	nonce := big.NewInt(473221)
	series := big.NewInt(0)
	allowed := []byte{0xde, 0xad, 0xbe, 0xef}

	orig := MakerTraitsOption{
		AllowPartialFills:     true,
		AllowMultipleFills:    false,
		NeedCheckEpochManager: true,
		UseBitInvalidator:     false,
		NeedPreInteraction:    true,
		NeedPostInteraction:   false,
		UnwrapWeth:            true,
		HasExtension:          false,
		IsPrivate:             true,
		Expiration:            exp,
		NonceOrEpoch:          nonce,
		Series:                series,
		AllowedSender:         allowed,
	}

	// marshal -> JSON bytes
	b, err := orig.Marshal()
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	if !json.Valid(b) {
		t.Fatalf("output is not valid JSON: %s", string(b))
	}
	t.Log(string(b))

	// Act: unmarshal -> struct
	got := MakerTraitsOption{}
	err = got.Unmarshal(b)
	if err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	// Assert: round-trip
	if got.AllowPartialFills != orig.AllowPartialFills ||
		got.AllowMultipleFills != orig.AllowMultipleFills ||
		got.NeedCheckEpochManager != orig.NeedCheckEpochManager ||
		got.UseBitInvalidator != orig.UseBitInvalidator ||
		got.NeedPreInteraction != orig.NeedPreInteraction ||
		got.NeedPostInteraction != orig.NeedPostInteraction ||
		got.UnwrapWeth != orig.UnwrapWeth ||
		got.HasExtension != orig.HasExtension ||
		got.IsPrivate != orig.IsPrivate {
		t.Fatalf("bool fields mismatch after round-trip")
	}
	if (got.Expiration == nil) != (orig.Expiration == nil) ||
		(got.Expiration != nil && got.Expiration.Cmp(orig.Expiration) != 0) {
		t.Fatalf("Expiration mismatch: got %v, orig %v", got.Expiration, orig.Expiration)
	}
	if (got.NonceOrEpoch == nil) != (orig.NonceOrEpoch == nil) ||
		(got.NonceOrEpoch != nil && got.NonceOrEpoch.Cmp(orig.NonceOrEpoch) != 0) {
		t.Fatalf("NonceOrEpoch mismatch")
	}
	if (got.Series == nil) != (orig.Series == nil) ||
		(got.Series != nil && got.Series.Cmp(orig.Series) != 0) {
		t.Fatalf("Series mismatch")
	}
	if !reflect.DeepEqual(got.AllowedSender, orig.AllowedSender) {
		t.Fatalf("AllowedSender mismatch: got %s, orig %s",
			hexutil.Encode(got.AllowedSender), hexutil.Encode(orig.AllowedSender))
	}
}
