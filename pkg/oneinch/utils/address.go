package utils

import "github.com/ethereum/go-ethereum/common"

func AddressFromFirstBytes(s []byte) common.Address {
	return common.BytesToAddress(s[:common.AddressLength])
}
