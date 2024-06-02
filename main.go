package main

import (
	"fmt"
)

func main() {
	// Create a blockchain
	myBlockchain := NewBlockchain()

	// Add blocks to the new blockchain
	myBlockchain.AddBlock("Send 1 ETH to Joseca")
	myBlockchain.AddBlock("Send 2 ETH to Giove")

	// Print attributes from each block
	for _, block := range myBlockchain.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println()
	}
}
