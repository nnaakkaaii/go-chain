package commands

import (
	"fmt"
	"github.com/nnaakkaaii/go-chain/internal/model"
)

func (cli *CLI) addBlock(data string) {
	transactions := model.NewCoinbaseTx("Ivan", data)
	cli.bc.AddBlock([]*model.Transaction{transactions})
	fmt.Println("Success!")
}
