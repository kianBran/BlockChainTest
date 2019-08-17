package main

import (


	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	//用户获取区块链数据
	db *bolt.DB

	//游标，用于不断索引
	currentHashPoint []byte

}

//定义迭代器

func (bc *BlockChain ) NewIterator()  *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,
		//最初指向区块链的最后一个区块，随着next的调用，不断变化
		bc.tail,
	}

}


//迭代器是属于区块链的
//Next方式是属于迭代器的
//next方法：1、返回当前区块; 2、指针前移
func (it *BlockChainIterator) Next() *Block  {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlockBucket))
		if bucket==nil{
			log.Panic("迭代器遍历时bucket不应该为空，请检查")
		}

		blockTmp := bucket.Get(it.currentHashPoint)

		//解码动作
		block = Deserialize(blockTmp)
		//游标哈希左移
		it.currentHashPoint=block.PrvHash

		return nil
	})
	return &block
}
