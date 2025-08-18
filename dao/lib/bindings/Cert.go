// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// CertMetaData contains all meta data concerning the Cert contract.
var CertMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"name\":\"Issued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Certificates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_grade\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Cert",
	Bin: "0x6080604052348015600e575f5ffd5b5060405161094d38038061094d833981016040819052602b9160b4565b806001600160a01b038116605857604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b605f816065565b505060df565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f6020828403121560c3575f5ffd5b81516001600160a01b038116811460d8575f5ffd5b9392505050565b610861806100ec5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063715018a6146100595780638da5cb5b146100635780639622c83614610082578063d1e7be26146100a5578063f2fde38b146100b8575b5f5ffd5b6100616100cb565b005b5f546040516001600160a01b0390911681526020015b60405180910390f35b6100956100903660046104a2565b6100de565b60405161007994939291906104e7565b6100616100b33660046105dd565b61031d565b6100616100c636600461069f565b6103e5565b6100d3610427565b6100dc5f610453565b565b60016020525f90815260409020805481906100f8906106cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610124906106cc565b801561016f5780601f106101465761010080835404028352916020019161016f565b820191905f5260205f20905b81548152906001019060200180831161015257829003601f168201915b505050505090806001018054610184906106cc565b80601f01602080910402602001604051908101604052809291908181526020018280546101b0906106cc565b80156101fb5780601f106101d2576101008083540402835291602001916101fb565b820191905f5260205f20905b8154815290600101906020018083116101de57829003601f168201915b505050505090806002018054610210906106cc565b80601f016020809104026020016040519081016040528092919081815260200182805461023c906106cc565b80156102875780601f1061025e57610100808354040283529160200191610287565b820191905f5260205f20905b81548152906001019060200180831161026a57829003601f168201915b50505050509080600301805461029c906106cc565b80601f01602080910402602001604051908101604052809291908181526020018280546102c8906106cc565b80156103135780601f106102ea57610100808354040283529160200191610313565b820191905f5260205f20905b8154815290600101906020018083116102f657829003601f168201915b5050505050905084565b610325610427565b604080516080810182528581526020808201869052818301859052606082018490525f888152600190915291909120815181906103629082610750565b50602082015160018201906103779082610750565b506040820151600282019061038c9082610750565b50606082015160038201906103a19082610750565b509050507f6f0d4ed09360dcc2710057310fc932744b9cfddf8c87f0ed6cf244b7abd2f9a085826040516103d692919061080b565b60405180910390a15050505050565b6103ed610427565b6001600160a01b03811661041b57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61042481610453565b50565b5f546001600160a01b031633146100dc5760405163118cdaa760e01b8152336004820152602401610412565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156104b2575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b608081525f6104f960808301876104b9565b828103602084015261050b81876104b9565b9050828103604084015261051f81866104b9565b9050828103606084015261053381856104b9565b979650505050505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610561575f5ffd5b813567ffffffffffffffff81111561057b5761057b61053e565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156105aa576105aa61053e565b6040528181528382016020018510156105c1575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f60a086880312156105f1575f5ffd5b85359450602086013567ffffffffffffffff81111561060e575f5ffd5b61061a88828901610552565b945050604086013567ffffffffffffffff811115610636575f5ffd5b61064288828901610552565b935050606086013567ffffffffffffffff81111561065e575f5ffd5b61066a88828901610552565b925050608086013567ffffffffffffffff811115610686575f5ffd5b61069288828901610552565b9150509295509295909350565b5f602082840312156106af575f5ffd5b81356001600160a01b03811681146106c5575f5ffd5b9392505050565b600181811c908216806106e057607f821691505b6020821081036106fe57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561074b57805f5260205f20601f840160051c810160208510156107295750805b601f840160051c820191505b81811015610748575f8155600101610735565b50505b505050565b815167ffffffffffffffff81111561076a5761076a61053e565b61077e8161077884546106cc565b84610704565b6020601f8211600181146107b0575f83156107995750848201515b5f19600385901b1c1916600184901b178455610748565b5f84815260208120601f198516915b828110156107df57878501518255602094850194600190920191016107bf565b50848210156107fc57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b828152604060208201525f61082360408301846104b9565b94935050505056fea26469706673582212209e1ee4a35404c8e85b70b458ca42dd4a810f0ea2e9535d2ebe5a5348dcac96f464736f6c634300081e0033",
}

// Cert is an auto generated Go binding around an Ethereum contract.
type Cert struct {
	abi abi.ABI
}

// NewCert creates a new instance of Cert.
func NewCert() *Cert {
	parsed, err := CertMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Cert{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Cert) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address initialOwner) returns()
func (cert *Cert) PackConstructor(initialOwner common.Address) []byte {
	enc, err := cert.abi.Pack("", initialOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCertificates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9622c836.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (cert *Cert) PackCertificates(arg0 *big.Int) []byte {
	enc, err := cert.abi.Pack("Certificates", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCertificates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9622c836.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (cert *Cert) TryPackCertificates(arg0 *big.Int) ([]byte, error) {
	return cert.abi.Pack("Certificates", arg0)
}

// CertificatesOutput serves as a container for the return parameters of contract
// method Certificates.
type CertificatesOutput struct {
	Name   string
	Course string
	Grade  string
	Date   string
}

// UnpackCertificates is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9622c836.
//
// Solidity: function Certificates(uint256 ) view returns(string name, string course, string grade, string date)
func (cert *Cert) UnpackCertificates(data []byte) (CertificatesOutput, error) {
	out, err := cert.abi.Unpack("Certificates", data)
	outstruct := new(CertificatesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Course = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Grade = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Date = *abi.ConvertType(out[3], new(string)).(*string)
	return *outstruct, nil
}

// PackIssue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1e7be26.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function issue(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (cert *Cert) PackIssue(id *big.Int, name string, course string, grade string, date string) []byte {
	enc, err := cert.abi.Pack("issue", id, name, course, grade, date)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIssue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1e7be26.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function issue(uint256 _id, string _name, string _course, string _grade, string _date) returns()
func (cert *Cert) TryPackIssue(id *big.Int, name string, course string, grade string, date string) ([]byte, error) {
	return cert.abi.Pack("issue", id, name, course, grade, date)
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (cert *Cert) PackOwner() []byte {
	enc, err := cert.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (cert *Cert) TryPackOwner() ([]byte, error) {
	return cert.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (cert *Cert) UnpackOwner(data []byte) (common.Address, error) {
	out, err := cert.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (cert *Cert) PackRenounceOwnership() []byte {
	enc, err := cert.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceOwnership() returns()
func (cert *Cert) TryPackRenounceOwnership() ([]byte, error) {
	return cert.abi.Pack("renounceOwnership")
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (cert *Cert) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := cert.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (cert *Cert) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return cert.abi.Pack("transferOwnership", newOwner)
}

// CertIssued represents a Issued event raised by the Cert contract.
type CertIssued struct {
	Id   *big.Int
	Date string
	Raw  *types.Log // Blockchain specific contextual infos
}

const CertIssuedEventName = "Issued"

// ContractEventName returns the user-defined event name.
func (CertIssued) ContractEventName() string {
	return CertIssuedEventName
}

// UnpackIssuedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Issued(uint256 id, string date)
func (cert *Cert) UnpackIssuedEvent(log *types.Log) (*CertIssued, error) {
	event := "Issued"
	if log.Topics[0] != cert.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CertIssued)
	if len(log.Data) > 0 {
		if err := cert.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cert.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// CertOwnershipTransferred represents a OwnershipTransferred event raised by the Cert contract.
type CertOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const CertOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (CertOwnershipTransferred) ContractEventName() string {
	return CertOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (cert *Cert) UnpackOwnershipTransferredEvent(log *types.Log) (*CertOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != cert.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CertOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := cert.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range cert.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (cert *Cert) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], cert.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return cert.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], cert.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return cert.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// CertOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the Cert contract.
type CertOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func CertOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (cert *Cert) UnpackOwnableInvalidOwnerError(raw []byte) (*CertOwnableInvalidOwner, error) {
	out := new(CertOwnableInvalidOwner)
	if err := cert.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// CertOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the Cert contract.
type CertOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func CertOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (cert *Cert) UnpackOwnableUnauthorizedAccountError(raw []byte) (*CertOwnableUnauthorizedAccount, error) {
	out := new(CertOwnableUnauthorizedAccount)
	if err := cert.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}
