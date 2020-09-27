package main

import (
	"block_chian/pow/Block"
	"block_chian/pow/Blockchain"
)

func main() {
	firest := Block.GenerateFirstBlock("创世区块")
	second := Block.GenerateNextBlock("新区块",firest)

	//创建链表
	header := Blockchain.CreatHeaderNode(&firest)
	//加入第二款区块
	Blockchain.AddNode(&second, header)
	Blockchain.ShowNodes(header)

}
