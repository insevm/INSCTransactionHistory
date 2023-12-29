package aabi

import (
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"strings"
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

// Erc721gpMetaData contains all meta data concerning the Erc721gp contract.
var Erc721gpMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"addLaunchpadSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"adminMintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"atransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanLaunchpadSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"gameFallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLaunchpadSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxLaunchpadSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"launchpad\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"launchpadTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_auth\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"launchpad_\",\"type\":\"address\"}],\"name\":\"setLaunchpad\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"setNameSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Erc721gpABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc721gpMetaData.ABI instead.
var Erc721gpABI = Erc721gpMetaData.ABI

// Erc721gp is an auto generated Go binding around an Ethereum contract.
type Erc721gp struct {
	Erc721gpCaller     // Read-only binding to the contract
	Erc721gpTransactor // Write-only binding to the contract
	Erc721gpFilterer   // Log filterer for contract events
}

// Erc721gpCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721gpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721gpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721gpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721gpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721gpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721gpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721gpSession struct {
	Contract     *Erc721gp         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721gpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721gpCallerSession struct {
	Contract *Erc721gpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Erc721gpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721gpTransactorSession struct {
	Contract     *Erc721gpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Erc721gpRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721gpRaw struct {
	Contract *Erc721gp // Generic contract binding to access the raw methods on
}

// Erc721gpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721gpCallerRaw struct {
	Contract *Erc721gpCaller // Generic read-only contract binding to access the raw methods on
}

// Erc721gpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721gpTransactorRaw struct {
	Contract *Erc721gpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721gp creates a new instance of Erc721gp, bound to a specific deployed contract.
func NewErc721gp(address common.Address, backend bind.ContractBackend) (*Erc721gp, error) {
	contract, err := bindErc721gp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721gp{Erc721gpCaller: Erc721gpCaller{contract: contract}, Erc721gpTransactor: Erc721gpTransactor{contract: contract}, Erc721gpFilterer: Erc721gpFilterer{contract: contract}}, nil
}

// NewErc721gpCaller creates a new read-only instance of Erc721gp, bound to a specific deployed contract.
func NewErc721gpCaller(address common.Address, caller bind.ContractCaller) (*Erc721gpCaller, error) {
	contract, err := bindErc721gp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721gpCaller{contract: contract}, nil
}

// NewErc721gpTransactor creates a new write-only instance of Erc721gp, bound to a specific deployed contract.
func NewErc721gpTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc721gpTransactor, error) {
	contract, err := bindErc721gp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721gpTransactor{contract: contract}, nil
}

// NewErc721gpFilterer creates a new log filterer instance of Erc721gp, bound to a specific deployed contract.
func NewErc721gpFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc721gpFilterer, error) {
	contract, err := bindErc721gp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721gpFilterer{contract: contract}, nil
}

// bindErc721gp binds a generic wrapper to an already deployed contract.
func bindErc721gp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721gpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721gp *Erc721gpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721gp.Contract.Erc721gpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721gp *Erc721gpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721gp.Contract.Erc721gpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721gp *Erc721gpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721gp.Contract.Erc721gpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721gp *Erc721gpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721gp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721gp *Erc721gpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721gp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721gp *Erc721gpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721gp.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Erc721gp *Erc721gpCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "admin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Erc721gp *Erc721gpSession) Admin(arg0 common.Address) (bool, error) {
	return _Erc721gp.Contract.Admin(&_Erc721gp.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_Erc721gp *Erc721gpCallerSession) Admin(arg0 common.Address) (bool, error) {
	return _Erc721gp.Contract.Admin(&_Erc721gp.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721gp *Erc721gpCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721gp *Erc721gpSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc721gp.Contract.BalanceOf(&_Erc721gp.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721gp *Erc721gpCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc721gp.Contract.BalanceOf(&_Erc721gp.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721gp *Erc721gpCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721gp *Erc721gpSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc721gp.Contract.GetApproved(&_Erc721gp.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721gp *Erc721gpCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc721gp.Contract.GetApproved(&_Erc721gp.CallOpts, tokenId)
}

// GetLaunchpadSupply is a free data retrieval call binding the contract method 0xeebb28b2.
//
// Solidity: function getLaunchpadSupply() view returns(uint256)
func (_Erc721gp *Erc721gpCaller) GetLaunchpadSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "getLaunchpadSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLaunchpadSupply is a free data retrieval call binding the contract method 0xeebb28b2.
//
// Solidity: function getLaunchpadSupply() view returns(uint256)
func (_Erc721gp *Erc721gpSession) GetLaunchpadSupply() (*big.Int, error) {
	return _Erc721gp.Contract.GetLaunchpadSupply(&_Erc721gp.CallOpts)
}

// GetLaunchpadSupply is a free data retrieval call binding the contract method 0xeebb28b2.
//
// Solidity: function getLaunchpadSupply() view returns(uint256)
func (_Erc721gp *Erc721gpCallerSession) GetLaunchpadSupply() (*big.Int, error) {
	return _Erc721gp.Contract.GetLaunchpadSupply(&_Erc721gp.CallOpts)
}

// GetMaxLaunchpadSupply is a free data retrieval call binding the contract method 0x5b43bba1.
//
// Solidity: function getMaxLaunchpadSupply() view returns(uint256)
func (_Erc721gp *Erc721gpCaller) GetMaxLaunchpadSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "getMaxLaunchpadSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxLaunchpadSupply is a free data retrieval call binding the contract method 0x5b43bba1.
//
// Solidity: function getMaxLaunchpadSupply() view returns(uint256)
func (_Erc721gp *Erc721gpSession) GetMaxLaunchpadSupply() (*big.Int, error) {
	return _Erc721gp.Contract.GetMaxLaunchpadSupply(&_Erc721gp.CallOpts)
}

// GetMaxLaunchpadSupply is a free data retrieval call binding the contract method 0x5b43bba1.
//
// Solidity: function getMaxLaunchpadSupply() view returns(uint256)
func (_Erc721gp *Erc721gpCallerSession) GetMaxLaunchpadSupply() (*big.Int, error) {
	return _Erc721gp.Contract.GetMaxLaunchpadSupply(&_Erc721gp.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721gp *Erc721gpCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721gp *Erc721gpSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc721gp.Contract.IsApprovedForAll(&_Erc721gp.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721gp *Erc721gpCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc721gp.Contract.IsApprovedForAll(&_Erc721gp.CallOpts, owner, operator)
}

// Launchpad is a free data retrieval call binding the contract method 0x02669b52.
//
// Solidity: function launchpad() view returns(address)
func (_Erc721gp *Erc721gpCaller) Launchpad(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "launchpad")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Launchpad is a free data retrieval call binding the contract method 0x02669b52.
//
// Solidity: function launchpad() view returns(address)
func (_Erc721gp *Erc721gpSession) Launchpad() (common.Address, error) {
	return _Erc721gp.Contract.Launchpad(&_Erc721gp.CallOpts)
}

// Launchpad is a free data retrieval call binding the contract method 0x02669b52.
//
// Solidity: function launchpad() view returns(address)
func (_Erc721gp *Erc721gpCallerSession) Launchpad() (common.Address, error) {
	return _Erc721gp.Contract.Launchpad(&_Erc721gp.CallOpts)
}

// LaunchpadTokenIds is a free data retrieval call binding the contract method 0xb3a57c33.
//
// Solidity: function launchpadTokenIds(uint256 ) view returns(uint256)
func (_Erc721gp *Erc721gpCaller) LaunchpadTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "launchpadTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaunchpadTokenIds is a free data retrieval call binding the contract method 0xb3a57c33.
//
// Solidity: function launchpadTokenIds(uint256 ) view returns(uint256)
func (_Erc721gp *Erc721gpSession) LaunchpadTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Erc721gp.Contract.LaunchpadTokenIds(&_Erc721gp.CallOpts, arg0)
}

// LaunchpadTokenIds is a free data retrieval call binding the contract method 0xb3a57c33.
//
// Solidity: function launchpadTokenIds(uint256 ) view returns(uint256)
func (_Erc721gp *Erc721gpCallerSession) LaunchpadTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Erc721gp.Contract.LaunchpadTokenIds(&_Erc721gp.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721gp *Erc721gpCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721gp *Erc721gpSession) Name() (string, error) {
	return _Erc721gp.Contract.Name(&_Erc721gp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721gp *Erc721gpCallerSession) Name() (string, error) {
	return _Erc721gp.Contract.Name(&_Erc721gp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721gp *Erc721gpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721gp *Erc721gpSession) Owner() (common.Address, error) {
	return _Erc721gp.Contract.Owner(&_Erc721gp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721gp *Erc721gpCallerSession) Owner() (common.Address, error) {
	return _Erc721gp.Contract.Owner(&_Erc721gp.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721gp *Erc721gpCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721gp *Erc721gpSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721gp.Contract.OwnerOf(&_Erc721gp.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721gp *Erc721gpCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721gp.Contract.OwnerOf(&_Erc721gp.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721gp *Erc721gpCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721gp *Erc721gpSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc721gp.Contract.SupportsInterface(&_Erc721gp.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721gp *Erc721gpCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc721gp.Contract.SupportsInterface(&_Erc721gp.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721gp *Erc721gpCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721gp *Erc721gpSession) Symbol() (string, error) {
	return _Erc721gp.Contract.Symbol(&_Erc721gp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721gp *Erc721gpCallerSession) Symbol() (string, error) {
	return _Erc721gp.Contract.Symbol(&_Erc721gp.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721gp *Erc721gpCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Erc721gp.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721gp *Erc721gpSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc721gp.Contract.TokenURI(&_Erc721gp.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721gp *Erc721gpCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc721gp.Contract.TokenURI(&_Erc721gp.CallOpts, tokenId)
}

// AddLaunchpadSupply is a paid mutator transaction binding the contract method 0x091dfac8.
//
// Solidity: function addLaunchpadSupply(uint256[] tokenIds) returns()
func (_Erc721gp *Erc721gpTransactor) AddLaunchpadSupply(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "addLaunchpadSupply", tokenIds)
}

// AddLaunchpadSupply is a paid mutator transaction binding the contract method 0x091dfac8.
//
// Solidity: function addLaunchpadSupply(uint256[] tokenIds) returns()
func (_Erc721gp *Erc721gpSession) AddLaunchpadSupply(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.AddLaunchpadSupply(&_Erc721gp.TransactOpts, tokenIds)
}

// AddLaunchpadSupply is a paid mutator transaction binding the contract method 0x091dfac8.
//
// Solidity: function addLaunchpadSupply(uint256[] tokenIds) returns()
func (_Erc721gp *Erc721gpTransactorSession) AddLaunchpadSupply(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.AddLaunchpadSupply(&_Erc721gp.TransactOpts, tokenIds)
}

// AdminMintTo is a paid mutator transaction binding the contract method 0x3d6a5745.
//
// Solidity: function adminMintTo(address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactor) AdminMintTo(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "adminMintTo", to, tokenId)
}

// AdminMintTo is a paid mutator transaction binding the contract method 0x3d6a5745.
//
// Solidity: function adminMintTo(address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpSession) AdminMintTo(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.AdminMintTo(&_Erc721gp.TransactOpts, to, tokenId)
}

// AdminMintTo is a paid mutator transaction binding the contract method 0x3d6a5745.
//
// Solidity: function adminMintTo(address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactorSession) AdminMintTo(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.AdminMintTo(&_Erc721gp.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.Approve(&_Erc721gp.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.Approve(&_Erc721gp.TransactOpts, to, tokenId)
}

// AtransferFrom is a paid mutator transaction binding the contract method 0x04d8c29f.
//
// Solidity: function atransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactor) AtransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "atransferFrom", from, to, tokenId)
}

// AtransferFrom is a paid mutator transaction binding the contract method 0x04d8c29f.
//
// Solidity: function atransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpSession) AtransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.AtransferFrom(&_Erc721gp.TransactOpts, from, to, tokenId)
}

// AtransferFrom is a paid mutator transaction binding the contract method 0x04d8c29f.
//
// Solidity: function atransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactorSession) AtransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.AtransferFrom(&_Erc721gp.TransactOpts, from, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _id) returns()
func (_Erc721gp *Erc721gpTransactor) Burn(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "burn", _id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _id) returns()
func (_Erc721gp *Erc721gpSession) Burn(_id *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.Burn(&_Erc721gp.TransactOpts, _id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _id) returns()
func (_Erc721gp *Erc721gpTransactorSession) Burn(_id *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.Burn(&_Erc721gp.TransactOpts, _id)
}

// CleanLaunchpadSupply is a paid mutator transaction binding the contract method 0x2f05be8f.
//
// Solidity: function cleanLaunchpadSupply() returns()
func (_Erc721gp *Erc721gpTransactor) CleanLaunchpadSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "cleanLaunchpadSupply")
}

// CleanLaunchpadSupply is a paid mutator transaction binding the contract method 0x2f05be8f.
//
// Solidity: function cleanLaunchpadSupply() returns()
func (_Erc721gp *Erc721gpSession) CleanLaunchpadSupply() (*types.Transaction, error) {
	return _Erc721gp.Contract.CleanLaunchpadSupply(&_Erc721gp.TransactOpts)
}

// CleanLaunchpadSupply is a paid mutator transaction binding the contract method 0x2f05be8f.
//
// Solidity: function cleanLaunchpadSupply() returns()
func (_Erc721gp *Erc721gpTransactorSession) CleanLaunchpadSupply() (*types.Transaction, error) {
	return _Erc721gp.Contract.CleanLaunchpadSupply(&_Erc721gp.TransactOpts)
}

// GameFallback is a paid mutator transaction binding the contract method 0x91518477.
//
// Solidity: function gameFallback(uint256 start, uint256 end) returns()
func (_Erc721gp *Erc721gpTransactor) GameFallback(opts *bind.TransactOpts, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "gameFallback", start, end)
}

// GameFallback is a paid mutator transaction binding the contract method 0x91518477.
//
// Solidity: function gameFallback(uint256 start, uint256 end) returns()
func (_Erc721gp *Erc721gpSession) GameFallback(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.GameFallback(&_Erc721gp.TransactOpts, start, end)
}

// GameFallback is a paid mutator transaction binding the contract method 0x91518477.
//
// Solidity: function gameFallback(uint256 start, uint256 end) returns()
func (_Erc721gp *Erc721gpTransactorSession) GameFallback(start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.GameFallback(&_Erc721gp.TransactOpts, start, end)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_Erc721gp *Erc721gpTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "initialize", _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_Erc721gp *Erc721gpSession) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _Erc721gp.Contract.Initialize(&_Erc721gp.TransactOpts, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_Erc721gp *Erc721gpTransactorSession) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _Erc721gp.Contract.Initialize(&_Erc721gp.TransactOpts, _name, _symbol)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "mint", tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) returns()
func (_Erc721gp *Erc721gpSession) Mint(tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.Mint(&_Erc721gp.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactorSession) Mint(tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.Mint(&_Erc721gp.TransactOpts, tokenId)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address to, uint256 size) returns()
func (_Erc721gp *Erc721gpTransactor) MintTo(opts *bind.TransactOpts, to common.Address, size *big.Int) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "mintTo", to, size)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address to, uint256 size) returns()
func (_Erc721gp *Erc721gpSession) MintTo(to common.Address, size *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.MintTo(&_Erc721gp.TransactOpts, to, size)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address to, uint256 size) returns()
func (_Erc721gp *Erc721gpTransactorSession) MintTo(to common.Address, size *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.MintTo(&_Erc721gp.TransactOpts, to, size)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc721gp *Erc721gpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc721gp *Erc721gpSession) RenounceOwnership() (*types.Transaction, error) {
	return _Erc721gp.Contract.RenounceOwnership(&_Erc721gp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc721gp *Erc721gpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Erc721gp.Contract.RenounceOwnership(&_Erc721gp.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.SafeTransferFrom(&_Erc721gp.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.SafeTransferFrom(&_Erc721gp.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Erc721gp *Erc721gpTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Erc721gp *Erc721gpSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721gp.Contract.SafeTransferFrom0(&_Erc721gp.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Erc721gp *Erc721gpTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Erc721gp.Contract.SafeTransferFrom0(&_Erc721gp.TransactOpts, from, to, tokenId, _data)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address user, bool _auth) returns()
func (_Erc721gp *Erc721gpTransactor) SetAdmin(opts *bind.TransactOpts, user common.Address, _auth bool) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "setAdmin", user, _auth)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address user, bool _auth) returns()
func (_Erc721gp *Erc721gpSession) SetAdmin(user common.Address, _auth bool) (*types.Transaction, error) {
	return _Erc721gp.Contract.SetAdmin(&_Erc721gp.TransactOpts, user, _auth)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(address user, bool _auth) returns()
func (_Erc721gp *Erc721gpTransactorSession) SetAdmin(user common.Address, _auth bool) (*types.Transaction, error) {
	return _Erc721gp.Contract.SetAdmin(&_Erc721gp.TransactOpts, user, _auth)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721gp *Erc721gpTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721gp *Erc721gpSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721gp.Contract.SetApprovalForAll(&_Erc721gp.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721gp *Erc721gpTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721gp.Contract.SetApprovalForAll(&_Erc721gp.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _uri) returns()
func (_Erc721gp *Erc721gpTransactor) SetBaseURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "setBaseURI", _uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _uri) returns()
func (_Erc721gp *Erc721gpSession) SetBaseURI(_uri string) (*types.Transaction, error) {
	return _Erc721gp.Contract.SetBaseURI(&_Erc721gp.TransactOpts, _uri)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _uri) returns()
func (_Erc721gp *Erc721gpTransactorSession) SetBaseURI(_uri string) (*types.Transaction, error) {
	return _Erc721gp.Contract.SetBaseURI(&_Erc721gp.TransactOpts, _uri)
}

// SetLaunchpad is a paid mutator transaction binding the contract method 0x36f4c0eb.
//
// Solidity: function setLaunchpad(address launchpad_) returns()
func (_Erc721gp *Erc721gpTransactor) SetLaunchpad(opts *bind.TransactOpts, launchpad_ common.Address) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "setLaunchpad", launchpad_)
}

// SetLaunchpad is a paid mutator transaction binding the contract method 0x36f4c0eb.
//
// Solidity: function setLaunchpad(address launchpad_) returns()
func (_Erc721gp *Erc721gpSession) SetLaunchpad(launchpad_ common.Address) (*types.Transaction, error) {
	return _Erc721gp.Contract.SetLaunchpad(&_Erc721gp.TransactOpts, launchpad_)
}

// SetLaunchpad is a paid mutator transaction binding the contract method 0x36f4c0eb.
//
// Solidity: function setLaunchpad(address launchpad_) returns()
func (_Erc721gp *Erc721gpTransactorSession) SetLaunchpad(launchpad_ common.Address) (*types.Transaction, error) {
	return _Erc721gp.Contract.SetLaunchpad(&_Erc721gp.TransactOpts, launchpad_)
}

// SetNameSymbol is a paid mutator transaction binding the contract method 0x504334c2.
//
// Solidity: function setNameSymbol(string _name, string _symbol) returns()
func (_Erc721gp *Erc721gpTransactor) SetNameSymbol(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "setNameSymbol", _name, _symbol)
}

// SetNameSymbol is a paid mutator transaction binding the contract method 0x504334c2.
//
// Solidity: function setNameSymbol(string _name, string _symbol) returns()
func (_Erc721gp *Erc721gpSession) SetNameSymbol(_name string, _symbol string) (*types.Transaction, error) {
	return _Erc721gp.Contract.SetNameSymbol(&_Erc721gp.TransactOpts, _name, _symbol)
}

// SetNameSymbol is a paid mutator transaction binding the contract method 0x504334c2.
//
// Solidity: function setNameSymbol(string _name, string _symbol) returns()
func (_Erc721gp *Erc721gpTransactorSession) SetNameSymbol(_name string, _symbol string) (*types.Transaction, error) {
	return _Erc721gp.Contract.SetNameSymbol(&_Erc721gp.TransactOpts, _name, _symbol)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.TransferFrom(&_Erc721gp.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721gp *Erc721gpTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721gp.Contract.TransferFrom(&_Erc721gp.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc721gp *Erc721gpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc721gp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc721gp *Erc721gpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc721gp.Contract.TransferOwnership(&_Erc721gp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc721gp *Erc721gpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc721gp.Contract.TransferOwnership(&_Erc721gp.TransactOpts, newOwner)
}

// Erc721gpApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc721gp contract.
type Erc721gpApprovalIterator struct {
	Event *Erc721gpApproval // Event containing the contract specifics and raw log

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
func (it *Erc721gpApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721gpApproval)
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
		it.Event = new(Erc721gpApproval)
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
func (it *Erc721gpApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721gpApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721gpApproval represents a Approval event raised by the Erc721gp contract.
type Erc721gpApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721gp *Erc721gpFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Erc721gpApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721gp.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721gpApprovalIterator{contract: _Erc721gp.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721gp *Erc721gpFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc721gpApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721gp.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721gpApproval)
				if err := _Erc721gp.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721gp *Erc721gpFilterer) ParseApproval(log types.Log) (*Erc721gpApproval, error) {
	event := new(Erc721gpApproval)
	if err := _Erc721gp.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721gpApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc721gp contract.
type Erc721gpApprovalForAllIterator struct {
	Event *Erc721gpApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Erc721gpApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721gpApprovalForAll)
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
		it.Event = new(Erc721gpApprovalForAll)
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
func (it *Erc721gpApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721gpApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721gpApprovalForAll represents a ApprovalForAll event raised by the Erc721gp contract.
type Erc721gpApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721gp *Erc721gpFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Erc721gpApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc721gp.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Erc721gpApprovalForAllIterator{contract: _Erc721gp.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721gp *Erc721gpFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc721gpApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc721gp.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721gpApprovalForAll)
				if err := _Erc721gp.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721gp *Erc721gpFilterer) ParseApprovalForAll(log types.Log) (*Erc721gpApprovalForAll, error) {
	event := new(Erc721gpApprovalForAll)
	if err := _Erc721gp.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721gpInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Erc721gp contract.
type Erc721gpInitializedIterator struct {
	Event *Erc721gpInitialized // Event containing the contract specifics and raw log

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
func (it *Erc721gpInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721gpInitialized)
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
		it.Event = new(Erc721gpInitialized)
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
func (it *Erc721gpInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721gpInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721gpInitialized represents a Initialized event raised by the Erc721gp contract.
type Erc721gpInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Erc721gp *Erc721gpFilterer) FilterInitialized(opts *bind.FilterOpts) (*Erc721gpInitializedIterator, error) {

	logs, sub, err := _Erc721gp.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Erc721gpInitializedIterator{contract: _Erc721gp.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Erc721gp *Erc721gpFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Erc721gpInitialized) (event.Subscription, error) {

	logs, sub, err := _Erc721gp.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721gpInitialized)
				if err := _Erc721gp.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Erc721gp *Erc721gpFilterer) ParseInitialized(log types.Log) (*Erc721gpInitialized, error) {
	event := new(Erc721gpInitialized)
	if err := _Erc721gp.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721gpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc721gp contract.
type Erc721gpOwnershipTransferredIterator struct {
	Event *Erc721gpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Erc721gpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721gpOwnershipTransferred)
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
		it.Event = new(Erc721gpOwnershipTransferred)
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
func (it *Erc721gpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721gpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721gpOwnershipTransferred represents a OwnershipTransferred event raised by the Erc721gp contract.
type Erc721gpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc721gp *Erc721gpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc721gpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc721gp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc721gpOwnershipTransferredIterator{contract: _Erc721gp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc721gp *Erc721gpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc721gpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc721gp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721gpOwnershipTransferred)
				if err := _Erc721gp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Erc721gp *Erc721gpFilterer) ParseOwnershipTransferred(log types.Log) (*Erc721gpOwnershipTransferred, error) {
	event := new(Erc721gpOwnershipTransferred)
	if err := _Erc721gp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721gpTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc721gp contract.
type Erc721gpTransferIterator struct {
	Event *Erc721gpTransfer // Event containing the contract specifics and raw log

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
func (it *Erc721gpTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721gpTransfer)
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
		it.Event = new(Erc721gpTransfer)
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
func (it *Erc721gpTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721gpTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721gpTransfer represents a Transfer event raised by the Erc721gp contract.
type Erc721gpTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721gp *Erc721gpFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Erc721gpTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721gp.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721gpTransferIterator{contract: _Erc721gp.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721gp *Erc721gpFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc721gpTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721gp.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721gpTransfer)
				if err := _Erc721gp.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721gp *Erc721gpFilterer) ParseTransfer(log types.Log) (*Erc721gpTransfer, error) {
	event := new(Erc721gpTransfer)
	if err := _Erc721gp.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
