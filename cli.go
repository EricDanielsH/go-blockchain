package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// CLI struct
type CLI struct {
  bc *Blockchain
}

// Helper method that prints the correct usage of the program
func (cli *CLI) printUsage() {
  fmt.Println("How to use the program: ")
  fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

// Helper method that validates the number of arguments in the program
func (cli *CLI) validateArgs() {
  // If there are less than two arguments ->
  if len(os.Args) < 2 {
    // Print the correct commands that user should use 
    cli.printUsage()
    // Close the program
    os.Exit(1)
  }
}

// Main runner method
func (cli *CLI) Run() {
  // Check if the number of arguments is correct
  cli.validateArgs()

  // Set the commands available with the flag package
  addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
  printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
  // Add a -data flag to the addBlock cmd that receives string inputs
    // name of flag + default value + description
  addBlockData := addBlockCmd.String("data", "", "Block data")

  // Pass(or Parse) the arguments to the different commands
  switch os.Args[1] {
    case "addblock":
      err := addBlockCmd.Parse(os.Args[2:])
      if err != nil {
        fmt.Println("Error parsing addblock args")
        os.Exit(1)
      }
    case "printchain":
      err := printChainCmd.Parse(os.Args[2:])
      if err != nil {
        fmt.Println("Error parsing printchain args")
        os.Exit(1)
      }

    // Wrong flag
    default:
      cli.printUsage()
      os.Exit(1)
  }

  // Run addBlock if it was populated
  if addBlockCmd.Parsed() {
    // Check that the data is not empty
    if *addBlockData == "" {
      cli.printUsage()
      os.Exit(1)
    }
    cli.addBlock(*addBlockData)
  }

  // Run printChain if it was populated
  if printChainCmd.Parsed() {
    cli.printChain()
  }

}

// Add block cli method
func (cli *CLI) addBlock(data string) {
  cli.bc.AddBlock(data)
  fmt.Println("New block added succesfully!")
}

// Print blockchain cli method
func (cli *CLI) printChain() {
  // Create a new iterator
  bci := cli.bc.NewBlockchainIterator()

  // Run through every block 
  for {
    block := bci.Next()

    fmt.Printf("Prev. Hash: %x\n", block.PrevHash)
    fmt.Printf("Hash: %x\n", block.Hash)
    fmt.Printf("Data: %s\n", block.Data)
    // Check the POW of this block
    pow := NewProofOfWork(block)
    fmt.Printf("Proof of Work: %s\n", strconv.FormatBool(pow.Validate()))
    fmt.Println()

    // Break the loop when the genesis block is reached
    if len(block.PrevHash) == 0 {
      fmt.Println("End of the blockchain!")
      break
    }
  }
}
