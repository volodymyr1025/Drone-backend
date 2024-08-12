// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Policy

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

// PolicyContractPolicy is an auto generated low-level Go binding around an user-defined struct.
type PolicyContractPolicy struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Exists    bool
}

// PolicyMetaData contains all meta data concerning the Policy contract.
var PolicyMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"}],\"name\":\"PolicyCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"PolicyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"}],\"name\":\"PolicyUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_endTime\",\"type\":\"string\"}],\"name\":\"createPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deletePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolicies\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structPolicyContract.Policy[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getPolicy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structPolicyContract.Policy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"}],\"name\":\"getPolicyByZone\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structPolicyContract.Policy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"policies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endTime\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_zone\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_startTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_endTime\",\"type\":\"string\"}],\"name\":\"updatePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506118878061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c8063882b413e11610059578063882b413e146101125780639d47b49c1461012e578063d3e894831461015e578063fed27c8f1461019257610086565b80632b07fce31461008a578063394d9578146100ba5780633b04f6f1146100d657806361b8ce8c146100f4575b5f80fd5b6100a4600480360381019061009f9190610d77565b6101ae565b6040516100b19190610ecd565b60405180910390f35b6100d460048036038101906100cf9190611043565b6103af565b005b6100de6104bb565b6040516100eb9190611200565b60405180910390f35b6100fc610772565b604051610109919061122f565b60405180910390f35b61012c60048036038101906101279190611248565b610778565b005b610148600480360381019061014391906112e4565b610885565b6040516101559190610ecd565b60405180910390f35b61017860048036038101906101739190610d77565b610ac9565b604051610189959493929190611375565b60405180910390f35b6101ac60048036038101906101a79190610d77565b610c21565b005b6101b6610d05565b600154821080156101f257505f82815481106101d5576101d46113d4565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b610231576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102289061144b565b60405180910390fd5b5f8281548110610244576102436113d4565b5b905f5260205f2090600502016040518060a00160405290815f82015481526020016001820154815260200160028201805461027e90611496565b80601f01602080910402602001604051908101604052809291908181526020018280546102aa90611496565b80156102f55780601f106102cc576101008083540402835291602001916102f5565b820191905f5260205f20905b8154815290600101906020018083116102d857829003601f168201915b5050505050815260200160038201805461030e90611496565b80601f016020809104026020016040519081016040528092919081815260200182805461033a90611496565b80156103855780601f1061035c57610100808354040283529160200191610385565b820191905f5260205f20905b81548152906001019060200180831161036857829003601f168201915b50505050508152602001600482015f9054906101000a900460ff1615151515815250509050919050565b5f6040518060a00160405280600154815260200185815260200184815260200183815260200160011515815250908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f01556020820151816001015560408201518160020190816104299190611663565b50606082015181600301908161043f9190611663565b506080820151816004015f6101000a81548160ff02191690831515021790555050506001547fe50a5fb0eabf749370e644cf8c82aed294141759342d76bfc73f485fa05b7f4e84848460405161049793929190611732565b60405180910390a260015f8154809291906104b1906117a2565b9190505550505050565b60605f805b5f8054905081101561051c575f81815481106104df576104de6113d4565b5b905f5260205f2090600502016004015f9054906101000a900460ff161561050f57818061050b906117a2565b9250505b80806001019150506104c0565b505f8167ffffffffffffffff81111561053857610537610f1f565b5b60405190808252806020026020018201604052801561057157816020015b61055e610d05565b8152602001906001900390816105565790505b5090505f805b5f80549050811015610768575f8181548110610596576105956113d4565b5b905f5260205f2090600502016004015f9054906101000a900460ff161561075b575f81815481106105ca576105c96113d4565b5b905f5260205f2090600502016040518060a00160405290815f82015481526020016001820154815260200160028201805461060490611496565b80601f016020809104026020016040519081016040528092919081815260200182805461063090611496565b801561067b5780601f106106525761010080835404028352916020019161067b565b820191905f5260205f20905b81548152906001019060200180831161065e57829003601f168201915b5050505050815260200160038201805461069490611496565b80601f01602080910402602001604051908101604052809291908181526020018280546106c090611496565b801561070b5780601f106106e25761010080835404028352916020019161070b565b820191905f5260205f20905b8154815290600101906020018083116106ee57829003601f168201915b50505050508152602001600482015f9054906101000a900460ff161515151581525050838381518110610741576107406113d4565b5b60200260200101819052508180610757906117a2565b9250505b8080600101915050610577565b5081935050505090565b60015481565b600154841080156107b457505f8481548110610797576107966113d4565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b6107f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ea9061144b565b60405180910390fd5b5f808581548110610807576108066113d4565b5b905f5260205f20906005020190508381600101819055508281600201908161082f9190611663565b50818160030190816108419190611663565b50847f89546db7a15619f9555052cd2c156972f7bce0ed4bcd25130a6a5ad2785d5f9185858560405161087693929190611732565b60405180910390a25050505050565b61088d610d05565b5f5b5f80549050811015610a8857825f82815481106108af576108ae6113d4565b5b905f5260205f209060050201600101541480156108f757505f81815481106108da576108d96113d4565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b15610a7b575f818154811061090f5761090e6113d4565b5b905f5260205f2090600502016040518060a00160405290815f82015481526020016001820154815260200160028201805461094990611496565b80601f016020809104026020016040519081016040528092919081815260200182805461097590611496565b80156109c05780601f10610997576101008083540402835291602001916109c0565b820191905f5260205f20905b8154815290600101906020018083116109a357829003601f168201915b505050505081526020016003820180546109d990611496565b80601f0160208091040260200160405190810160405280929190818152602001828054610a0590611496565b8015610a505780601f10610a2757610100808354040283529160200191610a50565b820191905f5260205f20905b815481529060010190602001808311610a3357829003601f168201915b50505050508152602001600482015f9054906101000a900460ff161515151581525050915050610ac4565b808060010191505061088f565b506040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610abb90611833565b60405180910390fd5b919050565b5f8181548110610ad7575f80fd5b905f5260205f2090600502015f91509050805f015490806001015490806002018054610b0290611496565b80601f0160208091040260200160405190810160405280929190818152602001828054610b2e90611496565b8015610b795780601f10610b5057610100808354040283529160200191610b79565b820191905f5260205f20905b815481529060010190602001808311610b5c57829003601f168201915b505050505090806003018054610b8e90611496565b80601f0160208091040260200160405190810160405280929190818152602001828054610bba90611496565b8015610c055780601f10610bdc57610100808354040283529160200191610c05565b820191905f5260205f20905b815481529060010190602001808311610be857829003601f168201915b505050505090806004015f9054906101000a900460ff16905085565b60015481108015610c5d57505f8181548110610c4057610c3f6113d4565b5b905f5260205f2090600502016004015f9054906101000a900460ff165b610c9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c939061144b565b60405180910390fd5b5f808281548110610cb057610caf6113d4565b5b905f5260205f2090600502016004015f6101000a81548160ff021916908315150217905550807f433ab09ab993382af577c4a70f77dcc7cf149b5b52eb40d19b869664bfb5496d60405160405180910390a250565b6040518060a001604052805f81526020015f815260200160608152602001606081526020015f151581525090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610d5681610d44565b8114610d60575f80fd5b50565b5f81359050610d7181610d4d565b92915050565b5f60208284031215610d8c57610d8b610d3c565b5b5f610d9984828501610d63565b91505092915050565b610dab81610d44565b82525050565b5f819050919050565b610dc381610db1565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e0b82610dc9565b610e158185610dd3565b9350610e25818560208601610de3565b610e2e81610df1565b840191505092915050565b5f8115159050919050565b610e4d81610e39565b82525050565b5f60a083015f830151610e685f860182610da2565b506020830151610e7b6020860182610dba565b5060408301518482036040860152610e938282610e01565b91505060608301518482036060860152610ead8282610e01565b9150506080830151610ec26080860182610e44565b508091505092915050565b5f6020820190508181035f830152610ee58184610e53565b905092915050565b610ef681610db1565b8114610f00575f80fd5b50565b5f81359050610f1181610eed565b92915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610f5582610df1565b810181811067ffffffffffffffff82111715610f7457610f73610f1f565b5b80604052505050565b5f610f86610d33565b9050610f928282610f4c565b919050565b5f67ffffffffffffffff821115610fb157610fb0610f1f565b5b610fba82610df1565b9050602081019050919050565b828183375f83830152505050565b5f610fe7610fe284610f97565b610f7d565b90508281526020810184848401111561100357611002610f1b565b5b61100e848285610fc7565b509392505050565b5f82601f83011261102a57611029610f17565b5b813561103a848260208601610fd5565b91505092915050565b5f805f6060848603121561105a57611059610d3c565b5b5f61106786828701610f03565b935050602084013567ffffffffffffffff81111561108857611087610d40565b5b61109486828701611016565b925050604084013567ffffffffffffffff8111156110b5576110b4610d40565b5b6110c186828701611016565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f60a083015f8301516111095f860182610da2565b50602083015161111c6020860182610dba565b50604083015184820360408601526111348282610e01565b9150506060830151848203606086015261114e8282610e01565b91505060808301516111636080860182610e44565b508091505092915050565b5f61117983836110f4565b905092915050565b5f602082019050919050565b5f611197826110cb565b6111a181856110d5565b9350836020820285016111b3856110e5565b805f5b858110156111ee57848403895281516111cf858261116e565b94506111da83611181565b925060208a019950506001810190506111b6565b50829750879550505050505092915050565b5f6020820190508181035f830152611218818461118d565b905092915050565b61122981610d44565b82525050565b5f6020820190506112425f830184611220565b92915050565b5f805f80608085870312156112605761125f610d3c565b5b5f61126d87828801610d63565b945050602061127e87828801610f03565b935050604085013567ffffffffffffffff81111561129f5761129e610d40565b5b6112ab87828801611016565b925050606085013567ffffffffffffffff8111156112cc576112cb610d40565b5b6112d887828801611016565b91505092959194509250565b5f602082840312156112f9576112f8610d3c565b5b5f61130684828501610f03565b91505092915050565b61131881610db1565b82525050565b5f82825260208201905092915050565b5f61133882610dc9565b611342818561131e565b9350611352818560208601610de3565b61135b81610df1565b840191505092915050565b61136f81610e39565b82525050565b5f60a0820190506113885f830188611220565b611395602083018761130f565b81810360408301526113a7818661132e565b905081810360608301526113bb818561132e565b90506113ca6080830184611366565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f506f6c69637920646f6573206e6f7420657869737400000000000000000000005f82015250565b5f61143560158361131e565b915061144082611401565b602082019050919050565b5f6020820190508181035f83015261146281611429565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806114ad57607f821691505b6020821081036114c0576114bf611469565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026115227fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826114e7565b61152c86836114e7565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61156761156261155d84610d44565b611544565b610d44565b9050919050565b5f819050919050565b6115808361154d565b61159461158c8261156e565b8484546114f3565b825550505050565b5f90565b6115a861159c565b6115b3818484611577565b505050565b5b818110156115d6576115cb5f826115a0565b6001810190506115b9565b5050565b601f82111561161b576115ec816114c6565b6115f5846114d8565b81016020851015611604578190505b611618611610856114d8565b8301826115b8565b50505b505050565b5f82821c905092915050565b5f61163b5f1984600802611620565b1980831691505092915050565b5f611653838361162c565b9150826002028217905092915050565b61166c82610dc9565b67ffffffffffffffff81111561168557611684610f1f565b5b61168f8254611496565b61169a8282856115da565b5f60209050601f8311600181146116cb575f84156116b9578287015190505b6116c38582611648565b86555061172a565b601f1984166116d9866114c6565b5f5b82811015611700578489015182556001820191506020850194506020810190506116db565b8683101561171d5784890151611719601f89168261162c565b8355505b6001600288020188555050505b505050505050565b5f6060820190506117455f83018661130f565b8181036020830152611757818561132e565b9050818103604083015261176b818461132e565b9050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6117ac82610d44565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117de576117dd611775565b5b600182019050919050565b7f506f6c696379206e6f7420666f756e64000000000000000000000000000000005f82015250565b5f61181d60108361131e565b9150611828826117e9565b602082019050919050565b5f6020820190508181035f83015261184a81611811565b905091905056fea2646970667358221220cf63286f9d56184341878fe7801950bd4c0aa68ce9926c37897ff275d3b773a164736f6c634300081a0033",
}

// PolicyABI is the input ABI used to generate the binding from.
// Deprecated: Use PolicyMetaData.ABI instead.
var PolicyABI = PolicyMetaData.ABI

// PolicyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolicyMetaData.Bin instead.
var PolicyBin = PolicyMetaData.Bin

// DeployPolicy deploys a new Ethereum contract, binding an instance of Policy to it.
func DeployPolicy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Policy, error) {
	parsed, err := PolicyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolicyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Policy{PolicyCaller: PolicyCaller{contract: contract}, PolicyTransactor: PolicyTransactor{contract: contract}, PolicyFilterer: PolicyFilterer{contract: contract}}, nil
}

// Policy is an auto generated Go binding around an Ethereum contract.
type Policy struct {
	PolicyCaller     // Read-only binding to the contract
	PolicyTransactor // Write-only binding to the contract
	PolicyFilterer   // Log filterer for contract events
}

// PolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolicySession struct {
	Contract     *Policy           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolicyCallerSession struct {
	Contract *PolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolicyTransactorSession struct {
	Contract     *PolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolicyRaw struct {
	Contract *Policy // Generic contract binding to access the raw methods on
}

// PolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolicyCallerRaw struct {
	Contract *PolicyCaller // Generic read-only contract binding to access the raw methods on
}

// PolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolicyTransactorRaw struct {
	Contract *PolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolicy creates a new instance of Policy, bound to a specific deployed contract.
func NewPolicy(address common.Address, backend bind.ContractBackend) (*Policy, error) {
	contract, err := bindPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Policy{PolicyCaller: PolicyCaller{contract: contract}, PolicyTransactor: PolicyTransactor{contract: contract}, PolicyFilterer: PolicyFilterer{contract: contract}}, nil
}

// NewPolicyCaller creates a new read-only instance of Policy, bound to a specific deployed contract.
func NewPolicyCaller(address common.Address, caller bind.ContractCaller) (*PolicyCaller, error) {
	contract, err := bindPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyCaller{contract: contract}, nil
}

// NewPolicyTransactor creates a new write-only instance of Policy, bound to a specific deployed contract.
func NewPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*PolicyTransactor, error) {
	contract, err := bindPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyTransactor{contract: contract}, nil
}

// NewPolicyFilterer creates a new log filterer instance of Policy, bound to a specific deployed contract.
func NewPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*PolicyFilterer, error) {
	contract, err := bindPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolicyFilterer{contract: contract}, nil
}

// bindPolicy binds a generic wrapper to an already deployed contract.
func bindPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolicyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.PolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transact(opts, method, params...)
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns((uint256,int256,string,string,bool)[])
func (_Policy *PolicyCaller) GetPolicies(opts *bind.CallOpts) ([]PolicyContractPolicy, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "getPolicies")

	if err != nil {
		return *new([]PolicyContractPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new([]PolicyContractPolicy)).(*[]PolicyContractPolicy)

	return out0, err

}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns((uint256,int256,string,string,bool)[])
func (_Policy *PolicySession) GetPolicies() ([]PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicies(&_Policy.CallOpts)
}

// GetPolicies is a free data retrieval call binding the contract method 0x3b04f6f1.
//
// Solidity: function getPolicies() view returns((uint256,int256,string,string,bool)[])
func (_Policy *PolicyCallerSession) GetPolicies() ([]PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicies(&_Policy.CallOpts)
}

// GetPolicy is a free data retrieval call binding the contract method 0x2b07fce3.
//
// Solidity: function getPolicy(uint256 _id) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicyCaller) GetPolicy(opts *bind.CallOpts, _id *big.Int) (PolicyContractPolicy, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "getPolicy", _id)

	if err != nil {
		return *new(PolicyContractPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new(PolicyContractPolicy)).(*PolicyContractPolicy)

	return out0, err

}

// GetPolicy is a free data retrieval call binding the contract method 0x2b07fce3.
//
// Solidity: function getPolicy(uint256 _id) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicySession) GetPolicy(_id *big.Int) (PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicy(&_Policy.CallOpts, _id)
}

// GetPolicy is a free data retrieval call binding the contract method 0x2b07fce3.
//
// Solidity: function getPolicy(uint256 _id) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicyCallerSession) GetPolicy(_id *big.Int) (PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicy(&_Policy.CallOpts, _id)
}

// GetPolicyByZone is a free data retrieval call binding the contract method 0x9d47b49c.
//
// Solidity: function getPolicyByZone(int256 _zone) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicyCaller) GetPolicyByZone(opts *bind.CallOpts, _zone *big.Int) (PolicyContractPolicy, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "getPolicyByZone", _zone)

	if err != nil {
		return *new(PolicyContractPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new(PolicyContractPolicy)).(*PolicyContractPolicy)

	return out0, err

}

// GetPolicyByZone is a free data retrieval call binding the contract method 0x9d47b49c.
//
// Solidity: function getPolicyByZone(int256 _zone) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicySession) GetPolicyByZone(_zone *big.Int) (PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicyByZone(&_Policy.CallOpts, _zone)
}

// GetPolicyByZone is a free data retrieval call binding the contract method 0x9d47b49c.
//
// Solidity: function getPolicyByZone(int256 _zone) view returns((uint256,int256,string,string,bool))
func (_Policy *PolicyCallerSession) GetPolicyByZone(_zone *big.Int) (PolicyContractPolicy, error) {
	return _Policy.Contract.GetPolicyByZone(&_Policy.CallOpts, _zone)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Policy *PolicyCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Policy *PolicySession) NextId() (*big.Int, error) {
	return _Policy.Contract.NextId(&_Policy.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Policy *PolicyCallerSession) NextId() (*big.Int, error) {
	return _Policy.Contract.NextId(&_Policy.CallOpts)
}

// Policies is a free data retrieval call binding the contract method 0xd3e89483.
//
// Solidity: function policies(uint256 ) view returns(uint256 id, int256 zone, string startTime, string endTime, bool exists)
func (_Policy *PolicyCaller) Policies(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Exists    bool
}, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "policies", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Zone      *big.Int
		StartTime string
		EndTime   string
		Exists    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Zone = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.EndTime = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Exists = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Policies is a free data retrieval call binding the contract method 0xd3e89483.
//
// Solidity: function policies(uint256 ) view returns(uint256 id, int256 zone, string startTime, string endTime, bool exists)
func (_Policy *PolicySession) Policies(arg0 *big.Int) (struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Exists    bool
}, error) {
	return _Policy.Contract.Policies(&_Policy.CallOpts, arg0)
}

// Policies is a free data retrieval call binding the contract method 0xd3e89483.
//
// Solidity: function policies(uint256 ) view returns(uint256 id, int256 zone, string startTime, string endTime, bool exists)
func (_Policy *PolicyCallerSession) Policies(arg0 *big.Int) (struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Exists    bool
}, error) {
	return _Policy.Contract.Policies(&_Policy.CallOpts, arg0)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x394d9578.
//
// Solidity: function createPolicy(int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicyTransactor) CreatePolicy(opts *bind.TransactOpts, _zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "createPolicy", _zone, _startTime, _endTime)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x394d9578.
//
// Solidity: function createPolicy(int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicySession) CreatePolicy(_zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.Contract.CreatePolicy(&_Policy.TransactOpts, _zone, _startTime, _endTime)
}

// CreatePolicy is a paid mutator transaction binding the contract method 0x394d9578.
//
// Solidity: function createPolicy(int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicyTransactorSession) CreatePolicy(_zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.Contract.CreatePolicy(&_Policy.TransactOpts, _zone, _startTime, _endTime)
}

// DeletePolicy is a paid mutator transaction binding the contract method 0xfed27c8f.
//
// Solidity: function deletePolicy(uint256 _id) returns()
func (_Policy *PolicyTransactor) DeletePolicy(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "deletePolicy", _id)
}

// DeletePolicy is a paid mutator transaction binding the contract method 0xfed27c8f.
//
// Solidity: function deletePolicy(uint256 _id) returns()
func (_Policy *PolicySession) DeletePolicy(_id *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.DeletePolicy(&_Policy.TransactOpts, _id)
}

// DeletePolicy is a paid mutator transaction binding the contract method 0xfed27c8f.
//
// Solidity: function deletePolicy(uint256 _id) returns()
func (_Policy *PolicyTransactorSession) DeletePolicy(_id *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.DeletePolicy(&_Policy.TransactOpts, _id)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0x882b413e.
//
// Solidity: function updatePolicy(uint256 _id, int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicyTransactor) UpdatePolicy(opts *bind.TransactOpts, _id *big.Int, _zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "updatePolicy", _id, _zone, _startTime, _endTime)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0x882b413e.
//
// Solidity: function updatePolicy(uint256 _id, int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicySession) UpdatePolicy(_id *big.Int, _zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.Contract.UpdatePolicy(&_Policy.TransactOpts, _id, _zone, _startTime, _endTime)
}

// UpdatePolicy is a paid mutator transaction binding the contract method 0x882b413e.
//
// Solidity: function updatePolicy(uint256 _id, int256 _zone, string _startTime, string _endTime) returns()
func (_Policy *PolicyTransactorSession) UpdatePolicy(_id *big.Int, _zone *big.Int, _startTime string, _endTime string) (*types.Transaction, error) {
	return _Policy.Contract.UpdatePolicy(&_Policy.TransactOpts, _id, _zone, _startTime, _endTime)
}

// PolicyPolicyCreatedIterator is returned from FilterPolicyCreated and is used to iterate over the raw logs and unpacked data for PolicyCreated events raised by the Policy contract.
type PolicyPolicyCreatedIterator struct {
	Event *PolicyPolicyCreated // Event containing the contract specifics and raw log

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
func (it *PolicyPolicyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyPolicyCreated)
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
		it.Event = new(PolicyPolicyCreated)
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
func (it *PolicyPolicyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyPolicyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyPolicyCreated represents a PolicyCreated event raised by the Policy contract.
type PolicyPolicyCreated struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPolicyCreated is a free log retrieval operation binding the contract event 0xe50a5fb0eabf749370e644cf8c82aed294141759342d76bfc73f485fa05b7f4e.
//
// Solidity: event PolicyCreated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) FilterPolicyCreated(opts *bind.FilterOpts, id []*big.Int) (*PolicyPolicyCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "PolicyCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &PolicyPolicyCreatedIterator{contract: _Policy.contract, event: "PolicyCreated", logs: logs, sub: sub}, nil
}

// WatchPolicyCreated is a free log subscription operation binding the contract event 0xe50a5fb0eabf749370e644cf8c82aed294141759342d76bfc73f485fa05b7f4e.
//
// Solidity: event PolicyCreated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) WatchPolicyCreated(opts *bind.WatchOpts, sink chan<- *PolicyPolicyCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "PolicyCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyPolicyCreated)
				if err := _Policy.contract.UnpackLog(event, "PolicyCreated", log); err != nil {
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

// ParsePolicyCreated is a log parse operation binding the contract event 0xe50a5fb0eabf749370e644cf8c82aed294141759342d76bfc73f485fa05b7f4e.
//
// Solidity: event PolicyCreated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) ParsePolicyCreated(log types.Log) (*PolicyPolicyCreated, error) {
	event := new(PolicyPolicyCreated)
	if err := _Policy.contract.UnpackLog(event, "PolicyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicyPolicyDeletedIterator is returned from FilterPolicyDeleted and is used to iterate over the raw logs and unpacked data for PolicyDeleted events raised by the Policy contract.
type PolicyPolicyDeletedIterator struct {
	Event *PolicyPolicyDeleted // Event containing the contract specifics and raw log

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
func (it *PolicyPolicyDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyPolicyDeleted)
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
		it.Event = new(PolicyPolicyDeleted)
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
func (it *PolicyPolicyDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyPolicyDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyPolicyDeleted represents a PolicyDeleted event raised by the Policy contract.
type PolicyPolicyDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPolicyDeleted is a free log retrieval operation binding the contract event 0x433ab09ab993382af577c4a70f77dcc7cf149b5b52eb40d19b869664bfb5496d.
//
// Solidity: event PolicyDeleted(uint256 indexed id)
func (_Policy *PolicyFilterer) FilterPolicyDeleted(opts *bind.FilterOpts, id []*big.Int) (*PolicyPolicyDeletedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "PolicyDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return &PolicyPolicyDeletedIterator{contract: _Policy.contract, event: "PolicyDeleted", logs: logs, sub: sub}, nil
}

// WatchPolicyDeleted is a free log subscription operation binding the contract event 0x433ab09ab993382af577c4a70f77dcc7cf149b5b52eb40d19b869664bfb5496d.
//
// Solidity: event PolicyDeleted(uint256 indexed id)
func (_Policy *PolicyFilterer) WatchPolicyDeleted(opts *bind.WatchOpts, sink chan<- *PolicyPolicyDeleted, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "PolicyDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyPolicyDeleted)
				if err := _Policy.contract.UnpackLog(event, "PolicyDeleted", log); err != nil {
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

// ParsePolicyDeleted is a log parse operation binding the contract event 0x433ab09ab993382af577c4a70f77dcc7cf149b5b52eb40d19b869664bfb5496d.
//
// Solidity: event PolicyDeleted(uint256 indexed id)
func (_Policy *PolicyFilterer) ParsePolicyDeleted(log types.Log) (*PolicyPolicyDeleted, error) {
	event := new(PolicyPolicyDeleted)
	if err := _Policy.contract.UnpackLog(event, "PolicyDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicyPolicyUpdatedIterator is returned from FilterPolicyUpdated and is used to iterate over the raw logs and unpacked data for PolicyUpdated events raised by the Policy contract.
type PolicyPolicyUpdatedIterator struct {
	Event *PolicyPolicyUpdated // Event containing the contract specifics and raw log

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
func (it *PolicyPolicyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyPolicyUpdated)
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
		it.Event = new(PolicyPolicyUpdated)
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
func (it *PolicyPolicyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyPolicyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyPolicyUpdated represents a PolicyUpdated event raised by the Policy contract.
type PolicyPolicyUpdated struct {
	Id        *big.Int
	Zone      *big.Int
	StartTime string
	EndTime   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPolicyUpdated is a free log retrieval operation binding the contract event 0x89546db7a15619f9555052cd2c156972f7bce0ed4bcd25130a6a5ad2785d5f91.
//
// Solidity: event PolicyUpdated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) FilterPolicyUpdated(opts *bind.FilterOpts, id []*big.Int) (*PolicyPolicyUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "PolicyUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &PolicyPolicyUpdatedIterator{contract: _Policy.contract, event: "PolicyUpdated", logs: logs, sub: sub}, nil
}

// WatchPolicyUpdated is a free log subscription operation binding the contract event 0x89546db7a15619f9555052cd2c156972f7bce0ed4bcd25130a6a5ad2785d5f91.
//
// Solidity: event PolicyUpdated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) WatchPolicyUpdated(opts *bind.WatchOpts, sink chan<- *PolicyPolicyUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "PolicyUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyPolicyUpdated)
				if err := _Policy.contract.UnpackLog(event, "PolicyUpdated", log); err != nil {
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

// ParsePolicyUpdated is a log parse operation binding the contract event 0x89546db7a15619f9555052cd2c156972f7bce0ed4bcd25130a6a5ad2785d5f91.
//
// Solidity: event PolicyUpdated(uint256 indexed id, int256 zone, string startTime, string endTime)
func (_Policy *PolicyFilterer) ParsePolicyUpdated(log types.Log) (*PolicyPolicyUpdated, error) {
	event := new(PolicyPolicyUpdated)
	if err := _Policy.contract.UnpackLog(event, "PolicyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
