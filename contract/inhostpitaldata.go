// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// InHospitalDataABI is the input ABI used to generate the binding from.
const InHospitalDataABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_personid\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_hospitalid\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_devid\",\"type\":\"string\"}],\"name\":\"dataPushed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"_personid\",\"type\":\"string\"}],\"name\":\"getDataByPersonId\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_personid\",\"type\":\"string\"},{\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"getDataByPersonidAndDate\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"string\"},{\"name\":\"_timestamp\",\"type\":\"string\"},{\"name\":\"_hospitalid\",\"type\":\"string\"},{\"name\":\"_devid\",\"type\":\"string\"},{\"name\":\"_personid\",\"type\":\"string\"},{\"name\":\"_personveindata\",\"type\":\"string\"},{\"name\":\"_personpicdata\",\"type\":\"string\"},{\"name\":\"_additionalinfo\",\"type\":\"string\"}],\"name\":\"putData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// InHospitalData is an auto generated Go binding around an Ethereum contract.
type InHospitalData struct {
	InHospitalDataCaller     // Read-only binding to the contract
	InHospitalDataTransactor // Write-only binding to the contract
	InHospitalDataFilterer   // Log filterer for contract events
}

// InHospitalDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type InHospitalDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InHospitalDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InHospitalDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InHospitalDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InHospitalDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InHospitalDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InHospitalDataSession struct {
	Contract     *InHospitalData   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InHospitalDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InHospitalDataCallerSession struct {
	Contract *InHospitalDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// InHospitalDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InHospitalDataTransactorSession struct {
	Contract     *InHospitalDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// InHospitalDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type InHospitalDataRaw struct {
	Contract *InHospitalData // Generic contract binding to access the raw methods on
}

// InHospitalDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InHospitalDataCallerRaw struct {
	Contract *InHospitalDataCaller // Generic read-only contract binding to access the raw methods on
}

// InHospitalDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InHospitalDataTransactorRaw struct {
	Contract *InHospitalDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInHospitalData creates a new instance of InHospitalData, bound to a specific deployed contract.
func NewInHospitalData(address common.Address, backend bind.ContractBackend) (*InHospitalData, error) {
	contract, err := bindInHospitalData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InHospitalData{InHospitalDataCaller: InHospitalDataCaller{contract: contract}, InHospitalDataTransactor: InHospitalDataTransactor{contract: contract}, InHospitalDataFilterer: InHospitalDataFilterer{contract: contract}}, nil
}

// NewInHospitalDataCaller creates a new read-only instance of InHospitalData, bound to a specific deployed contract.
func NewInHospitalDataCaller(address common.Address, caller bind.ContractCaller) (*InHospitalDataCaller, error) {
	contract, err := bindInHospitalData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InHospitalDataCaller{contract: contract}, nil
}

// NewInHospitalDataTransactor creates a new write-only instance of InHospitalData, bound to a specific deployed contract.
func NewInHospitalDataTransactor(address common.Address, transactor bind.ContractTransactor) (*InHospitalDataTransactor, error) {
	contract, err := bindInHospitalData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InHospitalDataTransactor{contract: contract}, nil
}

// NewInHospitalDataFilterer creates a new log filterer instance of InHospitalData, bound to a specific deployed contract.
func NewInHospitalDataFilterer(address common.Address, filterer bind.ContractFilterer) (*InHospitalDataFilterer, error) {
	contract, err := bindInHospitalData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InHospitalDataFilterer{contract: contract}, nil
}

// bindInHospitalData binds a generic wrapper to an already deployed contract.
func bindInHospitalData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InHospitalDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InHospitalData *InHospitalDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InHospitalData.Contract.InHospitalDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InHospitalData *InHospitalDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InHospitalData.Contract.InHospitalDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InHospitalData *InHospitalDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InHospitalData.Contract.InHospitalDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InHospitalData *InHospitalDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InHospitalData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InHospitalData *InHospitalDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InHospitalData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InHospitalData *InHospitalDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InHospitalData.Contract.contract.Transact(opts, method, params...)
}

// GetDataByPersonId is a free data retrieval call binding the contract method 0xd7e3724a.
//
// Solidity: function getDataByPersonId(_personid string) constant returns(string)
func (_InHospitalData *InHospitalDataCaller) GetDataByPersonId(opts *bind.CallOpts, _personid string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _InHospitalData.contract.Call(opts, out, "getDataByPersonId", _personid)
	return *ret0, err
}

// GetDataByPersonId is a free data retrieval call binding the contract method 0xd7e3724a.
//
// Solidity: function getDataByPersonId(_personid string) constant returns(string)
func (_InHospitalData *InHospitalDataSession) GetDataByPersonId(_personid string) (string, error) {
	return _InHospitalData.Contract.GetDataByPersonId(&_InHospitalData.CallOpts, _personid)
}

// GetDataByPersonId is a free data retrieval call binding the contract method 0xd7e3724a.
//
// Solidity: function getDataByPersonId(_personid string) constant returns(string)
func (_InHospitalData *InHospitalDataCallerSession) GetDataByPersonId(_personid string) (string, error) {
	return _InHospitalData.Contract.GetDataByPersonId(&_InHospitalData.CallOpts, _personid)
}

// GetDataByPersonidAndDate is a free data retrieval call binding the contract method 0x30a42947.
//
// Solidity: function getDataByPersonidAndDate(_personid string, _date string) constant returns(string)
func (_InHospitalData *InHospitalDataCaller) GetDataByPersonidAndDate(opts *bind.CallOpts, _personid string, _date string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _InHospitalData.contract.Call(opts, out, "getDataByPersonidAndDate", _personid, _date)
	return *ret0, err
}

// GetDataByPersonidAndDate is a free data retrieval call binding the contract method 0x30a42947.
//
// Solidity: function getDataByPersonidAndDate(_personid string, _date string) constant returns(string)
func (_InHospitalData *InHospitalDataSession) GetDataByPersonidAndDate(_personid string, _date string) (string, error) {
	return _InHospitalData.Contract.GetDataByPersonidAndDate(&_InHospitalData.CallOpts, _personid, _date)
}

// GetDataByPersonidAndDate is a free data retrieval call binding the contract method 0x30a42947.
//
// Solidity: function getDataByPersonidAndDate(_personid string, _date string) constant returns(string)
func (_InHospitalData *InHospitalDataCallerSession) GetDataByPersonidAndDate(_personid string, _date string) (string, error) {
	return _InHospitalData.Contract.GetDataByPersonidAndDate(&_InHospitalData.CallOpts, _personid, _date)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_InHospitalData *InHospitalDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _InHospitalData.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_InHospitalData *InHospitalDataSession) Owner() (common.Address, error) {
	return _InHospitalData.Contract.Owner(&_InHospitalData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_InHospitalData *InHospitalDataCallerSession) Owner() (common.Address, error) {
	return _InHospitalData.Contract.Owner(&_InHospitalData.CallOpts)
}

// PutData is a paid mutator transaction binding the contract method 0x96541295.
//
// Solidity: function putData(_key string, _timestamp string, _hospitalid string, _devid string, _personid string, _personveindata string, _personpicdata string, _additionalinfo string) returns()
func (_InHospitalData *InHospitalDataTransactor) PutData(opts *bind.TransactOpts, _key string, _timestamp string, _hospitalid string, _devid string, _personid string, _personveindata string, _personpicdata string, _additionalinfo string) (*types.Transaction, error) {
	return _InHospitalData.contract.Transact(opts, "putData", _key, _timestamp, _hospitalid, _devid, _personid, _personveindata, _personpicdata, _additionalinfo)
}

// PutData is a paid mutator transaction binding the contract method 0x96541295.
//
// Solidity: function putData(_key string, _timestamp string, _hospitalid string, _devid string, _personid string, _personveindata string, _personpicdata string, _additionalinfo string) returns()
func (_InHospitalData *InHospitalDataSession) PutData(_key string, _timestamp string, _hospitalid string, _devid string, _personid string, _personveindata string, _personpicdata string, _additionalinfo string) (*types.Transaction, error) {
	return _InHospitalData.Contract.PutData(&_InHospitalData.TransactOpts, _key, _timestamp, _hospitalid, _devid, _personid, _personveindata, _personpicdata, _additionalinfo)
}

// PutData is a paid mutator transaction binding the contract method 0x96541295.
//
// Solidity: function putData(_key string, _timestamp string, _hospitalid string, _devid string, _personid string, _personveindata string, _personpicdata string, _additionalinfo string) returns()
func (_InHospitalData *InHospitalDataTransactorSession) PutData(_key string, _timestamp string, _hospitalid string, _devid string, _personid string, _personveindata string, _personpicdata string, _additionalinfo string) (*types.Transaction, error) {
	return _InHospitalData.Contract.PutData(&_InHospitalData.TransactOpts, _key, _timestamp, _hospitalid, _devid, _personid, _personveindata, _personpicdata, _additionalinfo)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InHospitalData *InHospitalDataTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InHospitalData.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InHospitalData *InHospitalDataSession) RenounceOwnership() (*types.Transaction, error) {
	return _InHospitalData.Contract.RenounceOwnership(&_InHospitalData.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InHospitalData *InHospitalDataTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InHospitalData.Contract.RenounceOwnership(&_InHospitalData.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_InHospitalData *InHospitalDataTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _InHospitalData.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_InHospitalData *InHospitalDataSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _InHospitalData.Contract.TransferOwnership(&_InHospitalData.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_InHospitalData *InHospitalDataTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _InHospitalData.Contract.TransferOwnership(&_InHospitalData.TransactOpts, _newOwner)
}

// InHospitalDataOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the InHospitalData contract.
type InHospitalDataOwnershipRenouncedIterator struct {
	Event *InHospitalDataOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *InHospitalDataOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InHospitalDataOwnershipRenounced)
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
		it.Event = new(InHospitalDataOwnershipRenounced)
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
func (it *InHospitalDataOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InHospitalDataOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InHospitalDataOwnershipRenounced represents a OwnershipRenounced event raised by the InHospitalData contract.
type InHospitalDataOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_InHospitalData *InHospitalDataFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*InHospitalDataOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _InHospitalData.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InHospitalDataOwnershipRenouncedIterator{contract: _InHospitalData.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_InHospitalData *InHospitalDataFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *InHospitalDataOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _InHospitalData.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InHospitalDataOwnershipRenounced)
				if err := _InHospitalData.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// InHospitalDataOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InHospitalData contract.
type InHospitalDataOwnershipTransferredIterator struct {
	Event *InHospitalDataOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InHospitalDataOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InHospitalDataOwnershipTransferred)
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
		it.Event = new(InHospitalDataOwnershipTransferred)
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
func (it *InHospitalDataOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InHospitalDataOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InHospitalDataOwnershipTransferred represents a OwnershipTransferred event raised by the InHospitalData contract.
type InHospitalDataOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_InHospitalData *InHospitalDataFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InHospitalDataOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InHospitalData.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InHospitalDataOwnershipTransferredIterator{contract: _InHospitalData.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_InHospitalData *InHospitalDataFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InHospitalDataOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InHospitalData.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InHospitalDataOwnershipTransferred)
				if err := _InHospitalData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// InHospitalDataDataPushedIterator is returned from FilterDataPushed and is used to iterate over the raw logs and unpacked data for DataPushed events raised by the InHospitalData contract.
type InHospitalDataDataPushedIterator struct {
	Event *InHospitalDataDataPushed // Event containing the contract specifics and raw log

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
func (it *InHospitalDataDataPushedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InHospitalDataDataPushed)
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
		it.Event = new(InHospitalDataDataPushed)
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
func (it *InHospitalDataDataPushedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InHospitalDataDataPushedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InHospitalDataDataPushed represents a DataPushed event raised by the InHospitalData contract.
type InHospitalDataDataPushed struct {
	Personid   string
	Timestamp  string
	Hospitalid string
	Devid      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDataPushed is a free log retrieval operation binding the contract event 0x4b9a74499634246b716bd7303f9930cbe533330cfb4d38b42c33a4baed9770c0.
//
// Solidity: e dataPushed(_personid string, _timestamp string, _hospitalid string, _devid string)
func (_InHospitalData *InHospitalDataFilterer) FilterDataPushed(opts *bind.FilterOpts) (*InHospitalDataDataPushedIterator, error) {

	logs, sub, err := _InHospitalData.contract.FilterLogs(opts, "dataPushed")
	if err != nil {
		return nil, err
	}
	return &InHospitalDataDataPushedIterator{contract: _InHospitalData.contract, event: "dataPushed", logs: logs, sub: sub}, nil
}

// WatchDataPushed is a free log subscription operation binding the contract event 0x4b9a74499634246b716bd7303f9930cbe533330cfb4d38b42c33a4baed9770c0.
//
// Solidity: e dataPushed(_personid string, _timestamp string, _hospitalid string, _devid string)
func (_InHospitalData *InHospitalDataFilterer) WatchDataPushed(opts *bind.WatchOpts, sink chan<- *InHospitalDataDataPushed) (event.Subscription, error) {

	logs, sub, err := _InHospitalData.contract.WatchLogs(opts, "dataPushed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InHospitalDataDataPushed)
				if err := _InHospitalData.contract.UnpackLog(event, "dataPushed", log); err != nil {
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
