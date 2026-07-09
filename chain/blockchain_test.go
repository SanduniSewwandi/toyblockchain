package chain

import (
	"testing"
)

// TestGenesisCreatesInitialBalances verifies that the genesis block
// correctly creates the initial balances through SYSTEM transactions.
func TestGenesisCreatesInitialBalances(t *testing.T) {

	bc := NewBlockchain()

	ld := bc.BuildLedger()

	// Alice should receive 100 from the genesis block.
	if ld.GetBalance("Alice") != 100 {

		t.Errorf(
			"expected Alice balance 100, got %.2f",
			ld.GetBalance("Alice"),
		)
	}

	// Bob should receive 50 from the genesis block.
	if ld.GetBalance("Bob") != 50 {

		t.Errorf(
			"expected Bob balance 50, got %.2f",
			ld.GetBalance("Bob"),
		)
	}

	// Charlie should have no balance.
	if ld.GetBalance("Charlie") != 0 {

		t.Errorf(
			"expected Charlie balance 0, got %.2f",
			ld.GetBalance("Charlie"),
		)
	}
}
