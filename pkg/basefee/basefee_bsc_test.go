package basefee_test

import (
	"math/big"
	"testing"

	"github.com/KyberNetwork/tradinglib/pkg/basefee"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

// bscParent builds a BSC-shaped parent header: base fee fixed at zero by BEP-336, with gasUsed
// set relative to the gas target so a caller can pick which EIP-1559 branch would be taken if the
// Parlia short-circuit were missed.
func bscParent(gasUsed uint64) *types.Header {
	return &types.Header{
		Number:   big.NewInt(121_013_739),
		GasLimit: 55_000_000,
		GasUsed:  gasUsed,
		BaseFee:  big.NewInt(0),
	}
}

// TestCalcNextBaseFee_BSCIsAlwaysZero pins BSC's base fee to zero for every gasUsed branch.
//
// The over-target case is the one that regressed: with Ethereum's config the Parlia short-circuit
// is skipped, parentBaseFee*delta truncates below one, and CalcBaseFee returns parentBaseFee+1 —
// one wei instead of zero, and only on blocks that exceed their gas target. Measured against
// bsc-rpc.publicnode.com at block 121013739 (gasUsed 53881911, target 27607471): the chain gave 0,
// the Ethereum config gave 1.
//
// One wei is not a rounding nuisance downstream. A caller that treats a zero base fee as "the gas
// oracle is not warm yet" sees the guard trip or pass depending on how full the previous block
// was, which is worse than failing outright.
func TestCalcNextBaseFee_BSCIsAlwaysZero(t *testing.T) {
	t.Parallel()

	// GasLimit 55M with BSC's elasticity puts the target at 27.5M.
	for _, tc := range []struct {
		name    string
		gasUsed uint64
	}{
		{"over target", 53_881_911},
		{"at target", 27_500_000},
		{"under target", 21_333_840},
		{"empty block", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := basefee.CalcNextBaseFee(basefee.BscChainID, bscParent(tc.gasUsed))
			require.NoError(t, err)
			require.Zerof(t, got.Sign(), "BSC base fee is fixed at zero, got %v", got)
		})
	}
}

// TestCalcNextBaseFee_EthereumStillFollowsEIP1559 guards the other direction: the BSC fix must not
// flatten Ethereum, where an over-target parent raises the next base fee.
func TestCalcNextBaseFee_EthereumStillFollowsEIP1559(t *testing.T) {
	t.Parallel()

	parent := &types.Header{
		Number:   big.NewInt(23_637_674),
		GasLimit: 30_000_000,
		GasUsed:  30_000_000, // twice the 15M target
		BaseFee:  big.NewInt(1_000_000_000),
	}

	got, err := basefee.CalcNextBaseFee(basefee.EthChainID, parent)
	require.NoError(t, err)
	require.Positive(t, got.Cmp(parent.BaseFee), "a full Ethereum block raises the next base fee")
}

// TestCalcNextBaseFee_UnknownChainErrors: returning a nil fee with no error would read as a zero
// base fee, which on some chains is a legitimate value and so cannot signal "unsupported".
func TestCalcNextBaseFee_UnknownChainErrors(t *testing.T) {
	t.Parallel()

	_, err := basefee.CalcNextBaseFee(999_999, bscParent(0))
	require.Error(t, err)
}
