package main

import (
	"flag"
	"fmt"
	"os"
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
    cli.AddBlock(*addBlockData)
  }

  // Run printChain if it was populated
  if printChainCmd.Parsed() {
    cli.printChain()
  }

}
