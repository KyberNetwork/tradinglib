package fusionextention

import (
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"github.com/KyberNetwork/tradinglib/pkg/oneinch/util"
	"github.com/ethereum/go-ethereum/common"
)

type WhitelistItem struct {
	AddressHalf util.AddressHalf
	Delay       int64
}

type Whitelist struct {
	ResolvingStartTime int64
	Whitelist          []WhitelistItem
}

func DecodeWhitelist(iter *decode.BytesIterator) (Whitelist, error) {
	resolvingStartTime, err := iter.NextUint32()
	if err != nil {
		return Whitelist{}, fmt.Errorf("get resolving start time: %w", err)
	}

	size, err := iter.NextUint8()
	if err != nil {
		return Whitelist{}, fmt.Errorf("get resolving start time: %w", err)
	}

	whitelist := make([]WhitelistItem, 0, size)
	for range size {
		addressHalfBytes, err := iter.NextBytes(util.AddressHalfLength)
		if err != nil {
			return Whitelist{}, fmt.Errorf("get whitelist item address half: %w", err)
		}

		delay, err := iter.NextUint16()
		if err != nil {
			return Whitelist{}, fmt.Errorf("get whitelist item delay: %w", err)
		}

		whitelist = append(whitelist, WhitelistItem{
			AddressHalf: util.BytesToAddressHalf(addressHalfBytes),
			Delay:       int64(delay),
		})
	}

	return Whitelist{
		ResolvingStartTime: int64(resolvingStartTime),
		Whitelist:          whitelist,
	}, nil
}

func (wl Whitelist) Length() int {
	return len(wl.Whitelist)
}

func (wl Whitelist) CanExecuteAt(executor common.Address, executionTime int64) bool {
	addressHalf := util.HalfAddressFromAddress(executor)
	allowedFrom := wl.ResolvingStartTime

	for _, item := range wl.Whitelist {
		allowedFrom += item.Delay

		if addressHalf == item.AddressHalf {
			return executionTime >= allowedFrom
		} else if executionTime < allowedFrom {
			return false
		}
	}

	return false
}

func (wl Whitelist) CanAnySolverExecuteAt(executionTime int64) bool {
	allowedFrom := wl.ResolvingStartTime
	for _, item := range wl.Whitelist {
		allowedFrom += item.Delay
		if executionTime >= allowedFrom {
			return true
		}
	}
	return false
}

func (wl Whitelist) IsExclusivityPeriod(ts int64) bool {
	size := wl.Length()
	if size == 0 {
		return false
	}

	if size == 1 {
		return true
	}

	if wl.Whitelist[0].Delay == wl.Whitelist[1].Delay {
		return false
	}

	return ts <= wl.ResolvingStartTime+wl.Whitelist[1].Delay
}

func (wl Whitelist) IsExclusiveResolver(resolver common.Address) bool {
	size := wl.Length()
	if size == 0 {
		return false
	}

	addressHalf := util.HalfAddressFromAddress(resolver)
	if size == 1 {
		return addressHalf == wl.Whitelist[0].AddressHalf
	}

	if wl.Whitelist[0].Delay == wl.Whitelist[1].Delay {
		return false
	}

	return addressHalf == wl.Whitelist[0].AddressHalf
}

func (wl Whitelist) IsWhitelisted(taker common.Address) bool {
	addressHalf := util.HalfAddressFromAddress(taker)
	for _, item := range wl.Whitelist {
		if addressHalf == item.AddressHalf {
			return true
		}
	}

	return false
}
