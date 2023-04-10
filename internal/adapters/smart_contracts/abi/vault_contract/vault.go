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
	PositionsWorth     *big.Int
	CostBasis          *big.Int
	Pnl                *big.Int
}

// VaultContractMetaData contains all meta data concerning the VaultContract contract.
var VaultContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"positionManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"errorMessage\",\"type\":\"string\"}],\"name\":\"PositionManagerCannotRebalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TotalCompound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdcAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountToRebalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compoundedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialGasLeft\",\"type\":\"uint256\"}],\"name\":\"estimateGasCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvailableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositionManagers\",\"outputs\":[{\"internalType\":\"contractIPositionManager[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_vaultName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_usdcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceUtilsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethPriceFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rebalancePercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_wethTokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pnl\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"positionManagers\",\"outputs\":[{\"internalType\":\"contractIPositionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionsWorth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIPositionManager\",\"name\":\"positionManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"usdcAmountToHave\",\"type\":\"uint256\"}],\"internalType\":\"structRebalanceQueueData[]\",\"name\":\"rebalanceQueueData\",\"type\":\"tuple[]\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalancePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"phvTokenToBurn\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPositionManager[]\",\"name\":\"_positionManagers\",\"type\":\"address[]\"}],\"name\":\"setPositionManagers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIPositionManager\",\"name\":\"positionManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"usdcAmountToHave\",\"type\":\"uint256\"}],\"internalType\":\"structRebalanceQueueData[]\",\"name\":\"rebalanceQueueData\",\"type\":\"tuple[]\"}],\"name\":\"shouldRebalance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stats\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"positionManagerAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"positionWorth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costBasis\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structTokenExposure[]\",\"name\":\"tokenExposures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"leverage\",\"type\":\"uint256\"},{\"internalType\":\"enumPositionType\",\"name\":\"positionType\",\"type\":\"uint8\"}],\"internalType\":\"structTokenAllocation[]\",\"name\":\"tokenAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loanWorth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationLevel\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structPositionManagerStats[]\",\"name\":\"positionManagers\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"vaultWorth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionsWorth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"costBasis\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"}],\"internalType\":\"structVaultStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferToOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultCostBasis\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultPositionWorth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultWorth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	parsed, err := VaultContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// AmountToRebalance is a free data retrieval call binding the contract method 0x050ff61a.
//
// Solidity: function amountToRebalance() view returns(uint256)
func (_VaultContract *VaultContractCaller) AmountToRebalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "amountToRebalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountToRebalance is a free data retrieval call binding the contract method 0x050ff61a.
//
// Solidity: function amountToRebalance() view returns(uint256)
func (_VaultContract *VaultContractSession) AmountToRebalance() (*big.Int, error) {
	return _VaultContract.Contract.AmountToRebalance(&_VaultContract.CallOpts)
}

// AmountToRebalance is a free data retrieval call binding the contract method 0x050ff61a.
//
// Solidity: function amountToRebalance() view returns(uint256)
func (_VaultContract *VaultContractCallerSession) AmountToRebalance() (*big.Int, error) {
	return _VaultContract.Contract.AmountToRebalance(&_VaultContract.CallOpts)
}

// CompoundedAmount is a free data retrieval call binding the contract method 0x6b4c4574.
//
// Solidity: function compoundedAmount() view returns(uint256)
func (_VaultContract *VaultContractCaller) CompoundedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "compoundedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompoundedAmount is a free data retrieval call binding the contract method 0x6b4c4574.
//
// Solidity: function compoundedAmount() view returns(uint256)
func (_VaultContract *VaultContractSession) CompoundedAmount() (*big.Int, error) {
	return _VaultContract.Contract.CompoundedAmount(&_VaultContract.CallOpts)
}

// CompoundedAmount is a free data retrieval call binding the contract method 0x6b4c4574.
//
// Solidity: function compoundedAmount() view returns(uint256)
func (_VaultContract *VaultContractCallerSession) CompoundedAmount() (*big.Int, error) {
	return _VaultContract.Contract.CompoundedAmount(&_VaultContract.CallOpts)
}

// EstimateGasCost is a free data retrieval call binding the contract method 0xbd4725a6.
//
// Solidity: function estimateGasCost(uint256 initialGasLeft) view returns(uint256)
func (_VaultContract *VaultContractCaller) EstimateGasCost(opts *bind.CallOpts, initialGasLeft *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "estimateGasCost", initialGasLeft)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateGasCost is a free data retrieval call binding the contract method 0xbd4725a6.
//
// Solidity: function estimateGasCost(uint256 initialGasLeft) view returns(uint256)
func (_VaultContract *VaultContractSession) EstimateGasCost(initialGasLeft *big.Int) (*big.Int, error) {
	return _VaultContract.Contract.EstimateGasCost(&_VaultContract.CallOpts, initialGasLeft)
}

// EstimateGasCost is a free data retrieval call binding the contract method 0xbd4725a6.
//
// Solidity: function estimateGasCost(uint256 initialGasLeft) view returns(uint256)
func (_VaultContract *VaultContractCallerSession) EstimateGasCost(initialGasLeft *big.Int) (*big.Int, error) {
	return _VaultContract.Contract.EstimateGasCost(&_VaultContract.CallOpts, initialGasLeft)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VaultContract *VaultContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VaultContract *VaultContractSession) Owner() (common.Address, error) {
	return _VaultContract.Contract.Owner(&_VaultContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VaultContract *VaultContractCallerSession) Owner() (common.Address, error) {
	return _VaultContract.Contract.Owner(&_VaultContract.CallOpts)
}

// Pnl is a free data retrieval call binding the contract method 0x3056f68b.
//
// Solidity: function pnl() view returns(int256)
func (_VaultContract *VaultContractCaller) Pnl(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "pnl")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pnl is a free data retrieval call binding the contract method 0x3056f68b.
//
// Solidity: function pnl() view returns(int256)
func (_VaultContract *VaultContractSession) Pnl() (*big.Int, error) {
	return _VaultContract.Contract.Pnl(&_VaultContract.CallOpts)
}

// Pnl is a free data retrieval call binding the contract method 0x3056f68b.
//
// Solidity: function pnl() view returns(int256)
func (_VaultContract *VaultContractCallerSession) Pnl() (*big.Int, error) {
	return _VaultContract.Contract.Pnl(&_VaultContract.CallOpts)
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

// PositionsWorth is a free data retrieval call binding the contract method 0x043998bc.
//
// Solidity: function positionsWorth() view returns(uint256)
func (_VaultContract *VaultContractCaller) PositionsWorth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "positionsWorth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PositionsWorth is a free data retrieval call binding the contract method 0x043998bc.
//
// Solidity: function positionsWorth() view returns(uint256)
func (_VaultContract *VaultContractSession) PositionsWorth() (*big.Int, error) {
	return _VaultContract.Contract.PositionsWorth(&_VaultContract.CallOpts)
}

// PositionsWorth is a free data retrieval call binding the contract method 0x043998bc.
//
// Solidity: function positionsWorth() view returns(uint256)
func (_VaultContract *VaultContractCallerSession) PositionsWorth() (*big.Int, error) {
	return _VaultContract.Contract.PositionsWorth(&_VaultContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VaultContract *VaultContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VaultContract *VaultContractSession) ProxiableUUID() ([32]byte, error) {
	return _VaultContract.Contract.ProxiableUUID(&_VaultContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VaultContract *VaultContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VaultContract.Contract.ProxiableUUID(&_VaultContract.CallOpts)
}

// RebalancePercent is a free data retrieval call binding the contract method 0xf38c5ccc.
//
// Solidity: function rebalancePercent() view returns(uint256)
func (_VaultContract *VaultContractCaller) RebalancePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "rebalancePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebalancePercent is a free data retrieval call binding the contract method 0xf38c5ccc.
//
// Solidity: function rebalancePercent() view returns(uint256)
func (_VaultContract *VaultContractSession) RebalancePercent() (*big.Int, error) {
	return _VaultContract.Contract.RebalancePercent(&_VaultContract.CallOpts)
}

// RebalancePercent is a free data retrieval call binding the contract method 0xf38c5ccc.
//
// Solidity: function rebalancePercent() view returns(uint256)
func (_VaultContract *VaultContractCallerSession) RebalancePercent() (*big.Int, error) {
	return _VaultContract.Contract.RebalancePercent(&_VaultContract.CallOpts)
}

// ShouldRebalance is a free data retrieval call binding the contract method 0x13371519.
//
// Solidity: function shouldRebalance((address,uint256)[] rebalanceQueueData) view returns(bool, string)
func (_VaultContract *VaultContractCaller) ShouldRebalance(opts *bind.CallOpts, rebalanceQueueData []RebalanceQueueData) (bool, string, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "shouldRebalance", rebalanceQueueData)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// ShouldRebalance is a free data retrieval call binding the contract method 0x13371519.
//
// Solidity: function shouldRebalance((address,uint256)[] rebalanceQueueData) view returns(bool, string)
func (_VaultContract *VaultContractSession) ShouldRebalance(rebalanceQueueData []RebalanceQueueData) (bool, string, error) {
	return _VaultContract.Contract.ShouldRebalance(&_VaultContract.CallOpts, rebalanceQueueData)
}

// ShouldRebalance is a free data retrieval call binding the contract method 0x13371519.
//
// Solidity: function shouldRebalance((address,uint256)[] rebalanceQueueData) view returns(bool, string)
func (_VaultContract *VaultContractCallerSession) ShouldRebalance(rebalanceQueueData []RebalanceQueueData) (bool, string, error) {
	return _VaultContract.Contract.ShouldRebalance(&_VaultContract.CallOpts, rebalanceQueueData)
}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns((address,(address,string,uint256,uint256,int256,(int256,address,string)[],(uint256,address,string,uint256,uint8)[],uint256,uint256,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,int256))
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
// Solidity: function stats() view returns((address,(address,string,uint256,uint256,int256,(int256,address,string)[],(uint256,address,string,uint256,uint8)[],uint256,uint256,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,int256))
func (_VaultContract *VaultContractSession) Stats() (VaultStats, error) {
	return _VaultContract.Contract.Stats(&_VaultContract.CallOpts)
}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns((address,(address,string,uint256,uint256,int256,(int256,address,string)[],(uint256,address,string,uint256,uint8)[],uint256,uint256,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,int256))
func (_VaultContract *VaultContractCallerSession) Stats() (VaultStats, error) {
	return _VaultContract.Contract.Stats(&_VaultContract.CallOpts)
}

// VaultCostBasis is a free data retrieval call binding the contract method 0x67c4ca99.
//
// Solidity: function vaultCostBasis() view returns(uint256)
func (_VaultContract *VaultContractCaller) VaultCostBasis(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "vaultCostBasis")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultCostBasis is a free data retrieval call binding the contract method 0x67c4ca99.
//
// Solidity: function vaultCostBasis() view returns(uint256)
func (_VaultContract *VaultContractSession) VaultCostBasis() (*big.Int, error) {
	return _VaultContract.Contract.VaultCostBasis(&_VaultContract.CallOpts)
}

// VaultCostBasis is a free data retrieval call binding the contract method 0x67c4ca99.
//
// Solidity: function vaultCostBasis() view returns(uint256)
func (_VaultContract *VaultContractCallerSession) VaultCostBasis() (*big.Int, error) {
	return _VaultContract.Contract.VaultCostBasis(&_VaultContract.CallOpts)
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

// VaultPositionWorth is a free data retrieval call binding the contract method 0xab6cb37c.
//
// Solidity: function vaultPositionWorth() view returns(uint256)
func (_VaultContract *VaultContractCaller) VaultPositionWorth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "vaultPositionWorth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultPositionWorth is a free data retrieval call binding the contract method 0xab6cb37c.
//
// Solidity: function vaultPositionWorth() view returns(uint256)
func (_VaultContract *VaultContractSession) VaultPositionWorth() (*big.Int, error) {
	return _VaultContract.Contract.VaultPositionWorth(&_VaultContract.CallOpts)
}

// VaultPositionWorth is a free data retrieval call binding the contract method 0xab6cb37c.
//
// Solidity: function vaultPositionWorth() view returns(uint256)
func (_VaultContract *VaultContractCallerSession) VaultPositionWorth() (*big.Int, error) {
	return _VaultContract.Contract.VaultPositionWorth(&_VaultContract.CallOpts)
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

// WethToken is a free data retrieval call binding the contract method 0x4b57b0be.
//
// Solidity: function wethToken() view returns(address)
func (_VaultContract *VaultContractCaller) WethToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "wethToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethToken is a free data retrieval call binding the contract method 0x4b57b0be.
//
// Solidity: function wethToken() view returns(address)
func (_VaultContract *VaultContractSession) WethToken() (common.Address, error) {
	return _VaultContract.Contract.WethToken(&_VaultContract.CallOpts)
}

// WethToken is a free data retrieval call binding the contract method 0x4b57b0be.
//
// Solidity: function wethToken() view returns(address)
func (_VaultContract *VaultContractCallerSession) WethToken() (common.Address, error) {
	return _VaultContract.Contract.WethToken(&_VaultContract.CallOpts)
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

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns(uint256)
func (_VaultContract *VaultContractTransactor) Compound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "compound")
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns(uint256)
func (_VaultContract *VaultContractSession) Compound() (*types.Transaction, error) {
	return _VaultContract.Contract.Compound(&_VaultContract.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns(uint256)
func (_VaultContract *VaultContractTransactorSession) Compound() (*types.Transaction, error) {
	return _VaultContract.Contract.Compound(&_VaultContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x74b3bd8c.
//
// Solidity: function initialize(string _vaultName, address _usdcAddress, address _priceUtilsAddress, address _ethPriceFeedAddress, uint256 _rebalancePercent, address _wethTokenAddress) returns()
func (_VaultContract *VaultContractTransactor) Initialize(opts *bind.TransactOpts, _vaultName string, _usdcAddress common.Address, _priceUtilsAddress common.Address, _ethPriceFeedAddress common.Address, _rebalancePercent *big.Int, _wethTokenAddress common.Address) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "initialize", _vaultName, _usdcAddress, _priceUtilsAddress, _ethPriceFeedAddress, _rebalancePercent, _wethTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x74b3bd8c.
//
// Solidity: function initialize(string _vaultName, address _usdcAddress, address _priceUtilsAddress, address _ethPriceFeedAddress, uint256 _rebalancePercent, address _wethTokenAddress) returns()
func (_VaultContract *VaultContractSession) Initialize(_vaultName string, _usdcAddress common.Address, _priceUtilsAddress common.Address, _ethPriceFeedAddress common.Address, _rebalancePercent *big.Int, _wethTokenAddress common.Address) (*types.Transaction, error) {
	return _VaultContract.Contract.Initialize(&_VaultContract.TransactOpts, _vaultName, _usdcAddress, _priceUtilsAddress, _ethPriceFeedAddress, _rebalancePercent, _wethTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x74b3bd8c.
//
// Solidity: function initialize(string _vaultName, address _usdcAddress, address _priceUtilsAddress, address _ethPriceFeedAddress, uint256 _rebalancePercent, address _wethTokenAddress) returns()
func (_VaultContract *VaultContractTransactorSession) Initialize(_vaultName string, _usdcAddress common.Address, _priceUtilsAddress common.Address, _ethPriceFeedAddress common.Address, _rebalancePercent *big.Int, _wethTokenAddress common.Address) (*types.Transaction, error) {
	return _VaultContract.Contract.Initialize(&_VaultContract.TransactOpts, _vaultName, _usdcAddress, _priceUtilsAddress, _ethPriceFeedAddress, _rebalancePercent, _wethTokenAddress)
}

// Liquidate is a paid mutator transaction binding the contract method 0x28a07025.
//
// Solidity: function liquidate() returns()
func (_VaultContract *VaultContractTransactor) Liquidate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "liquidate")
}

// Liquidate is a paid mutator transaction binding the contract method 0x28a07025.
//
// Solidity: function liquidate() returns()
func (_VaultContract *VaultContractSession) Liquidate() (*types.Transaction, error) {
	return _VaultContract.Contract.Liquidate(&_VaultContract.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0x28a07025.
//
// Solidity: function liquidate() returns()
func (_VaultContract *VaultContractTransactorSession) Liquidate() (*types.Transaction, error) {
	return _VaultContract.Contract.Liquidate(&_VaultContract.TransactOpts)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VaultContract *VaultContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VaultContract *VaultContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _VaultContract.Contract.RenounceOwnership(&_VaultContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VaultContract *VaultContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VaultContract.Contract.RenounceOwnership(&_VaultContract.TransactOpts)
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

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VaultContract *VaultContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VaultContract *VaultContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VaultContract.Contract.TransferOwnership(&_VaultContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VaultContract *VaultContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VaultContract.Contract.TransferOwnership(&_VaultContract.TransactOpts, newOwner)
}

// TransferToOwner is a paid mutator transaction binding the contract method 0x2d90ae94.
//
// Solidity: function transferToOwner() returns()
func (_VaultContract *VaultContractTransactor) TransferToOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "transferToOwner")
}

// TransferToOwner is a paid mutator transaction binding the contract method 0x2d90ae94.
//
// Solidity: function transferToOwner() returns()
func (_VaultContract *VaultContractSession) TransferToOwner() (*types.Transaction, error) {
	return _VaultContract.Contract.TransferToOwner(&_VaultContract.TransactOpts)
}

// TransferToOwner is a paid mutator transaction binding the contract method 0x2d90ae94.
//
// Solidity: function transferToOwner() returns()
func (_VaultContract *VaultContractTransactorSession) TransferToOwner() (*types.Transaction, error) {
	return _VaultContract.Contract.TransferToOwner(&_VaultContract.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VaultContract *VaultContractTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VaultContract *VaultContractSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VaultContract.Contract.UpgradeTo(&_VaultContract.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VaultContract *VaultContractTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VaultContract.Contract.UpgradeTo(&_VaultContract.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VaultContract *VaultContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VaultContract *VaultContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VaultContract.Contract.UpgradeToAndCall(&_VaultContract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VaultContract *VaultContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VaultContract.Contract.UpgradeToAndCall(&_VaultContract.TransactOpts, newImplementation, data)
}

// VaultContractAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the VaultContract contract.
type VaultContractAdminChangedIterator struct {
	Event *VaultContractAdminChanged // Event containing the contract specifics and raw log

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
func (it *VaultContractAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultContractAdminChanged)
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
		it.Event = new(VaultContractAdminChanged)
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
func (it *VaultContractAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultContractAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultContractAdminChanged represents a AdminChanged event raised by the VaultContract contract.
type VaultContractAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VaultContract *VaultContractFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VaultContractAdminChangedIterator, error) {

	logs, sub, err := _VaultContract.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VaultContractAdminChangedIterator{contract: _VaultContract.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VaultContract *VaultContractFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VaultContractAdminChanged) (event.Subscription, error) {

	logs, sub, err := _VaultContract.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultContractAdminChanged)
				if err := _VaultContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VaultContract *VaultContractFilterer) ParseAdminChanged(log types.Log) (*VaultContractAdminChanged, error) {
	event := new(VaultContractAdminChanged)
	if err := _VaultContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultContractBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the VaultContract contract.
type VaultContractBeaconUpgradedIterator struct {
	Event *VaultContractBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VaultContractBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultContractBeaconUpgraded)
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
		it.Event = new(VaultContractBeaconUpgraded)
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
func (it *VaultContractBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultContractBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultContractBeaconUpgraded represents a BeaconUpgraded event raised by the VaultContract contract.
type VaultContractBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VaultContract *VaultContractFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VaultContractBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VaultContract.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VaultContractBeaconUpgradedIterator{contract: _VaultContract.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VaultContract *VaultContractFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VaultContractBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VaultContract.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultContractBeaconUpgraded)
				if err := _VaultContract.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VaultContract *VaultContractFilterer) ParseBeaconUpgraded(log types.Log) (*VaultContractBeaconUpgraded, error) {
	event := new(VaultContractBeaconUpgraded)
	if err := _VaultContract.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VaultContract contract.
type VaultContractInitializedIterator struct {
	Event *VaultContractInitialized // Event containing the contract specifics and raw log

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
func (it *VaultContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultContractInitialized)
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
		it.Event = new(VaultContractInitialized)
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
func (it *VaultContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultContractInitialized represents a Initialized event raised by the VaultContract contract.
type VaultContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VaultContract *VaultContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*VaultContractInitializedIterator, error) {

	logs, sub, err := _VaultContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VaultContractInitializedIterator{contract: _VaultContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VaultContract *VaultContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VaultContractInitialized) (event.Subscription, error) {

	logs, sub, err := _VaultContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultContractInitialized)
				if err := _VaultContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VaultContract *VaultContractFilterer) ParseInitialized(log types.Log) (*VaultContractInitialized, error) {
	event := new(VaultContractInitialized)
	if err := _VaultContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VaultContract contract.
type VaultContractOwnershipTransferredIterator struct {
	Event *VaultContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VaultContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultContractOwnershipTransferred)
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
		it.Event = new(VaultContractOwnershipTransferred)
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
func (it *VaultContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultContractOwnershipTransferred represents a OwnershipTransferred event raised by the VaultContract contract.
type VaultContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VaultContract *VaultContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VaultContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VaultContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VaultContractOwnershipTransferredIterator{contract: _VaultContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VaultContract *VaultContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VaultContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VaultContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultContractOwnershipTransferred)
				if err := _VaultContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VaultContract *VaultContractFilterer) ParseOwnershipTransferred(log types.Log) (*VaultContractOwnershipTransferred, error) {
	event := new(VaultContractOwnershipTransferred)
	if err := _VaultContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultContractTotalCompoundIterator is returned from FilterTotalCompound and is used to iterate over the raw logs and unpacked data for TotalCompound events raised by the VaultContract contract.
type VaultContractTotalCompoundIterator struct {
	Event *VaultContractTotalCompound // Event containing the contract specifics and raw log

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
func (it *VaultContractTotalCompoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultContractTotalCompound)
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
		it.Event = new(VaultContractTotalCompound)
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
func (it *VaultContractTotalCompoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultContractTotalCompoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultContractTotalCompound represents a TotalCompound event raised by the VaultContract contract.
type VaultContractTotalCompound struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTotalCompound is a free log retrieval operation binding the contract event 0x8fab91bb51324eadd82d52fc89b8bffa6e32bb695a6a40f9a60be92d6b1f80c3.
//
// Solidity: event TotalCompound(uint256 amount)
func (_VaultContract *VaultContractFilterer) FilterTotalCompound(opts *bind.FilterOpts) (*VaultContractTotalCompoundIterator, error) {

	logs, sub, err := _VaultContract.contract.FilterLogs(opts, "TotalCompound")
	if err != nil {
		return nil, err
	}
	return &VaultContractTotalCompoundIterator{contract: _VaultContract.contract, event: "TotalCompound", logs: logs, sub: sub}, nil
}

// WatchTotalCompound is a free log subscription operation binding the contract event 0x8fab91bb51324eadd82d52fc89b8bffa6e32bb695a6a40f9a60be92d6b1f80c3.
//
// Solidity: event TotalCompound(uint256 amount)
func (_VaultContract *VaultContractFilterer) WatchTotalCompound(opts *bind.WatchOpts, sink chan<- *VaultContractTotalCompound) (event.Subscription, error) {

	logs, sub, err := _VaultContract.contract.WatchLogs(opts, "TotalCompound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultContractTotalCompound)
				if err := _VaultContract.contract.UnpackLog(event, "TotalCompound", log); err != nil {
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

// ParseTotalCompound is a log parse operation binding the contract event 0x8fab91bb51324eadd82d52fc89b8bffa6e32bb695a6a40f9a60be92d6b1f80c3.
//
// Solidity: event TotalCompound(uint256 amount)
func (_VaultContract *VaultContractFilterer) ParseTotalCompound(log types.Log) (*VaultContractTotalCompound, error) {
	event := new(VaultContractTotalCompound)
	if err := _VaultContract.contract.UnpackLog(event, "TotalCompound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VaultContract contract.
type VaultContractUpgradedIterator struct {
	Event *VaultContractUpgraded // Event containing the contract specifics and raw log

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
func (it *VaultContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultContractUpgraded)
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
		it.Event = new(VaultContractUpgraded)
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
func (it *VaultContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultContractUpgraded represents a Upgraded event raised by the VaultContract contract.
type VaultContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VaultContract *VaultContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VaultContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VaultContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VaultContractUpgradedIterator{contract: _VaultContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VaultContract *VaultContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VaultContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VaultContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultContractUpgraded)
				if err := _VaultContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VaultContract *VaultContractFilterer) ParseUpgraded(log types.Log) (*VaultContractUpgraded, error) {
	event := new(VaultContractUpgraded)
	if err := _VaultContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
