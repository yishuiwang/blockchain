package blockchain

import "strconv"

type Block struct {
	Height    int    // 区块高度
	Hash      []byte // 区块哈希
	PreHash   []byte // 前一个区块哈希
	Data      []byte // 区块数据
	Timestamp int64  // 时间戳
}

func NewBlock(data string, preHash []byte, height int) *Block {
	block := &Block{
		Height:    height,
		Hash:      []byte{},
		PreHash:   preHash,
		Data:      []byte(data),
		Timestamp: 0,
	}
	return block
}

func (b *Block) SetHash() {
	timeString := strconv.FormatInt(b.Timestamp, 2)
	
}
