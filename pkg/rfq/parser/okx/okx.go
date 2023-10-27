// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package okx

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

// OrderRFQLibOrderRFQ is an auto generated low-level Go binding around an user-defined struct.
type OrderRFQLibOrderRFQ struct {
	Info          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
	Settler       common.Address
}

// OkxMetaData contains all meta data concerning the Okx contract.
var OkxMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilledRFQ\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalMask\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQCompact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filledMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderRFQTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filledMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderRFQToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"invalidatorForOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OkxABI is the input ABI used to generate the binding from.
// Deprecated: Use OkxMetaData.ABI instead.
var OkxABI = OkxMetaData.ABI

// Okx is an auto generated Go binding around an Ethereum contract.
type Okx struct {
	OkxCaller     // Read-only binding to the contract
	OkxTransactor // Write-only binding to the contract
	OkxFilterer   // Log filterer for contract events
}

// OkxCaller is an auto generated read-only Go binding around an Ethereum contract.
type OkxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OkxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OkxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OkxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OkxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OkxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OkxSession struct {
	Contract     *Okx              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OkxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OkxCallerSession struct {
	Contract *OkxCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OkxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OkxTransactorSession struct {
	Contract     *OkxTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OkxRaw is an auto generated low-level Go binding around an Ethereum contract.
type OkxRaw struct {
	Contract *Okx // Generic contract binding to access the raw methods on
}

// OkxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OkxCallerRaw struct {
	Contract *OkxCaller // Generic read-only contract binding to access the raw methods on
}

// OkxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OkxTransactorRaw struct {
	Contract *OkxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOkx creates a new instance of Okx, bound to a specific deployed contract.
func NewOkx(address common.Address, backend bind.ContractBackend) (*Okx, error) {
	contract, err := bindOkx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Okx{OkxCaller: OkxCaller{contract: contract}, OkxTransactor: OkxTransactor{contract: contract}, OkxFilterer: OkxFilterer{contract: contract}}, nil
}

// NewOkxCaller creates a new read-only instance of Okx, bound to a specific deployed contract.
func NewOkxCaller(address common.Address, caller bind.ContractCaller) (*OkxCaller, error) {
	contract, err := bindOkx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OkxCaller{contract: contract}, nil
}

// NewOkxTransactor creates a new write-only instance of Okx, bound to a specific deployed contract.
func NewOkxTransactor(address common.Address, transactor bind.ContractTransactor) (*OkxTransactor, error) {
	contract, err := bindOkx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OkxTransactor{contract: contract}, nil
}

// NewOkxFilterer creates a new log filterer instance of Okx, bound to a specific deployed contract.
func NewOkxFilterer(address common.Address, filterer bind.ContractFilterer) (*OkxFilterer, error) {
	contract, err := bindOkx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OkxFilterer{contract: contract}, nil
}

// bindOkx binds a generic wrapper to an already deployed contract.
func bindOkx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OkxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Okx *OkxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Okx.Contract.OkxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Okx *OkxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Okx.Contract.OkxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Okx *OkxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Okx.Contract.OkxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Okx *OkxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Okx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Okx *OkxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Okx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Okx *OkxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Okx.Contract.contract.Transact(opts, method, params...)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Okx *OkxCaller) InvalidatorForOrderRFQ(opts *bind.CallOpts, maker common.Address, slot *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Okx.contract.Call(opts, &out, "invalidatorForOrderRFQ", maker, slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Okx *OkxSession) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _Okx.Contract.InvalidatorForOrderRFQ(&_Okx.CallOpts, maker, slot)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Okx *OkxCallerSession) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _Okx.Contract.InvalidatorForOrderRFQ(&_Okx.CallOpts, maker, slot)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Okx *OkxTransactor) CancelOrderRFQ(opts *bind.TransactOpts, orderInfo *big.Int) (*types.Transaction, error) {
	return _Okx.contract.Transact(opts, "cancelOrderRFQ", orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Okx *OkxSession) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _Okx.Contract.CancelOrderRFQ(&_Okx.TransactOpts, orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Okx *OkxTransactorSession) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _Okx.Contract.CancelOrderRFQ(&_Okx.TransactOpts, orderInfo)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_Okx *OkxTransactor) CancelOrderRFQ0(opts *bind.TransactOpts, orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Okx.contract.Transact(opts, "cancelOrderRFQ0", orderInfo, additionalMask)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_Okx *OkxSession) CancelOrderRFQ0(orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Okx.Contract.CancelOrderRFQ0(&_Okx.TransactOpts, orderInfo, additionalMask)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_Okx *OkxTransactorSession) CancelOrderRFQ0(orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _Okx.Contract.CancelOrderRFQ0(&_Okx.TransactOpts, orderInfo, additionalMask)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0xe389f382.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256,address) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_Okx *OkxTransactor) FillOrderRFQ(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Okx.contract.Transact(opts, "fillOrderRFQ", order, signature, flagsAndAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0xe389f382.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256,address) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_Okx *OkxSession) FillOrderRFQ(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Okx.Contract.FillOrderRFQ(&_Okx.TransactOpts, order, signature, flagsAndAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0xe389f382.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256,address) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_Okx *OkxTransactorSession) FillOrderRFQ(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Okx.Contract.FillOrderRFQ(&_Okx.TransactOpts, order, signature, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x2cf240c8.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256,address) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Okx *OkxTransactor) FillOrderRFQCompact(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Okx.contract.Transact(opts, "fillOrderRFQCompact", order, r, vs, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x2cf240c8.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256,address) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Okx *OkxSession) FillOrderRFQCompact(order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Okx.Contract.FillOrderRFQCompact(&_Okx.TransactOpts, order, r, vs, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x2cf240c8.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256,address) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Okx *OkxTransactorSession) FillOrderRFQCompact(order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _Okx.Contract.FillOrderRFQCompact(&_Okx.TransactOpts, order, r, vs, flagsAndAmount)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x2db4e446.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256,address) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Okx *OkxTransactor) FillOrderRFQTo(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Okx.contract.Transact(opts, "fillOrderRFQTo", order, signature, flagsAndAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x2db4e446.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256,address) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Okx *OkxSession) FillOrderRFQTo(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Okx.Contract.FillOrderRFQTo(&_Okx.TransactOpts, order, signature, flagsAndAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x2db4e446.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256,address) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_Okx *OkxTransactorSession) FillOrderRFQTo(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Okx.Contract.FillOrderRFQTo(&_Okx.TransactOpts, order, signature, flagsAndAmount, target)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x6f6d26e9.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256,address) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Okx *OkxTransactor) FillOrderRFQToWithPermit(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Okx.contract.Transact(opts, "fillOrderRFQToWithPermit", order, signature, flagsAndAmount, target, permit)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x6f6d26e9.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256,address) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Okx *OkxSession) FillOrderRFQToWithPermit(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Okx.Contract.FillOrderRFQToWithPermit(&_Okx.TransactOpts, order, signature, flagsAndAmount, target, permit)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x6f6d26e9.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256,address) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_Okx *OkxTransactorSession) FillOrderRFQToWithPermit(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _Okx.Contract.FillOrderRFQToWithPermit(&_Okx.TransactOpts, order, signature, flagsAndAmount, target, permit)
}

// OkxOrderFilledRFQIterator is returned from FilterOrderFilledRFQ and is used to iterate over the raw logs and unpacked data for OrderFilledRFQ events raised by the Okx contract.
type OkxOrderFilledRFQIterator struct {
	Event *OkxOrderFilledRFQ // Event containing the contract specifics and raw log

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
func (it *OkxOrderFilledRFQIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkxOrderFilledRFQ)
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
		it.Event = new(OkxOrderFilledRFQ)
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
func (it *OkxOrderFilledRFQIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkxOrderFilledRFQIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkxOrderFilledRFQ represents a OrderFilledRFQ event raised by the Okx contract.
type OkxOrderFilledRFQ struct {
	OrderHash    [32]byte
	MakingAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderFilledRFQ is a free log retrieval operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_Okx *OkxFilterer) FilterOrderFilledRFQ(opts *bind.FilterOpts) (*OkxOrderFilledRFQIterator, error) {

	logs, sub, err := _Okx.contract.FilterLogs(opts, "OrderFilledRFQ")
	if err != nil {
		return nil, err
	}
	return &OkxOrderFilledRFQIterator{contract: _Okx.contract, event: "OrderFilledRFQ", logs: logs, sub: sub}, nil
}

// WatchOrderFilledRFQ is a free log subscription operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_Okx *OkxFilterer) WatchOrderFilledRFQ(opts *bind.WatchOpts, sink chan<- *OkxOrderFilledRFQ) (event.Subscription, error) {

	logs, sub, err := _Okx.contract.WatchLogs(opts, "OrderFilledRFQ")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkxOrderFilledRFQ)
				if err := _Okx.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
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

// ParseOrderFilledRFQ is a log parse operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_Okx *OkxFilterer) ParseOrderFilledRFQ(log types.Log) (*OkxOrderFilledRFQ, error) {
	event := new(OkxOrderFilledRFQ)
	if err := _Okx.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
