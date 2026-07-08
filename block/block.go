package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"toyblockchain/ledger"
)

// Block represents a single block in the blockchain.
type Block struct {
	Index        int
	Timestamp    int64
	Transactions []ledger.Transaction
	PreviousHash string
	Nonce        int
	Hash         string
}

// CalculateHash computes the SHA-256 hash of the block.
//
// The following fields are included in this order:
// 1. Index
// 2. Timestamp
// 3. Transactions
// 4. PreviousHash
// 5. Nonce
//
// The Hash field itself is NOT included.
func (b *Block) CalculateHash() string {

	data := struct {
		Index        int
		Timestamp    int64
		Transactions []ledger.Transaction
		PreviousHash string
		Nonce        int
	}{
		Index:        b.Index,
		Timestamp:    b.Timestamp,
		Transactions: b.Transactions,
		PreviousHash: b.PreviousHash,
		Nonce:        b.Nonce,
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(jsonBytes)

	return hex.EncodeToString(hash[:])
}

// NewBlock creates a new block.
//
// The block starts with Nonce = 0.
// Mining will later modify the nonce until the
// required Proof-of-Work difficulty is satisfied.
func NewBlock(index int, txs []ledger.Transaction, previousHash string) Block {

	newBlock := Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: txs,
		PreviousHash: previousHash,
		Nonce:        0,
	}

	newBlock.Hash = newBlock.CalculateHash()

	return newBlock
}
