package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/zukuo/zurok-blockchain/util"
)

// Blocks Implementation in Blockchain
type Block struct {
	Timestamp    int64
	Transactions []*Transaction
	Hash         []byte
	PrevHash     []byte
	Nonce        int
	Height       int
}

func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	merkleTree := NewMerkleTree(transactions)

	return merkleTree.RootNode.Data
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	util.HandleError(err)

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	util.HandleError(err)

	return &block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0)
}

func NewBlock(transactions []*Transaction, prevHash []byte, height int) *Block {
	block := &Block{time.Now().Unix(), transactions, []byte{}, prevHash, 0, height}

	// Setup PoW
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
