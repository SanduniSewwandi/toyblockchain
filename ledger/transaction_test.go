package ledger

import "testing"

func TestTransactionCreation(t *testing.T) {

	tx := Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   50,
	}

	if tx.Sender != "Alice" {
		t.Errorf("expected sender Alice, got %s", tx.Sender)
	}

	if tx.Receiver != "Bob" {
		t.Errorf("expected receiver Bob, got %s", tx.Receiver)
	}

	if tx.Amount != 50 {
		t.Errorf("expected amount 50, got %f", tx.Amount)
	}

}

func TestTransactionEmptySender(t *testing.T) {

	tx := Transaction{
		Receiver: "Alice",
	}

	if tx.Sender != "" {
		t.Error("expected empty sender")
	}

	if tx.Receiver != "Alice" {
		t.Error("receiver mismatch")
	}

}

func TestTransactionInvalidAmount(t *testing.T) {

	tx := Transaction{
		Amount: -5,
	}

	if tx.Amount >= 0 {
		t.Error(
			"negative transaction amount should be invalid",
		)
	}

}
