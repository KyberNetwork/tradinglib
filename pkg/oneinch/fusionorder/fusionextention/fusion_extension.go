package fusionextention

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/fusionorder/auctiondetail"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
)

var ErrInvalidExtension = errors.New("invalid extension")

type Extra struct {
	MakerPermit    limitorder.Interaction
	CustomReceiver common.Address
	Fees           limitorder.Fees
}

type FusionExtension struct {
	Address        common.Address
	AuctionDetails auctiondetail.AuctionDetails
	Whitelist      Whitelist
	Extra          Extra
	SurplusParam   SurplusParam
}

// public static fromExtension(extension: Extension): FusionExtension
//
//nolint:funlen,cyclop
func NewFusionExtensionFromExtension(extension limitorder.Extension) (FusionExtension, error) {
	settlementContract := util.AddressFromFirstBytes(extension.MakingAmountData)

	if util.AddressFromFirstBytes(extension.TakingAmountData) != settlementContract {
		return FusionExtension{},
			fmt.Errorf("%w: taking amount data settlement contract mismatch", ErrInvalidExtension)
	}
	if util.AddressFromFirstBytes(extension.PostInteraction) != settlementContract {
		return FusionExtension{},
			fmt.Errorf("%w: post interaction settlement contract mismatch", ErrInvalidExtension)
	}
	if !bytes.Equal(extension.TakingAmountData, extension.MakingAmountData) {
		return FusionExtension{},
			fmt.Errorf("%w: takingAmountData and makingAmountData not match", ErrInvalidExtension)
	}

	postInteractionData, err := DecodeSettlementPostInteractionData(extension.PostInteraction)
	if err != nil {
		return FusionExtension{}, fmt.Errorf("decode post interaction data: %w", err)
	}

	amountIter := decode.NewBytesIterator(extension.MakingAmountData)
	auctionDetails, err := auctiondetail.DecodeAuctionDetails(amountIter)
	if err != nil {
		return FusionExtension{}, fmt.Errorf("decode auctioncalculator details: %w", err)
	}

	amountData, err := ParseAmountData(amountIter)
	if err != nil {
		return FusionExtension{}, fmt.Errorf("decode amount data: %w", err)
	}

	whitelistLength, err := amountIter.NextUint8()
	if err != nil {
		return FusionExtension{}, fmt.Errorf("get whitelist address length: %w", err)
	}
	if int(whitelistLength) != postInteractionData.Whitelist.Length() {
		return FusionExtension{}, fmt.Errorf("%w: whitelist length not match", ErrInvalidExtension)
	}

	whitelistAddressesFromAmount := make([]util.AddressHalf, 0, whitelistLength)
	for range whitelistLength {
		addressHalfBytes, err := amountIter.NextBytes(util.AddressHalfLength)
		if err != nil {
			return FusionExtension{}, fmt.Errorf("decode whitelist address from amount: %w", err)
		}
		whitelistAddressesFromAmount = append(
			whitelistAddressesFromAmount, util.BytesToAddressHalf(addressHalfBytes))
	}

	var makerPermit limitorder.Interaction
	if extension.HasMakerPermit() {
		makerPermit, err = limitorder.DecodeInteraction(extension.MakerPermit)
		if err != nil {
			return FusionExtension{}, fmt.Errorf("decode maker permit: %w", err)
		}
	}

	if amountData.IntegratorFee != postInteractionData.InteractionData.IntegratorFee {
		return FusionExtension{}, fmt.Errorf("%w: integrator fee not match", ErrInvalidExtension)
	}
	if amountData.ResolverFee != postInteractionData.InteractionData.ResolverFee {
		return FusionExtension{}, fmt.Errorf("%w: resolver fee", ErrInvalidExtension)
	}
	if amountData.WhitelistDiscount != postInteractionData.InteractionData.WhitelistDiscount {
		return FusionExtension{}, fmt.Errorf("%w: whitelist discount not match", ErrInvalidExtension)
	}
	if amountData.IntegratorShare != postInteractionData.InteractionData.IntegratorShare {
		return FusionExtension{}, fmt.Errorf("%w: integrator share not match", ErrInvalidExtension)
	}
	for i, item := range postInteractionData.Whitelist.Whitelist {
		if item.AddressHalf != whitelistAddressesFromAmount[i] {
			return FusionExtension{}, fmt.Errorf("%w: whitelist address not match", ErrInvalidExtension)
		}
	}

	fusionExtension := FusionExtension{
		Address:        settlementContract,
		AuctionDetails: auctionDetails,
		Whitelist:      postInteractionData.Whitelist,
		Extra: Extra{
			MakerPermit:    makerPermit,
			CustomReceiver: postInteractionData.CustomReceiver,
		},
		SurplusParam: postInteractionData.SurplusParam,
	}
	if !postInteractionData.HasFees() {
		return fusionExtension, nil
	}

	var integratorFee limitorder.IntegratorFee
	if postInteractionData.InteractionData.IntegratorFee != 0 {
		integratorFee, err = limitorder.NewIntegratorFee(
			postInteractionData.IntegratorFeeRecipient,
			postInteractionData.ProtocolFeeRecipient,
			postInteractionData.InteractionData.IntegratorFee,
			postInteractionData.InteractionData.IntegratorShare,
		)
		if err != nil {
			return FusionExtension{}, err
		}
	}

	var resolverFee limitorder.ResolverFee
	if postInteractionData.InteractionData.ResolverFee != 0 {
		resolverFee, err = limitorder.NewResolverFee(
			postInteractionData.ProtocolFeeRecipient,
			postInteractionData.InteractionData.ResolverFee,
			postInteractionData.InteractionData.WhitelistDiscount,
		)
		if err != nil {
			return FusionExtension{}, err
		}
	}

	fusionExtension.Extra.Fees = limitorder.Fees{
		Integrator: integratorFee,
		Resolver:   resolverFee,
	}

	return fusionExtension, nil
}
