package BLC

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {
	BlockChain *BlockChain
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Parse error:", err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Parse error:", err)
		}
	default:
		cli.PrintUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) PrintUsage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("  addblock -data BLOCK_DATA - add a block to the blockchain\n")
	fmt.Printf("  printchain - print all the blocks of the blockchain\n")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.PrintUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.BlockChain.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	cli.BlockChain.PrintChain()
}
