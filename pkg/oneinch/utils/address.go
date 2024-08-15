package utils

import "github.com/ethereum/go-ethereum/common"

func AddressFromFirstBytes(s string) common.Address {
	const addressLength = 42
	return common.HexToAddress(s[:addressLength])
}
