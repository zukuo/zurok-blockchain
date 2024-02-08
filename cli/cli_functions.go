package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/zukuo/zurok-blockchain/blockchain"
	"github.com/zukuo/zurok-blockchain/network"
	"github.com/zukuo/zurok-blockchain/util"
	"github.com/zukuo/zurok-blockchain/wallet"
)

// Create a blockchain
func (cli *CLI) createBlockchain(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.CreateBlockchain(address, nodeID)
	defer bc.GetDB().Close()

	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}

func (cli *CLI) getBalance(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Invalid address")
	}
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{bc}
	defer bc.GetDB().Close()

	balance := 0
	pubKeyHash := wallet.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

// Print all blocks in the blockchain
func (cli *CLI) printChain(nodeID string) {
	bc := blockchain.NewBlockchain(nodeID)
	defer bc.GetDB().Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("========== Block - { %x } ==========\n", block.Hash)

		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		pow := blockchain.NewProofOfWork(block)
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
	if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is invalid")
	}
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is invalid")
	}

	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{bc}
	defer bc.GetDB().Close()

	wallets, err := wallet.NewWallets(nodeID)
	util.HandleError(err)
	wallet := wallets.GetWallet(from)
	tx := blockchain.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := blockchain.NewCoinbaseTX(from, "")
		txs := []*blockchain.Transaction{cbTx, tx}
		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		network.SendTx(network.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := wallet.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := wallet.NewWallets(nodeID)
	util.HandleError(err)
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)

	if len(minerAddress) > 0 {
		if wallet.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}

	network.StartServer(nodeID, minerAddress)
}
