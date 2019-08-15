package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	//前區塊哈希
	PrvHash []byte
	//當前區塊哈希
	Hash []byte
	//區塊數據
	Data []byte
}

//創建區塊
func NewBlock(data string,preBlockHash []byte) *Block{
	block:=Block{
		PrvHash: preBlockHash,
		Hash:    []byte{},
		Data:    []byte(data),
	}
	block.SetHash()
	return &block
}

//生成哈希
func (block*Block) SetHash()  {
	//1、拼裝數據
	blockInfo:=append(block.PrvHash,block.Data...)
	//2、sha256哈希
	hash := sha256.Sum256(blockInfo)
	block.Hash=hash[:]

}

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

func main()  {
	bc:=NewBlockChain()
	for i,block:=range bc.blocks{
		fmt.Printf("=============當前區塊高度: %d===========\n",i)
		fmt.Printf("前區塊哈希: %x\n",block.PrvHash)
		fmt.Printf("當前區塊哈希: %x\n",block.Hash)
		fmt.Printf("區塊數據: %s\n",block.Data)
	}

}
