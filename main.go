package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create a blockchain
	myBlockchain := NewBlockchain()

	// Add blocks to the new blockchain
	myBlockchain.AddBlock("Send 1 ETH to Joseca")
	myBlockchain.AddBlock("Send 2 ETH to Giove")

	// Print attributes from each block
	for _, block := range myBlockchain.blocks {

    // Create a new pow
    pow := NewProofOfWork(block)

		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
    fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
