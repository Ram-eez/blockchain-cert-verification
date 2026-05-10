// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package certificate

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

// CertificateMetaData contains all meta data concerning the Certificate contract.
var CertificateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pdfHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recipientName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"courseName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"issuingAuthority\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"issueDate\",\"type\":\"uint256\"}],\"name\":\"CertificateIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"pdfHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokeDate\",\"type\":\"uint256\"}],\"name\":\"CertificateRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pdfHash\",\"type\":\"bytes32\"}],\"name\":\"isCertificateValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pdfHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_recipientName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_courseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_issuingAuthority\",\"type\":\"string\"}],\"name\":\"issueCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pdfHash\",\"type\":\"bytes32\"}],\"name\":\"revokeCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pdfHash\",\"type\":\"bytes32\"}],\"name\":\"verifyCertificate\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"recipientName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"courseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"issuingAuthority\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"issueDate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CertificateABI is the input ABI used to generate the binding from.
// Deprecated: Use CertificateMetaData.ABI instead.
var CertificateABI = CertificateMetaData.ABI

// Certificate is an auto generated Go binding around an Ethereum contract.
type Certificate struct {
	CertificateCaller     // Read-only binding to the contract
	CertificateTransactor // Write-only binding to the contract
	CertificateFilterer   // Log filterer for contract events
}

// CertificateCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertificateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertificateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertificateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertificateSession struct {
	Contract     *Certificate      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertificateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertificateCallerSession struct {
	Contract *CertificateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CertificateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertificateTransactorSession struct {
	Contract     *CertificateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CertificateRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertificateRaw struct {
	Contract *Certificate // Generic contract binding to access the raw methods on
}

// CertificateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertificateCallerRaw struct {
	Contract *CertificateCaller // Generic read-only contract binding to access the raw methods on
}

// CertificateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertificateTransactorRaw struct {
	Contract *CertificateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificate creates a new instance of Certificate, bound to a specific deployed contract.
func NewCertificate(address common.Address, backend bind.ContractBackend) (*Certificate, error) {
	contract, err := bindCertificate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Certificate{CertificateCaller: CertificateCaller{contract: contract}, CertificateTransactor: CertificateTransactor{contract: contract}, CertificateFilterer: CertificateFilterer{contract: contract}}, nil
}

// NewCertificateCaller creates a new read-only instance of Certificate, bound to a specific deployed contract.
func NewCertificateCaller(address common.Address, caller bind.ContractCaller) (*CertificateCaller, error) {
	contract, err := bindCertificate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateCaller{contract: contract}, nil
}

// NewCertificateTransactor creates a new write-only instance of Certificate, bound to a specific deployed contract.
func NewCertificateTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificateTransactor, error) {
	contract, err := bindCertificate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateTransactor{contract: contract}, nil
}

// NewCertificateFilterer creates a new log filterer instance of Certificate, bound to a specific deployed contract.
func NewCertificateFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificateFilterer, error) {
	contract, err := bindCertificate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificateFilterer{contract: contract}, nil
}

// bindCertificate binds a generic wrapper to an already deployed contract.
func bindCertificate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertificateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Certificate *CertificateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Certificate.Contract.CertificateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Certificate *CertificateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Certificate.Contract.CertificateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Certificate *CertificateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Certificate.Contract.CertificateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Certificate *CertificateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Certificate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Certificate *CertificateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Certificate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Certificate *CertificateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Certificate.Contract.contract.Transact(opts, method, params...)
}

// IsCertificateValid is a free data retrieval call binding the contract method 0x05714099.
//
// Solidity: function isCertificateValid(bytes32 _pdfHash) view returns(bool)
func (_Certificate *CertificateCaller) IsCertificateValid(opts *bind.CallOpts, _pdfHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "isCertificateValid", _pdfHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCertificateValid is a free data retrieval call binding the contract method 0x05714099.
//
// Solidity: function isCertificateValid(bytes32 _pdfHash) view returns(bool)
func (_Certificate *CertificateSession) IsCertificateValid(_pdfHash [32]byte) (bool, error) {
	return _Certificate.Contract.IsCertificateValid(&_Certificate.CallOpts, _pdfHash)
}

// IsCertificateValid is a free data retrieval call binding the contract method 0x05714099.
//
// Solidity: function isCertificateValid(bytes32 _pdfHash) view returns(bool)
func (_Certificate *CertificateCallerSession) IsCertificateValid(_pdfHash [32]byte) (bool, error) {
	return _Certificate.Contract.IsCertificateValid(&_Certificate.CallOpts, _pdfHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Certificate *CertificateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Certificate *CertificateSession) Owner() (common.Address, error) {
	return _Certificate.Contract.Owner(&_Certificate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Certificate *CertificateCallerSession) Owner() (common.Address, error) {
	return _Certificate.Contract.Owner(&_Certificate.CallOpts)
}

// VerifyCertificate is a free data retrieval call binding the contract method 0x850c1768.
//
// Solidity: function verifyCertificate(bytes32 _pdfHash) view returns(string recipientName, string courseName, string grade, string issuingAuthority, uint256 issueDate, bool isValid)
func (_Certificate *CertificateCaller) VerifyCertificate(opts *bind.CallOpts, _pdfHash [32]byte) (struct {
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
	IssueDate        *big.Int
	IsValid          bool
}, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "verifyCertificate", _pdfHash)

	outstruct := new(struct {
		RecipientName    string
		CourseName       string
		Grade            string
		IssuingAuthority string
		IssueDate        *big.Int
		IsValid          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RecipientName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CourseName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Grade = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.IssuingAuthority = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IssueDate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsValid = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// VerifyCertificate is a free data retrieval call binding the contract method 0x850c1768.
//
// Solidity: function verifyCertificate(bytes32 _pdfHash) view returns(string recipientName, string courseName, string grade, string issuingAuthority, uint256 issueDate, bool isValid)
func (_Certificate *CertificateSession) VerifyCertificate(_pdfHash [32]byte) (struct {
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
	IssueDate        *big.Int
	IsValid          bool
}, error) {
	return _Certificate.Contract.VerifyCertificate(&_Certificate.CallOpts, _pdfHash)
}

// VerifyCertificate is a free data retrieval call binding the contract method 0x850c1768.
//
// Solidity: function verifyCertificate(bytes32 _pdfHash) view returns(string recipientName, string courseName, string grade, string issuingAuthority, uint256 issueDate, bool isValid)
func (_Certificate *CertificateCallerSession) VerifyCertificate(_pdfHash [32]byte) (struct {
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
	IssueDate        *big.Int
	IsValid          bool
}, error) {
	return _Certificate.Contract.VerifyCertificate(&_Certificate.CallOpts, _pdfHash)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0xdfd72020.
//
// Solidity: function issueCertificate(bytes32 _pdfHash, string _recipientName, string _courseName, string _grade, string _issuingAuthority) returns()
func (_Certificate *CertificateTransactor) IssueCertificate(opts *bind.TransactOpts, _pdfHash [32]byte, _recipientName string, _courseName string, _grade string, _issuingAuthority string) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "issueCertificate", _pdfHash, _recipientName, _courseName, _grade, _issuingAuthority)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0xdfd72020.
//
// Solidity: function issueCertificate(bytes32 _pdfHash, string _recipientName, string _courseName, string _grade, string _issuingAuthority) returns()
func (_Certificate *CertificateSession) IssueCertificate(_pdfHash [32]byte, _recipientName string, _courseName string, _grade string, _issuingAuthority string) (*types.Transaction, error) {
	return _Certificate.Contract.IssueCertificate(&_Certificate.TransactOpts, _pdfHash, _recipientName, _courseName, _grade, _issuingAuthority)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0xdfd72020.
//
// Solidity: function issueCertificate(bytes32 _pdfHash, string _recipientName, string _courseName, string _grade, string _issuingAuthority) returns()
func (_Certificate *CertificateTransactorSession) IssueCertificate(_pdfHash [32]byte, _recipientName string, _courseName string, _grade string, _issuingAuthority string) (*types.Transaction, error) {
	return _Certificate.Contract.IssueCertificate(&_Certificate.TransactOpts, _pdfHash, _recipientName, _courseName, _grade, _issuingAuthority)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xc6cbc52a.
//
// Solidity: function revokeCertificate(bytes32 _pdfHash) returns()
func (_Certificate *CertificateTransactor) RevokeCertificate(opts *bind.TransactOpts, _pdfHash [32]byte) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "revokeCertificate", _pdfHash)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xc6cbc52a.
//
// Solidity: function revokeCertificate(bytes32 _pdfHash) returns()
func (_Certificate *CertificateSession) RevokeCertificate(_pdfHash [32]byte) (*types.Transaction, error) {
	return _Certificate.Contract.RevokeCertificate(&_Certificate.TransactOpts, _pdfHash)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xc6cbc52a.
//
// Solidity: function revokeCertificate(bytes32 _pdfHash) returns()
func (_Certificate *CertificateTransactorSession) RevokeCertificate(_pdfHash [32]byte) (*types.Transaction, error) {
	return _Certificate.Contract.RevokeCertificate(&_Certificate.TransactOpts, _pdfHash)
}

// CertificateCertificateIssuedIterator is returned from FilterCertificateIssued and is used to iterate over the raw logs and unpacked data for CertificateIssued events raised by the Certificate contract.
type CertificateCertificateIssuedIterator struct {
	Event *CertificateCertificateIssued // Event containing the contract specifics and raw log

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
func (it *CertificateCertificateIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateCertificateIssued)
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
		it.Event = new(CertificateCertificateIssued)
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
func (it *CertificateCertificateIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateCertificateIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateCertificateIssued represents a CertificateIssued event raised by the Certificate contract.
type CertificateCertificateIssued struct {
	PdfHash          [32]byte
	RecipientName    string
	CourseName       string
	IssuingAuthority string
	IssueDate        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCertificateIssued is a free log retrieval operation binding the contract event 0x609da1dc5a288b98d4d6e35acd31d48b7967dea867215e051fb04d53fcd686eb.
//
// Solidity: event CertificateIssued(bytes32 indexed pdfHash, string recipientName, string courseName, string issuingAuthority, uint256 issueDate)
func (_Certificate *CertificateFilterer) FilterCertificateIssued(opts *bind.FilterOpts, pdfHash [][32]byte) (*CertificateCertificateIssuedIterator, error) {

	var pdfHashRule []interface{}
	for _, pdfHashItem := range pdfHash {
		pdfHashRule = append(pdfHashRule, pdfHashItem)
	}

	logs, sub, err := _Certificate.contract.FilterLogs(opts, "CertificateIssued", pdfHashRule)
	if err != nil {
		return nil, err
	}
	return &CertificateCertificateIssuedIterator{contract: _Certificate.contract, event: "CertificateIssued", logs: logs, sub: sub}, nil
}

// WatchCertificateIssued is a free log subscription operation binding the contract event 0x609da1dc5a288b98d4d6e35acd31d48b7967dea867215e051fb04d53fcd686eb.
//
// Solidity: event CertificateIssued(bytes32 indexed pdfHash, string recipientName, string courseName, string issuingAuthority, uint256 issueDate)
func (_Certificate *CertificateFilterer) WatchCertificateIssued(opts *bind.WatchOpts, sink chan<- *CertificateCertificateIssued, pdfHash [][32]byte) (event.Subscription, error) {

	var pdfHashRule []interface{}
	for _, pdfHashItem := range pdfHash {
		pdfHashRule = append(pdfHashRule, pdfHashItem)
	}

	logs, sub, err := _Certificate.contract.WatchLogs(opts, "CertificateIssued", pdfHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateCertificateIssued)
				if err := _Certificate.contract.UnpackLog(event, "CertificateIssued", log); err != nil {
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

// ParseCertificateIssued is a log parse operation binding the contract event 0x609da1dc5a288b98d4d6e35acd31d48b7967dea867215e051fb04d53fcd686eb.
//
// Solidity: event CertificateIssued(bytes32 indexed pdfHash, string recipientName, string courseName, string issuingAuthority, uint256 issueDate)
func (_Certificate *CertificateFilterer) ParseCertificateIssued(log types.Log) (*CertificateCertificateIssued, error) {
	event := new(CertificateCertificateIssued)
	if err := _Certificate.contract.UnpackLog(event, "CertificateIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateCertificateRevokedIterator is returned from FilterCertificateRevoked and is used to iterate over the raw logs and unpacked data for CertificateRevoked events raised by the Certificate contract.
type CertificateCertificateRevokedIterator struct {
	Event *CertificateCertificateRevoked // Event containing the contract specifics and raw log

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
func (it *CertificateCertificateRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateCertificateRevoked)
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
		it.Event = new(CertificateCertificateRevoked)
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
func (it *CertificateCertificateRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateCertificateRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateCertificateRevoked represents a CertificateRevoked event raised by the Certificate contract.
type CertificateCertificateRevoked struct {
	PdfHash    [32]byte
	RevokeDate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCertificateRevoked is a free log retrieval operation binding the contract event 0xba2ff7aceccd1742d00e386f05d483c89cbb4e8ebf02234c436df1ec854261fd.
//
// Solidity: event CertificateRevoked(bytes32 indexed pdfHash, uint256 revokeDate)
func (_Certificate *CertificateFilterer) FilterCertificateRevoked(opts *bind.FilterOpts, pdfHash [][32]byte) (*CertificateCertificateRevokedIterator, error) {

	var pdfHashRule []interface{}
	for _, pdfHashItem := range pdfHash {
		pdfHashRule = append(pdfHashRule, pdfHashItem)
	}

	logs, sub, err := _Certificate.contract.FilterLogs(opts, "CertificateRevoked", pdfHashRule)
	if err != nil {
		return nil, err
	}
	return &CertificateCertificateRevokedIterator{contract: _Certificate.contract, event: "CertificateRevoked", logs: logs, sub: sub}, nil
}

// WatchCertificateRevoked is a free log subscription operation binding the contract event 0xba2ff7aceccd1742d00e386f05d483c89cbb4e8ebf02234c436df1ec854261fd.
//
// Solidity: event CertificateRevoked(bytes32 indexed pdfHash, uint256 revokeDate)
func (_Certificate *CertificateFilterer) WatchCertificateRevoked(opts *bind.WatchOpts, sink chan<- *CertificateCertificateRevoked, pdfHash [][32]byte) (event.Subscription, error) {

	var pdfHashRule []interface{}
	for _, pdfHashItem := range pdfHash {
		pdfHashRule = append(pdfHashRule, pdfHashItem)
	}

	logs, sub, err := _Certificate.contract.WatchLogs(opts, "CertificateRevoked", pdfHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateCertificateRevoked)
				if err := _Certificate.contract.UnpackLog(event, "CertificateRevoked", log); err != nil {
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

// ParseCertificateRevoked is a log parse operation binding the contract event 0xba2ff7aceccd1742d00e386f05d483c89cbb4e8ebf02234c436df1ec854261fd.
//
// Solidity: event CertificateRevoked(bytes32 indexed pdfHash, uint256 revokeDate)
func (_Certificate *CertificateFilterer) ParseCertificateRevoked(log types.Log) (*CertificateCertificateRevoked, error) {
	event := new(CertificateCertificateRevoked)
	if err := _Certificate.contract.UnpackLog(event, "CertificateRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
