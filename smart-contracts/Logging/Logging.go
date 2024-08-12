// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Logging

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

// LoggingContractLogEntry is an auto generated low-level Go binding around an user-defined struct.
type LoggingContractLogEntry struct {
	User   common.Address
	Action string
	Data   string
}

// LoggingMetaData contains all meta data concerning the Logging contract.
var LoggingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"ActionLogged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLogs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structLoggingContract.LogEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"logAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"logs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50610cb88061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806397f8abd814610043578063e79899bd1461005f578063fe290ccb14610091575b5f80fd5b61005d60048036038101906100589190610648565b6100af565b005b610079600480360381019061007491906106f1565b6101cd565b604051610088939291906107bb565b60405180910390f35b61009961032c565b6040516100a69190610964565b60405180910390f35b5f60405180606001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816101609190610b7e565b5060408201518160020190816101769190610b7e565b5050503373ffffffffffffffffffffffffffffffffffffffff167fb7422199dd1457a590e77b44cc3cc853f37225f22ecdc4e0ae50fde56047a5de83836040516101c1929190610c4d565b60405180910390a25050565b5f81815481106101db575f80fd5b905f5260205f2090600302015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461021f906109b1565b80601f016020809104026020016040519081016040528092919081815260200182805461024b906109b1565b80156102965780601f1061026d57610100808354040283529160200191610296565b820191905f5260205f20905b81548152906001019060200180831161027957829003601f168201915b5050505050908060020180546102ab906109b1565b80601f01602080910402602001604051908101604052809291908181526020018280546102d7906109b1565b80156103225780601f106102f957610100808354040283529160200191610322565b820191905f5260205f20905b81548152906001019060200180831161030557829003601f168201915b5050505050905083565b60605f805480602002602001604051908101604052809291908181526020015f905b828210156104f2578382905f5260205f2090600302016040518060600160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820180546103d3906109b1565b80601f01602080910402602001604051908101604052809291908181526020018280546103ff906109b1565b801561044a5780601f106104215761010080835404028352916020019161044a565b820191905f5260205f20905b81548152906001019060200180831161042d57829003601f168201915b50505050508152602001600282018054610463906109b1565b80601f016020809104026020016040519081016040528092919081815260200182805461048f906109b1565b80156104da5780601f106104b1576101008083540402835291602001916104da565b820191905f5260205f20905b8154815290600101906020018083116104bd57829003601f168201915b5050505050815250508152602001906001019061034e565b50505050905090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61055a82610514565b810181811067ffffffffffffffff8211171561057957610578610524565b5b80604052505050565b5f61058b6104fb565b90506105978282610551565b919050565b5f67ffffffffffffffff8211156105b6576105b5610524565b5b6105bf82610514565b9050602081019050919050565b828183375f83830152505050565b5f6105ec6105e78461059c565b610582565b90508281526020810184848401111561060857610607610510565b5b6106138482856105cc565b509392505050565b5f82601f83011261062f5761062e61050c565b5b813561063f8482602086016105da565b91505092915050565b5f806040838503121561065e5761065d610504565b5b5f83013567ffffffffffffffff81111561067b5761067a610508565b5b6106878582860161061b565b925050602083013567ffffffffffffffff8111156106a8576106a7610508565b5b6106b48582860161061b565b9150509250929050565b5f819050919050565b6106d0816106be565b81146106da575f80fd5b50565b5f813590506106eb816106c7565b92915050565b5f6020828403121561070657610705610504565b5b5f610713848285016106dd565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6107458261071c565b9050919050565b6107558161073b565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61078d8261075b565b6107978185610765565b93506107a7818560208601610775565b6107b081610514565b840191505092915050565b5f6060820190506107ce5f83018661074c565b81810360208301526107e08185610783565b905081810360408301526107f48184610783565b9050949350505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6108308161073b565b82525050565b5f82825260208201905092915050565b5f6108508261075b565b61085a8185610836565b935061086a818560208601610775565b61087381610514565b840191505092915050565b5f606083015f8301516108935f860182610827565b50602083015184820360208601526108ab8282610846565b915050604083015184820360408601526108c58282610846565b9150508091505092915050565b5f6108dd838361087e565b905092915050565b5f602082019050919050565b5f6108fb826107fe565b6109058185610808565b93508360208202850161091785610818565b805f5b85811015610952578484038952815161093385826108d2565b945061093e836108e5565b925060208a0199505060018101905061091a565b50829750879550505050505092915050565b5f6020820190508181035f83015261097c81846108f1565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806109c857607f821691505b6020821081036109db576109da610984565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610a3d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a02565b610a478683610a02565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610a82610a7d610a78846106be565b610a5f565b6106be565b9050919050565b5f819050919050565b610a9b83610a68565b610aaf610aa782610a89565b848454610a0e565b825550505050565b5f90565b610ac3610ab7565b610ace818484610a92565b505050565b5b81811015610af157610ae65f82610abb565b600181019050610ad4565b5050565b601f821115610b3657610b07816109e1565b610b10846109f3565b81016020851015610b1f578190505b610b33610b2b856109f3565b830182610ad3565b50505b505050565b5f82821c905092915050565b5f610b565f1984600802610b3b565b1980831691505092915050565b5f610b6e8383610b47565b9150826002028217905092915050565b610b878261075b565b67ffffffffffffffff811115610ba057610b9f610524565b5b610baa82546109b1565b610bb5828285610af5565b5f60209050601f831160018114610be6575f8415610bd4578287015190505b610bde8582610b63565b865550610c45565b601f198416610bf4866109e1565b5f5b82811015610c1b57848901518255600182019150602085019450602081019050610bf6565b86831015610c385784890151610c34601f891682610b47565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f830152610c658185610783565b90508181036020830152610c798184610783565b9050939250505056fea26469706673582212209fe9ee22c72d3bc8ff3cd51f82ade0a46c0c6680053dec2e70e4accae743970864736f6c634300081a0033",
}

// LoggingABI is the input ABI used to generate the binding from.
// Deprecated: Use LoggingMetaData.ABI instead.
var LoggingABI = LoggingMetaData.ABI

// LoggingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LoggingMetaData.Bin instead.
var LoggingBin = LoggingMetaData.Bin

// DeployLogging deploys a new Ethereum contract, binding an instance of Logging to it.
func DeployLogging(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Logging, error) {
	parsed, err := LoggingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LoggingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Logging{LoggingCaller: LoggingCaller{contract: contract}, LoggingTransactor: LoggingTransactor{contract: contract}, LoggingFilterer: LoggingFilterer{contract: contract}}, nil
}

// Logging is an auto generated Go binding around an Ethereum contract.
type Logging struct {
	LoggingCaller     // Read-only binding to the contract
	LoggingTransactor // Write-only binding to the contract
	LoggingFilterer   // Log filterer for contract events
}

// LoggingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoggingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoggingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoggingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoggingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoggingSession struct {
	Contract     *Logging          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoggingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoggingCallerSession struct {
	Contract *LoggingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LoggingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoggingTransactorSession struct {
	Contract     *LoggingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LoggingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoggingRaw struct {
	Contract *Logging // Generic contract binding to access the raw methods on
}

// LoggingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoggingCallerRaw struct {
	Contract *LoggingCaller // Generic read-only contract binding to access the raw methods on
}

// LoggingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoggingTransactorRaw struct {
	Contract *LoggingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogging creates a new instance of Logging, bound to a specific deployed contract.
func NewLogging(address common.Address, backend bind.ContractBackend) (*Logging, error) {
	contract, err := bindLogging(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Logging{LoggingCaller: LoggingCaller{contract: contract}, LoggingTransactor: LoggingTransactor{contract: contract}, LoggingFilterer: LoggingFilterer{contract: contract}}, nil
}

// NewLoggingCaller creates a new read-only instance of Logging, bound to a specific deployed contract.
func NewLoggingCaller(address common.Address, caller bind.ContractCaller) (*LoggingCaller, error) {
	contract, err := bindLogging(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoggingCaller{contract: contract}, nil
}

// NewLoggingTransactor creates a new write-only instance of Logging, bound to a specific deployed contract.
func NewLoggingTransactor(address common.Address, transactor bind.ContractTransactor) (*LoggingTransactor, error) {
	contract, err := bindLogging(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoggingTransactor{contract: contract}, nil
}

// NewLoggingFilterer creates a new log filterer instance of Logging, bound to a specific deployed contract.
func NewLoggingFilterer(address common.Address, filterer bind.ContractFilterer) (*LoggingFilterer, error) {
	contract, err := bindLogging(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoggingFilterer{contract: contract}, nil
}

// bindLogging binds a generic wrapper to an already deployed contract.
func bindLogging(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoggingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logging *LoggingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logging.Contract.LoggingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logging *LoggingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logging.Contract.LoggingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logging *LoggingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logging.Contract.LoggingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logging *LoggingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logging.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logging *LoggingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logging.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logging *LoggingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logging.Contract.contract.Transact(opts, method, params...)
}

// GetLogs is a free data retrieval call binding the contract method 0xfe290ccb.
//
// Solidity: function getLogs() view returns((address,string,string)[])
func (_Logging *LoggingCaller) GetLogs(opts *bind.CallOpts) ([]LoggingContractLogEntry, error) {
	var out []interface{}
	err := _Logging.contract.Call(opts, &out, "getLogs")

	if err != nil {
		return *new([]LoggingContractLogEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]LoggingContractLogEntry)).(*[]LoggingContractLogEntry)

	return out0, err

}

// GetLogs is a free data retrieval call binding the contract method 0xfe290ccb.
//
// Solidity: function getLogs() view returns((address,string,string)[])
func (_Logging *LoggingSession) GetLogs() ([]LoggingContractLogEntry, error) {
	return _Logging.Contract.GetLogs(&_Logging.CallOpts)
}

// GetLogs is a free data retrieval call binding the contract method 0xfe290ccb.
//
// Solidity: function getLogs() view returns((address,string,string)[])
func (_Logging *LoggingCallerSession) GetLogs() ([]LoggingContractLogEntry, error) {
	return _Logging.Contract.GetLogs(&_Logging.CallOpts)
}

// Logs is a free data retrieval call binding the contract method 0xe79899bd.
//
// Solidity: function logs(uint256 ) view returns(address user, string action, string data)
func (_Logging *LoggingCaller) Logs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User   common.Address
	Action string
	Data   string
}, error) {
	var out []interface{}
	err := _Logging.contract.Call(opts, &out, "logs", arg0)

	outstruct := new(struct {
		User   common.Address
		Action string
		Data   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Action = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Data = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Logs is a free data retrieval call binding the contract method 0xe79899bd.
//
// Solidity: function logs(uint256 ) view returns(address user, string action, string data)
func (_Logging *LoggingSession) Logs(arg0 *big.Int) (struct {
	User   common.Address
	Action string
	Data   string
}, error) {
	return _Logging.Contract.Logs(&_Logging.CallOpts, arg0)
}

// Logs is a free data retrieval call binding the contract method 0xe79899bd.
//
// Solidity: function logs(uint256 ) view returns(address user, string action, string data)
func (_Logging *LoggingCallerSession) Logs(arg0 *big.Int) (struct {
	User   common.Address
	Action string
	Data   string
}, error) {
	return _Logging.Contract.Logs(&_Logging.CallOpts, arg0)
}

// LogAction is a paid mutator transaction binding the contract method 0x97f8abd8.
//
// Solidity: function logAction(string action, string data) returns()
func (_Logging *LoggingTransactor) LogAction(opts *bind.TransactOpts, action string, data string) (*types.Transaction, error) {
	return _Logging.contract.Transact(opts, "logAction", action, data)
}

// LogAction is a paid mutator transaction binding the contract method 0x97f8abd8.
//
// Solidity: function logAction(string action, string data) returns()
func (_Logging *LoggingSession) LogAction(action string, data string) (*types.Transaction, error) {
	return _Logging.Contract.LogAction(&_Logging.TransactOpts, action, data)
}

// LogAction is a paid mutator transaction binding the contract method 0x97f8abd8.
//
// Solidity: function logAction(string action, string data) returns()
func (_Logging *LoggingTransactorSession) LogAction(action string, data string) (*types.Transaction, error) {
	return _Logging.Contract.LogAction(&_Logging.TransactOpts, action, data)
}

// LoggingActionLoggedIterator is returned from FilterActionLogged and is used to iterate over the raw logs and unpacked data for ActionLogged events raised by the Logging contract.
type LoggingActionLoggedIterator struct {
	Event *LoggingActionLogged // Event containing the contract specifics and raw log

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
func (it *LoggingActionLoggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoggingActionLogged)
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
		it.Event = new(LoggingActionLogged)
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
func (it *LoggingActionLoggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoggingActionLoggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoggingActionLogged represents a ActionLogged event raised by the Logging contract.
type LoggingActionLogged struct {
	User   common.Address
	Action string
	Data   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterActionLogged is a free log retrieval operation binding the contract event 0xb7422199dd1457a590e77b44cc3cc853f37225f22ecdc4e0ae50fde56047a5de.
//
// Solidity: event ActionLogged(address indexed user, string action, string data)
func (_Logging *LoggingFilterer) FilterActionLogged(opts *bind.FilterOpts, user []common.Address) (*LoggingActionLoggedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Logging.contract.FilterLogs(opts, "ActionLogged", userRule)
	if err != nil {
		return nil, err
	}
	return &LoggingActionLoggedIterator{contract: _Logging.contract, event: "ActionLogged", logs: logs, sub: sub}, nil
}

// WatchActionLogged is a free log subscription operation binding the contract event 0xb7422199dd1457a590e77b44cc3cc853f37225f22ecdc4e0ae50fde56047a5de.
//
// Solidity: event ActionLogged(address indexed user, string action, string data)
func (_Logging *LoggingFilterer) WatchActionLogged(opts *bind.WatchOpts, sink chan<- *LoggingActionLogged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Logging.contract.WatchLogs(opts, "ActionLogged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoggingActionLogged)
				if err := _Logging.contract.UnpackLog(event, "ActionLogged", log); err != nil {
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

// ParseActionLogged is a log parse operation binding the contract event 0xb7422199dd1457a590e77b44cc3cc853f37225f22ecdc4e0ae50fde56047a5de.
//
// Solidity: event ActionLogged(address indexed user, string action, string data)
func (_Logging *LoggingFilterer) ParseActionLogged(log types.Log) (*LoggingActionLogged, error) {
	event := new(LoggingActionLogged)
	if err := _Logging.contract.UnpackLog(event, "ActionLogged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
