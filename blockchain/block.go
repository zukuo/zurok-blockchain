package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Blocks Implementation in Blockchain
type Block struct {
	Timestamp    int64
	Transactions []*Transaction
	Hash         []byte
	PrevHash     []byte
	Nonce        int
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
	HandleError(err)

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	HandleError(err)

	return &block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func NewBlock(transactions []*Transaction, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, []byte{}, prevHash, 0}

	// Setup PoW
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
