// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FaucetABI is the input ABI used to generate the binding from.
const FaucetABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"faucetStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"faucetName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"drip2000Token\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"drip5000Token\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"turnFaucetOff\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"drip1000Token\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"turnFaucetOn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_faucetName\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"OneKTokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"TwoKTokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"FiveKTokenSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"FaucetOn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"FaucetOff\",\"type\":\"event\"}]"

// FaucetBin is the compiled bytecode used for deploying new contracts.
const FaucetBin = `0x608060408190527f4f776e61626c6532303139303532383139333830304d4c000000000000000000600055610a3f388190039081908339810160405280516001805432600160a060020a031991821681179092556002805490911690911790550180516100739060049060208401906100c2565b506005805460ff1916600117908190556040805160ff9290921615158252517fd2a3a2d721c242e2bb5894f541c2a05972d28379a5ae20423a5419a42da3bf639181900360200190a15061015d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061010357805160ff1916838001178555610130565b82800160010185558215610130579182015b82811115610130578251825591602001919060010190610115565b5061013c929150610140565b5090565b61015a91905b8082111561013c5760008155600101610146565b90565b6108d38061016c6000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f81146100dc5780630dfec7961461010d57806312065fe014610136578063400207071461015d57806341c0e1b5146101e757806354fd4d50146101fc578063582b8511146102115780636d2e4b1b146102265780638da5cb5b146102475780639fef1ce21461025c578063af570c0414610271578063b412506b14610286578063cbf706fd1461029b578063ee8734d0146102b0578063f2fde38b146102c5575b005b3480156100e857600080fd5b506100f16102e6565b60408051600160a060020a039092168252519081900360200190f35b34801561011957600080fd5b506101226102f5565b604080519115158252519081900360200190f35b34801561014257600080fd5b5061014b6102fe565b60408051918252519081900360200190f35b34801561016957600080fd5b50610172610303565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ac578181015183820152602001610194565b50505050905090810190601f1680156101d95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101f357600080fd5b506100da610391565b34801561020857600080fd5b5061014b6103cd565b34801561021d57600080fd5b506100da6103d3565b34801561023257600080fd5b506100da600160a060020a036004351661044a565b34801561025357600080fd5b506100f16104a5565b34801561026857600080fd5b506100da6104b4565b34801561027d57600080fd5b506100f161052d565b34801561029257600080fd5b506100da61053c565b3480156102a757600080fd5b506100da6105bb565b3480156102bc57600080fd5b506100da610633565b3480156102d157600080fd5b506100da600160a060020a03600435166106bb565b600154600160a060020a031681565b60055460ff1681565b303190565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103895780601f1061035e57610100808354040283529160200191610389565b820191906000526020600020905b81548152906001019060200180831161036c57829003601f168201915b505050505081565b600254600160a060020a03163214806103b45750600254600160a060020a031633145b15156103bf57600080fd5b600254600160a060020a0316ff5b60005481565b60055460ff1615156103e457600080fd5b6103ed3361072d565b15156103f857600080fd5b61040a686c6b935b8bbd400000610784565b61041633611c20610886565b6040805133815290517eaacc7878c456f81effdd03a5982712003474ea96fceb9ded19748a0160a60c9181900360200190a1565b600154600160a060020a0316321461046157600080fd5b600160a060020a038116151561047657600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b60055460ff1615156104c557600080fd5b6104ce3361072d565b15156104d957600080fd5b6104ec69010f0cf064dd59200000610784565b6104f833614650610886565b6040805133815290517f0cecbc87a45437e3e1904e0a78549073436f9ce37badc9b8619e01328a0999929181900360200190a1565b600354600160a060020a031681565b600254600160a060020a031632148061055f5750600254600160a060020a031633145b151561056a57600080fd5b60055460ff16151561057b57600080fd5b6005805460ff19169055604080516000815290517f1f2881e73120325370afe6da38205f10340d474ae88deced4da5b739ca650bcf9181900360200190a1565b60055460ff1615156105cc57600080fd5b6105d53361072d565b15156105e057600080fd5b6105f2683635c9adc5dea00000610784565b6105fe33610e10610886565b6040805133815290517fb451b2fb4e1bc44234460580d26a4f71a2693d5c5a72e515b233f2e20235386c9181900360200190a1565b600254600160a060020a03163214806106565750600254600160a060020a031633145b151561066157600080fd5b60055460ff161561067157600080fd5b6005805460ff1916600117908190556040805160ff9290921615158252517fd2a3a2d721c242e2bb5894f541c2a05972d28379a5ae20423a5419a42da3bf639181900360200190a1565b600254600160a060020a03163214806106de5750600254600160a060020a031633145b15156106e957600080fd5b600160a060020a03811615156106fe57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a03811660009081526006602052604081205415156107545750600161077f565b600160a060020a038216600090815260066020526040902054421061077b5750600161077f565b5060005b919050565b303181111561081a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f496e73756666696369656e742062616c616e636520696e20666175636574206660448201527f6f72207769746864726177616c20726571756573740000000000000000000000606482015290519081900360840190fd5b604051339082156108fc029083906000818181858888f19350505050158015610847573d6000803e3d6000fd5b50604080513381526020810183905281517f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65929181900390910190a150565b600160a060020a0390911660009081526006602052604090204290910190555600a165627a7a72305820a4d295f44dc958ee20defe204e46d0c6fab77841c43b581756595d67468c1cbb0029`

// DeployFaucet deploys a new Ethereum contract, binding an instance of Faucet to it.
func DeployFaucet(auth *bind.TransactOpts, backend bind.ContractBackend, _faucetName string) (common.Address, *types.Transaction, *Faucet, error) {
	parsed, err := abi.JSON(strings.NewReader(FaucetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FaucetBin), backend, _faucetName)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Faucet{FaucetCaller: FaucetCaller{contract: contract}, FaucetTransactor: FaucetTransactor{contract: contract}, FaucetFilterer: FaucetFilterer{contract: contract}}, nil
}

// Faucet is an auto generated Go binding around an Ethereum contract.
type Faucet struct {
	FaucetCaller     // Read-only binding to the contract
	FaucetTransactor // Write-only binding to the contract
	FaucetFilterer   // Log filterer for contract events
}

// FaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaucetSession struct {
	Contract     *Faucet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaucetCallerSession struct {
	Contract *FaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaucetTransactorSession struct {
	Contract     *FaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaucetRaw struct {
	Contract *Faucet // Generic contract binding to access the raw methods on
}

// FaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaucetCallerRaw struct {
	Contract *FaucetCaller // Generic read-only contract binding to access the raw methods on
}

// FaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FaucetTransactorRaw struct {
	Contract *FaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaucet creates a new instance of Faucet, bound to a specific deployed contract.
func NewFaucet(address common.Address, backend bind.ContractBackend) (*Faucet, error) {
	contract, err := bindFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Faucet{FaucetCaller: FaucetCaller{contract: contract}, FaucetTransactor: FaucetTransactor{contract: contract}, FaucetFilterer: FaucetFilterer{contract: contract}}, nil
}

// NewFaucetCaller creates a new read-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetCaller(address common.Address, caller bind.ContractCaller) (*FaucetCaller, error) {
	contract, err := bindFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetCaller{contract: contract}, nil
}

// NewFaucetTransactor creates a new write-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*FaucetTransactor, error) {
	contract, err := bindFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetTransactor{contract: contract}, nil
}

// NewFaucetFilterer creates a new log filterer instance of Faucet, bound to a specific deployed contract.
func NewFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*FaucetFilterer, error) {
	contract, err := bindFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaucetFilterer{contract: contract}, nil
}

// bindFaucet binds a generic wrapper to an already deployed contract.
func bindFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FaucetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.FaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transact(opts, method, params...)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_Faucet *FaucetCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Faucet.contract.Call(opts, out, "contentSpace")
	return *ret0, err
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_Faucet *FaucetSession) ContentSpace() (common.Address, error) {
	return _Faucet.Contract.ContentSpace(&_Faucet.CallOpts)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_Faucet *FaucetCallerSession) ContentSpace() (common.Address, error) {
	return _Faucet.Contract.ContentSpace(&_Faucet.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Faucet *FaucetCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Faucet.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Faucet *FaucetSession) Creator() (common.Address, error) {
	return _Faucet.Contract.Creator(&_Faucet.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Faucet *FaucetCallerSession) Creator() (common.Address, error) {
	return _Faucet.Contract.Creator(&_Faucet.CallOpts)
}

// FaucetName is a free data retrieval call binding the contract method 0x40020707.
//
// Solidity: function faucetName() constant returns(string)
func (_Faucet *FaucetCaller) FaucetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Faucet.contract.Call(opts, out, "faucetName")
	return *ret0, err
}

// FaucetName is a free data retrieval call binding the contract method 0x40020707.
//
// Solidity: function faucetName() constant returns(string)
func (_Faucet *FaucetSession) FaucetName() (string, error) {
	return _Faucet.Contract.FaucetName(&_Faucet.CallOpts)
}

// FaucetName is a free data retrieval call binding the contract method 0x40020707.
//
// Solidity: function faucetName() constant returns(string)
func (_Faucet *FaucetCallerSession) FaucetName() (string, error) {
	return _Faucet.Contract.FaucetName(&_Faucet.CallOpts)
}

// FaucetStatus is a free data retrieval call binding the contract method 0x0dfec796.
//
// Solidity: function faucetStatus() constant returns(bool)
func (_Faucet *FaucetCaller) FaucetStatus(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Faucet.contract.Call(opts, out, "faucetStatus")
	return *ret0, err
}

// FaucetStatus is a free data retrieval call binding the contract method 0x0dfec796.
//
// Solidity: function faucetStatus() constant returns(bool)
func (_Faucet *FaucetSession) FaucetStatus() (bool, error) {
	return _Faucet.Contract.FaucetStatus(&_Faucet.CallOpts)
}

// FaucetStatus is a free data retrieval call binding the contract method 0x0dfec796.
//
// Solidity: function faucetStatus() constant returns(bool)
func (_Faucet *FaucetCallerSession) FaucetStatus() (bool, error) {
	return _Faucet.Contract.FaucetStatus(&_Faucet.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Faucet *FaucetCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Faucet.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Faucet *FaucetSession) GetBalance() (*big.Int, error) {
	return _Faucet.Contract.GetBalance(&_Faucet.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Faucet *FaucetCallerSession) GetBalance() (*big.Int, error) {
	return _Faucet.Contract.GetBalance(&_Faucet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Faucet *FaucetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Faucet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Faucet *FaucetSession) Owner() (common.Address, error) {
	return _Faucet.Contract.Owner(&_Faucet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Faucet *FaucetCallerSession) Owner() (common.Address, error) {
	return _Faucet.Contract.Owner(&_Faucet.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Faucet *FaucetCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Faucet.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Faucet *FaucetSession) Version() ([32]byte, error) {
	return _Faucet.Contract.Version(&_Faucet.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Faucet *FaucetCallerSession) Version() ([32]byte, error) {
	return _Faucet.Contract.Version(&_Faucet.CallOpts)
}

// Drip1000Token is a paid mutator transaction binding the contract method 0xcbf706fd.
//
// Solidity: function drip1000Token() returns()
func (_Faucet *FaucetTransactor) Drip1000Token(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "drip1000Token")
}

// Drip1000Token is a paid mutator transaction binding the contract method 0xcbf706fd.
//
// Solidity: function drip1000Token() returns()
func (_Faucet *FaucetSession) Drip1000Token() (*types.Transaction, error) {
	return _Faucet.Contract.Drip1000Token(&_Faucet.TransactOpts)
}

// Drip1000Token is a paid mutator transaction binding the contract method 0xcbf706fd.
//
// Solidity: function drip1000Token() returns()
func (_Faucet *FaucetTransactorSession) Drip1000Token() (*types.Transaction, error) {
	return _Faucet.Contract.Drip1000Token(&_Faucet.TransactOpts)
}

// Drip2000Token is a paid mutator transaction binding the contract method 0x582b8511.
//
// Solidity: function drip2000Token() returns()
func (_Faucet *FaucetTransactor) Drip2000Token(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "drip2000Token")
}

// Drip2000Token is a paid mutator transaction binding the contract method 0x582b8511.
//
// Solidity: function drip2000Token() returns()
func (_Faucet *FaucetSession) Drip2000Token() (*types.Transaction, error) {
	return _Faucet.Contract.Drip2000Token(&_Faucet.TransactOpts)
}

// Drip2000Token is a paid mutator transaction binding the contract method 0x582b8511.
//
// Solidity: function drip2000Token() returns()
func (_Faucet *FaucetTransactorSession) Drip2000Token() (*types.Transaction, error) {
	return _Faucet.Contract.Drip2000Token(&_Faucet.TransactOpts)
}

// Drip5000Token is a paid mutator transaction binding the contract method 0x9fef1ce2.
//
// Solidity: function drip5000Token() returns()
func (_Faucet *FaucetTransactor) Drip5000Token(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "drip5000Token")
}

// Drip5000Token is a paid mutator transaction binding the contract method 0x9fef1ce2.
//
// Solidity: function drip5000Token() returns()
func (_Faucet *FaucetSession) Drip5000Token() (*types.Transaction, error) {
	return _Faucet.Contract.Drip5000Token(&_Faucet.TransactOpts)
}

// Drip5000Token is a paid mutator transaction binding the contract method 0x9fef1ce2.
//
// Solidity: function drip5000Token() returns()
func (_Faucet *FaucetTransactorSession) Drip5000Token() (*types.Transaction, error) {
	return _Faucet.Contract.Drip5000Token(&_Faucet.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Faucet *FaucetTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Faucet *FaucetSession) Kill() (*types.Transaction, error) {
	return _Faucet.Contract.Kill(&_Faucet.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Faucet *FaucetTransactorSession) Kill() (*types.Transaction, error) {
	return _Faucet.Contract.Kill(&_Faucet.TransactOpts)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Faucet *FaucetTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Faucet *FaucetSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.TransferCreatorship(&_Faucet.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Faucet *FaucetTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.TransferCreatorship(&_Faucet.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Faucet *FaucetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Faucet *FaucetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.TransferOwnership(&_Faucet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Faucet *FaucetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.TransferOwnership(&_Faucet.TransactOpts, newOwner)
}

// TurnFaucetOff is a paid mutator transaction binding the contract method 0xb412506b.
//
// Solidity: function turnFaucetOff() returns()
func (_Faucet *FaucetTransactor) TurnFaucetOff(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "turnFaucetOff")
}

// TurnFaucetOff is a paid mutator transaction binding the contract method 0xb412506b.
//
// Solidity: function turnFaucetOff() returns()
func (_Faucet *FaucetSession) TurnFaucetOff() (*types.Transaction, error) {
	return _Faucet.Contract.TurnFaucetOff(&_Faucet.TransactOpts)
}

// TurnFaucetOff is a paid mutator transaction binding the contract method 0xb412506b.
//
// Solidity: function turnFaucetOff() returns()
func (_Faucet *FaucetTransactorSession) TurnFaucetOff() (*types.Transaction, error) {
	return _Faucet.Contract.TurnFaucetOff(&_Faucet.TransactOpts)
}

// TurnFaucetOn is a paid mutator transaction binding the contract method 0xee8734d0.
//
// Solidity: function turnFaucetOn() returns()
func (_Faucet *FaucetTransactor) TurnFaucetOn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "turnFaucetOn")
}

// TurnFaucetOn is a paid mutator transaction binding the contract method 0xee8734d0.
//
// Solidity: function turnFaucetOn() returns()
func (_Faucet *FaucetSession) TurnFaucetOn() (*types.Transaction, error) {
	return _Faucet.Contract.TurnFaucetOn(&_Faucet.TransactOpts)
}

// TurnFaucetOn is a paid mutator transaction binding the contract method 0xee8734d0.
//
// Solidity: function turnFaucetOn() returns()
func (_Faucet *FaucetTransactorSession) TurnFaucetOn() (*types.Transaction, error) {
	return _Faucet.Contract.TurnFaucetOn(&_Faucet.TransactOpts)
}

// FaucetFaucetOffIterator is returned from FilterFaucetOff and is used to iterate over the raw logs and unpacked data for FaucetOff events raised by the Faucet contract.
type FaucetFaucetOffIterator struct {
	Event *FaucetFaucetOff // Event containing the contract specifics and raw log

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
func (it *FaucetFaucetOffIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetFaucetOff)
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
		it.Event = new(FaucetFaucetOff)
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
func (it *FaucetFaucetOffIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetFaucetOffIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetFaucetOff represents a FaucetOff event raised by the Faucet contract.
type FaucetFaucetOff struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFaucetOff is a free log retrieval operation binding the contract event 0x1f2881e73120325370afe6da38205f10340d474ae88deced4da5b739ca650bcf.
//
// Solidity: event FaucetOff(bool status)
func (_Faucet *FaucetFilterer) FilterFaucetOff(opts *bind.FilterOpts) (*FaucetFaucetOffIterator, error) {

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "FaucetOff")
	if err != nil {
		return nil, err
	}
	return &FaucetFaucetOffIterator{contract: _Faucet.contract, event: "FaucetOff", logs: logs, sub: sub}, nil
}

// WatchFaucetOff is a free log subscription operation binding the contract event 0x1f2881e73120325370afe6da38205f10340d474ae88deced4da5b739ca650bcf.
//
// Solidity: event FaucetOff(bool status)
func (_Faucet *FaucetFilterer) WatchFaucetOff(opts *bind.WatchOpts, sink chan<- *FaucetFaucetOff) (event.Subscription, error) {

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "FaucetOff")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetFaucetOff)
				if err := _Faucet.contract.UnpackLog(event, "FaucetOff", log); err != nil {
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

// FaucetFaucetOnIterator is returned from FilterFaucetOn and is used to iterate over the raw logs and unpacked data for FaucetOn events raised by the Faucet contract.
type FaucetFaucetOnIterator struct {
	Event *FaucetFaucetOn // Event containing the contract specifics and raw log

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
func (it *FaucetFaucetOnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetFaucetOn)
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
		it.Event = new(FaucetFaucetOn)
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
func (it *FaucetFaucetOnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetFaucetOnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetFaucetOn represents a FaucetOn event raised by the Faucet contract.
type FaucetFaucetOn struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFaucetOn is a free log retrieval operation binding the contract event 0xd2a3a2d721c242e2bb5894f541c2a05972d28379a5ae20423a5419a42da3bf63.
//
// Solidity: event FaucetOn(bool status)
func (_Faucet *FaucetFilterer) FilterFaucetOn(opts *bind.FilterOpts) (*FaucetFaucetOnIterator, error) {

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "FaucetOn")
	if err != nil {
		return nil, err
	}
	return &FaucetFaucetOnIterator{contract: _Faucet.contract, event: "FaucetOn", logs: logs, sub: sub}, nil
}

// WatchFaucetOn is a free log subscription operation binding the contract event 0xd2a3a2d721c242e2bb5894f541c2a05972d28379a5ae20423a5419a42da3bf63.
//
// Solidity: event FaucetOn(bool status)
func (_Faucet *FaucetFilterer) WatchFaucetOn(opts *bind.WatchOpts, sink chan<- *FaucetFaucetOn) (event.Subscription, error) {

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "FaucetOn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetFaucetOn)
				if err := _Faucet.contract.UnpackLog(event, "FaucetOn", log); err != nil {
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

// FaucetFiveKTokenSentIterator is returned from FilterFiveKTokenSent and is used to iterate over the raw logs and unpacked data for FiveKTokenSent events raised by the Faucet contract.
type FaucetFiveKTokenSentIterator struct {
	Event *FaucetFiveKTokenSent // Event containing the contract specifics and raw log

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
func (it *FaucetFiveKTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetFiveKTokenSent)
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
		it.Event = new(FaucetFiveKTokenSent)
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
func (it *FaucetFiveKTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetFiveKTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetFiveKTokenSent represents a FiveKTokenSent event raised by the Faucet contract.
type FaucetFiveKTokenSent struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFiveKTokenSent is a free log retrieval operation binding the contract event 0x0cecbc87a45437e3e1904e0a78549073436f9ce37badc9b8619e01328a099992.
//
// Solidity: event FiveKTokenSent(address receiver)
func (_Faucet *FaucetFilterer) FilterFiveKTokenSent(opts *bind.FilterOpts) (*FaucetFiveKTokenSentIterator, error) {

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "FiveKTokenSent")
	if err != nil {
		return nil, err
	}
	return &FaucetFiveKTokenSentIterator{contract: _Faucet.contract, event: "FiveKTokenSent", logs: logs, sub: sub}, nil
}

// WatchFiveKTokenSent is a free log subscription operation binding the contract event 0x0cecbc87a45437e3e1904e0a78549073436f9ce37badc9b8619e01328a099992.
//
// Solidity: event FiveKTokenSent(address receiver)
func (_Faucet *FaucetFilterer) WatchFiveKTokenSent(opts *bind.WatchOpts, sink chan<- *FaucetFiveKTokenSent) (event.Subscription, error) {

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "FiveKTokenSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetFiveKTokenSent)
				if err := _Faucet.contract.UnpackLog(event, "FiveKTokenSent", log); err != nil {
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

// FaucetOneKTokenSentIterator is returned from FilterOneKTokenSent and is used to iterate over the raw logs and unpacked data for OneKTokenSent events raised by the Faucet contract.
type FaucetOneKTokenSentIterator struct {
	Event *FaucetOneKTokenSent // Event containing the contract specifics and raw log

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
func (it *FaucetOneKTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetOneKTokenSent)
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
		it.Event = new(FaucetOneKTokenSent)
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
func (it *FaucetOneKTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetOneKTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetOneKTokenSent represents a OneKTokenSent event raised by the Faucet contract.
type FaucetOneKTokenSent struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOneKTokenSent is a free log retrieval operation binding the contract event 0xb451b2fb4e1bc44234460580d26a4f71a2693d5c5a72e515b233f2e20235386c.
//
// Solidity: event OneKTokenSent(address receiver)
func (_Faucet *FaucetFilterer) FilterOneKTokenSent(opts *bind.FilterOpts) (*FaucetOneKTokenSentIterator, error) {

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "OneKTokenSent")
	if err != nil {
		return nil, err
	}
	return &FaucetOneKTokenSentIterator{contract: _Faucet.contract, event: "OneKTokenSent", logs: logs, sub: sub}, nil
}

// WatchOneKTokenSent is a free log subscription operation binding the contract event 0xb451b2fb4e1bc44234460580d26a4f71a2693d5c5a72e515b233f2e20235386c.
//
// Solidity: event OneKTokenSent(address receiver)
func (_Faucet *FaucetFilterer) WatchOneKTokenSent(opts *bind.WatchOpts, sink chan<- *FaucetOneKTokenSent) (event.Subscription, error) {

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "OneKTokenSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetOneKTokenSent)
				if err := _Faucet.contract.UnpackLog(event, "OneKTokenSent", log); err != nil {
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

// FaucetTwoKTokenSentIterator is returned from FilterTwoKTokenSent and is used to iterate over the raw logs and unpacked data for TwoKTokenSent events raised by the Faucet contract.
type FaucetTwoKTokenSentIterator struct {
	Event *FaucetTwoKTokenSent // Event containing the contract specifics and raw log

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
func (it *FaucetTwoKTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetTwoKTokenSent)
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
		it.Event = new(FaucetTwoKTokenSent)
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
func (it *FaucetTwoKTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetTwoKTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetTwoKTokenSent represents a TwoKTokenSent event raised by the Faucet contract.
type FaucetTwoKTokenSent struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTwoKTokenSent is a free log retrieval operation binding the contract event 0x00aacc7878c456f81effdd03a5982712003474ea96fceb9ded19748a0160a60c.
//
// Solidity: event TwoKTokenSent(address receiver)
func (_Faucet *FaucetFilterer) FilterTwoKTokenSent(opts *bind.FilterOpts) (*FaucetTwoKTokenSentIterator, error) {

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "TwoKTokenSent")
	if err != nil {
		return nil, err
	}
	return &FaucetTwoKTokenSentIterator{contract: _Faucet.contract, event: "TwoKTokenSent", logs: logs, sub: sub}, nil
}

// WatchTwoKTokenSent is a free log subscription operation binding the contract event 0x00aacc7878c456f81effdd03a5982712003474ea96fceb9ded19748a0160a60c.
//
// Solidity: event TwoKTokenSent(address receiver)
func (_Faucet *FaucetFilterer) WatchTwoKTokenSent(opts *bind.WatchOpts, sink chan<- *FaucetTwoKTokenSent) (event.Subscription, error) {

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "TwoKTokenSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetTwoKTokenSent)
				if err := _Faucet.contract.UnpackLog(event, "TwoKTokenSent", log); err != nil {
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

// FaucetWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Faucet contract.
type FaucetWithdrawalIterator struct {
	Event *FaucetWithdrawal // Event containing the contract specifics and raw log

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
func (it *FaucetWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaucetWithdrawal)
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
		it.Event = new(FaucetWithdrawal)
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
func (it *FaucetWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaucetWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaucetWithdrawal represents a Withdrawal event raised by the Faucet contract.
type FaucetWithdrawal struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address to, uint256 amount)
func (_Faucet *FaucetFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*FaucetWithdrawalIterator, error) {

	logs, sub, err := _Faucet.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &FaucetWithdrawalIterator{contract: _Faucet.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address to, uint256 amount)
func (_Faucet *FaucetFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *FaucetWithdrawal) (event.Subscription, error) {

	logs, sub, err := _Faucet.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaucetWithdrawal)
				if err := _Faucet.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newCreator\",\"type\":\"address\"}],\"name\":\"transferCreatorship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contentSpace\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x60806040527f4f776e61626c6532303139303532383139333830304d4c00000000000000000060005560018054600160a060020a0319908116329081179092556002805490911690911790556102c58061005a6000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302d05d3f811461008457806341c0e1b5146100b557806354fd4d50146100ca5780636d2e4b1b146100f15780638da5cb5b14610112578063af570c0414610127578063f2fde38b1461013c575b005b34801561009057600080fd5b5061009961015d565b60408051600160a060020a039092168252519081900360200190f35b3480156100c157600080fd5b5061008261016c565b3480156100d657600080fd5b506100df6101a8565b60408051918252519081900360200190f35b3480156100fd57600080fd5b50610082600160a060020a03600435166101ae565b34801561011e57600080fd5b50610099610209565b34801561013357600080fd5b50610099610218565b34801561014857600080fd5b50610082600160a060020a0360043516610227565b600154600160a060020a031681565b600254600160a060020a031632148061018f5750600254600160a060020a031633145b151561019a57600080fd5b600254600160a060020a0316ff5b60005481565b600154600160a060020a031632146101c557600080fd5b600160a060020a03811615156101da57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b600354600160a060020a031681565b600254600160a060020a031632148061024a5750600254600160a060020a031633145b151561025557600080fd5b600160a060020a038116151561026a57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820747d708664be33b1214aeab86aa309e2912acf3fa357d263ffd39f0d0b64fefb0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_Ownable *OwnableCaller) ContentSpace(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "contentSpace")
	return *ret0, err
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_Ownable *OwnableSession) ContentSpace() (common.Address, error) {
	return _Ownable.Contract.ContentSpace(&_Ownable.CallOpts)
}

// ContentSpace is a free data retrieval call binding the contract method 0xaf570c04.
//
// Solidity: function contentSpace() constant returns(address)
func (_Ownable *OwnableCallerSession) ContentSpace() (common.Address, error) {
	return _Ownable.Contract.ContentSpace(&_Ownable.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Ownable *OwnableCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Ownable *OwnableSession) Creator() (common.Address, error) {
	return _Ownable.Contract.Creator(&_Ownable.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Ownable *OwnableCallerSession) Creator() (common.Address, error) {
	return _Ownable.Contract.Creator(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Ownable *OwnableCaller) Version(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Ownable *OwnableSession) Version() ([32]byte, error) {
	return _Ownable.Contract.Version(&_Ownable.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(bytes32)
func (_Ownable *OwnableCallerSession) Version() ([32]byte, error) {
	return _Ownable.Contract.Version(&_Ownable.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ownable *OwnableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ownable *OwnableSession) Kill() (*types.Transaction, error) {
	return _Ownable.Contract.Kill(&_Ownable.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Ownable *OwnableTransactorSession) Kill() (*types.Transaction, error) {
	return _Ownable.Contract.Kill(&_Ownable.TransactOpts)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Ownable *OwnableTransactor) TransferCreatorship(opts *bind.TransactOpts, newCreator common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferCreatorship", newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Ownable *OwnableSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferCreatorship(&_Ownable.TransactOpts, newCreator)
}

// TransferCreatorship is a paid mutator transaction binding the contract method 0x6d2e4b1b.
//
// Solidity: function transferCreatorship(address newCreator) returns()
func (_Ownable *OwnableTransactorSession) TransferCreatorship(newCreator common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferCreatorship(&_Ownable.TransactOpts, newCreator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}
