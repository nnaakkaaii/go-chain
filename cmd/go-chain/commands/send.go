package commands

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/nnaakkaaii/go-chain/internal/server"
	"log"

	"github.com/nnaakkaaii/go-chain/internal/model"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !model.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !model.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := model.NewBlockchain(nodeID)
	UTXOSet := model.UTXOSet{bc}
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(bc.DB)

	wallets, err := model.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := model.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := model.NewCoinbaseTX(from, "")
		txs := []*model.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		server.SendTx(server.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}
