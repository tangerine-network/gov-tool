// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/tangerine-network/go-tangerine"
	"github.com/tangerine-network/go-tangerine/accounts/abi"
	"github.com/tangerine-network/go-tangerine/accounts/abi/bind"
	"github.com/tangerine-network/go-tangerine/common"
	"github.com/tangerine-network/go-tangerine/core/types"
	"github.com/tangerine-network/go-tangerine/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GovernanceABI is the input ABI used to generate the binding from.
const GovernanceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"dkgSuccesses\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notarySetSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"staked\",\"type\":\"uint256\"},{\"name\":\"fined\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"location\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"},{\"name\":\"unstaked\",\"type\":\"uint256\"},{\"name\":\"unstakedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notaryParamBeta\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"miningVelocity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lambdaBA\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"crsRound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notaryParamAlpha\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dkgSuccessesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"dkgFinalizeds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockGasLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dkgRound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodesOffsetByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"crs\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roundLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextHalvingSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dkgComplaints\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dkgMasterPublicKeyOffset\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"dkgMPKReadys\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastHalvedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"finedRecords\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lambdaDKG\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fineValues\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dkgMPKReadysCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBlockInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dkgMasterPublicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastProposedHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minGasPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dkgFinalizedsCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dkgComplaintsProposed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodesOffsetByNodeKeyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockupPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dkgResetCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ConfigurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"Round\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"CRS\",\"type\":\"bytes32\"}],\"name\":\"CRSProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"NewOwnerAddress\",\"type\":\"address\"}],\"name\":\"NodeOwnershipTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"PublicKey\",\"type\":\"bytes\"}],\"name\":\"NodePublicKeyReplaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"}],\"name\":\"NodeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Type\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"Arg1\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"Arg2\",\"type\":\"bytes\"}],\"name\":\"Reported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"Fined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"FinePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"Round\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"BlockHeight\",\"type\":\"uint256\"}],\"name\":\"DKGReset\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"NewOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"MinStake\",\"type\":\"uint256\"},{\"name\":\"LockupPeriod\",\"type\":\"uint256\"},{\"name\":\"MinGasPrice\",\"type\":\"uint256\"},{\"name\":\"BlockGasLimit\",\"type\":\"uint256\"},{\"name\":\"LambdaBA\",\"type\":\"uint256\"},{\"name\":\"LambdaDKG\",\"type\":\"uint256\"},{\"name\":\"NotaryParamAlpha\",\"type\":\"uint256\"},{\"name\":\"NotaryParamBeta\",\"type\":\"uint256\"},{\"name\":\"RoundLength\",\"type\":\"uint256\"},{\"name\":\"MinBlockInterval\",\"type\":\"uint256\"},{\"name\":\"FineValues\",\"type\":\"uint256[]\"}],\"name\":\"updateConfiguration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NewOwner\",\"type\":\"address\"}],\"name\":\"transferNodeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"OldOwner\",\"type\":\"address\"},{\"name\":\"NewOwner\",\"type\":\"address\"}],\"name\":\"transferNodeOwnershipByFoundation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NewPublicKey\",\"type\":\"bytes\"}],\"name\":\"replaceNodePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nodesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Round\",\"type\":\"uint256\"},{\"name\":\"SignedCRS\",\"type\":\"bytes\"}],\"name\":\"proposeCRS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Complaint\",\"type\":\"bytes\"}],\"name\":\"addDKGComplaint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"PublicKey\",\"type\":\"bytes\"}],\"name\":\"addDKGMasterPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"MPKReady\",\"type\":\"bytes\"}],\"name\":\"addDKGMPKReady\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Finalize\",\"type\":\"bytes\"}],\"name\":\"addDKGFinalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Success\",\"type\":\"bytes\"}],\"name\":\"addDKGSuccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"PublicKey\",\"type\":\"bytes\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"Location\",\"type\":\"string\"},{\"name\":\"Url\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NodeAddress\",\"type\":\"address\"}],\"name\":\"payFine\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Type\",\"type\":\"uint256\"},{\"name\":\"Arg1\",\"type\":\"bytes\"},{\"name\":\"Arg2\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NewSignedCRS\",\"type\":\"bytes\"}],\"name\":\"resetDKG\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() constant returns(uint256)
func (_Governance *GovernanceCaller) BlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "blockGasLimit")
	return *ret0, err
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() constant returns(uint256)
func (_Governance *GovernanceSession) BlockGasLimit() (*big.Int, error) {
	return _Governance.Contract.BlockGasLimit(&_Governance.CallOpts)
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() constant returns(uint256)
func (_Governance *GovernanceCallerSession) BlockGasLimit() (*big.Int, error) {
	return _Governance.Contract.BlockGasLimit(&_Governance.CallOpts)
}

// Crs is a free data retrieval call binding the contract method 0x8a591369.
//
// Solidity: function crs() constant returns(bytes32)
func (_Governance *GovernanceCaller) Crs(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "crs")
	return *ret0, err
}

// Crs is a free data retrieval call binding the contract method 0x8a591369.
//
// Solidity: function crs() constant returns(bytes32)
func (_Governance *GovernanceSession) Crs() ([32]byte, error) {
	return _Governance.Contract.Crs(&_Governance.CallOpts)
}

// Crs is a free data retrieval call binding the contract method 0x8a591369.
//
// Solidity: function crs() constant returns(bytes32)
func (_Governance *GovernanceCallerSession) Crs() ([32]byte, error) {
	return _Governance.Contract.Crs(&_Governance.CallOpts)
}

// CrsRound is a free data retrieval call binding the contract method 0x47731325.
//
// Solidity: function crsRound() constant returns(uint256)
func (_Governance *GovernanceCaller) CrsRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "crsRound")
	return *ret0, err
}

// CrsRound is a free data retrieval call binding the contract method 0x47731325.
//
// Solidity: function crsRound() constant returns(uint256)
func (_Governance *GovernanceSession) CrsRound() (*big.Int, error) {
	return _Governance.Contract.CrsRound(&_Governance.CallOpts)
}

// CrsRound is a free data retrieval call binding the contract method 0x47731325.
//
// Solidity: function crsRound() constant returns(uint256)
func (_Governance *GovernanceCallerSession) CrsRound() (*big.Int, error) {
	return _Governance.Contract.CrsRound(&_Governance.CallOpts)
}

// DkgComplaints is a free data retrieval call binding the contract method 0x8d90a159.
//
// Solidity: function dkgComplaints(uint256 ) constant returns(bytes)
func (_Governance *GovernanceCaller) DkgComplaints(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgComplaints", arg0)
	return *ret0, err
}

// DkgComplaints is a free data retrieval call binding the contract method 0x8d90a159.
//
// Solidity: function dkgComplaints(uint256 ) constant returns(bytes)
func (_Governance *GovernanceSession) DkgComplaints(arg0 *big.Int) ([]byte, error) {
	return _Governance.Contract.DkgComplaints(&_Governance.CallOpts, arg0)
}

// DkgComplaints is a free data retrieval call binding the contract method 0x8d90a159.
//
// Solidity: function dkgComplaints(uint256 ) constant returns(bytes)
func (_Governance *GovernanceCallerSession) DkgComplaints(arg0 *big.Int) ([]byte, error) {
	return _Governance.Contract.DkgComplaints(&_Governance.CallOpts, arg0)
}

// DkgComplaintsProposed is a free data retrieval call binding the contract method 0xe746714d.
//
// Solidity: function dkgComplaintsProposed(bytes32 ) constant returns(bool)
func (_Governance *GovernanceCaller) DkgComplaintsProposed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgComplaintsProposed", arg0)
	return *ret0, err
}

// DkgComplaintsProposed is a free data retrieval call binding the contract method 0xe746714d.
//
// Solidity: function dkgComplaintsProposed(bytes32 ) constant returns(bool)
func (_Governance *GovernanceSession) DkgComplaintsProposed(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.DkgComplaintsProposed(&_Governance.CallOpts, arg0)
}

// DkgComplaintsProposed is a free data retrieval call binding the contract method 0xe746714d.
//
// Solidity: function dkgComplaintsProposed(bytes32 ) constant returns(bool)
func (_Governance *GovernanceCallerSession) DkgComplaintsProposed(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.DkgComplaintsProposed(&_Governance.CallOpts, arg0)
}

// DkgFinalizeds is a free data retrieval call binding the contract method 0x70f3ac54.
//
// Solidity: function dkgFinalizeds(address ) constant returns(bool)
func (_Governance *GovernanceCaller) DkgFinalizeds(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgFinalizeds", arg0)
	return *ret0, err
}

// DkgFinalizeds is a free data retrieval call binding the contract method 0x70f3ac54.
//
// Solidity: function dkgFinalizeds(address ) constant returns(bool)
func (_Governance *GovernanceSession) DkgFinalizeds(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgFinalizeds(&_Governance.CallOpts, arg0)
}

// DkgFinalizeds is a free data retrieval call binding the contract method 0x70f3ac54.
//
// Solidity: function dkgFinalizeds(address ) constant returns(bool)
func (_Governance *GovernanceCallerSession) DkgFinalizeds(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgFinalizeds(&_Governance.CallOpts, arg0)
}

// DkgFinalizedsCount is a free data retrieval call binding the contract method 0xdd2083d0.
//
// Solidity: function dkgFinalizedsCount() constant returns(uint256)
func (_Governance *GovernanceCaller) DkgFinalizedsCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgFinalizedsCount")
	return *ret0, err
}

// DkgFinalizedsCount is a free data retrieval call binding the contract method 0xdd2083d0.
//
// Solidity: function dkgFinalizedsCount() constant returns(uint256)
func (_Governance *GovernanceSession) DkgFinalizedsCount() (*big.Int, error) {
	return _Governance.Contract.DkgFinalizedsCount(&_Governance.CallOpts)
}

// DkgFinalizedsCount is a free data retrieval call binding the contract method 0xdd2083d0.
//
// Solidity: function dkgFinalizedsCount() constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgFinalizedsCount() (*big.Int, error) {
	return _Governance.Contract.DkgFinalizedsCount(&_Governance.CallOpts)
}

// DkgMPKReadys is a free data retrieval call binding the contract method 0x9bc9f489.
//
// Solidity: function dkgMPKReadys(address ) constant returns(bool)
func (_Governance *GovernanceCaller) DkgMPKReadys(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgMPKReadys", arg0)
	return *ret0, err
}

// DkgMPKReadys is a free data retrieval call binding the contract method 0x9bc9f489.
//
// Solidity: function dkgMPKReadys(address ) constant returns(bool)
func (_Governance *GovernanceSession) DkgMPKReadys(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgMPKReadys(&_Governance.CallOpts, arg0)
}

// DkgMPKReadys is a free data retrieval call binding the contract method 0x9bc9f489.
//
// Solidity: function dkgMPKReadys(address ) constant returns(bool)
func (_Governance *GovernanceCallerSession) DkgMPKReadys(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgMPKReadys(&_Governance.CallOpts, arg0)
}

// DkgMPKReadysCount is a free data retrieval call binding the contract method 0xb03e75e4.
//
// Solidity: function dkgMPKReadysCount() constant returns(uint256)
func (_Governance *GovernanceCaller) DkgMPKReadysCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgMPKReadysCount")
	return *ret0, err
}

// DkgMPKReadysCount is a free data retrieval call binding the contract method 0xb03e75e4.
//
// Solidity: function dkgMPKReadysCount() constant returns(uint256)
func (_Governance *GovernanceSession) DkgMPKReadysCount() (*big.Int, error) {
	return _Governance.Contract.DkgMPKReadysCount(&_Governance.CallOpts)
}

// DkgMPKReadysCount is a free data retrieval call binding the contract method 0xb03e75e4.
//
// Solidity: function dkgMPKReadysCount() constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgMPKReadysCount() (*big.Int, error) {
	return _Governance.Contract.DkgMPKReadysCount(&_Governance.CallOpts)
}

// DkgMasterPublicKeyOffset is a free data retrieval call binding the contract method 0x99aadb9a.
//
// Solidity: function dkgMasterPublicKeyOffset(bytes32 ) constant returns(uint256)
func (_Governance *GovernanceCaller) DkgMasterPublicKeyOffset(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgMasterPublicKeyOffset", arg0)
	return *ret0, err
}

// DkgMasterPublicKeyOffset is a free data retrieval call binding the contract method 0x99aadb9a.
//
// Solidity: function dkgMasterPublicKeyOffset(bytes32 ) constant returns(uint256)
func (_Governance *GovernanceSession) DkgMasterPublicKeyOffset(arg0 [32]byte) (*big.Int, error) {
	return _Governance.Contract.DkgMasterPublicKeyOffset(&_Governance.CallOpts, arg0)
}

// DkgMasterPublicKeyOffset is a free data retrieval call binding the contract method 0x99aadb9a.
//
// Solidity: function dkgMasterPublicKeyOffset(bytes32 ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgMasterPublicKeyOffset(arg0 [32]byte) (*big.Int, error) {
	return _Governance.Contract.DkgMasterPublicKeyOffset(&_Governance.CallOpts, arg0)
}

// DkgMasterPublicKeys is a free data retrieval call binding the contract method 0xd21abb78.
//
// Solidity: function dkgMasterPublicKeys(uint256 ) constant returns(bytes)
func (_Governance *GovernanceCaller) DkgMasterPublicKeys(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgMasterPublicKeys", arg0)
	return *ret0, err
}

// DkgMasterPublicKeys is a free data retrieval call binding the contract method 0xd21abb78.
//
// Solidity: function dkgMasterPublicKeys(uint256 ) constant returns(bytes)
func (_Governance *GovernanceSession) DkgMasterPublicKeys(arg0 *big.Int) ([]byte, error) {
	return _Governance.Contract.DkgMasterPublicKeys(&_Governance.CallOpts, arg0)
}

// DkgMasterPublicKeys is a free data retrieval call binding the contract method 0xd21abb78.
//
// Solidity: function dkgMasterPublicKeys(uint256 ) constant returns(bytes)
func (_Governance *GovernanceCallerSession) DkgMasterPublicKeys(arg0 *big.Int) ([]byte, error) {
	return _Governance.Contract.DkgMasterPublicKeys(&_Governance.CallOpts, arg0)
}

// DkgResetCount is a free data retrieval call binding the contract method 0xf0771659.
//
// Solidity: function dkgResetCount(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCaller) DkgResetCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgResetCount", arg0)
	return *ret0, err
}

// DkgResetCount is a free data retrieval call binding the contract method 0xf0771659.
//
// Solidity: function dkgResetCount(uint256 ) constant returns(uint256)
func (_Governance *GovernanceSession) DkgResetCount(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.DkgResetCount(&_Governance.CallOpts, arg0)
}

// DkgResetCount is a free data retrieval call binding the contract method 0xf0771659.
//
// Solidity: function dkgResetCount(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgResetCount(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.DkgResetCount(&_Governance.CallOpts, arg0)
}

// DkgRound is a free data retrieval call binding the contract method 0x7a8c1c1e.
//
// Solidity: function dkgRound() constant returns(uint256)
func (_Governance *GovernanceCaller) DkgRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgRound")
	return *ret0, err
}

// DkgRound is a free data retrieval call binding the contract method 0x7a8c1c1e.
//
// Solidity: function dkgRound() constant returns(uint256)
func (_Governance *GovernanceSession) DkgRound() (*big.Int, error) {
	return _Governance.Contract.DkgRound(&_Governance.CallOpts)
}

// DkgRound is a free data retrieval call binding the contract method 0x7a8c1c1e.
//
// Solidity: function dkgRound() constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgRound() (*big.Int, error) {
	return _Governance.Contract.DkgRound(&_Governance.CallOpts)
}

// DkgSuccesses is a free data retrieval call binding the contract method 0x072bf11a.
//
// Solidity: function dkgSuccesses(address ) constant returns(bool)
func (_Governance *GovernanceCaller) DkgSuccesses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgSuccesses", arg0)
	return *ret0, err
}

// DkgSuccesses is a free data retrieval call binding the contract method 0x072bf11a.
//
// Solidity: function dkgSuccesses(address ) constant returns(bool)
func (_Governance *GovernanceSession) DkgSuccesses(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgSuccesses(&_Governance.CallOpts, arg0)
}

// DkgSuccesses is a free data retrieval call binding the contract method 0x072bf11a.
//
// Solidity: function dkgSuccesses(address ) constant returns(bool)
func (_Governance *GovernanceCallerSession) DkgSuccesses(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgSuccesses(&_Governance.CallOpts, arg0)
}

// DkgSuccessesCount is a free data retrieval call binding the contract method 0x4cb38c7a.
//
// Solidity: function dkgSuccessesCount() constant returns(uint256)
func (_Governance *GovernanceCaller) DkgSuccessesCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgSuccessesCount")
	return *ret0, err
}

// DkgSuccessesCount is a free data retrieval call binding the contract method 0x4cb38c7a.
//
// Solidity: function dkgSuccessesCount() constant returns(uint256)
func (_Governance *GovernanceSession) DkgSuccessesCount() (*big.Int, error) {
	return _Governance.Contract.DkgSuccessesCount(&_Governance.CallOpts)
}

// DkgSuccessesCount is a free data retrieval call binding the contract method 0x4cb38c7a.
//
// Solidity: function dkgSuccessesCount() constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgSuccessesCount() (*big.Int, error) {
	return _Governance.Contract.DkgSuccessesCount(&_Governance.CallOpts)
}

// FineValues is a free data retrieval call binding the contract method 0xae1f289d.
//
// Solidity: function fineValues(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCaller) FineValues(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "fineValues", arg0)
	return *ret0, err
}

// FineValues is a free data retrieval call binding the contract method 0xae1f289d.
//
// Solidity: function fineValues(uint256 ) constant returns(uint256)
func (_Governance *GovernanceSession) FineValues(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.FineValues(&_Governance.CallOpts, arg0)
}

// FineValues is a free data retrieval call binding the contract method 0xae1f289d.
//
// Solidity: function fineValues(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) FineValues(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.FineValues(&_Governance.CallOpts, arg0)
}

// FinedRecords is a free data retrieval call binding the contract method 0xa1e460eb.
//
// Solidity: function finedRecords(bytes32 ) constant returns(bool)
func (_Governance *GovernanceCaller) FinedRecords(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "finedRecords", arg0)
	return *ret0, err
}

// FinedRecords is a free data retrieval call binding the contract method 0xa1e460eb.
//
// Solidity: function finedRecords(bytes32 ) constant returns(bool)
func (_Governance *GovernanceSession) FinedRecords(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.FinedRecords(&_Governance.CallOpts, arg0)
}

// FinedRecords is a free data retrieval call binding the contract method 0xa1e460eb.
//
// Solidity: function finedRecords(bytes32 ) constant returns(bool)
func (_Governance *GovernanceCallerSession) FinedRecords(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.FinedRecords(&_Governance.CallOpts, arg0)
}

// LambdaBA is a free data retrieval call binding the contract method 0x2deb3316.
//
// Solidity: function lambdaBA() constant returns(uint256)
func (_Governance *GovernanceCaller) LambdaBA(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lambdaBA")
	return *ret0, err
}

// LambdaBA is a free data retrieval call binding the contract method 0x2deb3316.
//
// Solidity: function lambdaBA() constant returns(uint256)
func (_Governance *GovernanceSession) LambdaBA() (*big.Int, error) {
	return _Governance.Contract.LambdaBA(&_Governance.CallOpts)
}

// LambdaBA is a free data retrieval call binding the contract method 0x2deb3316.
//
// Solidity: function lambdaBA() constant returns(uint256)
func (_Governance *GovernanceCallerSession) LambdaBA() (*big.Int, error) {
	return _Governance.Contract.LambdaBA(&_Governance.CallOpts)
}

// LambdaDKG is a free data retrieval call binding the contract method 0xa9601a8d.
//
// Solidity: function lambdaDKG() constant returns(uint256)
func (_Governance *GovernanceCaller) LambdaDKG(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lambdaDKG")
	return *ret0, err
}

// LambdaDKG is a free data retrieval call binding the contract method 0xa9601a8d.
//
// Solidity: function lambdaDKG() constant returns(uint256)
func (_Governance *GovernanceSession) LambdaDKG() (*big.Int, error) {
	return _Governance.Contract.LambdaDKG(&_Governance.CallOpts)
}

// LambdaDKG is a free data retrieval call binding the contract method 0xa9601a8d.
//
// Solidity: function lambdaDKG() constant returns(uint256)
func (_Governance *GovernanceCallerSession) LambdaDKG() (*big.Int, error) {
	return _Governance.Contract.LambdaDKG(&_Governance.CallOpts)
}

// LastHalvedAmount is a free data retrieval call binding the contract method 0xa0488e55.
//
// Solidity: function lastHalvedAmount() constant returns(uint256)
func (_Governance *GovernanceCaller) LastHalvedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lastHalvedAmount")
	return *ret0, err
}

// LastHalvedAmount is a free data retrieval call binding the contract method 0xa0488e55.
//
// Solidity: function lastHalvedAmount() constant returns(uint256)
func (_Governance *GovernanceSession) LastHalvedAmount() (*big.Int, error) {
	return _Governance.Contract.LastHalvedAmount(&_Governance.CallOpts)
}

// LastHalvedAmount is a free data retrieval call binding the contract method 0xa0488e55.
//
// Solidity: function lastHalvedAmount() constant returns(uint256)
func (_Governance *GovernanceCallerSession) LastHalvedAmount() (*big.Int, error) {
	return _Governance.Contract.LastHalvedAmount(&_Governance.CallOpts)
}

// LastProposedHeight is a free data retrieval call binding the contract method 0xd767fdc9.
//
// Solidity: function lastProposedHeight(address ) constant returns(uint256)
func (_Governance *GovernanceCaller) LastProposedHeight(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lastProposedHeight", arg0)
	return *ret0, err
}

// LastProposedHeight is a free data retrieval call binding the contract method 0xd767fdc9.
//
// Solidity: function lastProposedHeight(address ) constant returns(uint256)
func (_Governance *GovernanceSession) LastProposedHeight(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.LastProposedHeight(&_Governance.CallOpts, arg0)
}

// LastProposedHeight is a free data retrieval call binding the contract method 0xd767fdc9.
//
// Solidity: function lastProposedHeight(address ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) LastProposedHeight(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.LastProposedHeight(&_Governance.CallOpts, arg0)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() constant returns(uint256)
func (_Governance *GovernanceCaller) LockupPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lockupPeriod")
	return *ret0, err
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() constant returns(uint256)
func (_Governance *GovernanceSession) LockupPeriod() (*big.Int, error) {
	return _Governance.Contract.LockupPeriod(&_Governance.CallOpts)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() constant returns(uint256)
func (_Governance *GovernanceCallerSession) LockupPeriod() (*big.Int, error) {
	return _Governance.Contract.LockupPeriod(&_Governance.CallOpts)
}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() constant returns(uint256)
func (_Governance *GovernanceCaller) MinBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "minBlockInterval")
	return *ret0, err
}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() constant returns(uint256)
func (_Governance *GovernanceSession) MinBlockInterval() (*big.Int, error) {
	return _Governance.Contract.MinBlockInterval(&_Governance.CallOpts)
}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MinBlockInterval() (*big.Int, error) {
	return _Governance.Contract.MinBlockInterval(&_Governance.CallOpts)
}

// MinGasPrice is a free data retrieval call binding the contract method 0xd96ed505.
//
// Solidity: function minGasPrice() constant returns(uint256)
func (_Governance *GovernanceCaller) MinGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "minGasPrice")
	return *ret0, err
}

// MinGasPrice is a free data retrieval call binding the contract method 0xd96ed505.
//
// Solidity: function minGasPrice() constant returns(uint256)
func (_Governance *GovernanceSession) MinGasPrice() (*big.Int, error) {
	return _Governance.Contract.MinGasPrice(&_Governance.CallOpts)
}

// MinGasPrice is a free data retrieval call binding the contract method 0xd96ed505.
//
// Solidity: function minGasPrice() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MinGasPrice() (*big.Int, error) {
	return _Governance.Contract.MinGasPrice(&_Governance.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_Governance *GovernanceCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "minStake")
	return *ret0, err
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_Governance *GovernanceSession) MinStake() (*big.Int, error) {
	return _Governance.Contract.MinStake(&_Governance.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MinStake() (*big.Int, error) {
	return _Governance.Contract.MinStake(&_Governance.CallOpts)
}

// MiningVelocity is a free data retrieval call binding the contract method 0x2357e72a.
//
// Solidity: function miningVelocity() constant returns(uint256)
func (_Governance *GovernanceCaller) MiningVelocity(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "miningVelocity")
	return *ret0, err
}

// MiningVelocity is a free data retrieval call binding the contract method 0x2357e72a.
//
// Solidity: function miningVelocity() constant returns(uint256)
func (_Governance *GovernanceSession) MiningVelocity() (*big.Int, error) {
	return _Governance.Contract.MiningVelocity(&_Governance.CallOpts)
}

// MiningVelocity is a free data retrieval call binding the contract method 0x2357e72a.
//
// Solidity: function miningVelocity() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MiningVelocity() (*big.Int, error) {
	return _Governance.Contract.MiningVelocity(&_Governance.CallOpts)
}

// NextHalvingSupply is a free data retrieval call binding the contract method 0x8ca55162.
//
// Solidity: function nextHalvingSupply() constant returns(uint256)
func (_Governance *GovernanceCaller) NextHalvingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "nextHalvingSupply")
	return *ret0, err
}

// NextHalvingSupply is a free data retrieval call binding the contract method 0x8ca55162.
//
// Solidity: function nextHalvingSupply() constant returns(uint256)
func (_Governance *GovernanceSession) NextHalvingSupply() (*big.Int, error) {
	return _Governance.Contract.NextHalvingSupply(&_Governance.CallOpts)
}

// NextHalvingSupply is a free data retrieval call binding the contract method 0x8ca55162.
//
// Solidity: function nextHalvingSupply() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NextHalvingSupply() (*big.Int, error) {
	return _Governance.Contract.NextHalvingSupply(&_Governance.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) constant returns(address owner, bytes publicKey, uint256 staked, uint256 fined, string name, string email, string location, string url, uint256 unstaked, uint256 unstakedAt)
func (_Governance *GovernanceCaller) Nodes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner      common.Address
	PublicKey  []byte
	Staked     *big.Int
	Fined      *big.Int
	Name       string
	Email      string
	Location   string
	Url        string
	Unstaked   *big.Int
	UnstakedAt *big.Int
}, error) {
	ret := new(struct {
		Owner      common.Address
		PublicKey  []byte
		Staked     *big.Int
		Fined      *big.Int
		Name       string
		Email      string
		Location   string
		Url        string
		Unstaked   *big.Int
		UnstakedAt *big.Int
	})
	out := ret
	err := _Governance.contract.Call(opts, out, "nodes", arg0)
	return *ret, err
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) constant returns(address owner, bytes publicKey, uint256 staked, uint256 fined, string name, string email, string location, string url, uint256 unstaked, uint256 unstakedAt)
func (_Governance *GovernanceSession) Nodes(arg0 *big.Int) (struct {
	Owner      common.Address
	PublicKey  []byte
	Staked     *big.Int
	Fined      *big.Int
	Name       string
	Email      string
	Location   string
	Url        string
	Unstaked   *big.Int
	UnstakedAt *big.Int
}, error) {
	return _Governance.Contract.Nodes(&_Governance.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) constant returns(address owner, bytes publicKey, uint256 staked, uint256 fined, string name, string email, string location, string url, uint256 unstaked, uint256 unstakedAt)
func (_Governance *GovernanceCallerSession) Nodes(arg0 *big.Int) (struct {
	Owner      common.Address
	PublicKey  []byte
	Staked     *big.Int
	Fined      *big.Int
	Name       string
	Email      string
	Location   string
	Url        string
	Unstaked   *big.Int
	UnstakedAt *big.Int
}, error) {
	return _Governance.Contract.Nodes(&_Governance.CallOpts, arg0)
}

// NodesLength is a free data retrieval call binding the contract method 0xf33de6c0.
//
// Solidity: function nodesLength() constant returns(uint256)
func (_Governance *GovernanceCaller) NodesLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "nodesLength")
	return *ret0, err
}

// NodesLength is a free data retrieval call binding the contract method 0xf33de6c0.
//
// Solidity: function nodesLength() constant returns(uint256)
func (_Governance *GovernanceSession) NodesLength() (*big.Int, error) {
	return _Governance.Contract.NodesLength(&_Governance.CallOpts)
}

// NodesLength is a free data retrieval call binding the contract method 0xf33de6c0.
//
// Solidity: function nodesLength() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NodesLength() (*big.Int, error) {
	return _Governance.Contract.NodesLength(&_Governance.CallOpts)
}

// NodesOffsetByAddress is a free data retrieval call binding the contract method 0x85031123.
//
// Solidity: function nodesOffsetByAddress(address ) constant returns(int256)
func (_Governance *GovernanceCaller) NodesOffsetByAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "nodesOffsetByAddress", arg0)
	return *ret0, err
}

// NodesOffsetByAddress is a free data retrieval call binding the contract method 0x85031123.
//
// Solidity: function nodesOffsetByAddress(address ) constant returns(int256)
func (_Governance *GovernanceSession) NodesOffsetByAddress(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.NodesOffsetByAddress(&_Governance.CallOpts, arg0)
}

// NodesOffsetByAddress is a free data retrieval call binding the contract method 0x85031123.
//
// Solidity: function nodesOffsetByAddress(address ) constant returns(int256)
func (_Governance *GovernanceCallerSession) NodesOffsetByAddress(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.NodesOffsetByAddress(&_Governance.CallOpts, arg0)
}

// NodesOffsetByNodeKeyAddress is a free data retrieval call binding the contract method 0xed3fbdf9.
//
// Solidity: function nodesOffsetByNodeKeyAddress(address ) constant returns(int256)
func (_Governance *GovernanceCaller) NodesOffsetByNodeKeyAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "nodesOffsetByNodeKeyAddress", arg0)
	return *ret0, err
}

// NodesOffsetByNodeKeyAddress is a free data retrieval call binding the contract method 0xed3fbdf9.
//
// Solidity: function nodesOffsetByNodeKeyAddress(address ) constant returns(int256)
func (_Governance *GovernanceSession) NodesOffsetByNodeKeyAddress(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.NodesOffsetByNodeKeyAddress(&_Governance.CallOpts, arg0)
}

// NodesOffsetByNodeKeyAddress is a free data retrieval call binding the contract method 0xed3fbdf9.
//
// Solidity: function nodesOffsetByNodeKeyAddress(address ) constant returns(int256)
func (_Governance *GovernanceCallerSession) NodesOffsetByNodeKeyAddress(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.NodesOffsetByNodeKeyAddress(&_Governance.CallOpts, arg0)
}

// NotaryParamAlpha is a free data retrieval call binding the contract method 0x4c523214.
//
// Solidity: function notaryParamAlpha() constant returns(uint256)
func (_Governance *GovernanceCaller) NotaryParamAlpha(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "notaryParamAlpha")
	return *ret0, err
}

// NotaryParamAlpha is a free data retrieval call binding the contract method 0x4c523214.
//
// Solidity: function notaryParamAlpha() constant returns(uint256)
func (_Governance *GovernanceSession) NotaryParamAlpha() (*big.Int, error) {
	return _Governance.Contract.NotaryParamAlpha(&_Governance.CallOpts)
}

// NotaryParamAlpha is a free data retrieval call binding the contract method 0x4c523214.
//
// Solidity: function notaryParamAlpha() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NotaryParamAlpha() (*big.Int, error) {
	return _Governance.Contract.NotaryParamAlpha(&_Governance.CallOpts)
}

// NotaryParamBeta is a free data retrieval call binding the contract method 0x1f29f5c8.
//
// Solidity: function notaryParamBeta() constant returns(uint256)
func (_Governance *GovernanceCaller) NotaryParamBeta(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "notaryParamBeta")
	return *ret0, err
}

// NotaryParamBeta is a free data retrieval call binding the contract method 0x1f29f5c8.
//
// Solidity: function notaryParamBeta() constant returns(uint256)
func (_Governance *GovernanceSession) NotaryParamBeta() (*big.Int, error) {
	return _Governance.Contract.NotaryParamBeta(&_Governance.CallOpts)
}

// NotaryParamBeta is a free data retrieval call binding the contract method 0x1f29f5c8.
//
// Solidity: function notaryParamBeta() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NotaryParamBeta() (*big.Int, error) {
	return _Governance.Contract.NotaryParamBeta(&_Governance.CallOpts)
}

// NotarySetSize is a free data retrieval call binding the contract method 0x0b3441dc.
//
// Solidity: function notarySetSize() constant returns(uint256)
func (_Governance *GovernanceCaller) NotarySetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "notarySetSize")
	return *ret0, err
}

// NotarySetSize is a free data retrieval call binding the contract method 0x0b3441dc.
//
// Solidity: function notarySetSize() constant returns(uint256)
func (_Governance *GovernanceSession) NotarySetSize() (*big.Int, error) {
	return _Governance.Contract.NotarySetSize(&_Governance.CallOpts)
}

// NotarySetSize is a free data retrieval call binding the contract method 0x0b3441dc.
//
// Solidity: function notarySetSize() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NotarySetSize() (*big.Int, error) {
	return _Governance.Contract.NotarySetSize(&_Governance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceCallerSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// RoundHeight is a free data retrieval call binding the contract method 0xaf58ef06.
//
// Solidity: function roundHeight(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCaller) RoundHeight(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "roundHeight", arg0)
	return *ret0, err
}

// RoundHeight is a free data retrieval call binding the contract method 0xaf58ef06.
//
// Solidity: function roundHeight(uint256 ) constant returns(uint256)
func (_Governance *GovernanceSession) RoundHeight(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.RoundHeight(&_Governance.CallOpts, arg0)
}

// RoundHeight is a free data retrieval call binding the contract method 0xaf58ef06.
//
// Solidity: function roundHeight(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) RoundHeight(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.RoundHeight(&_Governance.CallOpts, arg0)
}

// RoundLength is a free data retrieval call binding the contract method 0x8b649b94.
//
// Solidity: function roundLength() constant returns(uint256)
func (_Governance *GovernanceCaller) RoundLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "roundLength")
	return *ret0, err
}

// RoundLength is a free data retrieval call binding the contract method 0x8b649b94.
//
// Solidity: function roundLength() constant returns(uint256)
func (_Governance *GovernanceSession) RoundLength() (*big.Int, error) {
	return _Governance.Contract.RoundLength(&_Governance.CallOpts)
}

// RoundLength is a free data retrieval call binding the contract method 0x8b649b94.
//
// Solidity: function roundLength() constant returns(uint256)
func (_Governance *GovernanceCallerSession) RoundLength() (*big.Int, error) {
	return _Governance.Contract.RoundLength(&_Governance.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Governance *GovernanceCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "totalStaked")
	return *ret0, err
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Governance *GovernanceSession) TotalStaked() (*big.Int, error) {
	return _Governance.Contract.TotalStaked(&_Governance.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Governance *GovernanceCallerSession) TotalStaked() (*big.Int, error) {
	return _Governance.Contract.TotalStaked(&_Governance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Governance *GovernanceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Governance *GovernanceSession) TotalSupply() (*big.Int, error) {
	return _Governance.Contract.TotalSupply(&_Governance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Governance *GovernanceCallerSession) TotalSupply() (*big.Int, error) {
	return _Governance.Contract.TotalSupply(&_Governance.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() constant returns(bool)
func (_Governance *GovernanceCaller) Withdrawable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "withdrawable")
	return *ret0, err
}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() constant returns(bool)
func (_Governance *GovernanceSession) Withdrawable() (bool, error) {
	return _Governance.Contract.Withdrawable(&_Governance.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() constant returns(bool)
func (_Governance *GovernanceCallerSession) Withdrawable() (bool, error) {
	return _Governance.Contract.Withdrawable(&_Governance.CallOpts)
}

// AddDKGComplaint is a paid mutator transaction binding the contract method 0xab6a4013.
//
// Solidity: function addDKGComplaint(bytes Complaint) returns()
func (_Governance *GovernanceTransactor) AddDKGComplaint(opts *bind.TransactOpts, Complaint []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addDKGComplaint", Complaint)
}

// AddDKGComplaint is a paid mutator transaction binding the contract method 0xab6a4013.
//
// Solidity: function addDKGComplaint(bytes Complaint) returns()
func (_Governance *GovernanceSession) AddDKGComplaint(Complaint []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGComplaint(&_Governance.TransactOpts, Complaint)
}

// AddDKGComplaint is a paid mutator transaction binding the contract method 0xab6a4013.
//
// Solidity: function addDKGComplaint(bytes Complaint) returns()
func (_Governance *GovernanceTransactorSession) AddDKGComplaint(Complaint []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGComplaint(&_Governance.TransactOpts, Complaint)
}

// AddDKGFinalize is a paid mutator transaction binding the contract method 0x29891e74.
//
// Solidity: function addDKGFinalize(bytes Finalize) returns()
func (_Governance *GovernanceTransactor) AddDKGFinalize(opts *bind.TransactOpts, Finalize []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addDKGFinalize", Finalize)
}

// AddDKGFinalize is a paid mutator transaction binding the contract method 0x29891e74.
//
// Solidity: function addDKGFinalize(bytes Finalize) returns()
func (_Governance *GovernanceSession) AddDKGFinalize(Finalize []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGFinalize(&_Governance.TransactOpts, Finalize)
}

// AddDKGFinalize is a paid mutator transaction binding the contract method 0x29891e74.
//
// Solidity: function addDKGFinalize(bytes Finalize) returns()
func (_Governance *GovernanceTransactorSession) AddDKGFinalize(Finalize []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGFinalize(&_Governance.TransactOpts, Finalize)
}

// AddDKGMPKReady is a paid mutator transaction binding the contract method 0x22f8a889.
//
// Solidity: function addDKGMPKReady(bytes MPKReady) returns()
func (_Governance *GovernanceTransactor) AddDKGMPKReady(opts *bind.TransactOpts, MPKReady []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addDKGMPKReady", MPKReady)
}

// AddDKGMPKReady is a paid mutator transaction binding the contract method 0x22f8a889.
//
// Solidity: function addDKGMPKReady(bytes MPKReady) returns()
func (_Governance *GovernanceSession) AddDKGMPKReady(MPKReady []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGMPKReady(&_Governance.TransactOpts, MPKReady)
}

// AddDKGMPKReady is a paid mutator transaction binding the contract method 0x22f8a889.
//
// Solidity: function addDKGMPKReady(bytes MPKReady) returns()
func (_Governance *GovernanceTransactorSession) AddDKGMPKReady(MPKReady []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGMPKReady(&_Governance.TransactOpts, MPKReady)
}

// AddDKGMasterPublicKey is a paid mutator transaction binding the contract method 0x19cf10bc.
//
// Solidity: function addDKGMasterPublicKey(bytes PublicKey) returns()
func (_Governance *GovernanceTransactor) AddDKGMasterPublicKey(opts *bind.TransactOpts, PublicKey []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addDKGMasterPublicKey", PublicKey)
}

// AddDKGMasterPublicKey is a paid mutator transaction binding the contract method 0x19cf10bc.
//
// Solidity: function addDKGMasterPublicKey(bytes PublicKey) returns()
func (_Governance *GovernanceSession) AddDKGMasterPublicKey(PublicKey []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGMasterPublicKey(&_Governance.TransactOpts, PublicKey)
}

// AddDKGMasterPublicKey is a paid mutator transaction binding the contract method 0x19cf10bc.
//
// Solidity: function addDKGMasterPublicKey(bytes PublicKey) returns()
func (_Governance *GovernanceTransactorSession) AddDKGMasterPublicKey(PublicKey []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGMasterPublicKey(&_Governance.TransactOpts, PublicKey)
}

// AddDKGSuccess is a paid mutator transaction binding the contract method 0x4ccaa59d.
//
// Solidity: function addDKGSuccess(bytes Success) returns()
func (_Governance *GovernanceTransactor) AddDKGSuccess(opts *bind.TransactOpts, Success []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addDKGSuccess", Success)
}

// AddDKGSuccess is a paid mutator transaction binding the contract method 0x4ccaa59d.
//
// Solidity: function addDKGSuccess(bytes Success) returns()
func (_Governance *GovernanceSession) AddDKGSuccess(Success []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGSuccess(&_Governance.TransactOpts, Success)
}

// AddDKGSuccess is a paid mutator transaction binding the contract method 0x4ccaa59d.
//
// Solidity: function addDKGSuccess(bytes Success) returns()
func (_Governance *GovernanceTransactorSession) AddDKGSuccess(Success []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGSuccess(&_Governance.TransactOpts, Success)
}

// PayFine is a paid mutator transaction binding the contract method 0x3edfa229.
//
// Solidity: function payFine(address NodeAddress) returns()
func (_Governance *GovernanceTransactor) PayFine(opts *bind.TransactOpts, NodeAddress common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "payFine", NodeAddress)
}

// PayFine is a paid mutator transaction binding the contract method 0x3edfa229.
//
// Solidity: function payFine(address NodeAddress) returns()
func (_Governance *GovernanceSession) PayFine(NodeAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.PayFine(&_Governance.TransactOpts, NodeAddress)
}

// PayFine is a paid mutator transaction binding the contract method 0x3edfa229.
//
// Solidity: function payFine(address NodeAddress) returns()
func (_Governance *GovernanceTransactorSession) PayFine(NodeAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.PayFine(&_Governance.TransactOpts, NodeAddress)
}

// ProposeCRS is a paid mutator transaction binding the contract method 0xc448af34.
//
// Solidity: function proposeCRS(uint256 Round, bytes SignedCRS) returns()
func (_Governance *GovernanceTransactor) ProposeCRS(opts *bind.TransactOpts, Round *big.Int, SignedCRS []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "proposeCRS", Round, SignedCRS)
}

// ProposeCRS is a paid mutator transaction binding the contract method 0xc448af34.
//
// Solidity: function proposeCRS(uint256 Round, bytes SignedCRS) returns()
func (_Governance *GovernanceSession) ProposeCRS(Round *big.Int, SignedCRS []byte) (*types.Transaction, error) {
	return _Governance.Contract.ProposeCRS(&_Governance.TransactOpts, Round, SignedCRS)
}

// ProposeCRS is a paid mutator transaction binding the contract method 0xc448af34.
//
// Solidity: function proposeCRS(uint256 Round, bytes SignedCRS) returns()
func (_Governance *GovernanceTransactorSession) ProposeCRS(Round *big.Int, SignedCRS []byte) (*types.Transaction, error) {
	return _Governance.Contract.ProposeCRS(&_Governance.TransactOpts, Round, SignedCRS)
}

// Register is a paid mutator transaction binding the contract method 0xd242f4cc.
//
// Solidity: function register(bytes PublicKey, string Name, string Email, string Location, string Url) returns()
func (_Governance *GovernanceTransactor) Register(opts *bind.TransactOpts, PublicKey []byte, Name string, Email string, Location string, Url string) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "register", PublicKey, Name, Email, Location, Url)
}

// Register is a paid mutator transaction binding the contract method 0xd242f4cc.
//
// Solidity: function register(bytes PublicKey, string Name, string Email, string Location, string Url) returns()
func (_Governance *GovernanceSession) Register(PublicKey []byte, Name string, Email string, Location string, Url string) (*types.Transaction, error) {
	return _Governance.Contract.Register(&_Governance.TransactOpts, PublicKey, Name, Email, Location, Url)
}

// Register is a paid mutator transaction binding the contract method 0xd242f4cc.
//
// Solidity: function register(bytes PublicKey, string Name, string Email, string Location, string Url) returns()
func (_Governance *GovernanceTransactorSession) Register(PublicKey []byte, Name string, Email string, Location string, Url string) (*types.Transaction, error) {
	return _Governance.Contract.Register(&_Governance.TransactOpts, PublicKey, Name, Email, Location, Url)
}

// ReplaceNodePublicKey is a paid mutator transaction binding the contract method 0x3f007c20.
//
// Solidity: function replaceNodePublicKey(bytes NewPublicKey) returns()
func (_Governance *GovernanceTransactor) ReplaceNodePublicKey(opts *bind.TransactOpts, NewPublicKey []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "replaceNodePublicKey", NewPublicKey)
}

// ReplaceNodePublicKey is a paid mutator transaction binding the contract method 0x3f007c20.
//
// Solidity: function replaceNodePublicKey(bytes NewPublicKey) returns()
func (_Governance *GovernanceSession) ReplaceNodePublicKey(NewPublicKey []byte) (*types.Transaction, error) {
	return _Governance.Contract.ReplaceNodePublicKey(&_Governance.TransactOpts, NewPublicKey)
}

// ReplaceNodePublicKey is a paid mutator transaction binding the contract method 0x3f007c20.
//
// Solidity: function replaceNodePublicKey(bytes NewPublicKey) returns()
func (_Governance *GovernanceTransactorSession) ReplaceNodePublicKey(NewPublicKey []byte) (*types.Transaction, error) {
	return _Governance.Contract.ReplaceNodePublicKey(&_Governance.TransactOpts, NewPublicKey)
}

// Report is a paid mutator transaction binding the contract method 0x320c0826.
//
// Solidity: function report(uint256 Type, bytes Arg1, bytes Arg2) returns()
func (_Governance *GovernanceTransactor) Report(opts *bind.TransactOpts, Type *big.Int, Arg1 []byte, Arg2 []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "report", Type, Arg1, Arg2)
}

// Report is a paid mutator transaction binding the contract method 0x320c0826.
//
// Solidity: function report(uint256 Type, bytes Arg1, bytes Arg2) returns()
func (_Governance *GovernanceSession) Report(Type *big.Int, Arg1 []byte, Arg2 []byte) (*types.Transaction, error) {
	return _Governance.Contract.Report(&_Governance.TransactOpts, Type, Arg1, Arg2)
}

// Report is a paid mutator transaction binding the contract method 0x320c0826.
//
// Solidity: function report(uint256 Type, bytes Arg1, bytes Arg2) returns()
func (_Governance *GovernanceTransactorSession) Report(Type *big.Int, Arg1 []byte, Arg2 []byte) (*types.Transaction, error) {
	return _Governance.Contract.Report(&_Governance.TransactOpts, Type, Arg1, Arg2)
}

// ResetDKG is a paid mutator transaction binding the contract method 0x57316157.
//
// Solidity: function resetDKG(bytes NewSignedCRS) returns()
func (_Governance *GovernanceTransactor) ResetDKG(opts *bind.TransactOpts, NewSignedCRS []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "resetDKG", NewSignedCRS)
}

// ResetDKG is a paid mutator transaction binding the contract method 0x57316157.
//
// Solidity: function resetDKG(bytes NewSignedCRS) returns()
func (_Governance *GovernanceSession) ResetDKG(NewSignedCRS []byte) (*types.Transaction, error) {
	return _Governance.Contract.ResetDKG(&_Governance.TransactOpts, NewSignedCRS)
}

// ResetDKG is a paid mutator transaction binding the contract method 0x57316157.
//
// Solidity: function resetDKG(bytes NewSignedCRS) returns()
func (_Governance *GovernanceTransactorSession) ResetDKG(NewSignedCRS []byte) (*types.Transaction, error) {
	return _Governance.Contract.ResetDKG(&_Governance.TransactOpts, NewSignedCRS)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Governance *GovernanceTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Governance *GovernanceSession) Stake() (*types.Transaction, error) {
	return _Governance.Contract.Stake(&_Governance.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Governance *GovernanceTransactorSession) Stake() (*types.Transaction, error) {
	return _Governance.Contract.Stake(&_Governance.TransactOpts)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x4264ecf9.
//
// Solidity: function transferNodeOwnership(address NewOwner) returns()
func (_Governance *GovernanceTransactor) TransferNodeOwnership(opts *bind.TransactOpts, NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "transferNodeOwnership", NewOwner)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x4264ecf9.
//
// Solidity: function transferNodeOwnership(address NewOwner) returns()
func (_Governance *GovernanceSession) TransferNodeOwnership(NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferNodeOwnership(&_Governance.TransactOpts, NewOwner)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x4264ecf9.
//
// Solidity: function transferNodeOwnership(address NewOwner) returns()
func (_Governance *GovernanceTransactorSession) TransferNodeOwnership(NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferNodeOwnership(&_Governance.TransactOpts, NewOwner)
}

// TransferNodeOwnershipByFoundation is a paid mutator transaction binding the contract method 0x7ecfdeca.
//
// Solidity: function transferNodeOwnershipByFoundation(address OldOwner, address NewOwner) returns()
func (_Governance *GovernanceTransactor) TransferNodeOwnershipByFoundation(opts *bind.TransactOpts, OldOwner common.Address, NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "transferNodeOwnershipByFoundation", OldOwner, NewOwner)
}

// TransferNodeOwnershipByFoundation is a paid mutator transaction binding the contract method 0x7ecfdeca.
//
// Solidity: function transferNodeOwnershipByFoundation(address OldOwner, address NewOwner) returns()
func (_Governance *GovernanceSession) TransferNodeOwnershipByFoundation(OldOwner common.Address, NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferNodeOwnershipByFoundation(&_Governance.TransactOpts, OldOwner, NewOwner)
}

// TransferNodeOwnershipByFoundation is a paid mutator transaction binding the contract method 0x7ecfdeca.
//
// Solidity: function transferNodeOwnershipByFoundation(address OldOwner, address NewOwner) returns()
func (_Governance *GovernanceTransactorSession) TransferNodeOwnershipByFoundation(OldOwner common.Address, NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferNodeOwnershipByFoundation(&_Governance.TransactOpts, OldOwner, NewOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address NewOwner) returns()
func (_Governance *GovernanceTransactor) TransferOwnership(opts *bind.TransactOpts, NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "transferOwnership", NewOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address NewOwner) returns()
func (_Governance *GovernanceSession) TransferOwnership(NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, NewOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address NewOwner) returns()
func (_Governance *GovernanceTransactorSession) TransferOwnership(NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, NewOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 Amount) returns()
func (_Governance *GovernanceTransactor) Unstake(opts *bind.TransactOpts, Amount *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "unstake", Amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 Amount) returns()
func (_Governance *GovernanceSession) Unstake(Amount *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Unstake(&_Governance.TransactOpts, Amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 Amount) returns()
func (_Governance *GovernanceTransactorSession) Unstake(Amount *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Unstake(&_Governance.TransactOpts, Amount)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x11c7c321.
//
// Solidity: function updateConfiguration(uint256 MinStake, uint256 LockupPeriod, uint256 MinGasPrice, uint256 BlockGasLimit, uint256 LambdaBA, uint256 LambdaDKG, uint256 NotaryParamAlpha, uint256 NotaryParamBeta, uint256 RoundLength, uint256 MinBlockInterval, uint256[] FineValues) returns()
func (_Governance *GovernanceTransactor) UpdateConfiguration(opts *bind.TransactOpts, MinStake *big.Int, LockupPeriod *big.Int, MinGasPrice *big.Int, BlockGasLimit *big.Int, LambdaBA *big.Int, LambdaDKG *big.Int, NotaryParamAlpha *big.Int, NotaryParamBeta *big.Int, RoundLength *big.Int, MinBlockInterval *big.Int, FineValues []*big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "updateConfiguration", MinStake, LockupPeriod, MinGasPrice, BlockGasLimit, LambdaBA, LambdaDKG, NotaryParamAlpha, NotaryParamBeta, RoundLength, MinBlockInterval, FineValues)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x11c7c321.
//
// Solidity: function updateConfiguration(uint256 MinStake, uint256 LockupPeriod, uint256 MinGasPrice, uint256 BlockGasLimit, uint256 LambdaBA, uint256 LambdaDKG, uint256 NotaryParamAlpha, uint256 NotaryParamBeta, uint256 RoundLength, uint256 MinBlockInterval, uint256[] FineValues) returns()
func (_Governance *GovernanceSession) UpdateConfiguration(MinStake *big.Int, LockupPeriod *big.Int, MinGasPrice *big.Int, BlockGasLimit *big.Int, LambdaBA *big.Int, LambdaDKG *big.Int, NotaryParamAlpha *big.Int, NotaryParamBeta *big.Int, RoundLength *big.Int, MinBlockInterval *big.Int, FineValues []*big.Int) (*types.Transaction, error) {
	return _Governance.Contract.UpdateConfiguration(&_Governance.TransactOpts, MinStake, LockupPeriod, MinGasPrice, BlockGasLimit, LambdaBA, LambdaDKG, NotaryParamAlpha, NotaryParamBeta, RoundLength, MinBlockInterval, FineValues)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x11c7c321.
//
// Solidity: function updateConfiguration(uint256 MinStake, uint256 LockupPeriod, uint256 MinGasPrice, uint256 BlockGasLimit, uint256 LambdaBA, uint256 LambdaDKG, uint256 NotaryParamAlpha, uint256 NotaryParamBeta, uint256 RoundLength, uint256 MinBlockInterval, uint256[] FineValues) returns()
func (_Governance *GovernanceTransactorSession) UpdateConfiguration(MinStake *big.Int, LockupPeriod *big.Int, MinGasPrice *big.Int, BlockGasLimit *big.Int, LambdaBA *big.Int, LambdaDKG *big.Int, NotaryParamAlpha *big.Int, NotaryParamBeta *big.Int, RoundLength *big.Int, MinBlockInterval *big.Int, FineValues []*big.Int) (*types.Transaction, error) {
	return _Governance.Contract.UpdateConfiguration(&_Governance.TransactOpts, MinStake, LockupPeriod, MinGasPrice, BlockGasLimit, LambdaBA, LambdaDKG, NotaryParamAlpha, NotaryParamBeta, RoundLength, MinBlockInterval, FineValues)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Governance *GovernanceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Governance *GovernanceSession) Withdraw() (*types.Transaction, error) {
	return _Governance.Contract.Withdraw(&_Governance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Governance *GovernanceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Governance.Contract.Withdraw(&_Governance.TransactOpts)
}

// GovernanceCRSProposedIterator is returned from FilterCRSProposed and is used to iterate over the raw logs and unpacked data for CRSProposed events raised by the Governance contract.
type GovernanceCRSProposedIterator struct {
	Event *GovernanceCRSProposed // Event containing the contract specifics and raw log

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
func (it *GovernanceCRSProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceCRSProposed)
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
		it.Event = new(GovernanceCRSProposed)
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
func (it *GovernanceCRSProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceCRSProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceCRSProposed represents a CRSProposed event raised by the Governance contract.
type GovernanceCRSProposed struct {
	Round *big.Int
	CRS   [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCRSProposed is a free log retrieval operation binding the contract event 0xd9b8f7e45cd2897523743eb19fd2021ba458d2962753a38bc31223e9b05ad408.
//
// Solidity: event CRSProposed(uint256 indexed Round, bytes32 CRS)
func (_Governance *GovernanceFilterer) FilterCRSProposed(opts *bind.FilterOpts, Round []*big.Int) (*GovernanceCRSProposedIterator, error) {

	var RoundRule []interface{}
	for _, RoundItem := range Round {
		RoundRule = append(RoundRule, RoundItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "CRSProposed", RoundRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceCRSProposedIterator{contract: _Governance.contract, event: "CRSProposed", logs: logs, sub: sub}, nil
}

// WatchCRSProposed is a free log subscription operation binding the contract event 0xd9b8f7e45cd2897523743eb19fd2021ba458d2962753a38bc31223e9b05ad408.
//
// Solidity: event CRSProposed(uint256 indexed Round, bytes32 CRS)
func (_Governance *GovernanceFilterer) WatchCRSProposed(opts *bind.WatchOpts, sink chan<- *GovernanceCRSProposed, Round []*big.Int) (event.Subscription, error) {

	var RoundRule []interface{}
	for _, RoundItem := range Round {
		RoundRule = append(RoundRule, RoundItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "CRSProposed", RoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceCRSProposed)
				if err := _Governance.contract.UnpackLog(event, "CRSProposed", log); err != nil {
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

// GovernanceConfigurationChangedIterator is returned from FilterConfigurationChanged and is used to iterate over the raw logs and unpacked data for ConfigurationChanged events raised by the Governance contract.
type GovernanceConfigurationChangedIterator struct {
	Event *GovernanceConfigurationChanged // Event containing the contract specifics and raw log

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
func (it *GovernanceConfigurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceConfigurationChanged)
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
		it.Event = new(GovernanceConfigurationChanged)
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
func (it *GovernanceConfigurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceConfigurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceConfigurationChanged represents a ConfigurationChanged event raised by the Governance contract.
type GovernanceConfigurationChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterConfigurationChanged is a free log retrieval operation binding the contract event 0xb6aa5c479f96dcbdbe1d8b70883ec95e8799f8e602ee347ed972a22c974b14c0.
//
// Solidity: event ConfigurationChanged()
func (_Governance *GovernanceFilterer) FilterConfigurationChanged(opts *bind.FilterOpts) (*GovernanceConfigurationChangedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ConfigurationChanged")
	if err != nil {
		return nil, err
	}
	return &GovernanceConfigurationChangedIterator{contract: _Governance.contract, event: "ConfigurationChanged", logs: logs, sub: sub}, nil
}

// WatchConfigurationChanged is a free log subscription operation binding the contract event 0xb6aa5c479f96dcbdbe1d8b70883ec95e8799f8e602ee347ed972a22c974b14c0.
//
// Solidity: event ConfigurationChanged()
func (_Governance *GovernanceFilterer) WatchConfigurationChanged(opts *bind.WatchOpts, sink chan<- *GovernanceConfigurationChanged) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ConfigurationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceConfigurationChanged)
				if err := _Governance.contract.UnpackLog(event, "ConfigurationChanged", log); err != nil {
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

// GovernanceDKGResetIterator is returned from FilterDKGReset and is used to iterate over the raw logs and unpacked data for DKGReset events raised by the Governance contract.
type GovernanceDKGResetIterator struct {
	Event *GovernanceDKGReset // Event containing the contract specifics and raw log

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
func (it *GovernanceDKGResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceDKGReset)
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
		it.Event = new(GovernanceDKGReset)
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
func (it *GovernanceDKGResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceDKGResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceDKGReset represents a DKGReset event raised by the Governance contract.
type GovernanceDKGReset struct {
	Round       *big.Int
	BlockHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDKGReset is a free log retrieval operation binding the contract event 0x1ef9b10077100d06aefe929eb17ff481feebce38b8a770e5b7d744febfab9f50.
//
// Solidity: event DKGReset(uint256 indexed Round, uint256 BlockHeight)
func (_Governance *GovernanceFilterer) FilterDKGReset(opts *bind.FilterOpts, Round []*big.Int) (*GovernanceDKGResetIterator, error) {

	var RoundRule []interface{}
	for _, RoundItem := range Round {
		RoundRule = append(RoundRule, RoundItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "DKGReset", RoundRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceDKGResetIterator{contract: _Governance.contract, event: "DKGReset", logs: logs, sub: sub}, nil
}

// WatchDKGReset is a free log subscription operation binding the contract event 0x1ef9b10077100d06aefe929eb17ff481feebce38b8a770e5b7d744febfab9f50.
//
// Solidity: event DKGReset(uint256 indexed Round, uint256 BlockHeight)
func (_Governance *GovernanceFilterer) WatchDKGReset(opts *bind.WatchOpts, sink chan<- *GovernanceDKGReset, Round []*big.Int) (event.Subscription, error) {

	var RoundRule []interface{}
	for _, RoundItem := range Round {
		RoundRule = append(RoundRule, RoundItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "DKGReset", RoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceDKGReset)
				if err := _Governance.contract.UnpackLog(event, "DKGReset", log); err != nil {
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

// GovernanceFinePaidIterator is returned from FilterFinePaid and is used to iterate over the raw logs and unpacked data for FinePaid events raised by the Governance contract.
type GovernanceFinePaidIterator struct {
	Event *GovernanceFinePaid // Event containing the contract specifics and raw log

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
func (it *GovernanceFinePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceFinePaid)
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
		it.Event = new(GovernanceFinePaid)
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
func (it *GovernanceFinePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceFinePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceFinePaid represents a FinePaid event raised by the Governance contract.
type GovernanceFinePaid struct {
	NodeAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinePaid is a free log retrieval operation binding the contract event 0x34af8d99a30dafb3157ed162369d603f544a74be3858f8c7ac9e159829589c25.
//
// Solidity: event FinePaid(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) FilterFinePaid(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceFinePaidIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "FinePaid", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceFinePaidIterator{contract: _Governance.contract, event: "FinePaid", logs: logs, sub: sub}, nil
}

// WatchFinePaid is a free log subscription operation binding the contract event 0x34af8d99a30dafb3157ed162369d603f544a74be3858f8c7ac9e159829589c25.
//
// Solidity: event FinePaid(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) WatchFinePaid(opts *bind.WatchOpts, sink chan<- *GovernanceFinePaid, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "FinePaid", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceFinePaid)
				if err := _Governance.contract.UnpackLog(event, "FinePaid", log); err != nil {
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

// GovernanceFinedIterator is returned from FilterFined and is used to iterate over the raw logs and unpacked data for Fined events raised by the Governance contract.
type GovernanceFinedIterator struct {
	Event *GovernanceFined // Event containing the contract specifics and raw log

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
func (it *GovernanceFinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceFined)
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
		it.Event = new(GovernanceFined)
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
func (it *GovernanceFinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceFinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceFined represents a Fined event raised by the Governance contract.
type GovernanceFined struct {
	NodeAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFined is a free log retrieval operation binding the contract event 0x00913d46aef0f0d115d70ea1c7c23198505f577d1d1916cc60710ca2204ae6ae.
//
// Solidity: event Fined(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) FilterFined(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceFinedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Fined", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceFinedIterator{contract: _Governance.contract, event: "Fined", logs: logs, sub: sub}, nil
}

// WatchFined is a free log subscription operation binding the contract event 0x00913d46aef0f0d115d70ea1c7c23198505f577d1d1916cc60710ca2204ae6ae.
//
// Solidity: event Fined(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) WatchFined(opts *bind.WatchOpts, sink chan<- *GovernanceFined, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Fined", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceFined)
				if err := _Governance.contract.UnpackLog(event, "Fined", log); err != nil {
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

// GovernanceNodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the Governance contract.
type GovernanceNodeAddedIterator struct {
	Event *GovernanceNodeAdded // Event containing the contract specifics and raw log

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
func (it *GovernanceNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNodeAdded)
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
		it.Event = new(GovernanceNodeAdded)
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
func (it *GovernanceNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNodeAdded represents a NodeAdded event raised by the Governance contract.
type GovernanceNodeAdded struct {
	NodeAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address indexed NodeAddress)
func (_Governance *GovernanceFilterer) FilterNodeAdded(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceNodeAddedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NodeAdded", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNodeAddedIterator{contract: _Governance.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address indexed NodeAddress)
func (_Governance *GovernanceFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *GovernanceNodeAdded, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NodeAdded", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNodeAdded)
				if err := _Governance.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

// GovernanceNodeOwnershipTransferedIterator is returned from FilterNodeOwnershipTransfered and is used to iterate over the raw logs and unpacked data for NodeOwnershipTransfered events raised by the Governance contract.
type GovernanceNodeOwnershipTransferedIterator struct {
	Event *GovernanceNodeOwnershipTransfered // Event containing the contract specifics and raw log

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
func (it *GovernanceNodeOwnershipTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNodeOwnershipTransfered)
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
		it.Event = new(GovernanceNodeOwnershipTransfered)
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
func (it *GovernanceNodeOwnershipTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNodeOwnershipTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNodeOwnershipTransfered represents a NodeOwnershipTransfered event raised by the Governance contract.
type GovernanceNodeOwnershipTransfered struct {
	NodeAddress     common.Address
	NewOwnerAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeOwnershipTransfered is a free log retrieval operation binding the contract event 0x63616cf4caeee172ed91719b5e000da424d56855ee726aedb0a1d74fb3f1db2a.
//
// Solidity: event NodeOwnershipTransfered(address indexed NodeAddress, address indexed NewOwnerAddress)
func (_Governance *GovernanceFilterer) FilterNodeOwnershipTransfered(opts *bind.FilterOpts, NodeAddress []common.Address, NewOwnerAddress []common.Address) (*GovernanceNodeOwnershipTransferedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}
	var NewOwnerAddressRule []interface{}
	for _, NewOwnerAddressItem := range NewOwnerAddress {
		NewOwnerAddressRule = append(NewOwnerAddressRule, NewOwnerAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NodeOwnershipTransfered", NodeAddressRule, NewOwnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNodeOwnershipTransferedIterator{contract: _Governance.contract, event: "NodeOwnershipTransfered", logs: logs, sub: sub}, nil
}

// WatchNodeOwnershipTransfered is a free log subscription operation binding the contract event 0x63616cf4caeee172ed91719b5e000da424d56855ee726aedb0a1d74fb3f1db2a.
//
// Solidity: event NodeOwnershipTransfered(address indexed NodeAddress, address indexed NewOwnerAddress)
func (_Governance *GovernanceFilterer) WatchNodeOwnershipTransfered(opts *bind.WatchOpts, sink chan<- *GovernanceNodeOwnershipTransfered, NodeAddress []common.Address, NewOwnerAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}
	var NewOwnerAddressRule []interface{}
	for _, NewOwnerAddressItem := range NewOwnerAddress {
		NewOwnerAddressRule = append(NewOwnerAddressRule, NewOwnerAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NodeOwnershipTransfered", NodeAddressRule, NewOwnerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNodeOwnershipTransfered)
				if err := _Governance.contract.UnpackLog(event, "NodeOwnershipTransfered", log); err != nil {
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

// GovernanceNodePublicKeyReplacedIterator is returned from FilterNodePublicKeyReplaced and is used to iterate over the raw logs and unpacked data for NodePublicKeyReplaced events raised by the Governance contract.
type GovernanceNodePublicKeyReplacedIterator struct {
	Event *GovernanceNodePublicKeyReplaced // Event containing the contract specifics and raw log

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
func (it *GovernanceNodePublicKeyReplacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNodePublicKeyReplaced)
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
		it.Event = new(GovernanceNodePublicKeyReplaced)
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
func (it *GovernanceNodePublicKeyReplacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNodePublicKeyReplacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNodePublicKeyReplaced represents a NodePublicKeyReplaced event raised by the Governance contract.
type GovernanceNodePublicKeyReplaced struct {
	NodeAddress common.Address
	PublicKey   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodePublicKeyReplaced is a free log retrieval operation binding the contract event 0x8aa688c270fa0768c37a1f455f711b6ea1a925f04db9ef433e01dcba78b80432.
//
// Solidity: event NodePublicKeyReplaced(address indexed NodeAddress, bytes PublicKey)
func (_Governance *GovernanceFilterer) FilterNodePublicKeyReplaced(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceNodePublicKeyReplacedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NodePublicKeyReplaced", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNodePublicKeyReplacedIterator{contract: _Governance.contract, event: "NodePublicKeyReplaced", logs: logs, sub: sub}, nil
}

// WatchNodePublicKeyReplaced is a free log subscription operation binding the contract event 0x8aa688c270fa0768c37a1f455f711b6ea1a925f04db9ef433e01dcba78b80432.
//
// Solidity: event NodePublicKeyReplaced(address indexed NodeAddress, bytes PublicKey)
func (_Governance *GovernanceFilterer) WatchNodePublicKeyReplaced(opts *bind.WatchOpts, sink chan<- *GovernanceNodePublicKeyReplaced, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NodePublicKeyReplaced", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNodePublicKeyReplaced)
				if err := _Governance.contract.UnpackLog(event, "NodePublicKeyReplaced", log); err != nil {
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

// GovernanceNodeRemovedIterator is returned from FilterNodeRemoved and is used to iterate over the raw logs and unpacked data for NodeRemoved events raised by the Governance contract.
type GovernanceNodeRemovedIterator struct {
	Event *GovernanceNodeRemoved // Event containing the contract specifics and raw log

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
func (it *GovernanceNodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNodeRemoved)
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
		it.Event = new(GovernanceNodeRemoved)
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
func (it *GovernanceNodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNodeRemoved represents a NodeRemoved event raised by the Governance contract.
type GovernanceNodeRemoved struct {
	NodeAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeRemoved is a free log retrieval operation binding the contract event 0xcfc24166db4bb677e857cacabd1541fb2b30645021b27c5130419589b84db52b.
//
// Solidity: event NodeRemoved(address indexed NodeAddress)
func (_Governance *GovernanceFilterer) FilterNodeRemoved(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceNodeRemovedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NodeRemoved", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNodeRemovedIterator{contract: _Governance.contract, event: "NodeRemoved", logs: logs, sub: sub}, nil
}

// WatchNodeRemoved is a free log subscription operation binding the contract event 0xcfc24166db4bb677e857cacabd1541fb2b30645021b27c5130419589b84db52b.
//
// Solidity: event NodeRemoved(address indexed NodeAddress)
func (_Governance *GovernanceFilterer) WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *GovernanceNodeRemoved, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NodeRemoved", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNodeRemoved)
				if err := _Governance.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
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

// GovernanceReportedIterator is returned from FilterReported and is used to iterate over the raw logs and unpacked data for Reported events raised by the Governance contract.
type GovernanceReportedIterator struct {
	Event *GovernanceReported // Event containing the contract specifics and raw log

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
func (it *GovernanceReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceReported)
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
		it.Event = new(GovernanceReported)
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
func (it *GovernanceReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceReported represents a Reported event raised by the Governance contract.
type GovernanceReported struct {
	NodeAddress common.Address
	Type        *big.Int
	Arg1        []byte
	Arg2        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReported is a free log retrieval operation binding the contract event 0x0b4ded2614643cbc6be3ac9a1ed3188f2667636de75ea5250e190ad6af4b43d9.
//
// Solidity: event Reported(address indexed NodeAddress, uint256 Type, bytes Arg1, bytes Arg2)
func (_Governance *GovernanceFilterer) FilterReported(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceReportedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Reported", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceReportedIterator{contract: _Governance.contract, event: "Reported", logs: logs, sub: sub}, nil
}

// WatchReported is a free log subscription operation binding the contract event 0x0b4ded2614643cbc6be3ac9a1ed3188f2667636de75ea5250e190ad6af4b43d9.
//
// Solidity: event Reported(address indexed NodeAddress, uint256 Type, bytes Arg1, bytes Arg2)
func (_Governance *GovernanceFilterer) WatchReported(opts *bind.WatchOpts, sink chan<- *GovernanceReported, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Reported", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceReported)
				if err := _Governance.contract.UnpackLog(event, "Reported", log); err != nil {
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

// GovernanceStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Governance contract.
type GovernanceStakedIterator struct {
	Event *GovernanceStaked // Event containing the contract specifics and raw log

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
func (it *GovernanceStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceStaked)
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
		it.Event = new(GovernanceStaked)
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
func (it *GovernanceStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceStaked represents a Staked event raised by the Governance contract.
type GovernanceStaked struct {
	NodeAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) FilterStaked(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceStakedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Staked", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceStakedIterator{contract: _Governance.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *GovernanceStaked, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Staked", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceStaked)
				if err := _Governance.contract.UnpackLog(event, "Staked", log); err != nil {
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

// GovernanceUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Governance contract.
type GovernanceUnstakedIterator struct {
	Event *GovernanceUnstaked // Event containing the contract specifics and raw log

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
func (it *GovernanceUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceUnstaked)
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
		it.Event = new(GovernanceUnstaked)
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
func (it *GovernanceUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceUnstaked represents a Unstaked event raised by the Governance contract.
type GovernanceUnstaked struct {
	NodeAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) FilterUnstaked(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceUnstakedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Unstaked", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceUnstakedIterator{contract: _Governance.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *GovernanceUnstaked, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Unstaked", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceUnstaked)
				if err := _Governance.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// GovernanceWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Governance contract.
type GovernanceWithdrawnIterator struct {
	Event *GovernanceWithdrawn // Event containing the contract specifics and raw log

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
func (it *GovernanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceWithdrawn)
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
		it.Event = new(GovernanceWithdrawn)
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
func (it *GovernanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceWithdrawn represents a Withdrawn event raised by the Governance contract.
type GovernanceWithdrawn struct {
	NodeAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) FilterWithdrawn(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceWithdrawnIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Withdrawn", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceWithdrawnIterator{contract: _Governance.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GovernanceWithdrawn, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Withdrawn", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceWithdrawn)
				if err := _Governance.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
