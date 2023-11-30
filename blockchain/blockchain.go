package blockchain

import (
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Blockchain Type
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

func (bc *Blockchain) GetDB() *bolt.DB {
	return bc.db
}

func (bc *Blockchain) GetBlocks() [][]byte {
	var blocks [][]byte
	bci := bc.Iterator()
	
	for {
		block := bci.Next()
		blocks = append(blocks, block.Hash)
		if len(block.PrevHash) == 0 {
			break
		}
	}

	return blocks
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	// Get last block hash from DB (read-only)
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	HandleError(err)

	newBlock := NewBlock(data, lastHash)

	// After mining new block, save its serialized vals into DB
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("l"), newBlock.Hash)
		HandleError(err)
		bc.tip = newBlock.Hash
		return nil
	})
}

func NewBlockchain() *Blockchain {

	// Open DB
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	err = db.Update(func(tx *bolt.Tx) error {
		// Open read-write transaction
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			HandleError(err)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	HandleError(err)

	bc := Blockchain{tip, db}

	return &bc
}
