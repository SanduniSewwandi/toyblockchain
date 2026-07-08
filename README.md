# Toy Blockchain and Ledger Simulator

A simple blockchain and ledger simulator implemented in **Go**, developed as part of the **Backend Engineering Internship – Golang Developer Assessment**.

This project demonstrates the core concepts of blockchain technology, including deterministic hashing, Proof-of-Work mining, transaction validation, blockchain validation, and persistent storage, while remaining small, readable, and easy to understand.

---

# Features

- Deterministic Genesis Block
- SHA-256 Block Hashing
- Proof-of-Work (PoW) Mining
- Configurable Mining Difficulty
- Transaction Validation
- Account Balance Ledger
- Overspending Protection
- Blockchain Validation
- Tamper Detection
- Pending Transaction Pool
- JSON-based Persistence
- Command-Line Interface
- Unit Tests

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
│   ├── storage.go
│   ├── pending.go
│   ├── validation.go
│   ├── mining_test.go
│   ├── storage_test.go
│   └── validation_test.go
│
├── cli/
│   └── cli.go
│
├── ledger/
│   ├── ledger.go
│   ├── transaction.go
│   ├── transaction_test.go
│   └── ledger_test.go
│
├── main.go
├── go.mod
├── blockchain.json
├── pending.json
└── README.md
```

---

# Requirements

- Go 1.22 or newer

Check your Go version:

```bash
go version
```

---

# Build

Clone the repository and build the application.

```bash
go build
```

or run directly without building:

```bash
go run main.go
```

---

# Running the Application

## Display available commands

```bash
go run main.go
```

---

## Add a transaction

```bash
go run main.go add Alice Bob 25
```

---

## Mine pending transactions

Default difficulty:

```bash
go run main.go mine
```

Custom difficulty:

```bash
go run main.go mine 5
```

---

## Print the blockchain

```bash
go run main.go print
```

---

## Validate the blockchain

```bash
go run main.go validate
```

---

## Display account balances

```bash
go run main.go balance
```

---

## Run the demonstration

```bash
go run main.go demo
```

The demo:

- Creates sample transactions
- Mines blocks
- Prints the blockchain
- Prints ledger balances
- Demonstrates tamper detection

---

# Running Tests

Execute all unit tests using:

```bash
go test ./...
```

---

# Design Decisions

## Blockchain

Each block contains:

- Index
- Timestamp
- Transactions
- Previous Hash
- Nonce
- Hash

The blockchain starts with a deterministic genesis block whose previous hash consists of 64 zero characters.

---

## Hashing

Each block hash is generated using SHA-256.

The following fields are included in the hash calculation:

1. Index
2. Timestamp
3. Transactions
4. PreviousHash
5. Nonce

The Hash field itself is intentionally excluded to ensure deterministic hashing.

---

## Mining

Proof-of-Work mining repeatedly increments the nonce until the generated SHA-256 hash begins with the required number of leading zeros.

The mining difficulty is configurable through the command line.

Mining reports:

- Difficulty
- Nonce
- Generated hash
- Mining time

---

## Ledger

Account balances are reconstructed directly from the blockchain by replaying every transaction from the genesis block.

Initial balances are introduced using a faucet mechanism:

- Alice: 100
- Bob: 50

Transactions are rejected if:

- Amount is less than or equal to zero
- Sender has insufficient balance

---

## Validation

Blockchain validation checks:

- Block hash integrity
- Previous hash links
- Sequential block indexes
- Timestamp consistency
- Proof-of-Work difficulty

Any tampering with a block causes validation to fail and identifies the first invalid block.

---

## Persistence

The application stores data in JSON files.

Files:

- `blockchain.json` — blockchain data
- `pending.json` — pending transaction pool

The blockchain and pending transactions are automatically loaded when the application starts.

---

# Known Limitations

This project is intended for educational purposes and does not implement several features found in production blockchains.

Not implemented:

- Peer-to-peer networking
- Distributed consensus
- Digital signatures
- Public/private key cryptography
- Wallet management
- Merkle Trees
- Smart contracts
- Automatic difficulty adjustment
- Fork resolution

---

# Technologies Used

- Go 1.22+
- Standard Library

Packages used include:

- crypto/sha256
- encoding/json
- encoding/hex
- os
- time
- strings
- testing

No third-party libraries were used.

---

# Author

Developed as part of the **Golang Developer Assessment** for a Backend Engineering Internship.