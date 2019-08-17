package main

import (
	"bytes"
	_ "crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	//1、版本號
	Version uint64
	//2、前區塊哈希
	PrvHash []byte
	//3、Merkel根
	MerkelRoot []byte
	//4、時間戳
	TimeStamp uint64
	//5、難度值
	Difficulty uint64
	//6、隨機數，挖礦要找的數據
	Nonce uint64

	//正常比特幣區塊中沒有當前區塊哈希，我們爲了方便自己定義
	//a、當前區塊哈希
	Hash []byte
	//b、區塊數據
	Data []byte
}

//實現一個輔助函數，功能是將uint64轉成[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err!=nil{
		log.Panic(err)
	}
	return buffer.Bytes()
}

//創建區塊
func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrvHash:    preBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}
	//block.SetHash()
	//创建一个pow对象
	pow:=NewProofofWork(&block)
	//查找随机数，不停地进行哈希运算
	hash,nonce:=pow.Run()

	//根据挖矿结果对区块数据进行更行
	block.Hash=hash
	block.Nonce=nonce
	return &block
}
//序列化
func (block *Block)Serialize() []byte {
	//编码的数据放到buffer中
	var buffer bytes.Buffer

	//使用gob进行序列化（编码）得到字节流
	//1、定义一个编码器
	//2、使用编码器进行编码
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err !=nil{
		log.Panic("编码出错")
	}
	//fmt.Printf("编码后的小明：%x\n",buffer.Bytes())

	return buffer.Bytes()
}


//反序列化
func Deserialize(data []byte) Block {
	//使用gob进行反序列化（解码）的得到person结构
	//1、定义一个解码器
	decoder := gob.NewDecoder(bytes.NewReader(data))
	//2、使用解码器进行解码
	var block Block
	err := decoder.Decode(&block)
	if err!=nil{
		log.Panic("解码失败")
	}
	//fmt.Printf("解码后的大明：%v\n",daming)
	return block

}

//生成哈希
/*func (block *Block) SetHash() {
	/*var blockInfo []byte
	//1、拼裝數據
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrvHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Data...)*/
	
	/*tmp:=[][]byte{
		Uint64ToByte(block.Version),
		block.PrvHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}

	blockInfo := bytes.Join(tmp, []byte{})

	//2、sha256哈希
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]

}*/
