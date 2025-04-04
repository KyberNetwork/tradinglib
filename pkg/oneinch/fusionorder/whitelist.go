package fusionorder

import (
	"fmt"

	"github.com/KyberNetwork/tradinglib/pkg/oneinch/decode"
	"github.com/ethereum/go-ethereum/common"
)

type WhitelistItem struct {
	AddressHalf AddressHalf
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
		addressHalfBytes, err := iter.NextBytes(addressHalfLength)
		if err != nil {
			return Whitelist{}, fmt.Errorf("get whitelist item address half: %w", err)
		}

		delay, err := iter.NextUint16()
		if err != nil {
			return Whitelist{}, fmt.Errorf("get whitelist item delay: %w", err)
		}

		whitelist = append(whitelist, WhitelistItem{
			AddressHalf: BytesToAddressHalf(addressHalfBytes),
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
	addressHalf := HalfAddressFromAddress(executor)
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

	addressHalf := HalfAddressFromAddress(resolver)
	if size == 1 {
		return addressHalf == wl.Whitelist[0].AddressHalf
	}

	if wl.Whitelist[0].Delay == wl.Whitelist[1].Delay {
		return false
	}

	return addressHalf == wl.Whitelist[0].AddressHalf
}
