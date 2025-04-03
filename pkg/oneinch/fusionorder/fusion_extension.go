package fusionorder

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/limitorder"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrInvalidExtension     = errors.New("invalid extension")
	ErrInvalidIntegratorFee = errors.New("invalid integrator fee")
	ErrInvalidResolverFee   = errors.New("invalid resolver fee")
)

type IntegratorFee struct {
	Integrator common.Address
	Protocol   common.Address
	Fee        uint16 // in bps
	Share      uint16 // in bps
}

func NewIntegratorFee(
	integrator common.Address,
	protocol common.Address,
	fee uint16,
	share uint16,
) (IntegratorFee, error) {
	if fee == 0 {
		if share != 0 {
			return IntegratorFee{}, fmt.Errorf(
				"%w: share must be zero if fee is zero", ErrInvalidIntegratorFee)
		}
		if integrator != (common.Address{}) {
			return IntegratorFee{}, fmt.Errorf(
				"%w: integrator address must be zero if fee is zero", ErrInvalidIntegratorFee)
		}
		if protocol != (common.Address{}) {
			return IntegratorFee{}, fmt.Errorf(
				"%w: protocol address must be zero if fee is zero", ErrInvalidIntegratorFee)
		}
	}

	if (integrator == (common.Address{}) || protocol == (common.Address{})) && fee != 0 {
		return IntegratorFee{}, fmt.Errorf(
			"%w: fee must be zero if integrator or protocol is zero address", ErrInvalidIntegratorFee)
	}

	return IntegratorFee{
		Integrator: integrator,
		Protocol:   protocol,
		Fee:        fee,
		Share:      share,
	}, nil
}

type ResolverFee struct {
	Receiver          common.Address
	Fee               uint16 // in bps
	WhitelistDiscount uint16 // in bps
}

func NewResolverFee(receiver common.Address, fee uint16, whitelistDiscount uint16) (ResolverFee, error) {
	if receiver == (common.Address{}) && fee != 0 {
		return ResolverFee{}, fmt.Errorf(
			"%w: fee must be zero if receiver is zero address", ErrInvalidResolverFee)
	}
	if receiver != (common.Address{}) && fee == 0 {
		return ResolverFee{}, fmt.Errorf(
			"%w: receiver must be zero address if fee is zero", ErrInvalidResolverFee)
	}
	if fee == 0 && whitelistDiscount != 0 {
		return ResolverFee{}, fmt.Errorf(
			"%w: whitelist discount must be zero if fee is zero", ErrInvalidResolverFee)
	}

	return ResolverFee{
		Receiver:          receiver,
		Fee:               fee,
		WhitelistDiscount: whitelistDiscount,
	}, nil
}

type Fees struct {
	ResolverFee   ResolverFee
	IntegratorFee IntegratorFee
}

type Extra struct {
	MakerPermit    limitorder.Interaction
	CustomReceiver common.Address
	Fees           Fees
}

type FusionExtension struct {
	Address        common.Address
	AuctionDetails AuctionDetails
	Whitelist      Whitelist
	Extra          Extra
}

//nolint:funlen,cyclop
func NewFusionExtensionFromExtension(extension limitorder.Extension) (FusionExtension, error) {
	settlementContract := AddressFromFirstBytes(extension.MakingAmountData)

	if AddressFromFirstBytes(extension.TakingAmountData) != settlementContract {
		return FusionExtension{},
			fmt.Errorf("%w: taking amount data settlement contract mismatch", ErrInvalidExtension)
	}
	if AddressFromFirstBytes(extension.PostInteraction) != settlementContract {
		return FusionExtension{},
			fmt.Errorf("%w: post interaction settlement contract mismatch", ErrInvalidExtension)
	}
	if !bytes.Equal(extension.TakingAmountData, extension.MakingAmountData) {
		return FusionExtension{},
			fmt.Errorf("%w: takingAmountData and makingAmountData not match", ErrInvalidExtension)
	}

	postInteractionData, err := DecodeSettlementPostInteractionData(
		extension.PostInteraction[common.AddressLength:],
	)
	if err != nil {
		return FusionExtension{}, fmt.Errorf("decode post interaction data: %w", err)
	}

	amountIter := decode.NewBytesIterator(extension.MakingAmountData[common.AddressLength:])
	auctionDetails, err := DecodeAuctionDetails(amountIter)
	if err != nil {
		return FusionExtension{}, fmt.Errorf("decode auction details: %w", err)
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

	whitelistAddressesFromAmount := make([]AddressHalf, 0, whitelistLength)
	for range whitelistLength {
		addressHalfBytes, err := amountIter.NextBytes(addressHalfLength)
		if err != nil {
			return FusionExtension{}, fmt.Errorf("decode whitelist address from amount: %w", err)
		}
		whitelistAddressesFromAmount = append(
			whitelistAddressesFromAmount, BytesToAddressHalf(addressHalfBytes))
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
	}
	if postInteractionData.IntegratorFeeRecipient == (common.Address{}) &&
		postInteractionData.ProtocolFeeRecipient == (common.Address{}) {
		return fusionExtension, nil
	}

	var integratorFee IntegratorFee
	if postInteractionData.InteractionData.IntegratorFee != 0 {
		integratorFee, err = NewIntegratorFee(
			postInteractionData.IntegratorFeeRecipient,
			postInteractionData.ProtocolFeeRecipient,
			postInteractionData.InteractionData.IntegratorFee,
			postInteractionData.InteractionData.IntegratorShare,
		)
		if err != nil {
			return FusionExtension{}, err
		}
	}

	var resolverFee ResolverFee
	if postInteractionData.InteractionData.ResolverFee != 0 {
		resolverFee, err = NewResolverFee(
			postInteractionData.ProtocolFeeRecipient,
			postInteractionData.InteractionData.ResolverFee,
			postInteractionData.InteractionData.WhitelistDiscount,
		)
		if err != nil {
			return FusionExtension{}, err
		}
	}

	fusionExtension.Extra.Fees = Fees{
		IntegratorFee: integratorFee,
		ResolverFee:   resolverFee,
	}

	return fusionExtension, nil
}
