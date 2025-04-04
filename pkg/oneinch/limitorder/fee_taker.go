package limitorder

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const Base10000 = 10000

var (
	ErrInvalidResolverFee   = errors.New("invalid resolver fee")
	ErrInvalidIntegratorFee = errors.New("invalid integrator fee")
)

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

type Fees struct {
	Resolver   ResolverFee
	Integrator IntegratorFee
}

type IWhitelist interface {
	IsWhitelisted(taker common.Address) bool
}

type FeeCalculator struct {
	fees      Fees
	whitelist IWhitelist
}

func NewFeeCalculator(fees Fees, whitelist IWhitelist) FeeCalculator {
	return FeeCalculator{
		fees:      fees,
		whitelist: whitelist,
	}
}

func (c FeeCalculator) GetTakingAmount(taker common.Address, orderTakingAmount *big.Int) *big.Int {
	resolverFee, integratorFee := c.getFeesForTaker(taker)

	takingAmount := big.NewInt(Base10000 + resolverFee + integratorFee)
	takingAmount.Mul(orderTakingAmount, takingAmount)
	takingAmount.Add(takingAmount, big.NewInt(Base10000-1))
	return takingAmount.Div(takingAmount, big.NewInt(Base10000))
}

func (c FeeCalculator) GetMakingAmount(taker common.Address, orderMakingAmount *big.Int) *big.Int {
	resolverFee, integratorFee := c.getFeesForTaker(taker)

	makingAmount := big.NewInt(Base10000)
	makingAmount.Mul(orderMakingAmount, makingAmount)
	return makingAmount.Div(makingAmount, big.NewInt(Base10000+resolverFee+integratorFee))
}

func (c FeeCalculator) getFeesForTaker(taker common.Address) (int64, int64) {
	discountNumerator := uint16(Base10000)
	if c.whitelist.IsWhitelisted(taker) {
		discountNumerator = Base10000 - c.fees.Resolver.WhitelistDiscount
	}

	resolverFee := int64(discountNumerator) * int64(c.fees.Resolver.Fee) / Base10000

	return resolverFee, int64(c.fees.Integrator.Fee)
}
