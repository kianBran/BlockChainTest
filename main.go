package main

import "fmt"

func main()  {
	bc:= NewBlockChain()
	bc.AddBlock("這是我的第一個區塊")
	bc.AddBlock("這是我的第二個區塊")

	//创建迭代器
	it := bc.NewIterator()

	//调用迭代器,返回每一个区块的数据
	for{
		//返回区块数据，并且游标左移
		block := it.Next()
		fmt.Printf("前區塊哈希: %x\n",block.PrvHash)
		fmt.Printf("當前區塊哈希: %x\n",block.Hash)
		fmt.Printf("區塊數據: %s\n",block.Data)

		if len(block.PrvHash)==0{
			fmt.Printf("区块链遍历结束")
			break
		}


	}

}
