// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hashflowv3

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

// IHashflowPoolAuthorizedXChainPool is an auto generated low-level Go binding around an user-defined struct.
type IHashflowPoolAuthorizedXChainPool struct {
	ChainId uint16
	Pool    [32]byte
}

// IQuoteRFQMQuote is an auto generated low-level Go binding around an user-defined struct.
type IQuoteRFQMQuote struct {
	Pool             common.Address
	ExternalAccount  common.Address
	Trader           common.Address
	BaseToken        common.Address
	QuoteToken       common.Address
	BaseTokenAmount  *big.Int
	QuoteTokenAmount *big.Int
	QuoteExpiry      *big.Int
	Txid             [32]byte
	TakerSignature   []byte
	MakerSignature   []byte
}

// IQuoteRFQTQuote is an auto generated low-level Go binding around an user-defined struct.
type IQuoteRFQTQuote struct {
	Pool                     common.Address
	ExternalAccount          common.Address
	Trader                   common.Address
	EffectiveTrader          common.Address
	BaseToken                common.Address
	QuoteToken               common.Address
	EffectiveBaseTokenAmount *big.Int
	BaseTokenAmount          *big.Int
	QuoteTokenAmount         *big.Int
	QuoteExpiry              *big.Int
	Nonce                    *big.Int
	Txid                     [32]byte
	Signature                []byte
}

// IQuoteXChainRFQMQuote is an auto generated low-level Go binding around an user-defined struct.
type IQuoteXChainRFQMQuote struct {
	SrcChainId         uint16
	DstChainId         uint16
	SrcPool            common.Address
	DstPool            [32]byte
	SrcExternalAccount common.Address
	DstExternalAccount [32]byte
	Trader             common.Address
	DstTrader          [32]byte
	BaseToken          common.Address
	QuoteToken         [32]byte
	BaseTokenAmount    *big.Int
	QuoteTokenAmount   *big.Int
	QuoteExpiry        *big.Int
	Txid               [32]byte
	XChainMessenger    common.Address
	TakerSignature     []byte
	MakerSignature     []byte
}

// IQuoteXChainRFQTQuote is an auto generated low-level Go binding around an user-defined struct.
type IQuoteXChainRFQTQuote struct {
	SrcChainId               uint16
	DstChainId               uint16
	SrcPool                  common.Address
	DstPool                  [32]byte
	SrcExternalAccount       common.Address
	DstExternalAccount       [32]byte
	DstTrader                [32]byte
	BaseToken                common.Address
	QuoteToken               [32]byte
	EffectiveBaseTokenAmount *big.Int
	BaseTokenAmount          *big.Int
	QuoteTokenAmount         *big.Int
	QuoteExpiry              *big.Int
	Nonce                    *big.Int
	Txid                     [32]byte
	XChainMessenger          common.Address
	Signature                []byte
}

// Hashflowv3MetaData contains all meta data concerning the Hashflowv3 contract.
var Hashflowv3MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"effectiveTrader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteTokenAmount\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevSigner\",\"type\":\"address\"}],\"name\":\"UpdateSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"UpdateWithdrawalAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstPool\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstTrader\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"quoteToken\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteTokenAmount\",\"type\":\"uint256\"}],\"name\":\"XChainTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"}],\"name\":\"XChainTradeFill\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseTokenAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"externalAccount\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quoteTokenAmount\",\"type\":\"uint256\"}],\"name\":\"fillXChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseTokenAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operations\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"killswitchOperations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerConfiguration\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"externalAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerSignature\",\"type\":\"bytes\"}],\"internalType\":\"structIQuote.RFQMQuote\",\"name\":\"quote\",\"type\":\"tuple\"}],\"name\":\"tradeRFQM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"externalAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"effectiveTrader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"effectiveBaseTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIQuote.RFQTQuote\",\"name\":\"quote\",\"type\":\"tuple\"}],\"name\":\"tradeRFQT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"srcPool\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dstPool\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"srcExternalAccount\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dstExternalAccount\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dstTrader\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"quoteToken\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"baseTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteExpiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"xChainMessenger\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"takerSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"makerSignature\",\"type\":\"bytes\"}],\"internalType\":\"structIQuote.XChainRFQMQuote\",\"name\":\"quote\",\"type\":\"tuple\"}],\"name\":\"tradeXChainRFQM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"srcPool\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dstPool\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"srcExternalAccount\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dstExternalAccount\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dstTrader\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"quoteToken\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"effectiveBaseTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"xChainMessenger\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIQuote.XChainRFQTQuote\",\"name\":\"quote\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"}],\"name\":\"tradeXChainRFQT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"updateSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"withdrawalAccounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"updateWithdrawalAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"xChainMessenger\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"updateXChainMessengerAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"pool\",\"type\":\"bytes32\"}],\"internalType\":\"structIHashflowPool.AuthorizedXChainPool[]\",\"name\":\"pools\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"updateXChainPoolAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Hashflowv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Hashflowv3MetaData.ABI instead.
var Hashflowv3ABI = Hashflowv3MetaData.ABI

// Hashflowv3 is an auto generated Go binding around an Ethereum contract.
type Hashflowv3 struct {
	Hashflowv3Caller     // Read-only binding to the contract
	Hashflowv3Transactor // Write-only binding to the contract
	Hashflowv3Filterer   // Log filterer for contract events
}

// Hashflowv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Hashflowv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Hashflowv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Hashflowv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Hashflowv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Hashflowv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Hashflowv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Hashflowv3Session struct {
	Contract     *Hashflowv3       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Hashflowv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Hashflowv3CallerSession struct {
	Contract *Hashflowv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Hashflowv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Hashflowv3TransactorSession struct {
	Contract     *Hashflowv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Hashflowv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Hashflowv3Raw struct {
	Contract *Hashflowv3 // Generic contract binding to access the raw methods on
}

// Hashflowv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Hashflowv3CallerRaw struct {
	Contract *Hashflowv3Caller // Generic read-only contract binding to access the raw methods on
}

// Hashflowv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Hashflowv3TransactorRaw struct {
	Contract *Hashflowv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewHashflowv3 creates a new instance of Hashflowv3, bound to a specific deployed contract.
func NewHashflowv3(address common.Address, backend bind.ContractBackend) (*Hashflowv3, error) {
	contract, err := bindHashflowv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hashflowv3{Hashflowv3Caller: Hashflowv3Caller{contract: contract}, Hashflowv3Transactor: Hashflowv3Transactor{contract: contract}, Hashflowv3Filterer: Hashflowv3Filterer{contract: contract}}, nil
}

// NewHashflowv3Caller creates a new read-only instance of Hashflowv3, bound to a specific deployed contract.
func NewHashflowv3Caller(address common.Address, caller bind.ContractCaller) (*Hashflowv3Caller, error) {
	contract, err := bindHashflowv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Hashflowv3Caller{contract: contract}, nil
}

// NewHashflowv3Transactor creates a new write-only instance of Hashflowv3, bound to a specific deployed contract.
func NewHashflowv3Transactor(address common.Address, transactor bind.ContractTransactor) (*Hashflowv3Transactor, error) {
	contract, err := bindHashflowv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Hashflowv3Transactor{contract: contract}, nil
}

// NewHashflowv3Filterer creates a new log filterer instance of Hashflowv3, bound to a specific deployed contract.
func NewHashflowv3Filterer(address common.Address, filterer bind.ContractFilterer) (*Hashflowv3Filterer, error) {
	contract, err := bindHashflowv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Hashflowv3Filterer{contract: contract}, nil
}

// bindHashflowv3 binds a generic wrapper to an already deployed contract.
func bindHashflowv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Hashflowv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashflowv3 *Hashflowv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hashflowv3.Contract.Hashflowv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashflowv3 *Hashflowv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashflowv3.Contract.Hashflowv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashflowv3 *Hashflowv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashflowv3.Contract.Hashflowv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashflowv3 *Hashflowv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hashflowv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashflowv3 *Hashflowv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashflowv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashflowv3 *Hashflowv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashflowv3.Contract.contract.Transact(opts, method, params...)
}

// GetReserves is a free data retrieval call binding the contract method 0x3e99c1e4.
//
// Solidity: function getReserves(address token) view returns(uint256)
func (_Hashflowv3 *Hashflowv3Caller) GetReserves(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hashflowv3.contract.Call(opts, &out, "getReserves", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserves is a free data retrieval call binding the contract method 0x3e99c1e4.
//
// Solidity: function getReserves(address token) view returns(uint256)
func (_Hashflowv3 *Hashflowv3Session) GetReserves(token common.Address) (*big.Int, error) {
	return _Hashflowv3.Contract.GetReserves(&_Hashflowv3.CallOpts, token)
}

// GetReserves is a free data retrieval call binding the contract method 0x3e99c1e4.
//
// Solidity: function getReserves(address token) view returns(uint256)
func (_Hashflowv3 *Hashflowv3CallerSession) GetReserves(token common.Address) (*big.Int, error) {
	return _Hashflowv3.Contract.GetReserves(&_Hashflowv3.CallOpts, token)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_Hashflowv3 *Hashflowv3Caller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Hashflowv3.contract.Call(opts, &out, "isValidSignature", hash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_Hashflowv3 *Hashflowv3Session) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _Hashflowv3.Contract.IsValidSignature(&_Hashflowv3.CallOpts, hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_Hashflowv3 *Hashflowv3CallerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _Hashflowv3.Contract.IsValidSignature(&_Hashflowv3.CallOpts, hash, signature)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hashflowv3 *Hashflowv3Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hashflowv3.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hashflowv3 *Hashflowv3Session) Name() (string, error) {
	return _Hashflowv3.Contract.Name(&_Hashflowv3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Hashflowv3 *Hashflowv3CallerSession) Name() (string, error) {
	return _Hashflowv3.Contract.Name(&_Hashflowv3.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address trader) view returns(uint256)
func (_Hashflowv3 *Hashflowv3Caller) Nonces(opts *bind.CallOpts, trader common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hashflowv3.contract.Call(opts, &out, "nonces", trader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address trader) view returns(uint256)
func (_Hashflowv3 *Hashflowv3Session) Nonces(trader common.Address) (*big.Int, error) {
	return _Hashflowv3.Contract.Nonces(&_Hashflowv3.CallOpts, trader)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address trader) view returns(uint256)
func (_Hashflowv3 *Hashflowv3CallerSession) Nonces(trader common.Address) (*big.Int, error) {
	return _Hashflowv3.Contract.Nonces(&_Hashflowv3.CallOpts, trader)
}

// Operations is a free data retrieval call binding the contract method 0x8b33b4b2.
//
// Solidity: function operations() view returns(address)
func (_Hashflowv3 *Hashflowv3Caller) Operations(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hashflowv3.contract.Call(opts, &out, "operations")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operations is a free data retrieval call binding the contract method 0x8b33b4b2.
//
// Solidity: function operations() view returns(address)
func (_Hashflowv3 *Hashflowv3Session) Operations() (common.Address, error) {
	return _Hashflowv3.Contract.Operations(&_Hashflowv3.CallOpts)
}

// Operations is a free data retrieval call binding the contract method 0x8b33b4b2.
//
// Solidity: function operations() view returns(address)
func (_Hashflowv3 *Hashflowv3CallerSession) Operations() (common.Address, error) {
	return _Hashflowv3.Contract.Operations(&_Hashflowv3.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Hashflowv3 *Hashflowv3Caller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hashflowv3.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Hashflowv3 *Hashflowv3Session) Router() (common.Address, error) {
	return _Hashflowv3.Contract.Router(&_Hashflowv3.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Hashflowv3 *Hashflowv3CallerSession) Router() (common.Address, error) {
	return _Hashflowv3.Contract.Router(&_Hashflowv3.CallOpts)
}

// SignerConfiguration is a free data retrieval call binding the contract method 0xe0b45f38.
//
// Solidity: function signerConfiguration() view returns(address, bool)
func (_Hashflowv3 *Hashflowv3Caller) SignerConfiguration(opts *bind.CallOpts) (common.Address, bool, error) {
	var out []interface{}
	err := _Hashflowv3.contract.Call(opts, &out, "signerConfiguration")

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// SignerConfiguration is a free data retrieval call binding the contract method 0xe0b45f38.
//
// Solidity: function signerConfiguration() view returns(address, bool)
func (_Hashflowv3 *Hashflowv3Session) SignerConfiguration() (common.Address, bool, error) {
	return _Hashflowv3.Contract.SignerConfiguration(&_Hashflowv3.CallOpts)
}

// SignerConfiguration is a free data retrieval call binding the contract method 0xe0b45f38.
//
// Solidity: function signerConfiguration() view returns(address, bool)
func (_Hashflowv3 *Hashflowv3CallerSession) SignerConfiguration() (common.Address, bool, error) {
	return _Hashflowv3.Contract.SignerConfiguration(&_Hashflowv3.CallOpts)
}

// ApproveToken is a paid mutator transaction binding the contract method 0xda3e3397.
//
// Solidity: function approveToken(address token, address spender, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3Transactor) ApproveToken(opts *bind.TransactOpts, token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "approveToken", token, spender, amount)
}

// ApproveToken is a paid mutator transaction binding the contract method 0xda3e3397.
//
// Solidity: function approveToken(address token, address spender, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3Session) ApproveToken(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.Contract.ApproveToken(&_Hashflowv3.TransactOpts, token, spender, amount)
}

// ApproveToken is a paid mutator transaction binding the contract method 0xda3e3397.
//
// Solidity: function approveToken(address token, address spender, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) ApproveToken(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.Contract.ApproveToken(&_Hashflowv3.TransactOpts, token, spender, amount)
}

// DecreaseTokenAllowance is a paid mutator transaction binding the contract method 0x0bc33ce4.
//
// Solidity: function decreaseTokenAllowance(address token, address spender, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3Transactor) DecreaseTokenAllowance(opts *bind.TransactOpts, token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "decreaseTokenAllowance", token, spender, amount)
}

// DecreaseTokenAllowance is a paid mutator transaction binding the contract method 0x0bc33ce4.
//
// Solidity: function decreaseTokenAllowance(address token, address spender, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3Session) DecreaseTokenAllowance(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.Contract.DecreaseTokenAllowance(&_Hashflowv3.TransactOpts, token, spender, amount)
}

// DecreaseTokenAllowance is a paid mutator transaction binding the contract method 0x0bc33ce4.
//
// Solidity: function decreaseTokenAllowance(address token, address spender, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) DecreaseTokenAllowance(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.Contract.DecreaseTokenAllowance(&_Hashflowv3.TransactOpts, token, spender, amount)
}

// FillXChain is a paid mutator transaction binding the contract method 0x76a4f135.
//
// Solidity: function fillXChain(address externalAccount, bytes32 txid, address trader, address quoteToken, uint256 quoteTokenAmount) returns()
func (_Hashflowv3 *Hashflowv3Transactor) FillXChain(opts *bind.TransactOpts, externalAccount common.Address, txid [32]byte, trader common.Address, quoteToken common.Address, quoteTokenAmount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "fillXChain", externalAccount, txid, trader, quoteToken, quoteTokenAmount)
}

// FillXChain is a paid mutator transaction binding the contract method 0x76a4f135.
//
// Solidity: function fillXChain(address externalAccount, bytes32 txid, address trader, address quoteToken, uint256 quoteTokenAmount) returns()
func (_Hashflowv3 *Hashflowv3Session) FillXChain(externalAccount common.Address, txid [32]byte, trader common.Address, quoteToken common.Address, quoteTokenAmount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.Contract.FillXChain(&_Hashflowv3.TransactOpts, externalAccount, txid, trader, quoteToken, quoteTokenAmount)
}

// FillXChain is a paid mutator transaction binding the contract method 0x76a4f135.
//
// Solidity: function fillXChain(address externalAccount, bytes32 txid, address trader, address quoteToken, uint256 quoteTokenAmount) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) FillXChain(externalAccount common.Address, txid [32]byte, trader common.Address, quoteToken common.Address, quoteTokenAmount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.Contract.FillXChain(&_Hashflowv3.TransactOpts, externalAccount, txid, trader, quoteToken, quoteTokenAmount)
}

// IncreaseTokenAllowance is a paid mutator transaction binding the contract method 0xb39062e6.
//
// Solidity: function increaseTokenAllowance(address token, address spender, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3Transactor) IncreaseTokenAllowance(opts *bind.TransactOpts, token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "increaseTokenAllowance", token, spender, amount)
}

// IncreaseTokenAllowance is a paid mutator transaction binding the contract method 0xb39062e6.
//
// Solidity: function increaseTokenAllowance(address token, address spender, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3Session) IncreaseTokenAllowance(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.Contract.IncreaseTokenAllowance(&_Hashflowv3.TransactOpts, token, spender, amount)
}

// IncreaseTokenAllowance is a paid mutator transaction binding the contract method 0xb39062e6.
//
// Solidity: function increaseTokenAllowance(address token, address spender, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) IncreaseTokenAllowance(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.Contract.IncreaseTokenAllowance(&_Hashflowv3.TransactOpts, token, spender, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf34822b4.
//
// Solidity: function initialize(string name, address signer, address operations, address router) returns()
func (_Hashflowv3 *Hashflowv3Transactor) Initialize(opts *bind.TransactOpts, name string, signer common.Address, operations common.Address, router common.Address) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "initialize", name, signer, operations, router)
}

// Initialize is a paid mutator transaction binding the contract method 0xf34822b4.
//
// Solidity: function initialize(string name, address signer, address operations, address router) returns()
func (_Hashflowv3 *Hashflowv3Session) Initialize(name string, signer common.Address, operations common.Address, router common.Address) (*types.Transaction, error) {
	return _Hashflowv3.Contract.Initialize(&_Hashflowv3.TransactOpts, name, signer, operations, router)
}

// Initialize is a paid mutator transaction binding the contract method 0xf34822b4.
//
// Solidity: function initialize(string name, address signer, address operations, address router) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) Initialize(name string, signer common.Address, operations common.Address, router common.Address) (*types.Transaction, error) {
	return _Hashflowv3.Contract.Initialize(&_Hashflowv3.TransactOpts, name, signer, operations, router)
}

// KillswitchOperations is a paid mutator transaction binding the contract method 0xbe99a260.
//
// Solidity: function killswitchOperations(bool enabled) returns()
func (_Hashflowv3 *Hashflowv3Transactor) KillswitchOperations(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "killswitchOperations", enabled)
}

// KillswitchOperations is a paid mutator transaction binding the contract method 0xbe99a260.
//
// Solidity: function killswitchOperations(bool enabled) returns()
func (_Hashflowv3 *Hashflowv3Session) KillswitchOperations(enabled bool) (*types.Transaction, error) {
	return _Hashflowv3.Contract.KillswitchOperations(&_Hashflowv3.TransactOpts, enabled)
}

// KillswitchOperations is a paid mutator transaction binding the contract method 0xbe99a260.
//
// Solidity: function killswitchOperations(bool enabled) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) KillswitchOperations(enabled bool) (*types.Transaction, error) {
	return _Hashflowv3.Contract.KillswitchOperations(&_Hashflowv3.TransactOpts, enabled)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd752fab2.
//
// Solidity: function removeLiquidity(address token, address recipient, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3Transactor) RemoveLiquidity(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "removeLiquidity", token, recipient, amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd752fab2.
//
// Solidity: function removeLiquidity(address token, address recipient, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3Session) RemoveLiquidity(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.Contract.RemoveLiquidity(&_Hashflowv3.TransactOpts, token, recipient, amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xd752fab2.
//
// Solidity: function removeLiquidity(address token, address recipient, uint256 amount) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) RemoveLiquidity(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Hashflowv3.Contract.RemoveLiquidity(&_Hashflowv3.TransactOpts, token, recipient, amount)
}

// TradeRFQM is a paid mutator transaction binding the contract method 0x3020db05.
//
// Solidity: function tradeRFQM((address,address,address,address,address,uint256,uint256,uint256,bytes32,bytes,bytes) quote) returns()
func (_Hashflowv3 *Hashflowv3Transactor) TradeRFQM(opts *bind.TransactOpts, quote IQuoteRFQMQuote) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "tradeRFQM", quote)
}

// TradeRFQM is a paid mutator transaction binding the contract method 0x3020db05.
//
// Solidity: function tradeRFQM((address,address,address,address,address,uint256,uint256,uint256,bytes32,bytes,bytes) quote) returns()
func (_Hashflowv3 *Hashflowv3Session) TradeRFQM(quote IQuoteRFQMQuote) (*types.Transaction, error) {
	return _Hashflowv3.Contract.TradeRFQM(&_Hashflowv3.TransactOpts, quote)
}

// TradeRFQM is a paid mutator transaction binding the contract method 0x3020db05.
//
// Solidity: function tradeRFQM((address,address,address,address,address,uint256,uint256,uint256,bytes32,bytes,bytes) quote) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) TradeRFQM(quote IQuoteRFQMQuote) (*types.Transaction, error) {
	return _Hashflowv3.Contract.TradeRFQM(&_Hashflowv3.TransactOpts, quote)
}

// TradeRFQT is a paid mutator transaction binding the contract method 0xc52ac720.
//
// Solidity: function tradeRFQT((address,address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes) quote) payable returns()
func (_Hashflowv3 *Hashflowv3Transactor) TradeRFQT(opts *bind.TransactOpts, quote IQuoteRFQTQuote) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "tradeRFQT", quote)
}

// TradeRFQT is a paid mutator transaction binding the contract method 0xc52ac720.
//
// Solidity: function tradeRFQT((address,address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes) quote) payable returns()
func (_Hashflowv3 *Hashflowv3Session) TradeRFQT(quote IQuoteRFQTQuote) (*types.Transaction, error) {
	return _Hashflowv3.Contract.TradeRFQT(&_Hashflowv3.TransactOpts, quote)
}

// TradeRFQT is a paid mutator transaction binding the contract method 0xc52ac720.
//
// Solidity: function tradeRFQT((address,address,address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32,bytes) quote) payable returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) TradeRFQT(quote IQuoteRFQTQuote) (*types.Transaction, error) {
	return _Hashflowv3.Contract.TradeRFQT(&_Hashflowv3.TransactOpts, quote)
}

// TradeXChainRFQM is a paid mutator transaction binding the contract method 0xb86a5dd7.
//
// Solidity: function tradeXChainRFQM((uint16,uint16,address,bytes32,address,bytes32,address,bytes32,address,bytes32,uint256,uint256,uint256,bytes32,address,bytes,bytes) quote) returns()
func (_Hashflowv3 *Hashflowv3Transactor) TradeXChainRFQM(opts *bind.TransactOpts, quote IQuoteXChainRFQMQuote) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "tradeXChainRFQM", quote)
}

// TradeXChainRFQM is a paid mutator transaction binding the contract method 0xb86a5dd7.
//
// Solidity: function tradeXChainRFQM((uint16,uint16,address,bytes32,address,bytes32,address,bytes32,address,bytes32,uint256,uint256,uint256,bytes32,address,bytes,bytes) quote) returns()
func (_Hashflowv3 *Hashflowv3Session) TradeXChainRFQM(quote IQuoteXChainRFQMQuote) (*types.Transaction, error) {
	return _Hashflowv3.Contract.TradeXChainRFQM(&_Hashflowv3.TransactOpts, quote)
}

// TradeXChainRFQM is a paid mutator transaction binding the contract method 0xb86a5dd7.
//
// Solidity: function tradeXChainRFQM((uint16,uint16,address,bytes32,address,bytes32,address,bytes32,address,bytes32,uint256,uint256,uint256,bytes32,address,bytes,bytes) quote) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) TradeXChainRFQM(quote IQuoteXChainRFQMQuote) (*types.Transaction, error) {
	return _Hashflowv3.Contract.TradeXChainRFQM(&_Hashflowv3.TransactOpts, quote)
}

// TradeXChainRFQT is a paid mutator transaction binding the contract method 0xca3428d0.
//
// Solidity: function tradeXChainRFQT((uint16,uint16,address,bytes32,address,bytes32,bytes32,address,bytes32,uint256,uint256,uint256,uint256,uint256,bytes32,address,bytes) quote, address trader) payable returns()
func (_Hashflowv3 *Hashflowv3Transactor) TradeXChainRFQT(opts *bind.TransactOpts, quote IQuoteXChainRFQTQuote, trader common.Address) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "tradeXChainRFQT", quote, trader)
}

// TradeXChainRFQT is a paid mutator transaction binding the contract method 0xca3428d0.
//
// Solidity: function tradeXChainRFQT((uint16,uint16,address,bytes32,address,bytes32,bytes32,address,bytes32,uint256,uint256,uint256,uint256,uint256,bytes32,address,bytes) quote, address trader) payable returns()
func (_Hashflowv3 *Hashflowv3Session) TradeXChainRFQT(quote IQuoteXChainRFQTQuote, trader common.Address) (*types.Transaction, error) {
	return _Hashflowv3.Contract.TradeXChainRFQT(&_Hashflowv3.TransactOpts, quote, trader)
}

// TradeXChainRFQT is a paid mutator transaction binding the contract method 0xca3428d0.
//
// Solidity: function tradeXChainRFQT((uint16,uint16,address,bytes32,address,bytes32,bytes32,address,bytes32,uint256,uint256,uint256,uint256,uint256,bytes32,address,bytes) quote, address trader) payable returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) TradeXChainRFQT(quote IQuoteXChainRFQTQuote, trader common.Address) (*types.Transaction, error) {
	return _Hashflowv3.Contract.TradeXChainRFQT(&_Hashflowv3.TransactOpts, quote, trader)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address signer) returns()
func (_Hashflowv3 *Hashflowv3Transactor) UpdateSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "updateSigner", signer)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address signer) returns()
func (_Hashflowv3 *Hashflowv3Session) UpdateSigner(signer common.Address) (*types.Transaction, error) {
	return _Hashflowv3.Contract.UpdateSigner(&_Hashflowv3.TransactOpts, signer)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address signer) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) UpdateSigner(signer common.Address) (*types.Transaction, error) {
	return _Hashflowv3.Contract.UpdateSigner(&_Hashflowv3.TransactOpts, signer)
}

// UpdateWithdrawalAccount is a paid mutator transaction binding the contract method 0x83ade3dc.
//
// Solidity: function updateWithdrawalAccount(address[] withdrawalAccounts, bool authorized) returns()
func (_Hashflowv3 *Hashflowv3Transactor) UpdateWithdrawalAccount(opts *bind.TransactOpts, withdrawalAccounts []common.Address, authorized bool) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "updateWithdrawalAccount", withdrawalAccounts, authorized)
}

// UpdateWithdrawalAccount is a paid mutator transaction binding the contract method 0x83ade3dc.
//
// Solidity: function updateWithdrawalAccount(address[] withdrawalAccounts, bool authorized) returns()
func (_Hashflowv3 *Hashflowv3Session) UpdateWithdrawalAccount(withdrawalAccounts []common.Address, authorized bool) (*types.Transaction, error) {
	return _Hashflowv3.Contract.UpdateWithdrawalAccount(&_Hashflowv3.TransactOpts, withdrawalAccounts, authorized)
}

// UpdateWithdrawalAccount is a paid mutator transaction binding the contract method 0x83ade3dc.
//
// Solidity: function updateWithdrawalAccount(address[] withdrawalAccounts, bool authorized) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) UpdateWithdrawalAccount(withdrawalAccounts []common.Address, authorized bool) (*types.Transaction, error) {
	return _Hashflowv3.Contract.UpdateWithdrawalAccount(&_Hashflowv3.TransactOpts, withdrawalAccounts, authorized)
}

// UpdateXChainMessengerAuthorization is a paid mutator transaction binding the contract method 0x37cba5fb.
//
// Solidity: function updateXChainMessengerAuthorization(address xChainMessenger, bool authorized) returns()
func (_Hashflowv3 *Hashflowv3Transactor) UpdateXChainMessengerAuthorization(opts *bind.TransactOpts, xChainMessenger common.Address, authorized bool) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "updateXChainMessengerAuthorization", xChainMessenger, authorized)
}

// UpdateXChainMessengerAuthorization is a paid mutator transaction binding the contract method 0x37cba5fb.
//
// Solidity: function updateXChainMessengerAuthorization(address xChainMessenger, bool authorized) returns()
func (_Hashflowv3 *Hashflowv3Session) UpdateXChainMessengerAuthorization(xChainMessenger common.Address, authorized bool) (*types.Transaction, error) {
	return _Hashflowv3.Contract.UpdateXChainMessengerAuthorization(&_Hashflowv3.TransactOpts, xChainMessenger, authorized)
}

// UpdateXChainMessengerAuthorization is a paid mutator transaction binding the contract method 0x37cba5fb.
//
// Solidity: function updateXChainMessengerAuthorization(address xChainMessenger, bool authorized) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) UpdateXChainMessengerAuthorization(xChainMessenger common.Address, authorized bool) (*types.Transaction, error) {
	return _Hashflowv3.Contract.UpdateXChainMessengerAuthorization(&_Hashflowv3.TransactOpts, xChainMessenger, authorized)
}

// UpdateXChainPoolAuthorization is a paid mutator transaction binding the contract method 0x4eeaf16e.
//
// Solidity: function updateXChainPoolAuthorization((uint16,bytes32)[] pools, bool authorized) returns()
func (_Hashflowv3 *Hashflowv3Transactor) UpdateXChainPoolAuthorization(opts *bind.TransactOpts, pools []IHashflowPoolAuthorizedXChainPool, authorized bool) (*types.Transaction, error) {
	return _Hashflowv3.contract.Transact(opts, "updateXChainPoolAuthorization", pools, authorized)
}

// UpdateXChainPoolAuthorization is a paid mutator transaction binding the contract method 0x4eeaf16e.
//
// Solidity: function updateXChainPoolAuthorization((uint16,bytes32)[] pools, bool authorized) returns()
func (_Hashflowv3 *Hashflowv3Session) UpdateXChainPoolAuthorization(pools []IHashflowPoolAuthorizedXChainPool, authorized bool) (*types.Transaction, error) {
	return _Hashflowv3.Contract.UpdateXChainPoolAuthorization(&_Hashflowv3.TransactOpts, pools, authorized)
}

// UpdateXChainPoolAuthorization is a paid mutator transaction binding the contract method 0x4eeaf16e.
//
// Solidity: function updateXChainPoolAuthorization((uint16,bytes32)[] pools, bool authorized) returns()
func (_Hashflowv3 *Hashflowv3TransactorSession) UpdateXChainPoolAuthorization(pools []IHashflowPoolAuthorizedXChainPool, authorized bool) (*types.Transaction, error) {
	return _Hashflowv3.Contract.UpdateXChainPoolAuthorization(&_Hashflowv3.TransactOpts, pools, authorized)
}

// Hashflowv3RemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Hashflowv3 contract.
type Hashflowv3RemoveLiquidityIterator struct {
	Event *Hashflowv3RemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *Hashflowv3RemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Hashflowv3RemoveLiquidity)
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
		it.Event = new(Hashflowv3RemoveLiquidity)
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
func (it *Hashflowv3RemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Hashflowv3RemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Hashflowv3RemoveLiquidity represents a RemoveLiquidity event raised by the Hashflowv3 contract.
type Hashflowv3RemoveLiquidity struct {
	Token          common.Address
	Recipient      common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address token, address recipient, uint256 withdrawAmount)
func (_Hashflowv3 *Hashflowv3Filterer) FilterRemoveLiquidity(opts *bind.FilterOpts) (*Hashflowv3RemoveLiquidityIterator, error) {

	logs, sub, err := _Hashflowv3.contract.FilterLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return &Hashflowv3RemoveLiquidityIterator{contract: _Hashflowv3.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address token, address recipient, uint256 withdrawAmount)
func (_Hashflowv3 *Hashflowv3Filterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *Hashflowv3RemoveLiquidity) (event.Subscription, error) {

	logs, sub, err := _Hashflowv3.contract.WatchLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Hashflowv3RemoveLiquidity)
				if err := _Hashflowv3.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xd8ae9b9ba89e637bcb66a69ac91e8f688018e81d6f92c57e02226425c8efbdf6.
//
// Solidity: event RemoveLiquidity(address token, address recipient, uint256 withdrawAmount)
func (_Hashflowv3 *Hashflowv3Filterer) ParseRemoveLiquidity(log types.Log) (*Hashflowv3RemoveLiquidity, error) {
	event := new(Hashflowv3RemoveLiquidity)
	if err := _Hashflowv3.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Hashflowv3TradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Hashflowv3 contract.
type Hashflowv3TradeIterator struct {
	Event *Hashflowv3Trade // Event containing the contract specifics and raw log

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
func (it *Hashflowv3TradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Hashflowv3Trade)
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
		it.Event = new(Hashflowv3Trade)
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
func (it *Hashflowv3TradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Hashflowv3TradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Hashflowv3Trade represents a Trade event raised by the Hashflowv3 contract.
type Hashflowv3Trade struct {
	Trader           common.Address
	EffectiveTrader  common.Address
	Txid             [32]byte
	BaseToken        common.Address
	QuoteToken       common.Address
	BaseTokenAmount  *big.Int
	QuoteTokenAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb.
//
// Solidity: event Trade(address trader, address effectiveTrader, bytes32 txid, address baseToken, address quoteToken, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Hashflowv3 *Hashflowv3Filterer) FilterTrade(opts *bind.FilterOpts) (*Hashflowv3TradeIterator, error) {

	logs, sub, err := _Hashflowv3.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &Hashflowv3TradeIterator{contract: _Hashflowv3.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb.
//
// Solidity: event Trade(address trader, address effectiveTrader, bytes32 txid, address baseToken, address quoteToken, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Hashflowv3 *Hashflowv3Filterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *Hashflowv3Trade) (event.Subscription, error) {

	logs, sub, err := _Hashflowv3.contract.WatchLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Hashflowv3Trade)
				if err := _Hashflowv3.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0x34f57786fb01682fb4eec88d340387ef01a168fe345ea5b76f709d4e560c10eb.
//
// Solidity: event Trade(address trader, address effectiveTrader, bytes32 txid, address baseToken, address quoteToken, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Hashflowv3 *Hashflowv3Filterer) ParseTrade(log types.Log) (*Hashflowv3Trade, error) {
	event := new(Hashflowv3Trade)
	if err := _Hashflowv3.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Hashflowv3UpdateSignerIterator is returned from FilterUpdateSigner and is used to iterate over the raw logs and unpacked data for UpdateSigner events raised by the Hashflowv3 contract.
type Hashflowv3UpdateSignerIterator struct {
	Event *Hashflowv3UpdateSigner // Event containing the contract specifics and raw log

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
func (it *Hashflowv3UpdateSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Hashflowv3UpdateSigner)
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
		it.Event = new(Hashflowv3UpdateSigner)
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
func (it *Hashflowv3UpdateSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Hashflowv3UpdateSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Hashflowv3UpdateSigner represents a UpdateSigner event raised by the Hashflowv3 contract.
type Hashflowv3UpdateSigner struct {
	Signer     common.Address
	PrevSigner common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateSigner is a free log retrieval operation binding the contract event 0x72959271bae82854684905271432777342373a732ba892607d189cbf5049086f.
//
// Solidity: event UpdateSigner(address signer, address prevSigner)
func (_Hashflowv3 *Hashflowv3Filterer) FilterUpdateSigner(opts *bind.FilterOpts) (*Hashflowv3UpdateSignerIterator, error) {

	logs, sub, err := _Hashflowv3.contract.FilterLogs(opts, "UpdateSigner")
	if err != nil {
		return nil, err
	}
	return &Hashflowv3UpdateSignerIterator{contract: _Hashflowv3.contract, event: "UpdateSigner", logs: logs, sub: sub}, nil
}

// WatchUpdateSigner is a free log subscription operation binding the contract event 0x72959271bae82854684905271432777342373a732ba892607d189cbf5049086f.
//
// Solidity: event UpdateSigner(address signer, address prevSigner)
func (_Hashflowv3 *Hashflowv3Filterer) WatchUpdateSigner(opts *bind.WatchOpts, sink chan<- *Hashflowv3UpdateSigner) (event.Subscription, error) {

	logs, sub, err := _Hashflowv3.contract.WatchLogs(opts, "UpdateSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Hashflowv3UpdateSigner)
				if err := _Hashflowv3.contract.UnpackLog(event, "UpdateSigner", log); err != nil {
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

// ParseUpdateSigner is a log parse operation binding the contract event 0x72959271bae82854684905271432777342373a732ba892607d189cbf5049086f.
//
// Solidity: event UpdateSigner(address signer, address prevSigner)
func (_Hashflowv3 *Hashflowv3Filterer) ParseUpdateSigner(log types.Log) (*Hashflowv3UpdateSigner, error) {
	event := new(Hashflowv3UpdateSigner)
	if err := _Hashflowv3.contract.UnpackLog(event, "UpdateSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Hashflowv3UpdateWithdrawalAccountIterator is returned from FilterUpdateWithdrawalAccount and is used to iterate over the raw logs and unpacked data for UpdateWithdrawalAccount events raised by the Hashflowv3 contract.
type Hashflowv3UpdateWithdrawalAccountIterator struct {
	Event *Hashflowv3UpdateWithdrawalAccount // Event containing the contract specifics and raw log

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
func (it *Hashflowv3UpdateWithdrawalAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Hashflowv3UpdateWithdrawalAccount)
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
		it.Event = new(Hashflowv3UpdateWithdrawalAccount)
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
func (it *Hashflowv3UpdateWithdrawalAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Hashflowv3UpdateWithdrawalAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Hashflowv3UpdateWithdrawalAccount represents a UpdateWithdrawalAccount event raised by the Hashflowv3 contract.
type Hashflowv3UpdateWithdrawalAccount struct {
	Account    common.Address
	Authorized bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawalAccount is a free log retrieval operation binding the contract event 0xe7727546eb44fa587c611bae55b6e330f0e9184e3d9ea1f292a06cb88699d82c.
//
// Solidity: event UpdateWithdrawalAccount(address account, bool authorized)
func (_Hashflowv3 *Hashflowv3Filterer) FilterUpdateWithdrawalAccount(opts *bind.FilterOpts) (*Hashflowv3UpdateWithdrawalAccountIterator, error) {

	logs, sub, err := _Hashflowv3.contract.FilterLogs(opts, "UpdateWithdrawalAccount")
	if err != nil {
		return nil, err
	}
	return &Hashflowv3UpdateWithdrawalAccountIterator{contract: _Hashflowv3.contract, event: "UpdateWithdrawalAccount", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawalAccount is a free log subscription operation binding the contract event 0xe7727546eb44fa587c611bae55b6e330f0e9184e3d9ea1f292a06cb88699d82c.
//
// Solidity: event UpdateWithdrawalAccount(address account, bool authorized)
func (_Hashflowv3 *Hashflowv3Filterer) WatchUpdateWithdrawalAccount(opts *bind.WatchOpts, sink chan<- *Hashflowv3UpdateWithdrawalAccount) (event.Subscription, error) {

	logs, sub, err := _Hashflowv3.contract.WatchLogs(opts, "UpdateWithdrawalAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Hashflowv3UpdateWithdrawalAccount)
				if err := _Hashflowv3.contract.UnpackLog(event, "UpdateWithdrawalAccount", log); err != nil {
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

// ParseUpdateWithdrawalAccount is a log parse operation binding the contract event 0xe7727546eb44fa587c611bae55b6e330f0e9184e3d9ea1f292a06cb88699d82c.
//
// Solidity: event UpdateWithdrawalAccount(address account, bool authorized)
func (_Hashflowv3 *Hashflowv3Filterer) ParseUpdateWithdrawalAccount(log types.Log) (*Hashflowv3UpdateWithdrawalAccount, error) {
	event := new(Hashflowv3UpdateWithdrawalAccount)
	if err := _Hashflowv3.contract.UnpackLog(event, "UpdateWithdrawalAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Hashflowv3XChainTradeIterator is returned from FilterXChainTrade and is used to iterate over the raw logs and unpacked data for XChainTrade events raised by the Hashflowv3 contract.
type Hashflowv3XChainTradeIterator struct {
	Event *Hashflowv3XChainTrade // Event containing the contract specifics and raw log

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
func (it *Hashflowv3XChainTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Hashflowv3XChainTrade)
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
		it.Event = new(Hashflowv3XChainTrade)
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
func (it *Hashflowv3XChainTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Hashflowv3XChainTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Hashflowv3XChainTrade represents a XChainTrade event raised by the Hashflowv3 contract.
type Hashflowv3XChainTrade struct {
	DstChainId       uint16
	DstPool          [32]byte
	Trader           common.Address
	DstTrader        [32]byte
	Txid             [32]byte
	BaseToken        common.Address
	QuoteToken       [32]byte
	BaseTokenAmount  *big.Int
	QuoteTokenAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterXChainTrade is a free log retrieval operation binding the contract event 0x3f72b2a38919490277652bb34955c871b20e23068c243319c9fa5e27963d9e12.
//
// Solidity: event XChainTrade(uint16 dstChainId, bytes32 dstPool, address trader, bytes32 dstTrader, bytes32 txid, address baseToken, bytes32 quoteToken, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Hashflowv3 *Hashflowv3Filterer) FilterXChainTrade(opts *bind.FilterOpts) (*Hashflowv3XChainTradeIterator, error) {

	logs, sub, err := _Hashflowv3.contract.FilterLogs(opts, "XChainTrade")
	if err != nil {
		return nil, err
	}
	return &Hashflowv3XChainTradeIterator{contract: _Hashflowv3.contract, event: "XChainTrade", logs: logs, sub: sub}, nil
}

// WatchXChainTrade is a free log subscription operation binding the contract event 0x3f72b2a38919490277652bb34955c871b20e23068c243319c9fa5e27963d9e12.
//
// Solidity: event XChainTrade(uint16 dstChainId, bytes32 dstPool, address trader, bytes32 dstTrader, bytes32 txid, address baseToken, bytes32 quoteToken, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Hashflowv3 *Hashflowv3Filterer) WatchXChainTrade(opts *bind.WatchOpts, sink chan<- *Hashflowv3XChainTrade) (event.Subscription, error) {

	logs, sub, err := _Hashflowv3.contract.WatchLogs(opts, "XChainTrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Hashflowv3XChainTrade)
				if err := _Hashflowv3.contract.UnpackLog(event, "XChainTrade", log); err != nil {
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

// ParseXChainTrade is a log parse operation binding the contract event 0x3f72b2a38919490277652bb34955c871b20e23068c243319c9fa5e27963d9e12.
//
// Solidity: event XChainTrade(uint16 dstChainId, bytes32 dstPool, address trader, bytes32 dstTrader, bytes32 txid, address baseToken, bytes32 quoteToken, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Hashflowv3 *Hashflowv3Filterer) ParseXChainTrade(log types.Log) (*Hashflowv3XChainTrade, error) {
	event := new(Hashflowv3XChainTrade)
	if err := _Hashflowv3.contract.UnpackLog(event, "XChainTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Hashflowv3XChainTradeFillIterator is returned from FilterXChainTradeFill and is used to iterate over the raw logs and unpacked data for XChainTradeFill events raised by the Hashflowv3 contract.
type Hashflowv3XChainTradeFillIterator struct {
	Event *Hashflowv3XChainTradeFill // Event containing the contract specifics and raw log

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
func (it *Hashflowv3XChainTradeFillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Hashflowv3XChainTradeFill)
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
		it.Event = new(Hashflowv3XChainTradeFill)
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
func (it *Hashflowv3XChainTradeFillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Hashflowv3XChainTradeFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Hashflowv3XChainTradeFill represents a XChainTradeFill event raised by the Hashflowv3 contract.
type Hashflowv3XChainTradeFill struct {
	Txid [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterXChainTradeFill is a free log retrieval operation binding the contract event 0x51415fbbd9fd34f46b7cd4638c20246f7810a73254f3b06f167f7cf6ec9a0424.
//
// Solidity: event XChainTradeFill(bytes32 txid)
func (_Hashflowv3 *Hashflowv3Filterer) FilterXChainTradeFill(opts *bind.FilterOpts) (*Hashflowv3XChainTradeFillIterator, error) {

	logs, sub, err := _Hashflowv3.contract.FilterLogs(opts, "XChainTradeFill")
	if err != nil {
		return nil, err
	}
	return &Hashflowv3XChainTradeFillIterator{contract: _Hashflowv3.contract, event: "XChainTradeFill", logs: logs, sub: sub}, nil
}

// WatchXChainTradeFill is a free log subscription operation binding the contract event 0x51415fbbd9fd34f46b7cd4638c20246f7810a73254f3b06f167f7cf6ec9a0424.
//
// Solidity: event XChainTradeFill(bytes32 txid)
func (_Hashflowv3 *Hashflowv3Filterer) WatchXChainTradeFill(opts *bind.WatchOpts, sink chan<- *Hashflowv3XChainTradeFill) (event.Subscription, error) {

	logs, sub, err := _Hashflowv3.contract.WatchLogs(opts, "XChainTradeFill")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Hashflowv3XChainTradeFill)
				if err := _Hashflowv3.contract.UnpackLog(event, "XChainTradeFill", log); err != nil {
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

// ParseXChainTradeFill is a log parse operation binding the contract event 0x51415fbbd9fd34f46b7cd4638c20246f7810a73254f3b06f167f7cf6ec9a0424.
//
// Solidity: event XChainTradeFill(bytes32 txid)
func (_Hashflowv3 *Hashflowv3Filterer) ParseXChainTradeFill(log types.Log) (*Hashflowv3XChainTradeFill, error) {
	event := new(Hashflowv3XChainTradeFill)
	if err := _Hashflowv3.contract.UnpackLog(event, "XChainTradeFill", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
