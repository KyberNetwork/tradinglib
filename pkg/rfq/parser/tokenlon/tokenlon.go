// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenlon

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

// FillOrderMetaData contains all meta data concerning the FillOrder contract.
var FillOrderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"takerAssetAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"makerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"makerAssetAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiverAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settleAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeFactor\",\"type\":\"uint16\"}],\"name\":\"FillOrder\",\"type\":\"event\"}]",
}

// FillOrderABI is the input ABI used to generate the binding from.
// Deprecated: Use FillOrderMetaData.ABI instead.
var FillOrderABI = FillOrderMetaData.ABI

// FillOrder is an auto generated Go binding around an Ethereum contract.
type FillOrder struct {
	FillOrderCaller     // Read-only binding to the contract
	FillOrderTransactor // Write-only binding to the contract
	FillOrderFilterer   // Log filterer for contract events
}

// FillOrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type FillOrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FillOrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FillOrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FillOrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FillOrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FillOrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FillOrderSession struct {
	Contract     *FillOrder        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FillOrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FillOrderCallerSession struct {
	Contract *FillOrderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FillOrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FillOrderTransactorSession struct {
	Contract     *FillOrderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FillOrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type FillOrderRaw struct {
	Contract *FillOrder // Generic contract binding to access the raw methods on
}

// FillOrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FillOrderCallerRaw struct {
	Contract *FillOrderCaller // Generic read-only contract binding to access the raw methods on
}

// FillOrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FillOrderTransactorRaw struct {
	Contract *FillOrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFillOrder creates a new instance of FillOrder, bound to a specific deployed contract.
func NewFillOrder(address common.Address, backend bind.ContractBackend) (*FillOrder, error) {
	contract, err := bindFillOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FillOrder{FillOrderCaller: FillOrderCaller{contract: contract}, FillOrderTransactor: FillOrderTransactor{contract: contract}, FillOrderFilterer: FillOrderFilterer{contract: contract}}, nil
}

// NewFillOrderCaller creates a new read-only instance of FillOrder, bound to a specific deployed contract.
func NewFillOrderCaller(address common.Address, caller bind.ContractCaller) (*FillOrderCaller, error) {
	contract, err := bindFillOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FillOrderCaller{contract: contract}, nil
}

// NewFillOrderTransactor creates a new write-only instance of FillOrder, bound to a specific deployed contract.
func NewFillOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*FillOrderTransactor, error) {
	contract, err := bindFillOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FillOrderTransactor{contract: contract}, nil
}

// NewFillOrderFilterer creates a new log filterer instance of FillOrder, bound to a specific deployed contract.
func NewFillOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*FillOrderFilterer, error) {
	contract, err := bindFillOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FillOrderFilterer{contract: contract}, nil
}

// bindFillOrder binds a generic wrapper to an already deployed contract.
func bindFillOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FillOrderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FillOrder *FillOrderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FillOrder.Contract.FillOrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FillOrder *FillOrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FillOrder.Contract.FillOrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FillOrder *FillOrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FillOrder.Contract.FillOrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FillOrder *FillOrderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FillOrder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FillOrder *FillOrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FillOrder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FillOrder *FillOrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FillOrder.Contract.contract.Transact(opts, method, params...)
}

// FillOrderFillOrderIterator is returned from FilterFillOrder and is used to iterate over the raw logs and unpacked data for FillOrder events raised by the FillOrder contract.
type FillOrderFillOrderIterator struct {
	Event *FillOrderFillOrder // Event containing the contract specifics and raw log

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
func (it *FillOrderFillOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FillOrderFillOrder)
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
		it.Event = new(FillOrderFillOrder)
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
func (it *FillOrderFillOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FillOrderFillOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FillOrderFillOrder represents a FillOrder event raised by the FillOrder contract.
type FillOrderFillOrder struct {
	Source           string
	TransactionHash  [32]byte
	OrderHash        [32]byte
	UserAddr         common.Address
	TakerAssetAddr   common.Address
	TakerAssetAmount *big.Int
	MakerAddr        common.Address
	MakerAssetAddr   common.Address
	MakerAssetAmount *big.Int
	ReceiverAddr     common.Address
	SettleAmount     *big.Int
	FeeFactor        uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFillOrder is a free log retrieval operation binding the contract event 0x75d58426b26ab641a6a6a46f12fe35e17c570a1cd264c7248a73d90e3a8682ff.
//
// Solidity: event FillOrder(string source, bytes32 indexed transactionHash, bytes32 indexed orderHash, address indexed userAddr, address takerAssetAddr, uint256 takerAssetAmount, address makerAddr, address makerAssetAddr, uint256 makerAssetAmount, address receiverAddr, uint256 settleAmount, uint16 feeFactor)
func (_FillOrder *FillOrderFilterer) FilterFillOrder(opts *bind.FilterOpts, transactionHash [][32]byte, orderHash [][32]byte, userAddr []common.Address) (*FillOrderFillOrderIterator, error) {

	var transactionHashRule []interface{}
	for _, transactionHashItem := range transactionHash {
		transactionHashRule = append(transactionHashRule, transactionHashItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _FillOrder.contract.FilterLogs(opts, "FillOrder", transactionHashRule, orderHashRule, userAddrRule)
	if err != nil {
		return nil, err
	}
	return &FillOrderFillOrderIterator{contract: _FillOrder.contract, event: "FillOrder", logs: logs, sub: sub}, nil
}

// WatchFillOrder is a free log subscription operation binding the contract event 0x75d58426b26ab641a6a6a46f12fe35e17c570a1cd264c7248a73d90e3a8682ff.
//
// Solidity: event FillOrder(string source, bytes32 indexed transactionHash, bytes32 indexed orderHash, address indexed userAddr, address takerAssetAddr, uint256 takerAssetAmount, address makerAddr, address makerAssetAddr, uint256 makerAssetAmount, address receiverAddr, uint256 settleAmount, uint16 feeFactor)
func (_FillOrder *FillOrderFilterer) WatchFillOrder(opts *bind.WatchOpts, sink chan<- *FillOrderFillOrder, transactionHash [][32]byte, orderHash [][32]byte, userAddr []common.Address) (event.Subscription, error) {

	var transactionHashRule []interface{}
	for _, transactionHashItem := range transactionHash {
		transactionHashRule = append(transactionHashRule, transactionHashItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _FillOrder.contract.WatchLogs(opts, "FillOrder", transactionHashRule, orderHashRule, userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FillOrderFillOrder)
				if err := _FillOrder.contract.UnpackLog(event, "FillOrder", log); err != nil {
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

// ParseFillOrder is a log parse operation binding the contract event 0x75d58426b26ab641a6a6a46f12fe35e17c570a1cd264c7248a73d90e3a8682ff.
//
// Solidity: event FillOrder(string source, bytes32 indexed transactionHash, bytes32 indexed orderHash, address indexed userAddr, address takerAssetAddr, uint256 takerAssetAmount, address makerAddr, address makerAssetAddr, uint256 makerAssetAmount, address receiverAddr, uint256 settleAmount, uint16 feeFactor)
func (_FillOrder *FillOrderFilterer) ParseFillOrder(log types.Log) (*FillOrderFillOrder, error) {
	event := new(FillOrderFillOrder)
	if err := _FillOrder.contract.UnpackLog(event, "FillOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
