package encode

import "github.com/ethereum/go-ethereum/accounts/abi"

//nolint:gochecknoglobals
var (
	Uint256, _ = abi.NewType("uint256", "", nil)
	Address, _ = abi.NewType("address", "", nil)
	Bytes, _   = abi.NewType("bytes", "", nil)
)
