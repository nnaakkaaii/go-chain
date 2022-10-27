package model

// Blockchain is a database with a structure, ordered and back-linked.
type Blockchain struct {
	blocks []*Block
	// the map that keeps `hash -> block` pairs is necessary in a future.
}

// AddBlock will be used to add a new block to existing ones.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock creates a GenesisBlock, the first block in a chain.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain creates a blockchain with the genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
