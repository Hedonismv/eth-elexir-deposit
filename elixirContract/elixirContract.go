// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package elixirContract

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

// ElixirContractMetaData contains all meta data concerning the ElixirContract contract.
var ElixirContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DepositFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositsPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositsPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pauseDeposits\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ElixirContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ElixirContractMetaData.ABI instead.
var ElixirContractABI = ElixirContractMetaData.ABI

// ElixirContract is an auto generated Go binding around an Ethereum contract.
type ElixirContract struct {
	ElixirContractCaller     // Read-only binding to the contract
	ElixirContractTransactor // Write-only binding to the contract
	ElixirContractFilterer   // Log filterer for contract events
}

// ElixirContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElixirContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElixirContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElixirContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElixirContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElixirContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElixirContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElixirContractSession struct {
	Contract     *ElixirContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElixirContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElixirContractCallerSession struct {
	Contract *ElixirContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ElixirContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElixirContractTransactorSession struct {
	Contract     *ElixirContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ElixirContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElixirContractRaw struct {
	Contract *ElixirContract // Generic contract binding to access the raw methods on
}

// ElixirContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElixirContractCallerRaw struct {
	Contract *ElixirContractCaller // Generic read-only contract binding to access the raw methods on
}

// ElixirContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElixirContractTransactorRaw struct {
	Contract *ElixirContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElixirContract creates a new instance of ElixirContract, bound to a specific deployed contract.
func NewElixirContract(address common.Address, backend bind.ContractBackend) (*ElixirContract, error) {
	contract, err := bindElixirContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElixirContract{ElixirContractCaller: ElixirContractCaller{contract: contract}, ElixirContractTransactor: ElixirContractTransactor{contract: contract}, ElixirContractFilterer: ElixirContractFilterer{contract: contract}}, nil
}

// NewElixirContractCaller creates a new read-only instance of ElixirContract, bound to a specific deployed contract.
func NewElixirContractCaller(address common.Address, caller bind.ContractCaller) (*ElixirContractCaller, error) {
	contract, err := bindElixirContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElixirContractCaller{contract: contract}, nil
}

// NewElixirContractTransactor creates a new write-only instance of ElixirContract, bound to a specific deployed contract.
func NewElixirContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ElixirContractTransactor, error) {
	contract, err := bindElixirContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElixirContractTransactor{contract: contract}, nil
}

// NewElixirContractFilterer creates a new log filterer instance of ElixirContract, bound to a specific deployed contract.
func NewElixirContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ElixirContractFilterer, error) {
	contract, err := bindElixirContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElixirContractFilterer{contract: contract}, nil
}

// bindElixirContract binds a generic wrapper to an already deployed contract.
func bindElixirContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ElixirContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElixirContract *ElixirContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElixirContract.Contract.ElixirContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElixirContract *ElixirContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElixirContract.Contract.ElixirContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElixirContract *ElixirContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElixirContract.Contract.ElixirContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElixirContract *ElixirContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElixirContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElixirContract *ElixirContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElixirContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElixirContract *ElixirContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElixirContract.Contract.contract.Transact(opts, method, params...)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_ElixirContract *ElixirContractCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ElixirContract.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_ElixirContract *ElixirContractSession) Controller() (common.Address, error) {
	return _ElixirContract.Contract.Controller(&_ElixirContract.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_ElixirContract *ElixirContractCallerSession) Controller() (common.Address, error) {
	return _ElixirContract.Contract.Controller(&_ElixirContract.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address user) view returns(uint256 amount)
func (_ElixirContract *ElixirContractCaller) Deposits(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ElixirContract.contract.Call(opts, &out, "deposits", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address user) view returns(uint256 amount)
func (_ElixirContract *ElixirContractSession) Deposits(user common.Address) (*big.Int, error) {
	return _ElixirContract.Contract.Deposits(&_ElixirContract.CallOpts, user)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address user) view returns(uint256 amount)
func (_ElixirContract *ElixirContractCallerSession) Deposits(user common.Address) (*big.Int, error) {
	return _ElixirContract.Contract.Deposits(&_ElixirContract.CallOpts, user)
}

// DepositsPaused is a free data retrieval call binding the contract method 0x60da3e83.
//
// Solidity: function depositsPaused() view returns(bool)
func (_ElixirContract *ElixirContractCaller) DepositsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ElixirContract.contract.Call(opts, &out, "depositsPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositsPaused is a free data retrieval call binding the contract method 0x60da3e83.
//
// Solidity: function depositsPaused() view returns(bool)
func (_ElixirContract *ElixirContractSession) DepositsPaused() (bool, error) {
	return _ElixirContract.Contract.DepositsPaused(&_ElixirContract.CallOpts)
}

// DepositsPaused is a free data retrieval call binding the contract method 0x60da3e83.
//
// Solidity: function depositsPaused() view returns(bool)
func (_ElixirContract *ElixirContractCallerSession) DepositsPaused() (bool, error) {
	return _ElixirContract.Contract.DepositsPaused(&_ElixirContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ElixirContract *ElixirContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ElixirContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ElixirContract *ElixirContractSession) Owner() (common.Address, error) {
	return _ElixirContract.Contract.Owner(&_ElixirContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ElixirContract *ElixirContractCallerSession) Owner() (common.Address, error) {
	return _ElixirContract.Contract.Owner(&_ElixirContract.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ElixirContract *ElixirContractTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElixirContract.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ElixirContract *ElixirContractSession) Deposit() (*types.Transaction, error) {
	return _ElixirContract.Contract.Deposit(&_ElixirContract.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ElixirContract *ElixirContractTransactorSession) Deposit() (*types.Transaction, error) {
	return _ElixirContract.Contract.Deposit(&_ElixirContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool pauseDeposits) returns()
func (_ElixirContract *ElixirContractTransactor) Pause(opts *bind.TransactOpts, pauseDeposits bool) (*types.Transaction, error) {
	return _ElixirContract.contract.Transact(opts, "pause", pauseDeposits)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool pauseDeposits) returns()
func (_ElixirContract *ElixirContractSession) Pause(pauseDeposits bool) (*types.Transaction, error) {
	return _ElixirContract.Contract.Pause(&_ElixirContract.TransactOpts, pauseDeposits)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool pauseDeposits) returns()
func (_ElixirContract *ElixirContractTransactorSession) Pause(pauseDeposits bool) (*types.Transaction, error) {
	return _ElixirContract.Contract.Pause(&_ElixirContract.TransactOpts, pauseDeposits)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ElixirContract *ElixirContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElixirContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ElixirContract *ElixirContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _ElixirContract.Contract.RenounceOwnership(&_ElixirContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ElixirContract *ElixirContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ElixirContract.Contract.RenounceOwnership(&_ElixirContract.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_ElixirContract *ElixirContractTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _ElixirContract.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_ElixirContract *ElixirContractSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _ElixirContract.Contract.SetController(&_ElixirContract.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_ElixirContract *ElixirContractTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _ElixirContract.Contract.SetController(&_ElixirContract.TransactOpts, _controller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElixirContract *ElixirContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ElixirContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElixirContract *ElixirContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ElixirContract.Contract.TransferOwnership(&_ElixirContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElixirContract *ElixirContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ElixirContract.Contract.TransferOwnership(&_ElixirContract.TransactOpts, newOwner)
}

// ElixirContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ElixirContract contract.
type ElixirContractDepositIterator struct {
	Event *ElixirContractDeposit // Event containing the contract specifics and raw log

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
func (it *ElixirContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElixirContractDeposit)
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
		it.Event = new(ElixirContractDeposit)
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
func (it *ElixirContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElixirContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElixirContractDeposit represents a Deposit event raised by the ElixirContract contract.
type ElixirContractDeposit struct {
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed caller, uint256 indexed amount)
func (_ElixirContract *ElixirContractFilterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, amount []*big.Int) (*ElixirContractDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _ElixirContract.contract.FilterLogs(opts, "Deposit", callerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &ElixirContractDepositIterator{contract: _ElixirContract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed caller, uint256 indexed amount)
func (_ElixirContract *ElixirContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ElixirContractDeposit, caller []common.Address, amount []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _ElixirContract.contract.WatchLogs(opts, "Deposit", callerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElixirContractDeposit)
				if err := _ElixirContract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed caller, uint256 indexed amount)
func (_ElixirContract *ElixirContractFilterer) ParseDeposit(log types.Log) (*ElixirContractDeposit, error) {
	event := new(ElixirContractDeposit)
	if err := _ElixirContract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElixirContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ElixirContract contract.
type ElixirContractOwnershipTransferredIterator struct {
	Event *ElixirContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ElixirContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElixirContractOwnershipTransferred)
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
		it.Event = new(ElixirContractOwnershipTransferred)
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
func (it *ElixirContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElixirContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElixirContractOwnershipTransferred represents a OwnershipTransferred event raised by the ElixirContract contract.
type ElixirContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ElixirContract *ElixirContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ElixirContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ElixirContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ElixirContractOwnershipTransferredIterator{contract: _ElixirContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ElixirContract *ElixirContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ElixirContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ElixirContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElixirContractOwnershipTransferred)
				if err := _ElixirContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ElixirContract *ElixirContractFilterer) ParseOwnershipTransferred(log types.Log) (*ElixirContractOwnershipTransferred, error) {
	event := new(ElixirContractOwnershipTransferred)
	if err := _ElixirContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
