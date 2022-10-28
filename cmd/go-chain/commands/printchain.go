package commands

import (
	"fmt"
	"github.com/nnaakkaaii/go-chain/internal/model"
	"strconv"
)

func (cli *CLI) printChain() {
	bci := cli.bc.Iterate()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %x\n", block.HashTransactions())
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := model.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
