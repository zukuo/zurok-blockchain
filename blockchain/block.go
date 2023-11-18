package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Blocks Implementation
type Block struct {
	Timestamp int64
	Data      []byte
	Hash      []byte
	PrevHash  []byte
	Nonce     int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), []byte{}, prevHash, 0}
	// block.SetHash()

	// New PoW
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
