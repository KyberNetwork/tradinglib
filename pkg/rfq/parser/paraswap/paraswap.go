// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paraswap

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

// OrderFilledMetaData contains all meta data concerning the OrderFilled contract.
var OrderFilledMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"}]",
}

// OrderFilledABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderFilledMetaData.ABI instead.
var OrderFilledABI = OrderFilledMetaData.ABI

// OrderFilled is an auto generated Go binding around an Ethereum contract.
type OrderFilled struct {
	OrderFilledCaller     // Read-only binding to the contract
	OrderFilledTransactor // Write-only binding to the contract
	OrderFilledFilterer   // Log filterer for contract events
}

// OrderFilledCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderFilledCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFilledTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderFilledTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFilledFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderFilledFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderFilledSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderFilledSession struct {
	Contract     *OrderFilled      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderFilledCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderFilledCallerSession struct {
	Contract *OrderFilledCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OrderFilledTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderFilledTransactorSession struct {
	Contract     *OrderFilledTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OrderFilledRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderFilledRaw struct {
	Contract *OrderFilled // Generic contract binding to access the raw methods on
}

// OrderFilledCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderFilledCallerRaw struct {
	Contract *OrderFilledCaller // Generic read-only contract binding to access the raw methods on
}

// OrderFilledTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderFilledTransactorRaw struct {
	Contract *OrderFilledTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderFilled creates a new instance of OrderFilled, bound to a specific deployed contract.
func NewOrderFilled(address common.Address, backend bind.ContractBackend) (*OrderFilled, error) {
	contract, err := bindOrderFilled(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderFilled{OrderFilledCaller: OrderFilledCaller{contract: contract}, OrderFilledTransactor: OrderFilledTransactor{contract: contract}, OrderFilledFilterer: OrderFilledFilterer{contract: contract}}, nil
}

// NewOrderFilledCaller creates a new read-only instance of OrderFilled, bound to a specific deployed contract.
func NewOrderFilledCaller(address common.Address, caller bind.ContractCaller) (*OrderFilledCaller, error) {
	contract, err := bindOrderFilled(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderFilledCaller{contract: contract}, nil
}

// NewOrderFilledTransactor creates a new write-only instance of OrderFilled, bound to a specific deployed contract.
func NewOrderFilledTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderFilledTransactor, error) {
	contract, err := bindOrderFilled(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderFilledTransactor{contract: contract}, nil
}

// NewOrderFilledFilterer creates a new log filterer instance of OrderFilled, bound to a specific deployed contract.
func NewOrderFilledFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderFilledFilterer, error) {
	contract, err := bindOrderFilled(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderFilledFilterer{contract: contract}, nil
}

// bindOrderFilled binds a generic wrapper to an already deployed contract.
func bindOrderFilled(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrderFilledMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderFilled *OrderFilledRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderFilled.Contract.OrderFilledCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderFilled *OrderFilledRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderFilled.Contract.OrderFilledTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderFilled *OrderFilledRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderFilled.Contract.OrderFilledTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderFilled *OrderFilledCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderFilled.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderFilled *OrderFilledTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderFilled.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderFilled *OrderFilledTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderFilled.Contract.contract.Transact(opts, method, params...)
}

// OrderFilledOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the OrderFilled contract.
type OrderFilledOrderFilledIterator struct {
	Event *OrderFilledOrderFilled // Event containing the contract specifics and raw log

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
func (it *OrderFilledOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderFilledOrderFilled)
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
		it.Event = new(OrderFilledOrderFilled)
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
func (it *OrderFilledOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderFilledOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderFilledOrderFilled represents a OrderFilled event raised by the OrderFilled contract.
type OrderFilledOrderFilled struct {
	OrderHash   [32]byte
	Maker       common.Address
	MakerAsset  common.Address
	MakerAmount *big.Int
	Taker       common.Address
	TakerAsset  common.Address
	TakerAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address makerAsset, uint256 makerAmount, address indexed taker, address takerAsset, uint256 takerAmount)
func (_OrderFilled *OrderFilledFilterer) FilterOrderFilled(opts *bind.FilterOpts, orderHash [][32]byte, maker []common.Address, taker []common.Address) (*OrderFilledOrderFilledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _OrderFilled.contract.FilterLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &OrderFilledOrderFilledIterator{contract: _OrderFilled.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address makerAsset, uint256 makerAmount, address indexed taker, address takerAsset, uint256 takerAmount)
func (_OrderFilled *OrderFilledFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *OrderFilledOrderFilled, orderHash [][32]byte, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _OrderFilled.contract.WatchLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderFilledOrderFilled)
				if err := _OrderFilled.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0x6621486d9c28838df4a87d2cca5007bc2aaf6a5b5de083b1db8faf709302c473.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address makerAsset, uint256 makerAmount, address indexed taker, address takerAsset, uint256 takerAmount)
func (_OrderFilled *OrderFilledFilterer) ParseOrderFilled(log types.Log) (*OrderFilledOrderFilled, error) {
	event := new(OrderFilledOrderFilled)
	if err := _OrderFilled.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
