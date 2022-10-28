package commands

import (
	"flag"
	"fmt"
	"os"
)

const (
	getbalance       = "getbalance"
	createblockchain = "createblockchain"
	createwallet     = "createwallet"
	listaddresses    = "listaddresses"
	printchain       = "printchain"
	reindexutxo      = "reindexutxo"
	send             = "send"
	startnode        = "startnode"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  createwallet - Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  listaddresses - Lists all addresses from the wallet file")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  reindexutxo - Rebuilds the UTXO set")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT -mine - Send AMOUNT of coins from FROM address to TO. Mine on the same node, when -mine is set.")
	fmt.Println("  startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Println("NODE_ID env. var is not set")
		os.Exit(1)
	}

	getBalanceCmd := flag.NewFlagSet(getbalance, flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet(createblockchain, flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet(createwallet, flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet(listaddresses, flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(printchain, flag.ExitOnError)
	reindexUTXOCmd := flag.NewFlagSet(reindexutxo, flag.ExitOnError)
	sendCmd := flag.NewFlagSet(send, flag.ExitOnError)
	startNodeCmd := flag.NewFlagSet(startnode, flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")
	sendMine := sendCmd.Bool("mine", false, "Mine immediately on the same node")
	startNodeMiner := startNodeCmd.String("miner", "", "Enable mining mode and send reward to ADDRESS")

	switch os.Args[1] {
	case getbalance:
		if err := getBalanceCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress, nodeID)
	case createblockchain:
		if err := createBlockchainCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress, nodeID)
	case createwallet:
		if err := createWalletCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		cli.createWallet(nodeID)
	case listaddresses:
		if err := listAddressesCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		cli.listAddresses(nodeID)
	case printchain:
		if err := printChainCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		cli.printChain(nodeID)
	case reindexutxo:
		if err := reindexUTXOCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		cli.reindexUTXO(nodeID)
	case send:
		if err := sendCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFrom, *sendTo, *sendAmount, nodeID, *sendMine)
	case startnode:
		if err := startNodeCmd.Parse(os.Args[2:]); err != nil {
			panic(err)
		}
		cli.startNode(nodeID, *startNodeMiner)
	default:
		cli.printUsage()
		os.Exit(1)
	}
}
