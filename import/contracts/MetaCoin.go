// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MetaCoinABI is the input ABI used to generate the binding from.
const MetaCoinABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalanceInEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendCoin\",\"outputs\":[{\"name\":\"sufficient\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// MetaCoinBin is the compiled bytecode used for deploying new contracts.
const MetaCoinBin = `6060604052341561000f57600080fd5b5b6127106000803273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b5b6103bf806100666000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680637bd703e81461005457806390b98a11146100a1578063f8b2cb4f146100fb575b600080fd5b341561005f57600080fd5b61008b600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610148565b6040518082815260200191505060405180910390f35b34156100ac57600080fd5b6100e1600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506101f1565b604051808215151515815260200191505060405180910390f35b341561010657600080fd5b610132600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061034a565b6040518082815260200191505060405180910390f35b60007321ea709d33e96b60e2951ed887349e92037cd23e6396e4ee3d61016d8461034a565b60026000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808381526020018281526020019250505060206040518083038186803b15156101ce57600080fd5b6102c65a03f415156101df57600080fd5b5050506040518051905090505b919050565b6000816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156102425760009050610344565b816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190505b92915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b9190505600a165627a7a723058202d09bfff33bf8b06bb7a478a96b111bf475fdec3d25243d4b9f7c2d3d74951890029`

// DeployMetaCoin deploys a new Ethereum contract, binding an instance of MetaCoin to it.
func DeployMetaCoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MetaCoin, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaCoinABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MetaCoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MetaCoin{MetaCoinCaller: MetaCoinCaller{contract: contract}, MetaCoinTransactor: MetaCoinTransactor{contract: contract}}, nil
}

// MetaCoin is an auto generated Go binding around an Ethereum contract.
type MetaCoin struct {
	MetaCoinCaller     // Read-only binding to the contract
	MetaCoinTransactor // Write-only binding to the contract
}

// MetaCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaCoinSession struct {
	Contract     *MetaCoin         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaCoinCallerSession struct {
	Contract *MetaCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MetaCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaCoinTransactorSession struct {
	Contract     *MetaCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MetaCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaCoinRaw struct {
	Contract *MetaCoin // Generic contract binding to access the raw methods on
}

// MetaCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaCoinCallerRaw struct {
	Contract *MetaCoinCaller // Generic read-only contract binding to access the raw methods on
}

// MetaCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaCoinTransactorRaw struct {
	Contract *MetaCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaCoin creates a new instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoin(address common.Address, backend bind.ContractBackend) (*MetaCoin, error) {
	contract, err := bindMetaCoin(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaCoin{MetaCoinCaller: MetaCoinCaller{contract: contract}, MetaCoinTransactor: MetaCoinTransactor{contract: contract}}, nil
}

// NewMetaCoinCaller creates a new read-only instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoinCaller(address common.Address, caller bind.ContractCaller) (*MetaCoinCaller, error) {
	contract, err := bindMetaCoin(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MetaCoinCaller{contract: contract}, nil
}

// NewMetaCoinTransactor creates a new write-only instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaCoinTransactor, error) {
	contract, err := bindMetaCoin(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MetaCoinTransactor{contract: contract}, nil
}

// bindMetaCoin binds a generic wrapper to an already deployed contract.
func bindMetaCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaCoin *MetaCoinRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MetaCoin.Contract.MetaCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaCoin *MetaCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaCoin.Contract.MetaCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaCoin *MetaCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaCoin.Contract.MetaCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaCoin *MetaCoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MetaCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaCoin *MetaCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaCoin *MetaCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaCoin.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(addr address) constant returns(uint256)
func (_MetaCoin *MetaCoinCaller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MetaCoin.contract.Call(opts, out, "getBalance", addr)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(addr address) constant returns(uint256)
func (_MetaCoin *MetaCoinSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.GetBalance(&_MetaCoin.CallOpts, addr)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(addr address) constant returns(uint256)
func (_MetaCoin *MetaCoinCallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.GetBalance(&_MetaCoin.CallOpts, addr)
}

// GetBalanceInEth is a paid mutator transaction binding the contract method 0x7bd703e8.
//
// Solidity: function getBalanceInEth(addr address) returns(uint256)
func (_MetaCoin *MetaCoinTransactor) GetBalanceInEth(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _MetaCoin.contract.Transact(opts, "getBalanceInEth", addr)
}

// GetBalanceInEth is a paid mutator transaction binding the contract method 0x7bd703e8.
//
// Solidity: function getBalanceInEth(addr address) returns(uint256)
func (_MetaCoin *MetaCoinSession) GetBalanceInEth(addr common.Address) (*types.Transaction, error) {
	return _MetaCoin.Contract.GetBalanceInEth(&_MetaCoin.TransactOpts, addr)
}

// GetBalanceInEth is a paid mutator transaction binding the contract method 0x7bd703e8.
//
// Solidity: function getBalanceInEth(addr address) returns(uint256)
func (_MetaCoin *MetaCoinTransactorSession) GetBalanceInEth(addr common.Address) (*types.Transaction, error) {
	return _MetaCoin.Contract.GetBalanceInEth(&_MetaCoin.TransactOpts, addr)
}

// SendCoin is a paid mutator transaction binding the contract method 0x90b98a11.
//
// Solidity: function sendCoin(receiver address, amount uint256) returns(sufficient bool)
func (_MetaCoin *MetaCoinTransactor) SendCoin(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.contract.Transact(opts, "sendCoin", receiver, amount)
}

// SendCoin is a paid mutator transaction binding the contract method 0x90b98a11.
//
// Solidity: function sendCoin(receiver address, amount uint256) returns(sufficient bool)
func (_MetaCoin *MetaCoinSession) SendCoin(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.Contract.SendCoin(&_MetaCoin.TransactOpts, receiver, amount)
}

// SendCoin is a paid mutator transaction binding the contract method 0x90b98a11.
//
// Solidity: function sendCoin(receiver address, amount uint256) returns(sufficient bool)
func (_MetaCoin *MetaCoinTransactorSession) SendCoin(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.Contract.SendCoin(&_MetaCoin.TransactOpts, receiver, amount)
}
