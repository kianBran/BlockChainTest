package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//定义交易结构
type Transaction struct {
	TXID []byte //交易Id
	TXInputs []TXInput //交易输入数组
	TXOutputs []TXOput //交易输出数组
}

//定义交易输入
type TXInput struct {
	//引用的交易ID
	TXid []byte
	//引用的交易输出索引值
	Index int64
	//解锁脚本，我们用地址来模拟
	Sig string
}

//定义交易输出
type TXOput struct {
	//转账金额
	value float64
	//解锁脚本，我们用户地址来模拟
	PubkeyHansh string
}

//设置交易ID
func (tx *Transaction) Sethash()  {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err!=nil{
		log.Panic(err)
	}
	data:=buffer.Bytes()
	hash:=sha256.Sum256(data)
	tx.TXID=hash[:]

}
