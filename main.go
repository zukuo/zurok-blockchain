package main

import (
	"zurok/blockchain"
)

// Test Blockchain
func main() {
	bc := blockchain.NewBlockchain()
	defer bc.GetDB().Close()

	cli := blockchain.CreateCLI(bc)
	cli.Run()
}
