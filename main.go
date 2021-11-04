package main

import (
	"fmt"

	"github.com/aaydee1999/IBC/blockchain"
)

func main() {

	blockChain := blockchain.InitBlockChain()

	blockChain.AddBlock("First Block After Genesis")
	blockChain.AddBlock("Second Block After Genesis")
	blockChain.AddBlock("Third Block After Genesis")

	for _, block := range blockChain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("\n")
	}
}
