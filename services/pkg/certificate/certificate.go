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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"certificateId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recipientName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"courseName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"issueDate\",\"type\":\"uint256\"}],\"name\":\"CertificateIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"certificateId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"revokedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokeDate\",\"type\":\"uint256\"}],\"name\":\"CertificateRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"IssuerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"IssuerRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"addIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedIssuers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_seed\",\"type\":\"string\"}],\"name\":\"generateCertificateId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_certificateId\",\"type\":\"bytes32\"}],\"name\":\"isCertificateValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_recipientName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_courseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_uniqueSeed\",\"type\":\"string\"}],\"name\":\"issueCertificate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"removeIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_certificateId\",\"type\":\"bytes32\"}],\"name\":\"revokeCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_certificateId\",\"type\":\"bytes32\"}],\"name\":\"verifyCertificate\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"recipientName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"courseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"issueDate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuedBy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// AuthorizedIssuers is a free data retrieval call binding the contract method 0xf731fa0f.
//
// Solidity: function authorizedIssuers(address ) view returns(bool)
func (_Certificate *CertificateCaller) AuthorizedIssuers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "authorizedIssuers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedIssuers is a free data retrieval call binding the contract method 0xf731fa0f.
//
// Solidity: function authorizedIssuers(address ) view returns(bool)
func (_Certificate *CertificateSession) AuthorizedIssuers(arg0 common.Address) (bool, error) {
	return _Certificate.Contract.AuthorizedIssuers(&_Certificate.CallOpts, arg0)
}

// AuthorizedIssuers is a free data retrieval call binding the contract method 0xf731fa0f.
//
// Solidity: function authorizedIssuers(address ) view returns(bool)
func (_Certificate *CertificateCallerSession) AuthorizedIssuers(arg0 common.Address) (bool, error) {
	return _Certificate.Contract.AuthorizedIssuers(&_Certificate.CallOpts, arg0)
}

// GenerateCertificateId is a free data retrieval call binding the contract method 0xeef79cbc.
//
// Solidity: function generateCertificateId(string _seed) pure returns(bytes32)
func (_Certificate *CertificateCaller) GenerateCertificateId(opts *bind.CallOpts, _seed string) ([32]byte, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "generateCertificateId", _seed)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GenerateCertificateId is a free data retrieval call binding the contract method 0xeef79cbc.
//
// Solidity: function generateCertificateId(string _seed) pure returns(bytes32)
func (_Certificate *CertificateSession) GenerateCertificateId(_seed string) ([32]byte, error) {
	return _Certificate.Contract.GenerateCertificateId(&_Certificate.CallOpts, _seed)
}

// GenerateCertificateId is a free data retrieval call binding the contract method 0xeef79cbc.
//
// Solidity: function generateCertificateId(string _seed) pure returns(bytes32)
func (_Certificate *CertificateCallerSession) GenerateCertificateId(_seed string) ([32]byte, error) {
	return _Certificate.Contract.GenerateCertificateId(&_Certificate.CallOpts, _seed)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _addr) view returns(bool)
func (_Certificate *CertificateCaller) IsAuthorized(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "isAuthorized", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _addr) view returns(bool)
func (_Certificate *CertificateSession) IsAuthorized(_addr common.Address) (bool, error) {
	return _Certificate.Contract.IsAuthorized(&_Certificate.CallOpts, _addr)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _addr) view returns(bool)
func (_Certificate *CertificateCallerSession) IsAuthorized(_addr common.Address) (bool, error) {
	return _Certificate.Contract.IsAuthorized(&_Certificate.CallOpts, _addr)
}

// IsCertificateValid is a free data retrieval call binding the contract method 0x05714099.
//
// Solidity: function isCertificateValid(bytes32 _certificateId) view returns(bool)
func (_Certificate *CertificateCaller) IsCertificateValid(opts *bind.CallOpts, _certificateId [32]byte) (bool, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "isCertificateValid", _certificateId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCertificateValid is a free data retrieval call binding the contract method 0x05714099.
//
// Solidity: function isCertificateValid(bytes32 _certificateId) view returns(bool)
func (_Certificate *CertificateSession) IsCertificateValid(_certificateId [32]byte) (bool, error) {
	return _Certificate.Contract.IsCertificateValid(&_Certificate.CallOpts, _certificateId)
}

// IsCertificateValid is a free data retrieval call binding the contract method 0x05714099.
//
// Solidity: function isCertificateValid(bytes32 _certificateId) view returns(bool)
func (_Certificate *CertificateCallerSession) IsCertificateValid(_certificateId [32]byte) (bool, error) {
	return _Certificate.Contract.IsCertificateValid(&_Certificate.CallOpts, _certificateId)
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
// Solidity: function verifyCertificate(bytes32 _certificateId) view returns(string recipientName, string courseName, string grade, uint256 issueDate, address issuedBy, bool isValid)
func (_Certificate *CertificateCaller) VerifyCertificate(opts *bind.CallOpts, _certificateId [32]byte) (struct {
	RecipientName string
	CourseName    string
	Grade         string
	IssueDate     *big.Int
	IssuedBy      common.Address
	IsValid       bool
}, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "verifyCertificate", _certificateId)

	outstruct := new(struct {
		RecipientName string
		CourseName    string
		Grade         string
		IssueDate     *big.Int
		IssuedBy      common.Address
		IsValid       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RecipientName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CourseName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Grade = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.IssueDate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IssuedBy = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.IsValid = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// VerifyCertificate is a free data retrieval call binding the contract method 0x850c1768.
//
// Solidity: function verifyCertificate(bytes32 _certificateId) view returns(string recipientName, string courseName, string grade, uint256 issueDate, address issuedBy, bool isValid)
func (_Certificate *CertificateSession) VerifyCertificate(_certificateId [32]byte) (struct {
	RecipientName string
	CourseName    string
	Grade         string
	IssueDate     *big.Int
	IssuedBy      common.Address
	IsValid       bool
}, error) {
	return _Certificate.Contract.VerifyCertificate(&_Certificate.CallOpts, _certificateId)
}

// VerifyCertificate is a free data retrieval call binding the contract method 0x850c1768.
//
// Solidity: function verifyCertificate(bytes32 _certificateId) view returns(string recipientName, string courseName, string grade, uint256 issueDate, address issuedBy, bool isValid)
func (_Certificate *CertificateCallerSession) VerifyCertificate(_certificateId [32]byte) (struct {
	RecipientName string
	CourseName    string
	Grade         string
	IssueDate     *big.Int
	IssuedBy      common.Address
	IsValid       bool
}, error) {
	return _Certificate.Contract.VerifyCertificate(&_Certificate.CallOpts, _certificateId)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x20694db0.
//
// Solidity: function addIssuer(address _issuer) returns()
func (_Certificate *CertificateTransactor) AddIssuer(opts *bind.TransactOpts, _issuer common.Address) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "addIssuer", _issuer)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x20694db0.
//
// Solidity: function addIssuer(address _issuer) returns()
func (_Certificate *CertificateSession) AddIssuer(_issuer common.Address) (*types.Transaction, error) {
	return _Certificate.Contract.AddIssuer(&_Certificate.TransactOpts, _issuer)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x20694db0.
//
// Solidity: function addIssuer(address _issuer) returns()
func (_Certificate *CertificateTransactorSession) AddIssuer(_issuer common.Address) (*types.Transaction, error) {
	return _Certificate.Contract.AddIssuer(&_Certificate.TransactOpts, _issuer)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x818a2994.
//
// Solidity: function issueCertificate(string _recipientName, string _courseName, string _grade, string _uniqueSeed) returns(bytes32)
func (_Certificate *CertificateTransactor) IssueCertificate(opts *bind.TransactOpts, _recipientName string, _courseName string, _grade string, _uniqueSeed string) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "issueCertificate", _recipientName, _courseName, _grade, _uniqueSeed)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x818a2994.
//
// Solidity: function issueCertificate(string _recipientName, string _courseName, string _grade, string _uniqueSeed) returns(bytes32)
func (_Certificate *CertificateSession) IssueCertificate(_recipientName string, _courseName string, _grade string, _uniqueSeed string) (*types.Transaction, error) {
	return _Certificate.Contract.IssueCertificate(&_Certificate.TransactOpts, _recipientName, _courseName, _grade, _uniqueSeed)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x818a2994.
//
// Solidity: function issueCertificate(string _recipientName, string _courseName, string _grade, string _uniqueSeed) returns(bytes32)
func (_Certificate *CertificateTransactorSession) IssueCertificate(_recipientName string, _courseName string, _grade string, _uniqueSeed string) (*types.Transaction, error) {
	return _Certificate.Contract.IssueCertificate(&_Certificate.TransactOpts, _recipientName, _courseName, _grade, _uniqueSeed)
}

// RemoveIssuer is a paid mutator transaction binding the contract method 0x47bc7093.
//
// Solidity: function removeIssuer(address _issuer) returns()
func (_Certificate *CertificateTransactor) RemoveIssuer(opts *bind.TransactOpts, _issuer common.Address) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "removeIssuer", _issuer)
}

// RemoveIssuer is a paid mutator transaction binding the contract method 0x47bc7093.
//
// Solidity: function removeIssuer(address _issuer) returns()
func (_Certificate *CertificateSession) RemoveIssuer(_issuer common.Address) (*types.Transaction, error) {
	return _Certificate.Contract.RemoveIssuer(&_Certificate.TransactOpts, _issuer)
}

// RemoveIssuer is a paid mutator transaction binding the contract method 0x47bc7093.
//
// Solidity: function removeIssuer(address _issuer) returns()
func (_Certificate *CertificateTransactorSession) RemoveIssuer(_issuer common.Address) (*types.Transaction, error) {
	return _Certificate.Contract.RemoveIssuer(&_Certificate.TransactOpts, _issuer)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xc6cbc52a.
//
// Solidity: function revokeCertificate(bytes32 _certificateId) returns()
func (_Certificate *CertificateTransactor) RevokeCertificate(opts *bind.TransactOpts, _certificateId [32]byte) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "revokeCertificate", _certificateId)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xc6cbc52a.
//
// Solidity: function revokeCertificate(bytes32 _certificateId) returns()
func (_Certificate *CertificateSession) RevokeCertificate(_certificateId [32]byte) (*types.Transaction, error) {
	return _Certificate.Contract.RevokeCertificate(&_Certificate.TransactOpts, _certificateId)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xc6cbc52a.
//
// Solidity: function revokeCertificate(bytes32 _certificateId) returns()
func (_Certificate *CertificateTransactorSession) RevokeCertificate(_certificateId [32]byte) (*types.Transaction, error) {
	return _Certificate.Contract.RevokeCertificate(&_Certificate.TransactOpts, _certificateId)
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
	CertificateId [32]byte
	RecipientName string
	CourseName    string
	IssuedBy      common.Address
	IssueDate     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCertificateIssued is a free log retrieval operation binding the contract event 0xb3e83b59ef441b355e89ba1c1cf64fcbea183a70cda9a091f7515e6f0edd0efa.
//
// Solidity: event CertificateIssued(bytes32 indexed certificateId, string recipientName, string courseName, address indexed issuedBy, uint256 issueDate)
func (_Certificate *CertificateFilterer) FilterCertificateIssued(opts *bind.FilterOpts, certificateId [][32]byte, issuedBy []common.Address) (*CertificateCertificateIssuedIterator, error) {

	var certificateIdRule []interface{}
	for _, certificateIdItem := range certificateId {
		certificateIdRule = append(certificateIdRule, certificateIdItem)
	}

	var issuedByRule []interface{}
	for _, issuedByItem := range issuedBy {
		issuedByRule = append(issuedByRule, issuedByItem)
	}

	logs, sub, err := _Certificate.contract.FilterLogs(opts, "CertificateIssued", certificateIdRule, issuedByRule)
	if err != nil {
		return nil, err
	}
	return &CertificateCertificateIssuedIterator{contract: _Certificate.contract, event: "CertificateIssued", logs: logs, sub: sub}, nil
}

// WatchCertificateIssued is a free log subscription operation binding the contract event 0xb3e83b59ef441b355e89ba1c1cf64fcbea183a70cda9a091f7515e6f0edd0efa.
//
// Solidity: event CertificateIssued(bytes32 indexed certificateId, string recipientName, string courseName, address indexed issuedBy, uint256 issueDate)
func (_Certificate *CertificateFilterer) WatchCertificateIssued(opts *bind.WatchOpts, sink chan<- *CertificateCertificateIssued, certificateId [][32]byte, issuedBy []common.Address) (event.Subscription, error) {

	var certificateIdRule []interface{}
	for _, certificateIdItem := range certificateId {
		certificateIdRule = append(certificateIdRule, certificateIdItem)
	}

	var issuedByRule []interface{}
	for _, issuedByItem := range issuedBy {
		issuedByRule = append(issuedByRule, issuedByItem)
	}

	logs, sub, err := _Certificate.contract.WatchLogs(opts, "CertificateIssued", certificateIdRule, issuedByRule)
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

// ParseCertificateIssued is a log parse operation binding the contract event 0xb3e83b59ef441b355e89ba1c1cf64fcbea183a70cda9a091f7515e6f0edd0efa.
//
// Solidity: event CertificateIssued(bytes32 indexed certificateId, string recipientName, string courseName, address indexed issuedBy, uint256 issueDate)
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
	CertificateId [32]byte
	RevokedBy     common.Address
	RevokeDate    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCertificateRevoked is a free log retrieval operation binding the contract event 0x85fe6dd2dfc4757558754564c7ff020182ec3c918257c2ff3a217a63791ec24f.
//
// Solidity: event CertificateRevoked(bytes32 indexed certificateId, address indexed revokedBy, uint256 revokeDate)
func (_Certificate *CertificateFilterer) FilterCertificateRevoked(opts *bind.FilterOpts, certificateId [][32]byte, revokedBy []common.Address) (*CertificateCertificateRevokedIterator, error) {

	var certificateIdRule []interface{}
	for _, certificateIdItem := range certificateId {
		certificateIdRule = append(certificateIdRule, certificateIdItem)
	}
	var revokedByRule []interface{}
	for _, revokedByItem := range revokedBy {
		revokedByRule = append(revokedByRule, revokedByItem)
	}

	logs, sub, err := _Certificate.contract.FilterLogs(opts, "CertificateRevoked", certificateIdRule, revokedByRule)
	if err != nil {
		return nil, err
	}
	return &CertificateCertificateRevokedIterator{contract: _Certificate.contract, event: "CertificateRevoked", logs: logs, sub: sub}, nil
}

// WatchCertificateRevoked is a free log subscription operation binding the contract event 0x85fe6dd2dfc4757558754564c7ff020182ec3c918257c2ff3a217a63791ec24f.
//
// Solidity: event CertificateRevoked(bytes32 indexed certificateId, address indexed revokedBy, uint256 revokeDate)
func (_Certificate *CertificateFilterer) WatchCertificateRevoked(opts *bind.WatchOpts, sink chan<- *CertificateCertificateRevoked, certificateId [][32]byte, revokedBy []common.Address) (event.Subscription, error) {

	var certificateIdRule []interface{}
	for _, certificateIdItem := range certificateId {
		certificateIdRule = append(certificateIdRule, certificateIdItem)
	}
	var revokedByRule []interface{}
	for _, revokedByItem := range revokedBy {
		revokedByRule = append(revokedByRule, revokedByItem)
	}

	logs, sub, err := _Certificate.contract.WatchLogs(opts, "CertificateRevoked", certificateIdRule, revokedByRule)
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

// ParseCertificateRevoked is a log parse operation binding the contract event 0x85fe6dd2dfc4757558754564c7ff020182ec3c918257c2ff3a217a63791ec24f.
//
// Solidity: event CertificateRevoked(bytes32 indexed certificateId, address indexed revokedBy, uint256 revokeDate)
func (_Certificate *CertificateFilterer) ParseCertificateRevoked(log types.Log) (*CertificateCertificateRevoked, error) {
	event := new(CertificateCertificateRevoked)
	if err := _Certificate.contract.UnpackLog(event, "CertificateRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateIssuerAddedIterator is returned from FilterIssuerAdded and is used to iterate over the raw logs and unpacked data for IssuerAdded events raised by the Certificate contract.
type CertificateIssuerAddedIterator struct {
	Event *CertificateIssuerAdded // Event containing the contract specifics and raw log

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
func (it *CertificateIssuerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateIssuerAdded)
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
		it.Event = new(CertificateIssuerAdded)
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
func (it *CertificateIssuerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateIssuerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateIssuerAdded represents a IssuerAdded event raised by the Certificate contract.
type CertificateIssuerAdded struct {
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuerAdded is a free log retrieval operation binding the contract event 0x05e7c881d716bee8cb7ed92293133ba156704252439e5c502c277448f04e20c2.
//
// Solidity: event IssuerAdded(address indexed issuer)
func (_Certificate *CertificateFilterer) FilterIssuerAdded(opts *bind.FilterOpts, issuer []common.Address) (*CertificateIssuerAddedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Certificate.contract.FilterLogs(opts, "IssuerAdded", issuerRule)
	if err != nil {
		return nil, err
	}
	return &CertificateIssuerAddedIterator{contract: _Certificate.contract, event: "IssuerAdded", logs: logs, sub: sub}, nil
}

// WatchIssuerAdded is a free log subscription operation binding the contract event 0x05e7c881d716bee8cb7ed92293133ba156704252439e5c502c277448f04e20c2.
//
// Solidity: event IssuerAdded(address indexed issuer)
func (_Certificate *CertificateFilterer) WatchIssuerAdded(opts *bind.WatchOpts, sink chan<- *CertificateIssuerAdded, issuer []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Certificate.contract.WatchLogs(opts, "IssuerAdded", issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateIssuerAdded)
				if err := _Certificate.contract.UnpackLog(event, "IssuerAdded", log); err != nil {
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

// ParseIssuerAdded is a log parse operation binding the contract event 0x05e7c881d716bee8cb7ed92293133ba156704252439e5c502c277448f04e20c2.
//
// Solidity: event IssuerAdded(address indexed issuer)
func (_Certificate *CertificateFilterer) ParseIssuerAdded(log types.Log) (*CertificateIssuerAdded, error) {
	event := new(CertificateIssuerAdded)
	if err := _Certificate.contract.UnpackLog(event, "IssuerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateIssuerRemovedIterator is returned from FilterIssuerRemoved and is used to iterate over the raw logs and unpacked data for IssuerRemoved events raised by the Certificate contract.
type CertificateIssuerRemovedIterator struct {
	Event *CertificateIssuerRemoved // Event containing the contract specifics and raw log

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
func (it *CertificateIssuerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateIssuerRemoved)
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
		it.Event = new(CertificateIssuerRemoved)
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
func (it *CertificateIssuerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateIssuerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateIssuerRemoved represents a IssuerRemoved event raised by the Certificate contract.
type CertificateIssuerRemoved struct {
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuerRemoved is a free log retrieval operation binding the contract event 0xaf66545c919a3be306ee446d8f42a9558b5b022620df880517bc9593ec0f2d52.
//
// Solidity: event IssuerRemoved(address indexed issuer)
func (_Certificate *CertificateFilterer) FilterIssuerRemoved(opts *bind.FilterOpts, issuer []common.Address) (*CertificateIssuerRemovedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Certificate.contract.FilterLogs(opts, "IssuerRemoved", issuerRule)
	if err != nil {
		return nil, err
	}
	return &CertificateIssuerRemovedIterator{contract: _Certificate.contract, event: "IssuerRemoved", logs: logs, sub: sub}, nil
}

// WatchIssuerRemoved is a free log subscription operation binding the contract event 0xaf66545c919a3be306ee446d8f42a9558b5b022620df880517bc9593ec0f2d52.
//
// Solidity: event IssuerRemoved(address indexed issuer)
func (_Certificate *CertificateFilterer) WatchIssuerRemoved(opts *bind.WatchOpts, sink chan<- *CertificateIssuerRemoved, issuer []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Certificate.contract.WatchLogs(opts, "IssuerRemoved", issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateIssuerRemoved)
				if err := _Certificate.contract.UnpackLog(event, "IssuerRemoved", log); err != nil {
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

// ParseIssuerRemoved is a log parse operation binding the contract event 0xaf66545c919a3be306ee446d8f42a9558b5b022620df880517bc9593ec0f2d52.
//
// Solidity: event IssuerRemoved(address indexed issuer)
func (_Certificate *CertificateFilterer) ParseIssuerRemoved(log types.Log) (*CertificateIssuerRemoved, error) {
	event := new(CertificateIssuerRemoved)
	if err := _Certificate.contract.UnpackLog(event, "IssuerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
