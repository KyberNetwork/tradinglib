package types

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

var ErrInvalidBloxrouteBackrunConfig = errors.New("invalid bloxroute backrun config")

// BloxrouteBackrunConfig is the per-tx BackRunMe profit-split config from bloxroute's
// arbOnlyMEV stream. On the wire from bloxroute it arrives as a 5-element string array
// [contractSplit, targetRewardAddress, targetSplit, blxrRewardAddress, searcherSplit];
// mempool-explorer parses it once via ParseBloxrouteBackrunConfig and attaches the typed
// value to Message.BloxrouteBackrunConfig, so downstream consumers read it directly (nil
// when the tx is not a bloxroute backrun).
type BloxrouteBackrunConfig struct {
	ContractSplit       float64 `json:"contract_split"`
	TargetRewardAddress string  `json:"target_reward_address"`
	TargetSplit         float64 `json:"target_split"`
	BlxrRewardAddress   string  `json:"blxr_reward_address"`
	SearcherSplit       float64 `json:"searcher_split"`
}

// ParseBloxrouteBackrunConfig parses the 5-element backrunConfig array from the arbOnlyMEV
// stream into a validated BloxrouteBackrunConfig. Element order is
// [contractSplit, targetRewardAddress, targetSplit, blxrRewardAddress, searcherSplit].
func ParseBloxrouteBackrunConfig(parts []string) (*BloxrouteBackrunConfig, error) {
	if len(parts) != 5 {
		return nil, fmt.Errorf("%w: expected 5 elements, got %d", ErrInvalidBloxrouteBackrunConfig, len(parts))
	}

	contractSplit, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil, fmt.Errorf("%w: contract split: %w", ErrInvalidBloxrouteBackrunConfig, err)
	}
	targetSplit, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return nil, fmt.Errorf("%w: target split: %w", ErrInvalidBloxrouteBackrunConfig, err)
	}
	searcherSplit, err := strconv.ParseFloat(parts[4], 64)
	if err != nil {
		return nil, fmt.Errorf("%w: searcher split: %w", ErrInvalidBloxrouteBackrunConfig, err)
	}

	cfg := BloxrouteBackrunConfig{
		ContractSplit:       contractSplit,
		TargetRewardAddress: parts[1],
		TargetSplit:         targetSplit,
		BlxrRewardAddress:   parts[3],
		SearcherSplit:       searcherSplit,
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Validate checks the split percentages are non-negative, sum to <= 100, and the reward
// addresses are valid non-zero addresses.
func (c BloxrouteBackrunConfig) Validate() error {
	if c.ContractSplit < 0 || c.TargetSplit < 0 || c.SearcherSplit < 0 {
		return fmt.Errorf("%w: splits must be non-negative", ErrInvalidBloxrouteBackrunConfig)
	}
	if _, err := c.BlxrSplit(); err != nil {
		return err
	}
	if !isValidNonZeroAddress(c.TargetRewardAddress) {
		return fmt.Errorf("%w: invalid target reward address", ErrInvalidBloxrouteBackrunConfig)
	}
	if !isValidNonZeroAddress(c.BlxrRewardAddress) {
		return fmt.Errorf("%w: invalid blxr reward address", ErrInvalidBloxrouteBackrunConfig)
	}
	return nil
}

// BlxrSplit returns bloxroute's reward share: 100 - contract - target - searcher.
func (c BloxrouteBackrunConfig) BlxrSplit() (float64, error) {
	blxrSplit := 100 - c.ContractSplit - c.TargetSplit - c.SearcherSplit
	if blxrSplit < 0 {
		return 0, fmt.Errorf("%w: split sum exceeds 100", ErrInvalidBloxrouteBackrunConfig)
	}
	return blxrSplit, nil
}

func isValidNonZeroAddress(addr string) bool {
	return common.IsHexAddress(addr) && common.HexToAddress(addr) != (common.Address{})
}
