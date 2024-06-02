package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Create Block type
type Block struct {
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
	Data      []byte
}

// Set the hash of a block
func (b *Block) SetHash() {
	// Convert timestamp from int64 to byte slice. First convert to string and then to []byte
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// Join together Timestamp, PrevHash and Data to create a new header
	header := bytes.Join([][]byte{timestamp, b.PrevHash, b.Data}, []byte{})
	// Create a hash from the header. This creates an array of 32bytes
	hash := sha256.Sum256(header)
	// Put hash into the block, converting it before to a slice of bytes
	b.Hash = hash[:]
}

// Create a new Block
func NewBlock(data string, prevHash []byte) *Block {
	// Create a pointer to a new Block with the data
	block := &Block{time.Now().Unix(), prevHash, []byte{}, []byte(data)}
	// Compute the hash with setHash()
	block.SetHash()
	// Return block
	return block
}

// Create Genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
