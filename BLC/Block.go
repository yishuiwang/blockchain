package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	Height    int64  // 区块高度
	Hash      []byte // 区块哈希
	PreHash   []byte // 前一个区块哈希
	Data      []byte // 区块数据
	Timestamp int64  // 时间戳
	Nonce     int64  // 随机数
}

// CreateGenesisBlock 创建创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, []byte{}, 1)
}

func NewBlock(data string, preHash []byte, height int64) *Block {
	block := &Block{
		Height:    height,
		Hash:      []byte{},
		PreHash:   preHash,
		Data:      []byte(data),
		Timestamp: 0,
		Nonce:     0,
	}

	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

func (b *Block) SetHash() {
	// 1. Height -> byte
	heightBytes := IntToHex(b.Height)
	// 2. Timestamp -> byte
	timeString := strconv.FormatInt(b.Timestamp, 2)
	// 3. join all the bytes
	blockBytes := bytes.Join([][]byte{heightBytes, b.PreHash, b.Data, []byte(timeString)}, []byte{})
	// 4. generate hash
	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}
