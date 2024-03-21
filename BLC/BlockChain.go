package BLC

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockChain struct {
	Tip []byte // 最新区块的Hash
	DB  *bolt.DB
}

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// OpenOrCreateBlockChain 打开或创建区块链
func OpenOrCreateBlockChain() *BlockChain {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	// 从数据库中获取最新区块的 Hash
	var tip []byte
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			// 如果数据库中不存在区块链数据，则创建创世区块
			genesis := CreateGenesisBlock("Genesis Block")
			bucket, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}
			_ = bucket.Put(genesis.Hash, genesis.Serialize())
			_ = bucket.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			// 数据库中存在区块链数据，获取最新区块的 Hash
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

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
