package main

import (
	_ "crypto/sha256"
	_ "fmt"
	"github.com/boltdb/bolt"
	"log"
)


//引入區塊
type BlockChain struct{
	//定一個區塊數組
	//blocks []*Block
	db *bolt.DB
	tail []byte //存儲最後一個區塊的哈希
}
const BlockChainDB="blockChain.db"
const BlockBucket="blockBucket"

//定義一個區塊鏈
func NewBlockChain() *BlockChain  {
	//創建一個創世區塊，並作爲第一個區塊添加到區塊鏈中
	/*genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks:[]*Block{genesisBlock},
	}*/
	//最後一個區塊的hash，從數據庫中讀出來的
	var lastHash []byte
	//打開數據庫
	db,err:=bolt.Open("BlockChainDB",0600,nil)
	defer db.Close()

	if err!=nil{
		log.Panic("打開數據庫失敗")
	}

	//將要操作的數據庫改寫
	db.Update(func(tx *bolt.Tx) error {
		//找到抽屜bucket(如果沒有，就創建)
		bucket:=tx.Bucket([]byte("blockBucket"))
		if bucket==nil{
			//沒有抽屜，我們需要創建
			bucket,err=tx.CreateBucket([]byte("blockBucket"))
			if err !=nil{
				log.Panic("創建bucket(blockBucket)失敗")
			}
			//創建一個創世區塊，並作爲第一個區塊添加到區塊鏈中
			genesisBlock := GenesisBlock()

			//寫數據
			//hash作爲key，block字節流作爲value
			bucket.Put(genesisBlock.Hash,genesisBlock.toByte())
			bucket.Put([]byte("LastHashKey"),genesisBlock.Hash)
			lastHash=genesisBlock.Hash
		}else {
			//讀數據
			lastHash=bucket.Get([]byte("LastHashKey"))
		}

		return nil
	})
	return &BlockChain{db,lastHash}

}

//創世快
func GenesisBlock() *Block {
	return NewBlock("這是一個創世區塊",[]byte{})
}

//添加區塊
func (bc *BlockChain)AddBlock(data string)  {
	//如何獲取前區塊哈希
	
	/*//獲取最後一個區塊
	lastBlock:=bc.blocks[len(bc.blocks)-1]
	prvHash := lastBlock.Hash

	//創建新的區塊
	block:= NewBlock(data,prvHash)

	//添加到區塊鏈數組中
	bc.blocks=append(bc.blocks,block)*/
}
