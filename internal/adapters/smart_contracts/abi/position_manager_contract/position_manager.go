// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package position_manager_contract

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

// PositionManagerStats is an auto generated low-level Go binding around an user-defined struct.
type PositionManagerStats struct {
	PositionManagerAddress common.Address
	Name                   string
	PositionWorth          *big.Int
	CostBasis              *big.Int
	Pnl                    *big.Int
	TokenExposures         []TokenExposure
	TokenAllocations       []TokenAllocation
	Price                  *big.Int
	CollateralRatio        *big.Int
	LoanWorth              *big.Int
	LiquidationLevel       *big.Int
	Collateral             *big.Int
}

// TokenAllocation is an auto generated low-level Go binding around an user-defined struct.
type TokenAllocation struct {
	Percentage   *big.Int
	TokenAddress common.Address
	Symbol       string
	Leverage     *big.Int
	PositionType uint8
}

// TokenExposure is an auto generated low-level Go binding around an user-defined struct.
type TokenExposure struct {
	Amount *big.Int
	Token  common.Address
	Symbol string
}

// PositionManagerContractMetaData contains all meta data concerning the PositionManagerContract contract.
var PositionManagerContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"allocationByToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"enumPositionType\",\"name\":\"positionType\",\"type\":\"uint8\"}],\"internalType\":\"structTokenAllocation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocations\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"enumPositionType\",\"name\":\"positionType\",\"type\":\"uint8\"}],\"internalType\":\"structTokenAllocation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canCompound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"canRebalance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"costBasis\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exposures\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTokenExposure[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdcAmountToHave\",\"type\":\"uint256\"}],\"name\":\"getRebalanceAction\",\"outputs\":[{\"internalType\":\"enumRebalanceAction\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pnl\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionWorth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdcAmountToHave\",\"type\":\"uint256\"}],\"name\":\"rebalance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdcAmountToHave\",\"type\":\"uint256\"}],\"name\":\"rebalanceInfo\",\"outputs\":[{\"internalType\":\"enumRebalanceAction\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountToBuyOrSell\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stats\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"positionManagerAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"positionWorth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costBasis\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTokenExposure[]\",\"name\":\"tokenExposures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"enumPositionType\",\"name\":\"positionType\",\"type\":\"uint8\"}],\"internalType\":\"structTokenAllocation[]\",\"name\":\"tokenAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loanWorth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationLevel\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structPositionManagerStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PositionManagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use PositionManagerContractMetaData.ABI instead.
var PositionManagerContractABI = PositionManagerContractMetaData.ABI

// PositionManagerContract is an auto generated Go binding around an Ethereum contract.
type PositionManagerContract struct {
	PositionManagerContractCaller     // Read-only binding to the contract
	PositionManagerContractTransactor // Write-only binding to the contract
	PositionManagerContractFilterer   // Log filterer for contract events
}

// PositionManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PositionManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PositionManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PositionManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PositionManagerContractSession struct {
	Contract     *PositionManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PositionManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PositionManagerContractCallerSession struct {
	Contract *PositionManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// PositionManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PositionManagerContractTransactorSession struct {
	Contract     *PositionManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PositionManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PositionManagerContractRaw struct {
	Contract *PositionManagerContract // Generic contract binding to access the raw methods on
}

// PositionManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PositionManagerContractCallerRaw struct {
	Contract *PositionManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// PositionManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PositionManagerContractTransactorRaw struct {
	Contract *PositionManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPositionManagerContract creates a new instance of PositionManagerContract, bound to a specific deployed contract.
func NewPositionManagerContract(address common.Address, backend bind.ContractBackend) (*PositionManagerContract, error) {
	contract, err := bindPositionManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PositionManagerContract{PositionManagerContractCaller: PositionManagerContractCaller{contract: contract}, PositionManagerContractTransactor: PositionManagerContractTransactor{contract: contract}, PositionManagerContractFilterer: PositionManagerContractFilterer{contract: contract}}, nil
}

// NewPositionManagerContractCaller creates a new read-only instance of PositionManagerContract, bound to a specific deployed contract.
func NewPositionManagerContractCaller(address common.Address, caller bind.ContractCaller) (*PositionManagerContractCaller, error) {
	contract, err := bindPositionManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PositionManagerContractCaller{contract: contract}, nil
}

// NewPositionManagerContractTransactor creates a new write-only instance of PositionManagerContract, bound to a specific deployed contract.
func NewPositionManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PositionManagerContractTransactor, error) {
	contract, err := bindPositionManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PositionManagerContractTransactor{contract: contract}, nil
}

// NewPositionManagerContractFilterer creates a new log filterer instance of PositionManagerContract, bound to a specific deployed contract.
func NewPositionManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PositionManagerContractFilterer, error) {
	contract, err := bindPositionManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PositionManagerContractFilterer{contract: contract}, nil
}

// bindPositionManagerContract binds a generic wrapper to an already deployed contract.
func bindPositionManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PositionManagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionManagerContract *PositionManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionManagerContract.Contract.PositionManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionManagerContract *PositionManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManagerContract.Contract.PositionManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionManagerContract *PositionManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionManagerContract.Contract.PositionManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionManagerContract *PositionManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionManagerContract *PositionManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionManagerContract *PositionManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionManagerContract.Contract.contract.Transact(opts, method, params...)
}

// AllocationByToken is a free data retrieval call binding the contract method 0x0d95ee6a.
//
// Solidity: function allocationByToken(address tokenAddress) view returns((uint256,address,string,uint256,uint8))
func (_PositionManagerContract *PositionManagerContractCaller) AllocationByToken(opts *bind.CallOpts, tokenAddress common.Address) (TokenAllocation, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "allocationByToken", tokenAddress)

	if err != nil {
		return *new(TokenAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenAllocation)).(*TokenAllocation)

	return out0, err

}

// AllocationByToken is a free data retrieval call binding the contract method 0x0d95ee6a.
//
// Solidity: function allocationByToken(address tokenAddress) view returns((uint256,address,string,uint256,uint8))
func (_PositionManagerContract *PositionManagerContractSession) AllocationByToken(tokenAddress common.Address) (TokenAllocation, error) {
	return _PositionManagerContract.Contract.AllocationByToken(&_PositionManagerContract.CallOpts, tokenAddress)
}

// AllocationByToken is a free data retrieval call binding the contract method 0x0d95ee6a.
//
// Solidity: function allocationByToken(address tokenAddress) view returns((uint256,address,string,uint256,uint8))
func (_PositionManagerContract *PositionManagerContractCallerSession) AllocationByToken(tokenAddress common.Address) (TokenAllocation, error) {
	return _PositionManagerContract.Contract.AllocationByToken(&_PositionManagerContract.CallOpts, tokenAddress)
}

// Allocations is a free data retrieval call binding the contract method 0xc358234b.
//
// Solidity: function allocations() view returns((uint256,address,string,uint256,uint8)[])
func (_PositionManagerContract *PositionManagerContractCaller) Allocations(opts *bind.CallOpts) ([]TokenAllocation, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "allocations")

	if err != nil {
		return *new([]TokenAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new([]TokenAllocation)).(*[]TokenAllocation)

	return out0, err

}

// Allocations is a free data retrieval call binding the contract method 0xc358234b.
//
// Solidity: function allocations() view returns((uint256,address,string,uint256,uint8)[])
func (_PositionManagerContract *PositionManagerContractSession) Allocations() ([]TokenAllocation, error) {
	return _PositionManagerContract.Contract.Allocations(&_PositionManagerContract.CallOpts)
}

// Allocations is a free data retrieval call binding the contract method 0xc358234b.
//
// Solidity: function allocations() view returns((uint256,address,string,uint256,uint8)[])
func (_PositionManagerContract *PositionManagerContractCallerSession) Allocations() ([]TokenAllocation, error) {
	return _PositionManagerContract.Contract.Allocations(&_PositionManagerContract.CallOpts)
}

// CanCompound is a free data retrieval call binding the contract method 0xec8079d3.
//
// Solidity: function canCompound() view returns(bool)
func (_PositionManagerContract *PositionManagerContractCaller) CanCompound(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "canCompound")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCompound is a free data retrieval call binding the contract method 0xec8079d3.
//
// Solidity: function canCompound() view returns(bool)
func (_PositionManagerContract *PositionManagerContractSession) CanCompound() (bool, error) {
	return _PositionManagerContract.Contract.CanCompound(&_PositionManagerContract.CallOpts)
}

// CanCompound is a free data retrieval call binding the contract method 0xec8079d3.
//
// Solidity: function canCompound() view returns(bool)
func (_PositionManagerContract *PositionManagerContractCallerSession) CanCompound() (bool, error) {
	return _PositionManagerContract.Contract.CanCompound(&_PositionManagerContract.CallOpts)
}

// CanRebalance is a free data retrieval call binding the contract method 0x7e83cfec.
//
// Solidity: function canRebalance(uint256 ) view returns(bool, string)
func (_PositionManagerContract *PositionManagerContractCaller) CanRebalance(opts *bind.CallOpts, arg0 *big.Int) (bool, string, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "canRebalance", arg0)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// CanRebalance is a free data retrieval call binding the contract method 0x7e83cfec.
//
// Solidity: function canRebalance(uint256 ) view returns(bool, string)
func (_PositionManagerContract *PositionManagerContractSession) CanRebalance(arg0 *big.Int) (bool, string, error) {
	return _PositionManagerContract.Contract.CanRebalance(&_PositionManagerContract.CallOpts, arg0)
}

// CanRebalance is a free data retrieval call binding the contract method 0x7e83cfec.
//
// Solidity: function canRebalance(uint256 ) view returns(bool, string)
func (_PositionManagerContract *PositionManagerContractCallerSession) CanRebalance(arg0 *big.Int) (bool, string, error) {
	return _PositionManagerContract.Contract.CanRebalance(&_PositionManagerContract.CallOpts, arg0)
}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractCaller) CollateralRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "collateralRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractSession) CollateralRatio() (*big.Int, error) {
	return _PositionManagerContract.Contract.CollateralRatio(&_PositionManagerContract.CallOpts)
}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractCallerSession) CollateralRatio() (*big.Int, error) {
	return _PositionManagerContract.Contract.CollateralRatio(&_PositionManagerContract.CallOpts)
}

// CostBasis is a free data retrieval call binding the contract method 0x53e3cb5f.
//
// Solidity: function costBasis() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractCaller) CostBasis(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "costBasis")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CostBasis is a free data retrieval call binding the contract method 0x53e3cb5f.
//
// Solidity: function costBasis() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractSession) CostBasis() (*big.Int, error) {
	return _PositionManagerContract.Contract.CostBasis(&_PositionManagerContract.CallOpts)
}

// CostBasis is a free data retrieval call binding the contract method 0x53e3cb5f.
//
// Solidity: function costBasis() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractCallerSession) CostBasis() (*big.Int, error) {
	return _PositionManagerContract.Contract.CostBasis(&_PositionManagerContract.CallOpts)
}

// Exposures is a free data retrieval call binding the contract method 0x121cf9bc.
//
// Solidity: function exposures() view returns((int256,address,string)[])
func (_PositionManagerContract *PositionManagerContractCaller) Exposures(opts *bind.CallOpts) ([]TokenExposure, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "exposures")

	if err != nil {
		return *new([]TokenExposure), err
	}

	out0 := *abi.ConvertType(out[0], new([]TokenExposure)).(*[]TokenExposure)

	return out0, err

}

// Exposures is a free data retrieval call binding the contract method 0x121cf9bc.
//
// Solidity: function exposures() view returns((int256,address,string)[])
func (_PositionManagerContract *PositionManagerContractSession) Exposures() ([]TokenExposure, error) {
	return _PositionManagerContract.Contract.Exposures(&_PositionManagerContract.CallOpts)
}

// Exposures is a free data retrieval call binding the contract method 0x121cf9bc.
//
// Solidity: function exposures() view returns((int256,address,string)[])
func (_PositionManagerContract *PositionManagerContractCallerSession) Exposures() ([]TokenExposure, error) {
	return _PositionManagerContract.Contract.Exposures(&_PositionManagerContract.CallOpts)
}

// GetRebalanceAction is a free data retrieval call binding the contract method 0x6ca65f66.
//
// Solidity: function getRebalanceAction(uint256 usdcAmountToHave) view returns(uint8)
func (_PositionManagerContract *PositionManagerContractCaller) GetRebalanceAction(opts *bind.CallOpts, usdcAmountToHave *big.Int) (uint8, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "getRebalanceAction", usdcAmountToHave)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetRebalanceAction is a free data retrieval call binding the contract method 0x6ca65f66.
//
// Solidity: function getRebalanceAction(uint256 usdcAmountToHave) view returns(uint8)
func (_PositionManagerContract *PositionManagerContractSession) GetRebalanceAction(usdcAmountToHave *big.Int) (uint8, error) {
	return _PositionManagerContract.Contract.GetRebalanceAction(&_PositionManagerContract.CallOpts, usdcAmountToHave)
}

// GetRebalanceAction is a free data retrieval call binding the contract method 0x6ca65f66.
//
// Solidity: function getRebalanceAction(uint256 usdcAmountToHave) view returns(uint8)
func (_PositionManagerContract *PositionManagerContractCallerSession) GetRebalanceAction(usdcAmountToHave *big.Int) (uint8, error) {
	return _PositionManagerContract.Contract.GetRebalanceAction(&_PositionManagerContract.CallOpts, usdcAmountToHave)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractSession) Id() (*big.Int, error) {
	return _PositionManagerContract.Contract.Id(&_PositionManagerContract.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractCallerSession) Id() (*big.Int, error) {
	return _PositionManagerContract.Contract.Id(&_PositionManagerContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PositionManagerContract *PositionManagerContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PositionManagerContract *PositionManagerContractSession) Name() (string, error) {
	return _PositionManagerContract.Contract.Name(&_PositionManagerContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PositionManagerContract *PositionManagerContractCallerSession) Name() (string, error) {
	return _PositionManagerContract.Contract.Name(&_PositionManagerContract.CallOpts)
}

// Pnl is a free data retrieval call binding the contract method 0x3056f68b.
//
// Solidity: function pnl() view returns(int256)
func (_PositionManagerContract *PositionManagerContractCaller) Pnl(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "pnl")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pnl is a free data retrieval call binding the contract method 0x3056f68b.
//
// Solidity: function pnl() view returns(int256)
func (_PositionManagerContract *PositionManagerContractSession) Pnl() (*big.Int, error) {
	return _PositionManagerContract.Contract.Pnl(&_PositionManagerContract.CallOpts)
}

// Pnl is a free data retrieval call binding the contract method 0x3056f68b.
//
// Solidity: function pnl() view returns(int256)
func (_PositionManagerContract *PositionManagerContractCallerSession) Pnl() (*big.Int, error) {
	return _PositionManagerContract.Contract.Pnl(&_PositionManagerContract.CallOpts)
}

// PositionWorth is a free data retrieval call binding the contract method 0x9ac50e00.
//
// Solidity: function positionWorth() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractCaller) PositionWorth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "positionWorth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PositionWorth is a free data retrieval call binding the contract method 0x9ac50e00.
//
// Solidity: function positionWorth() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractSession) PositionWorth() (*big.Int, error) {
	return _PositionManagerContract.Contract.PositionWorth(&_PositionManagerContract.CallOpts)
}

// PositionWorth is a free data retrieval call binding the contract method 0x9ac50e00.
//
// Solidity: function positionWorth() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractCallerSession) PositionWorth() (*big.Int, error) {
	return _PositionManagerContract.Contract.PositionWorth(&_PositionManagerContract.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractSession) Price() (*big.Int, error) {
	return _PositionManagerContract.Contract.Price(&_PositionManagerContract.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_PositionManagerContract *PositionManagerContractCallerSession) Price() (*big.Int, error) {
	return _PositionManagerContract.Contract.Price(&_PositionManagerContract.CallOpts)
}

// RebalanceInfo is a free data retrieval call binding the contract method 0xbf92f18f.
//
// Solidity: function rebalanceInfo(uint256 usdcAmountToHave) view returns(uint8, uint256 amountToBuyOrSell)
func (_PositionManagerContract *PositionManagerContractCaller) RebalanceInfo(opts *bind.CallOpts, usdcAmountToHave *big.Int) (uint8, *big.Int, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "rebalanceInfo", usdcAmountToHave)

	if err != nil {
		return *new(uint8), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RebalanceInfo is a free data retrieval call binding the contract method 0xbf92f18f.
//
// Solidity: function rebalanceInfo(uint256 usdcAmountToHave) view returns(uint8, uint256 amountToBuyOrSell)
func (_PositionManagerContract *PositionManagerContractSession) RebalanceInfo(usdcAmountToHave *big.Int) (uint8, *big.Int, error) {
	return _PositionManagerContract.Contract.RebalanceInfo(&_PositionManagerContract.CallOpts, usdcAmountToHave)
}

// RebalanceInfo is a free data retrieval call binding the contract method 0xbf92f18f.
//
// Solidity: function rebalanceInfo(uint256 usdcAmountToHave) view returns(uint8, uint256 amountToBuyOrSell)
func (_PositionManagerContract *PositionManagerContractCallerSession) RebalanceInfo(usdcAmountToHave *big.Int) (uint8, *big.Int, error) {
	return _PositionManagerContract.Contract.RebalanceInfo(&_PositionManagerContract.CallOpts, usdcAmountToHave)
}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns((address,string,uint256,uint256,int256,(int256,address,string)[],(uint256,address,string,uint256,uint8)[],uint256,uint256,uint256,uint256,uint256))
func (_PositionManagerContract *PositionManagerContractCaller) Stats(opts *bind.CallOpts) (PositionManagerStats, error) {
	var out []interface{}
	err := _PositionManagerContract.contract.Call(opts, &out, "stats")

	if err != nil {
		return *new(PositionManagerStats), err
	}

	out0 := *abi.ConvertType(out[0], new(PositionManagerStats)).(*PositionManagerStats)

	return out0, err

}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns((address,string,uint256,uint256,int256,(int256,address,string)[],(uint256,address,string,uint256,uint8)[],uint256,uint256,uint256,uint256,uint256))
func (_PositionManagerContract *PositionManagerContractSession) Stats() (PositionManagerStats, error) {
	return _PositionManagerContract.Contract.Stats(&_PositionManagerContract.CallOpts)
}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns((address,string,uint256,uint256,int256,(int256,address,string)[],(uint256,address,string,uint256,uint8)[],uint256,uint256,uint256,uint256,uint256))
func (_PositionManagerContract *PositionManagerContractCallerSession) Stats() (PositionManagerStats, error) {
	return _PositionManagerContract.Contract.Stats(&_PositionManagerContract.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 ) returns(uint256)
func (_PositionManagerContract *PositionManagerContractTransactor) Buy(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _PositionManagerContract.contract.Transact(opts, "buy", arg0)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 ) returns(uint256)
func (_PositionManagerContract *PositionManagerContractSession) Buy(arg0 *big.Int) (*types.Transaction, error) {
	return _PositionManagerContract.Contract.Buy(&_PositionManagerContract.TransactOpts, arg0)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 ) returns(uint256)
func (_PositionManagerContract *PositionManagerContractTransactorSession) Buy(arg0 *big.Int) (*types.Transaction, error) {
	return _PositionManagerContract.Contract.Buy(&_PositionManagerContract.TransactOpts, arg0)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns(uint256)
func (_PositionManagerContract *PositionManagerContractTransactor) Compound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManagerContract.contract.Transact(opts, "compound")
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns(uint256)
func (_PositionManagerContract *PositionManagerContractSession) Compound() (*types.Transaction, error) {
	return _PositionManagerContract.Contract.Compound(&_PositionManagerContract.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns(uint256)
func (_PositionManagerContract *PositionManagerContractTransactorSession) Compound() (*types.Transaction, error) {
	return _PositionManagerContract.Contract.Compound(&_PositionManagerContract.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0x28a07025.
//
// Solidity: function liquidate() returns()
func (_PositionManagerContract *PositionManagerContractTransactor) Liquidate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManagerContract.contract.Transact(opts, "liquidate")
}

// Liquidate is a paid mutator transaction binding the contract method 0x28a07025.
//
// Solidity: function liquidate() returns()
func (_PositionManagerContract *PositionManagerContractSession) Liquidate() (*types.Transaction, error) {
	return _PositionManagerContract.Contract.Liquidate(&_PositionManagerContract.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0x28a07025.
//
// Solidity: function liquidate() returns()
func (_PositionManagerContract *PositionManagerContractTransactorSession) Liquidate() (*types.Transaction, error) {
	return _PositionManagerContract.Contract.Liquidate(&_PositionManagerContract.TransactOpts)
}

// Rebalance is a paid mutator transaction binding the contract method 0xf4993018.
//
// Solidity: function rebalance(uint256 usdcAmountToHave) returns(bool)
func (_PositionManagerContract *PositionManagerContractTransactor) Rebalance(opts *bind.TransactOpts, usdcAmountToHave *big.Int) (*types.Transaction, error) {
	return _PositionManagerContract.contract.Transact(opts, "rebalance", usdcAmountToHave)
}

// Rebalance is a paid mutator transaction binding the contract method 0xf4993018.
//
// Solidity: function rebalance(uint256 usdcAmountToHave) returns(bool)
func (_PositionManagerContract *PositionManagerContractSession) Rebalance(usdcAmountToHave *big.Int) (*types.Transaction, error) {
	return _PositionManagerContract.Contract.Rebalance(&_PositionManagerContract.TransactOpts, usdcAmountToHave)
}

// Rebalance is a paid mutator transaction binding the contract method 0xf4993018.
//
// Solidity: function rebalance(uint256 usdcAmountToHave) returns(bool)
func (_PositionManagerContract *PositionManagerContractTransactorSession) Rebalance(usdcAmountToHave *big.Int) (*types.Transaction, error) {
	return _PositionManagerContract.Contract.Rebalance(&_PositionManagerContract.TransactOpts, usdcAmountToHave)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 ) returns(uint256)
func (_PositionManagerContract *PositionManagerContractTransactor) Sell(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _PositionManagerContract.contract.Transact(opts, "sell", arg0)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 ) returns(uint256)
func (_PositionManagerContract *PositionManagerContractSession) Sell(arg0 *big.Int) (*types.Transaction, error) {
	return _PositionManagerContract.Contract.Sell(&_PositionManagerContract.TransactOpts, arg0)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 ) returns(uint256)
func (_PositionManagerContract *PositionManagerContractTransactorSession) Sell(arg0 *big.Int) (*types.Transaction, error) {
	return _PositionManagerContract.Contract.Sell(&_PositionManagerContract.TransactOpts, arg0)
}
