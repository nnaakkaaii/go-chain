package main

import (
	"flag"
	"github.com/nnaakkaaii/go-chain/internal/model"
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

	bc := model.NewBlockchain()

	for _, content := range blockContents {
		bc.AddBlock(content)
	}
}
