package main

import "fmt"

func (cli *CLI) AddBlock(data string)  {
	//cli.bc.AddBlock(txs)
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
		fmt.Printf("區塊數據: %s\n", block.Transactions[0].TXInputs[0].Sig)

		if len(block.PrvHash) == 0 {
			fmt.Printf("区块链遍历结束")
			break
		}

	}

}


func (cli *CLI)GetBalance(address string)  {
	utxos:=cli.bc.FindUTXOs(address)
	total:=0.0
	for _,utxos:=range utxos{
		total+=utxos.Value
	}
	fmt.Printf("%s餘額爲：%f\n",address,total)
}
