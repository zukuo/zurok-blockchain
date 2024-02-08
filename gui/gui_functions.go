package gui

import (
	"log"
	"os"

	"github.com/zukuo/zurok-blockchain/blockchain"
	"github.com/zukuo/zurok-blockchain/util"
	"github.com/zukuo/zurok-blockchain/wallet"
)

// ListAddresses returns a list of strings of all wallet addresses in the database at a given node
func (a *App) ListAddresses(nodeID string) []string {
	wallets, err := wallet.NewWallets(nodeID)
	util.HandleError(err)
	addresses := wallets.GetAddresses()

	return addresses
}

// getBalance returns an int of a given wallet address for a given node
func getBalance(address, nodeID string) int {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Invalid address")
	}
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	defer bc.GetDB().Close()

	balance := 0
	pubKeyHash := wallet.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}

// createWallet returns the wallet address of a newly generated wallet
func createWallet(nodeID string) string {
	wallets, _ := wallet.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	return address
}

// getHostname returns the hostname of the node
func getHostname() string {
	hostname, err := os.Hostname()
	util.HandleError(err)

	return hostname
}
