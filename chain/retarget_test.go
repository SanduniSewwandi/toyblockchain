package chain

import (
	"testing"

	"toyblockchain/block"
)

func createTestBlockchain(difficulty int, timestamps []int64) *Blockchain {

	bc := &Blockchain{}

	for i, timestamp := range timestamps {

		b := block.Block{
			Index:      i,
			Timestamp:  timestamp,
			Difficulty: difficulty,
		}

		bc.Blocks = append(bc.Blocks, b)
	}

	return bc
}

func TestDifficultyIncreaseWhenBlocksAreFast(t *testing.T) {

	bc := createTestBlockchain(4, []int64{
		0,
		1,
		2,
		3,
	})

	next := CalculateNextDifficulty(bc)

	if next != 5 {
		t.Fatalf("expected difficulty 5, got %d", next)
	}
}

func TestDifficultyDecreaseWhenBlocksAreSlow(t *testing.T) {

	bc := createTestBlockchain(5, []int64{
		0,
		5,
		10,
		20,
	})

	next := CalculateNextDifficulty(bc)

	if next != 4 {
		t.Fatalf("expected difficulty 4, got %d", next)
	}
}

func TestDifficultyDoesNotGoBelowMinimum(t *testing.T) {

	bc := createTestBlockchain(MinDifficulty, []int64{
		0,
		10,
		20,
		30,
	})

	next := CalculateNextDifficulty(bc)

	if next != MinDifficulty {
		t.Fatalf(
			"expected minimum difficulty %d, got %d",
			MinDifficulty,
			next,
		)
	}
}

func TestNoRetargetBeforeInterval(t *testing.T) {

	bc := createTestBlockchain(4, []int64{
		0,
		1,
		2,
	})

	next := CalculateNextDifficulty(bc)

	if next != 4 {
		t.Fatalf("expected difficulty unchanged, got %d", next)
	}
}

func TestNextDifficultyForHonorsRequestBeforeRetargeting(t *testing.T) {

	bc := createTestBlockchain(4, []int64{
		0,
		1,
	})

	got := NextDifficultyFor(bc, 7)

	if got != 7 {
		t.Fatalf(
			"expected requested difficulty 7 to be honored before retargeting kicks in, got %d",
			got,
		)
	}
}

func TestNextDifficultyForIgnoresRequestAfterRetargeting(t *testing.T) {

	bc := createTestBlockchain(4, []int64{
		0,
		1,
		2,
		3,
	})

	got := NextDifficultyFor(bc, 7)

	if got == 7 {
		t.Fatal(
			"expected requested difficulty to be overridden by retargeting once enough history exists",
		)
	}
}
