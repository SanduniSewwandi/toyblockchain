# Toy Blockchain and Ledger Simulator

A simple blockchain and ledger simulator implemented in **Go**, developed as part of the **Backend Engineering Internship – Golang Developer Assessment**.

The project demonstrates the core concepts behind blockchain technology, including deterministic hashing, Proof-of-Work (PoW) mining, blockchain validation, ledger reconstruction, transaction validation, persistence, and command-line interaction.

Rather than building a production blockchain, the objective is to implement the essential blockchain mechanisms in a clean, modular, and well-tested Go application.

---

# Features

- Deterministic Genesis Block
- Genesis Coinbase (Faucet) Transactions
- SHA-256 Deterministic Block Hashing
- Proof-of-Work (PoW) Mining
- Configurable Default Mining Difficulty
- Difficulty Stored Per Block
- Transaction Validation
- Ledger Reconstructed from Blockchain Transactions
- Overspending Protection
- Full Blockchain Validation
- Tamper Detection
- Pending Transaction Pool
- JSON Persistence
- Command-Line Interface (CLI)
- Comprehensive Unit Tests

---

# Project Structure

```
toyblockchain/
│
├── block/
│   ├── block.go
│   └── block_test.go
│
├── chain/
│   ├── blockchain.go
│   ├── mining.go
│   ├── validation.go
│   ├── storage.go
│   ├── pending.go
│   ├── blockchain_test.go
│   ├── mining_test.go
│   ├── validation_test.go
│   └── storage_test.go
│
├── cli/
│   └── cli.go
│
├── ledger/
│   ├── ledger.go
│   ├── transaction.go
│   ├── ledger_test.go
│   └── transaction_test.go
│
├── main.go
├── blockchain.json
├── pending.json
├── go.mod
└── README.md
```

---

# Requirements

- Go 1.22 or newer

Check your Go installation:

```bash
go version
```

---

# Building the Project

Clone the repository and build the application.

```bash
go build
```

or run directly:

```bash
go run main.go
```

---

# Running the Application

## Show available commands

```bash
go run main.go
```

---

## Add a transaction

```bash
go run main.go add Alice Bob 20
```

The transaction is validated before being added to the pending transaction pool.

Validation includes:

- Positive amount
- Sender has sufficient balance
- Pending transactions are considered

---

## Mine pending transactions

```bash
go run main.go mine
```

Mining uses the configured default difficulty.

During mining the application displays:

- Difficulty
- Nonce
- Generated hash
- Mining time

After mining:

- A new block is appended to the blockchain.
- The blockchain is saved to disk.
- Pending transactions are cleared.

---

## Print the blockchain

```bash
go run main.go print
```

Example output:

```
SYSTEM -> Alice : 100
SYSTEM -> Bob   : 50
Alice  -> Bob   : 20
```

---

## Display account balances

```bash
go run main.go balance
```

Example:

```
Alice : 80
Bob   : 70
```

---

## Validate the blockchain

```bash
go run main.go validate
```

Example:

```
Valid : true

Message: Chain is valid
```

---

## Demonstration Mode

```bash
go run main.go demo
```

The demonstration automatically:

- Creates transactions
- Mines blocks
- Prints the blockchain
- Displays balances
- Demonstrates blockchain tampering
- Runs validation before and after tampering

---

# Running Tests

Run every unit test:

```bash
go test ./...
```

The test suite covers:

- Deterministic hashing
- Block creation
- Genesis block creation
- Genesis coinbase balances
- Mining
- Proof-of-Work validation
- Nonce reproduction
- Blockchain validation
- Tamper detection
- Transaction validation
- Overspending rejection
- Persistence

---

# Design Decisions

## Block Structure

Each block contains:

- Index
- Timestamp
- Transactions
- Previous Hash
- Nonce
- Difficulty
- Hash

The hash field is excluded from hash calculation.

---

## Genesis Block

The blockchain starts with a deterministic genesis block.

Its previous hash is a fixed value consisting of 64 zero characters.

Unlike later blocks, the genesis block contains two special coinbase (faucet) transactions that introduce the initial currency supply into the blockchain.

```
SYSTEM → Alice : 100

SYSTEM → Bob : 50
```

These transactions allow the ledger to reconstruct balances entirely from blockchain data.

---

## Deterministic Hashing

Each block hash is computed using SHA-256 over a deterministic JSON serialization of the following fields:

1. Index
2. Timestamp
3. Transactions
4. PreviousHash
5. Nonce
6. Difficulty

The Hash field itself is intentionally excluded.

This guarantees that hashing the same block twice always produces the same hash.

---

## Proof-of-Work Mining

Mining repeatedly changes the nonce until the generated SHA-256 hash satisfies the required difficulty target.

The default mining difficulty is defined by:

```go
chain.DefaultDifficulty
```

Each mined block stores the difficulty used during mining.

Mining reports:

- Difficulty
- Nonce
- Generated hash
- Mining time

---

## Ledger

The application does **not** permanently store balances.

Instead, every balance is reconstructed by replaying all blockchain transactions beginning from the genesis block.

The genesis block introduces the initial balances through coinbase transactions.

Current initial balances:

| Account | Initial Balance |
|----------|----------------:|
| Alice | 100 |
| Bob | 50 |

Transactions are rejected when:

- Amount ≤ 0
- Sender has insufficient balance

This design ensures that account balances are always derived from blockchain history rather than maintained separately.

---

## Blockchain Validation

Validation verifies the entire blockchain.

For every block it checks:

- Stored hash equals recalculated hash
- Previous hash links are correct
- Block indexes are sequential
- Timestamps are consistent
- Proof-of-Work satisfies the stored difficulty

Validation immediately reports the first invalid block if any error is detected.

---

## Persistence

Blockchain data is stored as JSON.

Files used:

```
blockchain.json
```

Stores every block.

```
pending.json
```

Stores pending transactions waiting to be mined.

Both files are automatically loaded when the application starts.




---

# Known Limitations

This project is intended as an educational blockchain simulator.

The following production blockchain features are intentionally omitted:

- Peer-to-peer networking
- Distributed consensus
- Public/private key cryptography
- Digital signatures
- Wallet management
- Merkle trees
- Smart contracts
- Automatic difficulty adjustment
- Mining rewards beyond genesis coinbase transactions
- Chain forks
- Fork resolution

---

# Technologies Used

- Go 1.22+
- Standard Library

Packages include:

- crypto/sha256
- encoding/json
- encoding/hex
- fmt
- os
- strings
- strconv
- testing
- time

No third-party libraries were used.

---

# Assessment Notes

This project was implemented to satisfy the Backend Engineering Internship Golang Developer Assessment.

Key implementation decisions include:

- Deterministic genesis block.
- Initial balances created using genesis coinbase transactions.
- Ledger reconstructed entirely from blockchain history.
- SHA-256 deterministic hashing.
- Proof-of-Work mining with configurable default difficulty.
- Difficulty stored inside each block.
- Validation uses each block's stored difficulty.
- Full blockchain persistence using JSON.
- Command-line interface for all required operations.
- Unit tests covering the required functionality.

---

# Author

Developed as part of the **Backend Engineering Internship – Golang Developer Assessment** using the Go programming language.