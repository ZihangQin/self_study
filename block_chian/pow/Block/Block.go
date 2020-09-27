package Block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//通过代码实现挖矿

//1、区块声明结构体
type Block struct {
	//上一块的hash
	PreHash string
	//当前区块的hash
	HashCode string
	//时间戳
	TimeStamp string
	//当前网络的难读系数
	//控制hash有几个前导0的
	Diff int
	//存储交易信息
	Data string
	//区块高度
	Index int
	//随机值
	Nonce int
}

//创建创世区块
func GenerateFirstBlock(data string) (Block) {
	//声明并赋值一个区块，创建创世块
	var firstblock Block
	firstblock.PreHash = "0"
	firstblock.TimeStamp = time.Now().String()
	firstblock.Diff = 4
	firstblock.Data = data
	firstblock.Index = 1
	firstblock.Nonce = 0
	firstblock.HashCode = GenerationHashValue(firstblock)
	return firstblock
}


//生成区块的hash值
func GenerationHashValue(block Block) string {
	var hashdata = strconv.Itoa(block.Index)+
		strconv.Itoa(block.Nonce)+strconv.Itoa(block.Diff)+block.TimeStamp
	//哈希算法
	var sha = sha256.New()
	sha.Write([]byte(hashdata))
	hashed := sha.Sum(nil)
	//将字节转为字符串
	return hex.EncodeToString(hashed)
}


//产生新区块
func GenerateNextBlock(data string, oldBlock Block) Block {
	//产生一个新区块
	var newBlock Block
	newBlock.TimeStamp = time.Now().String()
	newBlock.Diff= 4
	newBlock.Index = 2
	newBlock.Data = data
	newBlock.PreHash = oldBlock.HashCode
	newBlock.Nonce = 0//矿工去调整的
	newBlock.HashCode = pow(newBlock.Diff,&newBlock)//利用pow算法进行
	return newBlock
}


//pow工作量证明算法进行hash碰撞,传指针保证是统一对象
func pow(diff int, block *Block) string {
	for   {
		hash := GenerationHashValue(*block)
		//没挖一次打印一次hash值
		fmt.Println(hash)
		//判断hash值前导
		if strings.HasPrefix(hash ,strings.Repeat("0",diff)) {
			fmt.Println("挖矿成功")
			return hash
		}else {
			//挖矿失败
			block.Nonce++
		}
	}
}




