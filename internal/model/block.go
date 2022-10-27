package model

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block contains its headers and some transactions.
// Timestamp means the timestamp when a block generated.
// Data means the actual data that includes some information.
// PrevBlockHash refers to the hash of a previous block.
// Hash stores the root hash of this block.
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// SetHash concatenate block fields and calculate SHA-256 hash.
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock creates one Block struct
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
	}
	block.SetHash()
	return block
}
