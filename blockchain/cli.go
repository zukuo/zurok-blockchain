package blockchain

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type CLI struct {}

// Create CLI
// func CreateCLI(bc *Blockchain) CLI {
//     return CLI{bc}
// }

// Validate and check inputs
func (cli *CLI) printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
    fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
    fmt.Println("  send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO")
    fmt.Println("  printchain - prints all the blocks in the chain")
}

func (cli *CLI) validateArgs() {
    if len(os.Args) < 2 {
        cli.printUsage()
        os.Exit(1)
    }
}

// Create a blockchain
func (cli *CLI) createBlockchain(address string) {
    bc := CreateBlockchain(address)
    bc.db.Close()
    fmt.Println("Done!")
}

func (cli *CLI) getBalance(address string) {
    bc := NewBlockchain(address)
    defer bc.db.Close()

    balance := 0
    UTXOs := bc.FindUTXO(address)

    for _, out := range UTXOs {
        balance += out.Value
    }

    fmt.Printf("Balance of '%s': %d\n", address, balance)
}

// Get the length of the blockchain
func (bc *Blockchain) getChainLength() int {
    bci := bc.Iterator()
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
    bc := NewBlockchain("")
    defer bc.db.Close()

    bci := bc.Iterator()
    blockNum := bc.getChainLength()

    for {
        block := bci.Next()

        blockNum--
        fmt.Printf("------ BLOCK %d -------\n", blockNum)

        fmt.Printf("Previous Hash: %x\n", block.PrevHash)
        fmt.Printf("Hash: %x\n", block.Hash)
        pow := NewProofOfWork(block)
        fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()

        if len(block.PrevHash) == 0 {
            break
        }
    }
}

func (cli *CLI) send(from, to string, amount int) {
    bc := NewBlockchain(from)
    defer bc.db.Close()

    tx := NewUTXOTransaction(from, to, amount, bc)
    bc.MineBlock([]*Transaction{tx})
    fmt.Println("Success!")
}

// Main CLI handler
func (cli *CLI) Run() {
	cli.validateArgs()

    // Define flags
    getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
    sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
    printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

    // Define params
    getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
    createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
    sendFrom := sendCmd.String("from", "", "Source wallet address")
    sendTo := sendCmd.String("to", "", "Destination wallet address")
    sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
    case "getbalance":
        err := getBalanceCmd.Parse(os.Args[2:])
        HandleError(err)
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
        HandleError(err)
	case "send":
		err := sendCmd.Parse(os.Args[2:])
        HandleError(err)
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
        HandleError(err)
	default:
		cli.printUsage()
		os.Exit(1)
	}

    if getBalanceCmd.Parsed() {
        if *getBalanceAddress == "" {
            getBalanceCmd.Usage()
            os.Exit(1)
        }
        cli.getBalance(*getBalanceAddress)
    }

    if createBlockchainCmd.Parsed() {
        if *createBlockchainAddress == "" {
            createBlockchainCmd.Usage()
            os.Exit(1)
        }
        cli.createBlockchain(*createBlockchainAddress)
    }

    if sendCmd.Parsed() {
        if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
            sendCmd.Usage()
            os.Exit(1)
        }
        cli.send(*sendFrom, *sendTo, *sendAmount)
    }

    if printChainCmd.Parsed() {
        cli.printChain()
    }
}
