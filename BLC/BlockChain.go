package BLC

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	genesisBlock := CreateGenesisBlock("Genesis Block")
	return &BlockChain{[]*Block{genesisBlock}}
}

func (bc *BlockChain) AddBlock(data string, height int64) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash, height)
	bc.Blocks = append(bc.Blocks, newBlock)
}
