package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/nnaakkaaii/go-chain/internal/model"
	"os"
)

const (
	addblock   = "addblock"
	printchain = "printchain"
)

var errUnknownCmd = errors.New("unknown command")

type CLI struct {
	bc *model.Blockchain
}

func NewCLI(bc *model.Blockchain) *CLI {
	return &CLI{bc: bc}
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() error {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet(addblock, flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(printchain, flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case addblock:
		if err := addBlockCmd.Parse(os.Args[2:]); err != nil {
			return err
		}
		if *addBlockData == "" {
			addBlockCmd.Usage()
			return errUnknownCmd
		}
		cli.addBlock(*addBlockData)
	case printchain:
		if err := printChainCmd.Parse(os.Args[2:]); err != nil {
			return err
		}
		cli.printChain()
	default:
		cli.printUsage()
		return errUnknownCmd
	}

	return nil
}
