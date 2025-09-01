package tests

import (
	"ethereum-developer-program/dao/lib/bindings"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
)

// Custom struct to ease the usage.
type govToken struct {
	address  *common.Address
	sim      *simulated.Backend
	owner    *bind.TransactOpts
	other    *bind.TransactOpts
	contract *bindings.GovToken
	instance *bind.BoundContract
}

// Method to advance the simulated backend with a given number of blocks.
func (g *govToken) skipBlocks(n int) {
	for range n {
		g.sim.Commit()
	}
}

// Function to initialize the simulated backend, multiple auths, contract instance and wrapper for GovToken.
func initializeGovToken() (*govToken, error) {
	// Create a random key for the deployer/owner account.
	ownerKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate key 1: %v", err)
	}

	// Create a random key for the other account.
	otherKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate key 2: %v", err)
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
	contract := bindings.NewGovToken()

	// Deploy the smart contract.
	ca, _, err := bind.DeployContract(ownerAuth, common.FromHex(bindings.GovTokenMetaData.Bin), sim.Client(), contract.PackConstructor(ownerAuth.From))
	if err != nil {
		return nil, fmt.Errorf("failed to deploy smart contract: %v", err)
	}

	// Commit the blocks so far.
	sim.Commit()

	return &govToken{
		address:  &ca,
		sim:      sim,
		owner:    ownerAuth,
		other:    otherAuth,
		contract: contract,
		instance: contract.Instance(sim.Client(), ca),
	}, nil
}

func TestGovTokenDeployment(t *testing.T) {
	g, err := initializeGovToken()
	if err != nil {
		t.Fatalf("Failed to initialize GovToken: %v", err)
	}

	// Owner should be the owner account by this point.
	owner, err := bind.Call(g.instance, nil, g.contract.PackOwner(), g.contract.UnpackOwner)
	if err != nil {
		t.Fatalf("Failed to call owner method: %v", err)
	}

	assert.Equal(t, g.owner.From, owner)

	// Total supply should be 1000000000000000000000.
	supply, err := bind.Call(g.instance, nil, g.contract.PackTotalSupply(), g.contract.UnpackTotalSupply)
	if err != nil {
		t.Fatalf("Failed to call totalSupply method: %v", err)
	}

	// Balance should be equal to the total supply by this point.
	balance, err := bind.Call(g.instance, nil, g.contract.PackBalanceOf(g.owner.From), g.contract.UnpackBalanceOf)
	if err != nil {
		t.Fatalf("Failed to call balanceOf method: %v", err)
	}

	assert.Equal(t, supply, balance)
}

func TestGovTokenTransfers(t *testing.T) {
	g, err := initializeGovToken()
	if err != nil {
		t.Fatalf("Failed to initialize GovToken: %v", err)
	}

	// Set up a transfer amount.
	transferVal := big.NewInt(100)

	// Transfer the amount from the owner account to the other account.
	if _, err = bind.Transact(g.instance, g.owner, g.contract.PackTransfer(g.other.From, transferVal)); err != nil {
		t.Fatalf("Failed to call transfer method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Check the balance of the other account. This is should be equal to the transfer amount.
	balance, err := bind.Call(g.instance, nil, g.contract.PackBalanceOf(g.other.From), g.contract.UnpackBalanceOf)
	if err != nil {
		t.Fatalf("Failed to call balanceOf method: %v", err)
	}

	assert.Equal(t, transferVal, balance)

	// Set up a transfer account.
	transferAccount := common.Address{0x3, 0x4, 0x5} // This means '0x0304050000000000000000000000000000000000'.

	// Transfer the amount from the other account to the transfer account.
	if _, err = bind.Transact(g.instance, g.other, g.contract.PackTransfer(transferAccount, transferVal)); err != nil {
		t.Fatalf("Failed to call transfer method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Check the balance of the transfer account. This is should be equal to the transfer amount.
	balance, err = bind.Call(g.instance, nil, g.contract.PackBalanceOf(transferAccount), g.contract.UnpackBalanceOf)
	if err != nil {
		t.Fatalf("Failed to call balanceOf method: %v", err)
	}

	assert.Equal(t, transferVal, balance)
}

func TestGovTokenDelegates(t *testing.T) {
	g, err := initializeGovToken()
	if err != nil {
		t.Fatalf("Failed to initialize GovToken: %v", err)
	}

	// Total supply should be 1000000000000000000000.
	supply, err := bind.Call(g.instance, nil, g.contract.PackTotalSupply(), g.contract.UnpackTotalSupply)
	if err != nil {
		t.Fatalf("Failed to call totalSupply method: %v", err)
	}

	// Set up a transfer amount.
	transferVal := big.NewInt(100)

	// Transfer the amount from the owner account to the other account.
	if _, err = bind.Transact(g.instance, g.owner, g.contract.PackTransfer(g.other.From, transferVal)); err != nil {
		t.Fatalf("Failed to call transfer method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Owner account delegates its voting power. Before delegation, the voting power is 0.
	if _, err = bind.Transact(g.instance, g.owner, g.contract.PackDelegate(g.owner.From)); err != nil {
		t.Fatalf("Failed to call delegate method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// Other account delegates its voting power.
	if _, err = bind.Transact(g.instance, g.other, g.contract.PackDelegate(g.other.From)); err != nil {
		t.Fatalf("Failed to call delegate method: %v", err)
	}

	// Commit the blocks so far.
	g.sim.Commit()

	// After the successful transfer, the voting power of the owner account should be total supply - transfer amount.
	votes, err := bind.Call(g.instance, nil, g.contract.PackGetVotes(g.owner.From), g.contract.UnpackGetVotes)
	if err != nil {
		t.Fatalf("Failed to fetch getVotes method: %v", err)
	}

	assert.Equal(t, supply.Sub(supply, transferVal), votes)

	// After the successful receive, the voting power of the other account should be the transfer amount.
	votes, err = bind.Call(g.instance, nil, g.contract.PackGetVotes(g.other.From), g.contract.UnpackGetVotes)
	if err != nil {
		t.Fatalf("Failed to fetch getVotes method: %v", err)
	}

	assert.Equal(t, transferVal, votes)
}
