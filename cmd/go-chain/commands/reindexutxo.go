package commands

import (
	"fmt"
	"github.com/nnaakkaaii/go-chain/internal/model"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := model.NewBlockchain(nodeID)
	UTXOSet := model.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
