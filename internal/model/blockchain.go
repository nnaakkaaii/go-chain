package model

import (
	"fmt"
	"github.com/boltdb/bolt"
)

const (
	dbFile              = "blockchain_%s.db"
	blocksBucket        = "blocks"
	genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
)

// Blockchain is a database with a structure, ordered and back-linked.
type Blockchain struct {
	tip []byte
	DB  *bolt.DB
}

// AddBlock will be used to add a new block to existing ones.
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	var lastHash []byte

	if err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	}); err != nil {
		panic(err)
	}

	newBlock := NewBlock(transactions, lastHash)

	if err := bc.DB.Update(func(tx *bolt.Tx) error {
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
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func CreateBlockchain(address string) *Blockchain {
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

		cbtx := NewCoinbaseTx(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)

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

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) Iterate() *BlockchainIterator {
	return &BlockchainIterator{
		currentHash: bc.tip,
		db:          bc.DB,
	}
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	if err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	}); err != nil {
		panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}
