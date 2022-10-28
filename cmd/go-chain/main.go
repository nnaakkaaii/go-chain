package main

import (
	"github.com/nnaakkaaii/go-chain/cmd/go-chain/commands"
	"github.com/nnaakkaaii/go-chain/internal/model"
	"os"
)

func main() {
	bc := model.CreateBlockchain("Ivan")
	defer bc.DB.Close()

	cli := commands.NewCLI(bc)
	if err := cli.Run(); err != nil {
		os.Exit(1)
	}
}
