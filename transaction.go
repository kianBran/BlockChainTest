package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const reword=12.5

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
	//引用的output索引值
	Index int64
	//解锁脚本，我们用地址来模拟
	Sig string
}

//定义交易输出
type TXOput struct {
	//转账金额
	Value float64
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

//实现一个函数，判断当前交易是否为挖矿交易
func (tx * Transaction)IsCoinbase()  bool {
	//1、交易input只有一个
	if len(tx.TXInputs)==1{
		input:=tx.TXInputs[0]
		//2、交易Id为空
		//3、交易的index为-1
		if bytes.Equal(input.TXid,[]byte{}) || input.Index !=-1{
			return false
		}
	}
	return true
}

//提供创建交易方法（挖矿交易）
func NewCoinBaseTX(address string,data string) *Transaction {
	//挖矿交易的特点
	//1、只有一个input
	//2、无需引用交易id
	//3、不需要index
	//旷工由于挖矿时无需指定签名，所以这个sig字段由矿工自由填写数据，一般是填写矿池的名字
	input:=TXInput{[]byte{},-1,data}
	output:=TXOput{reword,address}

	tx:=Transaction{[]byte{},[]TXInput{input},[]TXOput{output}}
	tx.Sethash()

	return &tx

}

//创建普通的转账交易


//3、创建outputs
//4、如果有零钱要找零

func NewTransaction(from,to string,amount float64,bc *BlockChain)*Transaction  {
	//1、找到最合理UTXO集合 map[string][]uint64
	utxos,resValue:=bc.FindNeedUTXOs(from,amount)

	if resValue<amount{
		fmt.Printf("余额不足，交易失败")
		return nil
	}

	var inputs []TXInput
	var outputs []TXOput

	//2、创建交易输入，将这些UTXO注意转成inputs
	for id,indexArray:=range utxos{
		for _,i:=range indexArray{
			input :=TXInput{[]byte(id),int64(i),from}
			inputs=append(inputs,input)
		}
	}

	//创建交易输出
	output :=TXOput{amount,to}
	outputs=append(outputs,output)

	if resValue>amount{
		//找零
		outputs=append(outputs,TXOput{resValue-amount,from})
	}

	tx:=Transaction{[]byte{},inputs,outputs}
	tx.Sethash()
	return &tx

}