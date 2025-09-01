# Go
GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")

# Constants
CONTRACTS = Cert GovToken MyGovernor TimeLock

# Directory paths
CURRENT_DIR := $(shell pwd)
BIN_DIR := $(CURRENT_DIR)/bin
CONTRACTS_DIR := $(CURRENT_DIR)/lib/contracts
ARTIFACTS_DIR := $(CURRENT_DIR)/lib/artifacts
BINDINGS_DIR := $(CURRENT_DIR)/lib/bindings

.PHONY: compile
## compile: Compile the contracts using solc.
compile:
	@cd $(CONTRACTS_DIR) && { \
		if [ ! -d "node_modules" ]; then \
			echo "@openzeppelin/contracts not found, running npm install..."; \
			npm install; \
			echo ""; \
		fi; \
	}
	@cd cmd/solc && { \
		if [ ! -d "node_modules" ]; then \
			echo "solc not found, running npm install..."; \
			npm install; \
			echo ""; \
		fi; \
		node .; \
	}

.PHONY: bindings
## bindings: Generate Go bindings of contracts using abigen.
bindings:
	@if [ ! -f "$(BIN_DIR)/abigen" ]; then \
		echo "Installing abigen..."; \
		mkdir -p $(BIN_DIR); \
		GOBIN=$(BIN_DIR) go install github.com/ethereum/go-ethereum/cmd/abigen@latest; \
	fi
	@mkdir -p $(BINDINGS_DIR)
	@for contract in $(CONTRACTS); do \
		echo "Generating binding for $$contract..."; \
		$(BIN_DIR)/abigen --v2 \
		--bin $(ARTIFACTS_DIR)/$$contract.bin \
		--abi $(ARTIFACTS_DIR)/$$contract.json \
		--pkg bindings --type $$contract \
		--out $(BINDINGS_DIR)/$$contract.go; \
	done
	@echo 'Done!'

.PHONY: test
## test: Run tests for contracts.
test:
	@$(GO) test -v ./...

.PHONY: deploy
## deploy: Deploy contracts.
deploy:
	@RPC_URL=$${RPC_URL:-http://127.0.0.1:8545} \
	PRIVATE_KEY=$${PRIVATE_KEY:-ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80} \
	$(GO) run ./cmd/contracts/deployer.go

.PHONY: tidy
## tidy: Clean and tidy dependencies.
tidy:
	@$(GO) mod tidy -v

.PHONY: fmt
## fmt: Format Go files.
fmt:
	@$(GOFMT) -w $(GOFILES)

.PHONY: help
## help: Display all targets.
help:
	@echo ''
	@echo 'Usage:'
	@echo '  make [target]'
	@echo ''
	@echo 'Targets:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.DEFAULT_GOAL := help