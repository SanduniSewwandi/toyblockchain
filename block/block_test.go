package block

import (
	"testing"

	"toyblockchain/ledger"
)

// Test NewBlock creation
func TestNewBlock(t *testing.T) {

	txs := []ledger.Transaction{
		{
			Sender:   "Alice",
			Receiver: "Bob",
			Amount:   10,
		},
	}

	block := NewBlock(
		1,
		txs,
		"0000previoushash",
	)

	// Check index
	if block.Index != 1 {
		t.Errorf(
			"Expected index 1, got %d",
			block.Index,
		)
	}

	// Check previous hash
	if block.PreviousHash != "0000previoushash" {

		t.Errorf(
			"Previous hash incorrect",
		)
	}

	// Check transactions
	if len(block.Transactions) != 1 {

		t.Errorf(
			"Expected 1 transaction, got %d",
			len(block.Transactions),
		)
	}

	// Hash should not be empty
	if block.Hash == "" {

		t.Errorf(
			"Block hash should not be empty",
		)
	}

	// Nonce starts at zero
	if block.Nonce != 0 {

		t.Errorf(
			"Expected nonce 0, got %d",
			block.Nonce,
		)
	}
}

// Test hash generation
func TestCalculateHash(t *testing.T) {

	block := Block{

		Index: 1,

		Timestamp: 1000,

		Transactions: []ledger.Transaction{

			{
				Sender:   "Alice",
				Receiver: "Bob",
				Amount:   5,
			},
		},

		PreviousHash: "abc",

		Nonce: 0,
	}

	hash1 := block.CalculateHash()

	if hash1 == "" {

		t.Error(
			"Hash should not be empty",
		)
	}

	// Same data should produce same hash
	hash2 := block.CalculateHash()

	if hash1 != hash2 {

		t.Error(
			"Hash should be deterministic",
		)
	}
}

// Test changing block data changes hash
func TestHashChangesWhenDataChanges(t *testing.T) {

	block := Block{

		Index: 1,

		Timestamp: 1000,

		Transactions: []ledger.Transaction{

			{
				Sender:   "Alice",
				Receiver: "Bob",
				Amount:   10,
			},
		},

		PreviousHash: "abc",

		Nonce: 0,
	}

	oldHash := block.CalculateHash()

	// Modify transaction
	block.Transactions[0].Amount = 20

	newHash := block.CalculateHash()

	if oldHash == newHash {

		t.Error(
			"Hash should change after block modification",
		)
	}
}

// Test that Hash field is not included
func TestHashFieldNotIncluded(t *testing.T) {

	block := Block{

		Index: 1,

		Timestamp: 1000,

		Transactions: []ledger.Transaction{},

		PreviousHash: "abc",

		Nonce: 5,

		Hash: "randomhash",
	}

	hash1 := block.CalculateHash()

	block.Hash = "anotherhash"

	hash2 := block.CalculateHash()

	if hash1 != hash2 {

		t.Error(
			"Hash calculation should ignore Hash field",
		)
	}
}
