package blocks

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
	"time"
	"encoding/json"
	"fmt"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Id            []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

//创建区块
func NewBlock(data, id string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0, []byte(id)}
	//block.SetHash()

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "1", []byte{})
}

func Verify(id string) (flag string) {
	bc := GetBlockchain()
	defer bc.DBClose()
	bci := bc.Iterator()
	for {
		block := bci.Next()
		blockdata,err := bc.GetBlock(block.Hash)
		if err != nil {
			return err.Error()
		}
		//json str 转struct
		var dat map[string]interface{}
		e := json.Unmarshal(blockdata.Data, &dat)
		if  e == nil {
			fmt.Println(dat)
			fmt.Println(dat["Id_"])
			fmt.Println(id)
			if dat["Id_"] == id {
				return "验证通过"
			}

		}
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return "验证失败：没有找到相关block"
}

//序列化数据结构
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	encoder.Encode(b)

	return result.Bytes()
}

//反序列化
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	decoder.Decode(&block)

	return &block
}
