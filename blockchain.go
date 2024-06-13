package main

import (
	"errors"
	"fmt"
	"os"

	"go.etcd.io/bbolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// OLD Create Blockchain type
// type Blockchain struct {
// 	blocks []*Block
// }

type Blockchain struct {
  tip []byte
  db *bbolt.DB
}

// Add block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	// Get the previous block
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}


// Old Create a blockchain
// func NewBlockchain() *Blockchain {
// 	return &Blockchain{[]*Block{NewGenesisBlock()}}
// }

// New blockchain builder
func NewBlockchain() *Blockchain {
  // Create var for the tip of the blockchain
  var tip []byte
  // Open the database connection (file, fileMode, otherOptions)
  db, err := bbolt.Open(dbFile, 0600, nil)

  if err != nil {
    fmt.Print("Error while opening the DB connection in blockchain creation")
    os.Exit(1)
  }

  // Write/Update the database
  err = db.Update(func(tx *bbolt.Tx) error {
    // Get a bucket from the database
    b := tx.Bucket([]byte(blocksBucket))
    
    // If the bucket doesn't exist
    if b == nil {
      // Generate a new GenesisBlock
      genesis := NewGenesisBlock()
      // Create a bucket with the blocksBucket key
      b, err = tx.CreateBucket([]byte(blocksBucket))
      if err != nil {
        fmt.Print("Error while creating new bucket in NewBlockchain()")
      }
      // Add the block data with the block hash as its key
      err = b.Put(genesis.Hash, genesis.SerialiseBlock())
      if err != nil {
        fmt.Print("Error while putting block into bucket")
        return errors.New("Error while putting block into bucket")
      }
      // Add the last hash with "l" key
      err = b.Put([]byte("l"), genesis.Hash)
      if err != nil {
        fmt.Print("Error while putting last hash into bucket")
        return errors.New("Error while putting last hash into bucket")
      }

    } else {
      // Get the last hash with the key "l"
      tip = b.Get([]byte("l"))
    }
    return nil

  })

  // Create a new blockchain pointer
  bc := &Blockchain{tip, db}

  return bc
}
