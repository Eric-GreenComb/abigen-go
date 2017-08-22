// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TestABI is the input ABI used to generate the binding from.
const TestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"string\"}],\"name\":\"ModifyValue\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]"

// TestBin is the compiled bytecode used for deploying new contracts.
const TestBin = `0x6060604052341561000f57600080fd5b5b6102a88061001f6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633fa4f2458114610048578063bc45be63146100d3575b600080fd5b341561005357600080fd5b61005b610126565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100985780820151818401525b60200161007f565b50505050905090810190601f1680156100c55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156100de57600080fd5b61012460046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506101c495505050505050565b005b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101bc5780601f10610191576101008083540402835291602001916101bc565b820191906000526020600020905b81548152906001019060200180831161019f57829003601f168201915b505050505081565b60008180516101d79291602001906101dc565b505b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061021d57805160ff191683800117855561024a565b8280016001018555821561024a579182015b8281111561024a57825182559160200191906001019061022f565b5b5061025792915061025b565b5090565b61027991905b808211156102575760008155600101610261565b5090565b905600a165627a7a723058200eb5c331c0c726e3a3f39a5a9a50fc459a89100add217d4cb580c525394a063e0029`

// DeployTest deploys a new Ethereum contract, binding an instance of Test to it.
func DeployTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}}, nil
}

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(string)
func (_Test *TestCaller) Value(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Test.contract.Call(opts, out, "value")
	return *ret0, err
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(string)
func (_Test *TestSession) Value() (string, error) {
	return _Test.Contract.Value(&_Test.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(string)
func (_Test *TestCallerSession) Value() (string, error) {
	return _Test.Contract.Value(&_Test.CallOpts)
}

// ModifyValue is a paid mutator transaction binding the contract method 0xbc45be63.
//
// Solidity: function ModifyValue(val string) returns()
func (_Test *TestTransactor) ModifyValue(opts *bind.TransactOpts, val string) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "ModifyValue", val)
}

// ModifyValue is a paid mutator transaction binding the contract method 0xbc45be63.
//
// Solidity: function ModifyValue(val string) returns()
func (_Test *TestSession) ModifyValue(val string) (*types.Transaction, error) {
	return _Test.Contract.ModifyValue(&_Test.TransactOpts, val)
}

// ModifyValue is a paid mutator transaction binding the contract method 0xbc45be63.
//
// Solidity: function ModifyValue(val string) returns()
func (_Test *TestTransactorSession) ModifyValue(val string) (*types.Transaction, error) {
	return _Test.Contract.ModifyValue(&_Test.TransactOpts, val)
}
