package limitorder_test

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
)

func TestDecodeTraits(t *testing.T) {
	type args struct {
		MakerTraits         string
		HasError            bool
		ExpectedMakerTraits limitorder.MakerTraitsOption
	}
	testCases := []args{
		{
			MakerTraits: "74438847047041592770632009233243045888",
			HasError:    false,
			ExpectedMakerTraits: limitorder.MakerTraitsOption{
				Expiration: int64(1719942657),
			},
		},
		{
			MakerTraits: "13289623586853032246658166462032491053056",
			HasError:    false,
			ExpectedMakerTraits: limitorder.MakerTraitsOption{
				Expiration: int64(1724667808),
			},
		},
		{
			MakerTraits: "",
			HasError:    true,
		},
	}
	for _, tt := range testCases {
		bb, ok := new(big.Int).SetString(tt.MakerTraits, 10)
		if !ok {
			if tt.HasError {
				continue
			}
			t.Error("big.Int SetString failed")
		}
		got, err := limitorder.DecodeMarkerTraits(bb)
		if tt.HasError {
			if err == nil {
				t.Errorf("DecodeTraits() error = %v, wantErr %v", err, tt.HasError)
				return
			}
		}
		if got.Expiration != tt.ExpectedMakerTraits.Expiration {
			t.Errorf("DecodeTraits() = %v, want %v", got.Expiration, tt.ExpectedMakerTraits.Expiration)
		}
	}
}
