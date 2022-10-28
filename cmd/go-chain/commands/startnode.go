package commands

import (
	"fmt"
	"log"

	"github.com/nnaakkaaii/go-chain/internal/model"
	"github.com/nnaakkaaii/go-chain/internal/server"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if model.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	server.StartServer(nodeID, minerAddress)
}
