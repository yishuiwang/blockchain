package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
	"time"
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
	return NewBlock(data, 0, []byte{})
}

func NewBlock(data string, height int64, preHash []byte) *Block {
	block := &Block{
		Height:    height,
		Hash:      []byte{},
		PreHash:   preHash,
		Data:      []byte(data),
		Timestamp: time.Now().Unix(),
		Nonce:     0,
	}

	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	fmt.Printf("\nNew Block: %x\n", block.Hash)

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

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(b); err != nil {
		log.Println(err)
	}
	return result.Bytes()
}

func DeserializeBlock(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&block); err != nil {
		log.Println(err)
	}
	return &block
}

func (b *Block) Print() {
	fmt.Printf("Height: %d\n", b.Height)
	fmt.Printf("PreHash: %x\n", b.PreHash)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Hash: %x\n", b.Hash)
	fmt.Printf("Timestamp: %s\n", time.Unix(b.Timestamp, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Println()
}
