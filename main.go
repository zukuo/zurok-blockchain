package main

import (
	"fmt"
	"strconv"
	"zurok/blockchain"
)

// Test Blockchain
func main() {
	bc := blockchain.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Eesa")
	bc.AddBlock("Send 2 more BTC to Eesa")

	for _, block := range bc.GetBlocks() {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
