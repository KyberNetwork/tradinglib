// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zxotc

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
)

// OtcOrderFilledMetaData contains all meta data concerning the OtcOrderFilled contract.
var OtcOrderFilledMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"makerToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"takerToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"makerTokenFilledAmount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"takerTokenFilledAmount\",\"type\":\"uint128\"}],\"name\":\"OtcOrderFilled\",\"type\":\"event\"}]",
}

// OtcOrderFilledABI is the input ABI used to generate the binding from.
// Deprecated: Use OtcOrderFilledMetaData.ABI instead.
var OtcOrderFilledABI = OtcOrderFilledMetaData.ABI

// OtcOrderFilled is an auto generated Go binding around an Ethereum contract.
type OtcOrderFilled struct {
	OtcOrderFilledCaller     // Read-only binding to the contract
	OtcOrderFilledTransactor // Write-only binding to the contract
	OtcOrderFilledFilterer   // Log filterer for contract events
}

// OtcOrderFilledCaller is an auto generated read-only Go binding around an Ethereum contract.
type OtcOrderFilledCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtcOrderFilledTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OtcOrderFilledTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtcOrderFilledFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OtcOrderFilledFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OtcOrderFilledSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OtcOrderFilledSession struct {
	Contract     *OtcOrderFilled   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OtcOrderFilledCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OtcOrderFilledCallerSession struct {
	Contract *OtcOrderFilledCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OtcOrderFilledTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OtcOrderFilledTransactorSession struct {
	Contract     *OtcOrderFilledTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OtcOrderFilledRaw is an auto generated low-level Go binding around an Ethereum contract.
type OtcOrderFilledRaw struct {
	Contract *OtcOrderFilled // Generic contract binding to access the raw methods on
}

// OtcOrderFilledCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OtcOrderFilledCallerRaw struct {
	Contract *OtcOrderFilledCaller // Generic read-only contract binding to access the raw methods on
}

// OtcOrderFilledTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OtcOrderFilledTransactorRaw struct {
	Contract *OtcOrderFilledTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOtcOrderFilled creates a new instance of OtcOrderFilled, bound to a specific deployed contract.
func NewOtcOrderFilled(address common.Address, backend bind.ContractBackend) (*OtcOrderFilled, error) {
	contract, err := bindOtcOrderFilled(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OtcOrderFilled{OtcOrderFilledCaller: OtcOrderFilledCaller{contract: contract}, OtcOrderFilledTransactor: OtcOrderFilledTransactor{contract: contract}, OtcOrderFilledFilterer: OtcOrderFilledFilterer{contract: contract}}, nil
}

// NewOtcOrderFilledCaller creates a new read-only instance of OtcOrderFilled, bound to a specific deployed contract.
func NewOtcOrderFilledCaller(address common.Address, caller bind.ContractCaller) (*OtcOrderFilledCaller, error) {
	contract, err := bindOtcOrderFilled(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OtcOrderFilledCaller{contract: contract}, nil
}

// NewOtcOrderFilledTransactor creates a new write-only instance of OtcOrderFilled, bound to a specific deployed contract.
func NewOtcOrderFilledTransactor(address common.Address, transactor bind.ContractTransactor) (*OtcOrderFilledTransactor, error) {
	contract, err := bindOtcOrderFilled(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OtcOrderFilledTransactor{contract: contract}, nil
}

// NewOtcOrderFilledFilterer creates a new log filterer instance of OtcOrderFilled, bound to a specific deployed contract.
func NewOtcOrderFilledFilterer(address common.Address, filterer bind.ContractFilterer) (*OtcOrderFilledFilterer, error) {
	contract, err := bindOtcOrderFilled(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OtcOrderFilledFilterer{contract: contract}, nil
}

// bindOtcOrderFilled binds a generic wrapper to an already deployed contract.
func bindOtcOrderFilled(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OtcOrderFilledABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtcOrderFilled *OtcOrderFilledRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OtcOrderFilled.Contract.OtcOrderFilledCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtcOrderFilled *OtcOrderFilledRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtcOrderFilled.Contract.OtcOrderFilledTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtcOrderFilled *OtcOrderFilledRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtcOrderFilled.Contract.OtcOrderFilledTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OtcOrderFilled *OtcOrderFilledCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OtcOrderFilled.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OtcOrderFilled *OtcOrderFilledTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OtcOrderFilled.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OtcOrderFilled *OtcOrderFilledTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OtcOrderFilled.Contract.contract.Transact(opts, method, params...)
}

// OtcOrderFilledOtcOrderFilledIterator is returned from FilterOtcOrderFilled and is used to iterate over the raw logs and unpacked data for OtcOrderFilled events raised by the OtcOrderFilled contract.
type OtcOrderFilledOtcOrderFilledIterator struct {
	Event *OtcOrderFilledOtcOrderFilled // Event containing the contract specifics and raw log

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
func (it *OtcOrderFilledOtcOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OtcOrderFilledOtcOrderFilled)
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
		it.Event = new(OtcOrderFilledOtcOrderFilled)
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
func (it *OtcOrderFilledOtcOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OtcOrderFilledOtcOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OtcOrderFilledOtcOrderFilled represents a OtcOrderFilled event raised by the OtcOrderFilled contract.
type OtcOrderFilledOtcOrderFilled struct {
	OrderHash              [32]byte
	Maker                  common.Address
	Taker                  common.Address
	MakerToken             common.Address
	TakerToken             common.Address
	MakerTokenFilledAmount *big.Int
	TakerTokenFilledAmount *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterOtcOrderFilled is a free log retrieval operation binding the contract event 0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f.
//
// Solidity: event OtcOrderFilled(bytes32 orderHash, address maker, address taker, address makerToken, address takerToken, uint128 makerTokenFilledAmount, uint128 takerTokenFilledAmount)
func (_OtcOrderFilled *OtcOrderFilledFilterer) FilterOtcOrderFilled(opts *bind.FilterOpts) (*OtcOrderFilledOtcOrderFilledIterator, error) {

	logs, sub, err := _OtcOrderFilled.contract.FilterLogs(opts, "OtcOrderFilled")
	if err != nil {
		return nil, err
	}
	return &OtcOrderFilledOtcOrderFilledIterator{contract: _OtcOrderFilled.contract, event: "OtcOrderFilled", logs: logs, sub: sub}, nil
}

// WatchOtcOrderFilled is a free log subscription operation binding the contract event 0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f.
//
// Solidity: event OtcOrderFilled(bytes32 orderHash, address maker, address taker, address makerToken, address takerToken, uint128 makerTokenFilledAmount, uint128 takerTokenFilledAmount)
func (_OtcOrderFilled *OtcOrderFilledFilterer) WatchOtcOrderFilled(opts *bind.WatchOpts, sink chan<- *OtcOrderFilledOtcOrderFilled) (event.Subscription, error) {

	logs, sub, err := _OtcOrderFilled.contract.WatchLogs(opts, "OtcOrderFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OtcOrderFilledOtcOrderFilled)
				if err := _OtcOrderFilled.contract.UnpackLog(event, "OtcOrderFilled", log); err != nil {
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

// ParseOtcOrderFilled is a log parse operation binding the contract event 0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f.
//
// Solidity: event OtcOrderFilled(bytes32 orderHash, address maker, address taker, address makerToken, address takerToken, uint128 makerTokenFilledAmount, uint128 takerTokenFilledAmount)
func (_OtcOrderFilled *OtcOrderFilledFilterer) ParseOtcOrderFilled(log types.Log) (*OtcOrderFilledOtcOrderFilled, error) {
	event := new(OtcOrderFilledOtcOrderFilled)
	if err := _OtcOrderFilled.contract.UnpackLog(event, "OtcOrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
