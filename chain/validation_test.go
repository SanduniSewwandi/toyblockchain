package chain

import (
	"testing"

	"toyblockchain/ledger"
)

// Test valid blockchain
func TestValidateValidChain(t *testing.T) {

	bc := NewBlockchain()

	tx := ledger.Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   10,
	}

	err := bc.AddBlock(
		[]ledger.Transaction{tx},
		DefaultDifficulty,
	)

	if err != nil {
		t.Fatal(err)
	}

	valid, message := bc.ValidateChain()

	if !valid {

		t.Errorf(
			"Expected valid chain, got invalid: %s",
			message,
		)
	}

}

// Test tampering detection
func TestValidateDetectsTampering(t *testing.T) {

	bc := NewBlockchain()

	tx := ledger.Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   10,
	}

	bc.AddBlock(
		[]ledger.Transaction{tx},
		DefaultDifficulty,
	)

	// Modify transaction data
	bc.Blocks[1].Transactions[0].Amount = 999

	valid, message := bc.ValidateChain()

	if valid {

		t.Error(
			"Blockchain should be invalid after tampering",
		)
	}

	if message == "" {

		t.Error(
			"Validation should return error message",
		)
	}

}

// Test invalid previous hash link
func TestValidateInvalidPreviousHash(t *testing.T) {

	bc := NewBlockchain()

	tx := ledger.Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   10,
	}

	bc.AddBlock(
		[]ledger.Transaction{tx},
		DefaultDifficulty,
	)

	// Break chain link
	bc.Blocks[1].PreviousHash = "wrong_hash"

	valid, _ := bc.ValidateChain()

	if valid {

		t.Error(
			"Blockchain should fail previous hash validation",
		)
	}

}

// Test invalid block index
func TestValidateInvalidIndex(t *testing.T) {

	bc := NewBlockchain()

	tx := ledger.Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   10,
	}

	bc.AddBlock(
		[]ledger.Transaction{tx},
		DefaultDifficulty,
	)

	// Change block height
	bc.Blocks[1].Index = 5

	valid, _ := bc.ValidateChain()

	if valid {

		t.Error(
			"Blockchain should fail index validation",
		)
	}

}

// Test invalid timestamp order
func TestValidateInvalidTimestamp(t *testing.T) {

	bc := NewBlockchain()

	tx := ledger.Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   10,
	}

	bc.AddBlock(
		[]ledger.Transaction{tx},
		DefaultDifficulty,
	)

	// Set future block timestamp lower than previous
	bc.Blocks[1].Timestamp = bc.Blocks[0].Timestamp - 100

	valid, _ := bc.ValidateChain()

	if valid {

		t.Error(
			"Blockchain should fail timestamp validation",
		)
	}

}

// Test invalid proof of work
func TestValidateInvalidProofOfWork(t *testing.T) {

	bc := NewBlockchain()

	tx := ledger.Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   10,
	}

	bc.AddBlock(
		[]ledger.Transaction{tx},
		DefaultDifficulty,
	)

	// Change hash so it does not satisfy difficulty
	bc.Blocks[1].Hash = "12345"

	valid, _ := bc.ValidateChain()

	if valid {

		t.Error(
			"Blockchain should fail proof-of-work validation",
		)
	}

}
