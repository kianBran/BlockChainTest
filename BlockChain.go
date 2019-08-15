package main

import (
	_ "crypto/sha256"
)


//引入區塊
type BlockChain struct{
	//定一個區塊數組
	blocks []*Block
}

//定義一個區塊鏈
func NewBlockChain() *BlockChain  {
	//創建一個創世區塊，並作爲第一個區塊添加到區塊鏈中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks:[]*Block{genesisBlock},
	}
}

//創世快
func GenesisBlock() *Block {
	return NewBlock("這是一個創世區塊",[]byte{})
}

//添加區塊
func (bc *BlockChain)AddBlock(data string)  {
	//如何獲取前區塊哈希
	
	//獲取最後一個區塊
	lastBlock:=bc.blocks[len(bc.blocks)-1]
	prvHash := lastBlock.PrvHash

	//創建新的區塊
	block:= NewBlock(data,prvHash)

	//添加到區塊鏈數組中
	bc.blocks=append(bc.blocks,block)
}
