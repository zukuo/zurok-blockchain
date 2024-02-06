package blockchain

import (
	"fmt"
	"log"
	"strconv"
)

// Create a blockchain
func (cli *CLI) createBlockchain(address, nodeID string) {
    if !ValidateAddress(address) {
        log.Panic("ERROR: Address is not valid")
    }
    bc := CreateBlockchain(address, nodeID)
    defer bc.db.Close()

    UTXOSet := UTXOSet{bc}
    UTXOSet.Reindex()

    fmt.Println("Done!")
}

func (cli *CLI) getBalance(address, nodeID string) {
    if !ValidateAddress(address) {
        log.Panic("ERROR: Invalid address")
    }
    bc := NewBlockchain(nodeID)
    UTXOSet := UTXOSet{bc}
    defer bc.db.Close()

    balance := 0
    pubKeyHash := Base58Decode([]byte(address))
    pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
    UTXOs := UTXOSet.FindUTXO(pubKeyHash)

    for _, out := range UTXOs {
        balance += out.Value
    }

    fmt.Printf("Balance of '%s': %d\n", address, balance)
}

// Print all blocks in the blockchain
func (cli *CLI) printChain(nodeID string) {
    bc := NewBlockchain(nodeID)
    defer bc.db.Close()

    bci := bc.Iterator()

    for {
        block := bci.Next()

        fmt.Printf("========== Block - { %x } ==========\n", block.Hash)

        fmt.Printf("Height: %d\n", block.Height)
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

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
    if !ValidateAddress(from) {
        log.Panic("ERROR: Sender address is invalid")
    }
    if !ValidateAddress(to) {
        log.Panic("ERROR: Recipient address is invalid")
    }

    bc := NewBlockchain(nodeID)
    UTXOSet := UTXOSet{bc}
    defer bc.db.Close()

    wallets, err := NewWallets(nodeID)
    HandleError(err)
    wallet := wallets.GetWallet(from)
    tx := NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

    if mineNow {
        cbTx := NewCoinbaseTX(from, "")
        txs := []*Transaction{cbTx, tx}
        newBlock := bc.MineBlock(txs)
        UTXOSet.Update(newBlock)
    } else {
        sendTx(knownNodes[0], tx)
    }

    fmt.Println("Success!")
}

func (cli *CLI) createWallet(nodeID string) {
    wallets, _ := NewWallets(nodeID)
    address := wallets.CreateWallet()
    wallets.SaveToFile(nodeID)

    fmt.Printf("Your new address: %s\n", address)
}

func (cli *CLI) listAddresses(nodeID string) {
    wallets, err := NewWallets(nodeID)
    HandleError(err)
    addresses := wallets.GetAddresses()

    for _, address := range addresses {
        fmt.Println(address)
    }
}

func (cli *CLI) reindexUTXO(nodeID string) {
    bc := NewBlockchain(nodeID)
    UTXOSet := UTXOSet{bc}
    UTXOSet.Reindex()

    count := UTXOSet.CountTransactions()
    fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}

func (cli *CLI) startNode(nodeID, minerAddress string) {
    fmt.Printf("Starting node %s\n", nodeID)

    if len(minerAddress) > 0 {
        if ValidateAddress(minerAddress) {
            fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
        } else {
            log.Panic("Wrong miner address!")
        }
    }

    StartServer(nodeID, minerAddress)
}
