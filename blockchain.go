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
	db  *bbolt.DB
}

// OLD Add block to the blockchain
// func (bc *Blockchain) AddBlock(data string) {
// 	// Get the previous block
// 	prevBlock := bc.blocks[len(bc.blocks)-1]
// 	newBlock := NewBlock(data, prevBlock.Hash)
// 	bc.blocks = append(bc.blocks, newBlock)
// }

// Add a block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	// Create var to save the hash of the last block
	var lastHash []byte

	// Get the last hash
	err := bc.db.View(func(tx *bbolt.Tx) error {
		// Get the bucket where the blockchain is
		b := tx.Bucket([]byte(blocksBucket))
		// Get the lastHash
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		fmt.Print("Error while getting lastHash")
	}

	// Create a newBlock
	newBlock := NewBlock(data, lastHash)

	// Add newBlock + update lastHash to bucket
	err = bc.db.Update(func(tx *bbolt.Tx) error {
		// Get the bucket that contains the blocks
		b := tx.Bucket([]byte(blocksBucket))
		// Add the new block to the bucket
		err = b.Put(newBlock.Hash, newBlock.SerialiseBlock())
		if err != nil {
			fmt.Print("Error while adding the newBlock to database")
		}
		// Update the last hash of the blockchain
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			fmt.Print("Error while updating the last hash in the database")
		}
		// Update the tip of the blockchain struct
		bc.tip = newBlock.Hash

		return nil
	})

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

// Blockchain Iterator struct
type BlockchainIterator struct {
	currentHash []byte
	db          *bbolt.DB
}

// Create a new blockchain iterator
func (bc *Blockchain) NewBlockchainIterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// Save and return the next block in the iterator
func (i *BlockchainIterator) Next() *Block {
  // Create a block var to save the next block
  var block *Block
  // Get the current block that the Iterator is pointing to
  err := i.db.View( func(tx *bbolt.Tx) error {
    // Open bucket
    b := tx.Bucket([]byte(blocksBucket))
    // Get the block data (in bytes) that iterator is pointing to
    encodedBlock := b.Get(i.currentHash)
    // Deserialise the block
    block = DeserialiseBlock(encodedBlock)

    return nil
  })
  if err != nil {
    fmt.Print("Error while taking iterator block from DB")
  }

  // Update the iterator to get the next block
  i.currentHash = block.PrevHash

  return block
}
