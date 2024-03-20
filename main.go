package main

import (
	"blockchain/BLC"
	"fmt"
)

func main() {
	bc := BLC.NewBlockChain()
	bc.AddBlock("Send 1 BTC to Ivan", 2)
	bc.AddBlock("Send 2 more BTC to Ivan", 3)
	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", block.PreHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
