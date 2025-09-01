package tests

import (
	"context"
	"ethereum-developer-program/dao/lib/bindings"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestMyGovernor(t *testing.T) {
	// First, we deploy the GovToken contract. We'll also use the same simulated backend to deploy other contracts as well.
	g, err := initializeGovToken()
	if err != nil {
		t.Fatalf("Failed to initialize GovToken: %v", err)
	}

	// Initialize a TimeLock instance.
	timelock := bindings.NewTimeLock()

	// Deploy the TimeLock contract. We assign deployer as the admin, proposer and executor.
	timelockAddress, _, err := bind.DeployContract(
		g.owner,
		common.FromHex(bindings.TimeLockMetaData.Bin),
		g.sim.Client(),
		timelock.PackConstructor(
			big.NewInt(0),
			[]common.Address{g.owner.From},
			[]common.Address{g.owner.From},
			g.owner.From,
		),
	)
	if err != nil {
		t.Fatalf("Failed to deploy TimeLock contract: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Initialize a Cert instance.
	cert := bindings.NewCert()

	// Deploy the Cert contract. We assign the TimeLock contract as the owner.
	certAddress, _, err := bind.DeployContract(
		g.owner,
		common.FromHex(bindings.CertMetaData.Bin),
		g.sim.Client(),
		cert.PackConstructor(timelockAddress),
	)
	if err != nil {
		t.Fatalf("Failed to deploy Cert contract: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Initialize a MyGovernor instance.
	myGovernor := bindings.NewMyGovernor()

	// Deploy the MyGovernor contract. We provide the GovToken contract address and TimeLock contract address as arguments.
	myGovernorAddress, _, err := bind.DeployContract(
		g.owner,
		common.FromHex(bindings.MyGovernorMetaData.Bin),
		g.sim.Client(),
		myGovernor.PackConstructor(*g.address, timelockAddress),
	)
	if err != nil {
		t.Fatalf("Failed to deploy MyGovernor contract: %v", err)
	}

	// Total supply should be '1000000000000000000000'.
	supply, err := bind.Call(g.instance, nil, g.contract.PackTotalSupply(), g.contract.UnpackTotalSupply)
	if err != nil {
		t.Fatalf("Failed to call totalSupply method: %v", err)
	}

	// Check the balance of deployer account. This is should be equal to the total supply.
	balance, err := bind.Call(g.instance, nil, g.contract.PackBalanceOf(g.owner.From), g.contract.UnpackBalanceOf)
	if err != nil {
		t.Fatalf("Failed to call balanceOf method: %v", err)
	}

	assert.Equal(t, supply, balance)

	// Before delegation, the voting power should be '0'.
	votes, err := bind.Call(g.instance, nil, g.contract.PackGetVotes(g.owner.From), g.contract.UnpackGetVotes)
	if err != nil {
		t.Fatalf("Failed to call getVotes method: %v", err)
	}

	assert.Equal(t, int64(0), votes.Int64())

	// Owner account delegates its voting power.
	if _, err = bind.Transact(g.instance, g.owner, g.contract.PackDelegate(g.owner.From)); err != nil {
		t.Fatalf("Failed to call delegate method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// After delegation, the voting power should be equal to the total supply.
	votes, err = bind.Call(g.instance, nil, g.contract.PackGetVotes(g.owner.From), g.contract.UnpackGetVotes)
	if err != nil {
		t.Fatalf("Failed to call getVotes method: %v", err)
	}

	assert.Equal(t, supply, votes)

	// Create a instance wrapper for the TimeLock contract.
	timelockInstance := timelock.Instance(g.sim.Client(), timelockAddress)

	// Set up roles for TimeLock contract.
	PROPOSER_ROLE := [32]byte(crypto.Keccak256([]byte("PROPOSER_ROLE")))
	EXECUTOR_ROLE := [32]byte(crypto.Keccak256([]byte("EXECUTOR_ROLE")))

	// Grant PROPOSER_ROLE to MyGovernor contract.
	if _, err = bind.Transact(timelockInstance, g.owner, timelock.PackGrantRole(PROPOSER_ROLE, myGovernorAddress)); err != nil {
		t.Fatalf("Failed to call grantRole method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Grant EXECUTOR_ROLE to MyGovernor contract.
	if _, err = bind.Transact(timelockInstance, g.owner, timelock.PackGrantRole(EXECUTOR_ROLE, myGovernorAddress)); err != nil {
		t.Fatalf("Failed to call grantRole method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Specify certificate data.
	id := big.NewInt(21)
	certificate := bindings.CertificatesOutput{
		Name:   "Shin",
		Course: "MBCC",
		Grade:  "S",
		Date:   "2025-09-01",
	}

	callData, err := cert.TryPackIssue(id, certificate.Name, certificate.Course, certificate.Grade, certificate.Date)
	if err != nil {
		t.Fatalf("Failed to pack parameters for issue method: %v", err)
	}

	// Create a instance wrapper for the MyGovernor contract.
	myGovernorInstance := myGovernor.Instance(g.sim.Client(), myGovernorAddress)

	// Create a description for the first proposal.
	description := fmt.Sprintf("Proposal #1: Issue certificate for ID #%d", id.Int64())

	// Create the first proposal.
	if _, err = bind.Transact(myGovernorInstance, g.owner,
		myGovernor.PackPropose(
			[]common.Address{certAddress},
			[]*big.Int{big.NewInt(0)},
			[][]byte{callData},
			description,
		),
	); err != nil {
		t.Fatalf("Failed to call propose method: %v", err)
	}

	// Commit the blocks so far. We store the latest block hash for log filtering.
	latestBlock := g.sim.Commit()

	// Filter events from the logs.
	logs, err := g.sim.Client().FilterLogs(context.Background(), ethereum.FilterQuery{BlockHash: &latestBlock})
	if err != nil {
		t.Fatalf("Failed to filter logs: %v", err)
	}

	// Create a variable to store the proposal ID.
	var proposalID *big.Int
	for _, l := range logs {
		event, err := myGovernor.UnpackProposalCreatedEvent(&l)
		if err != nil {
			t.Fatalf("Failed to unpack ProposalCreated event: %v", err)
		}
		proposalID = event.ProposalId
	}

	// Fetch the current state of the proposal. This should be '0' (pending).
	state, err := bind.Call(myGovernorInstance, nil, myGovernor.PackState(proposalID), myGovernor.UnpackState)
	if err != nil {
		t.Fatalf("Failed to call state method: %v", err)
	}

	assert.Equal(t, uint8(0), state)

	// Advance the chain for a given number (greater than the voting delay) of blocks.
	g.skipBlocks(11)

	// Fetch the current state of the proposal. This should be '1' (active) by this point.
	state, err = bind.Call(myGovernorInstance, nil, myGovernor.PackState(proposalID), myGovernor.UnpackState)
	if err != nil {
		t.Fatalf("Failed to call state method: %v", err)
	}

	assert.Equal(t, uint8(1), state)

	// Deployer casts a vote for the proposal.
	if _, err = bind.Transact(myGovernorInstance, g.owner, myGovernor.PackCastVote(proposalID, uint8(1))); err != nil {
		t.Fatalf("Failed to call castVote method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Fetch the current votes for the proposal. The 'for' votes should be equal to the total supply, while the others are '0'.
	proposalVotes, err := bind.Call(myGovernorInstance, nil, myGovernor.PackProposalVotes(proposalID), myGovernor.UnpackProposalVotes)
	if err != nil {
		t.Fatalf("Failed to call state method: %v", err)
	}

	assert.Equal(t, supply, proposalVotes.ForVotes)
	assert.Equal(t, int64(0), proposalVotes.AbstainVotes.Int64())
	assert.Equal(t, int64(0), proposalVotes.AgainstVotes.Int64())

	// Advance the chain for a given number (greater than the voting period) of blocks.
	g.skipBlocks(101)

	// Fetch the current state of the proposal. This should be '4' (succeeded) by this point.
	state, err = bind.Call(myGovernorInstance, nil, myGovernor.PackState(proposalID), myGovernor.UnpackState)
	if err != nil {
		t.Fatalf("Failed to call state method: %v", err)
	}

	assert.Equal(t, uint8(4), state)

	// Generate the bytes32 hash of the description.
	descriptionHash := [32]byte(crypto.Keccak256([]byte(description)))

	// Deployer, aka Executor, puts the proposal in the queue.
	if _, err = bind.Transact(myGovernorInstance, g.owner,
		myGovernor.PackQueue(
			[]common.Address{certAddress},
			[]*big.Int{big.NewInt(0)},
			[][]byte{callData},
			descriptionHash),
	); err != nil {
		t.Fatalf("Failed to call queue method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Fetch the current state of the proposal. This should be '5' (queued) by this point.
	state, err = bind.Call(myGovernorInstance, nil, myGovernor.PackState(proposalID), myGovernor.UnpackState)
	if err != nil {
		t.Fatalf("Failed to call state method: %v", err)
	}

	assert.Equal(t, uint8(5), state)

	// This is optional, since our minDelay is '0'. Use if otherwise.
	if err := g.sim.AdjustTime(time.Second); err != nil {
		t.Fatalf("Failed to adjust time of backend: %v", err)
	}

	// Executor executes the proposal from the queue.
	if _, err = bind.Transact(myGovernorInstance, g.owner,
		myGovernor.PackExecute(
			[]common.Address{certAddress},
			[]*big.Int{big.NewInt(0)},
			[][]byte{callData},
			descriptionHash),
	); err != nil {
		t.Fatalf("Failed to call execute method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Fetch the current state of the proposal. This should be '7' (executed) by this point.
	state, err = bind.Call(myGovernorInstance, nil, myGovernor.PackState(proposalID), myGovernor.UnpackState)
	if err != nil {
		t.Fatalf("Failed to call state method: %v", err)
	}

	assert.Equal(t, uint8(7), state)

	// Create a instance wrapper for the Cert contract.
	certInstance := cert.Instance(g.sim.Client(), certAddress)

	// Finally, fetch the certificate from the Cert contract corresponding to the ID.
	data, err := bind.Call(certInstance, nil, cert.PackCertificates(id), cert.UnpackCertificates)
	if err != nil {
		t.Fatalf("Failed to call Certificates method: %v", err)
	}

	if assert.NotNil(t, data) {
		assert.EqualValues(t, certificate, data)
	}
}
