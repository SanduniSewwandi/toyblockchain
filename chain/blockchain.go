package chain

import (
	"fmt"

	"toyblockchain/block"
	"toyblockchain/ledger"
)

// GenesisPreviousHash is the fixed previous hash used by the genesis block.
const GenesisPreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"


type Blockchain struct {
	Blocks []block.Block
}


func NewBlockchain() *Blockchain {

	genesis := block.NewBlock(
		0,
		[]ledger.Transaction{},
		GenesisPreviousHash,
	)

	return &Blockchain{
		Blocks: []block.Block{genesis},
	}
}

// GetLatestBlock returns the newest block.
func (bc *Blockchain) GetLatestBlock() block.Block {

	if len(bc.Blocks) == 0 {
		panic("blockchain is empty")
	}

	return bc.Blocks[len(bc.Blocks)-1]
}


func (bc *Blockchain) AddBlock(
	transactions []ledger.Transaction,
	difficulty int,
) error {

	latest := bc.GetLatestBlock()

	newBlock := block.NewBlock(
		latest.Index+1,
		transactions,
		latest.Hash,
	)

	// Mine block using selected difficulty.
	MineBlock(
		&newBlock,
		difficulty,
	)

	// Append mined block.
	bc.Blocks = append(
		bc.Blocks,
		newBlock,
	)

	// Save blockchain.
	if err := bc.SaveToFile(DefaultBlockchainFile); err != nil {
		return err
	}

	return nil
}

// Print displays the blockchain.
func (bc *Blockchain) Print() {

	fmt.Println("\n========== BLOCKCHAIN ==========")

	for _, b := range bc.Blocks {

		fmt.Println("--------------------------------")
		fmt.Printf("Index         : %d\n", b.Index)
		fmt.Printf("Timestamp     : %d\n", b.Timestamp)
		fmt.Printf("Previous Hash : %s\n", b.PreviousHash)
		fmt.Printf("Hash          : %s\n", b.Hash)
		fmt.Printf("Nonce         : %d\n", b.Nonce)

		fmt.Println("Transactions:")

		if len(b.Transactions) == 0 {

			fmt.Println("  No transactions")

		} else {

			for i, tx := range b.Transactions {

				fmt.Printf(
					"  %d. %s -> %s : %.2f\n",
					i+1,
					tx.Sender,
					tx.Receiver,
					tx.Amount,
				)
			}
		}

		fmt.Println("--------------------------------")
	}
}

// BuildLedger reconstructs account balances from the blockchain.
func (bc *Blockchain) BuildLedger() *ledger.Ledger {

	ld := ledger.NewLedger()

	// Initial balances (faucet)
	ld.Credit("Alice", 100)
	ld.Credit("Bob", 50)

	// Replay all transactions from blockchain
	for _, b := range bc.Blocks {

		for _, tx := range b.Transactions {

			if err := ld.ApplyTransaction(tx); err != nil {

				fmt.Println("Warning:", err)

			}
		}
	}

	return ld
}
