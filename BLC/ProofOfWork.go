package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	Block  *Block   // 当前要验证的区块
	Target *big.Int // 大数据存储
}

// 目标难度值
const TargetBits = 16

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TargetBits))

	pow := &ProofOfWork{block, target}

	return pow
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
	nonce := 0
	var hashInt big.Int
	var hash [32]byte

	for {
		data := pow.prepareData(int64(nonce))
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.Target) == -1 {
			break
		}
		nonce++
	}

	return hash[:], int64(nonce)
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PreHash,
		pow.Block.Data,
		IntToHex(pow.Block.Timestamp),
		IntToHex(TargetBits),
		IntToHex(nonce),
		IntToHex(pow.Block.Height),
	}, []byte{})

	return data
}
