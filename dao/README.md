# Sample DAO Project

Sample project of testing, deploying and interacting with DAO smart contract.

## 🛠 Built With

[![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)](https://go.dev/)
[![Geth Badge](https://img.shields.io/badge/Geth-3C3C3D?logo=ethereum&logoColor=fff&style=for-the-badge)](https://geth.ethereum.org/)
[![Solidity Badge](https://img.shields.io/badge/Solidity-363636?logo=solidity&logoColor=fff&style=for-the-badge)](https://soliditylang.org/)
[![OpenZeppelin Badge](https://img.shields.io/badge/OpenZeppelin-4E5EE4?logo=openzeppelin&logoColor=fff&style=for-the-badge)](https://docs.openzeppelin.com/)
[![React Badge](https://img.shields.io/badge/React-61DAFB?logo=react&logoColor=000&style=for-the-badge)](https://react.dev/)

## 📋 Prerequisites

- Go
- Make
- Any JavaScript runtime (preferably Node.js)

## ⚙️ Run Locally

### Backend

Compile the contracts using '**solc-js**'.

```sh
make compile
```

Generate Go bindings for the contracts using '**abigen**'.

```sh
make bindings
```

**Note**: Remove/comment the redeclared `Eip712DomainOutput` type definition in either [GovToken.go](./lib/bindings/GovToken.go) or [MyGovernor.go](./lib/bindings/MyGovernor.go). This is mandatory to move forward.

Run tests for the contracts.

```sh
make test
```

Export the environmental variables.

```sh
export RPC_URL=TBD
export PRIVATE_KEY=TBD
```

Deploy the contracts, apply the delegation and check the voting power.

```sh
make deploy
```

### Frontend

TODO
