package main

import (
	"crypto/sha256"
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

