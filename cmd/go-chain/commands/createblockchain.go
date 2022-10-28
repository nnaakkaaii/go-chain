package commands

import (
	"fmt"
	"github.com/boltdb/bolt"

	"github.com/nnaakkaaii/go-chain/internal/model"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !model.ValidateAddress(address) {
		panic("ERROR: Address is not valid")
	}
	bc := model.CreateBlockchain(address, nodeID)
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(bc.DB)

	UTXOSet := model.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
