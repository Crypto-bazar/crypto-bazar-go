// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// DAONFTNFTProposal is an auto generated low-level Go binding around an user-defined struct.
type DAONFTNFTProposal struct {
	TokenURI string
	Proposer common.Address
	Votes    *big.Int
	Minted   bool
}

// MycontractMetaData contains all meta data concerning the Mycontract contract.
var MycontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NFTMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"NFTProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposeNFT\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"minted\",\"type\":\"bool\"}],\"internalType\":\"structDAONFT.NFTProposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftProposals\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"minted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"proposeNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"voteForNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MycontractABI is the input ABI used to generate the binding from.
// Deprecated: Use MycontractMetaData.ABI instead.
var MycontractABI = MycontractMetaData.ABI

// Mycontract is an auto generated Go binding around an Ethereum contract.
type Mycontract struct {
	MycontractCaller     // Read-only binding to the contract
	MycontractTransactor // Write-only binding to the contract
	MycontractFilterer   // Log filterer for contract events
}

// MycontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MycontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MycontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MycontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MycontractSession struct {
	Contract     *Mycontract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MycontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MycontractCallerSession struct {
	Contract *MycontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MycontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MycontractTransactorSession struct {
	Contract     *MycontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MycontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MycontractRaw struct {
	Contract *Mycontract // Generic contract binding to access the raw methods on
}

// MycontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MycontractCallerRaw struct {
	Contract *MycontractCaller // Generic read-only contract binding to access the raw methods on
}

// MycontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MycontractTransactorRaw struct {
	Contract *MycontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMycontract creates a new instance of Mycontract, bound to a specific deployed contract.
func NewMycontract(address common.Address, backend bind.ContractBackend) (*Mycontract, error) {
	contract, err := bindMycontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mycontract{MycontractCaller: MycontractCaller{contract: contract}, MycontractTransactor: MycontractTransactor{contract: contract}, MycontractFilterer: MycontractFilterer{contract: contract}}, nil
}

// NewMycontractCaller creates a new read-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractCaller(address common.Address, caller bind.ContractCaller) (*MycontractCaller, error) {
	contract, err := bindMycontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractCaller{contract: contract}, nil
}

// NewMycontractTransactor creates a new write-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractTransactor(address common.Address, transactor bind.ContractTransactor) (*MycontractTransactor, error) {
	contract, err := bindMycontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractTransactor{contract: contract}, nil
}

// NewMycontractFilterer creates a new log filterer instance of Mycontract, bound to a specific deployed contract.
func NewMycontractFilterer(address common.Address, filterer bind.ContractFilterer) (*MycontractFilterer, error) {
	contract, err := bindMycontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MycontractFilterer{contract: contract}, nil
}

// bindMycontract binds a generic wrapper to an already deployed contract.
func bindMycontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MycontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.MycontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mycontract *MycontractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mycontract *MycontractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Mycontract.Contract.BalanceOf(&_Mycontract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mycontract *MycontractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Mycontract.Contract.BalanceOf(&_Mycontract.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Mycontract.Contract.GetApproved(&_Mycontract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Mycontract.Contract.GetApproved(&_Mycontract.CallOpts, tokenId)
}

// GetProposeNFT is a free data retrieval call binding the contract method 0x555356ab.
//
// Solidity: function getProposeNFT() view returns((string,address,uint256,bool)[])
func (_Mycontract *MycontractCaller) GetProposeNFT(opts *bind.CallOpts) ([]DAONFTNFTProposal, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "getProposeNFT")

	if err != nil {
		return *new([]DAONFTNFTProposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]DAONFTNFTProposal)).(*[]DAONFTNFTProposal)

	return out0, err

}

// GetProposeNFT is a free data retrieval call binding the contract method 0x555356ab.
//
// Solidity: function getProposeNFT() view returns((string,address,uint256,bool)[])
func (_Mycontract *MycontractSession) GetProposeNFT() ([]DAONFTNFTProposal, error) {
	return _Mycontract.Contract.GetProposeNFT(&_Mycontract.CallOpts)
}

// GetProposeNFT is a free data retrieval call binding the contract method 0x555356ab.
//
// Solidity: function getProposeNFT() view returns((string,address,uint256,bool)[])
func (_Mycontract *MycontractCallerSession) GetProposeNFT() ([]DAONFTNFTProposal, error) {
	return _Mycontract.Contract.GetProposeNFT(&_Mycontract.CallOpts)
}

// GovernanceToken is a free data retrieval call binding the contract method 0xf96dae0a.
//
// Solidity: function governanceToken() view returns(address)
func (_Mycontract *MycontractCaller) GovernanceToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "governanceToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceToken is a free data retrieval call binding the contract method 0xf96dae0a.
//
// Solidity: function governanceToken() view returns(address)
func (_Mycontract *MycontractSession) GovernanceToken() (common.Address, error) {
	return _Mycontract.Contract.GovernanceToken(&_Mycontract.CallOpts)
}

// GovernanceToken is a free data retrieval call binding the contract method 0xf96dae0a.
//
// Solidity: function governanceToken() view returns(address)
func (_Mycontract *MycontractCallerSession) GovernanceToken() (common.Address, error) {
	return _Mycontract.Contract.GovernanceToken(&_Mycontract.CallOpts)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 , address ) view returns(bool)
func (_Mycontract *MycontractCaller) HasVoted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "hasVoted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 , address ) view returns(bool)
func (_Mycontract *MycontractSession) HasVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Mycontract.Contract.HasVoted(&_Mycontract.CallOpts, arg0, arg1)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 , address ) view returns(bool)
func (_Mycontract *MycontractCallerSession) HasVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Mycontract.Contract.HasVoted(&_Mycontract.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mycontract *MycontractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mycontract *MycontractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Mycontract.Contract.IsApprovedForAll(&_Mycontract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mycontract *MycontractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Mycontract.Contract.IsApprovedForAll(&_Mycontract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycontract *MycontractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycontract *MycontractSession) Name() (string, error) {
	return _Mycontract.Contract.Name(&_Mycontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mycontract *MycontractCallerSession) Name() (string, error) {
	return _Mycontract.Contract.Name(&_Mycontract.CallOpts)
}

// NftProposals is a free data retrieval call binding the contract method 0x54ab9588.
//
// Solidity: function nftProposals(uint256 ) view returns(string tokenURI, address proposer, uint256 votes, bool minted)
func (_Mycontract *MycontractCaller) NftProposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenURI string
	Proposer common.Address
	Votes    *big.Int
	Minted   bool
}, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "nftProposals", arg0)

	outstruct := new(struct {
		TokenURI string
		Proposer common.Address
		Votes    *big.Int
		Minted   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenURI = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Proposer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Votes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Minted = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// NftProposals is a free data retrieval call binding the contract method 0x54ab9588.
//
// Solidity: function nftProposals(uint256 ) view returns(string tokenURI, address proposer, uint256 votes, bool minted)
func (_Mycontract *MycontractSession) NftProposals(arg0 *big.Int) (struct {
	TokenURI string
	Proposer common.Address
	Votes    *big.Int
	Minted   bool
}, error) {
	return _Mycontract.Contract.NftProposals(&_Mycontract.CallOpts, arg0)
}

// NftProposals is a free data retrieval call binding the contract method 0x54ab9588.
//
// Solidity: function nftProposals(uint256 ) view returns(string tokenURI, address proposer, uint256 votes, bool minted)
func (_Mycontract *MycontractCallerSession) NftProposals(arg0 *big.Int) (struct {
	TokenURI string
	Proposer common.Address
	Votes    *big.Int
	Minted   bool
}, error) {
	return _Mycontract.Contract.NftProposals(&_Mycontract.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Mycontract.Contract.OwnerOf(&_Mycontract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mycontract *MycontractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Mycontract.Contract.OwnerOf(&_Mycontract.CallOpts, tokenId)
}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_Mycontract *MycontractCaller) ProposalCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "proposalCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_Mycontract *MycontractSession) ProposalCounter() (*big.Int, error) {
	return _Mycontract.Contract.ProposalCounter(&_Mycontract.CallOpts)
}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_Mycontract *MycontractCallerSession) ProposalCounter() (*big.Int, error) {
	return _Mycontract.Contract.ProposalCounter(&_Mycontract.CallOpts)
}

// RequiredVotes is a free data retrieval call binding the contract method 0xbd31a4d8.
//
// Solidity: function requiredVotes() view returns(uint256)
func (_Mycontract *MycontractCaller) RequiredVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "requiredVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredVotes is a free data retrieval call binding the contract method 0xbd31a4d8.
//
// Solidity: function requiredVotes() view returns(uint256)
func (_Mycontract *MycontractSession) RequiredVotes() (*big.Int, error) {
	return _Mycontract.Contract.RequiredVotes(&_Mycontract.CallOpts)
}

// RequiredVotes is a free data retrieval call binding the contract method 0xbd31a4d8.
//
// Solidity: function requiredVotes() view returns(uint256)
func (_Mycontract *MycontractCallerSession) RequiredVotes() (*big.Int, error) {
	return _Mycontract.Contract.RequiredVotes(&_Mycontract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mycontract *MycontractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mycontract *MycontractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mycontract.Contract.SupportsInterface(&_Mycontract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mycontract *MycontractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mycontract.Contract.SupportsInterface(&_Mycontract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycontract *MycontractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycontract *MycontractSession) Symbol() (string, error) {
	return _Mycontract.Contract.Symbol(&_Mycontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mycontract *MycontractCallerSession) Symbol() (string, error) {
	return _Mycontract.Contract.Symbol(&_Mycontract.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mycontract *MycontractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Mycontract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mycontract *MycontractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Mycontract.Contract.TokenURI(&_Mycontract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mycontract *MycontractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Mycontract.Contract.TokenURI(&_Mycontract.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mycontract *MycontractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Approve(&_Mycontract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.Approve(&_Mycontract.TransactOpts, to, tokenId)
}

// ProposeNFT is a paid mutator transaction binding the contract method 0x69aecfd7.
//
// Solidity: function proposeNFT(string _tokenURI) returns()
func (_Mycontract *MycontractTransactor) ProposeNFT(opts *bind.TransactOpts, _tokenURI string) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "proposeNFT", _tokenURI)
}

// ProposeNFT is a paid mutator transaction binding the contract method 0x69aecfd7.
//
// Solidity: function proposeNFT(string _tokenURI) returns()
func (_Mycontract *MycontractSession) ProposeNFT(_tokenURI string) (*types.Transaction, error) {
	return _Mycontract.Contract.ProposeNFT(&_Mycontract.TransactOpts, _tokenURI)
}

// ProposeNFT is a paid mutator transaction binding the contract method 0x69aecfd7.
//
// Solidity: function proposeNFT(string _tokenURI) returns()
func (_Mycontract *MycontractTransactorSession) ProposeNFT(_tokenURI string) (*types.Transaction, error) {
	return _Mycontract.Contract.ProposeNFT(&_Mycontract.TransactOpts, _tokenURI)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SafeTransferFrom(&_Mycontract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.SafeTransferFrom(&_Mycontract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Mycontract *MycontractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Mycontract *MycontractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Mycontract.Contract.SafeTransferFrom0(&_Mycontract.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Mycontract *MycontractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Mycontract.Contract.SafeTransferFrom0(&_Mycontract.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mycontract *MycontractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mycontract *MycontractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mycontract.Contract.SetApprovalForAll(&_Mycontract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mycontract *MycontractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mycontract.Contract.SetApprovalForAll(&_Mycontract.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.TransferFrom(&_Mycontract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mycontract *MycontractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.TransferFrom(&_Mycontract.TransactOpts, from, to, tokenId)
}

// VoteForNFT is a paid mutator transaction binding the contract method 0x89ba52c4.
//
// Solidity: function voteForNFT(uint256 proposalId) returns()
func (_Mycontract *MycontractTransactor) VoteForNFT(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "voteForNFT", proposalId)
}

// VoteForNFT is a paid mutator transaction binding the contract method 0x89ba52c4.
//
// Solidity: function voteForNFT(uint256 proposalId) returns()
func (_Mycontract *MycontractSession) VoteForNFT(proposalId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.VoteForNFT(&_Mycontract.TransactOpts, proposalId)
}

// VoteForNFT is a paid mutator transaction binding the contract method 0x89ba52c4.
//
// Solidity: function voteForNFT(uint256 proposalId) returns()
func (_Mycontract *MycontractTransactorSession) VoteForNFT(proposalId *big.Int) (*types.Transaction, error) {
	return _Mycontract.Contract.VoteForNFT(&_Mycontract.TransactOpts, proposalId)
}

// MycontractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mycontract contract.
type MycontractApprovalIterator struct {
	Event *MycontractApproval // Event containing the contract specifics and raw log

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
func (it *MycontractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractApproval)
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
		it.Event = new(MycontractApproval)
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
func (it *MycontractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractApproval represents a Approval event raised by the Mycontract contract.
type MycontractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mycontract *MycontractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MycontractApprovalIterator, error) {

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

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MycontractApprovalIterator{contract: _Mycontract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mycontract *MycontractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MycontractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractApproval)
				if err := _Mycontract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Mycontract *MycontractFilterer) ParseApproval(log types.Log) (*MycontractApproval, error) {
	event := new(MycontractApproval)
	if err := _Mycontract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Mycontract contract.
type MycontractApprovalForAllIterator struct {
	Event *MycontractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MycontractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractApprovalForAll)
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
		it.Event = new(MycontractApprovalForAll)
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
func (it *MycontractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractApprovalForAll represents a ApprovalForAll event raised by the Mycontract contract.
type MycontractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mycontract *MycontractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MycontractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MycontractApprovalForAllIterator{contract: _Mycontract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mycontract *MycontractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MycontractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractApprovalForAll)
				if err := _Mycontract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Mycontract *MycontractFilterer) ParseApprovalForAll(log types.Log) (*MycontractApprovalForAll, error) {
	event := new(MycontractApprovalForAll)
	if err := _Mycontract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractNFTMintedIterator is returned from FilterNFTMinted and is used to iterate over the raw logs and unpacked data for NFTMinted events raised by the Mycontract contract.
type MycontractNFTMintedIterator struct {
	Event *MycontractNFTMinted // Event containing the contract specifics and raw log

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
func (it *MycontractNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractNFTMinted)
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
		it.Event = new(MycontractNFTMinted)
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
func (it *MycontractNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractNFTMinted represents a NFTMinted event raised by the Mycontract contract.
type MycontractNFTMinted struct {
	TokenId  *big.Int
	TokenURI string
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNFTMinted is a free log retrieval operation binding the contract event 0x411b345259a1495602cee098e5f501eb3412e0701c10b173ecd8ba7b585974a9.
//
// Solidity: event NFTMinted(uint256 indexed tokenId, string tokenURI, address owner)
func (_Mycontract *MycontractFilterer) FilterNFTMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*MycontractNFTMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "NFTMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MycontractNFTMintedIterator{contract: _Mycontract.contract, event: "NFTMinted", logs: logs, sub: sub}, nil
}

// WatchNFTMinted is a free log subscription operation binding the contract event 0x411b345259a1495602cee098e5f501eb3412e0701c10b173ecd8ba7b585974a9.
//
// Solidity: event NFTMinted(uint256 indexed tokenId, string tokenURI, address owner)
func (_Mycontract *MycontractFilterer) WatchNFTMinted(opts *bind.WatchOpts, sink chan<- *MycontractNFTMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "NFTMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractNFTMinted)
				if err := _Mycontract.contract.UnpackLog(event, "NFTMinted", log); err != nil {
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

// ParseNFTMinted is a log parse operation binding the contract event 0x411b345259a1495602cee098e5f501eb3412e0701c10b173ecd8ba7b585974a9.
//
// Solidity: event NFTMinted(uint256 indexed tokenId, string tokenURI, address owner)
func (_Mycontract *MycontractFilterer) ParseNFTMinted(log types.Log) (*MycontractNFTMinted, error) {
	event := new(MycontractNFTMinted)
	if err := _Mycontract.contract.UnpackLog(event, "NFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractNFTProposedIterator is returned from FilterNFTProposed and is used to iterate over the raw logs and unpacked data for NFTProposed events raised by the Mycontract contract.
type MycontractNFTProposedIterator struct {
	Event *MycontractNFTProposed // Event containing the contract specifics and raw log

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
func (it *MycontractNFTProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractNFTProposed)
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
		it.Event = new(MycontractNFTProposed)
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
func (it *MycontractNFTProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractNFTProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractNFTProposed represents a NFTProposed event raised by the Mycontract contract.
type MycontractNFTProposed struct {
	ProposalId *big.Int
	TokenURI   string
	Proposer   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTProposed is a free log retrieval operation binding the contract event 0x66fde1e5134898d49de0f081d598cf7d8fba2432595667ac3cbef6dd60dc0115.
//
// Solidity: event NFTProposed(uint256 indexed proposalId, string tokenURI, address proposer)
func (_Mycontract *MycontractFilterer) FilterNFTProposed(opts *bind.FilterOpts, proposalId []*big.Int) (*MycontractNFTProposedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "NFTProposed", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &MycontractNFTProposedIterator{contract: _Mycontract.contract, event: "NFTProposed", logs: logs, sub: sub}, nil
}

// WatchNFTProposed is a free log subscription operation binding the contract event 0x66fde1e5134898d49de0f081d598cf7d8fba2432595667ac3cbef6dd60dc0115.
//
// Solidity: event NFTProposed(uint256 indexed proposalId, string tokenURI, address proposer)
func (_Mycontract *MycontractFilterer) WatchNFTProposed(opts *bind.WatchOpts, sink chan<- *MycontractNFTProposed, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "NFTProposed", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractNFTProposed)
				if err := _Mycontract.contract.UnpackLog(event, "NFTProposed", log); err != nil {
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

// ParseNFTProposed is a log parse operation binding the contract event 0x66fde1e5134898d49de0f081d598cf7d8fba2432595667ac3cbef6dd60dc0115.
//
// Solidity: event NFTProposed(uint256 indexed proposalId, string tokenURI, address proposer)
func (_Mycontract *MycontractFilterer) ParseNFTProposed(log types.Log) (*MycontractNFTProposed, error) {
	event := new(MycontractNFTProposed)
	if err := _Mycontract.contract.UnpackLog(event, "NFTProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mycontract contract.
type MycontractTransferIterator struct {
	Event *MycontractTransfer // Event containing the contract specifics and raw log

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
func (it *MycontractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractTransfer)
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
		it.Event = new(MycontractTransfer)
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
func (it *MycontractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractTransfer represents a Transfer event raised by the Mycontract contract.
type MycontractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mycontract *MycontractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MycontractTransferIterator, error) {

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

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MycontractTransferIterator{contract: _Mycontract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mycontract *MycontractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MycontractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractTransfer)
				if err := _Mycontract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Mycontract *MycontractFilterer) ParseTransfer(log types.Log) (*MycontractTransfer, error) {
	event := new(MycontractTransfer)
	if err := _Mycontract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MycontractVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Mycontract contract.
type MycontractVotedIterator struct {
	Event *MycontractVoted // Event containing the contract specifics and raw log

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
func (it *MycontractVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractVoted)
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
		it.Event = new(MycontractVoted)
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
func (it *MycontractVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractVoted represents a Voted event raised by the Mycontract contract.
type MycontractVoted struct {
	ProposalId *big.Int
	Voter      common.Address
	TokenURI   string
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x3ae89ec44d6750209e1164a2659a766e3cfe3476e908a19cb5e928974ef5e4ec.
//
// Solidity: event Voted(uint256 indexed proposalId, address indexed voter, string tokenURI, uint256 amount)
func (_Mycontract *MycontractFilterer) FilterVoted(opts *bind.FilterOpts, proposalId []*big.Int, voter []common.Address) (*MycontractVotedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "Voted", proposalIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &MycontractVotedIterator{contract: _Mycontract.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x3ae89ec44d6750209e1164a2659a766e3cfe3476e908a19cb5e928974ef5e4ec.
//
// Solidity: event Voted(uint256 indexed proposalId, address indexed voter, string tokenURI, uint256 amount)
func (_Mycontract *MycontractFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *MycontractVoted, proposalId []*big.Int, voter []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "Voted", proposalIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractVoted)
				if err := _Mycontract.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x3ae89ec44d6750209e1164a2659a766e3cfe3476e908a19cb5e928974ef5e4ec.
//
// Solidity: event Voted(uint256 indexed proposalId, address indexed voter, string tokenURI, uint256 amount)
func (_Mycontract *MycontractFilterer) ParseVoted(log types.Log) (*MycontractVoted, error) {
	event := new(MycontractVoted)
	if err := _Mycontract.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
