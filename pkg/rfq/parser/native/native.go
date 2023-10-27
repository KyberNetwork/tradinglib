// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package native

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// INativePoolNewPoolConfig is an auto generated low-level Go binding around an user-defined struct.
type INativePoolNewPoolConfig struct {
	TreasuryAddress    common.Address
	PoolOwnerAddress   common.Address
	SignerAddress      common.Address
	RouterAddress      common.Address
	IsPublicTreasury   bool
	IsTreasuryContract bool
	Fees               []*big.Int
	TokenAs            []common.Address
	TokenBs            []common.Address
	PricingModelIds    []*big.Int
}

// NativeMetaData contains all meta data concerning the Native contract.
var NativeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"onlyOwnerOrPauserOrPoolFactoryCanCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"AddSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBlacklister\",\"type\":\"address\"}],\"name\":\"BlacklisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"RemovePair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"RemoveSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"SetRouter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"}],\"name\":\"SetTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasuryOwner\",\"type\":\"address\"}],\"name\":\"SetTreasuryOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amountIn\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"amountOut\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"quoteId\",\"type\":\"bytes16\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"UnBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeOld\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeNew\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricingModelIdOld\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricingModelIdNew\",\"type\":\"uint256\"}],\"name\":\"UpdatePair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CONSTANT_SUM_PRICE_MODEL_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FIXED_PRICE_MODEL_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PMM_PRICE_MODEL_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_V2_PRICE_MODEL_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"addSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blacklister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getPairFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getPairPricingModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pricingModelId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricingModelRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenBs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"treasuryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isPublicTreasury\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isTreasuryContract\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenAs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenBs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pricingModelIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structINativePool.NewPoolConfig\",\"name\":\"poolConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_pricingModelRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOnChainPricing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPmm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPublicTreasury\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTreasuryContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nonceMapping\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"pairExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricingModelRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"removingIdx\",\"type\":\"uint256\"}],\"name\":\"removePair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"removeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"setPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flexibleAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callback\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenAs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenBs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"unBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newBlacklister\",\"type\":\"address\"}],\"name\":\"updateBlacklister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokenAs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokenBs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_pricingModelIds\",\"type\":\"uint256[]\"}],\"name\":\"updatePairs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NativeABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeMetaData.ABI instead.
var NativeABI = NativeMetaData.ABI

// Native is an auto generated Go binding around an Ethereum contract.
type Native struct {
	NativeCaller     // Read-only binding to the contract
	NativeTransactor // Write-only binding to the contract
	NativeFilterer   // Log filterer for contract events
}

// NativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeSession struct {
	Contract     *Native           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeCallerSession struct {
	Contract *NativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTransactorSession struct {
	Contract     *NativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeRaw struct {
	Contract *Native // Generic contract binding to access the raw methods on
}

// NativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeCallerRaw struct {
	Contract *NativeCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTransactorRaw struct {
	Contract *NativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNative creates a new instance of Native, bound to a specific deployed contract.
func NewNative(address common.Address, backend bind.ContractBackend) (*Native, error) {
	contract, err := bindNative(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Native{NativeCaller: NativeCaller{contract: contract}, NativeTransactor: NativeTransactor{contract: contract}, NativeFilterer: NativeFilterer{contract: contract}}, nil
}

// NewNativeCaller creates a new read-only instance of Native, bound to a specific deployed contract.
func NewNativeCaller(address common.Address, caller bind.ContractCaller) (*NativeCaller, error) {
	contract, err := bindNative(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeCaller{contract: contract}, nil
}

// NewNativeTransactor creates a new write-only instance of Native, bound to a specific deployed contract.
func NewNativeTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTransactor, error) {
	contract, err := bindNative(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTransactor{contract: contract}, nil
}

// NewNativeFilterer creates a new log filterer instance of Native, bound to a specific deployed contract.
func NewNativeFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeFilterer, error) {
	contract, err := bindNative(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeFilterer{contract: contract}, nil
}

// bindNative binds a generic wrapper to an already deployed contract.
func bindNative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Native *NativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Native.Contract.NativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Native *NativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Native.Contract.NativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Native *NativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Native.Contract.NativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Native *NativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Native.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Native *NativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Native.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Native *NativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Native.Contract.contract.Transact(opts, method, params...)
}

// CONSTANTSUMPRICEMODELID is a free data retrieval call binding the contract method 0x2e5d509b.
//
// Solidity: function CONSTANT_SUM_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeCaller) CONSTANTSUMPRICEMODELID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "CONSTANT_SUM_PRICE_MODEL_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONSTANTSUMPRICEMODELID is a free data retrieval call binding the contract method 0x2e5d509b.
//
// Solidity: function CONSTANT_SUM_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeSession) CONSTANTSUMPRICEMODELID() (*big.Int, error) {
	return _Native.Contract.CONSTANTSUMPRICEMODELID(&_Native.CallOpts)
}

// CONSTANTSUMPRICEMODELID is a free data retrieval call binding the contract method 0x2e5d509b.
//
// Solidity: function CONSTANT_SUM_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeCallerSession) CONSTANTSUMPRICEMODELID() (*big.Int, error) {
	return _Native.Contract.CONSTANTSUMPRICEMODELID(&_Native.CallOpts)
}

// FIXEDPRICEMODELID is a free data retrieval call binding the contract method 0x8d484b24.
//
// Solidity: function FIXED_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeCaller) FIXEDPRICEMODELID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "FIXED_PRICE_MODEL_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FIXEDPRICEMODELID is a free data retrieval call binding the contract method 0x8d484b24.
//
// Solidity: function FIXED_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeSession) FIXEDPRICEMODELID() (*big.Int, error) {
	return _Native.Contract.FIXEDPRICEMODELID(&_Native.CallOpts)
}

// FIXEDPRICEMODELID is a free data retrieval call binding the contract method 0x8d484b24.
//
// Solidity: function FIXED_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeCallerSession) FIXEDPRICEMODELID() (*big.Int, error) {
	return _Native.Contract.FIXEDPRICEMODELID(&_Native.CallOpts)
}

// PMMPRICEMODELID is a free data retrieval call binding the contract method 0xe14da7ea.
//
// Solidity: function PMM_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeCaller) PMMPRICEMODELID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "PMM_PRICE_MODEL_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PMMPRICEMODELID is a free data retrieval call binding the contract method 0xe14da7ea.
//
// Solidity: function PMM_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeSession) PMMPRICEMODELID() (*big.Int, error) {
	return _Native.Contract.PMMPRICEMODELID(&_Native.CallOpts)
}

// PMMPRICEMODELID is a free data retrieval call binding the contract method 0xe14da7ea.
//
// Solidity: function PMM_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeCallerSession) PMMPRICEMODELID() (*big.Int, error) {
	return _Native.Contract.PMMPRICEMODELID(&_Native.CallOpts)
}

// UNISWAPV2PRICEMODELID is a free data retrieval call binding the contract method 0x51e8f37a.
//
// Solidity: function UNISWAP_V2_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeCaller) UNISWAPV2PRICEMODELID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "UNISWAP_V2_PRICE_MODEL_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNISWAPV2PRICEMODELID is a free data retrieval call binding the contract method 0x51e8f37a.
//
// Solidity: function UNISWAP_V2_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeSession) UNISWAPV2PRICEMODELID() (*big.Int, error) {
	return _Native.Contract.UNISWAPV2PRICEMODELID(&_Native.CallOpts)
}

// UNISWAPV2PRICEMODELID is a free data retrieval call binding the contract method 0x51e8f37a.
//
// Solidity: function UNISWAP_V2_PRICE_MODEL_ID() view returns(uint256)
func (_Native *NativeCallerSession) UNISWAPV2PRICEMODELID() (*big.Int, error) {
	return _Native.Contract.UNISWAPV2PRICEMODELID(&_Native.CallOpts)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Native *NativeCaller) Blacklister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "blacklister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Native *NativeSession) Blacklister() (common.Address, error) {
	return _Native.Contract.Blacklister(&_Native.CallOpts)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_Native *NativeCallerSession) Blacklister() (common.Address, error) {
	return _Native.Contract.Blacklister(&_Native.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Native *NativeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Native *NativeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Native.Contract.Eip712Domain(&_Native.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Native *NativeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Native.Contract.Eip712Domain(&_Native.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address _tokenIn, address _tokenOut) view returns(uint256 amountOut)
func (_Native *NativeCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, _tokenIn common.Address, _tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "getAmountOut", amountIn, _tokenIn, _tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address _tokenIn, address _tokenOut) view returns(uint256 amountOut)
func (_Native *NativeSession) GetAmountOut(amountIn *big.Int, _tokenIn common.Address, _tokenOut common.Address) (*big.Int, error) {
	return _Native.Contract.GetAmountOut(&_Native.CallOpts, amountIn, _tokenIn, _tokenOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address _tokenIn, address _tokenOut) view returns(uint256 amountOut)
func (_Native *NativeCallerSession) GetAmountOut(amountIn *big.Int, _tokenIn common.Address, _tokenOut common.Address) (*big.Int, error) {
	return _Native.Contract.GetAmountOut(&_Native.CallOpts, amountIn, _tokenIn, _tokenOut)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_Native *NativeCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_Native *NativeSession) GetImplementation() (common.Address, error) {
	return _Native.Contract.GetImplementation(&_Native.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_Native *NativeCallerSession) GetImplementation() (common.Address, error) {
	return _Native.Contract.GetImplementation(&_Native.CallOpts)
}

// GetPairFee is a free data retrieval call binding the contract method 0xf47e1b99.
//
// Solidity: function getPairFee(address tokenIn, address tokenOut) view returns(uint256 fee)
func (_Native *NativeCaller) GetPairFee(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "getPairFee", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPairFee is a free data retrieval call binding the contract method 0xf47e1b99.
//
// Solidity: function getPairFee(address tokenIn, address tokenOut) view returns(uint256 fee)
func (_Native *NativeSession) GetPairFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Native.Contract.GetPairFee(&_Native.CallOpts, tokenIn, tokenOut)
}

// GetPairFee is a free data retrieval call binding the contract method 0xf47e1b99.
//
// Solidity: function getPairFee(address tokenIn, address tokenOut) view returns(uint256 fee)
func (_Native *NativeCallerSession) GetPairFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Native.Contract.GetPairFee(&_Native.CallOpts, tokenIn, tokenOut)
}

// GetPairPricingModel is a free data retrieval call binding the contract method 0x4a763c09.
//
// Solidity: function getPairPricingModel(address tokenIn, address tokenOut) view returns(uint256 pricingModelId)
func (_Native *NativeCaller) GetPairPricingModel(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "getPairPricingModel", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPairPricingModel is a free data retrieval call binding the contract method 0x4a763c09.
//
// Solidity: function getPairPricingModel(address tokenIn, address tokenOut) view returns(uint256 pricingModelId)
func (_Native *NativeSession) GetPairPricingModel(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Native.Contract.GetPairPricingModel(&_Native.CallOpts, tokenIn, tokenOut)
}

// GetPairPricingModel is a free data retrieval call binding the contract method 0x4a763c09.
//
// Solidity: function getPairPricingModel(address tokenIn, address tokenOut) view returns(uint256 pricingModelId)
func (_Native *NativeCallerSession) GetPairPricingModel(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Native.Contract.GetPairPricingModel(&_Native.CallOpts, tokenIn, tokenOut)
}

// GetPricingModelRegistry is a free data retrieval call binding the contract method 0x6aa45a38.
//
// Solidity: function getPricingModelRegistry() view returns(address)
func (_Native *NativeCaller) GetPricingModelRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "getPricingModelRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPricingModelRegistry is a free data retrieval call binding the contract method 0x6aa45a38.
//
// Solidity: function getPricingModelRegistry() view returns(address)
func (_Native *NativeSession) GetPricingModelRegistry() (common.Address, error) {
	return _Native.Contract.GetPricingModelRegistry(&_Native.CallOpts)
}

// GetPricingModelRegistry is a free data retrieval call binding the contract method 0x6aa45a38.
//
// Solidity: function getPricingModelRegistry() view returns(address)
func (_Native *NativeCallerSession) GetPricingModelRegistry() (common.Address, error) {
	return _Native.Contract.GetPricingModelRegistry(&_Native.CallOpts)
}

// GetTokenAs is a free data retrieval call binding the contract method 0xc4e3145b.
//
// Solidity: function getTokenAs() view returns(address[])
func (_Native *NativeCaller) GetTokenAs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "getTokenAs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenAs is a free data retrieval call binding the contract method 0xc4e3145b.
//
// Solidity: function getTokenAs() view returns(address[])
func (_Native *NativeSession) GetTokenAs() ([]common.Address, error) {
	return _Native.Contract.GetTokenAs(&_Native.CallOpts)
}

// GetTokenAs is a free data retrieval call binding the contract method 0xc4e3145b.
//
// Solidity: function getTokenAs() view returns(address[])
func (_Native *NativeCallerSession) GetTokenAs() ([]common.Address, error) {
	return _Native.Contract.GetTokenAs(&_Native.CallOpts)
}

// GetTokenBs is a free data retrieval call binding the contract method 0xfd4223ad.
//
// Solidity: function getTokenBs() view returns(address[])
func (_Native *NativeCaller) GetTokenBs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "getTokenBs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenBs is a free data retrieval call binding the contract method 0xfd4223ad.
//
// Solidity: function getTokenBs() view returns(address[])
func (_Native *NativeSession) GetTokenBs() ([]common.Address, error) {
	return _Native.Contract.GetTokenBs(&_Native.CallOpts)
}

// GetTokenBs is a free data retrieval call binding the contract method 0xfd4223ad.
//
// Solidity: function getTokenBs() view returns(address[])
func (_Native *NativeCallerSession) GetTokenBs() ([]common.Address, error) {
	return _Native.Contract.GetTokenBs(&_Native.CallOpts)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Native *NativeCaller) IsBlacklisted(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "isBlacklisted", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Native *NativeSession) IsBlacklisted(_account common.Address) (bool, error) {
	return _Native.Contract.IsBlacklisted(&_Native.CallOpts, _account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_Native *NativeCallerSession) IsBlacklisted(_account common.Address) (bool, error) {
	return _Native.Contract.IsBlacklisted(&_Native.CallOpts, _account)
}

// IsOnChainPricing is a free data retrieval call binding the contract method 0xf05c1a15.
//
// Solidity: function isOnChainPricing() view returns(bool)
func (_Native *NativeCaller) IsOnChainPricing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "isOnChainPricing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOnChainPricing is a free data retrieval call binding the contract method 0xf05c1a15.
//
// Solidity: function isOnChainPricing() view returns(bool)
func (_Native *NativeSession) IsOnChainPricing() (bool, error) {
	return _Native.Contract.IsOnChainPricing(&_Native.CallOpts)
}

// IsOnChainPricing is a free data retrieval call binding the contract method 0xf05c1a15.
//
// Solidity: function isOnChainPricing() view returns(bool)
func (_Native *NativeCallerSession) IsOnChainPricing() (bool, error) {
	return _Native.Contract.IsOnChainPricing(&_Native.CallOpts)
}

// IsPmm is a free data retrieval call binding the contract method 0xbc23d308.
//
// Solidity: function isPmm() view returns(bool)
func (_Native *NativeCaller) IsPmm(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "isPmm")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPmm is a free data retrieval call binding the contract method 0xbc23d308.
//
// Solidity: function isPmm() view returns(bool)
func (_Native *NativeSession) IsPmm() (bool, error) {
	return _Native.Contract.IsPmm(&_Native.CallOpts)
}

// IsPmm is a free data retrieval call binding the contract method 0xbc23d308.
//
// Solidity: function isPmm() view returns(bool)
func (_Native *NativeCallerSession) IsPmm() (bool, error) {
	return _Native.Contract.IsPmm(&_Native.CallOpts)
}

// IsPublicTreasury is a free data retrieval call binding the contract method 0xa6f75997.
//
// Solidity: function isPublicTreasury() view returns(bool)
func (_Native *NativeCaller) IsPublicTreasury(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "isPublicTreasury")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicTreasury is a free data retrieval call binding the contract method 0xa6f75997.
//
// Solidity: function isPublicTreasury() view returns(bool)
func (_Native *NativeSession) IsPublicTreasury() (bool, error) {
	return _Native.Contract.IsPublicTreasury(&_Native.CallOpts)
}

// IsPublicTreasury is a free data retrieval call binding the contract method 0xa6f75997.
//
// Solidity: function isPublicTreasury() view returns(bool)
func (_Native *NativeCallerSession) IsPublicTreasury() (bool, error) {
	return _Native.Contract.IsPublicTreasury(&_Native.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_Native *NativeCaller) IsSigner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "isSigner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_Native *NativeSession) IsSigner(arg0 common.Address) (bool, error) {
	return _Native.Contract.IsSigner(&_Native.CallOpts, arg0)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_Native *NativeCallerSession) IsSigner(arg0 common.Address) (bool, error) {
	return _Native.Contract.IsSigner(&_Native.CallOpts, arg0)
}

// IsTreasuryContract is a free data retrieval call binding the contract method 0xbce8f539.
//
// Solidity: function isTreasuryContract() view returns(bool)
func (_Native *NativeCaller) IsTreasuryContract(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "isTreasuryContract")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTreasuryContract is a free data retrieval call binding the contract method 0xbce8f539.
//
// Solidity: function isTreasuryContract() view returns(bool)
func (_Native *NativeSession) IsTreasuryContract() (bool, error) {
	return _Native.Contract.IsTreasuryContract(&_Native.CallOpts)
}

// IsTreasuryContract is a free data retrieval call binding the contract method 0xbce8f539.
//
// Solidity: function isTreasuryContract() view returns(bool)
func (_Native *NativeCallerSession) IsTreasuryContract() (bool, error) {
	return _Native.Contract.IsTreasuryContract(&_Native.CallOpts)
}

// NonceMapping is a free data retrieval call binding the contract method 0xe2c4cee2.
//
// Solidity: function nonceMapping(address , uint256 ) view returns(bool)
func (_Native *NativeCaller) NonceMapping(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "nonceMapping", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonceMapping is a free data retrieval call binding the contract method 0xe2c4cee2.
//
// Solidity: function nonceMapping(address , uint256 ) view returns(bool)
func (_Native *NativeSession) NonceMapping(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Native.Contract.NonceMapping(&_Native.CallOpts, arg0, arg1)
}

// NonceMapping is a free data retrieval call binding the contract method 0xe2c4cee2.
//
// Solidity: function nonceMapping(address , uint256 ) view returns(bool)
func (_Native *NativeCallerSession) NonceMapping(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Native.Contract.NonceMapping(&_Native.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Native *NativeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Native *NativeSession) Owner() (common.Address, error) {
	return _Native.Contract.Owner(&_Native.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Native *NativeCallerSession) Owner() (common.Address, error) {
	return _Native.Contract.Owner(&_Native.CallOpts)
}

// PairCount is a free data retrieval call binding the contract method 0x173fd1db.
//
// Solidity: function pairCount() view returns(uint256)
func (_Native *NativeCaller) PairCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "pairCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairCount is a free data retrieval call binding the contract method 0x173fd1db.
//
// Solidity: function pairCount() view returns(uint256)
func (_Native *NativeSession) PairCount() (*big.Int, error) {
	return _Native.Contract.PairCount(&_Native.CallOpts)
}

// PairCount is a free data retrieval call binding the contract method 0x173fd1db.
//
// Solidity: function pairCount() view returns(uint256)
func (_Native *NativeCallerSession) PairCount() (*big.Int, error) {
	return _Native.Contract.PairCount(&_Native.CallOpts)
}

// PairExist is a free data retrieval call binding the contract method 0xbcd4e9b1.
//
// Solidity: function pairExist(address tokenIn, address tokenOut) view returns(bool exist)
func (_Native *NativeCaller) PairExist(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (bool, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "pairExist", tokenIn, tokenOut)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PairExist is a free data retrieval call binding the contract method 0xbcd4e9b1.
//
// Solidity: function pairExist(address tokenIn, address tokenOut) view returns(bool exist)
func (_Native *NativeSession) PairExist(tokenIn common.Address, tokenOut common.Address) (bool, error) {
	return _Native.Contract.PairExist(&_Native.CallOpts, tokenIn, tokenOut)
}

// PairExist is a free data retrieval call binding the contract method 0xbcd4e9b1.
//
// Solidity: function pairExist(address tokenIn, address tokenOut) view returns(bool exist)
func (_Native *NativeCallerSession) PairExist(tokenIn common.Address, tokenOut common.Address) (bool, error) {
	return _Native.Contract.PairExist(&_Native.CallOpts, tokenIn, tokenOut)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Native *NativeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Native *NativeSession) Paused() (bool, error) {
	return _Native.Contract.Paused(&_Native.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Native *NativeCallerSession) Paused() (bool, error) {
	return _Native.Contract.Paused(&_Native.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Native *NativeCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Native *NativeSession) Pauser() (common.Address, error) {
	return _Native.Contract.Pauser(&_Native.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Native *NativeCallerSession) Pauser() (common.Address, error) {
	return _Native.Contract.Pauser(&_Native.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Native *NativeCaller) PoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "poolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Native *NativeSession) PoolFactory() (common.Address, error) {
	return _Native.Contract.PoolFactory(&_Native.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_Native *NativeCallerSession) PoolFactory() (common.Address, error) {
	return _Native.Contract.PoolFactory(&_Native.CallOpts)
}

// PricingModelRegistry is a free data retrieval call binding the contract method 0x4edc9f3e.
//
// Solidity: function pricingModelRegistry() view returns(address)
func (_Native *NativeCaller) PricingModelRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "pricingModelRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PricingModelRegistry is a free data retrieval call binding the contract method 0x4edc9f3e.
//
// Solidity: function pricingModelRegistry() view returns(address)
func (_Native *NativeSession) PricingModelRegistry() (common.Address, error) {
	return _Native.Contract.PricingModelRegistry(&_Native.CallOpts)
}

// PricingModelRegistry is a free data retrieval call binding the contract method 0x4edc9f3e.
//
// Solidity: function pricingModelRegistry() view returns(address)
func (_Native *NativeCallerSession) PricingModelRegistry() (common.Address, error) {
	return _Native.Contract.PricingModelRegistry(&_Native.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Native *NativeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Native *NativeSession) ProxiableUUID() ([32]byte, error) {
	return _Native.Contract.ProxiableUUID(&_Native.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Native *NativeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Native.Contract.ProxiableUUID(&_Native.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Native *NativeCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Native *NativeSession) Router() (common.Address, error) {
	return _Native.Contract.Router(&_Native.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Native *NativeCallerSession) Router() (common.Address, error) {
	return _Native.Contract.Router(&_Native.CallOpts)
}

// TokenAs is a free data retrieval call binding the contract method 0xccd0e418.
//
// Solidity: function tokenAs(uint256 ) view returns(address)
func (_Native *NativeCaller) TokenAs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "tokenAs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAs is a free data retrieval call binding the contract method 0xccd0e418.
//
// Solidity: function tokenAs(uint256 ) view returns(address)
func (_Native *NativeSession) TokenAs(arg0 *big.Int) (common.Address, error) {
	return _Native.Contract.TokenAs(&_Native.CallOpts, arg0)
}

// TokenAs is a free data retrieval call binding the contract method 0xccd0e418.
//
// Solidity: function tokenAs(uint256 ) view returns(address)
func (_Native *NativeCallerSession) TokenAs(arg0 *big.Int) (common.Address, error) {
	return _Native.Contract.TokenAs(&_Native.CallOpts, arg0)
}

// TokenBs is a free data retrieval call binding the contract method 0xbcc341ca.
//
// Solidity: function tokenBs(uint256 ) view returns(address)
func (_Native *NativeCaller) TokenBs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "tokenBs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenBs is a free data retrieval call binding the contract method 0xbcc341ca.
//
// Solidity: function tokenBs(uint256 ) view returns(address)
func (_Native *NativeSession) TokenBs(arg0 *big.Int) (common.Address, error) {
	return _Native.Contract.TokenBs(&_Native.CallOpts, arg0)
}

// TokenBs is a free data retrieval call binding the contract method 0xbcc341ca.
//
// Solidity: function tokenBs(uint256 ) view returns(address)
func (_Native *NativeCallerSession) TokenBs(arg0 *big.Int) (common.Address, error) {
	return _Native.Contract.TokenBs(&_Native.CallOpts, arg0)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Native *NativeCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Native *NativeSession) Treasury() (common.Address, error) {
	return _Native.Contract.Treasury(&_Native.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Native *NativeCallerSession) Treasury() (common.Address, error) {
	return _Native.Contract.Treasury(&_Native.CallOpts)
}

// TreasuryOwner is a free data retrieval call binding the contract method 0x629d3f8f.
//
// Solidity: function treasuryOwner() view returns(address)
func (_Native *NativeCaller) TreasuryOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Native.contract.Call(opts, &out, "treasuryOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryOwner is a free data retrieval call binding the contract method 0x629d3f8f.
//
// Solidity: function treasuryOwner() view returns(address)
func (_Native *NativeSession) TreasuryOwner() (common.Address, error) {
	return _Native.Contract.TreasuryOwner(&_Native.CallOpts)
}

// TreasuryOwner is a free data retrieval call binding the contract method 0x629d3f8f.
//
// Solidity: function treasuryOwner() view returns(address)
func (_Native *NativeCallerSession) TreasuryOwner() (common.Address, error) {
	return _Native.Contract.TreasuryOwner(&_Native.CallOpts)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address _signer) returns()
func (_Native *NativeTransactor) AddSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "addSigner", _signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address _signer) returns()
func (_Native *NativeSession) AddSigner(_signer common.Address) (*types.Transaction, error) {
	return _Native.Contract.AddSigner(&_Native.TransactOpts, _signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address _signer) returns()
func (_Native *NativeTransactorSession) AddSigner(_signer common.Address) (*types.Transaction, error) {
	return _Native.Contract.AddSigner(&_Native.TransactOpts, _signer)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Native *NativeTransactor) Blacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "blacklist", _account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Native *NativeSession) Blacklist(_account common.Address) (*types.Transaction, error) {
	return _Native.Contract.Blacklist(&_Native.TransactOpts, _account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_Native *NativeTransactorSession) Blacklist(_account common.Address) (*types.Transaction, error) {
	return _Native.Contract.Blacklist(&_Native.TransactOpts, _account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8b1efc71.
//
// Solidity: function initialize((address,address,address,address,bool,bool,uint256[],address[],address[],uint256[]) poolConfig, address _pricingModelRegistry) returns()
func (_Native *NativeTransactor) Initialize(opts *bind.TransactOpts, poolConfig INativePoolNewPoolConfig, _pricingModelRegistry common.Address) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "initialize", poolConfig, _pricingModelRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x8b1efc71.
//
// Solidity: function initialize((address,address,address,address,bool,bool,uint256[],address[],address[],uint256[]) poolConfig, address _pricingModelRegistry) returns()
func (_Native *NativeSession) Initialize(poolConfig INativePoolNewPoolConfig, _pricingModelRegistry common.Address) (*types.Transaction, error) {
	return _Native.Contract.Initialize(&_Native.TransactOpts, poolConfig, _pricingModelRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x8b1efc71.
//
// Solidity: function initialize((address,address,address,address,bool,bool,uint256[],address[],address[],uint256[]) poolConfig, address _pricingModelRegistry) returns()
func (_Native *NativeTransactorSession) Initialize(poolConfig INativePoolNewPoolConfig, _pricingModelRegistry common.Address) (*types.Transaction, error) {
	return _Native.Contract.Initialize(&_Native.TransactOpts, poolConfig, _pricingModelRegistry)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Native *NativeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Native *NativeSession) Pause() (*types.Transaction, error) {
	return _Native.Contract.Pause(&_Native.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Native *NativeTransactorSession) Pause() (*types.Transaction, error) {
	return _Native.Contract.Pause(&_Native.TransactOpts)
}

// RemovePair is a paid mutator transaction binding the contract method 0x2b42dc62.
//
// Solidity: function removePair(uint256 removingIdx) returns()
func (_Native *NativeTransactor) RemovePair(opts *bind.TransactOpts, removingIdx *big.Int) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "removePair", removingIdx)
}

// RemovePair is a paid mutator transaction binding the contract method 0x2b42dc62.
//
// Solidity: function removePair(uint256 removingIdx) returns()
func (_Native *NativeSession) RemovePair(removingIdx *big.Int) (*types.Transaction, error) {
	return _Native.Contract.RemovePair(&_Native.TransactOpts, removingIdx)
}

// RemovePair is a paid mutator transaction binding the contract method 0x2b42dc62.
//
// Solidity: function removePair(uint256 removingIdx) returns()
func (_Native *NativeTransactorSession) RemovePair(removingIdx *big.Int) (*types.Transaction, error) {
	return _Native.Contract.RemovePair(&_Native.TransactOpts, removingIdx)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address _signer) returns()
func (_Native *NativeTransactor) RemoveSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "removeSigner", _signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address _signer) returns()
func (_Native *NativeSession) RemoveSigner(_signer common.Address) (*types.Transaction, error) {
	return _Native.Contract.RemoveSigner(&_Native.TransactOpts, _signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address _signer) returns()
func (_Native *NativeTransactorSession) RemoveSigner(_signer common.Address) (*types.Transaction, error) {
	return _Native.Contract.RemoveSigner(&_Native.TransactOpts, _signer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Native *NativeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Native *NativeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Native.Contract.RenounceOwnership(&_Native.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Native *NativeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Native.Contract.RenounceOwnership(&_Native.TransactOpts)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Native *NativeTransactor) SetPauser(opts *bind.TransactOpts, _pauser common.Address) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "setPauser", _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Native *NativeSession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _Native.Contract.SetPauser(&_Native.TransactOpts, _pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address _pauser) returns()
func (_Native *NativeTransactorSession) SetPauser(_pauser common.Address) (*types.Transaction, error) {
	return _Native.Contract.SetPauser(&_Native.TransactOpts, _pauser)
}

// Swap is a paid mutator transaction binding the contract method 0xd025fdfa.
//
// Solidity: function swap(bytes order, bytes signature, uint256 flexibleAmount, address recipient, bytes callback) returns(int256, int256)
func (_Native *NativeTransactor) Swap(opts *bind.TransactOpts, order []byte, signature []byte, flexibleAmount *big.Int, recipient common.Address, callback []byte) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "swap", order, signature, flexibleAmount, recipient, callback)
}

// Swap is a paid mutator transaction binding the contract method 0xd025fdfa.
//
// Solidity: function swap(bytes order, bytes signature, uint256 flexibleAmount, address recipient, bytes callback) returns(int256, int256)
func (_Native *NativeSession) Swap(order []byte, signature []byte, flexibleAmount *big.Int, recipient common.Address, callback []byte) (*types.Transaction, error) {
	return _Native.Contract.Swap(&_Native.TransactOpts, order, signature, flexibleAmount, recipient, callback)
}

// Swap is a paid mutator transaction binding the contract method 0xd025fdfa.
//
// Solidity: function swap(bytes order, bytes signature, uint256 flexibleAmount, address recipient, bytes callback) returns(int256, int256)
func (_Native *NativeTransactorSession) Swap(order []byte, signature []byte, flexibleAmount *big.Int, recipient common.Address, callback []byte) (*types.Transaction, error) {
	return _Native.Contract.Swap(&_Native.TransactOpts, order, signature, flexibleAmount, recipient, callback)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Native *NativeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Native *NativeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Native.Contract.TransferOwnership(&_Native.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Native *NativeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Native.Contract.TransferOwnership(&_Native.TransactOpts, newOwner)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Native *NativeTransactor) UnBlacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "unBlacklist", _account)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Native *NativeSession) UnBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Native.Contract.UnBlacklist(&_Native.TransactOpts, _account)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_Native *NativeTransactorSession) UnBlacklist(_account common.Address) (*types.Transaction, error) {
	return _Native.Contract.UnBlacklist(&_Native.TransactOpts, _account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Native *NativeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Native *NativeSession) Unpause() (*types.Transaction, error) {
	return _Native.Contract.Unpause(&_Native.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Native *NativeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Native.Contract.Unpause(&_Native.TransactOpts)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Native *NativeTransactor) UpdateBlacklister(opts *bind.TransactOpts, _newBlacklister common.Address) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "updateBlacklister", _newBlacklister)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Native *NativeSession) UpdateBlacklister(_newBlacklister common.Address) (*types.Transaction, error) {
	return _Native.Contract.UpdateBlacklister(&_Native.TransactOpts, _newBlacklister)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_Native *NativeTransactorSession) UpdateBlacklister(_newBlacklister common.Address) (*types.Transaction, error) {
	return _Native.Contract.UpdateBlacklister(&_Native.TransactOpts, _newBlacklister)
}

// UpdatePairs is a paid mutator transaction binding the contract method 0xc11fbae4.
//
// Solidity: function updatePairs(uint256[] _fees, address[] _tokenAs, address[] _tokenBs, uint256[] _pricingModelIds) returns()
func (_Native *NativeTransactor) UpdatePairs(opts *bind.TransactOpts, _fees []*big.Int, _tokenAs []common.Address, _tokenBs []common.Address, _pricingModelIds []*big.Int) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "updatePairs", _fees, _tokenAs, _tokenBs, _pricingModelIds)
}

// UpdatePairs is a paid mutator transaction binding the contract method 0xc11fbae4.
//
// Solidity: function updatePairs(uint256[] _fees, address[] _tokenAs, address[] _tokenBs, uint256[] _pricingModelIds) returns()
func (_Native *NativeSession) UpdatePairs(_fees []*big.Int, _tokenAs []common.Address, _tokenBs []common.Address, _pricingModelIds []*big.Int) (*types.Transaction, error) {
	return _Native.Contract.UpdatePairs(&_Native.TransactOpts, _fees, _tokenAs, _tokenBs, _pricingModelIds)
}

// UpdatePairs is a paid mutator transaction binding the contract method 0xc11fbae4.
//
// Solidity: function updatePairs(uint256[] _fees, address[] _tokenAs, address[] _tokenBs, uint256[] _pricingModelIds) returns()
func (_Native *NativeTransactorSession) UpdatePairs(_fees []*big.Int, _tokenAs []common.Address, _tokenBs []common.Address, _pricingModelIds []*big.Int) (*types.Transaction, error) {
	return _Native.Contract.UpdatePairs(&_Native.TransactOpts, _fees, _tokenAs, _tokenBs, _pricingModelIds)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Native *NativeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Native *NativeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Native.Contract.UpgradeTo(&_Native.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Native *NativeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Native.Contract.UpgradeTo(&_Native.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Native *NativeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Native.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Native *NativeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Native.Contract.UpgradeToAndCall(&_Native.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Native *NativeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Native.Contract.UpgradeToAndCall(&_Native.TransactOpts, newImplementation, data)
}

// NativeAddSignerIterator is returned from FilterAddSigner and is used to iterate over the raw logs and unpacked data for AddSigner events raised by the Native contract.
type NativeAddSignerIterator struct {
	Event *NativeAddSigner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeAddSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeAddSigner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeAddSigner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeAddSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeAddSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeAddSigner represents a AddSigner event raised by the Native contract.
type NativeAddSigner struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddSigner is a free log retrieval operation binding the contract event 0x637c77a2d598a51d085d4a2413332c45a235a25ee855bf3dfcdc2c8fcf02860f.
//
// Solidity: event AddSigner(address signer)
func (_Native *NativeFilterer) FilterAddSigner(opts *bind.FilterOpts) (*NativeAddSignerIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "AddSigner")
	if err != nil {
		return nil, err
	}
	return &NativeAddSignerIterator{contract: _Native.contract, event: "AddSigner", logs: logs, sub: sub}, nil
}

// WatchAddSigner is a free log subscription operation binding the contract event 0x637c77a2d598a51d085d4a2413332c45a235a25ee855bf3dfcdc2c8fcf02860f.
//
// Solidity: event AddSigner(address signer)
func (_Native *NativeFilterer) WatchAddSigner(opts *bind.WatchOpts, sink chan<- *NativeAddSigner) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "AddSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeAddSigner)
				if err := _Native.contract.UnpackLog(event, "AddSigner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddSigner is a log parse operation binding the contract event 0x637c77a2d598a51d085d4a2413332c45a235a25ee855bf3dfcdc2c8fcf02860f.
//
// Solidity: event AddSigner(address signer)
func (_Native *NativeFilterer) ParseAddSigner(log types.Log) (*NativeAddSigner, error) {
	event := new(NativeAddSigner)
	if err := _Native.contract.UnpackLog(event, "AddSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Native contract.
type NativeAdminChangedIterator struct {
	Event *NativeAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeAdminChanged represents a AdminChanged event raised by the Native contract.
type NativeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Native *NativeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NativeAdminChangedIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NativeAdminChangedIterator{contract: _Native.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Native *NativeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NativeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeAdminChanged)
				if err := _Native.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Native *NativeFilterer) ParseAdminChanged(log types.Log) (*NativeAdminChanged, error) {
	event := new(NativeAdminChanged)
	if err := _Native.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Native contract.
type NativeBeaconUpgradedIterator struct {
	Event *NativeBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeBeaconUpgraded represents a BeaconUpgraded event raised by the Native contract.
type NativeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Native *NativeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NativeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Native.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NativeBeaconUpgradedIterator{contract: _Native.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Native *NativeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NativeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Native.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeBeaconUpgraded)
				if err := _Native.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Native *NativeFilterer) ParseBeaconUpgraded(log types.Log) (*NativeBeaconUpgraded, error) {
	event := new(NativeBeaconUpgraded)
	if err := _Native.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeBlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the Native contract.
type NativeBlacklistedIterator struct {
	Event *NativeBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeBlacklisted represents a Blacklisted event raised by the Native contract.
type NativeBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Native *NativeFilterer) FilterBlacklisted(opts *bind.FilterOpts, _account []common.Address) (*NativeBlacklistedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Native.contract.FilterLogs(opts, "Blacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return &NativeBlacklistedIterator{contract: _Native.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Native *NativeFilterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *NativeBlacklisted, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Native.contract.WatchLogs(opts, "Blacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeBlacklisted)
				if err := _Native.contract.UnpackLog(event, "Blacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlacklisted is a log parse operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_Native *NativeFilterer) ParseBlacklisted(log types.Log) (*NativeBlacklisted, error) {
	event := new(NativeBlacklisted)
	if err := _Native.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeBlacklisterChangedIterator is returned from FilterBlacklisterChanged and is used to iterate over the raw logs and unpacked data for BlacklisterChanged events raised by the Native contract.
type NativeBlacklisterChangedIterator struct {
	Event *NativeBlacklisterChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeBlacklisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeBlacklisterChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeBlacklisterChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeBlacklisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeBlacklisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeBlacklisterChanged represents a BlacklisterChanged event raised by the Native contract.
type NativeBlacklisterChanged struct {
	NewBlacklister common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlacklisterChanged is a free log retrieval operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Native *NativeFilterer) FilterBlacklisterChanged(opts *bind.FilterOpts, newBlacklister []common.Address) (*NativeBlacklisterChangedIterator, error) {

	var newBlacklisterRule []interface{}
	for _, newBlacklisterItem := range newBlacklister {
		newBlacklisterRule = append(newBlacklisterRule, newBlacklisterItem)
	}

	logs, sub, err := _Native.contract.FilterLogs(opts, "BlacklisterChanged", newBlacklisterRule)
	if err != nil {
		return nil, err
	}
	return &NativeBlacklisterChangedIterator{contract: _Native.contract, event: "BlacklisterChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklisterChanged is a free log subscription operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Native *NativeFilterer) WatchBlacklisterChanged(opts *bind.WatchOpts, sink chan<- *NativeBlacklisterChanged, newBlacklister []common.Address) (event.Subscription, error) {

	var newBlacklisterRule []interface{}
	for _, newBlacklisterItem := range newBlacklister {
		newBlacklisterRule = append(newBlacklisterRule, newBlacklisterItem)
	}

	logs, sub, err := _Native.contract.WatchLogs(opts, "BlacklisterChanged", newBlacklisterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeBlacklisterChanged)
				if err := _Native.contract.UnpackLog(event, "BlacklisterChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlacklisterChanged is a log parse operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_Native *NativeFilterer) ParseBlacklisterChanged(log types.Log) (*NativeBlacklisterChanged, error) {
	event := new(NativeBlacklisterChanged)
	if err := _Native.contract.UnpackLog(event, "BlacklisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Native contract.
type NativeEIP712DomainChangedIterator struct {
	Event *NativeEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeEIP712DomainChanged represents a EIP712DomainChanged event raised by the Native contract.
type NativeEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Native *NativeFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*NativeEIP712DomainChangedIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &NativeEIP712DomainChangedIterator{contract: _Native.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Native *NativeFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *NativeEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeEIP712DomainChanged)
				if err := _Native.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Native *NativeFilterer) ParseEIP712DomainChanged(log types.Log) (*NativeEIP712DomainChanged, error) {
	event := new(NativeEIP712DomainChanged)
	if err := _Native.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Native contract.
type NativeInitializedIterator struct {
	Event *NativeInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeInitialized represents a Initialized event raised by the Native contract.
type NativeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Native *NativeFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeInitializedIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeInitializedIterator{contract: _Native.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Native *NativeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeInitialized) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeInitialized)
				if err := _Native.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Native *NativeFilterer) ParseInitialized(log types.Log) (*NativeInitialized, error) {
	event := new(NativeInitialized)
	if err := _Native.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Native contract.
type NativeOwnershipTransferredIterator struct {
	Event *NativeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeOwnershipTransferred represents a OwnershipTransferred event raised by the Native contract.
type NativeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Native *NativeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NativeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Native.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NativeOwnershipTransferredIterator{contract: _Native.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Native *NativeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NativeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Native.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeOwnershipTransferred)
				if err := _Native.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Native *NativeFilterer) ParseOwnershipTransferred(log types.Log) (*NativeOwnershipTransferred, error) {
	event := new(NativeOwnershipTransferred)
	if err := _Native.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Native contract.
type NativePausedIterator struct {
	Event *NativePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativePaused represents a Paused event raised by the Native contract.
type NativePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Native *NativeFilterer) FilterPaused(opts *bind.FilterOpts) (*NativePausedIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NativePausedIterator{contract: _Native.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Native *NativeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NativePaused) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativePaused)
				if err := _Native.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Native *NativeFilterer) ParsePaused(log types.Log) (*NativePaused, error) {
	event := new(NativePaused)
	if err := _Native.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeRemovePairIterator is returned from FilterRemovePair and is used to iterate over the raw logs and unpacked data for RemovePair events raised by the Native contract.
type NativeRemovePairIterator struct {
	Event *NativeRemovePair // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeRemovePairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeRemovePair)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeRemovePair)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeRemovePairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeRemovePairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeRemovePair represents a RemovePair event raised by the Native contract.
type NativeRemovePair struct {
	TokenA common.Address
	TokenB common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemovePair is a free log retrieval operation binding the contract event 0xf8846cbe66fa5bf657d787c9d1fa1e0c49d164118e743861a349b498d21f9152.
//
// Solidity: event RemovePair(address tokenA, address tokenB)
func (_Native *NativeFilterer) FilterRemovePair(opts *bind.FilterOpts) (*NativeRemovePairIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "RemovePair")
	if err != nil {
		return nil, err
	}
	return &NativeRemovePairIterator{contract: _Native.contract, event: "RemovePair", logs: logs, sub: sub}, nil
}

// WatchRemovePair is a free log subscription operation binding the contract event 0xf8846cbe66fa5bf657d787c9d1fa1e0c49d164118e743861a349b498d21f9152.
//
// Solidity: event RemovePair(address tokenA, address tokenB)
func (_Native *NativeFilterer) WatchRemovePair(opts *bind.WatchOpts, sink chan<- *NativeRemovePair) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "RemovePair")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeRemovePair)
				if err := _Native.contract.UnpackLog(event, "RemovePair", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemovePair is a log parse operation binding the contract event 0xf8846cbe66fa5bf657d787c9d1fa1e0c49d164118e743861a349b498d21f9152.
//
// Solidity: event RemovePair(address tokenA, address tokenB)
func (_Native *NativeFilterer) ParseRemovePair(log types.Log) (*NativeRemovePair, error) {
	event := new(NativeRemovePair)
	if err := _Native.contract.UnpackLog(event, "RemovePair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeRemoveSignerIterator is returned from FilterRemoveSigner and is used to iterate over the raw logs and unpacked data for RemoveSigner events raised by the Native contract.
type NativeRemoveSignerIterator struct {
	Event *NativeRemoveSigner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeRemoveSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeRemoveSigner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeRemoveSigner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeRemoveSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeRemoveSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeRemoveSigner represents a RemoveSigner event raised by the Native contract.
type NativeRemoveSigner struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveSigner is a free log retrieval operation binding the contract event 0x1803740ef72fc16e647c10fe2d31cf61a1578081960c2e3fb7f5aa957e82f550.
//
// Solidity: event RemoveSigner(address signer)
func (_Native *NativeFilterer) FilterRemoveSigner(opts *bind.FilterOpts) (*NativeRemoveSignerIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "RemoveSigner")
	if err != nil {
		return nil, err
	}
	return &NativeRemoveSignerIterator{contract: _Native.contract, event: "RemoveSigner", logs: logs, sub: sub}, nil
}

// WatchRemoveSigner is a free log subscription operation binding the contract event 0x1803740ef72fc16e647c10fe2d31cf61a1578081960c2e3fb7f5aa957e82f550.
//
// Solidity: event RemoveSigner(address signer)
func (_Native *NativeFilterer) WatchRemoveSigner(opts *bind.WatchOpts, sink chan<- *NativeRemoveSigner) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "RemoveSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeRemoveSigner)
				if err := _Native.contract.UnpackLog(event, "RemoveSigner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveSigner is a log parse operation binding the contract event 0x1803740ef72fc16e647c10fe2d31cf61a1578081960c2e3fb7f5aa957e82f550.
//
// Solidity: event RemoveSigner(address signer)
func (_Native *NativeFilterer) ParseRemoveSigner(log types.Log) (*NativeRemoveSigner, error) {
	event := new(NativeRemoveSigner)
	if err := _Native.contract.UnpackLog(event, "RemoveSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeSetRouterIterator is returned from FilterSetRouter and is used to iterate over the raw logs and unpacked data for SetRouter events raised by the Native contract.
type NativeSetRouterIterator struct {
	Event *NativeSetRouter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeSetRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeSetRouter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeSetRouter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeSetRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeSetRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeSetRouter represents a SetRouter event raised by the Native contract.
type NativeSetRouter struct {
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetRouter is a free log retrieval operation binding the contract event 0x6de4326a8b9054d72d9dbab97d27bc4edffadee7d966f5af9cc4eafdaf8e5455.
//
// Solidity: event SetRouter(address router)
func (_Native *NativeFilterer) FilterSetRouter(opts *bind.FilterOpts) (*NativeSetRouterIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "SetRouter")
	if err != nil {
		return nil, err
	}
	return &NativeSetRouterIterator{contract: _Native.contract, event: "SetRouter", logs: logs, sub: sub}, nil
}

// WatchSetRouter is a free log subscription operation binding the contract event 0x6de4326a8b9054d72d9dbab97d27bc4edffadee7d966f5af9cc4eafdaf8e5455.
//
// Solidity: event SetRouter(address router)
func (_Native *NativeFilterer) WatchSetRouter(opts *bind.WatchOpts, sink chan<- *NativeSetRouter) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "SetRouter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeSetRouter)
				if err := _Native.contract.UnpackLog(event, "SetRouter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetRouter is a log parse operation binding the contract event 0x6de4326a8b9054d72d9dbab97d27bc4edffadee7d966f5af9cc4eafdaf8e5455.
//
// Solidity: event SetRouter(address router)
func (_Native *NativeFilterer) ParseSetRouter(log types.Log) (*NativeSetRouter, error) {
	event := new(NativeSetRouter)
	if err := _Native.contract.UnpackLog(event, "SetRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeSetTreasuryIterator is returned from FilterSetTreasury and is used to iterate over the raw logs and unpacked data for SetTreasury events raised by the Native contract.
type NativeSetTreasuryIterator struct {
	Event *NativeSetTreasury // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeSetTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeSetTreasury)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeSetTreasury)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeSetTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeSetTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeSetTreasury represents a SetTreasury event raised by the Native contract.
type NativeSetTreasury struct {
	Treasury common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetTreasury is a free log retrieval operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address treasury)
func (_Native *NativeFilterer) FilterSetTreasury(opts *bind.FilterOpts) (*NativeSetTreasuryIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "SetTreasury")
	if err != nil {
		return nil, err
	}
	return &NativeSetTreasuryIterator{contract: _Native.contract, event: "SetTreasury", logs: logs, sub: sub}, nil
}

// WatchSetTreasury is a free log subscription operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address treasury)
func (_Native *NativeFilterer) WatchSetTreasury(opts *bind.WatchOpts, sink chan<- *NativeSetTreasury) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "SetTreasury")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeSetTreasury)
				if err := _Native.contract.UnpackLog(event, "SetTreasury", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTreasury is a log parse operation binding the contract event 0xcb7ef3e545f5cdb893f5c568ba710fe08f336375a2d9fd66e161033f8fc09ef3.
//
// Solidity: event SetTreasury(address treasury)
func (_Native *NativeFilterer) ParseSetTreasury(log types.Log) (*NativeSetTreasury, error) {
	event := new(NativeSetTreasury)
	if err := _Native.contract.UnpackLog(event, "SetTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeSetTreasuryOwnerIterator is returned from FilterSetTreasuryOwner and is used to iterate over the raw logs and unpacked data for SetTreasuryOwner events raised by the Native contract.
type NativeSetTreasuryOwnerIterator struct {
	Event *NativeSetTreasuryOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeSetTreasuryOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeSetTreasuryOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeSetTreasuryOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeSetTreasuryOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeSetTreasuryOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeSetTreasuryOwner represents a SetTreasuryOwner event raised by the Native contract.
type NativeSetTreasuryOwner struct {
	TreasuryOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTreasuryOwner is a free log retrieval operation binding the contract event 0x7a763988f90310111b5423a29fd3edfadcca9bc03ea4d3f9feeebb59dd9291f3.
//
// Solidity: event SetTreasuryOwner(address treasuryOwner)
func (_Native *NativeFilterer) FilterSetTreasuryOwner(opts *bind.FilterOpts) (*NativeSetTreasuryOwnerIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "SetTreasuryOwner")
	if err != nil {
		return nil, err
	}
	return &NativeSetTreasuryOwnerIterator{contract: _Native.contract, event: "SetTreasuryOwner", logs: logs, sub: sub}, nil
}

// WatchSetTreasuryOwner is a free log subscription operation binding the contract event 0x7a763988f90310111b5423a29fd3edfadcca9bc03ea4d3f9feeebb59dd9291f3.
//
// Solidity: event SetTreasuryOwner(address treasuryOwner)
func (_Native *NativeFilterer) WatchSetTreasuryOwner(opts *bind.WatchOpts, sink chan<- *NativeSetTreasuryOwner) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "SetTreasuryOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeSetTreasuryOwner)
				if err := _Native.contract.UnpackLog(event, "SetTreasuryOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTreasuryOwner is a log parse operation binding the contract event 0x7a763988f90310111b5423a29fd3edfadcca9bc03ea4d3f9feeebb59dd9291f3.
//
// Solidity: event SetTreasuryOwner(address treasuryOwner)
func (_Native *NativeFilterer) ParseSetTreasuryOwner(log types.Log) (*NativeSetTreasuryOwner, error) {
	event := new(NativeSetTreasuryOwner)
	if err := _Native.contract.UnpackLog(event, "SetTreasuryOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Native contract.
type NativeSwapIterator struct {
	Event *NativeSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeSwap represents a Swap event raised by the Native contract.
type NativeSwap struct {
	Sender    common.Address
	Recipient common.Address
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Fee       *big.Int
	QuoteId   [16]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xe3a54b69726c85299f4e794bac96150af56af801be76cafd11947a1103b6308a.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, address tokenIn, address tokenOut, int256 amountIn, int256 amountOut, uint256 fee, bytes16 quoteId)
func (_Native *NativeFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*NativeSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Native.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeSwapIterator{contract: _Native.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xe3a54b69726c85299f4e794bac96150af56af801be76cafd11947a1103b6308a.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, address tokenIn, address tokenOut, int256 amountIn, int256 amountOut, uint256 fee, bytes16 quoteId)
func (_Native *NativeFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *NativeSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Native.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeSwap)
				if err := _Native.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0xe3a54b69726c85299f4e794bac96150af56af801be76cafd11947a1103b6308a.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, address tokenIn, address tokenOut, int256 amountIn, int256 amountOut, uint256 fee, bytes16 quoteId)
func (_Native *NativeFilterer) ParseSwap(log types.Log) (*NativeSwap, error) {
	event := new(NativeSwap)
	if err := _Native.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeUnBlacklistedIterator is returned from FilterUnBlacklisted and is used to iterate over the raw logs and unpacked data for UnBlacklisted events raised by the Native contract.
type NativeUnBlacklistedIterator struct {
	Event *NativeUnBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeUnBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeUnBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeUnBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeUnBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeUnBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeUnBlacklisted represents a UnBlacklisted event raised by the Native contract.
type NativeUnBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnBlacklisted is a free log retrieval operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Native *NativeFilterer) FilterUnBlacklisted(opts *bind.FilterOpts, _account []common.Address) (*NativeUnBlacklistedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Native.contract.FilterLogs(opts, "UnBlacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return &NativeUnBlacklistedIterator{contract: _Native.contract, event: "UnBlacklisted", logs: logs, sub: sub}, nil
}

// WatchUnBlacklisted is a free log subscription operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Native *NativeFilterer) WatchUnBlacklisted(opts *bind.WatchOpts, sink chan<- *NativeUnBlacklisted, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Native.contract.WatchLogs(opts, "UnBlacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeUnBlacklisted)
				if err := _Native.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnBlacklisted is a log parse operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_Native *NativeFilterer) ParseUnBlacklisted(log types.Log) (*NativeUnBlacklisted, error) {
	event := new(NativeUnBlacklisted)
	if err := _Native.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Native contract.
type NativeUnpausedIterator struct {
	Event *NativeUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeUnpaused represents a Unpaused event raised by the Native contract.
type NativeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Native *NativeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NativeUnpausedIterator, error) {

	logs, sub, err := _Native.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NativeUnpausedIterator{contract: _Native.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Native *NativeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NativeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Native.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeUnpaused)
				if err := _Native.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Native *NativeFilterer) ParseUnpaused(log types.Log) (*NativeUnpaused, error) {
	event := new(NativeUnpaused)
	if err := _Native.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeUpdatePairIterator is returned from FilterUpdatePair and is used to iterate over the raw logs and unpacked data for UpdatePair events raised by the Native contract.
type NativeUpdatePairIterator struct {
	Event *NativeUpdatePair // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeUpdatePairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeUpdatePair)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeUpdatePair)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeUpdatePairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeUpdatePairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeUpdatePair represents a UpdatePair event raised by the Native contract.
type NativeUpdatePair struct {
	TokenA            common.Address
	TokenB            common.Address
	FeeOld            *big.Int
	FeeNew            *big.Int
	PricingModelIdOld *big.Int
	PricingModelIdNew *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdatePair is a free log retrieval operation binding the contract event 0x2d498c7d2a877bb8a3c6ec493a1743ddf467889264eb046fbfe2ca5ee2b4e752.
//
// Solidity: event UpdatePair(address indexed tokenA, address indexed tokenB, uint256 feeOld, uint256 feeNew, uint256 pricingModelIdOld, uint256 pricingModelIdNew)
func (_Native *NativeFilterer) FilterUpdatePair(opts *bind.FilterOpts, tokenA []common.Address, tokenB []common.Address) (*NativeUpdatePairIterator, error) {

	var tokenARule []interface{}
	for _, tokenAItem := range tokenA {
		tokenARule = append(tokenARule, tokenAItem)
	}
	var tokenBRule []interface{}
	for _, tokenBItem := range tokenB {
		tokenBRule = append(tokenBRule, tokenBItem)
	}

	logs, sub, err := _Native.contract.FilterLogs(opts, "UpdatePair", tokenARule, tokenBRule)
	if err != nil {
		return nil, err
	}
	return &NativeUpdatePairIterator{contract: _Native.contract, event: "UpdatePair", logs: logs, sub: sub}, nil
}

// WatchUpdatePair is a free log subscription operation binding the contract event 0x2d498c7d2a877bb8a3c6ec493a1743ddf467889264eb046fbfe2ca5ee2b4e752.
//
// Solidity: event UpdatePair(address indexed tokenA, address indexed tokenB, uint256 feeOld, uint256 feeNew, uint256 pricingModelIdOld, uint256 pricingModelIdNew)
func (_Native *NativeFilterer) WatchUpdatePair(opts *bind.WatchOpts, sink chan<- *NativeUpdatePair, tokenA []common.Address, tokenB []common.Address) (event.Subscription, error) {

	var tokenARule []interface{}
	for _, tokenAItem := range tokenA {
		tokenARule = append(tokenARule, tokenAItem)
	}
	var tokenBRule []interface{}
	for _, tokenBItem := range tokenB {
		tokenBRule = append(tokenBRule, tokenBItem)
	}

	logs, sub, err := _Native.contract.WatchLogs(opts, "UpdatePair", tokenARule, tokenBRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeUpdatePair)
				if err := _Native.contract.UnpackLog(event, "UpdatePair", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatePair is a log parse operation binding the contract event 0x2d498c7d2a877bb8a3c6ec493a1743ddf467889264eb046fbfe2ca5ee2b4e752.
//
// Solidity: event UpdatePair(address indexed tokenA, address indexed tokenB, uint256 feeOld, uint256 feeNew, uint256 pricingModelIdOld, uint256 pricingModelIdNew)
func (_Native *NativeFilterer) ParseUpdatePair(log types.Log) (*NativeUpdatePair, error) {
	event := new(NativeUpdatePair)
	if err := _Native.contract.UnpackLog(event, "UpdatePair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Native contract.
type NativeUpgradedIterator struct {
	Event *NativeUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NativeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NativeUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NativeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeUpgraded represents a Upgraded event raised by the Native contract.
type NativeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Native *NativeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NativeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Native.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NativeUpgradedIterator{contract: _Native.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Native *NativeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NativeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Native.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeUpgraded)
				if err := _Native.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Native *NativeFilterer) ParseUpgraded(log types.Log) (*NativeUpgraded, error) {
	event := new(NativeUpgraded)
	if err := _Native.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
