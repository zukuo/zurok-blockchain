package blockchain

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {}

// Validate and check inputs
func (cli *CLI) printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  createwallet - Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  listaddresses - Lists all addresses from the wallet file")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO")
}

func (cli *CLI) validateArgs() {
    if len(os.Args) < 2 {
        cli.printUsage()
        os.Exit(1)
    }
}

// Main CLI handler
func (cli *CLI) Run() {
	cli.validateArgs()

    // Define flags
    getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)
    sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
    printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

    // Set command outputs
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
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
        HandleError(err)
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
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

    if createWalletCmd.Parsed() {
        cli.createWallet()
    }

    if listAddressesCmd.Parsed() {
        cli.listAddresses()
    }

    if printChainCmd.Parsed() {
        cli.printChain()
    }
}
