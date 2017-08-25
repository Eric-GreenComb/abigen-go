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

// ConvertLibABI is the input ABI used to generate the binding from.
const ConvertLibABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"conversionRate\",\"type\":\"uint256\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"convertedAmount\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"}]"

// ConvertLibBin is the compiled bytecode used for deploying new contracts.
const ConvertLibBin = `0x60606040523415600e57600080fd5b5b60908061001d6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166396e4ee3d8114603c575b600080fd5b6048600435602435605a565b60405190815260200160405180910390f35b8181025b929150505600a165627a7a723058202707dff741e9ad19d17a8472212a6a1fd34304f76724f21aa402e2f9a644ba650029`

// DeployConvertLib deploys a new Ethereum contract, binding an instance of ConvertLib to it.
func DeployConvertLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ConvertLib, error) {
	parsed, err := abi.JSON(strings.NewReader(ConvertLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConvertLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConvertLib{ConvertLibCaller: ConvertLibCaller{contract: contract}, ConvertLibTransactor: ConvertLibTransactor{contract: contract}}, nil
}

// ConvertLib is an auto generated Go binding around an Ethereum contract.
type ConvertLib struct {
	ConvertLibCaller     // Read-only binding to the contract
	ConvertLibTransactor // Write-only binding to the contract
}

// ConvertLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConvertLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvertLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConvertLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvertLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConvertLibSession struct {
	Contract     *ConvertLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConvertLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConvertLibCallerSession struct {
	Contract *ConvertLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ConvertLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConvertLibTransactorSession struct {
	Contract     *ConvertLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ConvertLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConvertLibRaw struct {
	Contract *ConvertLib // Generic contract binding to access the raw methods on
}

// ConvertLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConvertLibCallerRaw struct {
	Contract *ConvertLibCaller // Generic read-only contract binding to access the raw methods on
}

// ConvertLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConvertLibTransactorRaw struct {
	Contract *ConvertLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConvertLib creates a new instance of ConvertLib, bound to a specific deployed contract.
func NewConvertLib(address common.Address, backend bind.ContractBackend) (*ConvertLib, error) {
	contract, err := bindConvertLib(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConvertLib{ConvertLibCaller: ConvertLibCaller{contract: contract}, ConvertLibTransactor: ConvertLibTransactor{contract: contract}}, nil
}

// NewConvertLibCaller creates a new read-only instance of ConvertLib, bound to a specific deployed contract.
func NewConvertLibCaller(address common.Address, caller bind.ContractCaller) (*ConvertLibCaller, error) {
	contract, err := bindConvertLib(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ConvertLibCaller{contract: contract}, nil
}

// NewConvertLibTransactor creates a new write-only instance of ConvertLib, bound to a specific deployed contract.
func NewConvertLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ConvertLibTransactor, error) {
	contract, err := bindConvertLib(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ConvertLibTransactor{contract: contract}, nil
}

// bindConvertLib binds a generic wrapper to an already deployed contract.
func bindConvertLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConvertLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvertLib *ConvertLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ConvertLib.Contract.ConvertLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvertLib *ConvertLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvertLib.Contract.ConvertLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvertLib *ConvertLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvertLib.Contract.ConvertLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvertLib *ConvertLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ConvertLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvertLib *ConvertLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvertLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvertLib *ConvertLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvertLib.Contract.contract.Transact(opts, method, params...)
}

// Convert is a paid mutator transaction binding the contract method 0x96e4ee3d.
//
// Solidity: function convert(amount uint256, conversionRate uint256) returns(convertedAmount uint256)
func (_ConvertLib *ConvertLibTransactor) Convert(opts *bind.TransactOpts, amount *big.Int, conversionRate *big.Int) (*types.Transaction, error) {
	return _ConvertLib.contract.Transact(opts, "convert", amount, conversionRate)
}

// Convert is a paid mutator transaction binding the contract method 0x96e4ee3d.
//
// Solidity: function convert(amount uint256, conversionRate uint256) returns(convertedAmount uint256)
func (_ConvertLib *ConvertLibSession) Convert(amount *big.Int, conversionRate *big.Int) (*types.Transaction, error) {
	return _ConvertLib.Contract.Convert(&_ConvertLib.TransactOpts, amount, conversionRate)
}

// Convert is a paid mutator transaction binding the contract method 0x96e4ee3d.
//
// Solidity: function convert(amount uint256, conversionRate uint256) returns(convertedAmount uint256)
func (_ConvertLib *ConvertLibTransactorSession) Convert(amount *big.Int, conversionRate *big.Int) (*types.Transaction, error) {
	return _ConvertLib.Contract.Convert(&_ConvertLib.TransactOpts, amount, conversionRate)
}

// MetaCoinABI is the input ABI used to generate the binding from.
const MetaCoinABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalanceInEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendCoin\",\"outputs\":[{\"name\":\"sufficient\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// MetaCoinBin is the compiled bytecode used for deploying new contracts.
const MetaCoinBin = `0x6060604052341561000f57600080fd5b5b600160a060020a033216600090815260208190526040902061271090555b5b6102728061003e6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637bd703e8811461005357806390b98a1114610084578063f8b2cb4f146100ba575b600080fd5b341561005e57600080fd5b610072600160a060020a03600435166100eb565b60405190815260200160405180910390f35b341561008f57600080fd5b6100a6600160a060020a036004351660243561018f565b604051901515815260200160405180910390f35b34156100c557600080fd5b610072600160a060020a0360043516610227565b60405190815260200160405180910390f35b600073__ConvertLib.sol:ConvertLib_____________6396e4ee3d61011084610227565b60026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff85160281526004810192909252602482015260440160206040518083038186803b151561016d57600080fd5b6102c65a03f4151561017e57600080fd5b50505060405180519150505b919050565b600160a060020a033316600090815260208190526040812054829010156101b857506000610221565b600160a060020a033381166000818152602081905260408082208054879003905592861680825290839020805486019055917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060015b92915050565b600160a060020a0381166000908152602081905260409020545b9190505600a165627a7a72305820603f671a68f91083fa50ee029053f2063db8daf843fbaba2d831e74b5622d9a20029`

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
