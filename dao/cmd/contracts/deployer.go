package main

import (
	"context"
	"ethereum-developer-program/dao/lib/bindings"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ctx = context.Background()

	// Variables to store the contract addresses.
	govTokenAddr   common.Address
	timeLockAddr   common.Address
	certAddr       common.Address
	myGovernorAddr common.Address
)

// Main config.
type deployerConfig struct {
	addr   common.Address
	client *ethclient.Client
	opts   *bind.TransactOpts
}

// Method to update the transaction authorization data.
func (d *deployerConfig) newOpts() error {
	nonce, err := d.client.PendingNonceAt(ctx, d.addr)
	if err != nil {
		return fmt.Errorf("failed to fetch nonce: %v", err)
	}

	d.opts.Nonce = big.NewInt(int64(nonce))
	d.opts.GasLimit = uint64(9200000)

	gasPrice, err := d.client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch gas price: %v", err)
	}
	d.opts.GasPrice = gasPrice

	return nil
}

// Method to deploy contracts.
func (d *deployerConfig) deployContract(metadata *bind.MetaData, input []byte) (common.Address, error) {
	if err := d.newOpts(); err != nil {
		return common.Address{}, err
	}

	deployParams := bind.DeploymentParams{
		Contracts: []*bind.MetaData{metadata},
		Inputs: map[string][]byte{
			metadata.ID: input,
		},
	}

	deployer := bind.DefaultDeployer(d.opts, d.client)
	result, err := bind.LinkAndDeploy(&deployParams, deployer)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy %s: %v", metadata.ID, err)
	}

	if _, err = bind.WaitDeployed(ctx, d.client, result.Txs[metadata.ID].Hash()); err != nil {
		return common.Address{}, fmt.Errorf("failed to confirm deployment for %s: %v", metadata.ID, err)
	}

	return result.Addresses[metadata.ID], nil
}

// Method used for 'write' operations with a contract.
func (d *deployerConfig) transact(wrapper *bind.BoundContract, data []byte) error {
	if err := d.newOpts(); err != nil {
		return err
	}

	tx, err := bind.Transact(wrapper, d.opts, data)
	if err != nil {
		return fmt.Errorf("failed to transact: %v", err)
	}

	if _, err := bind.WaitMined(ctx, d.client, tx.Hash()); err != nil {
		return fmt.Errorf("failed to confirm transaction: %v", err)
	}

	return nil
}

// Function to create the main config.
func initialize(rawurl, key string) (*deployerConfig, error) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	addr := crypto.PubkeyToAddress(privateKey.PublicKey)

	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect client: %v", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve chain ID: %v", err)
	}

	opts := bind.NewKeyedTransactor(privateKey, chainID)

	return &deployerConfig{
		addr:   addr,
		client: client,
		opts:   opts,
	}, nil
}

func main() {
	rpcURL := os.Getenv("RPC_URL")
	privateKey := os.Getenv("PRIVATE_KEY")

	config, err := initialize(rpcURL, privateKey)
	if err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m %v", err)
	}

	// Initialize a GovToken instance.
	govToken := bindings.NewGovToken()

	// Deploy GovToken contract.
	govTokenAddr, err = config.deployContract(
		&bindings.GovTokenMetaData,
		govToken.PackConstructor(config.addr),
	)
	if err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m %v", err)
	}

	// Initialize a TimeLock instance.
	timeLock := bindings.NewTimeLock()

	// Deploy TimeLock contract.
	timeLockAddr, err = config.deployContract(
		&bindings.TimeLockMetaData,
		timeLock.PackConstructor(
			big.NewInt(0),
			[]common.Address{config.addr},
			[]common.Address{config.addr},
			config.addr,
		),
	)
	if err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m %v", err)
	}

	// Initialize a Cert instance.
	cert := bindings.NewCert()

	// Deploy Cert contract.
	certAddr, err = config.deployContract(
		&bindings.CertMetaData,
		cert.PackConstructor(timeLockAddr),
	)
	if err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m %v", err)
	}

	// Initialize a Cert instance.
	myGovernor := bindings.NewMyGovernor()

	// Deploy MyGovernor contract.
	myGovernorAddr, err = config.deployContract(
		&bindings.CertMetaData,
		myGovernor.PackConstructor(govTokenAddr, timeLockAddr),
	)
	if err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m %v", err)
	}

	log.Printf("\033[32m[INFO]\033[0m GovToken deployed at: \033[34m%s\033[0m\n", govTokenAddr)
	log.Printf("\033[32m[INFO]\033[0m TimeLock deployed at: \033[34m%s\033[0m\n", timeLockAddr)
	log.Printf("\033[32m[INFO]\033[0m Cert deployed at: \033[34m%s\033[0m\n", certAddr)
	log.Printf("\033[32m[INFO]\033[0m MyGovernor deployed at: \033[34m%s\033[0m\n", myGovernorAddr)

	// Create a instance wrapper for the GovToken contract.
	govTokenInstance := govToken.Instance(config.client, govTokenAddr)

	// Fetch the balance of deployer account.
	balance, err := bind.Call(govTokenInstance, nil, govToken.PackBalanceOf(config.addr), govToken.UnpackBalanceOf)
	if err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m Failed to call balanceOf method: %v", err)
	}

	log.Printf("\033[32m[INFO]\033[0m Balance of deployer: \033[33m%s\033[0m\n", balance)

	// Fetch votes before delegation.
	votes, err := bind.Call(govTokenInstance, nil, govToken.PackGetVotes(config.addr), govToken.UnpackGetVotes)
	if err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m Failed to fetch getVotes method: %v", err)
	}

	log.Printf("\033[32m[INFO]\033[0m Votes of deployer before delegation: \033[33m%s\033[0m\n", votes)

	// Delegate voting power.
	if err := config.transact(govTokenInstance, govToken.PackDelegate(config.addr)); err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m %v", err)
	}

	// Fetch votes after delegation.
	votes, err = bind.Call(govTokenInstance, nil, govToken.PackGetVotes(config.addr), govToken.UnpackGetVotes)
	if err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m Failed to fetch getVotes method: %v", err)
	}

	log.Printf("\033[32m[INFO]\033[0m Votes of deployer after delegation: \033[33m%s\033[0m\n", votes)

	// Create the role identifiers.
	PROPOSER_ROLE := [32]byte(crypto.Keccak256([]byte("PROPOSER_ROLE")))
	EXECUTOR_ROLE := [32]byte(crypto.Keccak256([]byte("EXECUTOR_ROLE")))

	// Create a instance wrapper for the TimeLock contract.
	timeLockInstance := timeLock.Instance(config.client, timeLockAddr)

	// Grant 'PROPOSER_ROLE' to MyGovernor contract.
	if err := config.transact(timeLockInstance, timeLock.PackGrantRole(PROPOSER_ROLE, myGovernorAddr)); err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m %v", err)
	}

	log.Printf("\033[32m[INFO]\033[0m Granted \033[35m%s\033[0m to MyGovernor contract.", "PROPOSER_ROLE")

	// Grant 'EXECUTOR_ROLE' to MyGovernor contract.
	if err := config.transact(timeLockInstance, timeLock.PackGrantRole(EXECUTOR_ROLE, myGovernorAddr)); err != nil {
		log.Fatalf("\033[31m[ERROR]\033[0m %v", err)
	}

	log.Printf("\033[32m[INFO]\033[0m Granted \033[35m%s\033[0m to MyGovernor contract.", "EXECUTOR_ROLE")
}
