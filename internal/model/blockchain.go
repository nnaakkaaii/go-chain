package model

import (
	"fmt"
	"github.com/boltdb/bolt"
)

const (
	dbFile       = "blockchain_%s.db"
	blocksBucket = "blocks"
)

// Blockchain is a database with a structure, ordered and back-linked.
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

// AddBlock will be used to add a new block to existing ones.
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	if err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	}); err != nil {
		panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	if err := bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if err := b.Put(newBlock.Hash, newBlock.Serialize()); err != nil {
			return err
		}
		if err := b.Put([]byte("l"), newBlock.Hash); err != nil {
			return err
		}

		bc.tip = newBlock.Hash

		return nil
	}); err != nil {
		panic(err)
	}
}

// NewGenesisBlock creates a GenesisBlock, the first block in a chain.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain creates a blockchain with the genesis block
func NewBlockchain() *Blockchain {
	var tip []byte

	db, err := bolt.Open(fmt.Sprintf(dbFile, "1"), 0600, nil)
	if err != nil {
		panic(err)
	}

	if err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b != nil {
			tip = b.Get([]byte("l"))
			return nil
		}
		genesis := NewGenesisBlock()

		if b, err = tx.CreateBucket([]byte(blocksBucket)); err != nil {
			return err
		}
		if err = b.Put(genesis.Hash, genesis.Serialize()); err != nil {
			return err
		}
		if err = b.Put([]byte("l"), genesis.Hash); err != nil {
			return err
		}
		tip = genesis.Hash

		return nil
	}); err != nil {
		panic(err)
	}

	return &Blockchain{tip, db}
}
