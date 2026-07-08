package chain

import (
	"fmt"
	"strings"
	"time"

	"toyblockchain/block"
)

// Default mining difficulty.
// A valid block hash must start with this many zeros.
const DefaultDifficulty = 4

// MineBlock performs the Proof-of-Work algorithm.
// It repeatedly increments the nonce until the
// block hash satisfies the required difficulty.
func MineBlock(b *block.Block, difficulty int) {

	// Target prefix (example: "0000")
	target := strings.Repeat("0", difficulty)

	fmt.Println("\n===================================")
	fmt.Println("Mining Block...")
	fmt.Printf("Difficulty : %d leading zeros\n", difficulty)

	start := time.Now()

	for {

		// Calculate current hash.
		hash := b.CalculateHash()

		// Check difficulty requirement.
		if strings.HasPrefix(hash, target) {

			b.Hash = hash
			break
		}

		// Try next nonce.
		b.Nonce++
	}

	elapsed := time.Since(start)

	fmt.Println("Mining Successful!")
	fmt.Printf("Nonce        : %d\n", b.Nonce)
	fmt.Printf("Hash         : %s\n", b.Hash)
	fmt.Printf("Mining Time  : %v\n", elapsed)
	fmt.Println("===================================")
}
