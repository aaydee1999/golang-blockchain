package main

import (
	"fmt"
	"strconv"

	"github.com/aaydee1999/IBC/blockchain"
)

func main() {

	blockChain := blockchain.InitBlockChain()

	blockChain.AddBlock("First Block After Genesis")
	blockChain.AddBlock("Second Block After Genesis")
	blockChain.AddBlock("Third Block After Genesis")

	for _, block := range blockChain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		proofOfWork := blockchain.NewProof(block)
		fmt.Printf("Proof of Work: %s\n", strconv.FormatBool(proofOfWork.Validate()))
		fmt.Println()
	}
}
