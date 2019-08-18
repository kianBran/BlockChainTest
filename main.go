package main

func main()  {
	bc:= NewBlockChain()
	cli:=CLI{bc}
	cli.Run()
	/*bc.AddBlock("這是我的第一個區塊")
	bc.AddBlock("這是我的第二個區塊")



	}*/

}
