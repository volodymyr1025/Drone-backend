// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Attribute

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

// AttributeContractAttribute is an auto generated low-level Go binding around an user-defined struct.
type AttributeContractAttribute struct {
	Id     *big.Int
	Name   string
	Value  []*big.Int
	Exists bool
}

// AttributeMetaData contains all meta data concerning the Attribute contract.
var AttributeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"name\":\"AttributeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AttributeDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"name\":\"AttributeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"attributes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"_value\",\"type\":\"uint256[]\"}],\"name\":\"createAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"deleteAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getAttribute\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structAttributeContract.Attribute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAttributes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structAttributeContract.Attribute[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAttributesByName\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structAttributeContract.Attribute[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"_value\",\"type\":\"uint256[]\"}],\"name\":\"updateAttribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50611ac58061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c8063493fe80f11610059578063493fe80f1461011057806361b8ce8c146101405780637238b41f1461015e578063d05dcc6a1461017a57610086565b806306516a4a1461008a578063152583de146100ba5780631736a62b146100d85780632822800c146100f4575b5f80fd5b6100a4600480360381019061009f9190610fb0565b6101ac565b6040516100b19190611253565b60405180910390f35b6100c2610595565b6040516100cf9190611253565b60405180910390f35b6100f260048036038101906100ed9190611361565b610808565b005b61010e600480360381019061010991906113e9565b610910565b005b61012a6004803603810190610125919061145f565b610a10565b60405161013791906114f1565b60405180910390f35b610148610bcd565b6040516101559190611520565b60405180910390f35b6101786004803603810190610173919061145f565b610bd3565b005b610194600480360381019061018f919061145f565b610cb7565b6040516101a393929190611590565b60405180910390f35b60605f805b5f805490508110156102c7575f81815481106101d0576101cf6115cc565b5b905f5260205f2090600402016003015f9054906101000a900460ff1680156102a657506102a55f8281548110610209576102086115cc565b5b905f5260205f209060040201600101805461022390611626565b80601f016020809104026020016040519081016040528092919081815260200182805461024f90611626565b801561029a5780601f106102715761010080835404028352916020019161029a565b820191905f5260205f20905b81548152906001019060200180831161027d57829003601f168201915b505050505085610d7d565b5b156102ba5781806102b690611683565b9250505b80806001019150506101b1565b505f8167ffffffffffffffff8111156102e3576102e2610e8c565b5b60405190808252806020026020018201604052801561031c57816020015b610309610dd5565b8152602001906001900390816103015790505b5090505f805b5f80549050811015610589575f8181548110610341576103406115cc565b5b905f5260205f2090600402016003015f9054906101000a900460ff16801561041757506104165f828154811061037a576103796115cc565b5b905f5260205f209060040201600101805461039490611626565b80601f01602080910402602001604051908101604052809291908181526020018280546103c090611626565b801561040b5780601f106103e25761010080835404028352916020019161040b565b820191905f5260205f20905b8154815290600101906020018083116103ee57829003601f168201915b505050505087610d7d565b5b1561057c575f818154811061042f5761042e6115cc565b5b905f5260205f2090600402016040518060800160405290815f820154815260200160018201805461045f90611626565b80601f016020809104026020016040519081016040528092919081815260200182805461048b90611626565b80156104d65780601f106104ad576101008083540402835291602001916104d6565b820191905f5260205f20905b8154815290600101906020018083116104b957829003601f168201915b505050505081526020016002820180548060200260200160405190810160405280929190818152602001828054801561052c57602002820191905f5260205f20905b815481526020019060010190808311610518575b50505050508152602001600382015f9054906101000a900460ff161515151581525050838381518110610562576105616115cc565b5b6020026020010181905250818061057890611683565b9250505b8080600101915050610322565b50819350505050919050565b60605f805b5f805490508110156105f6575f81815481106105b9576105b86115cc565b5b905f5260205f2090600402016003015f9054906101000a900460ff16156105e95781806105e590611683565b9250505b808060010191505061059a565b505f8167ffffffffffffffff81111561061257610611610e8c565b5b60405190808252806020026020018201604052801561064b57816020015b610638610dd5565b8152602001906001900390816106305790505b5090505f805b5f805490508110156107fe575f81815481106106705761066f6115cc565b5b905f5260205f2090600402016003015f9054906101000a900460ff16156107f1575f81815481106106a4576106a36115cc565b5b905f5260205f2090600402016040518060800160405290815f82015481526020016001820180546106d490611626565b80601f016020809104026020016040519081016040528092919081815260200182805461070090611626565b801561074b5780601f106107225761010080835404028352916020019161074b565b820191905f5260205f20905b81548152906001019060200180831161072e57829003601f168201915b50505050508152602001600282018054806020026020016040519081016040528092919081815260200182805480156107a157602002820191905f5260205f20905b81548152602001906001019080831161078d575b50505050508152602001600382015f9054906101000a900460ff1615151515815250508383815181106107d7576107d66115cc565b5b602002602001018190525081806107ed90611683565b9250505b8080600101915050610651565b5081935050505090565b6001548310801561084457505f8381548110610827576108266115cc565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b610883576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087a90611714565b60405180910390fd5b5f808481548110610897576108966115cc565b5b905f5260205f2090600402019050828160010190816108b691906118cf565b50818160020190805190602001906108cf929190610dfd565b50837f8f26e9a323213025e76e7c58dbf6338a9c36708b0f933b7e5babc20acca4e8d18484604051610902929190611a0a565b60405180910390a250505050565b5f6040518060800160405280600154815260200184815260200183815260200160011515815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f0155602082015181600101908161097a91906118cf565b506040820151816002019080519060200190610997929190610dfd565b506060820151816003015f6101000a81548160ff02191690831515021790555050506001547f92ff6eabb2f26553137551f9eaac5f9e399114f11d71898144a4e7e71f8689f183836040516109ed929190611a0a565b60405180910390a260015f815480929190610a0790611683565b91905055505050565b610a18610dd5565b60015482108015610a5457505f8281548110610a3757610a366115cc565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b610a93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8a90611714565b60405180910390fd5b5f8281548110610aa657610aa56115cc565b5b905f5260205f2090600402016040518060800160405290815f8201548152602001600182018054610ad690611626565b80601f0160208091040260200160405190810160405280929190818152602001828054610b0290611626565b8015610b4d5780601f10610b2457610100808354040283529160200191610b4d565b820191905f5260205f20905b815481529060010190602001808311610b3057829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020018280548015610ba357602002820191905f5260205f20905b815481526020019060010190808311610b8f575b50505050508152602001600382015f9054906101000a900460ff1615151515815250509050919050565b60015481565b60015481108015610c0f57505f8181548110610bf257610bf16115cc565b5b905f5260205f2090600402016003015f9054906101000a900460ff165b610c4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4590611714565b60405180910390fd5b5f808281548110610c6257610c616115cc565b5b905f5260205f2090600402016003015f6101000a81548160ff021916908315150217905550807ffe96137d7877697f141204379bad0e6617ff7023d7608f9db3be956d7ceef71160405160405180910390a250565b5f8181548110610cc5575f80fd5b905f5260205f2090600402015f91509050805f015490806001018054610cea90611626565b80601f0160208091040260200160405190810160405280929190818152602001828054610d1690611626565b8015610d615780601f10610d3857610100808354040283529160200191610d61565b820191905f5260205f20905b815481529060010190602001808311610d4457829003601f168201915b505050505090806003015f9054906101000a900460ff16905083565b5f81604051602001610d8f9190611a79565b6040516020818303038152906040528051906020012083604051602001610db69190611a79565b6040516020818303038152906040528051906020012014905092915050565b60405180608001604052805f815260200160608152602001606081526020015f151581525090565b828054828255905f5260205f20908101928215610e37579160200282015b82811115610e36578251825591602001919060010190610e1b565b5b509050610e449190610e48565b5090565b5b80821115610e5f575f815f905550600101610e49565b5090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610ec282610e7c565b810181811067ffffffffffffffff82111715610ee157610ee0610e8c565b5b80604052505050565b5f610ef3610e63565b9050610eff8282610eb9565b919050565b5f67ffffffffffffffff821115610f1e57610f1d610e8c565b5b610f2782610e7c565b9050602081019050919050565b828183375f83830152505050565b5f610f54610f4f84610f04565b610eea565b905082815260208101848484011115610f7057610f6f610e78565b5b610f7b848285610f34565b509392505050565b5f82601f830112610f9757610f96610e74565b5b8135610fa7848260208601610f42565b91505092915050565b5f60208284031215610fc557610fc4610e6c565b5b5f82013567ffffffffffffffff811115610fe257610fe1610e70565b5b610fee84828501610f83565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b61103281611020565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61106a82611038565b6110748185611042565b9350611084818560208601611052565b61108d81610e7c565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6110cc8383611029565b60208301905092915050565b5f602082019050919050565b5f6110ee82611098565b6110f881856110a2565b9350611103836110b2565b805f5b8381101561113357815161111a88826110c1565b9750611125836110d8565b925050600181019050611106565b5085935050505092915050565b5f8115159050919050565b61115481611140565b82525050565b5f608083015f83015161116f5f860182611029565b50602083015184820360208601526111878282611060565b915050604083015184820360408601526111a182826110e4565b91505060608301516111b6606086018261114b565b508091505092915050565b5f6111cc838361115a565b905092915050565b5f602082019050919050565b5f6111ea82610ff7565b6111f48185611001565b93508360208202850161120685611011565b805f5b85811015611241578484038952815161122285826111c1565b945061122d836111d4565b925060208a01995050600181019050611209565b50829750879550505050505092915050565b5f6020820190508181035f83015261126b81846111e0565b905092915050565b61127c81611020565b8114611286575f80fd5b50565b5f8135905061129781611273565b92915050565b5f67ffffffffffffffff8211156112b7576112b6610e8c565b5b602082029050602081019050919050565b5f80fd5b5f6112de6112d98461129d565b610eea565b90508083825260208201905060208402830185811115611301576113006112c8565b5b835b8181101561132a57806113168882611289565b845260208401935050602081019050611303565b5050509392505050565b5f82601f83011261134857611347610e74565b5b81356113588482602086016112cc565b91505092915050565b5f805f6060848603121561137857611377610e6c565b5b5f61138586828701611289565b935050602084013567ffffffffffffffff8111156113a6576113a5610e70565b5b6113b286828701610f83565b925050604084013567ffffffffffffffff8111156113d3576113d2610e70565b5b6113df86828701611334565b9150509250925092565b5f80604083850312156113ff576113fe610e6c565b5b5f83013567ffffffffffffffff81111561141c5761141b610e70565b5b61142885828601610f83565b925050602083013567ffffffffffffffff81111561144957611448610e70565b5b61145585828601611334565b9150509250929050565b5f6020828403121561147457611473610e6c565b5b5f61148184828501611289565b91505092915050565b5f608083015f83015161149f5f860182611029565b50602083015184820360208601526114b78282611060565b915050604083015184820360408601526114d182826110e4565b91505060608301516114e6606086018261114b565b508091505092915050565b5f6020820190508181035f830152611509818461148a565b905092915050565b61151a81611020565b82525050565b5f6020820190506115335f830184611511565b92915050565b5f82825260208201905092915050565b5f61155382611038565b61155d8185611539565b935061156d818560208601611052565b61157681610e7c565b840191505092915050565b61158a81611140565b82525050565b5f6060820190506115a35f830186611511565b81810360208301526115b58185611549565b90506115c46040830184611581565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061163d57607f821691505b6020821081036116505761164f6115f9565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61168d82611020565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116bf576116be611656565b5b600182019050919050565b7f41747472696275746520646f6573206e6f7420657869737400000000000000005f82015250565b5f6116fe601883611539565b9150611709826116ca565b602082019050919050565b5f6020820190508181035f83015261172b816116f2565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261178e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611753565b6117988683611753565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6117d36117ce6117c984611020565b6117b0565b611020565b9050919050565b5f819050919050565b6117ec836117b9565b6118006117f8826117da565b84845461175f565b825550505050565b5f90565b611814611808565b61181f8184846117e3565b505050565b5b81811015611842576118375f8261180c565b600181019050611825565b5050565b601f8211156118875761185881611732565b61186184611744565b81016020851015611870578190505b61188461187c85611744565b830182611824565b50505b505050565b5f82821c905092915050565b5f6118a75f198460080261188c565b1980831691505092915050565b5f6118bf8383611898565b9150826002028217905092915050565b6118d882611038565b67ffffffffffffffff8111156118f1576118f0610e8c565b5b6118fb8254611626565b611906828285611846565b5f60209050601f831160018114611937575f8415611925578287015190505b61192f85826118b4565b865550611996565b601f19841661194586611732565b5f5b8281101561196c57848901518255600182019150602085019450602081019050611947565b868310156119895784890151611985601f891682611898565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b5f6119b882611098565b6119c2818561199e565b93506119cd836110b2565b805f5b838110156119fd5781516119e488826110c1565b97506119ef836110d8565b9250506001810190506119d0565b5085935050505092915050565b5f6040820190508181035f830152611a228185611549565b90508181036020830152611a3681846119ae565b90509392505050565b5f81905092915050565b5f611a5382611038565b611a5d8185611a3f565b9350611a6d818560208601611052565b80840191505092915050565b5f611a848284611a49565b91508190509291505056fea2646970667358221220808ca2e471a9298a7e083829238a0306f4eef6f89b99ccc1f4dfe6f0557954b864736f6c634300081a0033",
}

// AttributeABI is the input ABI used to generate the binding from.
// Deprecated: Use AttributeMetaData.ABI instead.
var AttributeABI = AttributeMetaData.ABI

// AttributeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttributeMetaData.Bin instead.
var AttributeBin = AttributeMetaData.Bin

// DeployAttribute deploys a new Ethereum contract, binding an instance of Attribute to it.
func DeployAttribute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Attribute, error) {
	parsed, err := AttributeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttributeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Attribute{AttributeCaller: AttributeCaller{contract: contract}, AttributeTransactor: AttributeTransactor{contract: contract}, AttributeFilterer: AttributeFilterer{contract: contract}}, nil
}

// Attribute is an auto generated Go binding around an Ethereum contract.
type Attribute struct {
	AttributeCaller     // Read-only binding to the contract
	AttributeTransactor // Write-only binding to the contract
	AttributeFilterer   // Log filterer for contract events
}

// AttributeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttributeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttributeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttributeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttributeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttributeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttributeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttributeSession struct {
	Contract     *Attribute        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttributeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttributeCallerSession struct {
	Contract *AttributeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AttributeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttributeTransactorSession struct {
	Contract     *AttributeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AttributeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttributeRaw struct {
	Contract *Attribute // Generic contract binding to access the raw methods on
}

// AttributeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttributeCallerRaw struct {
	Contract *AttributeCaller // Generic read-only contract binding to access the raw methods on
}

// AttributeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttributeTransactorRaw struct {
	Contract *AttributeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttribute creates a new instance of Attribute, bound to a specific deployed contract.
func NewAttribute(address common.Address, backend bind.ContractBackend) (*Attribute, error) {
	contract, err := bindAttribute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Attribute{AttributeCaller: AttributeCaller{contract: contract}, AttributeTransactor: AttributeTransactor{contract: contract}, AttributeFilterer: AttributeFilterer{contract: contract}}, nil
}

// NewAttributeCaller creates a new read-only instance of Attribute, bound to a specific deployed contract.
func NewAttributeCaller(address common.Address, caller bind.ContractCaller) (*AttributeCaller, error) {
	contract, err := bindAttribute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttributeCaller{contract: contract}, nil
}

// NewAttributeTransactor creates a new write-only instance of Attribute, bound to a specific deployed contract.
func NewAttributeTransactor(address common.Address, transactor bind.ContractTransactor) (*AttributeTransactor, error) {
	contract, err := bindAttribute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttributeTransactor{contract: contract}, nil
}

// NewAttributeFilterer creates a new log filterer instance of Attribute, bound to a specific deployed contract.
func NewAttributeFilterer(address common.Address, filterer bind.ContractFilterer) (*AttributeFilterer, error) {
	contract, err := bindAttribute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttributeFilterer{contract: contract}, nil
}

// bindAttribute binds a generic wrapper to an already deployed contract.
func bindAttribute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AttributeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attribute *AttributeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attribute.Contract.AttributeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attribute *AttributeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attribute.Contract.AttributeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attribute *AttributeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attribute.Contract.AttributeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attribute *AttributeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attribute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attribute *AttributeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attribute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attribute *AttributeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attribute.Contract.contract.Transact(opts, method, params...)
}

// Attributes is a free data retrieval call binding the contract method 0xd05dcc6a.
//
// Solidity: function attributes(uint256 ) view returns(uint256 id, string name, bool exists)
func (_Attribute *AttributeCaller) Attributes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id     *big.Int
	Name   string
	Exists bool
}, error) {
	var out []interface{}
	err := _Attribute.contract.Call(opts, &out, "attributes", arg0)

	outstruct := new(struct {
		Id     *big.Int
		Name   string
		Exists bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Exists = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Attributes is a free data retrieval call binding the contract method 0xd05dcc6a.
//
// Solidity: function attributes(uint256 ) view returns(uint256 id, string name, bool exists)
func (_Attribute *AttributeSession) Attributes(arg0 *big.Int) (struct {
	Id     *big.Int
	Name   string
	Exists bool
}, error) {
	return _Attribute.Contract.Attributes(&_Attribute.CallOpts, arg0)
}

// Attributes is a free data retrieval call binding the contract method 0xd05dcc6a.
//
// Solidity: function attributes(uint256 ) view returns(uint256 id, string name, bool exists)
func (_Attribute *AttributeCallerSession) Attributes(arg0 *big.Int) (struct {
	Id     *big.Int
	Name   string
	Exists bool
}, error) {
	return _Attribute.Contract.Attributes(&_Attribute.CallOpts, arg0)
}

// GetAttribute is a free data retrieval call binding the contract method 0x493fe80f.
//
// Solidity: function getAttribute(uint256 _id) view returns((uint256,string,uint256[],bool))
func (_Attribute *AttributeCaller) GetAttribute(opts *bind.CallOpts, _id *big.Int) (AttributeContractAttribute, error) {
	var out []interface{}
	err := _Attribute.contract.Call(opts, &out, "getAttribute", _id)

	if err != nil {
		return *new(AttributeContractAttribute), err
	}

	out0 := *abi.ConvertType(out[0], new(AttributeContractAttribute)).(*AttributeContractAttribute)

	return out0, err

}

// GetAttribute is a free data retrieval call binding the contract method 0x493fe80f.
//
// Solidity: function getAttribute(uint256 _id) view returns((uint256,string,uint256[],bool))
func (_Attribute *AttributeSession) GetAttribute(_id *big.Int) (AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttribute(&_Attribute.CallOpts, _id)
}

// GetAttribute is a free data retrieval call binding the contract method 0x493fe80f.
//
// Solidity: function getAttribute(uint256 _id) view returns((uint256,string,uint256[],bool))
func (_Attribute *AttributeCallerSession) GetAttribute(_id *big.Int) (AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttribute(&_Attribute.CallOpts, _id)
}

// GetAttributes is a free data retrieval call binding the contract method 0x152583de.
//
// Solidity: function getAttributes() view returns((uint256,string,uint256[],bool)[])
func (_Attribute *AttributeCaller) GetAttributes(opts *bind.CallOpts) ([]AttributeContractAttribute, error) {
	var out []interface{}
	err := _Attribute.contract.Call(opts, &out, "getAttributes")

	if err != nil {
		return *new([]AttributeContractAttribute), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttributeContractAttribute)).(*[]AttributeContractAttribute)

	return out0, err

}

// GetAttributes is a free data retrieval call binding the contract method 0x152583de.
//
// Solidity: function getAttributes() view returns((uint256,string,uint256[],bool)[])
func (_Attribute *AttributeSession) GetAttributes() ([]AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttributes(&_Attribute.CallOpts)
}

// GetAttributes is a free data retrieval call binding the contract method 0x152583de.
//
// Solidity: function getAttributes() view returns((uint256,string,uint256[],bool)[])
func (_Attribute *AttributeCallerSession) GetAttributes() ([]AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttributes(&_Attribute.CallOpts)
}

// GetAttributesByName is a free data retrieval call binding the contract method 0x06516a4a.
//
// Solidity: function getAttributesByName(string _name) view returns((uint256,string,uint256[],bool)[])
func (_Attribute *AttributeCaller) GetAttributesByName(opts *bind.CallOpts, _name string) ([]AttributeContractAttribute, error) {
	var out []interface{}
	err := _Attribute.contract.Call(opts, &out, "getAttributesByName", _name)

	if err != nil {
		return *new([]AttributeContractAttribute), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttributeContractAttribute)).(*[]AttributeContractAttribute)

	return out0, err

}

// GetAttributesByName is a free data retrieval call binding the contract method 0x06516a4a.
//
// Solidity: function getAttributesByName(string _name) view returns((uint256,string,uint256[],bool)[])
func (_Attribute *AttributeSession) GetAttributesByName(_name string) ([]AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttributesByName(&_Attribute.CallOpts, _name)
}

// GetAttributesByName is a free data retrieval call binding the contract method 0x06516a4a.
//
// Solidity: function getAttributesByName(string _name) view returns((uint256,string,uint256[],bool)[])
func (_Attribute *AttributeCallerSession) GetAttributesByName(_name string) ([]AttributeContractAttribute, error) {
	return _Attribute.Contract.GetAttributesByName(&_Attribute.CallOpts, _name)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Attribute *AttributeCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attribute.contract.Call(opts, &out, "nextId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Attribute *AttributeSession) NextId() (*big.Int, error) {
	return _Attribute.Contract.NextId(&_Attribute.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() view returns(uint256)
func (_Attribute *AttributeCallerSession) NextId() (*big.Int, error) {
	return _Attribute.Contract.NextId(&_Attribute.CallOpts)
}

// CreateAttribute is a paid mutator transaction binding the contract method 0x2822800c.
//
// Solidity: function createAttribute(string _name, uint256[] _value) returns()
func (_Attribute *AttributeTransactor) CreateAttribute(opts *bind.TransactOpts, _name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.contract.Transact(opts, "createAttribute", _name, _value)
}

// CreateAttribute is a paid mutator transaction binding the contract method 0x2822800c.
//
// Solidity: function createAttribute(string _name, uint256[] _value) returns()
func (_Attribute *AttributeSession) CreateAttribute(_name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.CreateAttribute(&_Attribute.TransactOpts, _name, _value)
}

// CreateAttribute is a paid mutator transaction binding the contract method 0x2822800c.
//
// Solidity: function createAttribute(string _name, uint256[] _value) returns()
func (_Attribute *AttributeTransactorSession) CreateAttribute(_name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.CreateAttribute(&_Attribute.TransactOpts, _name, _value)
}

// DeleteAttribute is a paid mutator transaction binding the contract method 0x7238b41f.
//
// Solidity: function deleteAttribute(uint256 _id) returns()
func (_Attribute *AttributeTransactor) DeleteAttribute(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Attribute.contract.Transact(opts, "deleteAttribute", _id)
}

// DeleteAttribute is a paid mutator transaction binding the contract method 0x7238b41f.
//
// Solidity: function deleteAttribute(uint256 _id) returns()
func (_Attribute *AttributeSession) DeleteAttribute(_id *big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.DeleteAttribute(&_Attribute.TransactOpts, _id)
}

// DeleteAttribute is a paid mutator transaction binding the contract method 0x7238b41f.
//
// Solidity: function deleteAttribute(uint256 _id) returns()
func (_Attribute *AttributeTransactorSession) DeleteAttribute(_id *big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.DeleteAttribute(&_Attribute.TransactOpts, _id)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x1736a62b.
//
// Solidity: function updateAttribute(uint256 _id, string _name, uint256[] _value) returns()
func (_Attribute *AttributeTransactor) UpdateAttribute(opts *bind.TransactOpts, _id *big.Int, _name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.contract.Transact(opts, "updateAttribute", _id, _name, _value)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x1736a62b.
//
// Solidity: function updateAttribute(uint256 _id, string _name, uint256[] _value) returns()
func (_Attribute *AttributeSession) UpdateAttribute(_id *big.Int, _name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.UpdateAttribute(&_Attribute.TransactOpts, _id, _name, _value)
}

// UpdateAttribute is a paid mutator transaction binding the contract method 0x1736a62b.
//
// Solidity: function updateAttribute(uint256 _id, string _name, uint256[] _value) returns()
func (_Attribute *AttributeTransactorSession) UpdateAttribute(_id *big.Int, _name string, _value []*big.Int) (*types.Transaction, error) {
	return _Attribute.Contract.UpdateAttribute(&_Attribute.TransactOpts, _id, _name, _value)
}

// AttributeAttributeCreatedIterator is returned from FilterAttributeCreated and is used to iterate over the raw logs and unpacked data for AttributeCreated events raised by the Attribute contract.
type AttributeAttributeCreatedIterator struct {
	Event *AttributeAttributeCreated // Event containing the contract specifics and raw log

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
func (it *AttributeAttributeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttributeAttributeCreated)
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
		it.Event = new(AttributeAttributeCreated)
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
func (it *AttributeAttributeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttributeAttributeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttributeAttributeCreated represents a AttributeCreated event raised by the Attribute contract.
type AttributeAttributeCreated struct {
	Id    *big.Int
	Name  string
	Value []*big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAttributeCreated is a free log retrieval operation binding the contract event 0x92ff6eabb2f26553137551f9eaac5f9e399114f11d71898144a4e7e71f8689f1.
//
// Solidity: event AttributeCreated(uint256 indexed id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) FilterAttributeCreated(opts *bind.FilterOpts, id []*big.Int) (*AttributeAttributeCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Attribute.contract.FilterLogs(opts, "AttributeCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &AttributeAttributeCreatedIterator{contract: _Attribute.contract, event: "AttributeCreated", logs: logs, sub: sub}, nil
}

// WatchAttributeCreated is a free log subscription operation binding the contract event 0x92ff6eabb2f26553137551f9eaac5f9e399114f11d71898144a4e7e71f8689f1.
//
// Solidity: event AttributeCreated(uint256 indexed id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) WatchAttributeCreated(opts *bind.WatchOpts, sink chan<- *AttributeAttributeCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Attribute.contract.WatchLogs(opts, "AttributeCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttributeAttributeCreated)
				if err := _Attribute.contract.UnpackLog(event, "AttributeCreated", log); err != nil {
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

// ParseAttributeCreated is a log parse operation binding the contract event 0x92ff6eabb2f26553137551f9eaac5f9e399114f11d71898144a4e7e71f8689f1.
//
// Solidity: event AttributeCreated(uint256 indexed id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) ParseAttributeCreated(log types.Log) (*AttributeAttributeCreated, error) {
	event := new(AttributeAttributeCreated)
	if err := _Attribute.contract.UnpackLog(event, "AttributeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttributeAttributeDeletedIterator is returned from FilterAttributeDeleted and is used to iterate over the raw logs and unpacked data for AttributeDeleted events raised by the Attribute contract.
type AttributeAttributeDeletedIterator struct {
	Event *AttributeAttributeDeleted // Event containing the contract specifics and raw log

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
func (it *AttributeAttributeDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttributeAttributeDeleted)
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
		it.Event = new(AttributeAttributeDeleted)
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
func (it *AttributeAttributeDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttributeAttributeDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttributeAttributeDeleted represents a AttributeDeleted event raised by the Attribute contract.
type AttributeAttributeDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAttributeDeleted is a free log retrieval operation binding the contract event 0xfe96137d7877697f141204379bad0e6617ff7023d7608f9db3be956d7ceef711.
//
// Solidity: event AttributeDeleted(uint256 indexed id)
func (_Attribute *AttributeFilterer) FilterAttributeDeleted(opts *bind.FilterOpts, id []*big.Int) (*AttributeAttributeDeletedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Attribute.contract.FilterLogs(opts, "AttributeDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return &AttributeAttributeDeletedIterator{contract: _Attribute.contract, event: "AttributeDeleted", logs: logs, sub: sub}, nil
}

// WatchAttributeDeleted is a free log subscription operation binding the contract event 0xfe96137d7877697f141204379bad0e6617ff7023d7608f9db3be956d7ceef711.
//
// Solidity: event AttributeDeleted(uint256 indexed id)
func (_Attribute *AttributeFilterer) WatchAttributeDeleted(opts *bind.WatchOpts, sink chan<- *AttributeAttributeDeleted, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Attribute.contract.WatchLogs(opts, "AttributeDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttributeAttributeDeleted)
				if err := _Attribute.contract.UnpackLog(event, "AttributeDeleted", log); err != nil {
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

// ParseAttributeDeleted is a log parse operation binding the contract event 0xfe96137d7877697f141204379bad0e6617ff7023d7608f9db3be956d7ceef711.
//
// Solidity: event AttributeDeleted(uint256 indexed id)
func (_Attribute *AttributeFilterer) ParseAttributeDeleted(log types.Log) (*AttributeAttributeDeleted, error) {
	event := new(AttributeAttributeDeleted)
	if err := _Attribute.contract.UnpackLog(event, "AttributeDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttributeAttributeUpdatedIterator is returned from FilterAttributeUpdated and is used to iterate over the raw logs and unpacked data for AttributeUpdated events raised by the Attribute contract.
type AttributeAttributeUpdatedIterator struct {
	Event *AttributeAttributeUpdated // Event containing the contract specifics and raw log

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
func (it *AttributeAttributeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttributeAttributeUpdated)
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
		it.Event = new(AttributeAttributeUpdated)
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
func (it *AttributeAttributeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttributeAttributeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttributeAttributeUpdated represents a AttributeUpdated event raised by the Attribute contract.
type AttributeAttributeUpdated struct {
	Id    *big.Int
	Name  string
	Value []*big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAttributeUpdated is a free log retrieval operation binding the contract event 0x8f26e9a323213025e76e7c58dbf6338a9c36708b0f933b7e5babc20acca4e8d1.
//
// Solidity: event AttributeUpdated(uint256 indexed id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) FilterAttributeUpdated(opts *bind.FilterOpts, id []*big.Int) (*AttributeAttributeUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Attribute.contract.FilterLogs(opts, "AttributeUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &AttributeAttributeUpdatedIterator{contract: _Attribute.contract, event: "AttributeUpdated", logs: logs, sub: sub}, nil
}

// WatchAttributeUpdated is a free log subscription operation binding the contract event 0x8f26e9a323213025e76e7c58dbf6338a9c36708b0f933b7e5babc20acca4e8d1.
//
// Solidity: event AttributeUpdated(uint256 indexed id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) WatchAttributeUpdated(opts *bind.WatchOpts, sink chan<- *AttributeAttributeUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Attribute.contract.WatchLogs(opts, "AttributeUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttributeAttributeUpdated)
				if err := _Attribute.contract.UnpackLog(event, "AttributeUpdated", log); err != nil {
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

// ParseAttributeUpdated is a log parse operation binding the contract event 0x8f26e9a323213025e76e7c58dbf6338a9c36708b0f933b7e5babc20acca4e8d1.
//
// Solidity: event AttributeUpdated(uint256 indexed id, string name, uint256[] value)
func (_Attribute *AttributeFilterer) ParseAttributeUpdated(log types.Log) (*AttributeAttributeUpdated, error) {
	event := new(AttributeAttributeUpdated)
	if err := _Attribute.contract.UnpackLog(event, "AttributeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
