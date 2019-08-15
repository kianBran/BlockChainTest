package proofofwork

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
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
	block.SetHash()
	return &block
}

//生成哈希
func (block *Block) SetHash() {
	/*var blockInfo []byte
	//1、拼裝數據
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrvHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Data...)*/
	
	tmp:=[][]byte{
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

}
