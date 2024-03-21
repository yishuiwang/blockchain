package BLC

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockChain struct {
	Tip []byte // 最新区块的Hash
	DB  *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			// 将创世区块存入数据库
			genesis := CreateGenesisBlock("Genesis Block")
			bucket, _ := tx.CreateBucket([]byte(blocksBucket))
			key := genesis.Hash
			value := genesis.Serialize()
			_ = bucket.Put(key, value)
			_ = bucket.Put([]byte("l"), key)
			tip = key
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	return &BlockChain{tip, db}
}

func (bc *BlockChain) AddBlock(data string) {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		if bucket != nil {
			// 获取最新区块
			lastHash := bucket.Get([]byte("l"))
			lastBlockBytes := bucket.Get(lastHash)
			lastBlock := DeserializeBlock(lastBlockBytes)

			// 创建新区块
			newBlock := NewBlock(data, lastBlock.Height+1, lastBlock.Hash)
			key := newBlock.Hash
			value := newBlock.Serialize()
			_ = bucket.Put(key, value)
			_ = bucket.Put([]byte("l"), key)
			bc.Tip = key
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (bc *BlockChain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.DB}

	return bci
}

func (bc *BlockChain) PrintChain() {
	bci := bc.Iterator()
	for {
		block := bci.Next()
		block.Print()
		if len(block.PreHash) == 0 {
			break
		}
	}
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PreHash

	return block
}
