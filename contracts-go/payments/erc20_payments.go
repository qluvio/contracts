// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payments

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"strings"

	c "github.com/eluv-io/contracts/contracts-go/events"

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Map of ABI names to *abi.ABI
// ABI names are constants starting with K_
var ParsedABIS = map[string]*abi.ABI{}

// Map of ABI names to *bind.BoundContract for log parsing only
// ABI names are constants starting with K_
var BoundContracts = map[string]*bind.BoundContract{}

// Map of Unique events names to *EventInfo.
// Unique events names are constants starting with E_
var UniqueEvents = map[string]*EventInfo{}

// Map of Unique events types to *EventInfo
var EventsByType = map[reflect.Type]*EventInfo{}

// Map of Unique events IDs to *EventInfo
var EventsByID = map[common.Hash]*EventInfo{}

// JSON returns a parsed ABI interface and error if it failed.
func JSON(reader io.Reader) (*abi.ABI, error) {
	dec := json.NewDecoder(reader)

	var anAbi abi.ABI
	if err := dec.Decode(&anAbi); err != nil {
		return nil, err
	}

	return &anAbi, nil
}

func parseABI(name string) (*abi.ABI, error) {
	sabi := ABIS[name]
	if sabi == "" {
		return nil, fmt.Errorf("no such ABI %s", name)
	}
	return JSON(strings.NewReader(sabi))
}

func ParsedABI(name string) (*abi.ABI, error) {
	pabi, ok := ParsedABIS[name]
	if ok {
		return pabi, nil
	}
	return parseABI(name)
}

// ERC20PaymentsPaymentID is an auto generated low-level Go binding around an user-defined struct.
type ERC20PaymentsPaymentID struct {
	RefId    [16]byte
	OracleId common.Address
}

func BoundContract(name string) *bind.BoundContract {
	bc, ok := BoundContracts[name]
	if !ok {
		anABI, err := ParsedABI(name)
		if err != nil {
			panic(err)
		}
		bc = bind.NewBoundContract(common.Address{}, *anABI, nil, nil, nil)
		BoundContracts[name] = bc
	}
	return bc
}

// Type names of contract binding
const (
	K_Payments = "Payments"
)

var ABIS = map[string]string{

	K_Payments: PaymentsABI,
}

// Unique events names.
// Unique events are events whose ID and name are unique across contracts.
const (
	E_Canceled  = "Canceled"
	E_Created   = "Created"
	E_Withdrawn = "Withdrawn"
)

type EventInfo = c.EventInfo
type EventType = c.EventType

func init() {
	for name, _ := range ABIS {
		a, err := parseABI(name)
		if err == nil {
			ParsedABIS[name] = a
		}
	}
	var ev *EventInfo

	ev = &EventInfo{
		Name: "Canceled",
		ID:   common.HexToHash("0x25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea264"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*Canceled)(nil)),
				BoundContract: BoundContract(K_Payments),
			},
		},
	}
	UniqueEvents[E_Canceled] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "Created",
		ID:   common.HexToHash("0xf5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e2"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*Created)(nil)),
				BoundContract: BoundContract(K_Payments),
			},
		},
	}
	UniqueEvents[E_Created] = ev
	EventsByType[ev.Types[0].Type] = ev

	ev = &EventInfo{
		Name: "Withdrawn",
		ID:   common.HexToHash("0xe2676bf5a4d68036e2fc017546a033acda64f74e24c77a7aacea919422b891e9"),
		Types: []EventType{
			{
				Type:          reflect.TypeOf((*Withdrawn)(nil)),
				BoundContract: BoundContract(K_Payments),
			},
		},
	}
	UniqueEvents[E_Withdrawn] = ev
	EventsByType[ev.Types[0].Type] = ev

}

// Unique events structs

// Canceled event with ID 0x25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea264
type Canceled struct {
	ContractId [16]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// Created event with ID 0xf5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e2
type Created struct {
	ContractId    [16]byte
	Sender        common.Address
	Receivers     []common.Address
	TokenContract common.Address
	Amounts       []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// Withdrawn event with ID 0xe2676bf5a4d68036e2fc017546a033acda64f74e24c77a7aacea919422b891e9
type Withdrawn struct {
	ContractId [16]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// PaymentsABI is the input ABI used to generate the binding from.
const PaymentsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"contractId\",\"type\":\"bytes16\"}],\"name\":\"Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"contractId\",\"type\":\"bytes16\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes16\",\"name\":\"contractId\",\"type\":\"bytes16\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_refId\",\"type\":\"bytes16\"}],\"name\":\"cancelPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_refId\",\"type\":\"bytes16\"}],\"name\":\"claimPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_receivers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"ref_id\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"oracle_id\",\"type\":\"address\"}],\"internalType\":\"structERC20Payments.PaymentID\",\"name\":\"_paymentId\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"createPayment\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"refId\",\"type\":\"bytes16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"_refId\",\"type\":\"bytes16\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"enumERC20Payments.PaymentState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PaymentsBin is the compiled bytecode used for deploying new contracts.
var PaymentsBin = "0x608060405234801561001057600080fd5b50611487806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063732760a11461005c578063820b71ca1461008d5780639c6173be146100b0578063e2f5f998146100d5578063f2c9f603146100f6575b600080fd5b61006f61006a36600461107b565b610109565b6040516001600160801b031990911681526020015b60405180910390f35b6100a061009b36600461111b565b61050d565b6040519015158152602001610084565b6100c36100be36600461111b565b61089c565b60405161008496959493929190611184565b6100e86100e3366004611253565b610a10565b604051908152602001610084565b6100a061010436600461111b565b610a5c565b6000610118602086018661111b565b61012181610d70565b156101735760405162461bcd60e51b815260206004820152601860248201527f7061796d656e74494420616c726561647920657869737473000000000000000060448201526064015b60405180910390fd5b610180602087018761111b565b915060006101c28585808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250610d97915050565b90506101cd83610d70565b1561021a5760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420616c726561647920657869737473000000000000000000604482015260640161016a565b6040516323b872dd60e01b8152336004820152306024820152604481018290526001600160a01b038716906323b872dd906064016020604051808303816000875af115801561026d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102919190611311565b6102e85760405162461bcd60e51b815260206004820152602260248201527f7472616e7366657246726f6d2073656e64657220746f2074686973206661696c604482015261195960f21b606482015260840161016a565b6040518060c00160405280336001600160a01b031681526020018a8a808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505050908252506001600160a01b0388166020808301919091526040805188830281810184018352898252919093019291899189918291908501908490808284376000920191909152505050908252506020908101906103979060408b01908b01611333565b6001600160a01b03168152602001600090526001600160801b03198416600090815260208181526040909120825181546001600160a01b0319166001600160a01b0390911617815582820151805191926103f992600185019290910190610f5e565b5060408201516002820180546001600160a01b0319166001600160a01b039092169190911790556060820151805161043b916003840191602090910190610fc3565b5060808201516004820180546001600160a01b039092166001600160a01b031983168117825560a0850151926001600160a81b03191617600160a01b8360028111156104895761048961114c565b0217905550506040516104a091508a908a9061134e565b6040518091039020336001600160a01b0316846fffffffffffffffffffffffffffffffff19167ff5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e28989896040516104f99392919061138e565b60405180910390a450509695505050505050565b60008161051981610d70565b6105605760405162461bcd60e51b81526020600482015260186024820152771c185e5b595b9d125108191bd95cc81b9bdd08195e1a5cdd60421b604482015260640161016a565b8260016001600160801b03198216600090815260208190526040902060040154600160a01b900460ff16600281111561059b5761059b61114c565b036105e85760405162461bcd60e51b815260206004820152601f60248201527f776974686472617761626c653a20616c72656164792077697468647261776e00604482015260640161016a565b60026001600160801b03198216600090815260208190526040902060040154600160a01b900460ff1660028111156106225761062261114c565b0361066f5760405162461bcd60e51b815260206004820152601e60248201527f776974686472617761626c653a20616c726561647920726566756e6465640000604482015260640161016a565b6001600160801b0319841660009081526020819052604090206004015484906001600160a01b031633146106f75760405162461bcd60e51b815260206004820152602960248201527f6f6e6c794f7261636c653a206d6573736167652073656e646572206d757374206044820152686265206f7261636c6560b81b606482015260840161016a565b6001600160801b03198516600090815260208190526040812060048101805460ff60a01b1916600160a01b179055905b600182015481101561085a5760028201546001830180546001600160a01b039092169163a9059cbb919084908110610761576107616113db565b6000918252602090912001546003850180546001600160a01b03909216918590811061078f5761078f6113db565b6000918252602090912001546040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af11580156107e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080a9190611311565b6108485760405162461bcd60e51b815260206004820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b604482015260640161016a565b8061085281611407565b915050610727565b506040516001600160801b03198716907fe2676bf5a4d68036e2fc017546a033acda64f74e24c77a7aacea919422b891e990600090a250600195945050505050565b60006060600060606000806108b087610d70565b15156000036109015760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420646f6573206e6f74206578697374000000000000000000604482015260640161016a565b6001600160801b031987166000908152602081815260409182902080546002820154600483015460018401805487518188028101880190985280885294966001600160a01b0394851696919593851694600389019490841693600160a01b900460ff16928791908301828280156109a157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610983575b50505050509450828054806020026020016040519081016040528092919081815260200182805480156109f357602002820191906000526020600020905b8154815260200190600101908083116109df575b505050505092509650965096509650965096505091939550919395565b6000805b8251811015610a5657828181518110610a2f57610a2f6113db565b602002602001015182610a429190611420565b915080610a4e81611407565b915050610a14565b50919050565b600081610a6881610d70565b610aaf5760405162461bcd60e51b81526020600482015260186024820152771c185e5b595b9d125108191bd95cc81b9bdd08195e1a5cdd60421b604482015260640161016a565b6001600160801b0319831660009081526020819052604090206004015483906001600160a01b03163314610b1e5760405162461bcd60e51b8152602060048201526016602482015275726566756e6461626c653a206e6f74206f7261636c6560501b604482015260640161016a565b60026001600160801b03198216600090815260208190526040902060040154600160a01b900460ff166002811115610b5857610b5861114c565b03610ba55760405162461bcd60e51b815260206004820152601c60248201527f726566756e6461626c653a20616c726561647920726566756e64656400000000604482015260640161016a565b60016001600160801b03198216600090815260208190526040902060040154600160a01b900460ff166002811115610bdf57610bdf61114c565b03610c2c5760405162461bcd60e51b815260206004820152601d60248201527f726566756e6461626c653a20616c72656164792077697468647261776e000000604482015260640161016a565b6001600160801b0319841660009081526020818152604080832060048101805460ff60a01b1916600160a11b179055600381018054835181860281018601909452808452919493610cb2939290830182828015610ca857602002820191906000526020600020905b815481526020019060010190808311610c94575b5050505050610a10565b6002830154835460405163a9059cbb60e01b81526001600160a01b03918216600482015260248101849052929350169063a9059cbb906044016020604051808303816000875af1158015610d0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2e9190611311565b506040516001600160801b03198716907f25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea26490600090a250600195945050505050565b6001600160801b0319166000908152602081905260409020546001600160a01b0316151590565b600080835111610ddf5760405162461bcd60e51b81526020600482015260136024820152721b9bc8185b5bdd5b9d1cc81c1c9bdd9a591959606a1b604482015260640161016a565b506000805b8351811015610e8f576000848281518110610e0157610e016113db565b602002602001015111610e565760405162461bcd60e51b815260206004820152601d60248201527f616d6f756e74206d7573742062652067726561746572207468616e2030000000604482015260640161016a565b838181518110610e6857610e686113db565b602002602001015182610e7b9190611420565b915080610e8781611407565b915050610de4565b50604051636eb1769f60e11b815233600482015230602482015281906001600160a01b0384169063dd62ed3e90604401602060405180830381865afa158015610edc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f009190611438565b1015610f585760405162461bcd60e51b815260206004820152602160248201527f616c6c6f77616e6365206d757374206265203e3d2073756d28616d6f756e74736044820152602960f81b606482015260840161016a565b92915050565b828054828255906000526020600020908101928215610fb3579160200282015b82811115610fb357825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610f7e565b50610fbf929150610ffe565b5090565b828054828255906000526020600020908101928215610fb3579160200282015b82811115610fb3578251825591602001919060010190610fe3565b5b80821115610fbf5760008155600101610fff565b60008083601f84011261102557600080fd5b50813567ffffffffffffffff81111561103d57600080fd5b6020830191508360208260051b850101111561105857600080fd5b9250929050565b80356001600160a01b038116811461107657600080fd5b919050565b60008060008060008086880360a081121561109557600080fd5b873567ffffffffffffffff808211156110ad57600080fd5b6110b98b838c01611013565b90995097508791506040601f19840112156110d357600080fd5b60208a0196506110e560608b0161105f565b955060808a01359250808311156110fb57600080fd5b505061110989828a01611013565b979a9699509497509295939492505050565b60006020828403121561112d57600080fd5b81356001600160801b03198116811461114557600080fd5b9392505050565b634e487b7160e01b600052602160045260246000fd5b6003811061118057634e487b7160e01b600052602160045260246000fd5b9052565b6001600160a01b03878116825260c0602080840182905288519184018290526000928982019290919060e0860190855b818110156111d25785518516835294830194918301916001016111b4565b5050918916604086015284820360608601528751808352918101925080880160005b83811015611210578151855293820193908201906001016111f4565b5050506001600160a01b038616608085015250905061123260a0830184611162565b979650505050505050565b634e487b7160e01b600052604160045260246000fd5b6000602080838503121561126657600080fd5b823567ffffffffffffffff8082111561127e57600080fd5b818501915085601f83011261129257600080fd5b8135818111156112a4576112a461123d565b8060051b604051601f19603f830116810181811085821117156112c9576112c961123d565b6040529182528482019250838101850191888311156112e757600080fd5b938501935b82851015611305578435845293850193928501926112ec565b98975050505050505050565b60006020828403121561132357600080fd5b8151801515811461114557600080fd5b60006020828403121561134557600080fd5b6111458261105f565b60008184825b85811015611383576001600160a01b0361136d8361105f565b1683526020928301929190910190600101611354565b509095945050505050565b6001600160a01b0384168152604060208201819052810182905260006001600160fb1b038311156113be57600080fd5b8260051b8085606085013760009201606001918252509392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611419576114196113f1565b5060010190565b60008219821115611433576114336113f1565b500190565b60006020828403121561144a57600080fd5b505191905056fea2646970667358221220b2da5b414cbc65def9095a490826355ec9c7b70b71674e3f2d1ee04b9b2eef6c64736f6c634300080d0033"

// DeployPayments deploys a new Ethereum contract, binding an instance of Payments to it.
func DeployPayments(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Payments, error) {
	parsed, err := ParsedABI(K_Payments)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PaymentsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Payments{PaymentsCaller: PaymentsCaller{contract: contract}, PaymentsTransactor: PaymentsTransactor{contract: contract}, PaymentsFilterer: PaymentsFilterer{contract: contract}}, nil
}

// Payments is an auto generated Go binding around an Ethereum contract.
type Payments struct {
	PaymentsCaller     // Read-only binding to the contract
	PaymentsTransactor // Write-only binding to the contract
	PaymentsFilterer   // Log filterer for contract events
}

// PaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewPayments creates a new instance of Payments, bound to a specific deployed contract.
func NewPayments(address common.Address, backend bind.ContractBackend) (*Payments, error) {
	contract, err := bindPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Payments{PaymentsCaller: PaymentsCaller{contract: contract}, PaymentsTransactor: PaymentsTransactor{contract: contract}, PaymentsFilterer: PaymentsFilterer{contract: contract}}, nil
}

// NewPaymentsCaller creates a new read-only instance of Payments, bound to a specific deployed contract.
func NewPaymentsCaller(address common.Address, caller bind.ContractCaller) (*PaymentsCaller, error) {
	contract, err := bindPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentsCaller{contract: contract}, nil
}

// NewPaymentsTransactor creates a new write-only instance of Payments, bound to a specific deployed contract.
func NewPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentsTransactor, error) {
	contract, err := bindPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentsTransactor{contract: contract}, nil
}

// NewPaymentsFilterer creates a new log filterer instance of Payments, bound to a specific deployed contract.
func NewPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentsFilterer, error) {
	contract, err := bindPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentsFilterer{contract: contract}, nil
}

// bindPayments binds a generic wrapper to an already deployed contract.
func bindPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParsedABI(K_Payments)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// CalculateTotal is a free data retrieval call binding the contract method 0xe2f5f998.
//
// Solidity: function calculateTotal(uint256[] _amounts) pure returns(uint256 total)
func (_Payments *PaymentsCaller) CalculateTotal(opts *bind.CallOpts, _amounts []*big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Payments.contract.Call(opts, out, "calculateTotal", _amounts)
	return *ret0, err
}

// GetContract is a free data retrieval call binding the contract method 0x9c6173be.
//
// Solidity: function getContract(bytes16 _refId) view returns(address sender, address[] receivers, address tokenContract, uint256[] amounts, address oracle, uint8 state)
func (_Payments *PaymentsCaller) GetContract(opts *bind.CallOpts, _refId [16]byte) (struct {
	Sender        common.Address
	Receivers     []common.Address
	TokenContract common.Address
	Amounts       []*big.Int
	Oracle        common.Address
	State         uint8
}, error) {
	ret := new(struct {
		Sender        common.Address
		Receivers     []common.Address
		TokenContract common.Address
		Amounts       []*big.Int
		Oracle        common.Address
		State         uint8
	})
	out := ret
	err := _Payments.contract.Call(opts, out, "getContract", _refId)
	return *ret, err
}

// CancelPayment is a paid mutator transaction binding the contract method 0xf2c9f603.
//
// Solidity: function cancelPayment(bytes16 _refId) returns(bool)
func (_Payments *PaymentsTransactor) CancelPayment(opts *bind.TransactOpts, _refId [16]byte) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "cancelPayment", _refId)
}

// ClaimPayment is a paid mutator transaction binding the contract method 0x820b71ca.
//
// Solidity: function claimPayment(bytes16 _refId) returns(bool)
func (_Payments *PaymentsTransactor) ClaimPayment(opts *bind.TransactOpts, _refId [16]byte) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "claimPayment", _refId)
}

// CreatePayment is a paid mutator transaction binding the contract method 0x732760a1.
//
// Solidity: function createPayment(address[] _receivers, (bytes16,address) _paymentId, address _tokenContract, uint256[] _amounts) returns(bytes16 refId)
func (_Payments *PaymentsTransactor) CreatePayment(opts *bind.TransactOpts, _receivers []common.Address, _paymentId ERC20PaymentsPaymentID, _tokenContract common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "createPayment", _receivers, _paymentId, _tokenContract, _amounts)
}

// PaymentsCanceledIterator is returned from FilterCanceled and is used to iterate over the raw logs and unpacked data for Canceled events raised by the Payments contract.
type PaymentsCanceledIterator struct {
	Event *PaymentsCanceled // Event containing the contract specifics and raw log

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
func (it *PaymentsCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsCanceled)
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
		it.Event = new(PaymentsCanceled)
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
func (it *PaymentsCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsCanceled represents a Canceled event raised by the Payments contract.
type PaymentsCanceled struct {
	ContractId [16]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCanceled is a free log retrieval operation binding the contract event 0x25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea264.
//
// Solidity: event Canceled(bytes16 indexed contractId)
func (_Payments *PaymentsFilterer) FilterCanceled(opts *bind.FilterOpts, contractId [][16]byte) (*PaymentsCanceledIterator, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Canceled", contractIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsCanceledIterator{contract: _Payments.contract, event: "Canceled", logs: logs, sub: sub}, nil
}

// WatchCanceled is a free log subscription operation binding the contract event 0x25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea264.
//
// Solidity: event Canceled(bytes16 indexed contractId)
func (_Payments *PaymentsFilterer) WatchCanceled(opts *bind.WatchOpts, sink chan<- *PaymentsCanceled, contractId [][16]byte) (event.Subscription, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Canceled", contractIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsCanceled)
				if err := _Payments.contract.UnpackLog(event, "Canceled", log); err != nil {
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

// ParseCanceled is a log parse operation binding the contract event 0x25ccef20460325159f2f1853d4e2f53d50e508519fd950a1b57ca5f4f92ea264.
//
// Solidity: event Canceled(bytes16 indexed contractId)
func (_Payments *PaymentsFilterer) ParseCanceled(log types.Log) (*PaymentsCanceled, error) {
	event := new(PaymentsCanceled)
	if err := _Payments.contract.UnpackLog(event, "Canceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentsCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the Payments contract.
type PaymentsCreatedIterator struct {
	Event *PaymentsCreated // Event containing the contract specifics and raw log

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
func (it *PaymentsCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsCreated)
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
		it.Event = new(PaymentsCreated)
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
func (it *PaymentsCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsCreated represents a Created event raised by the Payments contract.
type PaymentsCreated struct {
	ContractId    [16]byte
	Sender        common.Address
	Receivers     []common.Address
	TokenContract common.Address
	Amounts       []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0xf5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e2.
//
// Solidity: event Created(bytes16 indexed contractId, address indexed sender, address[] indexed receivers, address tokenContract, uint256[] amounts)
func (_Payments *PaymentsFilterer) FilterCreated(opts *bind.FilterOpts, contractId [][16]byte, sender []common.Address, receivers [][]common.Address) (*PaymentsCreatedIterator, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiversRule []interface{}
	for _, receiversItem := range receivers {
		receiversRule = append(receiversRule, receiversItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Created", contractIdRule, senderRule, receiversRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsCreatedIterator{contract: _Payments.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0xf5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e2.
//
// Solidity: event Created(bytes16 indexed contractId, address indexed sender, address[] indexed receivers, address tokenContract, uint256[] amounts)
func (_Payments *PaymentsFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *PaymentsCreated, contractId [][16]byte, sender []common.Address, receivers [][]common.Address) (event.Subscription, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiversRule []interface{}
	for _, receiversItem := range receivers {
		receiversRule = append(receiversRule, receiversItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Created", contractIdRule, senderRule, receiversRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsCreated)
				if err := _Payments.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0xf5b628ae04a2e801c4184fb9b99c711c49643bbf8e1c6a0a70c10e3f5c32c0e2.
//
// Solidity: event Created(bytes16 indexed contractId, address indexed sender, address[] indexed receivers, address tokenContract, uint256[] amounts)
func (_Payments *PaymentsFilterer) ParseCreated(log types.Log) (*PaymentsCreated, error) {
	event := new(PaymentsCreated)
	if err := _Payments.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Payments contract.
type PaymentsWithdrawnIterator struct {
	Event *PaymentsWithdrawn // Event containing the contract specifics and raw log

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
func (it *PaymentsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsWithdrawn)
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
		it.Event = new(PaymentsWithdrawn)
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
func (it *PaymentsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsWithdrawn represents a Withdrawn event raised by the Payments contract.
type PaymentsWithdrawn struct {
	ContractId [16]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xe2676bf5a4d68036e2fc017546a033acda64f74e24c77a7aacea919422b891e9.
//
// Solidity: event Withdrawn(bytes16 indexed contractId)
func (_Payments *PaymentsFilterer) FilterWithdrawn(opts *bind.FilterOpts, contractId [][16]byte) (*PaymentsWithdrawnIterator, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Withdrawn", contractIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsWithdrawnIterator{contract: _Payments.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xe2676bf5a4d68036e2fc017546a033acda64f74e24c77a7aacea919422b891e9.
//
// Solidity: event Withdrawn(bytes16 indexed contractId)
func (_Payments *PaymentsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *PaymentsWithdrawn, contractId [][16]byte) (event.Subscription, error) {

	var contractIdRule []interface{}
	for _, contractIdItem := range contractId {
		contractIdRule = append(contractIdRule, contractIdItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Withdrawn", contractIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsWithdrawn)
				if err := _Payments.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xe2676bf5a4d68036e2fc017546a033acda64f74e24c77a7aacea919422b891e9.
//
// Solidity: event Withdrawn(bytes16 indexed contractId)
func (_Payments *PaymentsFilterer) ParseWithdrawn(log types.Log) (*PaymentsWithdrawn, error) {
	event := new(PaymentsWithdrawn)
	if err := _Payments.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}