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

// MyGovernorMetaData contains all meta data concerning the MyGovernor contract.
var MyGovernorMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVotes\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractTimelockController\",\"name\":\"_timelock\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorAlreadyCastVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorAlreadyQueuedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorDisabledDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"GovernorInsufficientProposerVotes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldatas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidProposalLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorumNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quorumDenominator\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidQuorumFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidVotingPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNonexistentProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNotQueuedProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorQueueNotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"GovernorRestrictedProposer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorUnableToCancel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"current\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"GovernorUnexpectedProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTimelock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"TimelockChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"getProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalNeedsQueuing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTimelockController\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"updateTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "MyGovernor",
	Bin: "0x610180604052348015610010575f5ffd5b506040516148b53803806148b583398101604081905261002f916105b1565b806004836040518060400160405280600a81526020016926bca3b7bb32b93737b960b11b8152508061006561014960201b60201c565b61006f825f610164565b6101205261007e816001610164565b61014052815160208084019190912060e052815190820120610100524660a05261010a60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c052600361011f8282610681565b50506001600160a01b03166101605261013781610196565b5061014181610230565b5050506107d7565b6040805180820190915260018152603160f81b602082015290565b5f60208351101561017f5761017883610299565b9050610190565b8161018a8482610681565b5060ff90505b92915050565b6064808211156101c85760405163243e544560e01b815260048101839052602481018290526044015b60405180910390fd5b5f6101d16102d6565b90506101f06101de6102ef565b6101e785610369565b600891906103a0565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b600954604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a1600980546001600160a01b0319166001600160a01b0392909216919091179055565b5f5f829050601f815111156102c3578260405163305a27a960e01b81526004016101bf919061073b565b80516102ce82610770565b179392505050565b5f6102e160086103ba565b6001600160d01b0316905090565b5f6102fa6101605190565b6001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610353575060408051601f3d908101601f1916820190925261035091810190610793565b60015b6103645761035f610402565b905090565b919050565b5f6001600160d01b0382111561039c576040516306dfcc6560e41b815260d06004820152602481018390526044016101bf565b5090565b5f806103ad85858561040c565b915091505b935093915050565b80545f9080156103f9576103e0836103d36001846107b8565b5f91825260209091200190565b54660100000000000090046001600160d01b03166103fb565b5f5b9392505050565b5f61035f43610568565b82545f908190801561050b575f610428876103d36001856107b8565b805490915065ffffffffffff80821691660100000000000090046001600160d01b031690881682111561046e57604051632520601d60e01b815260040160405180910390fd5b8765ffffffffffff168265ffffffffffff16036104aa57825465ffffffffffff1666010000000000006001600160d01b038916021783556104fd565b6040805180820190915265ffffffffffff808a1682526001600160d01b03808a1660208085019182528d54600181018f555f8f815291909120945191519092166601000000000000029216919091179101555b94508593506103b292505050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a555f8a81529182209551925190931666010000000000000291909316179201919091559050816103b2565b5f65ffffffffffff82111561039c576040516306dfcc6560e41b815260306004820152602481018390526044016101bf565b6001600160a01b03811681146105ae575f5ffd5b50565b5f5f604083850312156105c2575f5ffd5b82516105cd8161059a565b60208401519092506105de8161059a565b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061061157607f821691505b60208210810361062f57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561067c57805f5260205f20601f840160051c8101602085101561065a5750805b601f840160051c820191505b81811015610679575f8155600101610666565b50505b505050565b81516001600160401b0381111561069a5761069a6105e9565b6106ae816106a884546105fd565b84610635565b6020601f8211600181146106e0575f83156106c95750848201515b5f19600385901b1c1916600184901b178455610679565b5f84815260208120601f198516915b8281101561070f57878501518255602094850194600190920191016106ef565b508482101561072c57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8051602080830151919081101561062f575f1960209190910360031b1b16919050565b5f602082840312156107a3575f5ffd5b815165ffffffffffff811681146103fb575f5ffd5b8181038181111561019057634e487b7160e01b5f52601160045260245ffd5b60805160a05160c05160e0516101005161012051610140516101605161406761084e5f395f81816108fd01528181610d530152818161109b015281816113340152611cb301525f611bf701525f611bcb01525f612b0c01525f612ae401525f612a3f01525f612a6901525f612a9301526140675ff3fe608060405260043610610278575f3560e01c80637ecebe001161014a578063b58131b0116100be578063dd4e2ba511610078578063dd4e2ba51461081a578063deaaa7cc1461085f578063eb9019d414610892578063f23a6e61146108b1578063f8ce560a146108d0578063fc0c546a146108ef575f5ffd5b8063b58131b01461077b578063bc197c811461078d578063c01f9e37146107ac578063c28bc2fa146107cb578063c59057e4146107de578063d33219b4146107fd575f5ffd5b80639a802a6d1161010f5780639a802a6d146106b5578063a7713a70146106d4578063a890c910146106e8578063a8f8a66814610707578063a9a9529414610726578063ab58fb8e14610745575f5ffd5b80637ecebe001461061057806384b0196e146106445780638ff262e31461066b57806391ddadf41461068a57806397c3d334146102e7575f5ffd5b80633e4f49e6116101ec57806356781388116101a657806356781388146105565780635b8d0e0d146105755780635f398a141461059457806360c4247f146105b35780637b3c71d3146105d25780637d5e81e2146105f1575f5ffd5b80633e4f49e61461044c5780634385963214610478578063452115d6146104c05780634bf5d7e9146104df578063544ffc9c146104f357806354fd4d501461052d575f5ffd5b8063150b7a021161023d578063150b7a0214610390578063160cbed7146103c85780632656227d146103e75780632d63f693146103fa5780632fe3e261146104195780633932abb1146102e7575f5ffd5b806301ffc9a7146102b357806302a251a3146102e757806306f3f9e61461030457806306fdde0314610323578063143489d014610344575f5ffd5b366102af5730610286610921565b6001600160a01b0316146102ad57604051637485328f60e11b815260040160405180910390fd5b005b5f5ffd5b3480156102be575f5ffd5b506102d26102cd3660046130b1565b610939565b60405190151581526020015b60405180910390f35b3480156102f2575f5ffd5b5060645b6040519081526020016102de565b34801561030f575f5ffd5b506102ad61031e3660046130d8565b6109a5565b34801561032e575f5ffd5b506103376109b9565b6040516102de919061311d565b34801561034f575f5ffd5b5061037861035e3660046130d8565b5f908152600460205260409020546001600160a01b031690565b6040516001600160a01b0390911681526020016102de565b34801561039b575f5ffd5b506103af6103aa366004613206565b610a49565b6040516001600160e01b031990911681526020016102de565b3480156103d3575f5ffd5b506102f66103e23660046133cc565b610a8b565b6102f66103f53660046133cc565b610b57565b348015610405575f5ffd5b506102f66104143660046130d8565b610cbf565b348015610424575f5ffd5b506102f67f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b348015610457575f5ffd5b5061046b6104663660046130d8565b610cdf565b6040516102de9190613493565b348015610483575f5ffd5b506102d26104923660046134a1565b5f8281526007602090815260408083206001600160a01b038516845260030190915290205460ff1692915050565b3480156104cb575f5ffd5b506102f66104da3660046133cc565b610ce9565b3480156104ea575f5ffd5b50610337610d4f565b3480156104fe575f5ffd5b5061051261050d3660046130d8565b610e0f565b604080519384526020840192909252908201526060016102de565b348015610538575f5ffd5b506040805180820190915260018152603160f81b6020820152610337565b348015610561575f5ffd5b506102f66105703660046134df565b610e34565b348015610580575f5ffd5b506102f661058f36600461354d565b610e5b565b34801561059f575f5ffd5b506102f66105ae366004613609565b610f18565b3480156105be575f5ffd5b506102f66105cd3660046130d8565b610f60565b3480156105dd575f5ffd5b506102f66105ec36600461368a565b610f6c565b3480156105fc575f5ffd5b506102f661060b3660046136df565b610fbc565b34801561061b575f5ffd5b506102f661062a36600461379b565b6001600160a01b03165f9081526002602052604090205490565b34801561064f575f5ffd5b50610658610ffe565b6040516102de97969594939291906137f0565b348015610676575f5ffd5b506102f661068536600461385f565b611040565b348015610695575f5ffd5b5061069e611098565b60405165ffffffffffff90911681526020016102de565b3480156106c0575f5ffd5b506102f66106cf3660046138ac565b61111f565b3480156106df575f5ffd5b506102f6611135565b3480156106f3575f5ffd5b506102ad61070236600461379b565b61114e565b348015610712575f5ffd5b506102f66107213660046133cc565b61115f565b348015610731575f5ffd5b506102d26107403660046130d8565b61116c565b348015610750575f5ffd5b506102f661075f3660046130d8565b5f9081526004602052604090206001015465ffffffffffff1690565b348015610786575f5ffd5b505f6102f6565b348015610798575f5ffd5b506103af6107a7366004613900565b611174565b3480156107b7575f5ffd5b506102f66107c63660046130d8565b6111b7565b6102ad6107d9366004613997565b6111f9565b3480156107e9575f5ffd5b506102f66107f83660046133cc565b611275565b348015610808575f5ffd5b506009546001600160a01b0316610378565b348015610825575f5ffd5b506040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e90820152610337565b34801561086a575f5ffd5b506102f67ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b34801561089d575f5ffd5b506102f66108ac3660046139d6565b6112ae565b3480156108bc575f5ffd5b506103af6108cb366004613a00565b6112cd565b3480156108db575f5ffd5b506102f66108ea3660046130d8565b611310565b3480156108fa575f5ffd5b507f0000000000000000000000000000000000000000000000000000000000000000610378565b5f6109346009546001600160a01b031690565b905090565b5f6001600160e01b031982166366defe7760e11b148061096957506001600160e01b031982166332a2ad4360e11b145b8061098457506001600160e01b03198216630271189760e51b145b8061099f57506301ffc9a760e01b6001600160e01b03198316145b92915050565b6109ad6113ad565b6109b681611426565b50565b6060600380546109c890613a57565b80601f01602080910402602001604051908101604052809291908181526020018280546109f490613a57565b8015610a3f5780601f10610a1657610100808354040283529160200191610a3f565b820191905f5260205f20905b815481529060010190602001808311610a2257829003601f168201915b5050505050905090565b5f30610a53610921565b6001600160a01b031614610a7a57604051637485328f60e11b815260040160405180910390fd5b50630a85bd0160e11b949350505050565b5f5f610a998686868661115f565b9050610aae81610aa960046114bb565b6114dd565b505f610abd828888888861151a565b905065ffffffffffff811615610b34575f82815260046020908152604091829020600101805465ffffffffffff191665ffffffffffff85169081179091558251858152918201527f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892910160405180910390a1610b4d565b604051634844252360e11b815260040160405180910390fd5b5095945050505050565b5f5f610b658686868661115f565b9050610b8581610b7560056114bb565b610b7f60046114bb565b176114dd565b505f818152600460205260409020805460ff60f01b1916600160f01b17905530610bad610921565b6001600160a01b031614610c36575f5b8651811015610c3457306001600160a01b0316878281518110610be257610be2613a8f565b60200260200101516001600160a01b031603610c2c57610c2c858281518110610c0d57610c0d613a8f565b602002602001015180519060200120600561152890919063ffffffff16565b600101610bbd565b505b610c438187878787611589565b30610c4c610921565b6001600160a01b031614158015610c7857506005546001600160801b03808216600160801b9092041614155b15610c82575f6005555b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906020015b60405180910390a195945050505050565b5f90815260046020526040902054600160a01b900465ffffffffffff1690565b5f61099f8261159d565b5f5f610cf78686868661115f565b905033610d0482826116d6565b610d3857604051638fe5d8a960e01b8152600481018390526001600160a01b03821660248201526044015b60405180910390fd5b610d448787878761171b565b979650505050505050565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634bf5d7e96040518163ffffffff1660e01b81526004015f60405180830381865afa925050508015610dce57506040513d5f823e601f3d908101601f19168201604052610dcb9190810190613aa3565b60015b610e0a575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b5f818152600760205260409020805460018201546002909201549091905b9193909250565b5f80339050610e5384828560405180602001604052805f815250611728565b949350505050565b5f610ea188888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a92508991506117499050565b610ec9576040516394ab6c0760e01b81526001600160a01b0387166004820152602401610d2f565b610f0c88878988888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a9250611816915050565b98975050505050505050565b5f80339050610d4487828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a9250611816915050565b5f61099f6008836118f4565b5f80339050610fb286828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061172892505050565b9695505050505050565b5f33610fc88184611941565b610ff05760405163d9b3955760e01b81526001600160a01b0382166004820152602401610d2f565b5f610d4487878787866119c5565b5f6060805f5f5f606061100f611bc4565b611017611bf0565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f61104d85858585611c1d565b611075576040516394ab6c0760e01b81526001600160a01b0384166004820152602401610d2f565b61108f85848660405180602001604052805f815250611728565b95945050505050565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611113575060408051601f3d908101601f1916820190925261111091810190613b2b565b60015b610e0a57610934611ca6565b5f61112b848484611cb0565b90505b9392505050565b5f6111406008611d43565b6001600160d01b0316905090565b6111566113ad565b6109b681611d87565b5f61108f85858585611275565b5f600161099f565b5f3061117e610921565b6001600160a01b0316146111a557604051637485328f60e11b815260040160405180910390fd5b5063bc197c8160e01b95945050505050565b5f818152600460205260408120546111eb90600160d01b810463ffffffff1690600160a01b900465ffffffffffff16613b50565b65ffffffffffff1692915050565b6112016113ad565b5f5f856001600160a01b031685858560405161121e929190613b6e565b5f6040518083038185875af1925050503d805f8114611258576040519150601f19603f3d011682016040523d82523d5f602084013e61125d565b606091505b509150915061126c8282611df0565b50505050505050565b5f8484848460405160200161128d9493929190613c10565b60408051601f19818403018152919052805160209091012095945050505050565b5f61112e83836112c860408051602081019091525f815290565b611cb0565b5f306112d7610921565b6001600160a01b0316146112fe57604051637485328f60e11b815260040160405180910390fd5b5063f23a6e6160e01b95945050505050565b604051632394e7a360e21b8152600481018290525f9061099f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638e539e8c90602401602060405180830381865afa158015611379573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061139d9190613c5a565b6113a684610f60565b6064611e0c565b336113b6610921565b6001600160a01b0316146113df576040516347096e4760e01b8152336004820152602401610d2f565b306113e8610921565b6001600160a01b031614611424575f8036604051611407929190613b6e565b604051809103902090505b8061141d6005611ebc565b0361141257505b565b6064808211156114535760405163243e544560e01b81526004810183905260248101829052604401610d2f565b5f61145c611135565b905061147b611469611098565b61147285611f29565b60089190611f60565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b5f8160078111156114ce576114ce61345f565b600160ff919091161b92915050565b5f5f6114e884610cdf565b90505f836114f5836114bb565b160361112e578381846040516331b75e4d60e01b8152600401610d2f93929190613c71565b5f610fb28686868686611f7a565b81546001600160801b03600160801b82048116918116600183019091160361155457611554604161210b565b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b611596858585858561211c565b5050505050565b5f5f6115a8836121ac565b905060058160078111156115be576115be61345f565b146115c95792915050565b5f838152600a602052604090819020546009549151632c258a9f60e11b81526004810182905290916001600160a01b03169063584b153e90602401602060405180830381865afa15801561161f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116439190613c93565b15611652575060059392505050565b600954604051632ab0f52960e01b8152600481018390526001600160a01b0390911690632ab0f52990602401602060405180830381865afa158015611699573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116bd9190613c93565b156116cc575060079392505050565b5060029392505050565b5f806116e184610cdf565b60078111156116f2576116f261345f565b14801561112e5750505f91825260046020526040909120546001600160a01b0391821691161490565b5f61108f858585856122dd565b5f61108f8585858561174460408051602081019091525f815290565b611816565b5f610d44856118107f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118a8a8a61179b8c6001600160a01b03165f90815260026020526040902080546001810190915590565b8b516020808e01919091208c518d8301206040516117f598979695949301968752602087019590955260ff9390931660408601526001600160a01b03919091166060850152608084015260a083015260c082015260e00190565b60405160208183030381529060405280519060200120612374565b846123a0565b5f61182586610aa960016114bb565b505f61183a8661183489610cbf565b85611cb0565b90505f61184a8888888588612410565b905083515f036118a057866001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4898884896040516118939493929190613cb2565b60405180910390a2610d44565b866001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb871289888489896040516118e1959493929190613cd9565b60405180910390a2979650505050505050565b5f5f5f6119008561250c565b9250925050838265ffffffffffff16111561192d5761192861192185612564565b8690612596565b61192f565b805b6001600160d01b031695945050505050565b80515f90603481101561195857600191505061099f565b60131981840101516001600160b01b03198116692370726f706f7365723d60b01b146119895760019250505061099f565b5f5f61199986602a860386612638565b91509150811580610d445750866001600160a01b0316816001600160a01b031614979650505050505050565b5f6119d9868686868051906020012061115f565b9050845186511415806119ee57508351865114155b806119f857508551155b15611a2d57855184518651604051630447b05d60e41b8152600481019390935260248301919091526044820152606401610d2f565b5f81815260046020526040902054600160a01b900465ffffffffffff1615611a765780611a5982610cdf565b6040516331b75e4d60e01b8152610d2f9291905f90600401613c71565b5f6064611a81611098565b65ffffffffffff16611a939190613d12565b90505f60645f84815260046020526040902080546001600160a01b0319166001600160a01b038716178155909150611aca83612564565b815465ffffffffffff91909116600160a01b0265ffffffffffff60a01b19909116178155611af7826126e1565b815463ffffffff91909116600160d01b0263ffffffff60d01b1990911617815588517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e090859087908c908c906001600160401b03811115611b5a57611b5a613143565b604051908082528060200260200182016040528015611b8d57816020015b6060815260200190600190039081611b785790505b508c89611b9a8a82613d12565b8e604051611bb099989796959493929190613d25565b60405180910390a150505095945050505050565b60606109347f00000000000000000000000000000000000000000000000000000000000000005f612711565b60606109347f00000000000000000000000000000000000000000000000000000000000000006001612711565b5f61108f836118107ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d7888888611c6f8a6001600160a01b03165f90815260026020526040902080546001810190915590565b60408051602081019690965285019390935260ff90911660608401526001600160a01b0316608083015260a082015260c0016117f5565b5f61093443612564565b5f7f0000000000000000000000000000000000000000000000000000000000000000604051630748d63560e31b81526001600160a01b038681166004830152602482018690529190911690633a46b1a890604401602060405180830381865afa158015611d1f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061112b9190613c5a565b80545f908015611d7f57611d6983611d5c600184613e00565b5f91825260209091200190565b54600160301b90046001600160d01b031661112e565b5f9392505050565b600954604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a1600980546001600160a01b0319166001600160a01b0392909216919091179055565b606082611e0557611e00826127ba565b61099f565b508061099f565b5f5f5f611e1986866127e2565b91509150815f03611e3d57838181611e3357611e33613e13565b049250505061112e565b818411611e5457611e54600385150260111861210b565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010185841190960395909502919093039390930492909217029150509392505050565b80545f906001600160801b0380821691600160801b9004168103611ee457611ee4603161210b565b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b5f6001600160d01b03821115611f5c576040516306dfcc6560e41b815260d0600482015260248101839052604401610d2f565b5090565b5f80611f6d8585856127fe565b915091505b935093915050565b5f5f60095f9054906101000a90046001600160a01b03166001600160a01b031663f27a0c926040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fcc573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ff09190613c5a565b90505f3060601b6bffffffffffffffffffffffff1916841860095460405163b1c5f42760e01b81529192506001600160a01b03169063b1c5f42790612041908a908a908a905f908890600401613e27565b602060405180830381865afa15801561205c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120809190613c5a565b5f898152600a60205260408082209290925560095491516308f2a0bb60e41b81526001600160a01b0390921691638f2a0bb0916120ca918b918b918b919088908a90600401613e74565b5f604051808303815f87803b1580156120e1575f5ffd5b505af11580156120f3573d5f5f3e3d5ffd5b50505050610f0c82426121069190613d12565b612564565b634e487b715f52806020526024601cfd5b6009546001600160a01b031663e38335e5348686865f3060601b6bffffffffffffffffffffffff191688186040518763ffffffff1660e01b8152600401612167959493929190613e27565b5f604051808303818588803b15801561217e575f5ffd5b505af1158015612190573d5f5f3e3d5ffd5b5050505f9687525050600a602052505060408320929092555050565b5f818152600460205260408120805460ff600160f01b8204811691600160f81b90041681156121e057506007949350505050565b80156121f157506002949350505050565b5f6121fb86610cbf565b9050805f0361222057604051636ad0607560e01b815260048101879052602401610d2f565b5f612229611098565b65ffffffffffff16905080821061224657505f9695505050505050565b5f612250886111b7565b905081811061226757506001979650505050505050565b6122708861294e565b158061228f57505f888152600760205260409020805460019091015411155b156122a257506003979650505050505050565b5f8881526004602052604090206001015465ffffffffffff165f036122cf57506004979650505050505050565b506005979650505050505050565b5f5f6122eb86868686612984565b5f818152600a60205260409020549091508015610b4d5760095460405163c4d252f560e01b8152600481018390526001600160a01b039091169063c4d252f5906024015f604051808303815f87803b158015612345575f5ffd5b505af1158015612357573d5f5f3e3d5ffd5b5050505f838152600a602052604081205550509050949350505050565b5f61099f612380612a33565b8360405161190160f01b8152600281019290925260228201526042902090565b5f836001600160a01b03163b5f036123fe575f5f6123be8585612b5c565b5090925090505f8160038111156123d7576123d761345f565b1480156123f55750856001600160a01b0316826001600160a01b0316145b9250505061112e565b612409848484612ba5565b905061112e565b5f8581526007602090815260408083206001600160a01b03881684526003810190925282205460ff1615612462576040516371c6af4960e01b81526001600160a01b0387166004820152602401610d2f565b6001600160a01b0386165f9081526003820160205260409020805460ff1916600117905560ff85166124ab5783815f015f8282546124a09190613d12565b909155506125019050565b5f1960ff8616016124c95783816001015f8282546124a09190613d12565b60011960ff8616016124e85783816002015f8282546124a09190613d12565b6040516303599be160e11b815260040160405180910390fd5b509195945050505050565b80545f908190819080820361252a575f5f5f93509350935050610e2d565b5f61253a86611d5c600185613e00565b546001955065ffffffffffff81169450600160301b90046001600160d01b03169250610e2d915050565b5f65ffffffffffff821115611f5c576040516306dfcc6560e41b81526030600482015260248101839052604401610d2f565b81545f90818160058111156125f2575f6125af84612c7b565b6125b99085613e00565b5f8881526020902090915081015465ffffffffffff90811690871610156125e2578091506125f0565b6125ed816001613d12565b92505b505b5f6125ff87878585612dd3565b9050801561262c5761261687611d5c600184613e00565b54600160301b90046001600160d01b0316610d44565b505f9695505050505050565b5f5f845183118061264857508284115b1561265757505f905080611f72565b5f612663856001613d12565b8411801561268b575061060f60f31b61267f8787016020015190565b6001600160f01b031916145b90505f61269b8215156002613ecb565b6126a6906028613d12565b9050806126b38787613e00565b036126d4575f5f6126c5898989612e32565b9096509450611f729350505050565b5f5f935093505050611f72565b5f63ffffffff821115611f5c576040516306dfcc6560e41b81526020600482015260248101839052604401610d2f565b606060ff831461272b5761272483612ef3565b905061099f565b81805461273790613a57565b80601f016020809104026020016040519081016040528092919081815260200182805461276390613a57565b80156127ae5780601f10612785576101008083540402835291602001916127ae565b820191905f5260205f20905b81548152906001019060200180831161279157829003601f168201915b5050505050905061099f565b8051156127c957805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f805f1983850993909202808410938190039390930393915050565b82545f90819080156128f4575f61281a87611d5c600185613e00565b805490915065ffffffffffff80821691600160301b90046001600160d01b031690881682111561285d57604051632520601d60e01b815260040160405180910390fd5b8765ffffffffffff168265ffffffffffff160361289657825465ffffffffffff16600160301b6001600160d01b038916021783556128e6565b6040805180820190915265ffffffffffff808a1682526001600160d01b03808a1660208085019182528d54600181018f555f8f81529190912094519151909216600160301b029216919091179101555b9450859350611f7292505050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a555f8a815291822095519251909316600160301b029190931617920191909155905081611f72565b5f8181526007602052604081206002810154600182015461296f9190613d12565b61297b6108ea85610cbf565b11159392505050565b5f5f6129928686868661115f565b90506129e0816129a260076114bb565b6129ac60066114bb565b6129b660026114bb565b60016129c3600782613ee2565b6129ce906002613fd6565b6129d89190613e00565b1818186114dd565b505f818152600460205260409081902080546001600160f81b0316600160f81b179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c90610cae9083815260200190565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015612a8b57507f000000000000000000000000000000000000000000000000000000000000000046145b15612ab557507f000000000000000000000000000000000000000000000000000000000000000090565b610934604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f5f5f8351604103612b93576020840151604085015160608601515f1a612b8588828585612f30565b955095509550505050612b9e565b505081515f91506002905b9250925092565b5f5f5f856001600160a01b03168585604051602401612bc5929190613fe4565b60408051601f198184030181529181526020820180516001600160e01b0316630b135d3f60e11b17905251612bfa9190613ffc565b5f60405180830381855afa9150503d805f8114612c32576040519150601f19603f3d011682016040523d82523d5f602084013e612c37565b606091505b5091509150818015612c4b57506020815110155b8015610fb257508051630b135d3f60e11b90612c709083016020908101908401613c5a565b149695505050505050565b5f60018211612c88575090565b816001600160801b8210612ca15760809190911c9060401b5b680100000000000000008210612cbc5760409190911c9060201b5b6401000000008210612cd35760209190911c9060101b5b620100008210612ce85760109190911c9060081b5b6101008210612cfc5760089190911c9060041b5b60108210612d0f5760049190911c9060021b5b60048210612d1b5760011b5b600302600190811c90818581612d3357612d33613e13565b048201901c90506001818581612d4b57612d4b613e13565b048201901c90506001818581612d6357612d63613e13565b048201901c90506001818581612d7b57612d7b613e13565b048201901c90506001818581612d9357612d93613e13565b048201901c90506001818581612dab57612dab613e13565b048201901c9050612dca818581612dc457612dc4613e13565b04821190565b90039392505050565b5f5b81831015612e2a575f612de88484612ff8565b5f8781526020902090915065ffffffffffff86169082015465ffffffffffff161115612e1657809250612e24565b612e21816001613d12565b93505b50612dd5565b509392505050565b5f808481612e41866001613d12565b85118015612e69575061060f60f31b612e5d8388016020015190565b6001600160f01b031916145b90505f612e798215156002613ecb565b90505f80612e87838a613d12565b90505b87811015612ee2575f612ea8612ea38784016020015190565b613012565b9050600f8160ff161115612ec7575f5f97509750505050505050611f72565b612ed2601084613ecb565b60ff909116019150600101612e8a565b506001999098509650505050505050565b60605f612eff8361308a565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115612f6957505f91506003905082612fee565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015612fba573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116612fe557505f925060019150829050612fee565b92505f91508190505b9450945094915050565b5f6130066002848418614012565b61112e90848416613d12565b5f60f882901c602f8111801561302b5750603a8160ff16105b1561303957602f190161099f565b60608160ff1611801561304f575060678160ff16105b1561305d576056190161099f565b60408160ff16118015613073575060478160ff16105b15613081576036190161099f565b5060ff92915050565b5f60ff8216601f81111561099f57604051632cd44ac360e21b815260040160405180910390fd5b5f602082840312156130c1575f5ffd5b81356001600160e01b03198116811461112e575f5ffd5b5f602082840312156130e8575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61112e60208301846130ef565b6001600160a01b03811681146109b6575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561317f5761317f613143565b604052919050565b5f6001600160401b0382111561319f5761319f613143565b50601f01601f191660200190565b5f6131bf6131ba84613187565b613157565b90508281528383830111156131d2575f5ffd5b828260208301375f602084830101529392505050565b5f82601f8301126131f7575f5ffd5b61112e838335602085016131ad565b5f5f5f5f60808587031215613219575f5ffd5b84356132248161312f565b935060208501356132348161312f565b92506040850135915060608501356001600160401b03811115613255575f5ffd5b613261878288016131e8565b91505092959194509250565b5f6001600160401b0382111561328557613285613143565b5060051b60200190565b5f82601f83011261329e575f5ffd5b81356132ac6131ba8261326d565b8082825260208201915060208360051b8601019250858311156132cd575f5ffd5b602085015b83811015610b4d5780356132e58161312f565b8352602092830192016132d2565b5f82601f830112613302575f5ffd5b81356133106131ba8261326d565b8082825260208201915060208360051b860101925085831115613331575f5ffd5b602085015b83811015610b4d578035835260209283019201613336565b5f82601f83011261335d575f5ffd5b813561336b6131ba8261326d565b8082825260208201915060208360051b86010192508583111561338c575f5ffd5b602085015b83811015610b4d5780356001600160401b038111156133ae575f5ffd5b6133bd886020838a01016131e8565b84525060209283019201613391565b5f5f5f5f608085870312156133df575f5ffd5b84356001600160401b038111156133f4575f5ffd5b6134008782880161328f565b94505060208501356001600160401b0381111561341b575f5ffd5b613427878288016132f3565b93505060408501356001600160401b03811115613442575f5ffd5b61344e8782880161334e565b949793965093946060013593505050565b634e487b7160e01b5f52602160045260245ffd5b6008811061348f57634e487b7160e01b5f52602160045260245ffd5b9052565b6020810161099f8284613473565b5f5f604083850312156134b2575f5ffd5b8235915060208301356134c48161312f565b809150509250929050565b803560ff81168114610e0a575f5ffd5b5f5f604083850312156134f0575f5ffd5b82359150613500602084016134cf565b90509250929050565b5f5f83601f840112613519575f5ffd5b5081356001600160401b0381111561352f575f5ffd5b602083019150836020828501011115613546575f5ffd5b9250929050565b5f5f5f5f5f5f5f60c0888a031215613563575f5ffd5b87359650613573602089016134cf565b955060408801356135838161312f565b945060608801356001600160401b0381111561359d575f5ffd5b6135a98a828b01613509565b90955093505060808801356001600160401b038111156135c7575f5ffd5b6135d38a828b016131e8565b92505060a08801356001600160401b038111156135ee575f5ffd5b6135fa8a828b016131e8565b91505092959891949750929550565b5f5f5f5f5f6080868803121561361d575f5ffd5b8535945061362d602087016134cf565b935060408601356001600160401b03811115613647575f5ffd5b61365388828901613509565b90945092505060608601356001600160401b03811115613671575f5ffd5b61367d888289016131e8565b9150509295509295909350565b5f5f5f5f6060858703121561369d575f5ffd5b843593506136ad602086016134cf565b925060408501356001600160401b038111156136c7575f5ffd5b6136d387828801613509565b95989497509550505050565b5f5f5f5f608085870312156136f2575f5ffd5b84356001600160401b03811115613707575f5ffd5b6137138782880161328f565b94505060208501356001600160401b0381111561372e575f5ffd5b61373a878288016132f3565b93505060408501356001600160401b03811115613755575f5ffd5b6137618782880161334e565b92505060608501356001600160401b0381111561377c575f5ffd5b8501601f8101871361378c575f5ffd5b613261878235602084016131ad565b5f602082840312156137ab575f5ffd5b813561112e8161312f565b5f8151808452602084019350602083015f5b828110156137e65781518652602095860195909101906001016137c8565b5093949350505050565b60ff60f81b8816815260e060208201525f61380e60e08301896130ef565b828103604084015261382081896130ef565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152905061385181856137b6565b9a9950505050505050505050565b5f5f5f5f60808587031215613872575f5ffd5b84359350613882602086016134cf565b925060408501356138928161312f565b915060608501356001600160401b03811115613255575f5ffd5b5f5f5f606084860312156138be575f5ffd5b83356138c98161312f565b92506020840135915060408401356001600160401b038111156138ea575f5ffd5b6138f6868287016131e8565b9150509250925092565b5f5f5f5f5f60a08688031215613914575f5ffd5b853561391f8161312f565b9450602086013561392f8161312f565b935060408601356001600160401b03811115613949575f5ffd5b613955888289016132f3565b93505060608601356001600160401b03811115613970575f5ffd5b61397c888289016132f3565b92505060808601356001600160401b03811115613671575f5ffd5b5f5f5f5f606085870312156139aa575f5ffd5b84356139b58161312f565b93506020850135925060408501356001600160401b038111156136c7575f5ffd5b5f5f604083850312156139e7575f5ffd5b82356139f28161312f565b946020939093013593505050565b5f5f5f5f5f60a08688031215613a14575f5ffd5b8535613a1f8161312f565b94506020860135613a2f8161312f565b9350604086013592506060860135915060808601356001600160401b03811115613671575f5ffd5b600181811c90821680613a6b57607f821691505b602082108103613a8957634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215613ab3575f5ffd5b81516001600160401b03811115613ac8575f5ffd5b8201601f81018413613ad8575f5ffd5b8051613ae66131ba82613187565b818152856020838501011115613afa575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b634e487b7160e01b5f52601160045260245ffd5b5f60208284031215613b3b575f5ffd5b815165ffffffffffff8116811461112e575f5ffd5b65ffffffffffff818116838216019081111561099f5761099f613b17565b818382375f9101908152919050565b5f8151808452602084019350602083015f5b828110156137e65781516001600160a01b0316865260209586019590910190600101613b8f565b5f82825180855260208501945060208160051b830101602085015f5b83811015613c0457601f19858403018852613bee8383516130ef565b6020988901989093509190910190600101613bd2565b50909695505050505050565b608081525f613c226080830187613b7d565b8281036020840152613c3481876137b6565b90508281036040840152613c488186613bb6565b91505082606083015295945050505050565b5f60208284031215613c6a575f5ffd5b5051919050565b83815260608101613c856020830185613473565b826040830152949350505050565b5f60208284031215613ca3575f5ffd5b8151801515811461112e575f5ffd5b84815260ff84166020820152826040820152608060608201525f610fb260808301846130ef565b85815260ff8516602082015283604082015260a060608201525f613d0060a08301856130ef565b8281036080840152610f0c81856130ef565b8082018082111561099f5761099f613b17565b8981526001600160a01b0389166020820152610120604082018190525f90613d4f9083018a613b7d565b8281036060840152613d61818a6137b6565b9050828103608084015280885180835260208301915060208160051b84010160208b015f5b83811015613db857601f19868403018552613da28383516130ef565b6020958601959093509190910190600101613d86565b505085810360a0870152613dcc818b613bb6565b93505050508560c08401528460e0840152828103610100840152613df081856130ef565b9c9b505050505050505050505050565b8181038181111561099f5761099f613b17565b634e487b7160e01b5f52601260045260245ffd5b60a081525f613e3960a0830188613b7d565b8281036020840152613e4b81886137b6565b90508281036040840152613e5f8187613bb6565b60608401959095525050608001529392505050565b60c081525f613e8660c0830189613b7d565b8281036020840152613e9881896137b6565b90508281036040840152613eac8188613bb6565b60608401969096525050608081019290925260a0909101529392505050565b808202811582820484141761099f5761099f613b17565b60ff818116838216019081111561099f5761099f613b17565b6001815b6001841115611f7257808504811115613f1a57613f1a613b17565b6001841615613f2857908102905b60019390931c928002613eff565b5f82613f445750600161099f565b81613f5057505f61099f565b8160018114613f665760028114613f7057613f8c565b600191505061099f565b60ff841115613f8157613f81613b17565b50506001821b61099f565b5060208310610133831016604e8410600b8410161715613faf575081810a61099f565b613fbb5f198484613efb565b805f1904821115613fce57613fce613b17565b029392505050565b5f61112e60ff841683613f36565b828152604060208201525f61112b60408301846130ef565b5f82518060208501845e5f920191825250919050565b5f8261402c57634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220a606f8db09abd57c02d0c6fddb8f7d5920856040c94fa564c18993742f01560c64736f6c634300081e0033",
}

// MyGovernor is an auto generated Go binding around an Ethereum contract.
type MyGovernor struct {
	abi abi.ABI
}

// NewMyGovernor creates a new instance of MyGovernor.
func NewMyGovernor() *MyGovernor {
	parsed, err := MyGovernorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MyGovernor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MyGovernor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _token, address _timelock) returns()
func (myGovernor *MyGovernor) PackConstructor(_token common.Address, _timelock common.Address) []byte {
	enc, err := myGovernor.abi.Pack("", _token, _timelock)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdeaaa7cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) PackBALLOTTYPEHASH() []byte {
	enc, err := myGovernor.abi.Pack("BALLOT_TYPEHASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdeaaa7cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) TryPackBALLOTTYPEHASH() ([]byte, error) {
	return myGovernor.abi.Pack("BALLOT_TYPEHASH")
}

// UnpackBALLOTTYPEHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) UnpackBALLOTTYPEHASH(data []byte) ([32]byte, error) {
	out, err := myGovernor.abi.Unpack("BALLOT_TYPEHASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackCLOCKMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4bf5d7e9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (myGovernor *MyGovernor) PackCLOCKMODE() []byte {
	enc, err := myGovernor.abi.Pack("CLOCK_MODE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCLOCKMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4bf5d7e9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (myGovernor *MyGovernor) TryPackCLOCKMODE() ([]byte, error) {
	return myGovernor.abi.Pack("CLOCK_MODE")
}

// UnpackCLOCKMODE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (myGovernor *MyGovernor) UnpackCLOCKMODE(data []byte) (string, error) {
	out, err := myGovernor.abi.Unpack("CLOCK_MODE", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackCOUNTINGMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd4e2ba5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (myGovernor *MyGovernor) PackCOUNTINGMODE() []byte {
	enc, err := myGovernor.abi.Pack("COUNTING_MODE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCOUNTINGMODE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd4e2ba5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (myGovernor *MyGovernor) TryPackCOUNTINGMODE() ([]byte, error) {
	return myGovernor.abi.Pack("COUNTING_MODE")
}

// UnpackCOUNTINGMODE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (myGovernor *MyGovernor) UnpackCOUNTINGMODE(data []byte) (string, error) {
	out, err := myGovernor.abi.Unpack("COUNTING_MODE", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackEXTENDEDBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2fe3e261.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) PackEXTENDEDBALLOTTYPEHASH() []byte {
	enc, err := myGovernor.abi.Pack("EXTENDED_BALLOT_TYPEHASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEXTENDEDBALLOTTYPEHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2fe3e261.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) TryPackEXTENDEDBALLOTTYPEHASH() ([]byte, error) {
	return myGovernor.abi.Pack("EXTENDED_BALLOT_TYPEHASH")
}

// UnpackEXTENDEDBALLOTTYPEHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (myGovernor *MyGovernor) UnpackEXTENDEDBALLOTTYPEHASH(data []byte) ([32]byte, error) {
	out, err := myGovernor.abi.Unpack("EXTENDED_BALLOT_TYPEHASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x452115d6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) PackCancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := myGovernor.abi.Pack("cancel", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x452115d6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) TryPackCancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) ([]byte, error) {
	return myGovernor.abi.Pack("cancel", targets, values, calldatas, descriptionHash)
}

// UnpackCancel is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) UnpackCancel(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("cancel", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCastVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56781388.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (myGovernor *MyGovernor) PackCastVote(proposalId *big.Int, support uint8) []byte {
	enc, err := myGovernor.abi.Pack("castVote", proposalId, support)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56781388.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (myGovernor *MyGovernor) TryPackCastVote(proposalId *big.Int, support uint8) ([]byte, error) {
	return myGovernor.abi.Pack("castVote", proposalId, support)
}

// UnpackCastVote is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (myGovernor *MyGovernor) UnpackCastVote(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("castVote", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCastVoteBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ff262e3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) PackCastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) []byte {
	enc, err := myGovernor.abi.Pack("castVoteBySig", proposalId, support, voter, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVoteBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ff262e3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) TryPackCastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) ([]byte, error) {
	return myGovernor.abi.Pack("castVoteBySig", proposalId, support, voter, signature)
}

// UnpackCastVoteBySig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) UnpackCastVoteBySig(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("castVoteBySig", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCastVoteWithReason is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b3c71d3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (myGovernor *MyGovernor) PackCastVoteWithReason(proposalId *big.Int, support uint8, reason string) []byte {
	enc, err := myGovernor.abi.Pack("castVoteWithReason", proposalId, support, reason)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVoteWithReason is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7b3c71d3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (myGovernor *MyGovernor) TryPackCastVoteWithReason(proposalId *big.Int, support uint8, reason string) ([]byte, error) {
	return myGovernor.abi.Pack("castVoteWithReason", proposalId, support, reason)
}

// UnpackCastVoteWithReason is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (myGovernor *MyGovernor) UnpackCastVoteWithReason(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("castVoteWithReason", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCastVoteWithReasonAndParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f398a14.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (myGovernor *MyGovernor) PackCastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) []byte {
	enc, err := myGovernor.abi.Pack("castVoteWithReasonAndParams", proposalId, support, reason, params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVoteWithReasonAndParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f398a14.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (myGovernor *MyGovernor) TryPackCastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) ([]byte, error) {
	return myGovernor.abi.Pack("castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// UnpackCastVoteWithReasonAndParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (myGovernor *MyGovernor) UnpackCastVoteWithReasonAndParams(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("castVoteWithReasonAndParams", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCastVoteWithReasonAndParamsBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b8d0e0d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) PackCastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) []byte {
	enc, err := myGovernor.abi.Pack("castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCastVoteWithReasonAndParamsBySig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b8d0e0d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) TryPackCastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) ([]byte, error) {
	return myGovernor.abi.Pack("castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
}

// UnpackCastVoteWithReasonAndParamsBySig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (myGovernor *MyGovernor) UnpackCastVoteWithReasonAndParamsBySig(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("castVoteWithReasonAndParamsBySig", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackClock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91ddadf4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clock() view returns(uint48)
func (myGovernor *MyGovernor) PackClock() []byte {
	enc, err := myGovernor.abi.Pack("clock")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91ddadf4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clock() view returns(uint48)
func (myGovernor *MyGovernor) TryPackClock() ([]byte, error) {
	return myGovernor.abi.Pack("clock")
}

// UnpackClock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (myGovernor *MyGovernor) UnpackClock(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("clock", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (myGovernor *MyGovernor) PackEip712Domain() []byte {
	enc, err := myGovernor.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (myGovernor *MyGovernor) TryPackEip712Domain() ([]byte, error) {
	return myGovernor.abi.Pack("eip712Domain")
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
// type Eip712DomainOutput struct {
// 	Fields            [1]byte
// 	Name              string
// 	Version           string
// 	ChainId           *big.Int
// 	VerifyingContract common.Address
// 	Salt              [32]byte
// 	Extensions        []*big.Int
// }

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (myGovernor *MyGovernor) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := myGovernor.abi.Unpack("eip712Domain", data)
	outstruct := new(Eip712DomainOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, nil
}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2656227d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (myGovernor *MyGovernor) PackExecute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := myGovernor.abi.Pack("execute", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2656227d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (myGovernor *MyGovernor) TryPackExecute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) ([]byte, error) {
	return myGovernor.abi.Pack("execute", targets, values, calldatas, descriptionHash)
}

// UnpackExecute is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (myGovernor *MyGovernor) UnpackExecute(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("execute", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetProposalId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8f8a668.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (myGovernor *MyGovernor) PackGetProposalId(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := myGovernor.abi.Pack("getProposalId", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProposalId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8f8a668.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (myGovernor *MyGovernor) TryPackGetProposalId(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) ([]byte, error) {
	return myGovernor.abi.Pack("getProposalId", targets, values, calldatas, descriptionHash)
}

// UnpackGetProposalId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa8f8a668.
//
// Solidity: function getProposalId(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) view returns(uint256)
func (myGovernor *MyGovernor) UnpackGetProposalId(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("getProposalId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb9019d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) PackGetVotes(account common.Address, timepoint *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("getVotes", account, timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb9019d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) TryPackGetVotes(account common.Address, timepoint *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("getVotes", account, timepoint)
}

// UnpackGetVotes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) UnpackGetVotes(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("getVotes", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetVotesWithParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a802a6d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (myGovernor *MyGovernor) PackGetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) []byte {
	enc, err := myGovernor.abi.Pack("getVotesWithParams", account, timepoint, params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVotesWithParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a802a6d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (myGovernor *MyGovernor) TryPackGetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) ([]byte, error) {
	return myGovernor.abi.Pack("getVotesWithParams", account, timepoint, params)
}

// UnpackGetVotesWithParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (myGovernor *MyGovernor) UnpackGetVotesWithParams(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("getVotesWithParams", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackHasVoted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43859632.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (myGovernor *MyGovernor) PackHasVoted(proposalId *big.Int, account common.Address) []byte {
	enc, err := myGovernor.abi.Pack("hasVoted", proposalId, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHasVoted is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43859632.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (myGovernor *MyGovernor) TryPackHasVoted(proposalId *big.Int, account common.Address) ([]byte, error) {
	return myGovernor.abi.Pack("hasVoted", proposalId, account)
}

// UnpackHasVoted is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (myGovernor *MyGovernor) UnpackHasVoted(data []byte) (bool, error) {
	out, err := myGovernor.abi.Unpack("hasVoted", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHashProposal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc59057e4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (myGovernor *MyGovernor) PackHashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := myGovernor.abi.Pack("hashProposal", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHashProposal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc59057e4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (myGovernor *MyGovernor) TryPackHashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) ([]byte, error) {
	return myGovernor.abi.Pack("hashProposal", targets, values, calldatas, descriptionHash)
}

// UnpackHashProposal is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (myGovernor *MyGovernor) UnpackHashProposal(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("hashProposal", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name() view returns(string)
func (myGovernor *MyGovernor) PackName() []byte {
	enc, err := myGovernor.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function name() view returns(string)
func (myGovernor *MyGovernor) TryPackName() ([]byte, error) {
	return myGovernor.abi.Pack("name")
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (myGovernor *MyGovernor) UnpackName(data []byte) (string, error) {
	out, err := myGovernor.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (myGovernor *MyGovernor) PackNonces(owner common.Address) []byte {
	enc, err := myGovernor.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (myGovernor *MyGovernor) TryPackNonces(owner common.Address) ([]byte, error) {
	return myGovernor.abi.Pack("nonces", owner)
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (myGovernor *MyGovernor) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOnERC1155BatchReceived is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc197c81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) PackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) []byte {
	enc, err := myGovernor.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
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
func (myGovernor *MyGovernor) TryPackOnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([]byte, error) {
	return myGovernor.abi.Pack("onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155BatchReceived is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) UnpackOnERC1155BatchReceived(data []byte) ([4]byte, error) {
	out, err := myGovernor.abi.Unpack("onERC1155BatchReceived", data)
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
func (myGovernor *MyGovernor) PackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) []byte {
	enc, err := myGovernor.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
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
func (myGovernor *MyGovernor) TryPackOnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([]byte, error) {
	return myGovernor.abi.Pack("onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// UnpackOnERC1155Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) UnpackOnERC1155Received(data []byte) ([4]byte, error) {
	out, err := myGovernor.abi.Unpack("onERC1155Received", data)
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
func (myGovernor *MyGovernor) PackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) []byte {
	enc, err := myGovernor.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
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
func (myGovernor *MyGovernor) TryPackOnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([]byte, error) {
	return myGovernor.abi.Pack("onERC721Received", arg0, arg1, arg2, arg3)
}

// UnpackOnERC721Received is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (myGovernor *MyGovernor) UnpackOnERC721Received(data []byte) ([4]byte, error) {
	out, err := myGovernor.abi.Unpack("onERC721Received", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, nil
}

// PackProposalDeadline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc01f9e37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) PackProposalDeadline(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalDeadline", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalDeadline is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc01f9e37.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) TryPackProposalDeadline(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalDeadline", proposalId)
}

// UnpackProposalDeadline is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) UnpackProposalDeadline(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("proposalDeadline", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalEta is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab58fb8e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) PackProposalEta(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalEta", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalEta is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab58fb8e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) TryPackProposalEta(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalEta", proposalId)
}

// UnpackProposalEta is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) UnpackProposalEta(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("proposalEta", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalNeedsQueuing is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9a95294.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (myGovernor *MyGovernor) PackProposalNeedsQueuing(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalNeedsQueuing", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalNeedsQueuing is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9a95294.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (myGovernor *MyGovernor) TryPackProposalNeedsQueuing(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalNeedsQueuing", proposalId)
}

// UnpackProposalNeedsQueuing is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (myGovernor *MyGovernor) UnpackProposalNeedsQueuing(data []byte) (bool, error) {
	out, err := myGovernor.abi.Unpack("proposalNeedsQueuing", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProposalProposer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x143489d0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (myGovernor *MyGovernor) PackProposalProposer(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalProposer", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalProposer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x143489d0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (myGovernor *MyGovernor) TryPackProposalProposer(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalProposer", proposalId)
}

// UnpackProposalProposer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (myGovernor *MyGovernor) UnpackProposalProposer(data []byte) (common.Address, error) {
	out, err := myGovernor.abi.Unpack("proposalProposer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackProposalSnapshot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d63f693.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) PackProposalSnapshot(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalSnapshot", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalSnapshot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d63f693.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) TryPackProposalSnapshot(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalSnapshot", proposalId)
}

// UnpackProposalSnapshot is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (myGovernor *MyGovernor) UnpackProposalSnapshot(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("proposalSnapshot", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb58131b0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (myGovernor *MyGovernor) PackProposalThreshold() []byte {
	enc, err := myGovernor.abi.Pack("proposalThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb58131b0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (myGovernor *MyGovernor) TryPackProposalThreshold() ([]byte, error) {
	return myGovernor.abi.Pack("proposalThreshold")
}

// UnpackProposalThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb58131b0.
//
// Solidity: function proposalThreshold() pure returns(uint256)
func (myGovernor *MyGovernor) UnpackProposalThreshold(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("proposalThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposalVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x544ffc9c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (myGovernor *MyGovernor) PackProposalVotes(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("proposalVotes", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposalVotes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x544ffc9c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (myGovernor *MyGovernor) TryPackProposalVotes(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("proposalVotes", proposalId)
}

// ProposalVotesOutput serves as a container for the return parameters of contract
// method ProposalVotes.
type ProposalVotesOutput struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}

// UnpackProposalVotes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (myGovernor *MyGovernor) UnpackProposalVotes(data []byte) (ProposalVotesOutput, error) {
	out, err := myGovernor.abi.Unpack("proposalVotes", data)
	outstruct := new(ProposalVotesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AgainstVotes = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.ForVotes = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.AbstainVotes = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackPropose is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d5e81e2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (myGovernor *MyGovernor) PackPropose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) []byte {
	enc, err := myGovernor.abi.Pack("propose", targets, values, calldatas, description)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPropose is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d5e81e2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (myGovernor *MyGovernor) TryPackPropose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) ([]byte, error) {
	return myGovernor.abi.Pack("propose", targets, values, calldatas, description)
}

// UnpackPropose is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (myGovernor *MyGovernor) UnpackPropose(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("propose", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQueue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x160cbed7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) PackQueue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) []byte {
	enc, err := myGovernor.abi.Pack("queue", targets, values, calldatas, descriptionHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQueue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x160cbed7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) TryPackQueue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) ([]byte, error) {
	return myGovernor.abi.Pack("queue", targets, values, calldatas, descriptionHash)
}

// UnpackQueue is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (myGovernor *MyGovernor) UnpackQueue(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("queue", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQuorum is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ce560a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) PackQuorum(timepoint *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("quorum", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuorum is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8ce560a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) TryPackQuorum(timepoint *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("quorum", timepoint)
}

// UnpackQuorum is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) UnpackQuorum(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("quorum", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQuorumDenominator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x97c3d334.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (myGovernor *MyGovernor) PackQuorumDenominator() []byte {
	enc, err := myGovernor.abi.Pack("quorumDenominator")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuorumDenominator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x97c3d334.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (myGovernor *MyGovernor) TryPackQuorumDenominator() ([]byte, error) {
	return myGovernor.abi.Pack("quorumDenominator")
}

// UnpackQuorumDenominator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (myGovernor *MyGovernor) UnpackQuorumDenominator(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("quorumDenominator", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQuorumNumerator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60c4247f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) PackQuorumNumerator(timepoint *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("quorumNumerator", timepoint)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuorumNumerator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60c4247f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) TryPackQuorumNumerator(timepoint *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("quorumNumerator", timepoint)
}

// UnpackQuorumNumerator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (myGovernor *MyGovernor) UnpackQuorumNumerator(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("quorumNumerator", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackQuorumNumerator0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7713a70.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (myGovernor *MyGovernor) PackQuorumNumerator0() []byte {
	enc, err := myGovernor.abi.Pack("quorumNumerator0")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuorumNumerator0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7713a70.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (myGovernor *MyGovernor) TryPackQuorumNumerator0() ([]byte, error) {
	return myGovernor.abi.Pack("quorumNumerator0")
}

// UnpackQuorumNumerator0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (myGovernor *MyGovernor) UnpackQuorumNumerator0(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("quorumNumerator0", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc28bc2fa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (myGovernor *MyGovernor) PackRelay(target common.Address, value *big.Int, data []byte) []byte {
	enc, err := myGovernor.abi.Pack("relay", target, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc28bc2fa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (myGovernor *MyGovernor) TryPackRelay(target common.Address, value *big.Int, data []byte) ([]byte, error) {
	return myGovernor.abi.Pack("relay", target, value, data)
}

// PackState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e4f49e6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (myGovernor *MyGovernor) PackState(proposalId *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("state", proposalId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e4f49e6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (myGovernor *MyGovernor) TryPackState(proposalId *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("state", proposalId)
}

// UnpackState is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (myGovernor *MyGovernor) UnpackState(data []byte) (uint8, error) {
	out, err := myGovernor.abi.Unpack("state", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (myGovernor *MyGovernor) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := myGovernor.abi.Pack("supportsInterface", interfaceId)
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
func (myGovernor *MyGovernor) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return myGovernor.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (myGovernor *MyGovernor) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := myGovernor.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd33219b4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelock() view returns(address)
func (myGovernor *MyGovernor) PackTimelock() []byte {
	enc, err := myGovernor.abi.Pack("timelock")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd33219b4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function timelock() view returns(address)
func (myGovernor *MyGovernor) TryPackTimelock() ([]byte, error) {
	return myGovernor.abi.Pack("timelock")
}

// UnpackTimelock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (myGovernor *MyGovernor) UnpackTimelock(data []byte) (common.Address, error) {
	out, err := myGovernor.abi.Unpack("timelock", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc0c546a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function token() view returns(address)
func (myGovernor *MyGovernor) PackToken() []byte {
	enc, err := myGovernor.abi.Pack("token")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc0c546a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function token() view returns(address)
func (myGovernor *MyGovernor) TryPackToken() ([]byte, error) {
	return myGovernor.abi.Pack("token")
}

// UnpackToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (myGovernor *MyGovernor) UnpackToken(data []byte) (common.Address, error) {
	out, err := myGovernor.abi.Unpack("token", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackUpdateQuorumNumerator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06f3f9e6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (myGovernor *MyGovernor) PackUpdateQuorumNumerator(newQuorumNumerator *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("updateQuorumNumerator", newQuorumNumerator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateQuorumNumerator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06f3f9e6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (myGovernor *MyGovernor) TryPackUpdateQuorumNumerator(newQuorumNumerator *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("updateQuorumNumerator", newQuorumNumerator)
}

// PackUpdateTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa890c910.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (myGovernor *MyGovernor) PackUpdateTimelock(newTimelock common.Address) []byte {
	enc, err := myGovernor.abi.Pack("updateTimelock", newTimelock)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateTimelock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa890c910.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (myGovernor *MyGovernor) TryPackUpdateTimelock(newTimelock common.Address) ([]byte, error) {
	return myGovernor.abi.Pack("updateTimelock", newTimelock)
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(string)
func (myGovernor *MyGovernor) PackVersion() []byte {
	enc, err := myGovernor.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function version() view returns(string)
func (myGovernor *MyGovernor) TryPackVersion() ([]byte, error) {
	return myGovernor.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (myGovernor *MyGovernor) UnpackVersion(data []byte) (string, error) {
	out, err := myGovernor.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3932abb1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function votingDelay() pure returns(uint256)
func (myGovernor *MyGovernor) PackVotingDelay() []byte {
	enc, err := myGovernor.abi.Pack("votingDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3932abb1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function votingDelay() pure returns(uint256)
func (myGovernor *MyGovernor) TryPackVotingDelay() ([]byte, error) {
	return myGovernor.abi.Pack("votingDelay")
}

// UnpackVotingDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (myGovernor *MyGovernor) UnpackVotingDelay(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("votingDelay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02a251a3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (myGovernor *MyGovernor) PackVotingPeriod() []byte {
	enc, err := myGovernor.abi.Pack("votingPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02a251a3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (myGovernor *MyGovernor) TryPackVotingPeriod() ([]byte, error) {
	return myGovernor.abi.Pack("votingPeriod")
}

// UnpackVotingPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (myGovernor *MyGovernor) UnpackVotingPeriod(data []byte) (*big.Int, error) {
	out, err := myGovernor.abi.Unpack("votingPeriod", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// MyGovernorEIP712DomainChanged represents a EIP712DomainChanged event raised by the MyGovernor contract.
type MyGovernorEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const MyGovernorEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (MyGovernorEIP712DomainChanged) ContractEventName() string {
	return MyGovernorEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (myGovernor *MyGovernor) UnpackEIP712DomainChangedEvent(log *types.Log) (*MyGovernorEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorProposalCanceled represents a ProposalCanceled event raised by the MyGovernor contract.
type MyGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const MyGovernorProposalCanceledEventName = "ProposalCanceled"

// ContractEventName returns the user-defined event name.
func (MyGovernorProposalCanceled) ContractEventName() string {
	return MyGovernorProposalCanceledEventName
}

// UnpackProposalCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (myGovernor *MyGovernor) UnpackProposalCanceledEvent(log *types.Log) (*MyGovernorProposalCanceled, error) {
	event := "ProposalCanceled"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorProposalCanceled)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorProposalCreated represents a ProposalCreated event raised by the MyGovernor contract.
type MyGovernorProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         *types.Log // Blockchain specific contextual infos
}

const MyGovernorProposalCreatedEventName = "ProposalCreated"

// ContractEventName returns the user-defined event name.
func (MyGovernorProposalCreated) ContractEventName() string {
	return MyGovernorProposalCreatedEventName
}

// UnpackProposalCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (myGovernor *MyGovernor) UnpackProposalCreatedEvent(log *types.Log) (*MyGovernorProposalCreated, error) {
	event := "ProposalCreated"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorProposalCreated)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorProposalExecuted represents a ProposalExecuted event raised by the MyGovernor contract.
type MyGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const MyGovernorProposalExecutedEventName = "ProposalExecuted"

// ContractEventName returns the user-defined event name.
func (MyGovernorProposalExecuted) ContractEventName() string {
	return MyGovernorProposalExecutedEventName
}

// UnpackProposalExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (myGovernor *MyGovernor) UnpackProposalExecutedEvent(log *types.Log) (*MyGovernorProposalExecuted, error) {
	event := "ProposalExecuted"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorProposalExecuted)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorProposalQueued represents a ProposalQueued event raised by the MyGovernor contract.
type MyGovernorProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const MyGovernorProposalQueuedEventName = "ProposalQueued"

// ContractEventName returns the user-defined event name.
func (MyGovernorProposalQueued) ContractEventName() string {
	return MyGovernorProposalQueuedEventName
}

// UnpackProposalQueuedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (myGovernor *MyGovernor) UnpackProposalQueuedEvent(log *types.Log) (*MyGovernorProposalQueued, error) {
	event := "ProposalQueued"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorProposalQueued)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the MyGovernor contract.
type MyGovernorQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                *types.Log // Blockchain specific contextual infos
}

const MyGovernorQuorumNumeratorUpdatedEventName = "QuorumNumeratorUpdated"

// ContractEventName returns the user-defined event name.
func (MyGovernorQuorumNumeratorUpdated) ContractEventName() string {
	return MyGovernorQuorumNumeratorUpdatedEventName
}

// UnpackQuorumNumeratorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (myGovernor *MyGovernor) UnpackQuorumNumeratorUpdatedEvent(log *types.Log) (*MyGovernorQuorumNumeratorUpdated, error) {
	event := "QuorumNumeratorUpdated"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorQuorumNumeratorUpdated)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorTimelockChange represents a TimelockChange event raised by the MyGovernor contract.
type MyGovernorTimelockChange struct {
	OldTimelock common.Address
	NewTimelock common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const MyGovernorTimelockChangeEventName = "TimelockChange"

// ContractEventName returns the user-defined event name.
func (MyGovernorTimelockChange) ContractEventName() string {
	return MyGovernorTimelockChangeEventName
}

// UnpackTimelockChangeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (myGovernor *MyGovernor) UnpackTimelockChangeEvent(log *types.Log) (*MyGovernorTimelockChange, error) {
	event := "TimelockChange"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorTimelockChange)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorVoteCast represents a VoteCast event raised by the MyGovernor contract.
type MyGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        *types.Log // Blockchain specific contextual infos
}

const MyGovernorVoteCastEventName = "VoteCast"

// ContractEventName returns the user-defined event name.
func (MyGovernorVoteCast) ContractEventName() string {
	return MyGovernorVoteCastEventName
}

// UnpackVoteCastEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (myGovernor *MyGovernor) UnpackVoteCastEvent(log *types.Log) (*MyGovernorVoteCast, error) {
	event := "VoteCast"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorVoteCast)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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

// MyGovernorVoteCastWithParams represents a VoteCastWithParams event raised by the MyGovernor contract.
type MyGovernorVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const MyGovernorVoteCastWithParamsEventName = "VoteCastWithParams"

// ContractEventName returns the user-defined event name.
func (MyGovernorVoteCastWithParams) ContractEventName() string {
	return MyGovernorVoteCastWithParamsEventName
}

// UnpackVoteCastWithParamsEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (myGovernor *MyGovernor) UnpackVoteCastWithParamsEvent(log *types.Log) (*MyGovernorVoteCastWithParams, error) {
	event := "VoteCastWithParams"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorVoteCastWithParams)
	if len(log.Data) > 0 {
		if err := myGovernor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range myGovernor.abi.Events[event].Inputs {
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
func (myGovernor *MyGovernor) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["CheckpointUnorderedInsertion"].ID.Bytes()[:4]) {
		return myGovernor.UnpackCheckpointUnorderedInsertionError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return myGovernor.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorAlreadyCastVote"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorAlreadyCastVoteError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorAlreadyQueuedProposal"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorAlreadyQueuedProposalError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorDisabledDeposit"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorDisabledDepositError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInsufficientProposerVotes"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInsufficientProposerVotesError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidProposalLength"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidProposalLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidQuorumFraction"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidQuorumFractionError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidSignature"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidVoteParams"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidVoteParamsError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidVoteType"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidVoteTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorInvalidVotingPeriod"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorInvalidVotingPeriodError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorNonexistentProposal"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorNonexistentProposalError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorNotQueuedProposal"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorNotQueuedProposalError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorOnlyExecutor"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorQueueNotImplemented"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorQueueNotImplementedError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorRestrictedProposer"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorRestrictedProposerError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorUnableToCancel"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorUnableToCancelError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["GovernorUnexpectedProposalState"].ID.Bytes()[:4]) {
		return myGovernor.UnpackGovernorUnexpectedProposalStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return myGovernor.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return myGovernor.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return myGovernor.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], myGovernor.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return myGovernor.UnpackStringTooLongError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// MyGovernorCheckpointUnorderedInsertion represents a CheckpointUnorderedInsertion error raised by the MyGovernor contract.
type MyGovernorCheckpointUnorderedInsertion struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CheckpointUnorderedInsertion()
func MyGovernorCheckpointUnorderedInsertionErrorID() common.Hash {
	return common.HexToHash("0x2520601d9d60b717c34a36ad270857824c5a1ebbfd08ff39aba6930089495cfa")
}

// UnpackCheckpointUnorderedInsertionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CheckpointUnorderedInsertion()
func (myGovernor *MyGovernor) UnpackCheckpointUnorderedInsertionError(raw []byte) (*MyGovernorCheckpointUnorderedInsertion, error) {
	out := new(MyGovernorCheckpointUnorderedInsertion)
	if err := myGovernor.abi.UnpackIntoInterface(out, "CheckpointUnorderedInsertion", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorFailedCall represents a FailedCall error raised by the MyGovernor contract.
type MyGovernorFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func MyGovernorFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (myGovernor *MyGovernor) UnpackFailedCallError(raw []byte) (*MyGovernorFailedCall, error) {
	out := new(MyGovernorFailedCall)
	if err := myGovernor.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorAlreadyCastVote represents a GovernorAlreadyCastVote error raised by the MyGovernor contract.
type MyGovernorGovernorAlreadyCastVote struct {
	Voter common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorAlreadyCastVote(address voter)
func MyGovernorGovernorAlreadyCastVoteErrorID() common.Hash {
	return common.HexToHash("0x71c6af4932ed543cdb181fcbb800f4b9940a2ccceeaee5e6e141de8c50856ede")
}

// UnpackGovernorAlreadyCastVoteError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorAlreadyCastVote(address voter)
func (myGovernor *MyGovernor) UnpackGovernorAlreadyCastVoteError(raw []byte) (*MyGovernorGovernorAlreadyCastVote, error) {
	out := new(MyGovernorGovernorAlreadyCastVote)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorAlreadyCastVote", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorAlreadyQueuedProposal represents a GovernorAlreadyQueuedProposal error raised by the MyGovernor contract.
type MyGovernorGovernorAlreadyQueuedProposal struct {
	ProposalId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorAlreadyQueuedProposal(uint256 proposalId)
func MyGovernorGovernorAlreadyQueuedProposalErrorID() common.Hash {
	return common.HexToHash("0xf20e7d374e58691196b2e081c7753efc94203ab3520c58efe153076e279fde0a")
}

// UnpackGovernorAlreadyQueuedProposalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorAlreadyQueuedProposal(uint256 proposalId)
func (myGovernor *MyGovernor) UnpackGovernorAlreadyQueuedProposalError(raw []byte) (*MyGovernorGovernorAlreadyQueuedProposal, error) {
	out := new(MyGovernorGovernorAlreadyQueuedProposal)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorAlreadyQueuedProposal", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorDisabledDeposit represents a GovernorDisabledDeposit error raised by the MyGovernor contract.
type MyGovernorGovernorDisabledDeposit struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorDisabledDeposit()
func MyGovernorGovernorDisabledDepositErrorID() common.Hash {
	return common.HexToHash("0xe90a651e5fdea7022846d10b5f36994c22e1f46db4b5021aa3e26ad83b24bfd8")
}

// UnpackGovernorDisabledDepositError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorDisabledDeposit()
func (myGovernor *MyGovernor) UnpackGovernorDisabledDepositError(raw []byte) (*MyGovernorGovernorDisabledDeposit, error) {
	out := new(MyGovernorGovernorDisabledDeposit)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorDisabledDeposit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInsufficientProposerVotes represents a GovernorInsufficientProposerVotes error raised by the MyGovernor contract.
type MyGovernorGovernorInsufficientProposerVotes struct {
	Proposer  common.Address
	Votes     *big.Int
	Threshold *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInsufficientProposerVotes(address proposer, uint256 votes, uint256 threshold)
func MyGovernorGovernorInsufficientProposerVotesErrorID() common.Hash {
	return common.HexToHash("0xc242ee16ab08d11dbce60e744efdbd91b4e07ac4c074d993992519795a6324d0")
}

// UnpackGovernorInsufficientProposerVotesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInsufficientProposerVotes(address proposer, uint256 votes, uint256 threshold)
func (myGovernor *MyGovernor) UnpackGovernorInsufficientProposerVotesError(raw []byte) (*MyGovernorGovernorInsufficientProposerVotes, error) {
	out := new(MyGovernorGovernorInsufficientProposerVotes)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInsufficientProposerVotes", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidProposalLength represents a GovernorInvalidProposalLength error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidProposalLength struct {
	Targets   *big.Int
	Calldatas *big.Int
	Values    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidProposalLength(uint256 targets, uint256 calldatas, uint256 values)
func MyGovernorGovernorInvalidProposalLengthErrorID() common.Hash {
	return common.HexToHash("0x447b05d0c41e339e22932be48ca2091a47f3c39df25e2152ad11a8729753b2b4")
}

// UnpackGovernorInvalidProposalLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidProposalLength(uint256 targets, uint256 calldatas, uint256 values)
func (myGovernor *MyGovernor) UnpackGovernorInvalidProposalLengthError(raw []byte) (*MyGovernorGovernorInvalidProposalLength, error) {
	out := new(MyGovernorGovernorInvalidProposalLength)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidProposalLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidQuorumFraction represents a GovernorInvalidQuorumFraction error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidQuorumFraction struct {
	QuorumNumerator   *big.Int
	QuorumDenominator *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidQuorumFraction(uint256 quorumNumerator, uint256 quorumDenominator)
func MyGovernorGovernorInvalidQuorumFractionErrorID() common.Hash {
	return common.HexToHash("0x243e5445050913bb1c3de7a2f82eba0c3b0b8a55c2aacf660392fa35087a1919")
}

// UnpackGovernorInvalidQuorumFractionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidQuorumFraction(uint256 quorumNumerator, uint256 quorumDenominator)
func (myGovernor *MyGovernor) UnpackGovernorInvalidQuorumFractionError(raw []byte) (*MyGovernorGovernorInvalidQuorumFraction, error) {
	out := new(MyGovernorGovernorInvalidQuorumFraction)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidQuorumFraction", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidSignature represents a GovernorInvalidSignature error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidSignature struct {
	Voter common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidSignature(address voter)
func MyGovernorGovernorInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0x94ab6c07905fb25046d2809e4563b09690f891c9495bfe19551d602e7eddbb1b")
}

// UnpackGovernorInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidSignature(address voter)
func (myGovernor *MyGovernor) UnpackGovernorInvalidSignatureError(raw []byte) (*MyGovernorGovernorInvalidSignature, error) {
	out := new(MyGovernorGovernorInvalidSignature)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidVoteParams represents a GovernorInvalidVoteParams error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidVoteParams struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidVoteParams()
func MyGovernorGovernorInvalidVoteParamsErrorID() common.Hash {
	return common.HexToHash("0x867db7717d0cc803be5726127d33cc17ae07776705d8ba6659af8aaa5b418dd8")
}

// UnpackGovernorInvalidVoteParamsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidVoteParams()
func (myGovernor *MyGovernor) UnpackGovernorInvalidVoteParamsError(raw []byte) (*MyGovernorGovernorInvalidVoteParams, error) {
	out := new(MyGovernorGovernorInvalidVoteParams)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidVoteParams", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidVoteType represents a GovernorInvalidVoteType error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidVoteType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidVoteType()
func MyGovernorGovernorInvalidVoteTypeErrorID() common.Hash {
	return common.HexToHash("0x06b337c26289d63178b4d4ed5fd513f38a1d8832d0edd309ef07bfc9ba5caf49")
}

// UnpackGovernorInvalidVoteTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidVoteType()
func (myGovernor *MyGovernor) UnpackGovernorInvalidVoteTypeError(raw []byte) (*MyGovernorGovernorInvalidVoteType, error) {
	out := new(MyGovernorGovernorInvalidVoteType)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidVoteType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorInvalidVotingPeriod represents a GovernorInvalidVotingPeriod error raised by the MyGovernor contract.
type MyGovernorGovernorInvalidVotingPeriod struct {
	VotingPeriod *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorInvalidVotingPeriod(uint256 votingPeriod)
func MyGovernorGovernorInvalidVotingPeriodErrorID() common.Hash {
	return common.HexToHash("0xf1cfbf057db43f9730bc42a52728d66da9151a5c6929758ee824e299f82f5689")
}

// UnpackGovernorInvalidVotingPeriodError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorInvalidVotingPeriod(uint256 votingPeriod)
func (myGovernor *MyGovernor) UnpackGovernorInvalidVotingPeriodError(raw []byte) (*MyGovernorGovernorInvalidVotingPeriod, error) {
	out := new(MyGovernorGovernorInvalidVotingPeriod)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorInvalidVotingPeriod", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorNonexistentProposal represents a GovernorNonexistentProposal error raised by the MyGovernor contract.
type MyGovernorGovernorNonexistentProposal struct {
	ProposalId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorNonexistentProposal(uint256 proposalId)
func MyGovernorGovernorNonexistentProposalErrorID() common.Hash {
	return common.HexToHash("0x6ad06075316ea071ccae80931b756598be5aad3433b2c47b38607a8eec344a70")
}

// UnpackGovernorNonexistentProposalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorNonexistentProposal(uint256 proposalId)
func (myGovernor *MyGovernor) UnpackGovernorNonexistentProposalError(raw []byte) (*MyGovernorGovernorNonexistentProposal, error) {
	out := new(MyGovernorGovernorNonexistentProposal)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorNonexistentProposal", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorNotQueuedProposal represents a GovernorNotQueuedProposal error raised by the MyGovernor contract.
type MyGovernorGovernorNotQueuedProposal struct {
	ProposalId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorNotQueuedProposal(uint256 proposalId)
func MyGovernorGovernorNotQueuedProposalErrorID() common.Hash {
	return common.HexToHash("0xd5ddb8255ec3d5fb4ee2dd5d919eb1db6a2f1e4420bb74c3741830500cfb6a4f")
}

// UnpackGovernorNotQueuedProposalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorNotQueuedProposal(uint256 proposalId)
func (myGovernor *MyGovernor) UnpackGovernorNotQueuedProposalError(raw []byte) (*MyGovernorGovernorNotQueuedProposal, error) {
	out := new(MyGovernorGovernorNotQueuedProposal)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorNotQueuedProposal", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorOnlyExecutor represents a GovernorOnlyExecutor error raised by the MyGovernor contract.
type MyGovernorGovernorOnlyExecutor struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorOnlyExecutor(address account)
func MyGovernorGovernorOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x47096e4749c231af946d5efc74a7fd817371e756031e98787f18bf70aaa7753c")
}

// UnpackGovernorOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorOnlyExecutor(address account)
func (myGovernor *MyGovernor) UnpackGovernorOnlyExecutorError(raw []byte) (*MyGovernorGovernorOnlyExecutor, error) {
	out := new(MyGovernorGovernorOnlyExecutor)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorOnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorQueueNotImplemented represents a GovernorQueueNotImplemented error raised by the MyGovernor contract.
type MyGovernorGovernorQueueNotImplemented struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorQueueNotImplemented()
func MyGovernorGovernorQueueNotImplementedErrorID() common.Hash {
	return common.HexToHash("0x90884a46490684db0a73766419939e5ace793ae0f80195a286e159115c6628a0")
}

// UnpackGovernorQueueNotImplementedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorQueueNotImplemented()
func (myGovernor *MyGovernor) UnpackGovernorQueueNotImplementedError(raw []byte) (*MyGovernorGovernorQueueNotImplemented, error) {
	out := new(MyGovernorGovernorQueueNotImplemented)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorQueueNotImplemented", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorRestrictedProposer represents a GovernorRestrictedProposer error raised by the MyGovernor contract.
type MyGovernorGovernorRestrictedProposer struct {
	Proposer common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorRestrictedProposer(address proposer)
func MyGovernorGovernorRestrictedProposerErrorID() common.Hash {
	return common.HexToHash("0xd9b395579c6f1566cc7608394c53136f366f7fa719d01a941bee075ef8c704f4")
}

// UnpackGovernorRestrictedProposerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorRestrictedProposer(address proposer)
func (myGovernor *MyGovernor) UnpackGovernorRestrictedProposerError(raw []byte) (*MyGovernorGovernorRestrictedProposer, error) {
	out := new(MyGovernorGovernorRestrictedProposer)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorRestrictedProposer", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorUnableToCancel represents a GovernorUnableToCancel error raised by the MyGovernor contract.
type MyGovernorGovernorUnableToCancel struct {
	ProposalId *big.Int
	Account    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorUnableToCancel(uint256 proposalId, address account)
func MyGovernorGovernorUnableToCancelErrorID() common.Hash {
	return common.HexToHash("0x8fe5d8a9809b4b1a3569a8d98ce25e21fb956a89b6b9a2741d15bc699f46ef07")
}

// UnpackGovernorUnableToCancelError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorUnableToCancel(uint256 proposalId, address account)
func (myGovernor *MyGovernor) UnpackGovernorUnableToCancelError(raw []byte) (*MyGovernorGovernorUnableToCancel, error) {
	out := new(MyGovernorGovernorUnableToCancel)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorUnableToCancel", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorGovernorUnexpectedProposalState represents a GovernorUnexpectedProposalState error raised by the MyGovernor contract.
type MyGovernorGovernorUnexpectedProposalState struct {
	ProposalId     *big.Int
	Current        uint8
	ExpectedStates [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernorUnexpectedProposalState(uint256 proposalId, uint8 current, bytes32 expectedStates)
func MyGovernorGovernorUnexpectedProposalStateErrorID() common.Hash {
	return common.HexToHash("0x31b75e4d4f8317c390cf01cbc79dfe4f67ce2d27f65a099074fdc67f00f76908")
}

// UnpackGovernorUnexpectedProposalStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernorUnexpectedProposalState(uint256 proposalId, uint8 current, bytes32 expectedStates)
func (myGovernor *MyGovernor) UnpackGovernorUnexpectedProposalStateError(raw []byte) (*MyGovernorGovernorUnexpectedProposalState, error) {
	out := new(MyGovernorGovernorUnexpectedProposalState)
	if err := myGovernor.abi.UnpackIntoInterface(out, "GovernorUnexpectedProposalState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorInvalidAccountNonce represents a InvalidAccountNonce error raised by the MyGovernor contract.
type MyGovernorInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func MyGovernorInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (myGovernor *MyGovernor) UnpackInvalidAccountNonceError(raw []byte) (*MyGovernorInvalidAccountNonce, error) {
	out := new(MyGovernorInvalidAccountNonce)
	if err := myGovernor.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorInvalidShortString represents a InvalidShortString error raised by the MyGovernor contract.
type MyGovernorInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func MyGovernorInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (myGovernor *MyGovernor) UnpackInvalidShortStringError(raw []byte) (*MyGovernorInvalidShortString, error) {
	out := new(MyGovernorInvalidShortString)
	if err := myGovernor.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the MyGovernor contract.
type MyGovernorSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func MyGovernorSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (myGovernor *MyGovernor) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*MyGovernorSafeCastOverflowedUintDowncast, error) {
	out := new(MyGovernorSafeCastOverflowedUintDowncast)
	if err := myGovernor.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MyGovernorStringTooLong represents a StringTooLong error raised by the MyGovernor contract.
type MyGovernorStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func MyGovernorStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (myGovernor *MyGovernor) UnpackStringTooLongError(raw []byte) (*MyGovernorStringTooLong, error) {
	out := new(MyGovernorStringTooLong)
	if err := myGovernor.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}
