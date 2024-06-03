package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
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
  block.Hash = hash

	// Return block
	return block
}

// Create Genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Serialise the block
func (b *Block) SerialiseBlock() []byte {
  // Create a buffer to hold the block information
  var buffer bytes.Buffer
  // Create an encoder that saves encodings into the buffer
  encoder := gob.NewEncoder(&buffer)
  // Encode the block, which will be saved in buffer. Returns error
  err := encoder.Encode(b)
  if err != nil {
    fmt.Print("Error while encoding block")
    os.Exit(1)
  }
  // Convert the enconding into a slices of bytes
  return buffer.Bytes()
}


// Deserialise the block
func  SerialiseBlock(d []byte) *Block {
  // Create a block var where the data will be deserialised
  var block Block
  // Create a deserialiser that contains a reader with the data
  decoder := gob.NewDecoder(bytes.NewReader(d))
  // Decode data into the block var
  err := decoder.Decode(&block)
  if err != nil {
    fmt.Print("Error while dencoding block")
    os.Exit(1)
  }
  // Return the address of the block with the decoded data
  return &block
}
