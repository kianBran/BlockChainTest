package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofofWork struct {
	//a、block
	block*Block
	//b、目標值
	//一個非常大的數，他有很多方法
	target*big.Int
}

//提供創建pow的函數
func NewProofofWork(block *Block) *ProofofWork {
	pow:=ProofofWork{
		block:block,
	}

	//我們指定的難度值，現在是一個string類型，需要進行轉換
	targetStr:="0000100000000000000000000000000000000000000000000000000000000000"

	//引入輔助變量，目的是將上面的難度值轉成big.int
	tmpInt:=big.Int{}
	//將難度值複製給big。int,指定16進制格式
	tmpInt.SetString(targetStr,16)

	pow.target=&tmpInt
	return &pow
}

//3、提供不断计算hash的函数
func (pow *ProofofWork)Run() ([]byte,uint64) {
	block:=pow.block
	var nonce uint64
	var hash [32]byte
	//1、拼装数据(区块的数据，还有不断变化的随机数)
	for{
		tmp:=[][]byte{
			Uint64ToByte(block.Version),
			block.PrvHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			//只對去塊頭做hash，區塊體通過Merkelroot產生影響
			//block.Data,
		}
		block.MerkelRoot=block.MakeMerkelRoot()

		blockInfo := bytes.Join(tmp, []byte{})
		//2、做哈希运算
		hash = sha256.Sum256(blockInfo)
		//3、与pow中的target进行比较
		tmpInt:=big.Int{}
		tmpInt.SetBytes(hash[:])

		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if tmpInt.Cmp(pow.target)==-1{
			//a、找到了。退出返回
			fmt.Printf("挖矿成功! hash:%x,nonce:%d\n",hash,nonce)
			return hash[:],nonce
		}else {
			//b、没找到，继续找，随机数加1
			nonce++
		}



		//return []byte("HelloWorld"),10
		//return hash[:],nonce
	}


}

