package blockchain

type BlockChain struct {
	Blocks []*Block
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (blockChain *BlockChain) AddBlock(data string) {
	prevBlock := blockChain.Blocks[len(blockChain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	blockChain.Blocks = append(blockChain.Blocks, new)

}
