package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (block *Block) DeriveHash() {
	info := bytes.Join([][]byte{block.Data, block.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	block.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (blockChain *BlockChain) AddBlock(data string) {
	prevBlock := blockChain.blocks[len(blockChain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	blockChain.blocks = append(blockChain.blocks, new)

}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
