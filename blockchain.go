package main

// Create Blockchain type
type Blockchain struct {
	blocks []*Block
}

// Add block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	// Get the previous block
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}


// Create a blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
