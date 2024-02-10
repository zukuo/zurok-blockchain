package gui

import (
	"log"
	"os"
	"sort"

	"github.com/zukuo/zurok-blockchain/blockchain"
	"github.com/zukuo/zurok-blockchain/util"
	"github.com/zukuo/zurok-blockchain/wallet"
)

type balances struct {
	Key     int    `json:"key"`
	Address string `json:"address"`
	Balance int    `json:"balance"`
}

func (a *App) GetAddressesWithBalances(nodeID string) []balances {
	addresses := listAddresses(nodeID)
	addrBal := make([]balances, len(addresses))

	for i, addr := range addresses {
		addrBal[i].Key = i + 1
		addrBal[i].Address = addr
		addrBal[i].Balance = getBalance(addr, nodeID)
	}

	return addrBal
}

// GetAddresses is used for getting all addresses
func (a *App) GetAddresses(nodeID string) []string {
	wallets, err := wallet.NewWallets(nodeID)
	util.HandleError(err)
	addresses := wallets.GetAddresses()
	sort.Strings(addresses)

	return addresses
}

// listAddresses returns a list of strings of all wallet addresses in the database at a given node
func listAddresses(nodeID string) []string {
	wallets, err := wallet.NewWallets(nodeID)
	util.HandleError(err)
	addresses := wallets.GetAddresses()
	sort.Strings(addresses)

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
