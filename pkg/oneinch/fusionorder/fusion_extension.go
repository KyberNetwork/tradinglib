package fusionorder

import (
	"errors"
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/common"
)

type FusionExtension struct {
	Address             common.Address
	AuctionDetails      AuctionDetails
	PostInteractionData SettlementPostInteractionData
	MakerPermit         limitorder.Interaction
}

var ErrSettlementContractMismatch = errors.New("settlement contract mismatch")

func NewFusionExtensionFromExtension(extension limitorder.Extension) (FusionExtension, error) {
	settlementContract := AddressFromFirstBytes(extension.MakingAmountData)

	if AddressFromFirstBytes(extension.TakingAmountData) != settlementContract {
		return FusionExtension{},
			fmt.Errorf("taking amount data settlement contract mismatch: %w", ErrSettlementContractMismatch)
	}
	if AddressFromFirstBytes(extension.PostInteraction) != settlementContract {
		return FusionExtension{},
			fmt.Errorf("post interaction settlement contract mismatch: %w", ErrSettlementContractMismatch)
	}

	auctionDetails, err := DecodeAuctionDetails(extension.MakingAmountData[common.AddressLength:])
	if err != nil {
		return FusionExtension{}, fmt.Errorf("decode auction details: %w", err)
	}

	postInteractionData, err := DecodeSettlementPostInteractionData(
		extension.PostInteraction[common.AddressLength:],
	)
	if err != nil {
		return FusionExtension{}, fmt.Errorf("decode post interaction data: %w", err)
	}

	var makerPermit limitorder.Interaction
	if extension.HasMakerPermit() {
		makerPermit = limitorder.DecodeInteraction(extension.MakerPermit)
	}

	return FusionExtension{
		Address:             settlementContract,
		AuctionDetails:      auctionDetails,
		PostInteractionData: postInteractionData,
		MakerPermit:         makerPermit,
	}, nil
}
