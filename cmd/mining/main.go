package main

import (
	"flag"
	"fmt"
	"github.com/nnaakkaaii/go-chain/internal/model"
	"strconv"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "array flags"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var blockContents arrayFlags
	flag.Var(&blockContents, "c", "specify multiple contents")
	flag.Parse()

	bc := model.CreateBlockchain("")

	for _, content := range blockContents {
		transactions := model.NewCoinbaseTx("Michael", content)
		bc.AddBlock([]*model.Transaction{transactions})
	}

	i := bc.Iterate()
	for {
		b := i.Next()
		fmt.Printf("Prev. hash: %x\n", b.PrevBlockHash)
		fmt.Printf("Data: %x\n", b.HashTransactions())
		fmt.Printf("Hash: %x\n", b.Hash)
		pow := model.NewProofOfWork(b)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(b.PrevBlockHash) == 0 {
			break
		}
	}
}
