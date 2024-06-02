package main

import (
	"time"
)

// Create Block type
type Block struct {
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
	Data      []byte
	Nonce     int
}

// Now hash is set with ProofOfWork!!!

// Create a new Block
func NewBlock(data string, prevHash []byte) *Block {
	// Create a pointer to a new Block with the data
	block := &Block{time.Now().Unix(), prevHash, []byte{}, []byte(data), 0}
	// Compute the hash and nonce with POW
  pow := NewProofOfWork(block)
  nonce, hash := pow.Run()

  block.Nonce = nonce
  block.Hash = hash[:]

	// Return block
	return block
}

// Create Genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
