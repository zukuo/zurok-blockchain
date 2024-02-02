package blockchain

import (
	"fmt"
	"log"
	"strconv"
)

// Create a blockchain
func (cli *CLI) createBlockchain(address string) {
    if !ValidateAddress(address) {
        log.Panic("ERROR: Address is not valid")
    }
    bc := CreateBlockchain(address)
    bc.db.Close()
    fmt.Println("Done!")
}

func (cli *CLI) getBalance(address string) {
    if !ValidateAddress(address) {
        log.Panic("ERROR: Invalid address")
    }
    bc := NewBlockchain(address)
    defer bc.db.Close()

    balance := 0
    pubKeyHash := Base58Decode([]byte(address))
    pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
    UTXOs := bc.FindUTXO(pubKeyHash)

    for _, out := range UTXOs {
        balance += out.Value
    }

    fmt.Printf("Balance of '%s': %d\n", address, balance)
}

// Print all blocks in the blockchain
func (cli *CLI) printChain() {
    bc := NewBlockchain("")
    defer bc.db.Close()

    bci := bc.Iterator()

    for {
        block := bci.Next()

        fmt.Printf("========== Block - { %x } ==========\n", block.Hash)

        fmt.Printf("Previous Hash: %x\n", block.PrevHash)
        pow := NewProofOfWork(block)
        fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()
        for _, tx := range block.Transactions {
            fmt.Println(tx)
        }

        if len(block.PrevHash) == 0 {
            break
        }
    }
}

func (cli *CLI) send(from, to string, amount int) {
    if !ValidateAddress(from) {
        log.Panic("ERROR: Sender address is invalid")
    }
    if !ValidateAddress(to) {
        log.Panic("ERROR: Recipient address is invalid")
    }

    bc := NewBlockchain(from)
    defer bc.db.Close()

    tx := NewUTXOTransaction(from, to, amount, bc)
    bc.MineBlock([]*Transaction{tx})
    fmt.Println("Success!")
}

func (cli *CLI) createWallet() {
    wallets, _ := NewWallets()
    address := wallets.CreateWallet()
    wallets.SaveToFile()

    fmt.Printf("Your new address: %s\n", address)
}

func (cli *CLI) listAddresses() {
    wallets, err := NewWallets()
    HandleError(err)
    addresses := wallets.GetAddresses()

    for _, address := range addresses {
        fmt.Println(address)
    }
}
