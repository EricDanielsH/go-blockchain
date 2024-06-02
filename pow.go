package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// Set the POW target bits
const targetBits = 24
// Set maxNonce
const maxNonce = math.MaxInt64

// Create POW struct
type ProofOfWork struct {
  block *Block
  target *big.Int
}

// Create a new POW
func NewProofOfWork(b *Block) *ProofOfWork{
  // Create a big int
  num := big.NewInt(1)
  // Shift to the left 256 - target
  num.Lsh(num, uint(256 - targetBits))
  // Create a pointer to a new POW
  pow := &ProofOfWork{b, num}
  return pow
}

// Prepare data method with a nonce (counter)
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
  // Join Timestamp, PrevHash, Data, TargetBits and Nonce
  data := bytes.Join([][]byte{
    IntToHex(pow.block.Timestamp),
    pow.block.PrevHash,
    pow.block.Data,
    IntToHex(int64(targetBits)),
    IntToHex(int64(nonce)),
    }, []byte{})
    
  return data
}

// Run method
func (pow *ProofOfWork) Run() (int, []byte) {
  // Create loop that adds +1 to the nonce and checks
  var bigHash big.Int
  var hash [32]byte
  nonce := 0

  fmt.Printf("Currently printing block with data '%s'\n", pow.block.Data)
  for nonce < maxNonce {
    // Create data
    data:= pow.PrepareData(nonce)
    // Generate hash with the d
    hash = sha256.Sum256(data)
    // Insert hash into a bigInt
    bigHash.SetBytes(hash[:])

    if (bigHash.Cmp(pow.target) == -1) {
      break
    } else {
      nonce++
    }
  }
  return nonce, hash[:]
}

// Validate POW
func (pow *ProofOfWork) Validate() bool {
  var bigHash big.Int

  data := pow.PrepareData(pow.block.Nonce)
  hash := sha256.Sum256(data)
  bigHash.SetBytes(hash[:])
  isValid := bigHash.Cmp(pow.target) == -1


  return isValid
}
