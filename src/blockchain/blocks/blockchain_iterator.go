package blocks

import "github.com/boltdb/bolt"

//检查区块链
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}


//返回区块链中的下一个区块
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {

	}

	i.currentHash = block.PrevBlockHash

	return block
}
