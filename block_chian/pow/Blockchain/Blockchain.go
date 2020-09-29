package Blockchain

import (
	"block_chian/pow/Block"
	"fmt"
)

//通过链表的形式维护区块链中的业务

type Node struct {
	//指针域
	NextNode *Node
	//数据域
	Data *Block.Block
}

//创建头结点，用于保存创世区块
func CreatHeaderNode(data *Block.Block) *Node {
	//创建头结点
	headNode := new(Node)
	//将头结点的指针域指向nil
	headNode.NextNode = nil
	//数据域保存传入的data
	headNode.Data = data
	//返回头结点
	return headNode
}

//添加结点
//挖矿成功添加区块
func AddNode(data *Block.Block, node *Node) *Node {
	//创建新节点
	var newNode *Node = new(Node)
	//指针域为nil
	newNode.NextNode = nil
	//数据域为data
	newNode.Data = data
	node.NextNode = newNode
	return newNode

}

//查看链表中的数据
func ShowNodes(node *Node) {
	n := node
	for {
		//如果有数据就打印,没有下个结点就结束循环
		if n.NextNode == nil {
			fmt.Println(n.Data)
			break
		} else {
			fmt.Println(n.Data)
			n = n.NextNode
		}
	}
}
