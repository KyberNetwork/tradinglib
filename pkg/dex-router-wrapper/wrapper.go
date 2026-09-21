// Package dexrouterwrapper encodes calls to, and provides the deployed
// bytecode of, DexRouterWrapper.sol — a contract that forwards an arbitrary
// call to a DEX router and reports the resulting output-token amount and gas
// used. It is designed to be run via an eth_call/eth_estimateGas state
// override rather than deployed on-chain: see StateOverride.
package dexrouterwrapper

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

const methodWrap = "wrap"

// NativeTokenAddress is the sentinel address meaning "the native chain
// asset" wherever an ERC20 token address is otherwise expected, matching
// DexRouterWrapper.sol's NATIVE constant.
var NativeTokenAddress = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

// EncodeWrapCalldata packs a call to DexRouterWrapper.wrap(target, data,
// outputToken, recipient). If target's call needs native ETH as input,
// set the Value field on the caller's ethereum.CallMsg — it is forwarded by
// the wrapper as msg.value, not encoded in this calldata.
func EncodeWrapCalldata(target common.Address, data []byte, outputToken, recipient common.Address) ([]byte, error) {
	packed, err := WrapperABI.Pack(methodWrap, target, data, outputToken, recipient)
	if err != nil {
		return nil, fmt.Errorf("pack wrap calldata: %w", err)
	}

	return packed, nil
}

// DecodeWrapOutput unpacks the (returnAmount, gasUsed) return values of a
// DexRouterWrapper.wrap call.
func DecodeWrapOutput(data []byte) (*big.Int, uint64, error) {
	var res [2]any
	if err := WrapperABI.UnpackIntoInterface(&res, methodWrap, data); err != nil {
		return nil, 0, fmt.Errorf("unpack wrap output: %w", err)
	}

	returnAmount := *abi.ConvertType(res[0], new(*big.Int)).(**big.Int)       //nolint:forcetypeassert
	gasUsed := (*abi.ConvertType(res[1], new(*big.Int)).(**big.Int)).Uint64() //nolint:forcetypeassert

	return returnAmount, gasUsed, nil
}

// DeployedBytecode returns DexRouterWrapper's runtime (deployed) bytecode —
// the code that runs when the contract is called, as opposed to the
// constructor code that produces it. This is what belongs in a state
// override's Code field; see StateOverride.
func DeployedBytecode() []byte {
	return append([]byte(nil), wrapperDeployedBytecode...)
}

// CreationBytecode returns DexRouterWrapper's constructor (creation)
// bytecode, for actually deploying the contract — e.g. in tests against an
// in-process EVM.
func CreationBytecode() []byte {
	return append([]byte(nil), wrapperCreationBytecode...)
}

// StateOverride builds the eth_call/eth_estimateGas state override that
// makes overrideAddr behave as DexRouterWrapper for the duration of one
// simulated call, without any real deployment.
func StateOverride(overrideAddr common.Address) map[common.Address]gethclient.OverrideAccount {
	return map[common.Address]gethclient.OverrideAccount{
		overrideAddr: {Code: DeployedBytecode()},
	}
}
