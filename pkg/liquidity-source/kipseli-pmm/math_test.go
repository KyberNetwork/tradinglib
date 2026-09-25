//nolint:testpackage
package kipselipmm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_removeFirstLevelsByAmount(t *testing.T) {
	type args struct {
		priceLevels []PriceLevel
		amount      float64
	}
	tests := []struct {
		name string
		args args
		want []PriceLevel
	}{
		{
			name: "remove nothing",
			args: args{
				priceLevels: []PriceLevel{
					{Rate: 2, Quantity: 10},
					{Rate: 1, Quantity: 20},
				},
				amount: 0,
			},
			want: []PriceLevel{
				{Rate: 2, Quantity: 10},
				{Rate: 1, Quantity: 20},
			},
		},
		{
			name: "remove some quantity from first level",
			args: args{
				priceLevels: []PriceLevel{
					{Rate: 2, Quantity: 10},
					{Rate: 1, Quantity: 20},
				},
				amount: 5,
			},
			want: []PriceLevel{
				{Rate: 2, Quantity: 5},
				{Rate: 1, Quantity: 20},
			},
		},
		{
			name: "remove first level completely",
			args: args{
				priceLevels: []PriceLevel{
					{Rate: 2, Quantity: 10},
					{Rate: 1, Quantity: 20},
				},
				amount: 10,
			},
			want: []PriceLevel{
				{Rate: 1, Quantity: 20},
			},
		},
		{
			name: "remove all levels",
			args: args{
				priceLevels: []PriceLevel{
					{Rate: 2, Quantity: 10},
					{Rate: 1, Quantity: 20},
				},
				amount: 40,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeFirstLevelsByAmount(tt.args.priceLevels, tt.args.amount)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

// nolint: dupl
func Test_calcQuoteAmount(t *testing.T) {
	type args struct {
		priceLevels []PriceLevel
		baseAmount  float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr error
	}{
		{
			name: "normal case",
			args: args{
				priceLevels: []PriceLevel{
					{Rate: 3, Quantity: 10},
					{Rate: 2, Quantity: 20},
				},
				baseAmount: 20,
			},
			want: 50, // 10*3 + 10*2
		},
		{
			name: "empty price levels",
			args: args{
				priceLevels: []PriceLevel{},
				baseAmount:  10,
			},
			wantErr: ErrEmptyPriceLevels,
		},
		{
			name: "insufficient liquidity",
			args: args{
				priceLevels: []PriceLevel{
					{Rate: 3, Quantity: 5},
				},
				baseAmount: 10,
			},
			wantErr: ErrInsufficientLiquidity,
		},
		{
			name: "exact liquidity",
			args: args{
				priceLevels: []PriceLevel{
					{Rate: 3, Quantity: 5},
					{Rate: 2, Quantity: 5},
				},
				baseAmount: 10,
			},
			want: 25, // 5*3 + 5*2
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcQuoteAmount(tt.args.priceLevels, tt.args.baseAmount)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)
			assert.InEpsilon(t, tt.want, got, 10e-18)
		})
	}
}

// nolint: dupl
func Test_calcBaseAmount(t *testing.T) {
	type args struct {
		priceLevels []PriceLevel
		quoteAmount float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr error
	}{
		{
			name: "normal case",
			args: args{
				priceLevels: []PriceLevel{
					{Rate: 2, Quantity: 10}, // can provide 20 quote
					{Rate: 1, Quantity: 20}, // can provide 20 quote
				},
				quoteAmount: 30,
			},
			want: 20, // 10 + 10
		},
		{
			name: "empty price levels",
			args: args{
				priceLevels: []PriceLevel{},
				quoteAmount: 10,
			},
			wantErr: ErrEmptyPriceLevels,
		},
		{
			name: "insufficient liquidity",
			args: args{
				priceLevels: []PriceLevel{
					{Rate: 2, Quantity: 5}, // can provide 10 quote
				},
				quoteAmount: 20,
			},
			wantErr: ErrInsufficientLiquidity,
		},
		{
			name: "exact liquidity",
			args: args{
				priceLevels: []PriceLevel{
					{Rate: 2, Quantity: 5}, // can provide 10 quote
					{Rate: 1, Quantity: 5}, // can provide 5 quote
				},
				quoteAmount: 15,
			},
			want: 10, // 5 + 5
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcBaseAmount(tt.args.priceLevels, tt.args.quoteAmount)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)
			assert.InEpsilon(t, tt.want, got, 10e-18)
		})
	}
}
