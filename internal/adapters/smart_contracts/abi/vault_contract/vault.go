// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault_contract

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
)

// PositionManagerStats is an auto generated low-level Go binding around an user-defined struct.
type PositionManagerStats struct {
	PositionManagerAddress common.Address
	PositionWorth          *big.Int
	CostBasis              *big.Int
	Pnl                    *big.Int
	TokenExposures         []TokenExposure
	TokenAllocation        []TokenAllocation
	CanRebalance           bool
}

// RebalanceQueueData is an auto generated low-level Go binding around an user-defined struct.
type RebalanceQueueData struct {
	PositionManager  common.Address
	UsdcAmountToHave *big.Int
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

// VaultStats is an auto generated low-level Go binding around an user-defined struct.
type VaultStats struct {
	VaultAddress       common.Address
	PositionManagers   []PositionManagerStats
	VaultWorth         *big.Int
	AvailableLiquidity *big.Int
}

// VaultContractMetaData contains all meta data concerning the VaultContract contract.
var VaultContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_vaultName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_usdcAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdcAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvailableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositionManagers\",\"outputs\":[{\"internalType\":\"contractIPositionManager[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"positionManagers\",\"outputs\":[{\"internalType\":\"contractIPositionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIPositionManager\",\"name\":\"positionManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"usdcAmountToHave\",\"type\":\"uint256\"}],\"internalType\":\"structRebalanceQueueData[]\",\"name\":\"rebalanceQueueData\",\"type\":\"tuple[]\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"phvTokenToBurn\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPositionManager[]\",\"name\":\"_positionManagers\",\"type\":\"address[]\"}],\"name\":\"setPositionManagers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stats\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"positionManagerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"positionWorth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costBasis\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTokenExposure[]\",\"name\":\"tokenExposures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"enumPositionType\",\"name\":\"positionType\",\"type\":\"uint8\"}],\"internalType\":\"structTokenAllocation[]\",\"name\":\"tokenAllocation\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"canRebalance\",\"type\":\"bool\"}],\"internalType\":\"structPositionManagerStats[]\",\"name\":\"positionManagers\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"vaultWorth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"}],\"internalType\":\"structVaultStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultWorth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VaultContractABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultContractMetaData.ABI instead.
var VaultContractABI = VaultContractMetaData.ABI

// VaultContract is an auto generated Go binding around an Ethereum contract.
type VaultContract struct {
	VaultContractCaller     // Read-only binding to the contract
	VaultContractTransactor // Write-only binding to the contract
	VaultContractFilterer   // Log filterer for contract events
}

// VaultContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultContractSession struct {
	Contract     *VaultContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultContractCallerSession struct {
	Contract *VaultContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VaultContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultContractTransactorSession struct {
	Contract     *VaultContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VaultContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultContractRaw struct {
	Contract *VaultContract // Generic contract binding to access the raw methods on
}

// VaultContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultContractCallerRaw struct {
	Contract *VaultContractCaller // Generic read-only contract binding to access the raw methods on
}

// VaultContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultContractTransactorRaw struct {
	Contract *VaultContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultContract creates a new instance of VaultContract, bound to a specific deployed contract.
func NewVaultContract(address common.Address, backend bind.ContractBackend) (*VaultContract, error) {
	contract, err := bindVaultContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultContract{VaultContractCaller: VaultContractCaller{contract: contract}, VaultContractTransactor: VaultContractTransactor{contract: contract}, VaultContractFilterer: VaultContractFilterer{contract: contract}}, nil
}

// NewVaultContractCaller creates a new read-only instance of VaultContract, bound to a specific deployed contract.
func NewVaultContractCaller(address common.Address, caller bind.ContractCaller) (*VaultContractCaller, error) {
	contract, err := bindVaultContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultContractCaller{contract: contract}, nil
}

// NewVaultContractTransactor creates a new write-only instance of VaultContract, bound to a specific deployed contract.
func NewVaultContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultContractTransactor, error) {
	contract, err := bindVaultContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultContractTransactor{contract: contract}, nil
}

// NewVaultContractFilterer creates a new log filterer instance of VaultContract, bound to a specific deployed contract.
func NewVaultContractFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultContractFilterer, error) {
	contract, err := bindVaultContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultContractFilterer{contract: contract}, nil
}

// bindVaultContract binds a generic wrapper to an already deployed contract.
func bindVaultContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultContract *VaultContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultContract.Contract.VaultContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultContract *VaultContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultContract.Contract.VaultContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultContract *VaultContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultContract.Contract.VaultContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultContract *VaultContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultContract *VaultContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultContract *VaultContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultContract.Contract.contract.Transact(opts, method, params...)
}

// GetAvailableLiquidity is a free data retrieval call binding the contract method 0x9b745aec.
//
// Solidity: function getAvailableLiquidity() view returns(uint256)
func (_VaultContract *VaultContractCaller) GetAvailableLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "getAvailableLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableLiquidity is a free data retrieval call binding the contract method 0x9b745aec.
//
// Solidity: function getAvailableLiquidity() view returns(uint256)
func (_VaultContract *VaultContractSession) GetAvailableLiquidity() (*big.Int, error) {
	return _VaultContract.Contract.GetAvailableLiquidity(&_VaultContract.CallOpts)
}

// GetAvailableLiquidity is a free data retrieval call binding the contract method 0x9b745aec.
//
// Solidity: function getAvailableLiquidity() view returns(uint256)
func (_VaultContract *VaultContractCallerSession) GetAvailableLiquidity() (*big.Int, error) {
	return _VaultContract.Contract.GetAvailableLiquidity(&_VaultContract.CallOpts)
}

// GetPositionManagers is a free data retrieval call binding the contract method 0xfdbe02ba.
//
// Solidity: function getPositionManagers() view returns(address[])
func (_VaultContract *VaultContractCaller) GetPositionManagers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "getPositionManagers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPositionManagers is a free data retrieval call binding the contract method 0xfdbe02ba.
//
// Solidity: function getPositionManagers() view returns(address[])
func (_VaultContract *VaultContractSession) GetPositionManagers() ([]common.Address, error) {
	return _VaultContract.Contract.GetPositionManagers(&_VaultContract.CallOpts)
}

// GetPositionManagers is a free data retrieval call binding the contract method 0xfdbe02ba.
//
// Solidity: function getPositionManagers() view returns(address[])
func (_VaultContract *VaultContractCallerSession) GetPositionManagers() ([]common.Address, error) {
	return _VaultContract.Contract.GetPositionManagers(&_VaultContract.CallOpts)
}

// PositionManagers is a free data retrieval call binding the contract method 0x20e467f6.
//
// Solidity: function positionManagers(uint256 ) view returns(address)
func (_VaultContract *VaultContractCaller) PositionManagers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "positionManagers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManagers is a free data retrieval call binding the contract method 0x20e467f6.
//
// Solidity: function positionManagers(uint256 ) view returns(address)
func (_VaultContract *VaultContractSession) PositionManagers(arg0 *big.Int) (common.Address, error) {
	return _VaultContract.Contract.PositionManagers(&_VaultContract.CallOpts, arg0)
}

// PositionManagers is a free data retrieval call binding the contract method 0x20e467f6.
//
// Solidity: function positionManagers(uint256 ) view returns(address)
func (_VaultContract *VaultContractCallerSession) PositionManagers(arg0 *big.Int) (common.Address, error) {
	return _VaultContract.Contract.PositionManagers(&_VaultContract.CallOpts, arg0)
}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns((address,(address,uint256,uint256,int256,(int256,address,string)[],(uint256,address,string,uint256,uint8)[],bool)[],uint256,uint256))
func (_VaultContract *VaultContractCaller) Stats(opts *bind.CallOpts) (VaultStats, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "stats")

	if err != nil {
		return *new(VaultStats), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultStats)).(*VaultStats)

	return out0, err

}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns((address,(address,uint256,uint256,int256,(int256,address,string)[],(uint256,address,string,uint256,uint8)[],bool)[],uint256,uint256))
func (_VaultContract *VaultContractSession) Stats() (VaultStats, error) {
	return _VaultContract.Contract.Stats(&_VaultContract.CallOpts)
}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns((address,(address,uint256,uint256,int256,(int256,address,string)[],(uint256,address,string,uint256,uint8)[],bool)[],uint256,uint256))
func (_VaultContract *VaultContractCallerSession) Stats() (VaultStats, error) {
	return _VaultContract.Contract.Stats(&_VaultContract.CallOpts)
}

// VaultName is a free data retrieval call binding the contract method 0x0ace9ca0.
//
// Solidity: function vaultName() view returns(string)
func (_VaultContract *VaultContractCaller) VaultName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "vaultName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VaultName is a free data retrieval call binding the contract method 0x0ace9ca0.
//
// Solidity: function vaultName() view returns(string)
func (_VaultContract *VaultContractSession) VaultName() (string, error) {
	return _VaultContract.Contract.VaultName(&_VaultContract.CallOpts)
}

// VaultName is a free data retrieval call binding the contract method 0x0ace9ca0.
//
// Solidity: function vaultName() view returns(string)
func (_VaultContract *VaultContractCallerSession) VaultName() (string, error) {
	return _VaultContract.Contract.VaultName(&_VaultContract.CallOpts)
}

// VaultWorth is a free data retrieval call binding the contract method 0x7254a078.
//
// Solidity: function vaultWorth() view returns(uint256)
func (_VaultContract *VaultContractCaller) VaultWorth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "vaultWorth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultWorth is a free data retrieval call binding the contract method 0x7254a078.
//
// Solidity: function vaultWorth() view returns(uint256)
func (_VaultContract *VaultContractSession) VaultWorth() (*big.Int, error) {
	return _VaultContract.Contract.VaultWorth(&_VaultContract.CallOpts)
}

// VaultWorth is a free data retrieval call binding the contract method 0x7254a078.
//
// Solidity: function vaultWorth() view returns(uint256)
func (_VaultContract *VaultContractCallerSession) VaultWorth() (*big.Int, error) {
	return _VaultContract.Contract.VaultWorth(&_VaultContract.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 usdcAmount) returns()
func (_VaultContract *VaultContractTransactor) AddLiquidity(opts *bind.TransactOpts, usdcAmount *big.Int) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "addLiquidity", usdcAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 usdcAmount) returns()
func (_VaultContract *VaultContractSession) AddLiquidity(usdcAmount *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.AddLiquidity(&_VaultContract.TransactOpts, usdcAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 usdcAmount) returns()
func (_VaultContract *VaultContractTransactorSession) AddLiquidity(usdcAmount *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.AddLiquidity(&_VaultContract.TransactOpts, usdcAmount)
}

// Rebalance is a paid mutator transaction binding the contract method 0xf009a1e0.
//
// Solidity: function rebalance((address,uint256)[] rebalanceQueueData) returns()
func (_VaultContract *VaultContractTransactor) Rebalance(opts *bind.TransactOpts, rebalanceQueueData []RebalanceQueueData) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "rebalance", rebalanceQueueData)
}

// Rebalance is a paid mutator transaction binding the contract method 0xf009a1e0.
//
// Solidity: function rebalance((address,uint256)[] rebalanceQueueData) returns()
func (_VaultContract *VaultContractSession) Rebalance(rebalanceQueueData []RebalanceQueueData) (*types.Transaction, error) {
	return _VaultContract.Contract.Rebalance(&_VaultContract.TransactOpts, rebalanceQueueData)
}

// Rebalance is a paid mutator transaction binding the contract method 0xf009a1e0.
//
// Solidity: function rebalance((address,uint256)[] rebalanceQueueData) returns()
func (_VaultContract *VaultContractTransactorSession) Rebalance(rebalanceQueueData []RebalanceQueueData) (*types.Transaction, error) {
	return _VaultContract.Contract.Rebalance(&_VaultContract.TransactOpts, rebalanceQueueData)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 phvTokenToBurn) returns()
func (_VaultContract *VaultContractTransactor) RemoveLiquidity(opts *bind.TransactOpts, phvTokenToBurn *big.Int) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "removeLiquidity", phvTokenToBurn)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 phvTokenToBurn) returns()
func (_VaultContract *VaultContractSession) RemoveLiquidity(phvTokenToBurn *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.RemoveLiquidity(&_VaultContract.TransactOpts, phvTokenToBurn)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 phvTokenToBurn) returns()
func (_VaultContract *VaultContractTransactorSession) RemoveLiquidity(phvTokenToBurn *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.RemoveLiquidity(&_VaultContract.TransactOpts, phvTokenToBurn)
}

// SetPositionManagers is a paid mutator transaction binding the contract method 0xb38e85f8.
//
// Solidity: function setPositionManagers(address[] _positionManagers) returns()
func (_VaultContract *VaultContractTransactor) SetPositionManagers(opts *bind.TransactOpts, _positionManagers []common.Address) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "setPositionManagers", _positionManagers)
}

// SetPositionManagers is a paid mutator transaction binding the contract method 0xb38e85f8.
//
// Solidity: function setPositionManagers(address[] _positionManagers) returns()
func (_VaultContract *VaultContractSession) SetPositionManagers(_positionManagers []common.Address) (*types.Transaction, error) {
	return _VaultContract.Contract.SetPositionManagers(&_VaultContract.TransactOpts, _positionManagers)
}

// SetPositionManagers is a paid mutator transaction binding the contract method 0xb38e85f8.
//
// Solidity: function setPositionManagers(address[] _positionManagers) returns()
func (_VaultContract *VaultContractTransactorSession) SetPositionManagers(_positionManagers []common.Address) (*types.Transaction, error) {
	return _VaultContract.Contract.SetPositionManagers(&_VaultContract.TransactOpts, _positionManagers)
}
