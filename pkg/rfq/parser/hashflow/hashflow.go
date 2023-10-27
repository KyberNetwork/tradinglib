// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hashflow

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

// HashflowMetaData contains all meta data concerning the Hashflow contract.
var HashflowMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteTokenAmount\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"}]",
}

// HashflowABI is the input ABI used to generate the binding from.
// Deprecated: Use HashflowMetaData.ABI instead.
var HashflowABI = HashflowMetaData.ABI

// Hashflow is an auto generated Go binding around an Ethereum contract.
type Hashflow struct {
	HashflowCaller     // Read-only binding to the contract
	HashflowTransactor // Write-only binding to the contract
	HashflowFilterer   // Log filterer for contract events
}

// HashflowCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashflowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashflowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashflowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashflowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashflowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashflowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashflowSession struct {
	Contract     *Hashflow         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashflowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashflowCallerSession struct {
	Contract *HashflowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HashflowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashflowTransactorSession struct {
	Contract     *HashflowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HashflowRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashflowRaw struct {
	Contract *Hashflow // Generic contract binding to access the raw methods on
}

// HashflowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashflowCallerRaw struct {
	Contract *HashflowCaller // Generic read-only contract binding to access the raw methods on
}

// HashflowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashflowTransactorRaw struct {
	Contract *HashflowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashflow creates a new instance of Hashflow, bound to a specific deployed contract.
func NewHashflow(address common.Address, backend bind.ContractBackend) (*Hashflow, error) {
	contract, err := bindHashflow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hashflow{HashflowCaller: HashflowCaller{contract: contract}, HashflowTransactor: HashflowTransactor{contract: contract}, HashflowFilterer: HashflowFilterer{contract: contract}}, nil
}

// NewHashflowCaller creates a new read-only instance of Hashflow, bound to a specific deployed contract.
func NewHashflowCaller(address common.Address, caller bind.ContractCaller) (*HashflowCaller, error) {
	contract, err := bindHashflow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashflowCaller{contract: contract}, nil
}

// NewHashflowTransactor creates a new write-only instance of Hashflow, bound to a specific deployed contract.
func NewHashflowTransactor(address common.Address, transactor bind.ContractTransactor) (*HashflowTransactor, error) {
	contract, err := bindHashflow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashflowTransactor{contract: contract}, nil
}

// NewHashflowFilterer creates a new log filterer instance of Hashflow, bound to a specific deployed contract.
func NewHashflowFilterer(address common.Address, filterer bind.ContractFilterer) (*HashflowFilterer, error) {
	contract, err := bindHashflow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashflowFilterer{contract: contract}, nil
}

// bindHashflow binds a generic wrapper to an already deployed contract.
func bindHashflow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashflowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashflow *HashflowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hashflow.Contract.HashflowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashflow *HashflowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashflow.Contract.HashflowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashflow *HashflowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashflow.Contract.HashflowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashflow *HashflowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hashflow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashflow *HashflowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashflow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashflow *HashflowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashflow.Contract.contract.Transact(opts, method, params...)
}

// HashflowTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Hashflow contract.
type HashflowTradeIterator struct {
	Event *HashflowTrade // Event containing the contract specifics and raw log

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
func (it *HashflowTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashflowTrade)
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
		it.Event = new(HashflowTrade)
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
func (it *HashflowTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashflowTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashflowTrade represents a Trade event raised by the Hashflow contract.
type HashflowTrade struct {
	Trader           common.Address
	Txid             [32]byte
	BaseToken        common.Address
	QuoteToken       common.Address
	BaseTokenAmount  *big.Int
	QuoteTokenAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x8cf3dec1929508e5677d7db003124e74802bfba7250a572205a9986d86ca9f1e.
//
// Solidity: event Trade(address trader, bytes32 txid, address baseToken, address quoteToken, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Hashflow *HashflowFilterer) FilterTrade(opts *bind.FilterOpts) (*HashflowTradeIterator, error) {

	logs, sub, err := _Hashflow.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &HashflowTradeIterator{contract: _Hashflow.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x8cf3dec1929508e5677d7db003124e74802bfba7250a572205a9986d86ca9f1e.
//
// Solidity: event Trade(address trader, bytes32 txid, address baseToken, address quoteToken, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Hashflow *HashflowFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *HashflowTrade) (event.Subscription, error) {

	logs, sub, err := _Hashflow.contract.WatchLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashflowTrade)
				if err := _Hashflow.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0x8cf3dec1929508e5677d7db003124e74802bfba7250a572205a9986d86ca9f1e.
//
// Solidity: event Trade(address trader, bytes32 txid, address baseToken, address quoteToken, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Hashflow *HashflowFilterer) ParseTrade(log types.Log) (*HashflowTrade, error) {
	event := new(HashflowTrade)
	if err := _Hashflow.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
