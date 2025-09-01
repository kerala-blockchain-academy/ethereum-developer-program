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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVotes\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractTimelockController\",\"name\":\"_timelock\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorAlreadyCastVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorAlreadyQueuedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorDisabledDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"GovernorInsufficientProposerVotes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldatas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidProposalLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorumNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quorumDenominator\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidQuorumFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidVotingPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNonexistentProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNotQueuedProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorQueueNotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"GovernorRestrictedProposer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorUnableToCancel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"current\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"GovernorUnexpectedProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"ProposalThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTimelock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"TimelockChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"VotingDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"VotingPeriodSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"getProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalNeedsQueuing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"setProposalThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newVotingDelay\",\"type\":\"uint48\"}],\"name\":\"setVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newVotingPeriod\",\"type\":\"uint32\"}],\"name\":\"setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTimelockController\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"updateTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "MyGovernor",
	Bin: "0x610180604052348015610010575f5ffd5b50604051614d68380380614d6883398101604081905261002f9161071b565b80600483600a60645f6040518060400160405280600a81526020016926bca3b7bb32b93737b960b11b8152508061006a61016d60201b60201c565b610074825f610188565b61012052610083816001610188565b61014052815160208084019190912060e052815190820120610100524660a05261010f60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c052600361012482826107eb565b506101309050836101ba565b61013982610220565b610142816102c4565b5050506001600160a01b03166101605261015b81610305565b506101658161039a565b505050610941565b6040805180820190915260018152603160f81b602082015290565b5f6020835110156101a35761019c83610403565b90506101b4565b816101ae84826107eb565b5060ff90505b92915050565b6008546040805165ffffffffffff928316815291831660208301527fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93910160405180910390a16008805465ffffffffffff191665ffffffffffff92909216919091179055565b8063ffffffff165f0361024d5760405163f1cfbf0560e01b81525f60048201526024015b60405180910390fd5b6008546040805163ffffffff66010000000000009093048316815291831660208301527f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828910160405180910390a16008805463ffffffff90921666010000000000000263ffffffff60301b19909216919091179055565b60075460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461910160405180910390a1600755565b6064808211156103325760405163243e544560e01b81526004810183905260248101829052604401610244565b5f61033b610440565b905061035a610348610459565b610351856104d3565b600a919061050a565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b600b54604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a1600b80546001600160a01b0319166001600160a01b0392909216919091179055565b5f5f829050601f8151111561042d578260405163305a27a960e01b815260040161024491906108a5565b8051610438826108da565b179392505050565b5f61044b600a610524565b6001600160d01b0316905090565b5f6104646101605190565b6001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156104bd575060408051601f3d908101601f191682019092526104ba918101906108fd565b60015b6104ce576104c961056c565b905090565b919050565b5f6001600160d01b03821115610506576040516306dfcc6560e41b815260d0600482015260248101839052604401610244565b5090565b5f80610517858585610576565b915091505b935093915050565b80545f9080156105635761054a8361053d600184610922565b5f91825260209091200190565b54660100000000000090046001600160d01b0316610565565b5f5b9392505050565b5f6104c9436106d2565b82545f9081908015610675575f6105928761053d600185610922565b805490915065ffffffffffff80821691660100000000000090046001600160d01b03169088168211156105d857604051632520601d60e01b815260040160405180910390fd5b8765ffffffffffff168265ffffffffffff160361061457825465ffffffffffff1666010000000000006001600160d01b03891602178355610667565b6040805180820190915265ffffffffffff808a1682526001600160d01b03808a1660208085019182528d54600181018f555f8f815291909120945191519092166601000000000000029216919091179101555b945085935061051c92505050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a555f8a815291822095519251909316660100000000000002919093161792019190915590508161051c565b5f65ffffffffffff821115610506576040516306dfcc6560e41b81526030600482015260248101839052604401610244565b6001600160a01b0381168114610718575f5ffd5b50565b5f5f6040838503121561072c575f5ffd5b825161073781610704565b602084015190925061074881610704565b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061077b57607f821691505b60208210810361079957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156107e657805f5260205f20601f840160051c810160208510156107c45750805b601f840160051c820191505b818110156107e3575f81556001016107d0565b50505b505050565b81516001600160401b0381111561080457610804610753565b610818816108128454610767565b8461079f565b6020601f82116001811461084a575f83156108335750848201515b5f19600385901b1c1916600184901b1784556107e3565b5f84815260208120601f198516915b828110156108795787850151825560209485019460019092019101610859565b508482101561089657868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610799575f1960209190910360031b1b16919050565b5f6020828403121561090d575f5ffd5b815165ffffffffffff81168114610565575f5ffd5b818103818111156101b457634e487b7160e01b5f52601160045260245ffd5b60805160a05160c05160e051610100516101205161014051610160516143b06109b85f395f81816109c901528181610e1f015281816111ec015281816114b10152611eba01525f611dfe01525f611dd201525f612df001525f612dc801525f612d2301525f612d4d01525f612d7701526143b05ff3fe6080604052600436106102a8575f3560e01c80637ecebe001161016f578063bc197c81116100d8578063deaaa7cc11610092578063ece40cc11161006d578063ece40cc11461095e578063f23a6e611461097d578063f8ce560a1461099c578063fc0c546a146109bb575f5ffd5b8063deaaa7cc146108ed578063e540d01d14610920578063eb9019d41461093f575f5ffd5b8063bc197c811461081b578063c01f9e371461083a578063c28bc2fa14610859578063c59057e41461086c578063d33219b41461088b578063dd4e2ba5146108a8575f5ffd5b8063a7713a7011610129578063a7713a7014610760578063a890c91014610774578063a8f8a66814610793578063a9a95294146107b2578063ab58fb8e146107d1578063b58131b014610807575f5ffd5b80637ecebe001461068957806384b0196e146106bd5780638ff262e3146106e457806391ddadf41461070357806397c3d3341461072e5780639a802a6d14610741575f5ffd5b806343859632116102115780635b8d0e0d116101cb5780635b8d0e0d146105cf5780635f398a14146105ee57806360c4247f1461060d578063790518871461062c5780637b3c71d31461064b5780637d5e81e21461066a575f5ffd5b806343859632146104d2578063452115d61461051a5780634bf5d7e914610539578063544ffc9c1461054d57806354fd4d501461058757806356781388146105b0575f5ffd5b8063160cbed711610262578063160cbed7146104065780632656227d146104255780632d63f693146104385780632fe3e261146104575780633932abb11461048a5780633e4f49e6146104a6575f5ffd5b806301ffc9a7146102e357806302a251a31461031757806306f3f9e61461034257806306fdde0314610361578063143489d014610382578063150b7a02146103ce575f5ffd5b366102df57306102b66109ed565b6001600160a01b0316146102dd57604051637485328f60e11b815260040160405180910390fd5b005b5f5ffd5b3480156102ee575f5ffd5b506103026102fd366004613395565b610a05565b60405190151581526020015b60405180910390f35b348015610322575f5ffd5b50600854600160301b900463ffffffff165b60405190815260200161030e565b34801561034d575f5ffd5b506102dd61035c3660046133bc565b610a71565b34801561036c575f5ffd5b50610375610a85565b60405161030e9190613401565b34801561038d575f5ffd5b506103b661039c3660046133bc565b5f908152600460205260409020546001600160a01b031690565b6040516001600160a01b03909116815260200161030e565b3480156103d9575f5ffd5b506103ed6103e83660046134ea565b610b15565b6040516001600160e01b0319909116815260200161030e565b348015610411575f5ffd5b506103346104203660046136b0565b610b57565b6103346104333660046136b0565b610c23565b348015610443575f5ffd5b506103346104523660046133bc565b610d8b565b348015610462575f5ffd5b506103347f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b348015610495575f5ffd5b5060085465ffffffffffff16610334565b3480156104b1575f5ffd5b506104c56104c03660046133bc565b610dab565b60405161030e9190613777565b3480156104dd575f5ffd5b506103026104ec366004613785565b5f8281526009602090815260408083206001600160a01b038516845260030190915290205460ff1692915050565b348015610525575f5ffd5b506103346105343660046136b0565b610db5565b348015610544575f5ffd5b50610375610e1b565b348015610558575f5ffd5b5061056c6105673660046133bc565b610edb565b6040805193845260208401929092529082015260600161030e565b348015610592575f5ffd5b506040805180820190915260018152603160f81b6020820152610375565b3480156105bb575f5ffd5b506103346105ca3660046137c3565b610f00565b3480156105da575f5ffd5b506103346105e9366004613831565b610f27565b3480156105f9575f5ffd5b506103346106083660046138ed565b610fe4565b348015610618575f5ffd5b506103346106273660046133bc565b61102c565b348015610637575f5ffd5b506102dd610646366004613981565b611038565b348015610656575f5ffd5b5061033461066536600461399c565b611049565b348015610675575f5ffd5b506103346106843660046139f1565b611099565b348015610694575f5ffd5b506103346106a3366004613aad565b6001600160a01b03165f9081526002602052604090205490565b3480156106c8575f5ffd5b506106d161114f565b60405161030e9796959493929190613b02565b3480156106ef575f5ffd5b506103346106fe366004613b71565b611191565b34801561070e575f5ffd5b506107176111e9565b60405165ffffffffffff909116815260200161030e565b348015610739575f5ffd5b506064610334565b34801561074c575f5ffd5b5061033461075b366004613bbe565b611270565b34801561076b575f5ffd5b50610334611286565b34801561077f575f5ffd5b506102dd61078e366004613aad565b61129f565b34801561079e575f5ffd5b506103346107ad3660046136b0565b6112b0565b3480156107bd575f5ffd5b506103026107cc3660046133bc565b6112bd565b3480156107dc575f5ffd5b506103346107eb3660046133bc565b5f9081526004602052604090206001015465ffffffffffff1690565b348015610812575f5ffd5b506103346112c5565b348015610826575f5ffd5b506103ed610835366004613c12565b6112cf565b348015610845575f5ffd5b506103346108543660046133bc565b611312565b6102dd610867366004613ca9565b611354565b348015610877575f5ffd5b506103346108863660046136b0565b6113d0565b348015610896575f5ffd5b50600b546001600160a01b03166103b6565b3480156108b3575f5ffd5b506040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e90820152610375565b3480156108f8575f5ffd5b506103347ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b34801561092b575f5ffd5b506102dd61093a366004613ce8565b611409565b34801561094a575f5ffd5b50610334610959366004613d0b565b61141a565b348015610969575f5ffd5b506102dd6109783660046133bc565b611439565b348015610988575f5ffd5b506103ed610997366004613d35565b61144a565b3480156109a7575f5ffd5b506103346109b63660046133bc565b61148d565b3480156109c6575f5ffd5b507f00000000000000000000000000000000000000000000000000000000000000006103b6565b5f610a00600b546001600160a01b031690565b905090565b5f6001600160e01b031982166366defe7760e11b1480610a3557506001600160e01b031982166332a2ad4360e11b145b80610a5057506001600160e01b03198216630271189760e51b145b80610a6b57506301ffc9a760e01b6001600160e01b03198316145b92915050565b610a7961152a565b610a82816115a3565b50565b606060038054610a9490613d8c565b80601f0160208091040260200160405190810160405280929190818152602001828054610ac090613d8c565b8015610b0b5780601f10610ae257610100808354040283529160200191610b0b565b820191905f5260205f20905b815481529060010190602001808311610aee57829003601f168201915b5050505050905090565b5f30610b1f6109ed565b6001600160a01b031614610b4657604051637485328f60e11b815260040160405180910390fd5b50630a85bd0160e11b949350505050565b5f5f610b65868686866112b0565b9050610b7a81610b756004611638565b61165a565b505f610b898288888888611697565b905065ffffffffffff811615610c00575f82815260046020908152604091829020600101805465ffffffffffff191665ffffffffffff85169081179091558251858152918201527f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892910160405180910390a1610c19565b604051634844252360e11b815260040160405180910390fd5b5095945050505050565b5f5f610c31868686866112b0565b9050610c5181610c416005611638565b610c4b6004611638565b1761165a565b505f818152600460205260409020805460ff60f01b1916600160f01b17905530610c796109ed565b6001600160a01b031614610d02575f5b8651811015610d0057306001600160a01b0316878281518110610cae57610cae613dc4565b60200260200101516001600160a01b031603610cf857610cf8858281518110610cd957610cd9613dc4565b60200260200101518051906020012060056116a590919063ffffffff16565b600101610c89565b505b610d0f8187878787611706565b30610d186109ed565b6001600160a01b031614158015610d4457506005546001600160801b03808216600160801b9092041614155b15610d4e575f6005555b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906020015b60405180910390a195945050505050565b5f90815260046020526040902054600160a01b900465ffffffffffff1690565b5f610a6b8261171a565b5f5f610dc3868686866112b0565b905033610dd08282611853565b610e0457604051638fe5d8a960e01b8152600481018390526001600160a01b03821660248201526044015b60405180910390fd5b610e1087878787611898565b979650505050505050565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634bf5d7e96040518163ffffffff1660e01b81526004015f60405180830381865afa925050508015610e9a57506040513d5f823e601f3d908101601f19168201604052610e979190810190613dd8565b60015b610ed6575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b5f818152600960205260409020805460018201546002909201549091905b9193909250565b5f80339050610f1f84828560405180602001604052805f8152506118a5565b949350505050565b5f610f6d88888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a92508991506118c69050565b610f95576040516394ab6c0760e01b81526001600160a01b0387166004820152602401610dfb565b610fd888878988888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a9250611993915050565b98975050505050505050565b5f80339050610e1087828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a9250611993915050565b5f610a6b600a83611a71565b61104061152a565b610a8281611abe565b5f8033905061108f86828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506118a592505050565b9695505050505050565b5f336110a58184611b24565b6110cd5760405163d9b3955760e01b81526001600160a01b0382166004820152602401610dfb565b5f6110d66112c5565b90508015611142575f6111048360016110ed6111e9565b6110f79190613e60565b65ffffffffffff1661141a565b90508181101561114057604051636121770b60e11b81526001600160a01b03841660048201526024810182905260448101839052606401610dfb565b505b610e108787878786611ba8565b5f6060805f5f5f6060611160611dcb565b611168611df7565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f61119e85858585611e24565b6111c6576040516394ab6c0760e01b81526001600160a01b0384166004820152602401610dfb565b6111e085848660405180602001604052805f8152506118a5565b95945050505050565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611264575060408051601f3d908101601f1916820190925261126191810190613e7e565b60015b610ed657610a00611ead565b5f61127c848484611eb7565b90505b9392505050565b5f611291600a611f4a565b6001600160d01b0316905090565b6112a761152a565b610a8281611f8e565b5f6111e0858585856113d0565b5f6001610a6b565b5f610a0060075490565b5f306112d96109ed565b6001600160a01b03161461130057604051637485328f60e11b815260040160405180910390fd5b5063bc197c8160e01b95945050505050565b5f8181526004602052604081205461134690600160d01b810463ffffffff1690600160a01b900465ffffffffffff16613e99565b65ffffffffffff1692915050565b61135c61152a565b5f5f856001600160a01b0316858585604051611379929190613eb7565b5f6040518083038185875af1925050503d805f81146113b3576040519150601f19603f3d011682016040523d82523d5f602084013e6113b8565b606091505b50915091506113c78282611ff7565b50505050505050565b5f848484846040516020016113e89493929190613f59565b60408051601f19818403018152919052805160209091012095945050505050565b61141161152a565b610a8281612013565b5f61127f838361143460408051602081019091525f815290565b611eb7565b61144161152a565b610a82816120af565b5f306114546109ed565b6001600160a01b03161461147b57604051637485328f60e11b815260040160405180910390fd5b5063f23a6e6160e01b95945050505050565b604051632394e7a360e21b8152600481018290525f90610a6b906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638e539e8c90602401602060405180830381865afa1580156114f6573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061151a9190613fa3565b6115238461102c565b60646120f0565b336115336109ed565b6001600160a01b03161461155c576040516347096e4760e01b8152336004820152602401610dfb565b306115656109ed565b6001600160a01b0316146115a1575f8036604051611584929190613eb7565b604051809103902090505b8061159a60056121a0565b0361158f57505b565b6064808211156115d05760405163243e544560e01b81526004810183905260248101829052604401610dfb565b5f6115d9611286565b90506115f86115e66111e9565b6115ef8561220d565b600a9190612244565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b5f81600781111561164b5761164b613743565b600160ff919091161b92915050565b5f5f61166584610dab565b90505f8361167283611638565b160361127f578381846040516331b75e4d60e01b8152600401610dfb93929190613fba565b5f61108f868686868661225e565b81546001600160801b03600160801b8204811691811660018301909116036116d1576116d160416123ef565b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b6117138585858585612400565b5050505050565b5f5f61172583612490565b9050600581600781111561173b5761173b613743565b146117465792915050565b5f838152600c60205260409081902054600b549151632c258a9f60e11b81526004810182905290916001600160a01b03169063584b153e90602401602060405180830381865afa15801561179c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117c09190613fdc565b156117cf575060059392505050565b600b54604051632ab0f52960e01b8152600481018390526001600160a01b0390911690632ab0f52990602401602060405180830381865afa158015611816573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061183a9190613fdc565b15611849575060079392505050565b5060029392505050565b5f8061185e84610dab565b600781111561186f5761186f613743565b14801561127f5750505f91825260046020526040909120546001600160a01b0391821691161490565b5f6111e0858585856125c1565b5f6111e0858585856118c160408051602081019091525f815290565b611993565b5f610e108561198d7f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118a8a8a6119188c6001600160a01b03165f90815260026020526040902080546001810190915590565b8b516020808e01919091208c518d83012060405161197298979695949301968752602087019590955260ff9390931660408601526001600160a01b03919091166060850152608084015260a083015260c082015260e00190565b60405160208183030381529060405280519060200120612658565b84612684565b5f6119a286610b756001611638565b505f6119b7866119b189610d8b565b85611eb7565b90505f6119c788888885886126f4565b905083515f03611a1d57866001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda489888489604051611a109493929190613ffb565b60405180910390a2610e10565b866001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb87128988848989604051611a5e959493929190614022565b60405180910390a2979650505050505050565b5f5f5f611a7d856127f0565b9250925050838265ffffffffffff161115611aaa57611aa5611a9e85612848565b869061287a565b611aac565b805b6001600160d01b031695945050505050565b6008546040805165ffffffffffff928316815291831660208301527fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93910160405180910390a16008805465ffffffffffff191665ffffffffffff92909216919091179055565b80515f906034811015611b3b576001915050610a6b565b60131981840101516001600160b01b03198116692370726f706f7365723d60b01b14611b6c57600192505050610a6b565b5f5f611b7c86602a86038661291c565b91509150811580610e105750866001600160a01b0316816001600160a01b031614979650505050505050565b5f611bbc86868686805190602001206112b0565b905084518651141580611bd157508351865114155b80611bdb57508551155b15611c1057855184518651604051630447b05d60e41b8152600481019390935260248301919091526044820152606401610dfb565b5f81815260046020526040902054600160a01b900465ffffffffffff1615611c595780611c3c82610dab565b6040516331b75e4d60e01b8152610dfb9291905f90600401613fba565b5f611c6b60085465ffffffffffff1690565b611c736111e9565b65ffffffffffff16611c85919061405b565b90505f611c9f60085463ffffffff600160301b9091041690565b5f84815260046020526040902080546001600160a01b0319166001600160a01b038716178155909150611cd183612848565b815465ffffffffffff91909116600160a01b0265ffffffffffff60a01b19909116178155611cfe826129c5565b815463ffffffff91909116600160d01b0263ffffffff60d01b1990911617815588517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e090859087908c908c906001600160401b03811115611d6157611d61613427565b604051908082528060200260200182016040528015611d9457816020015b6060815260200190600190039081611d7f5790505b508c89611da18a8261405b565b8e604051611db79998979695949392919061406e565b60405180910390a150505095945050505050565b6060610a007f00000000000000000000000000000000000000000000000000000000000000005f6129f5565b6060610a007f000000000000000000000000000000000000000000000000000000000000000060016129f5565b5f6111e08361198d7ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d7888888611e768a6001600160a01b03165f90815260026020526040902080546001810190915590565b60408051602081019690965285019390935260ff90911660608401526001600160a01b0316608083015260a082015260c001611972565b5f610a0043612848565b5f7f0000000000000000000000000000000000000000000000000000000000000000604051630748d63560e31b81526001600160a01b038681166004830152602482018690529190911690633a46b1a890604401602060405180830381865afa158015611f26573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061127c9190613fa3565b80545f908015611f8657611f7083611f63600184614149565b5f91825260209091200190565b54600160301b90046001600160d01b031661127f565b5f9392505050565b600b54604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a1600b80546001600160a01b0319166001600160a01b0392909216919091179055565b60608261200c5761200782612a9e565b610a6b565b5080610a6b565b8063ffffffff165f0361203b5760405163f1cfbf0560e01b81525f6004820152602401610dfb565b6008546040805163ffffffff600160301b9093048316815291831660208301527f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828910160405180910390a16008805463ffffffff909216600160301b0269ffffffff00000000000019909216919091179055565b60075460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461910160405180910390a1600755565b5f5f5f6120fd8686612ac6565b91509150815f03612121578381816121175761211761415c565b049250505061127f565b8184116121385761213860038515026011186123ef565b5f848688095f868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010185841190960395909502919093039390930492909217029150509392505050565b80545f906001600160801b0380821691600160801b90041681036121c8576121c860316123ef565b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b5f6001600160d01b03821115612240576040516306dfcc6560e41b815260d0600482015260248101839052604401610dfb565b5090565b5f80612251858585612ae2565b915091505b935093915050565b5f5f600b5f9054906101000a90046001600160a01b03166001600160a01b031663f27a0c926040518163ffffffff1660e01b8152600401602060405180830381865afa1580156122b0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122d49190613fa3565b90505f3060601b6bffffffffffffffffffffffff19168418600b5460405163b1c5f42760e01b81529192506001600160a01b03169063b1c5f42790612325908a908a908a905f908890600401614170565b602060405180830381865afa158015612340573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123649190613fa3565b5f898152600c602052604080822092909255600b5491516308f2a0bb60e41b81526001600160a01b0390921691638f2a0bb0916123ae918b918b918b919088908a906004016141bd565b5f604051808303815f87803b1580156123c5575f5ffd5b505af11580156123d7573d5f5f3e3d5ffd5b50505050610fd882426123ea919061405b565b612848565b634e487b715f52806020526024601cfd5b600b546001600160a01b031663e38335e5348686865f3060601b6bffffffffffffffffffffffff191688186040518763ffffffff1660e01b815260040161244b959493929190614170565b5f604051808303818588803b158015612462575f5ffd5b505af1158015612474573d5f5f3e3d5ffd5b5050505f9687525050600c602052505060408320929092555050565b5f818152600460205260408120805460ff600160f01b8204811691600160f81b90041681156124c457506007949350505050565b80156124d557506002949350505050565b5f6124df86610d8b565b9050805f0361250457604051636ad0607560e01b815260048101879052602401610dfb565b5f61250d6111e9565b65ffffffffffff16905080821061252a57505f9695505050505050565b5f61253488611312565b905081811061254b57506001979650505050505050565b61255488612c32565b158061257357505f888152600960205260409020805460019091015411155b1561258657506003979650505050505050565b5f8881526004602052604090206001015465ffffffffffff165f036125b357506004979650505050505050565b506005979650505050505050565b5f5f6125cf86868686612c68565b5f818152600c60205260409020549091508015610c1957600b5460405163c4d252f560e01b8152600481018390526001600160a01b039091169063c4d252f5906024015f604051808303815f87803b158015612629575f5ffd5b505af115801561263b573d5f5f3e3d5ffd5b5050505f838152600c602052604081205550509050949350505050565b5f610a6b612664612d17565b8360405161190160f01b8152600281019290925260228201526042902090565b5f836001600160a01b03163b5f036126e2575f5f6126a28585612e40565b5090925090505f8160038111156126bb576126bb613743565b1480156126d95750856001600160a01b0316826001600160a01b0316145b9250505061127f565b6126ed848484612e89565b905061127f565b5f8581526009602090815260408083206001600160a01b03881684526003810190925282205460ff1615612746576040516371c6af4960e01b81526001600160a01b0387166004820152602401610dfb565b6001600160a01b0386165f9081526003820160205260409020805460ff1916600117905560ff851661278f5783815f015f828254612784919061405b565b909155506127e59050565b5f1960ff8616016127ad5783816001015f828254612784919061405b565b60011960ff8616016127cc5783816002015f828254612784919061405b565b6040516303599be160e11b815260040160405180910390fd5b509195945050505050565b80545f908190819080820361280e575f5f5f93509350935050610ef9565b5f61281e86611f63600185614149565b546001955065ffffffffffff81169450600160301b90046001600160d01b03169250610ef9915050565b5f65ffffffffffff821115612240576040516306dfcc6560e41b81526030600482015260248101839052604401610dfb565b81545f90818160058111156128d6575f61289384612f5f565b61289d9085614149565b5f8881526020902090915081015465ffffffffffff90811690871610156128c6578091506128d4565b6128d181600161405b565b92505b505b5f6128e3878785856130b7565b90508015612910576128fa87611f63600184614149565b54600160301b90046001600160d01b0316610e10565b505f9695505050505050565b5f5f845183118061292c57508284115b1561293b57505f905080612256565b5f61294785600161405b565b8411801561296f575061060f60f31b6129638787016020015190565b6001600160f01b031916145b90505f61297f8215156002614214565b61298a90602861405b565b9050806129978787614149565b036129b8575f5f6129a9898989613116565b90965094506122569350505050565b5f5f935093505050612256565b5f63ffffffff821115612240576040516306dfcc6560e41b81526020600482015260248101839052604401610dfb565b606060ff8314612a0f57612a08836131d7565b9050610a6b565b818054612a1b90613d8c565b80601f0160208091040260200160405190810160405280929190818152602001828054612a4790613d8c565b8015612a925780601f10612a6957610100808354040283529160200191612a92565b820191905f5260205f20905b815481529060010190602001808311612a7557829003601f168201915b50505050509050610a6b565b805115612aad57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f805f1983850993909202808410938190039390930393915050565b82545f9081908015612bd8575f612afe87611f63600185614149565b805490915065ffffffffffff80821691600160301b90046001600160d01b0316908816821115612b4157604051632520601d60e01b815260040160405180910390fd5b8765ffffffffffff168265ffffffffffff1603612b7a57825465ffffffffffff16600160301b6001600160d01b03891602178355612bca565b6040805180820190915265ffffffffffff808a1682526001600160d01b03808a1660208085019182528d54600181018f555f8f81529190912094519151909216600160301b029216919091179101555b945085935061225692505050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a555f8a815291822095519251909316600160301b029190931617920191909155905081612256565b5f81815260096020526040812060028101546001820154612c53919061405b565b612c5f6109b685610d8b565b11159392505050565b5f5f612c76868686866112b0565b9050612cc481612c866007611638565b612c906006611638565b612c9a6002611638565b6001612ca760078261422b565b612cb290600261431f565b612cbc9190614149565b18181861165a565b505f818152600460205260409081902080546001600160f81b0316600160f81b179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c90610d7a9083815260200190565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015612d6f57507f000000000000000000000000000000000000000000000000000000000000000046145b15612d9957507f000000000000000000000000000000000000000000000000000000000000000090565b610a00604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f5f5f8351604103612e77576020840151604085015160608601515f1a612e6988828585613214565b955095509550505050612e82565b505081515f91506002905b9250925092565b5f5f5f856001600160a01b03168585604051602401612ea992919061432d565b60408051601f198184030181529181526020820180516001600160e01b0316630b135d3f60e11b17905251612ede9190614345565b5f60405180830381855afa9150503d805f8114612f16576040519150601f19603f3d011682016040523d82523d5f602084013e612f1b565b606091505b5091509150818015612f2f57506020815110155b801561108f57508051630b135d3f60e11b90612f549083016020908101908401613fa3565b149695505050505050565b5f60018211612f6c575090565b816001600160801b8210612f855760809190911c9060401b5b680100000000000000008210612fa05760409190911c9060201b5b6401000000008210612fb75760209190911c9060101b5b620100008210612fcc5760109190911c9060081b5b6101008210612fe05760089190911c9060041b5b60108210612ff35760049190911c9060021b5b60048210612fff5760011b5b600302600190811c908185816130175761301761415c565b048201901c9050600181858161302f5761302f61415c565b048201901c905060018185816130475761304761415c565b048201901c9050600181858161305f5761305f61415c565b048201901c905060018185816130775761307761415c565b048201901c9050600181858161308f5761308f61415c565b048201901c90506130ae8185816130a8576130a861415c565b04821190565b90039392505050565b5f5b8183101561310e575f6130cc84846132dc565b5f8781526020902090915065ffffffffffff86169082015465ffffffffffff1611156130fa57809250613108565b61310581600161405b565b93505b506130b9565b509392505050565b5f80848161312586600161405b565b8511801561314d575061060f60f31b6131418388016020015190565b6001600160f01b031916145b90505f61315d8215156002614214565b90505f8061316b838a61405b565b90505b878110156131c6575f61318c6131878784016020015190565b6132f6565b9050600f8160ff1611156131ab575f5f97509750505050505050612256565b6131b6601084614214565b60ff90911601915060010161316e565b506001999098509650505050505050565b60605f6131e38361336e565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561324d57505f915060039050826132d2565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561329e573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166132c957505f9250600191508290506132d2565b92505f91508190505b9450945094915050565b5f6132ea600284841861435b565b61127f9084841661405b565b5f60f882901c602f8111801561330f5750603a8160ff16105b1561331d57602f1901610a6b565b60608160ff16118015613333575060678160ff16105b156133415760561901610a6b565b60408160ff16118015613357575060478160ff16105b156133655760361901610a6b565b5060ff92915050565b5f60ff8216601f811115610a6b57604051632cd44ac360e21b815260040160405180910390fd5b5f602082840312156133a5575f5ffd5b81356001600160e01b03198116811461127f575f5ffd5b5f602082840312156133cc575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61127f60208301846133d3565b6001600160a01b0381168114610a82575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561346357613463613427565b604052919050565b5f6001600160401b0382111561348357613483613427565b50601f01601f191660200190565b5f6134a361349e8461346b565b61343b565b90508281528383830111156134b6575f5ffd5b828260208301375f602084830101529392505050565b5f82601f8301126134db575f5ffd5b61127f83833560208501613491565b5f5f5f5f608085870312156134fd575f5ffd5b843561350881613413565b9350602085013561351881613413565b92506040850135915060608501356001600160401b03811115613539575f5ffd5b613545878288016134cc565b91505092959194509250565b5f6001600160401b0382111561356957613569613427565b5060051b60200190565b5f82601f830112613582575f5ffd5b813561359061349e82613551565b8082825260208201915060208360051b8601019250858311156135b1575f5ffd5b602085015b83811015610c195780356135c981613413565b8352602092830192016135b6565b5f82601f8301126135e6575f5ffd5b81356135f461349e82613551565b8082825260208201915060208360051b860101925085831115613615575f5ffd5b602085015b83811015610c1957803583526020928301920161361a565b5f82601f830112613641575f5ffd5b813561364f61349e82613551565b8082825260208201915060208360051b860101925085831115613670575f5ffd5b602085015b83811015610c195780356001600160401b03811115613692575f5ffd5b6136a1886020838a01016134cc565b84525060209283019201613675565b5f5f5f5f608085870312156136c3575f5ffd5b84356001600160401b038111156136d8575f5ffd5b6136e487828801613573565b94505060208501356001600160401b038111156136ff575f5ffd5b61370b878288016135d7565b93505060408501356001600160401b03811115613726575f5ffd5b61373287828801613632565b949793965093946060013593505050565b634e487b7160e01b5f52602160045260245ffd5b6008811061377357634e487b7160e01b5f52602160045260245ffd5b9052565b60208101610a6b8284613757565b5f5f60408385031215613796575f5ffd5b8235915060208301356137a881613413565b809150509250929050565b803560ff81168114610ed6575f5ffd5b5f5f604083850312156137d4575f5ffd5b823591506137e4602084016137b3565b90509250929050565b5f5f83601f8401126137fd575f5ffd5b5081356001600160401b03811115613813575f5ffd5b60208301915083602082850101111561382a575f5ffd5b9250929050565b5f5f5f5f5f5f5f60c0888a031215613847575f5ffd5b87359650613857602089016137b3565b9550604088013561386781613413565b945060608801356001600160401b03811115613881575f5ffd5b61388d8a828b016137ed565b90955093505060808801356001600160401b038111156138ab575f5ffd5b6138b78a828b016134cc565b92505060a08801356001600160401b038111156138d2575f5ffd5b6138de8a828b016134cc565b91505092959891949750929550565b5f5f5f5f5f60808688031215613901575f5ffd5b85359450613911602087016137b3565b935060408601356001600160401b0381111561392b575f5ffd5b613937888289016137ed565b90945092505060608601356001600160401b03811115613955575f5ffd5b613961888289016134cc565b9150509295509295909350565b65ffffffffffff81168114610a82575f5ffd5b5f60208284031215613991575f5ffd5b813561127f8161396e565b5f5f5f5f606085870312156139af575f5ffd5b843593506139bf602086016137b3565b925060408501356001600160401b038111156139d9575f5ffd5b6139e5878288016137ed565b95989497509550505050565b5f5f5f5f60808587031215613a04575f5ffd5b84356001600160401b03811115613a19575f5ffd5b613a2587828801613573565b94505060208501356001600160401b03811115613a40575f5ffd5b613a4c878288016135d7565b93505060408501356001600160401b03811115613a67575f5ffd5b613a7387828801613632565b92505060608501356001600160401b03811115613a8e575f5ffd5b8501601f81018713613a9e575f5ffd5b61354587823560208401613491565b5f60208284031215613abd575f5ffd5b813561127f81613413565b5f8151808452602084019350602083015f5b82811015613af8578151865260209586019590910190600101613ada565b5093949350505050565b60ff60f81b8816815260e060208201525f613b2060e08301896133d3565b8281036040840152613b3281896133d3565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050613b638185613ac8565b9a9950505050505050505050565b5f5f5f5f60808587031215613b84575f5ffd5b84359350613b94602086016137b3565b92506040850135613ba481613413565b915060608501356001600160401b03811115613539575f5ffd5b5f5f5f60608486031215613bd0575f5ffd5b8335613bdb81613413565b92506020840135915060408401356001600160401b03811115613bfc575f5ffd5b613c08868287016134cc565b9150509250925092565b5f5f5f5f5f60a08688031215613c26575f5ffd5b8535613c3181613413565b94506020860135613c4181613413565b935060408601356001600160401b03811115613c5b575f5ffd5b613c67888289016135d7565b93505060608601356001600160401b03811115613c82575f5ffd5b613c8e888289016135d7565b92505060808601356001600160401b03811115613955575f5ffd5b5f5f5f5f60608587031215613cbc575f5ffd5b8435613cc781613413565b93506020850135925060408501356001600160401b038111156139d9575f5ffd5b5f60208284031215613cf8575f5ffd5b813563ffffffff8116811461127f575f5ffd5b5f5f60408385031215613d1c575f5ffd5b8235613d2781613413565b946020939093013593505050565b5f5f5f5f5f60a08688031215613d49575f5ffd5b8535613d5481613413565b94506020860135613d6481613413565b9350604086013592506060860135915060808601356001600160401b03811115613955575f5ffd5b600181811c90821680613da057607f821691505b602082108103613dbe57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215613de8575f5ffd5b81516001600160401b03811115613dfd575f5ffd5b8201601f81018413613e0d575f5ffd5b8051613e1b61349e8261346b565b818152856020838501011115613e2f575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b634e487b7160e01b5f52601160045260245ffd5b65ffffffffffff8281168282160390811115610a6b57610a6b613e4c565b5f60208284031215613e8e575f5ffd5b815161127f8161396e565b65ffffffffffff8181168382160190811115610a6b57610a6b613e4c565b818382375f9101908152919050565b5f8151808452602084019350602083015f5b82811015613af85781516001600160a01b0316865260209586019590910190600101613ed8565b5f82825180855260208501945060208160051b830101602085015f5b83811015613f4d57601f19858403018852613f378383516133d3565b6020988901989093509190910190600101613f1b565b50909695505050505050565b608081525f613f6b6080830187613ec6565b8281036020840152613f7d8187613ac8565b90508281036040840152613f918186613eff565b91505082606083015295945050505050565b5f60208284031215613fb3575f5ffd5b5051919050565b83815260608101613fce6020830185613757565b826040830152949350505050565b5f60208284031215613fec575f5ffd5b8151801515811461127f575f5ffd5b84815260ff84166020820152826040820152608060608201525f61108f60808301846133d3565b85815260ff8516602082015283604082015260a060608201525f61404960a08301856133d3565b8281036080840152610fd881856133d3565b80820180821115610a6b57610a6b613e4c565b8981526001600160a01b0389166020820152610120604082018190525f906140989083018a613ec6565b82810360608401526140aa818a613ac8565b9050828103608084015280885180835260208301915060208160051b84010160208b015f5b8381101561410157601f198684030185526140eb8383516133d3565b60209586019590935091909101906001016140cf565b505085810360a0870152614115818b613eff565b93505050508560c08401528460e084015282810361010084015261413981856133d3565b9c9b505050505050505050505050565b81810381811115610a6b57610a6b613e4c565b634e487b7160e01b5f52601260045260245ffd5b60a081525f61418260a0830188613ec6565b82810360208401526141948188613ac8565b905082810360408401526141a88187613eff565b60608401959095525050608001529392505050565b60c081525f6141cf60c0830189613ec6565b82810360208401526141e18189613ac8565b905082810360408401526141f58188613eff565b60608401969096525050608081019290925260a0909101529392505050565b8082028115828204841417610a6b57610a6b613e4c565b60ff8181168382160190811115610a6b57610a6b613e4c565b6001815b60018411156122565780850481111561426357614263613e4c565b600184161561427157908102905b60019390931c928002614248565b5f8261428d57506001610a6b565b8161429957505f610a6b565b81600181146142af57600281146142b9576142d5565b6001915050610a6b565b60ff8411156142ca576142ca613e4c565b50506001821b610a6b565b5060208310610133831016604e8410600b84101617156142f8575081810a610a6b565b6143045f198484614244565b805f190482111561431757614317613e4c565b029392505050565b5f61127f60ff84168361427f565b828152604060208201525f61127c60408301846133d3565b5f82518060208501845e5f920191825250919050565b5f8261437557634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220a98bd205eeeaa47e2ab454b6f128d517165046281c68bd736303a721a900302164736f6c634300081e0033",
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
// Solidity: function proposalThreshold() view returns(uint256)
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
// Solidity: function proposalThreshold() view returns(uint256)
func (myGovernor *MyGovernor) TryPackProposalThreshold() ([]byte, error) {
	return myGovernor.abi.Pack("proposalThreshold")
}

// UnpackProposalThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
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

// PackSetProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xece40cc1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (myGovernor *MyGovernor) PackSetProposalThreshold(newProposalThreshold *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("setProposalThreshold", newProposalThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetProposalThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xece40cc1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (myGovernor *MyGovernor) TryPackSetProposalThreshold(newProposalThreshold *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("setProposalThreshold", newProposalThreshold)
}

// PackSetVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79051887.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (myGovernor *MyGovernor) PackSetVotingDelay(newVotingDelay *big.Int) []byte {
	enc, err := myGovernor.abi.Pack("setVotingDelay", newVotingDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetVotingDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79051887.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (myGovernor *MyGovernor) TryPackSetVotingDelay(newVotingDelay *big.Int) ([]byte, error) {
	return myGovernor.abi.Pack("setVotingDelay", newVotingDelay)
}

// PackSetVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe540d01d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (myGovernor *MyGovernor) PackSetVotingPeriod(newVotingPeriod uint32) []byte {
	enc, err := myGovernor.abi.Pack("setVotingPeriod", newVotingPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetVotingPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe540d01d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (myGovernor *MyGovernor) TryPackSetVotingPeriod(newVotingPeriod uint32) ([]byte, error) {
	return myGovernor.abi.Pack("setVotingPeriod", newVotingPeriod)
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
// Solidity: function votingDelay() view returns(uint256)
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
// Solidity: function votingDelay() view returns(uint256)
func (myGovernor *MyGovernor) TryPackVotingDelay() ([]byte, error) {
	return myGovernor.abi.Pack("votingDelay")
}

// UnpackVotingDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
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
// Solidity: function votingPeriod() view returns(uint256)
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
// Solidity: function votingPeriod() view returns(uint256)
func (myGovernor *MyGovernor) TryPackVotingPeriod() ([]byte, error) {
	return myGovernor.abi.Pack("votingPeriod")
}

// UnpackVotingPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
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

// MyGovernorProposalThresholdSet represents a ProposalThresholdSet event raised by the MyGovernor contract.
type MyGovernorProposalThresholdSet struct {
	OldProposalThreshold *big.Int
	NewProposalThreshold *big.Int
	Raw                  *types.Log // Blockchain specific contextual infos
}

const MyGovernorProposalThresholdSetEventName = "ProposalThresholdSet"

// ContractEventName returns the user-defined event name.
func (MyGovernorProposalThresholdSet) ContractEventName() string {
	return MyGovernorProposalThresholdSetEventName
}

// UnpackProposalThresholdSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (myGovernor *MyGovernor) UnpackProposalThresholdSetEvent(log *types.Log) (*MyGovernorProposalThresholdSet, error) {
	event := "ProposalThresholdSet"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorProposalThresholdSet)
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

// MyGovernorVotingDelaySet represents a VotingDelaySet event raised by the MyGovernor contract.
type MyGovernorVotingDelaySet struct {
	OldVotingDelay *big.Int
	NewVotingDelay *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const MyGovernorVotingDelaySetEventName = "VotingDelaySet"

// ContractEventName returns the user-defined event name.
func (MyGovernorVotingDelaySet) ContractEventName() string {
	return MyGovernorVotingDelaySetEventName
}

// UnpackVotingDelaySetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (myGovernor *MyGovernor) UnpackVotingDelaySetEvent(log *types.Log) (*MyGovernorVotingDelaySet, error) {
	event := "VotingDelaySet"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorVotingDelaySet)
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

// MyGovernorVotingPeriodSet represents a VotingPeriodSet event raised by the MyGovernor contract.
type MyGovernorVotingPeriodSet struct {
	OldVotingPeriod *big.Int
	NewVotingPeriod *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const MyGovernorVotingPeriodSetEventName = "VotingPeriodSet"

// ContractEventName returns the user-defined event name.
func (MyGovernorVotingPeriodSet) ContractEventName() string {
	return MyGovernorVotingPeriodSetEventName
}

// UnpackVotingPeriodSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (myGovernor *MyGovernor) UnpackVotingPeriodSetEvent(log *types.Log) (*MyGovernorVotingPeriodSet, error) {
	event := "VotingPeriodSet"
	if log.Topics[0] != myGovernor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MyGovernorVotingPeriodSet)
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
