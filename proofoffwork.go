package main

import "math/big"

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
	return []byte("HelloWorld"),10

}

