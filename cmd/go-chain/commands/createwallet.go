package commands

import (
	"fmt"

	"github.com/nnaakkaaii/go-chain/internal/model"
)

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := model.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
