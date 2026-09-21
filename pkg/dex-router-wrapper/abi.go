package dexrouterwrapper

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	// WrapperABI is the parsed ABI of DexRouterWrapper.sol.
	WrapperABI abi.ABI

	wrapperCreationBytecode []byte
	wrapperDeployedBytecode []byte
)

//nolint:gochecknoinits
func init() {
	var err error
	WrapperABI, err = abi.JSON(bytes.NewReader(wrapperABIJSON))
	if err != nil {
		panic(err)
	}

	wrapperCreationBytecode, err = hex.DecodeString(strings.TrimSpace(wrapperCreationBytecodeHex))
	if err != nil {
		panic(err)
	}

	wrapperDeployedBytecode, err = hex.DecodeString(strings.TrimSpace(wrapperDeployedBytecodeHex))
	if err != nil {
		panic(err)
	}
}
