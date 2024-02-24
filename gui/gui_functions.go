package gui

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/zukuo/zurok-blockchain/blockchain"
	"github.com/zukuo/zurok-blockchain/network"
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

// CreateWallet returns the wallet address of a newly generated wallet
func (a *App) CreateWallet(nodeID string) string {
	wallets, _ := wallet.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	return address
}

// Send Transactions - NEEDS UPDATING
func (a *App) SendTransaction(from, to string, amount int, nodeID string, mineNow bool) {
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

func (a *App) StartNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)

	if len(minerAddress) > 0 {
		if wallet.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}

	network.StartServerWithTime(nodeID, minerAddress, 3)
}

//type transactions struct {
//	Key         int    `json:"key"`
//	Transaction string `json:"transaction"`
//	Amount      int    `json:"balance"`
//}
//
//func (a *App) GetTransactionNames(nodeID string) []transactions {
//	bc := blockchain.NewBlockchain(nodeID)
//	defer bc.GetDB().Close()
//	bci := bc.Iterator()
//
//	for {
//		block := bci.Next()
//	}
//}

type blocks struct {
	Key       int    `json:"key"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevhash"`
	Height    int    `json:"height"`
	Timestamp string `json:"timestamp"`
	Nonce     int    `json:"nonce"`
	Pow       bool   `json:"pow"`
}

func (a *App) GetBlockInfos(nodeID string) []blocks {
	bc := blockchain.NewBlockchain(nodeID)
	defer bc.GetDB().Close()
	bci := bc.Iterator()
	var blocksArr []blocks

	i := 0
	for {
		block := bci.Next()
		pow := blockchain.NewProofOfWork(block)

		data := blocks{
			Key:       i + 1,
			Hash:      fmt.Sprintf("%x", block.Hash),
			PrevHash:  fmt.Sprintf("%x", block.PrevHash),
			Height:    block.Height,
			Timestamp: fmt.Sprintf("%s", time.Unix(block.Timestamp, 0)),
			Nonce:     block.Nonce,
			Pow:       pow.Validate(),
		}

		blocksArr = append(blocksArr, data)
		i++

		if len(block.PrevHash) == 0 {
			break
		}
	}

	return blocksArr
}

//func (a *App) GetAddressesWithBalances(nodeID string) []balances {
//	addresses := listAddresses(nodeID)
//	addrBal := make([]balances, len(addresses))
//
//	for i, addr := range addresses {
//		addrBal[i].Key = i + 1
//		addrBal[i].Address = addr
//		addrBal[i].Balance = getBalance(addr, nodeID)
//	}
//
//	return addrBal
//}

//func (a *App) RefreshNode(nodeID string) {
//	network.StartServer(nodeID, "")
//}

//func (a *App) StopNode(nodeID, minerAddress string) {
//
//}

// getHostname returns the hostname of the node
func getHostname() string {
	hostname, err := os.Hostname()
	util.HandleError(err)

	return hostname
}
