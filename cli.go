package main

import (
	"fmt"
	"os"
)

// CLI struct
type CLI struct {
  bc *Blockchain
}

func (cli *CLI) printUsage() {
  fmt.Println("How to use the program: ")
  fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
  // If there are less than two arguments ->
  if len(os.Args) < 2 {
    // Print the correct commands that user should use 
    cli.printUsage()
    // Close the program
    os.Exit(1)
  }
}

