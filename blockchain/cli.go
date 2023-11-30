package blockchain

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	bc *Blockchain
}

// Create CLI
func CreateCLI(bc *Blockchain) CLI {
    return CLI{bc}
}

// Validate and check inputs
func (cli *CLI) printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  addblock -data BLOCK_DATA - adds a block to the blockchain")
    fmt.Println("  printchain - prints all the blocks in the chain")
}

func (cli *CLI) validateArgs() {
    if len(os.Args) < 2 {
        cli.printUsage()
        os.Exit(1)
    }
}

// Add block
func (cli *CLI) addBlock(data string) {
    cli.bc.AddBlock(data)
    fmt.Println("Success!")
}

// Get the length of the blockchain
func (cli *CLI) getChainLength() int {
    bci := cli.bc.Iterator()
    blockNum := 0
    
    for {
        block := bci.Next()
        blockNum++
        if len(block.PrevHash) == 0 {
            break
        }
    }

    return blockNum
}

// Print all blocks in the blockchain
func (cli *CLI) printChain() {
    bci := cli.bc.Iterator()

    blockNum := cli.getChainLength()

    for {
        block := bci.Next()

        blockNum--
        fmt.Printf("------ BLOCK %d -------\n", blockNum)

        fmt.Printf("Previous Hash: %x\n", block.PrevHash)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
        pow := NewProofOfWork(block)
        fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()

        if len(block.PrevHash) == 0 {
            break
        }
    }
}

// Main CLI handler
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
        HandleError(err)
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
        HandleError(err)
	default:
		cli.printUsage()
		os.Exit(1)
	}

    if addBlockCmd.Parsed() {
        if *addBlockData == "" {
            addBlockCmd.Usage()
            os.Exit(1)
        }
        cli.addBlock(*addBlockData)
    }

    if printChainCmd.Parsed() {
        cli.printChain()
    }
}
