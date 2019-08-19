package main

import (
	_ "crypto/sha256"
	"fmt"
	_ "fmt"
	"github.com/boltdb/bolt"
	"log"
)


//引入區塊
type BlockChain struct{
	//定一個區塊數組
	//blocks []*Block
	db *bolt.DB
	tail []byte //存儲最後一個區塊的哈希
}
const BlockChainDB="blockChain.db"
const BlockBucket="blockBucket"

//定義一個區塊鏈
func NewBlockChain(address string) *BlockChain  {
	//創建一個創世區塊，並作爲第一個區塊添加到區塊鏈中
	/*genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks:[]*Block{genesisBlock},
	}*/
	//最後一個區塊的hash，從數據庫中讀出來的
	var lastHash []byte
	//打開數據庫
	db,err:=bolt.Open("BlockChainDB",0600,nil)
	//defer db.Close()

	if err!=nil{
		log.Panic("打開數據庫失敗")
	}

	//將要操作的數據庫改寫
	db.Update(func(tx *bolt.Tx) error {
		//找到抽屜bucket(如果沒有，就創建)
		bucket:=tx.Bucket([]byte("blockBucket"))
		if bucket==nil{
			//沒有抽屜，我們需要創建
			bucket,err=tx.CreateBucket([]byte("blockBucket"))
			if err !=nil{
				log.Panic("創建bucket(blockBucket)失敗")
			}
			//創建一個創世區塊，並作爲第一個區塊添加到區塊鏈中
			genesisBlock := GenesisBlock(address)

			//寫數據
			//hash作爲key，block字節流作爲value
			bucket.Put(genesisBlock.Hash,genesisBlock.Serialize())
			bucket.Put([]byte("LastHashKey"),genesisBlock.Hash)
			lastHash=genesisBlock.Hash
		}else {
			//讀數據
			lastHash=bucket.Get([]byte("LastHashKey"))
		}

		return nil
	})
	return &BlockChain{db,lastHash}

}

//創世快
func GenesisBlock(address string) *Block {
	coinBase:=NewCoinBaseTX(address,"這是一個創世區塊")
	return NewBlock([]*Transaction{coinBase},[]byte{})
}

//添加區塊
func (bc *BlockChain)AddBlock(txs []*Transaction)  {
	//如何獲取前區塊哈希
	db:=bc.db //区块链数据库
	lasthash := bc.tail //最后一个区块的哈希

	//獲取最後一個區塊数据
	db.Update(func(tx *bolt.Tx) error {
		//完成数据添加
		bucket := tx.Bucket([]byte(BlockBucket))
		if bucket==nil{
			log.Panic("bucket不应该为空，请检查")
		}

		//創建新的區塊
		block := NewBlock(txs, lasthash)

		//添加到區塊鏈數組中
		//hash作爲key，block字節流作爲value
		bucket.Put(block.Hash,block.Serialize())
		bucket.Put([]byte("LastHashKey"),block.Hash)

		//更新内存中的lasthash
		bc.tail=block.Hash
		return nil
	})





}

//找到指定地址的所有UTXO
func (bc *BlockChain)FindUTXOs(address string) []TXOput {
	var UTXO []TXOput
	//定義一個map來保存消費過的output，key是這個output的交易ID，value是這個交易中索引的數組（多個交易）
	//map[交易id][]int64
	spentOutputs:=make(map[string][]int64)
	//創建迭代器
	it:=bc.NewIterator()

	for{
		//1、便利區塊
		block:=it.Next()
		//2、便利交易
		for _,tx:=range block.Transactions{
			fmt.Printf("current txid:%x\n",tx.TXID)
			//3、遍歷output,找到和自己相關的utxo（在添加output之前檢查一下是否消耗過）
			OUTPUT:
			for i,output:=range tx.TXOutputs{
				fmt.Printf("current i:%d\n",i)
				//在這裏做一個過濾，將所有消耗過的outputs和當前的所即將添加的output對比一下
				//如果相同，則跳過，否則添加
				//如果當前的交易id存在於我們已經表示的map，那麼說明這個交易裏面有消耗過的output
				//map[2222]=[]inte64{0}
				//map[3333]=[]inte64{0,1}
				//不是創世區塊
				if spentOutputs[string(tx.TXID)]!=nil{
					for _,j:=range spentOutputs[string(tx.TXID)]{
						//inte64{0,1},j:0,1
						if int64(i)==j{
							//當前準備添加的output已經消耗過了，不要再加了
							continue OUTPUT
						}
					}
				}
				//這個output和我們的目標地址相同，滿足條件，加到返回utxo數組中
				if output.PubkeyHansh==address{
					UTXO = append(UTXO, output)
				}
			}
			//如果当前交易是挖矿交易的话，那么不做遍历，直接跳过
			if !tx.IsCoinbase(){
				//4、遍歷input，找到自己花費過的utxo集合（把自己消耗過得標示出來）
				for _,input:=range tx.TXInputs{
					//判斷一下當前這個input和目標（李四）是否一致，如果相同，如果相同，說明這個是李四消耗過的
					if input.Sig==address{
						//spentOutputs:=make(map[string][]int64)
						indexArray:=spentOutputs[string(input.TXid)]
						indexArray = append(indexArray, input.Index)
					}
				}
			}else {
				fmt.Printf("这是普通交易，不做判断")
			}



		}





		if len(block.PrvHash)==0{
			break
			fmt.Printf("區塊遍歷完成")
		}
	}
	return UTXO
}

func (bc *BlockChain)FindNeedUTXOs(from string,amount float64) (map[string][]uint64,float64)  {
	//找到的合理的utxos集合
	var utxos map[string][]uint64
	//找到的utxos里面包含钱的总数
	var calc float64

	return utxos,calc
}

