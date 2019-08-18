package main

import "fmt"

func (cli *CLI) AddBlock(data string)  {
	cli.bc.AddBlock(data)
	fmt.Printf("添加区块链成功\n")

}

func (cli*CLI) PrintBlockChain() {
	//创建迭代器
	bc := cli.bc
	it := bc.NewIterator()

	//调用迭代器,返回每一个区块的数据
	for {
		//返回区块数据，并且游标左移
		block := it.Next()
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前區塊哈希: %x\n", block.PrvHash)
		fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
		fmt.Printf("时间戳: %v\n", block.TimeStamp)
		fmt.Printf("难度值: %d\n", block.Difficulty)
		fmt.Printf("随机数: %d\n", block.Nonce)
		fmt.Printf("當前區塊哈希: %x\n", block.Hash)
		fmt.Printf("區塊數據: %s\n", block.Data)

		if len(block.PrvHash) == 0 {
			fmt.Printf("区块链遍历结束")
			break
		}

	}

}
