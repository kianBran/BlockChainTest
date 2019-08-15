package proofofwork

import "fmt"

func main()  {
	bc:= NewBlockChain()
	bc.AddBlock("這是我的第一個區塊")
	bc.AddBlock("這是我的第二個區塊")
	for i,block:=range bc.blocks{
		fmt.Printf("=============當前區塊高度: %d===========\n",i)
		fmt.Printf("前區塊哈希: %x\n",block.PrvHash)
		fmt.Printf("當前區塊哈希: %x\n",block.Hash)
		fmt.Printf("區塊數據: %s\n",block.Data)
	}

}
