package blocks

import (
	"github.com/boltdb/bolt"
)
const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
	Blocks []*Block
	tip []byte
	db  *bolt.DB
}

//添加区块链 数组形式
//func (bc *Blockchain) AddBlock(data string) {
//	prevBlock := bc.Blocks[len(bc.Blocks)-1]
//	newBlock := NewBlock(data, prevBlock.Hash)
//	bc.Blocks = append(bc.Blocks, newBlock)
//}

//添加区块链 数据库形式
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	newBlock := NewBlock(data, lastHash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		//存储序列化的代表数据到数据库里
		err = b.Put(newBlock.Hash, newBlock.Serialize())
		//更新 l 键的hash
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
}
//创建初始区块链
func NewBlockchain() *Blockchain {
	//return &Blockchain{[]*Block{NewGenesisBlock()}}
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err = tx.CreateBucket([]byte(blocksBucket))
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip: tip, db: db}

	return &bc
}

//迭代器初始化
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}