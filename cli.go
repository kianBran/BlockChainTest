package main

import (
	"fmt"
	"os"
)

//这是一个接受命令行参数并且控制区块链操作的文件
type CLI struct {
	bc *BlockChain
}

const Usage=`
	addBlock --data DATA    "add data to blockChain"
	printChain              "print all biockChain data"

	`

//接受参数的动作我们放到一个函数中
func (cli *CLI)Run()  {


	//./block printChain
	//1、得到所有的参数
	args := os.Args
	if len(args)<2{
		fmt.Printf(Usage)
		return
	}

	//分析命令
	cmd:=args[1]
	switch cmd{
	//执行相应操作
		case "addBlock":
			fmt.Printf("//添加区块\n")

			//确保命令有效
			// ./block addBlock --data "helloworld"
			if len(args)==4 && args[2]=="--data"{
				//获取命令的数据
				//a、获取数据
				data:=args[3]
				//b、使用bc添加区块 AddBlock()
				cli.AddBlock(data)
			}else {
				fmt.Printf("参数使用不当，请检查")
				fmt.Printf(Usage)
			}

		case "printChain":
			fmt.Printf("//打印区块链\n")
			cli.PrintBlockChain()


	default:
		fmt.Print("无效命令")
		fmt.Print(Usage)
	}



}
