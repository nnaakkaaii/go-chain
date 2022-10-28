package commands

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/nnaakkaaii/go-chain/pkg/base58"
	"log"

	"github.com/nnaakkaaii/go-chain/internal/model"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !model.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := model.NewBlockchain(nodeID)
	UTXOSet := model.UTXOSet{bc}
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(bc.DB)

	balance := 0
	pubKeyHash := base58.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
