package chain

import (
	"fmt"
	"strings"
)

func (bc *Blockchain) ValidateChain() (bool, string) {

	for i := 0; i < len(bc.Blocks); i++ {

		current := bc.Blocks[i]

		// Verify block hash
		if current.CalculateHash() != current.Hash {

			return false, fmt.Sprintf(
				"Block %d: hash mismatch (data tampered)",
				i,
			)
		}

		// Genesis block checks
		if i == 0 {

			if current.Index != 0 {

				return false,
					"Genesis block has invalid index"
			}

			if current.PreviousHash != GenesisPreviousHash {

				return false,
					"Genesis block has invalid previous hash"
			}

			continue
		}

		previous := bc.Blocks[i-1]

		// Verify previous hash link
		if current.PreviousHash != previous.Hash {

			return false, fmt.Sprintf(
				"Block %d: invalid previous hash link",
				i,
			)
		}

		// Verify block index
		if current.Index != previous.Index+1 {

			return false, fmt.Sprintf(
				"Block %d: invalid block index",
				i,
			)
		}

		// Verify timestamp order
		if current.Timestamp < previous.Timestamp {

			return false, fmt.Sprintf(
				"Block %d: invalid timestamp",
				i,
			)
		}

		// Verify Proof-of-Work using THIS BLOCK'S difficulty
		target := strings.Repeat(
			"0",
			current.Difficulty,
		)

		if !strings.HasPrefix(
			current.Hash,
			target,
		) {

			return false, fmt.Sprintf(
				"Block %d: invalid proof-of-work",
				i,
			)
		}
	}

	return true, "Chain is valid"
}
