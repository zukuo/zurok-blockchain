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
	//network.StartServer(nodeID, minerAddress)
}

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

	for {
		block := bci.Next()
		pow := blockchain.NewProofOfWork(block)

		data := blocks{
			Key:       block.Height + 1,
			Hash:      fmt.Sprintf("%x", block.Hash),
			PrevHash:  fmt.Sprintf("%x", block.PrevHash),
			Height:    block.Height,
			Timestamp: fmt.Sprintf("%s", time.Unix(block.Timestamp, 0)),
			Nonce:     block.Nonce,
			Pow:       pow.Validate(),
		}

		blocksArr = append(blocksArr, data)

		if len(block.PrevHash) == 0 {
			break
		}
	}

	sort.Slice(blocksArr, func(i, j int) bool {
		return blocksArr[i].Key > blocksArr[j].Key
	})

	return blocksArr
}

type transactions struct {
	Key         int      `json:"key"`
	Transaction string   `json:"transaction"`
	Amount      int      `json:"amount"`
	Block       string   `json:"block"`
	Height      int      `json:"height"`
	Inputs      []*txin  `json:"inputs"`
	Outputs     []*txout `json:"outputs"`
}

type txin struct {
	Txid      string `json:"txid"`
	Vout      int    `json:"vout"`
	Signature string `json:"signature"`
	PubKey    string `json:"pubkey"`
}

type txout struct {
	Value      int    `json:"value"`
	PubKeyHash string `json:"pubkeyhash"`
}

func (a *App) GetTransactions(nodeID string) []transactions {
	bc := blockchain.NewBlockchain(nodeID)
	defer bc.GetDB().Close()
	bci := bc.Iterator()
	var txs []transactions

	i := 1
	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			data := transactions{
				Key:         i,
				Transaction: fmt.Sprintf("%x", tx.ID),
				Amount:      tx.Vout[0].Value,
				Block:       fmt.Sprintf("%x", block.Hash),
				Height:      block.Height,
				Inputs:      make([]*txin, 0),
				Outputs:     make([]*txout, 0),
			}

			for _, input := range tx.Vin {
				in := txin{
					Txid:      util.BytesToString(input.Txid),
					Vout:      input.Vout,
					Signature: util.BytesToString(input.Signature),
					PubKey:    util.BytesToString(input.PubKey),
				}
				data.Inputs = append(data.Inputs, &in)
			}

			for _, output := range tx.Vout {
				out := txout{
					Value:      output.Value,
					PubKeyHash: util.BytesToString(output.PubKeyHash),
				}
				data.Outputs = append(data.Outputs, &out)
			}

			txs = append(txs, data)
			i++
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}

	sort.Slice(txs, func(i, j int) bool {
		return txs[i].Key < txs[j].Key
	})

	return txs
}

// getHostname returns the hostname of the node
func getHostname() string {
	hostname, err := os.Hostname()
	util.HandleError(err)

	return hostname
}
