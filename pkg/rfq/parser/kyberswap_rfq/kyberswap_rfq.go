// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kyberswaprfq

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

// IRFQOrderRFQ is an auto generated low-level Go binding around an user-defined struct.
type IRFQOrderRFQ struct {
	Info          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
}

// KyberswaprfqMetaData contains all meta data concerning the Kyberswaprfq contract.
var KyberswaprfqMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilledRFQ\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RescueFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"TransferAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"grantOrRevoke\",\"type\":\"bool\"}],\"name\":\"UpdateOperator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIMIT_ORDER_RFQ_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIRFQ.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIRFQ.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderRFQTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"invalidatorForOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"grantOrRevoke\",\"type\":\"bool\"}],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// KyberswaprfqABI is the input ABI used to generate the binding from.
// Deprecated: Use KyberswaprfqMetaData.ABI instead.
var KyberswaprfqABI = KyberswaprfqMetaData.ABI

// Kyberswaprfq is an auto generated Go binding around an Ethereum contract.
type Kyberswaprfq struct {
	KyberswaprfqCaller     // Read-only binding to the contract
	KyberswaprfqTransactor // Write-only binding to the contract
	KyberswaprfqFilterer   // Log filterer for contract events
}

// KyberswaprfqCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberswaprfqCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberswaprfqTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberswaprfqTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberswaprfqFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberswaprfqFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberswaprfqSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberswaprfqSession struct {
	Contract     *Kyberswaprfq     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberswaprfqCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberswaprfqCallerSession struct {
	Contract *KyberswaprfqCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KyberswaprfqTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberswaprfqTransactorSession struct {
	Contract     *KyberswaprfqTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KyberswaprfqRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberswaprfqRaw struct {
	Contract *Kyberswaprfq // Generic contract binding to access the raw methods on
}

// KyberswaprfqCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberswaprfqCallerRaw struct {
	Contract *KyberswaprfqCaller // Generic read-only contract binding to access the raw methods on
}

// KyberswaprfqTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberswaprfqTransactorRaw struct {
	Contract *KyberswaprfqTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberswaprfq creates a new instance of Kyberswaprfq, bound to a specific deployed contract.
func NewKyberswaprfq(address common.Address, backend bind.ContractBackend) (*Kyberswaprfq, error) {
	contract, err := bindKyberswaprfq(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kyberswaprfq{KyberswaprfqCaller: KyberswaprfqCaller{contract: contract}, KyberswaprfqTransactor: KyberswaprfqTransactor{contract: contract}, KyberswaprfqFilterer: KyberswaprfqFilterer{contract: contract}}, nil
}

// NewKyberswaprfqCaller creates a new read-only instance of Kyberswaprfq, bound to a specific deployed contract.
func NewKyberswaprfqCaller(address common.Address, caller bind.ContractCaller) (*KyberswaprfqCaller, error) {
	contract, err := bindKyberswaprfq(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberswaprfqCaller{contract: contract}, nil
}

// NewKyberswaprfqTransactor creates a new write-only instance of Kyberswaprfq, bound to a specific deployed contract.
func NewKyberswaprfqTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberswaprfqTransactor, error) {
	contract, err := bindKyberswaprfq(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberswaprfqTransactor{contract: contract}, nil
}

// NewKyberswaprfqFilterer creates a new log filterer instance of Kyberswaprfq, bound to a specific deployed contract.
func NewKyberswaprfqFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberswaprfqFilterer, error) {
	contract, err := bindKyberswaprfq(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberswaprfqFilterer{contract: contract}, nil
}

// bindKyberswaprfq binds a generic wrapper to an already deployed contract.
func bindKyberswaprfq(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KyberswaprfqMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kyberswaprfq *KyberswaprfqRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kyberswaprfq.Contract.KyberswaprfqCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kyberswaprfq *KyberswaprfqRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.KyberswaprfqTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kyberswaprfq *KyberswaprfqRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.KyberswaprfqTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kyberswaprfq *KyberswaprfqCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kyberswaprfq.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kyberswaprfq *KyberswaprfqTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kyberswaprfq *KyberswaprfqTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Kyberswaprfq *KyberswaprfqCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Kyberswaprfq.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Kyberswaprfq *KyberswaprfqSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Kyberswaprfq.Contract.DOMAINSEPARATOR(&_Kyberswaprfq.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Kyberswaprfq *KyberswaprfqCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Kyberswaprfq.Contract.DOMAINSEPARATOR(&_Kyberswaprfq.CallOpts)
}

// LIMITORDERRFQTYPEHASH is a free data retrieval call binding the contract method 0x06bf53d0.
//
// Solidity: function LIMIT_ORDER_RFQ_TYPEHASH() view returns(bytes32)
func (_Kyberswaprfq *KyberswaprfqCaller) LIMITORDERRFQTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Kyberswaprfq.contract.Call(opts, &out, "LIMIT_ORDER_RFQ_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LIMITORDERRFQTYPEHASH is a free data retrieval call binding the contract method 0x06bf53d0.
//
// Solidity: function LIMIT_ORDER_RFQ_TYPEHASH() view returns(bytes32)
func (_Kyberswaprfq *KyberswaprfqSession) LIMITORDERRFQTYPEHASH() ([32]byte, error) {
	return _Kyberswaprfq.Contract.LIMITORDERRFQTYPEHASH(&_Kyberswaprfq.CallOpts)
}

// LIMITORDERRFQTYPEHASH is a free data retrieval call binding the contract method 0x06bf53d0.
//
// Solidity: function LIMIT_ORDER_RFQ_TYPEHASH() view returns(bytes32)
func (_Kyberswaprfq *KyberswaprfqCallerSession) LIMITORDERRFQTYPEHASH() ([32]byte, error) {
	return _Kyberswaprfq.Contract.LIMITORDERRFQTYPEHASH(&_Kyberswaprfq.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Kyberswaprfq *KyberswaprfqCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kyberswaprfq.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Kyberswaprfq *KyberswaprfqSession) Admin() (common.Address, error) {
	return _Kyberswaprfq.Contract.Admin(&_Kyberswaprfq.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Kyberswaprfq *KyberswaprfqCallerSession) Admin() (common.Address, error) {
	return _Kyberswaprfq.Contract.Admin(&_Kyberswaprfq.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Kyberswaprfq *KyberswaprfqCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Kyberswaprfq.contract.Call(opts, &out, "eip712Domain")

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
func (_Kyberswaprfq *KyberswaprfqSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Kyberswaprfq.Contract.Eip712Domain(&_Kyberswaprfq.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Kyberswaprfq *KyberswaprfqCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Kyberswaprfq.Contract.Eip712Domain(&_Kyberswaprfq.CallOpts)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Kyberswaprfq *KyberswaprfqCaller) InvalidatorForOrderRFQ(opts *bind.CallOpts, maker common.Address, slot *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kyberswaprfq.contract.Call(opts, &out, "invalidatorForOrderRFQ", maker, slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Kyberswaprfq *KyberswaprfqSession) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _Kyberswaprfq.Contract.InvalidatorForOrderRFQ(&_Kyberswaprfq.CallOpts, maker, slot)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_Kyberswaprfq *KyberswaprfqCallerSession) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _Kyberswaprfq.Contract.InvalidatorForOrderRFQ(&_Kyberswaprfq.CallOpts, maker, slot)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Kyberswaprfq *KyberswaprfqCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Kyberswaprfq.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Kyberswaprfq *KyberswaprfqSession) Operators(arg0 common.Address) (bool, error) {
	return _Kyberswaprfq.Contract.Operators(&_Kyberswaprfq.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Kyberswaprfq *KyberswaprfqCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _Kyberswaprfq.Contract.Operators(&_Kyberswaprfq.CallOpts, arg0)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Kyberswaprfq *KyberswaprfqTransactor) CancelOrderRFQ(opts *bind.TransactOpts, orderInfo *big.Int) (*types.Transaction, error) {
	return _Kyberswaprfq.contract.Transact(opts, "cancelOrderRFQ", orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Kyberswaprfq *KyberswaprfqSession) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.CancelOrderRFQ(&_Kyberswaprfq.TransactOpts, orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_Kyberswaprfq *KyberswaprfqTransactorSession) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.CancelOrderRFQ(&_Kyberswaprfq.TransactOpts, orderInfo)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0xd0a3b665.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount) payable returns(uint256, uint256)
func (_Kyberswaprfq *KyberswaprfqTransactor) FillOrderRFQ(opts *bind.TransactOpts, order IRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int) (*types.Transaction, error) {
	return _Kyberswaprfq.contract.Transact(opts, "fillOrderRFQ", order, signature, makingAmount, takingAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0xd0a3b665.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount) payable returns(uint256, uint256)
func (_Kyberswaprfq *KyberswaprfqSession) FillOrderRFQ(order IRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.FillOrderRFQ(&_Kyberswaprfq.TransactOpts, order, signature, makingAmount, takingAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0xd0a3b665.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount) payable returns(uint256, uint256)
func (_Kyberswaprfq *KyberswaprfqTransactorSession) FillOrderRFQ(order IRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.FillOrderRFQ(&_Kyberswaprfq.TransactOpts, order, signature, makingAmount, takingAmount)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0xbaba5855.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, address target) payable returns(uint256, uint256)
func (_Kyberswaprfq *KyberswaprfqTransactor) FillOrderRFQTo(opts *bind.TransactOpts, order IRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Kyberswaprfq.contract.Transact(opts, "fillOrderRFQTo", order, signature, makingAmount, takingAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0xbaba5855.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, address target) payable returns(uint256, uint256)
func (_Kyberswaprfq *KyberswaprfqSession) FillOrderRFQTo(order IRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.FillOrderRFQTo(&_Kyberswaprfq.TransactOpts, order, signature, makingAmount, takingAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0xbaba5855.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, address target) payable returns(uint256, uint256)
func (_Kyberswaprfq *KyberswaprfqTransactorSession) FillOrderRFQTo(order IRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.FillOrderRFQTo(&_Kyberswaprfq.TransactOpts, order, signature, makingAmount, takingAmount, target)
}

// RescueFund is a paid mutator transaction binding the contract method 0x6e261df1.
//
// Solidity: function rescueFund(address token, uint256 amount) returns()
func (_Kyberswaprfq *KyberswaprfqTransactor) RescueFund(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kyberswaprfq.contract.Transact(opts, "rescueFund", token, amount)
}

// RescueFund is a paid mutator transaction binding the contract method 0x6e261df1.
//
// Solidity: function rescueFund(address token, uint256 amount) returns()
func (_Kyberswaprfq *KyberswaprfqSession) RescueFund(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.RescueFund(&_Kyberswaprfq.TransactOpts, token, amount)
}

// RescueFund is a paid mutator transaction binding the contract method 0x6e261df1.
//
// Solidity: function rescueFund(address token, uint256 amount) returns()
func (_Kyberswaprfq *KyberswaprfqTransactorSession) RescueFund(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.RescueFund(&_Kyberswaprfq.TransactOpts, token, amount)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address _admin) returns()
func (_Kyberswaprfq *KyberswaprfqTransactor) TransferAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Kyberswaprfq.contract.Transact(opts, "transferAdmin", _admin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address _admin) returns()
func (_Kyberswaprfq *KyberswaprfqSession) TransferAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.TransferAdmin(&_Kyberswaprfq.TransactOpts, _admin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(address _admin) returns()
func (_Kyberswaprfq *KyberswaprfqTransactorSession) TransferAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.TransferAdmin(&_Kyberswaprfq.TransactOpts, _admin)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0x6d44a3b2.
//
// Solidity: function updateOperator(address user, bool grantOrRevoke) returns()
func (_Kyberswaprfq *KyberswaprfqTransactor) UpdateOperator(opts *bind.TransactOpts, user common.Address, grantOrRevoke bool) (*types.Transaction, error) {
	return _Kyberswaprfq.contract.Transact(opts, "updateOperator", user, grantOrRevoke)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0x6d44a3b2.
//
// Solidity: function updateOperator(address user, bool grantOrRevoke) returns()
func (_Kyberswaprfq *KyberswaprfqSession) UpdateOperator(user common.Address, grantOrRevoke bool) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.UpdateOperator(&_Kyberswaprfq.TransactOpts, user, grantOrRevoke)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0x6d44a3b2.
//
// Solidity: function updateOperator(address user, bool grantOrRevoke) returns()
func (_Kyberswaprfq *KyberswaprfqTransactorSession) UpdateOperator(user common.Address, grantOrRevoke bool) (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.UpdateOperator(&_Kyberswaprfq.TransactOpts, user, grantOrRevoke)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kyberswaprfq *KyberswaprfqTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswaprfq.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kyberswaprfq *KyberswaprfqSession) Receive() (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.Receive(&_Kyberswaprfq.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kyberswaprfq *KyberswaprfqTransactorSession) Receive() (*types.Transaction, error) {
	return _Kyberswaprfq.Contract.Receive(&_Kyberswaprfq.TransactOpts)
}

// KyberswaprfqEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Kyberswaprfq contract.
type KyberswaprfqEIP712DomainChangedIterator struct {
	Event *KyberswaprfqEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *KyberswaprfqEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswaprfqEIP712DomainChanged)
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
		it.Event = new(KyberswaprfqEIP712DomainChanged)
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
func (it *KyberswaprfqEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswaprfqEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswaprfqEIP712DomainChanged represents a EIP712DomainChanged event raised by the Kyberswaprfq contract.
type KyberswaprfqEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Kyberswaprfq *KyberswaprfqFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*KyberswaprfqEIP712DomainChangedIterator, error) {

	logs, sub, err := _Kyberswaprfq.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &KyberswaprfqEIP712DomainChangedIterator{contract: _Kyberswaprfq.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Kyberswaprfq *KyberswaprfqFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *KyberswaprfqEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Kyberswaprfq.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswaprfqEIP712DomainChanged)
				if err := _Kyberswaprfq.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_Kyberswaprfq *KyberswaprfqFilterer) ParseEIP712DomainChanged(log types.Log) (*KyberswaprfqEIP712DomainChanged, error) {
	event := new(KyberswaprfqEIP712DomainChanged)
	if err := _Kyberswaprfq.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswaprfqOrderFilledRFQIterator is returned from FilterOrderFilledRFQ and is used to iterate over the raw logs and unpacked data for OrderFilledRFQ events raised by the Kyberswaprfq contract.
type KyberswaprfqOrderFilledRFQIterator struct {
	Event *KyberswaprfqOrderFilledRFQ // Event containing the contract specifics and raw log

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
func (it *KyberswaprfqOrderFilledRFQIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswaprfqOrderFilledRFQ)
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
		it.Event = new(KyberswaprfqOrderFilledRFQ)
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
func (it *KyberswaprfqOrderFilledRFQIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswaprfqOrderFilledRFQIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswaprfqOrderFilledRFQ represents a OrderFilledRFQ event raised by the Kyberswaprfq contract.
type KyberswaprfqOrderFilledRFQ struct {
	OrderHash    [32]byte
	Maker        common.Address
	Taker        common.Address
	MakerAsset   common.Address
	TakerAsset   common.Address
	MakingAmount *big.Int
	TakingAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderFilledRFQ is a free log retrieval operation binding the contract event 0x71cea972ff7e532f1478ece362ee7d7e5be0654905843c97ecca781a3a0c9d4a.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, address indexed maker, address indexed taker, address makerAsset, address takerAsset, uint256 makingAmount, uint256 takingAmount)
func (_Kyberswaprfq *KyberswaprfqFilterer) FilterOrderFilledRFQ(opts *bind.FilterOpts, maker []common.Address, taker []common.Address) (*KyberswaprfqOrderFilledRFQIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Kyberswaprfq.contract.FilterLogs(opts, "OrderFilledRFQ", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &KyberswaprfqOrderFilledRFQIterator{contract: _Kyberswaprfq.contract, event: "OrderFilledRFQ", logs: logs, sub: sub}, nil
}

// WatchOrderFilledRFQ is a free log subscription operation binding the contract event 0x71cea972ff7e532f1478ece362ee7d7e5be0654905843c97ecca781a3a0c9d4a.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, address indexed maker, address indexed taker, address makerAsset, address takerAsset, uint256 makingAmount, uint256 takingAmount)
func (_Kyberswaprfq *KyberswaprfqFilterer) WatchOrderFilledRFQ(opts *bind.WatchOpts, sink chan<- *KyberswaprfqOrderFilledRFQ, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Kyberswaprfq.contract.WatchLogs(opts, "OrderFilledRFQ", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswaprfqOrderFilledRFQ)
				if err := _Kyberswaprfq.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
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

// ParseOrderFilledRFQ is a log parse operation binding the contract event 0x71cea972ff7e532f1478ece362ee7d7e5be0654905843c97ecca781a3a0c9d4a.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, address indexed maker, address indexed taker, address makerAsset, address takerAsset, uint256 makingAmount, uint256 takingAmount)
func (_Kyberswaprfq *KyberswaprfqFilterer) ParseOrderFilledRFQ(log types.Log) (*KyberswaprfqOrderFilledRFQ, error) {
	event := new(KyberswaprfqOrderFilledRFQ)
	if err := _Kyberswaprfq.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswaprfqRescueFundIterator is returned from FilterRescueFund and is used to iterate over the raw logs and unpacked data for RescueFund events raised by the Kyberswaprfq contract.
type KyberswaprfqRescueFundIterator struct {
	Event *KyberswaprfqRescueFund // Event containing the contract specifics and raw log

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
func (it *KyberswaprfqRescueFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswaprfqRescueFund)
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
		it.Event = new(KyberswaprfqRescueFund)
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
func (it *KyberswaprfqRescueFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswaprfqRescueFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswaprfqRescueFund represents a RescueFund event raised by the Kyberswaprfq contract.
type KyberswaprfqRescueFund struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRescueFund is a free log retrieval operation binding the contract event 0x125fbdbcf3a7becd38f03d012558074dc0f2dbf28942defd577fea7832369872.
//
// Solidity: event RescueFund(address indexed token, uint256 amount)
func (_Kyberswaprfq *KyberswaprfqFilterer) FilterRescueFund(opts *bind.FilterOpts, token []common.Address) (*KyberswaprfqRescueFundIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Kyberswaprfq.contract.FilterLogs(opts, "RescueFund", tokenRule)
	if err != nil {
		return nil, err
	}
	return &KyberswaprfqRescueFundIterator{contract: _Kyberswaprfq.contract, event: "RescueFund", logs: logs, sub: sub}, nil
}

// WatchRescueFund is a free log subscription operation binding the contract event 0x125fbdbcf3a7becd38f03d012558074dc0f2dbf28942defd577fea7832369872.
//
// Solidity: event RescueFund(address indexed token, uint256 amount)
func (_Kyberswaprfq *KyberswaprfqFilterer) WatchRescueFund(opts *bind.WatchOpts, sink chan<- *KyberswaprfqRescueFund, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Kyberswaprfq.contract.WatchLogs(opts, "RescueFund", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswaprfqRescueFund)
				if err := _Kyberswaprfq.contract.UnpackLog(event, "RescueFund", log); err != nil {
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

// ParseRescueFund is a log parse operation binding the contract event 0x125fbdbcf3a7becd38f03d012558074dc0f2dbf28942defd577fea7832369872.
//
// Solidity: event RescueFund(address indexed token, uint256 amount)
func (_Kyberswaprfq *KyberswaprfqFilterer) ParseRescueFund(log types.Log) (*KyberswaprfqRescueFund, error) {
	event := new(KyberswaprfqRescueFund)
	if err := _Kyberswaprfq.contract.UnpackLog(event, "RescueFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswaprfqTransferAdminIterator is returned from FilterTransferAdmin and is used to iterate over the raw logs and unpacked data for TransferAdmin events raised by the Kyberswaprfq contract.
type KyberswaprfqTransferAdminIterator struct {
	Event *KyberswaprfqTransferAdmin // Event containing the contract specifics and raw log

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
func (it *KyberswaprfqTransferAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswaprfqTransferAdmin)
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
		it.Event = new(KyberswaprfqTransferAdmin)
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
func (it *KyberswaprfqTransferAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswaprfqTransferAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswaprfqTransferAdmin represents a TransferAdmin event raised by the Kyberswaprfq contract.
type KyberswaprfqTransferAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferAdmin is a free log retrieval operation binding the contract event 0xda7b0a7bc965abdec8a1a995575a891838264c2968e14bd456c5391827b7aa30.
//
// Solidity: event TransferAdmin(address indexed admin)
func (_Kyberswaprfq *KyberswaprfqFilterer) FilterTransferAdmin(opts *bind.FilterOpts, admin []common.Address) (*KyberswaprfqTransferAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Kyberswaprfq.contract.FilterLogs(opts, "TransferAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &KyberswaprfqTransferAdminIterator{contract: _Kyberswaprfq.contract, event: "TransferAdmin", logs: logs, sub: sub}, nil
}

// WatchTransferAdmin is a free log subscription operation binding the contract event 0xda7b0a7bc965abdec8a1a995575a891838264c2968e14bd456c5391827b7aa30.
//
// Solidity: event TransferAdmin(address indexed admin)
func (_Kyberswaprfq *KyberswaprfqFilterer) WatchTransferAdmin(opts *bind.WatchOpts, sink chan<- *KyberswaprfqTransferAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Kyberswaprfq.contract.WatchLogs(opts, "TransferAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswaprfqTransferAdmin)
				if err := _Kyberswaprfq.contract.UnpackLog(event, "TransferAdmin", log); err != nil {
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

// ParseTransferAdmin is a log parse operation binding the contract event 0xda7b0a7bc965abdec8a1a995575a891838264c2968e14bd456c5391827b7aa30.
//
// Solidity: event TransferAdmin(address indexed admin)
func (_Kyberswaprfq *KyberswaprfqFilterer) ParseTransferAdmin(log types.Log) (*KyberswaprfqTransferAdmin, error) {
	event := new(KyberswaprfqTransferAdmin)
	if err := _Kyberswaprfq.contract.UnpackLog(event, "TransferAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswaprfqUpdateOperatorIterator is returned from FilterUpdateOperator and is used to iterate over the raw logs and unpacked data for UpdateOperator events raised by the Kyberswaprfq contract.
type KyberswaprfqUpdateOperatorIterator struct {
	Event *KyberswaprfqUpdateOperator // Event containing the contract specifics and raw log

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
func (it *KyberswaprfqUpdateOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswaprfqUpdateOperator)
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
		it.Event = new(KyberswaprfqUpdateOperator)
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
func (it *KyberswaprfqUpdateOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswaprfqUpdateOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswaprfqUpdateOperator represents a UpdateOperator event raised by the Kyberswaprfq contract.
type KyberswaprfqUpdateOperator struct {
	User          common.Address
	GrantOrRevoke bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateOperator is a free log retrieval operation binding the contract event 0x2ee52be9d342458b3d25e07faada7ff9bc06723b4aa24edb6321ac1316b8a9dd.
//
// Solidity: event UpdateOperator(address indexed user, bool grantOrRevoke)
func (_Kyberswaprfq *KyberswaprfqFilterer) FilterUpdateOperator(opts *bind.FilterOpts, user []common.Address) (*KyberswaprfqUpdateOperatorIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Kyberswaprfq.contract.FilterLogs(opts, "UpdateOperator", userRule)
	if err != nil {
		return nil, err
	}
	return &KyberswaprfqUpdateOperatorIterator{contract: _Kyberswaprfq.contract, event: "UpdateOperator", logs: logs, sub: sub}, nil
}

// WatchUpdateOperator is a free log subscription operation binding the contract event 0x2ee52be9d342458b3d25e07faada7ff9bc06723b4aa24edb6321ac1316b8a9dd.
//
// Solidity: event UpdateOperator(address indexed user, bool grantOrRevoke)
func (_Kyberswaprfq *KyberswaprfqFilterer) WatchUpdateOperator(opts *bind.WatchOpts, sink chan<- *KyberswaprfqUpdateOperator, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Kyberswaprfq.contract.WatchLogs(opts, "UpdateOperator", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswaprfqUpdateOperator)
				if err := _Kyberswaprfq.contract.UnpackLog(event, "UpdateOperator", log); err != nil {
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

// ParseUpdateOperator is a log parse operation binding the contract event 0x2ee52be9d342458b3d25e07faada7ff9bc06723b4aa24edb6321ac1316b8a9dd.
//
// Solidity: event UpdateOperator(address indexed user, bool grantOrRevoke)
func (_Kyberswaprfq *KyberswaprfqFilterer) ParseUpdateOperator(log types.Log) (*KyberswaprfqUpdateOperator, error) {
	event := new(KyberswaprfqUpdateOperator)
	if err := _Kyberswaprfq.contract.UnpackLog(event, "UpdateOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
