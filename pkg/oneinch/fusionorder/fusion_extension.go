package fusionorder

import (
	"errors"
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/utils"
	"github.com/ethereum/go-ethereum/common"
)

type FusionExtension struct {
	Address             common.Address
	AuctionDetails      AuctionDetails
	PostInteractionData SettlementPostInteractionData
	MakerPermit         limitorder.Interaction
}

var (
	ErrSettlementContractMismatch = errors.New("settlement contract mismatch")
)

func NewFusionExtensionFromExtension(extension limitorder.Extension) (FusionExtension, error) {
	settlementContract := utils.AddressFromFirstBytes(extension.MakingAmountData)

	if utils.AddressFromFirstBytes(extension.TakingAmountData) != settlementContract {
		return FusionExtension{},
			fmt.Errorf("taking amount data settlement contract mismatch: %w", ErrSettlementContractMismatch)
	}
	if utils.AddressFromFirstBytes(extension.PostInteraction) != settlementContract {
		return FusionExtension{},
			fmt.Errorf("post interaction settlement contract mismatch: %w", ErrSettlementContractMismatch)
	}

	auctionDetails, err := DecodeAuctionDetails(
		utils.Add0x(extension.MakingAmountData[42:]),
	)
	if err != nil {
		return FusionExtension{}, fmt.Errorf("decode auction details: %w", err)
	}

	postInteractionData, err := DecodeSettlementPostInteractionData(
		utils.Add0x(extension.PostInteraction[42:]),
	)
	if err != nil {
		return FusionExtension{}, fmt.Errorf("decode post interaction data: %w", err)
	}

	var makerPermit limitorder.Interaction
	if extension.HasMakerPermit() {
		makerPermit, err = limitorder.DecodeInteraction(extension.MakerPermit)
		if err != nil {
			return FusionExtension{}, fmt.Errorf("decode maker permit: %w", err)
		}
	}

	return FusionExtension{
		Address:             settlementContract,
		AuctionDetails:      auctionDetails,
		PostInteractionData: postInteractionData,
		MakerPermit:         makerPermit,
	}, nil
}
