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

// TimeLockMetaData contains all meta data concerning the TimeLock contract.
var TimeLockMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minDelay\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"proposers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"executors\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDelay\",\"type\":\"uint256\"}],\"name\":\"TimelockInsufficientDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payloads\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"TimelockInvalidOperationLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TimelockUnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"}],\"name\":\"TimelockUnexecutedPredecessor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operationId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"TimelockUnexpectedOperationState\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"CallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"CallSalt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"CallScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Cancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"MinDelayChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CANCELLER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getOperationState\",\"outputs\":[{\"internalType\":\"enumTimelockController.OperationState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"hashOperation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"hashOperationBatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"schedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"scheduleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"updateDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "TimeLock",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611d33380380611d3383398101604081905261002e916102fe565b8383838361003c5f30610183565b506001600160a01b03811615610058576100565f82610183565b505b5f5b83518110156100ec576100ac7fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc18583815181106100995761009961037d565b602002602001015161018360201b60201c565b506100e37ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f7838583815181106100995761009961037d565b5060010161005a565b505f5b82518110156101375761012e7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e638483815181106100995761009961037d565b506001016100ef565b506002849055604080515f8152602081018690527f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5910160405180910390a15050505050505050610391565b5f828152602081815260408083206001600160a01b038516845290915281205460ff16610223575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556101db3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610226565b505f5b92915050565b634e487b7160e01b5f52604160045260245ffd5b80516001600160a01b0381168114610256575f5ffd5b919050565b5f82601f83011261026a575f5ffd5b81516001600160401b038111156102835761028361022c565b604051600582901b90603f8201601f191681016001600160401b03811182821017156102b1576102b161022c565b6040529182526020818501810192908101868411156102ce575f5ffd5b6020860192505b838310156102f4576102e683610240565b8152602092830192016102d5565b5095945050505050565b5f5f5f5f60808587031215610311575f5ffd5b845160208601519094506001600160401b0381111561032e575f5ffd5b61033a8782880161025b565b604087015190945090506001600160401b03811115610357575f5ffd5b6103638782880161025b565b92505061037260608601610240565b905092959194509250565b634e487b7160e01b5f52603260045260245ffd5b6119958061039e5f395ff3fe6080604052600436106101b2575f3560e01c80638065657f116100e7578063bc197c8111610087578063d547741f11610062578063d547741f14610546578063e38335e514610565578063f23a6e6114610578578063f27a0c92146105a3575f5ffd5b8063bc197c81146104d1578063c4d252f5146104fc578063d45c44351461051b575f5ffd5b806391d14854116100c257806391d148541461044d578063a217fddf1461046c578063b08e51c01461047f578063b1c5f427146104b2575f5ffd5b80638065657f146103dc5780638f2a0bb0146103fb5780638f61f4f51461041a575f5ffd5b80632ab0f5291161015257806336568abe1161012d57806336568abe14610353578063584b153e1461037257806364d62353146103915780637958004c146103b0575f5ffd5b80632ab0f529146102f65780632f2ff15d1461031557806331d5075014610334575f5ffd5b8063134008d31161018d578063134008d31461025357806313bc9f2014610266578063150b7a0214610285578063248a9ca3146102c8575f5ffd5b806301d5062a146101bd57806301ffc9a7146101de57806307bd026514610212575f5ffd5b366101b957005b5f5ffd5b3480156101c8575f5ffd5b506101dc6101d7366004611162565b6105b7565b005b3480156101e9575f5ffd5b506101fd6101f83660046111d0565b61068b565b60405190151581526020015b60405180910390f35b34801561021d575f5ffd5b506102457fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6381565b604051908152602001610209565b6101dc6102613660046111f7565b61069b565b348015610271575f5ffd5b506101fd61028036600461125d565b61074d565b348015610290575f5ffd5b506102af61029f366004611323565b630a85bd0160e11b949350505050565b6040516001600160e01b03199091168152602001610209565b3480156102d3575f5ffd5b506102456102e236600461125d565b5f9081526020819052604090206001015490565b348015610301575f5ffd5b506101fd61031036600461125d565b610772565b348015610320575f5ffd5b506101dc61032f366004611386565b61077a565b34801561033f575f5ffd5b506101fd61034e36600461125d565b6107a4565b34801561035e575f5ffd5b506101dc61036d366004611386565b6107c8565b34801561037d575f5ffd5b506101fd61038c36600461125d565b610800565b34801561039c575f5ffd5b506101dc6103ab36600461125d565b610845565b3480156103bb575f5ffd5b506103cf6103ca36600461125d565b6108b8565b60405161020991906113c4565b3480156103e7575f5ffd5b506102456103f63660046111f7565b610900565b348015610406575f5ffd5b506101dc61041536600461142a565b61093e565b348015610425575f5ffd5b506102457fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc181565b348015610458575f5ffd5b506101fd610467366004611386565b610aca565b348015610477575f5ffd5b506102455f81565b34801561048a575f5ffd5b506102457ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f78381565b3480156104bd575f5ffd5b506102456104cc3660046114dc565b610af2565b3480156104dc575f5ffd5b506102af6104eb366004611605565b63bc197c8160e01b95945050505050565b348015610507575f5ffd5b506101dc61051636600461125d565b610b36565b348015610526575f5ffd5b5061024561053536600461125d565b5f9081526001602052604090205490565b348015610551575f5ffd5b506101dc610560366004611386565b610be0565b6101dc6105733660046114dc565b610c04565b348015610583575f5ffd5b506102af6105923660046116b1565b63f23a6e6160e01b95945050505050565b3480156105ae575f5ffd5b50600254610245565b7fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc16105e181610d85565b5f6105f0898989898989610900565b90506105fc8184610d92565b5f817f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca8b8b8b8b8b8a6040516106379695949392919061172c565b60405180910390a3831561068057807f20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d03878560405161067791815260200190565b60405180910390a25b505050505050505050565b5f61069582610e23565b92915050565b7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e636106c6815f610aca565b6106d4576106d48133610e47565b5f6106e3888888888888610900565b90506106ef8185610e84565b6106fb88888888610ed2565b5f817fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b588a8a8a8a6040516107329493929190611768565b60405180910390a361074381610f46565b5050505050505050565b5f60025b61075a836108b8565b600381111561076b5761076b6113b0565b1492915050565b5f6003610751565b5f8281526020819052604090206001015461079481610d85565b61079e8383610f71565b50505050565b5f806107af836108b8565b60038111156107c0576107c06113b0565b141592915050565b6001600160a01b03811633146107f15760405163334bd91960e11b815260040160405180910390fd5b6107fb8282611000565b505050565b5f5f61080b836108b8565b90506001816003811115610821576108216113b0565b148061083e5750600281600381111561083c5761083c6113b0565b145b9392505050565b333081146108765760405163e2850c5960e01b81526001600160a01b03821660048201526024015b60405180910390fd5b60025460408051918252602082018490527f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5910160405180910390a150600255565b5f81815260016020526040812054805f036108d557505f92915050565b600181036108e65750600392915050565b428111156108f75750600192915050565b50600292915050565b5f86868686868660405160200161091c9695949392919061172c565b6040516020818303038152906040528051906020012090509695505050505050565b7fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc161096881610d85565b88871415806109775750888514155b156109a9576040516001624fcdef60e01b03198152600481018a9052602481018690526044810188905260640161086d565b5f6109ba8b8b8b8b8b8b8b8b610af2565b90506109c68184610d92565b5f5b8a811015610a7b5780827f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca8e8e85818110610a0557610a0561178f565b9050602002016020810190610a1a91906117a3565b8d8d86818110610a2c57610a2c61178f565b905060200201358c8c87818110610a4557610a4561178f565b9050602002810190610a5791906117bc565b8c8b604051610a6b9695949392919061172c565b60405180910390a36001016109c8565b508315610abd57807f20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d038785604051610ab491815260200190565b60405180910390a25b5050505050505050505050565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b5f8888888888888888604051602001610b12989796959493929190611892565b60405160208183030381529060405280519060200120905098975050505050505050565b7ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f783610b6081610d85565b610b6982610800565b610ba55781610b786002611069565b610b826001611069565b604051635ead8eb560e01b8152600481019390935217602482015260440161086d565b5f828152600160205260408082208290555183917fbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb7091a25050565b5f82815260208190526040902060010154610bfa81610d85565b61079e8383611000565b7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63610c2f815f610aca565b610c3d57610c3d8133610e47565b8786141580610c4c5750878414155b15610c7e576040516001624fcdef60e01b0319815260048101899052602481018590526044810187905260640161086d565b5f610c8f8a8a8a8a8a8a8a8a610af2565b9050610c9b8185610e84565b5f5b89811015610d6f575f8b8b83818110610cb857610cb861178f565b9050602002016020810190610ccd91906117a3565b90505f8a8a84818110610ce257610ce261178f565b905060200201359050365f8a8a86818110610cff57610cff61178f565b9050602002810190610d1191906117bc565b91509150610d2184848484610ed2565b84867fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b5886868686604051610d589493929190611768565b60405180910390a350505050806001019050610c9d565b50610d7981610f46565b50505050505050505050565b610d8f8133610e47565b50565b610d9b826107a4565b15610dcc5781610daa5f611069565b604051635ead8eb560e01b81526004810192909252602482015260440161086d565b5f610dd660025490565b905080821015610e0357604051635433660960e01b8152600481018390526024810182905260440161086d565b610e0d8242611931565b5f93845260016020526040909320929092555050565b5f6001600160e01b03198216630271189760e51b148061069557506106958261108b565b610e518282610aca565b610e805760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161086d565b5050565b610e8d8261074d565b610e9c5781610daa6002611069565b8015801590610eb15750610eaf81610772565b155b15610e805760405163121534c360e31b81526004810182905260240161086d565b5f5f856001600160a01b0316858585604051610eef929190611950565b5f6040518083038185875af1925050503d805f8114610f29576040519150601f19603f3d011682016040523d82523d5f602084013e610f2e565b606091505b5091509150610f3d82826110bf565b50505050505050565b610f4f8161074d565b610f5e5780610daa6002611069565b5f90815260016020819052604090912055565b5f610f7c8383610aca565b610ff9575f838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055610fb13390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610695565b505f610695565b5f61100b8383610aca565b15610ff9575f838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610695565b5f81600381111561107c5761107c6113b0565b600160ff919091161b92915050565b5f6001600160e01b03198216637965db0b60e01b148061069557506301ffc9a760e01b6001600160e01b0319831614610695565b6060826110d4576110cf826110db565b610695565b5080610695565b8051156110ea57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b80356001600160a01b0381168114611119575f5ffd5b919050565b5f5f83601f84011261112e575f5ffd5b5081356001600160401b03811115611144575f5ffd5b60208301915083602082850101111561115b575f5ffd5b9250929050565b5f5f5f5f5f5f5f60c0888a031215611178575f5ffd5b61118188611103565b96506020880135955060408801356001600160401b038111156111a2575f5ffd5b6111ae8a828b0161111e565b989b979a50986060810135976080820135975060a09091013595509350505050565b5f602082840312156111e0575f5ffd5b81356001600160e01b03198116811461083e575f5ffd5b5f5f5f5f5f5f60a0878903121561120c575f5ffd5b61121587611103565b95506020870135945060408701356001600160401b03811115611236575f5ffd5b61124289828a0161111e565b979a9699509760608101359660809091013595509350505050565b5f6020828403121561126d575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156112b0576112b0611274565b604052919050565b5f82601f8301126112c7575f5ffd5b81356001600160401b038111156112e0576112e0611274565b6112f3601f8201601f1916602001611288565b818152846020838601011115611307575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f60808587031215611336575f5ffd5b61133f85611103565b935061134d60208601611103565b92506040850135915060608501356001600160401b0381111561136e575f5ffd5b61137a878288016112b8565b91505092959194509250565b5f5f60408385031215611397575f5ffd5b823591506113a760208401611103565b90509250929050565b634e487b7160e01b5f52602160045260245ffd5b60208101600483106113e457634e487b7160e01b5f52602160045260245ffd5b91905290565b5f5f83601f8401126113fa575f5ffd5b5081356001600160401b03811115611410575f5ffd5b6020830191508360208260051b850101111561115b575f5ffd5b5f5f5f5f5f5f5f5f5f60c08a8c031215611442575f5ffd5b89356001600160401b03811115611457575f5ffd5b6114638c828d016113ea565b909a5098505060208a01356001600160401b03811115611481575f5ffd5b61148d8c828d016113ea565b90985096505060408a01356001600160401b038111156114ab575f5ffd5b6114b78c828d016113ea565b9a9d999c50979a969997986060880135976080810135975060a0013595509350505050565b5f5f5f5f5f5f5f5f60a0898b0312156114f3575f5ffd5b88356001600160401b03811115611508575f5ffd5b6115148b828c016113ea565b90995097505060208901356001600160401b03811115611532575f5ffd5b61153e8b828c016113ea565b90975095505060408901356001600160401b0381111561155c575f5ffd5b6115688b828c016113ea565b999c989b509699959896976060870135966080013595509350505050565b5f82601f830112611595575f5ffd5b81356001600160401b038111156115ae576115ae611274565b8060051b6115be60208201611288565b918252602081850181019290810190868411156115d9575f5ffd5b6020860192505b838310156115fb5782358252602092830192909101906115e0565b9695505050505050565b5f5f5f5f5f60a08688031215611619575f5ffd5b61162286611103565b945061163060208701611103565b935060408601356001600160401b0381111561164a575f5ffd5b61165688828901611586565b93505060608601356001600160401b03811115611671575f5ffd5b61167d88828901611586565b92505060808601356001600160401b03811115611698575f5ffd5b6116a4888289016112b8565b9150509295509295909350565b5f5f5f5f5f60a086880312156116c5575f5ffd5b6116ce86611103565b94506116dc60208701611103565b9350604086013592506060860135915060808601356001600160401b03811115611698575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60018060a01b038716815285602082015260a060408201525f61175360a083018688611704565b60608301949094525060800152949350505050565b60018060a01b0385168152836020820152606060408201525f6115fb606083018486611704565b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156117b3575f5ffd5b61083e82611103565b5f5f8335601e198436030181126117d1575f5ffd5b8301803591506001600160401b038211156117ea575f5ffd5b60200191503681900382131561115b575f5ffd5b5f8383855260208501945060208460051b820101835f5b8681101561188657838303601f19018852813536879003601e1901811261183a575f5ffd5b86016020810190356001600160401b03811115611855575f5ffd5b803603821315611863575f5ffd5b61186e858284611704565b60209a8b019a90955093909301925050600101611815565b50909695505050505050565b60a080825281018890525f8960c08301825b8b8110156118d2576001600160a01b036118bd84611103565b168252602092830192909101906001016118a4565b5083810360208501528881526001600160fb1b038911156118f1575f5ffd5b8860051b9150818a6020830137018281036020908101604085015261191990820187896117fe565b60608401959095525050608001529695505050505050565b8082018082111561069557634e487b7160e01b5f52601160045260245ffd5b818382375f910190815291905056fea26469706673582212209bebabafa3a92546922f86be8c7f89953d5838a4b536a815440fba8f71043e2964736f6c634300081e0033",
}

// TimeLock is an auto generated Go binding around an Ethereum contract.
type TimeLock struct {
	abi abi.ABI
}

// NewTimeLock creates a new instance of TimeLock.
func NewTimeLock() *TimeLock {
	parsed, err := TimeLockMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TimeLock{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TimeLock) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(uint256 minDelay, address[] proposers, address[] executors, address admin) returns()
func (timeLock *TimeLock) PackConstructor(minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) []byte {
	enc, err := timeLock.abi.Pack("", minDelay, proposers, executors, admin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCANCELLERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb08e51c0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (timeLock *TimeLock) PackCANCELLERROLE() []byte {
	enc, err := timeLock.abi.Pack("CANCELLER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCANCELLERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb08e51c0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (timeLock *TimeLock) TryPackCANCELLERROLE() ([]byte, error) {
	return timeLock.abi.Pack("CANCELLER_ROLE")
}

// UnpackCANCELLERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (timeLock *TimeLock) UnpackCANCELLERROLE(data []byte) ([32]byte, error) {
	out, err := timeLock.abi.Unpack("CANCELLER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (timeLock *TimeLock) PackDEFAULTADMINROLE() []byte {
	enc, err := timeLock.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (timeLock *TimeLock) TryPackDEFAULTADMINROLE() ([]byte, error) {
	return timeLock.abi.Pack("DEFAULT_ADMIN_ROLE")
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (timeLock *TimeLock) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := timeLock.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackEXECUTORROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07bd0265.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (timeLock *TimeLock) PackEXECUTORROLE() []byte {
	enc, err := timeLock.abi.Pack("EXECUTOR_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEXECUTORROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07bd0265.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (timeLock *TimeLock) TryPackEXECUTORROLE() ([]byte, error) {
	return timeLock.abi.Pack("EXECUTOR_ROLE")
}

// UnpackEXECUTORROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (timeLock *TimeLock) UnpackEXECUTORROLE(data []byte) ([32]byte, error) {
	out, err := timeLock.abi.Unpack("EXECUTOR_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackPROPOSERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f61f4f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (timeLock *TimeLock) PackPROPOSERROLE() []byte {
	enc, err := timeLock.abi.Pack("PROPOSER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPROPOSERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f61f4f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (timeLock *TimeLock) TryPackPROPOSERROLE() ([]byte, error) {
	return timeLock.abi.Pack("PROPOSER_ROLE")
}

// UnpackPROPOSERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (timeLock *TimeLock) UnpackPROPOSERROLE(data []byte) ([32]byte, error) {
	out, err := timeLock.abi.Unpack("PROPOSER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d252f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancel(bytes32 id) returns()
func (timeLock *TimeLock) PackCancel(id [32]byte) []byte {
	enc, err := timeLock.abi.Pack("cancel", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d252f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancel(bytes32 id) returns()
func (timeLock *TimeLock) TryPackCancel(id [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("cancel", id)
}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x134008d3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (timeLock *TimeLock) PackExecute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := timeLock.abi.Pack("execute", target, value, payload, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x134008d3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (timeLock *TimeLock) TryPackExecute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("execute", target, value, payload, predecessor, salt)
}

// PackExecuteBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe38335e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (timeLock *TimeLock) PackExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := timeLock.abi.Pack("executeBatch", targets, values, payloads, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe38335e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (timeLock *TimeLock) TryPackExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("executeBatch", targets, values, payloads, predecessor, salt)
}

// PackGetMinDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf27a0c92.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinDelay() view returns(uint256)
func (timeLock *TimeLock) PackGetMinDelay() []byte {
	enc, err := timeLock.abi.Pack("getMinDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf27a0c92.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinDelay() view returns(uint256)
func (timeLock *TimeLock) TryPackGetMinDelay() ([]byte, error) {
	return timeLock.abi.Pack("getMinDelay")
}

// UnpackGetMinDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (timeLock *TimeLock) UnpackGetMinDelay(data []byte) (*big.Int, error) {
	out, err := timeLock.abi.Unpack("getMinDelay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetOperationState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7958004c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (timeLock *TimeLock) PackGetOperationState(id [32]byte) []byte {
	enc, err := timeLock.abi.Pack("getOperationState", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetOperationState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7958004c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (timeLock *TimeLock) TryPackGetOperationState(id [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("getOperationState", id)
}

// UnpackGetOperationState is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (timeLock *TimeLock) UnpackGetOperationState(data []byte) (uint8, error) {
	out, err := timeLock.abi.Unpack("getOperationState", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (timeLock *TimeLock) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := timeLock.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (timeLock *TimeLock) TryPackGetRoleAdmin(role [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("getRoleAdmin", role)
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (timeLock *TimeLock) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := timeLock.abi.Unpack("getRoleAdmin", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd45c4435.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (timeLock *TimeLock) PackGetTimestamp(id [32]byte) []byte {
	enc, err := timeLock.abi.Pack("getTimestamp", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTimestamp is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd45c4435.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (timeLock *TimeLock) TryPackGetTimestamp(id [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("getTimestamp", id)
}

// UnpackGetTimestamp is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (timeLock *TimeLock) UnpackGetTimestamp(data []byte) (*big.Int, error) {
	out, err := timeLock.abi.Unpack("getTimestamp", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (timeLock *TimeLock) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := timeLock.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (timeLock *TimeLock) TryPackGrantRole(role [32]byte, account common.Address) ([]byte, error) {
	return timeLock.abi.Pack("grantRole", role, account)
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (timeLock *TimeLock) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := timeLock.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (timeLock *TimeLock) TryPackHasRole(role [32]byte, account common.Address) ([]byte, error) {
	return timeLock.abi.Pack("hasRole", role, account)
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (timeLock *TimeLock) UnpackHasRole(data []byte) (bool, error) {
	out, err := timeLock.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHashOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8065657f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timeLock *TimeLock) PackHashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := timeLock.abi.Pack("hashOperation", target, value, data, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHashOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8065657f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timeLock *TimeLock) TryPackHashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("hashOperation", target, value, data, predecessor, salt)
}

// UnpackHashOperation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timeLock *TimeLock) UnpackHashOperation(data []byte) ([32]byte, error) {
	out, err := timeLock.abi.Unpack("hashOperation", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackHashOperationBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1c5f427.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timeLock *TimeLock) PackHashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) []byte {
	enc, err := timeLock.abi.Pack("hashOperationBatch", targets, values, payloads, predecessor, salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHashOperationBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1c5f427.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timeLock *TimeLock) TryPackHashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("hashOperationBatch", targets, values, payloads, predecessor, salt)
}

// UnpackHashOperationBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (timeLock *TimeLock) UnpackHashOperationBatch(data []byte) ([32]byte, error) {
	out, err := timeLock.abi.Unpack("hashOperationBatch", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackIsOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31d50750.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (timeLock *TimeLock) PackIsOperation(id [32]byte) []byte {
	enc, err := timeLock.abi.Pack("isOperation", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x31d50750.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (timeLock *TimeLock) TryPackIsOperation(id [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("isOperation", id)
}

// UnpackIsOperation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (timeLock *TimeLock) UnpackIsOperation(data []byte) (bool, error) {
	out, err := timeLock.abi.Unpack("isOperation", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsOperationDone is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab0f529.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (timeLock *TimeLock) PackIsOperationDone(id [32]byte) []byte {
	enc, err := timeLock.abi.Pack("isOperationDone", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperationDone is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ab0f529.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (timeLock *TimeLock) TryPackIsOperationDone(id [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("isOperationDone", id)
}

// UnpackIsOperationDone is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (timeLock *TimeLock) UnpackIsOperationDone(data []byte) (bool, error) {
	out, err := timeLock.abi.Unpack("isOperationDone", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsOperationPending is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x584b153e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (timeLock *TimeLock) PackIsOperationPending(id [32]byte) []byte {
	enc, err := timeLock.abi.Pack("isOperationPending", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperationPending is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x584b153e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (timeLock *TimeLock) TryPackIsOperationPending(id [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("isOperationPending", id)
}

// UnpackIsOperationPending is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (timeLock *TimeLock) UnpackIsOperationPending(data []byte) (bool, error) {
	out, err := timeLock.abi.Unpack("isOperationPending", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsOperationReady is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13bc9f20.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (timeLock *TimeLock) PackIsOperationReady(id [32]byte) []byte {
	enc, err := timeLock.abi.Pack("isOperationReady", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperationReady is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13bc9f20.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (timeLock *TimeLock) TryPackIsOperationReady(id [32]byte) ([]byte, error) {
	return timeLock.abi.Pack("isOperationReady", id)
}

// UnpackIsOperationReady is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (timeLock *TimeLock) UnpackIsOperationReady(data []byte) (bool, error) {
	out, err := timeLock.abi.Unpack("isOperationReady", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (timeLock *TimeLock) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := timeLock.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (timeLock *TimeLock) TryPackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([]byte, error) {
	return timeLock.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (timeLock *TimeLock) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := timeLock.abi.Unpack("onERC1155BatchReceived", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (timeLock *TimeLock) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := timeLock.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC1155Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf23a6e61.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (timeLock *TimeLock) TryPackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return timeLock.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (timeLock *TimeLock) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := timeLock.abi.Unpack("onERC1155Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (timeLock *TimeLock) PackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) []byte {
	enc, err := timeLock.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOnERC721Received is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x150b7a02.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (timeLock *TimeLock) TryPackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([]byte, error) {
	return timeLock.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (timeLock *TimeLock) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := timeLock.abi.Unpack("onERC721Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (timeLock *TimeLock) PackRenounceRole(role [32]byte, callerConfirmation common.Address) []byte {
	enc, err := timeLock.abi.Pack("renounceRole", role, callerConfirmation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (timeLock *TimeLock) TryPackRenounceRole(role [32]byte, callerConfirmation common.Address) ([]byte, error) {
	return timeLock.abi.Pack("renounceRole", role, callerConfirmation)
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (timeLock *TimeLock) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := timeLock.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (timeLock *TimeLock) TryPackRevokeRole(role [32]byte, account common.Address) ([]byte, error) {
	return timeLock.abi.Pack("revokeRole", role, account)
}

// PackSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01d5062a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (timeLock *TimeLock) PackSchedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) []byte {
	enc, err := timeLock.abi.Pack("schedule", target, value, data, predecessor, salt, delay)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01d5062a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (timeLock *TimeLock) TryPackSchedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) ([]byte, error) {
	return timeLock.abi.Pack("schedule", target, value, data, predecessor, salt, delay)
}

// PackScheduleBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f2a0bb0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (timeLock *TimeLock) PackScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) []byte {
	enc, err := timeLock.abi.Pack("scheduleBatch", targets, values, payloads, predecessor, salt, delay)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackScheduleBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f2a0bb0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (timeLock *TimeLock) TryPackScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) ([]byte, error) {
	return timeLock.abi.Pack("scheduleBatch", targets, values, payloads, predecessor, salt, delay)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (timeLock *TimeLock) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := timeLock.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (timeLock *TimeLock) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return timeLock.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (timeLock *TimeLock) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := timeLock.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackUpdateDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64d62353.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (timeLock *TimeLock) PackUpdateDelay(newDelay *big.Int) []byte {
	enc, err := timeLock.abi.Pack("updateDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64d62353.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (timeLock *TimeLock) TryPackUpdateDelay(newDelay *big.Int) ([]byte, error) {
	return timeLock.abi.Pack("updateDelay", newDelay)
}

// TimeLockCallExecuted represents a CallExecuted event raised by the TimeLock contract.
type TimeLockCallExecuted struct {
	Id     [32]byte
	Index  *big.Int
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    *types.Log // Blockchain specific contextual infos
}

const TimeLockCallExecutedEventName = "CallExecuted"

// ContractEventName returns the user-defined event name.
func (TimeLockCallExecuted) ContractEventName() string {
	return TimeLockCallExecutedEventName
}

// UnpackCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (timeLock *TimeLock) UnpackCallExecutedEvent(log *types.Log) (*TimeLockCallExecuted, error) {
	event := "CallExecuted"
	if log.Topics[0] != timeLock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimeLockCallExecuted)
	if len(log.Data) > 0 {
		if err := timeLock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timeLock.abi.Events[event].Inputs {
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

// TimeLockCallSalt represents a CallSalt event raised by the TimeLock contract.
type TimeLockCallSalt struct {
	Id   [32]byte
	Salt [32]byte
	Raw  *types.Log // Blockchain specific contextual infos
}

const TimeLockCallSaltEventName = "CallSalt"

// ContractEventName returns the user-defined event name.
func (TimeLockCallSalt) ContractEventName() string {
	return TimeLockCallSaltEventName
}

// UnpackCallSaltEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (timeLock *TimeLock) UnpackCallSaltEvent(log *types.Log) (*TimeLockCallSalt, error) {
	event := "CallSalt"
	if log.Topics[0] != timeLock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimeLockCallSalt)
	if len(log.Data) > 0 {
		if err := timeLock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timeLock.abi.Events[event].Inputs {
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

// TimeLockCallScheduled represents a CallScheduled event raised by the TimeLock contract.
type TimeLockCallScheduled struct {
	Id          [32]byte
	Index       *big.Int
	Target      common.Address
	Value       *big.Int
	Data        []byte
	Predecessor [32]byte
	Delay       *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const TimeLockCallScheduledEventName = "CallScheduled"

// ContractEventName returns the user-defined event name.
func (TimeLockCallScheduled) ContractEventName() string {
	return TimeLockCallScheduledEventName
}

// UnpackCallScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (timeLock *TimeLock) UnpackCallScheduledEvent(log *types.Log) (*TimeLockCallScheduled, error) {
	event := "CallScheduled"
	if log.Topics[0] != timeLock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimeLockCallScheduled)
	if len(log.Data) > 0 {
		if err := timeLock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timeLock.abi.Events[event].Inputs {
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

// TimeLockCancelled represents a Cancelled event raised by the TimeLock contract.
type TimeLockCancelled struct {
	Id  [32]byte
	Raw *types.Log // Blockchain specific contextual infos
}

const TimeLockCancelledEventName = "Cancelled"

// ContractEventName returns the user-defined event name.
func (TimeLockCancelled) ContractEventName() string {
	return TimeLockCancelledEventName
}

// UnpackCancelledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (timeLock *TimeLock) UnpackCancelledEvent(log *types.Log) (*TimeLockCancelled, error) {
	event := "Cancelled"
	if log.Topics[0] != timeLock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimeLockCancelled)
	if len(log.Data) > 0 {
		if err := timeLock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timeLock.abi.Events[event].Inputs {
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

// TimeLockMinDelayChange represents a MinDelayChange event raised by the TimeLock contract.
type TimeLockMinDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const TimeLockMinDelayChangeEventName = "MinDelayChange"

// ContractEventName returns the user-defined event name.
func (TimeLockMinDelayChange) ContractEventName() string {
	return TimeLockMinDelayChangeEventName
}

// UnpackMinDelayChangeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (timeLock *TimeLock) UnpackMinDelayChangeEvent(log *types.Log) (*TimeLockMinDelayChange, error) {
	event := "MinDelayChange"
	if log.Topics[0] != timeLock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimeLockMinDelayChange)
	if len(log.Data) > 0 {
		if err := timeLock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timeLock.abi.Events[event].Inputs {
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

// TimeLockRoleAdminChanged represents a RoleAdminChanged event raised by the TimeLock contract.
type TimeLockRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const TimeLockRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (TimeLockRoleAdminChanged) ContractEventName() string {
	return TimeLockRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (timeLock *TimeLock) UnpackRoleAdminChangedEvent(log *types.Log) (*TimeLockRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != timeLock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimeLockRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := timeLock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timeLock.abi.Events[event].Inputs {
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

// TimeLockRoleGranted represents a RoleGranted event raised by the TimeLock contract.
type TimeLockRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TimeLockRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (TimeLockRoleGranted) ContractEventName() string {
	return TimeLockRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (timeLock *TimeLock) UnpackRoleGrantedEvent(log *types.Log) (*TimeLockRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != timeLock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimeLockRoleGranted)
	if len(log.Data) > 0 {
		if err := timeLock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timeLock.abi.Events[event].Inputs {
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

// TimeLockRoleRevoked represents a RoleRevoked event raised by the TimeLock contract.
type TimeLockRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const TimeLockRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (TimeLockRoleRevoked) ContractEventName() string {
	return TimeLockRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (timeLock *TimeLock) UnpackRoleRevokedEvent(log *types.Log) (*TimeLockRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != timeLock.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TimeLockRoleRevoked)
	if len(log.Data) > 0 {
		if err := timeLock.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range timeLock.abi.Events[event].Inputs {
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
func (timeLock *TimeLock) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], timeLock.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return timeLock.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], timeLock.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return timeLock.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], timeLock.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return timeLock.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], timeLock.abi.Errors["TimelockInsufficientDelay"].ID.Bytes()[:4]) {
		return timeLock.UnpackTimelockInsufficientDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], timeLock.abi.Errors["TimelockInvalidOperationLength"].ID.Bytes()[:4]) {
		return timeLock.UnpackTimelockInvalidOperationLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], timeLock.abi.Errors["TimelockUnauthorizedCaller"].ID.Bytes()[:4]) {
		return timeLock.UnpackTimelockUnauthorizedCallerError(raw[4:])
	}
	if bytes.Equal(raw[:4], timeLock.abi.Errors["TimelockUnexecutedPredecessor"].ID.Bytes()[:4]) {
		return timeLock.UnpackTimelockUnexecutedPredecessorError(raw[4:])
	}
	if bytes.Equal(raw[:4], timeLock.abi.Errors["TimelockUnexpectedOperationState"].ID.Bytes()[:4]) {
		return timeLock.UnpackTimelockUnexpectedOperationStateError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TimeLockAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the TimeLock contract.
type TimeLockAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func TimeLockAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (timeLock *TimeLock) UnpackAccessControlBadConfirmationError(raw []byte) (*TimeLockAccessControlBadConfirmation, error) {
	out := new(TimeLockAccessControlBadConfirmation)
	if err := timeLock.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TimeLockAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the TimeLock contract.
type TimeLockAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func TimeLockAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (timeLock *TimeLock) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*TimeLockAccessControlUnauthorizedAccount, error) {
	out := new(TimeLockAccessControlUnauthorizedAccount)
	if err := timeLock.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TimeLockFailedCall represents a FailedCall error raised by the TimeLock contract.
type TimeLockFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TimeLockFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (timeLock *TimeLock) UnpackFailedCallError(raw []byte) (*TimeLockFailedCall, error) {
	out := new(TimeLockFailedCall)
	if err := timeLock.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TimeLockTimelockInsufficientDelay represents a TimelockInsufficientDelay error raised by the TimeLock contract.
type TimeLockTimelockInsufficientDelay struct {
	Delay    *big.Int
	MinDelay *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockInsufficientDelay(uint256 delay, uint256 minDelay)
func TimeLockTimelockInsufficientDelayErrorID() common.Hash {
	return common.HexToHash("0x543366097fc46a718a4890a59e5216cbbc872f61973a9e2a6666ecccaa1a07ca")
}

// UnpackTimelockInsufficientDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockInsufficientDelay(uint256 delay, uint256 minDelay)
func (timeLock *TimeLock) UnpackTimelockInsufficientDelayError(raw []byte) (*TimeLockTimelockInsufficientDelay, error) {
	out := new(TimeLockTimelockInsufficientDelay)
	if err := timeLock.abi.UnpackIntoInterface(out, "TimelockInsufficientDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TimeLockTimelockInvalidOperationLength represents a TimelockInvalidOperationLength error raised by the TimeLock contract.
type TimeLockTimelockInvalidOperationLength struct {
	Targets  *big.Int
	Payloads *big.Int
	Values   *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockInvalidOperationLength(uint256 targets, uint256 payloads, uint256 values)
func TimeLockTimelockInvalidOperationLengthErrorID() common.Hash {
	return common.HexToHash("0xffb0321166292613ddbe36e2eb9b9b1e8877aa314505f6d35f0f4ae660e8ada3")
}

// UnpackTimelockInvalidOperationLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockInvalidOperationLength(uint256 targets, uint256 payloads, uint256 values)
func (timeLock *TimeLock) UnpackTimelockInvalidOperationLengthError(raw []byte) (*TimeLockTimelockInvalidOperationLength, error) {
	out := new(TimeLockTimelockInvalidOperationLength)
	if err := timeLock.abi.UnpackIntoInterface(out, "TimelockInvalidOperationLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TimeLockTimelockUnauthorizedCaller represents a TimelockUnauthorizedCaller error raised by the TimeLock contract.
type TimeLockTimelockUnauthorizedCaller struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockUnauthorizedCaller(address caller)
func TimeLockTimelockUnauthorizedCallerErrorID() common.Hash {
	return common.HexToHash("0xe2850c5900ceb2d1b367e63ffd96510279f191862fece2dde12a1b1bce580ebd")
}

// UnpackTimelockUnauthorizedCallerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockUnauthorizedCaller(address caller)
func (timeLock *TimeLock) UnpackTimelockUnauthorizedCallerError(raw []byte) (*TimeLockTimelockUnauthorizedCaller, error) {
	out := new(TimeLockTimelockUnauthorizedCaller)
	if err := timeLock.abi.UnpackIntoInterface(out, "TimelockUnauthorizedCaller", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TimeLockTimelockUnexecutedPredecessor represents a TimelockUnexecutedPredecessor error raised by the TimeLock contract.
type TimeLockTimelockUnexecutedPredecessor struct {
	PredecessorId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockUnexecutedPredecessor(bytes32 predecessorId)
func TimeLockTimelockUnexecutedPredecessorErrorID() common.Hash {
	return common.HexToHash("0x90a9a618cfd84aabd2505bb50bbc6c95924a5d4f10d2bf500768fcbf91f51cab")
}

// UnpackTimelockUnexecutedPredecessorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockUnexecutedPredecessor(bytes32 predecessorId)
func (timeLock *TimeLock) UnpackTimelockUnexecutedPredecessorError(raw []byte) (*TimeLockTimelockUnexecutedPredecessor, error) {
	out := new(TimeLockTimelockUnexecutedPredecessor)
	if err := timeLock.abi.UnpackIntoInterface(out, "TimelockUnexecutedPredecessor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TimeLockTimelockUnexpectedOperationState represents a TimelockUnexpectedOperationState error raised by the TimeLock contract.
type TimeLockTimelockUnexpectedOperationState struct {
	OperationId    [32]byte
	ExpectedStates [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockUnexpectedOperationState(bytes32 operationId, bytes32 expectedStates)
func TimeLockTimelockUnexpectedOperationStateErrorID() common.Hash {
	return common.HexToHash("0x5ead8eb51d1c1b7ef2eb1ef3ec2f009cfba25e924d04ccc853f25803ea40b489")
}

// UnpackTimelockUnexpectedOperationStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockUnexpectedOperationState(bytes32 operationId, bytes32 expectedStates)
func (timeLock *TimeLock) UnpackTimelockUnexpectedOperationStateError(raw []byte) (*TimeLockTimelockUnexpectedOperationState, error) {
	out := new(TimeLockTimelockUnexpectedOperationState)
	if err := timeLock.abi.UnpackIntoInterface(out, "TimelockUnexpectedOperationState", raw); err != nil {
		return nil, err
	}
	return out, nil
}
