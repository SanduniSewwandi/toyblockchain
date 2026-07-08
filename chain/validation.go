package chain

import (
	"fmt"
	"strings"
)

// ValidateChain checks the integrity of the entire blockchain.
// It ensures:
// 1. Each block's hash is correct (no tampering)
// 2. Each block correctly links to the previous block
// 3. Each mined block satisfies Proof-of-Work difficulty
// 4. Block indexes are sequential
// 5. Timestamps are consistent
func (bc *Blockchain) ValidateChain() (bool, string) {

	// Use default mining difficulty
	target := strings.Repeat("0", DefaultDifficulty)

	for i := 0; i < len(bc.Blocks); i++ {

		current := bc.Blocks[i]

		// --------------------------------
		// 1. Verify hash integrity
		// --------------------------------
		if current.CalculateHash() != current.Hash {

			return false, fmt.Sprintf(
				"Block %d: hash mismatch (data tampered)",
				i,
			)
		}

		// --------------------------------
		// 2. Validate Genesis Block
		// --------------------------------
		if i == 0 {

			if current.Index != 0 {

				return false,
					"Genesis block has invalid index"
			}

			continue
		}

		previous := bc.Blocks[i-1]

		// --------------------------------
		// 3. Verify previous hash link
		// --------------------------------
		if current.PreviousHash != previous.Hash {

			return false, fmt.Sprintf(
				"Block %d: invalid previous hash link",
				i,
			)
		}

		// --------------------------------
		// 4. Verify block index/height
		// --------------------------------
		if current.Index != previous.Index+1 {

			return false, fmt.Sprintf(
				"Block %d: invalid block index",
				i,
			)
		}

		// --------------------------------
		// 5. Verify timestamp order
		// --------------------------------
		if current.Timestamp < previous.Timestamp {

			return false, fmt.Sprintf(
				"Block %d: invalid timestamp",
				i,
			)
		}

		// --------------------------------
		// 6. Verify Proof-of-Work
		// --------------------------------
		if !strings.HasPrefix(current.Hash, target) {

			return false, fmt.Sprintf(
				"Block %d: invalid proof-of-work",
				i,
			)
		}
	}

	return true, "Chain is valid"
}
