package commands

import (
	"fmt"
	"log"

	"github.com/nnaakkaaii/go-chain/internal/model"
)

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := model.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
