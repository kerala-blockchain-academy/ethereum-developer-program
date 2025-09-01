package tests

import (
	"context"
	"ethereum-developer-program/dao/lib/bindings"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
)

func TestCert(t *testing.T) {
	// Create a random key for the deployer/owner account.
	ownerKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate key 1: %v", err)
	}

	// Create a random key for the other account.
	otherKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate key 2: %v", err)
	}

	// Set up the simulated blockchain with enough balance for the owner and other accounts.
	chainID := params.AllDevChainProtocolChanges.ChainID
	ownerAuth := bind.NewKeyedTransactor(ownerKey, chainID)
	otherAuth := bind.NewKeyedTransactor(otherKey, chainID)
	sim := simulated.NewBackend(map[common.Address]types.Account{
		ownerAuth.From: {Balance: big.NewInt(9e18)},
		otherAuth.From: {Balance: big.NewInt(9e18)},
	})

	// Initialize a instance for the contract.
	cert := bindings.NewCert()

	// Deploy the smart contract.
	ca, _, err := bind.DeployContract(ownerAuth, common.FromHex(bindings.CertMetaData.Bin), sim.Client(), cert.PackConstructor(ownerAuth.From))
	if err != nil {
		t.Fatalf("Failed to deploy smart contract: %v", err)
	}

	// Commit the blocks so far.
	sim.Commit()

	// Create a instance wrapper for the deployed contract.
	instance := cert.Instance(sim.Client(), ca)

	// Specify certificate data.
	id := big.NewInt(5)
	ct := bindings.CertificatesOutput{
		Name:   "Catherine",
		Course: "MBCC",
		Grade:  "S",
		Date:   "2025-08-11",
	}

	// Issue certificate using the owner account. This should succeed.
	if _, err = bind.Transact(instance, ownerAuth, cert.PackIssue(id, ct.Name, ct.Course, ct.Grade, ct.Date)); err != nil {
		t.Fatalf("Failed to call issue method: %v", err)
	}

	// Commit the blocks so far.
	sim.Commit()

	// Issue certificate using the other account. This should fail.
	_, err = bind.Transact(instance, otherAuth, cert.PackIssue(id, ct.Name, ct.Course, "C", ct.Date))
	assert.NotNil(t, err)

	// Filter logs for the deployed contract.
	logs, err := sim.Client().FilterLogs(context.Background(), ethereum.FilterQuery{Addresses: []common.Address{ca}})
	if err != nil {
		t.Fatalf("Failed to filter logs: %v", err)
	}

	// Unpack the first event and validate the parameters.
	ownershipTransferred, err := cert.UnpackOwnershipTransferredEvent(&logs[0])
	if err != nil {
		t.Fatalf("Failed to unpack OwnershipTransferred event: %v", err)
	}

	assert.Equal(t, ownershipTransferred.PreviousOwner, common.Address{})
	assert.Equal(t, ownershipTransferred.NewOwner, ownerAuth.From)

	// Unpack the second event and validate the parameters.
	issued, err := cert.UnpackIssuedEvent(&logs[1])
	if err != nil {
		t.Fatalf("Failed to unpack Issued event: %v", err)
	}

	assert.Equal(t, issued.Id, id)
	assert.Equal(t, issued.Date, ct.Date)

	// Fetch the certificate data from the contract and validate the values.
	val, err := bind.Call(instance, nil, cert.PackCertificates(id), cert.UnpackCertificates)
	if err != nil {
		t.Fatalf("Failed to call Certificates method: %v", err)
	}

	if assert.NotNil(t, val) {
		assert.EqualValues(t, val, ct)
	}

	// Transfer ownership to the other account.
	if _, err = bind.Transact(instance, ownerAuth, cert.PackTransferOwnership(otherAuth.From)); err != nil {
		t.Fatalf("Failed to call transferOwnership method: %v", err)
	}

	// Commit the blocks so far.
	sim.Commit()

	// Owner should be the other account by this point.
	owner, err := bind.Call(instance, nil, cert.PackOwner(), cert.UnpackOwner)
	if err != nil {
		t.Fatalf("Failed to call owner method: %v", err)
	}

	assert.Equal(t, otherAuth.From, owner)

	// Other account renounces the contract.
	if _, err = bind.Transact(instance, otherAuth, cert.PackRenounceOwnership()); err != nil {
		t.Fatalf("Failed to call renounceOwnership method: %v", err)
	}

	// Commit the blocks so far.
	sim.Commit()

	// Owner should be address(0) by this point.
	owner, err = bind.Call(instance, nil, cert.PackOwner(), cert.UnpackOwner)
	if err != nil {
		t.Fatalf("Failed to call owner method: %v", err)
	}

	assert.Equal(t, common.Address{}, owner)
}
