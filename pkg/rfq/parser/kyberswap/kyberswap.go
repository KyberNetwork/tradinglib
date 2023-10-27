// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kyberswap

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

// SwappedMetaData contains all meta data concerning the Swapped contract.
var SwappedMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"}]",
}

// SwappedABI is the input ABI used to generate the binding from.
// Deprecated: Use SwappedMetaData.ABI instead.
var SwappedABI = SwappedMetaData.ABI

// Swapped is an auto generated Go binding around an Ethereum contract.
type Swapped struct {
	SwappedCaller     // Read-only binding to the contract
	SwappedTransactor // Write-only binding to the contract
	SwappedFilterer   // Log filterer for contract events
}

// SwappedCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwappedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwappedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwappedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwappedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwappedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwappedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwappedSession struct {
	Contract     *Swapped          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwappedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwappedCallerSession struct {
	Contract *SwappedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SwappedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwappedTransactorSession struct {
	Contract     *SwappedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SwappedRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwappedRaw struct {
	Contract *Swapped // Generic contract binding to access the raw methods on
}

// SwappedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwappedCallerRaw struct {
	Contract *SwappedCaller // Generic read-only contract binding to access the raw methods on
}

// SwappedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwappedTransactorRaw struct {
	Contract *SwappedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapped creates a new instance of Swapped, bound to a specific deployed contract.
func NewSwapped(address common.Address, backend bind.ContractBackend) (*Swapped, error) {
	contract, err := bindSwapped(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swapped{SwappedCaller: SwappedCaller{contract: contract}, SwappedTransactor: SwappedTransactor{contract: contract}, SwappedFilterer: SwappedFilterer{contract: contract}}, nil
}

// NewSwappedCaller creates a new read-only instance of Swapped, bound to a specific deployed contract.
func NewSwappedCaller(address common.Address, caller bind.ContractCaller) (*SwappedCaller, error) {
	contract, err := bindSwapped(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwappedCaller{contract: contract}, nil
}

// NewSwappedTransactor creates a new write-only instance of Swapped, bound to a specific deployed contract.
func NewSwappedTransactor(address common.Address, transactor bind.ContractTransactor) (*SwappedTransactor, error) {
	contract, err := bindSwapped(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwappedTransactor{contract: contract}, nil
}

// NewSwappedFilterer creates a new log filterer instance of Swapped, bound to a specific deployed contract.
func NewSwappedFilterer(address common.Address, filterer bind.ContractFilterer) (*SwappedFilterer, error) {
	contract, err := bindSwapped(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwappedFilterer{contract: contract}, nil
}

// bindSwapped binds a generic wrapper to an already deployed contract.
func bindSwapped(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwappedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapped *SwappedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapped.Contract.SwappedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapped *SwappedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapped.Contract.SwappedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapped *SwappedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapped.Contract.SwappedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapped *SwappedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapped.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapped *SwappedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapped.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapped *SwappedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapped.Contract.contract.Transact(opts, method, params...)
}

// SwappedSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Swapped contract.
type SwappedSwappedIterator struct {
	Event *SwappedSwapped // Event containing the contract specifics and raw log

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
func (it *SwappedSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwappedSwapped)
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
		it.Event = new(SwappedSwapped)
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
func (it *SwappedSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwappedSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwappedSwapped represents a Swapped event raised by the Swapped contract.
type SwappedSwapped struct {
	Sender       common.Address
	SrcToken     common.Address
	DstToken     common.Address
	DstReceiver  common.Address
	SpentAmount  *big.Int
	ReturnAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Swapped *SwappedFilterer) FilterSwapped(opts *bind.FilterOpts) (*SwappedSwappedIterator, error) {

	logs, sub, err := _Swapped.contract.FilterLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return &SwappedSwappedIterator{contract: _Swapped.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Swapped *SwappedFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *SwappedSwapped) (event.Subscription, error) {

	logs, sub, err := _Swapped.contract.WatchLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwappedSwapped)
				if err := _Swapped.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Swapped *SwappedFilterer) ParseSwapped(log types.Log) (*SwappedSwapped, error) {
	event := new(SwappedSwapped)
	if err := _Swapped.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
