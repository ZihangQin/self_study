package Leven

import "fmt"

type Iterator interface {
	//判断是否有下一个值
	Next() bool
	//遍历键值
	Key() []byte
	Value() []byte
}

//键值的结构体
type Pair struct {
	Key   []byte
	Value []byte
}

//迭代器的结构
type DefaultIterartor struct {
	data   []Pair //键值对
	index  int    //索引
	length int    //长度
}

//创建默认的迭代器
func NewDefauItIterartor(data map[string][]byte) *DefaultIterartor {
	self := new(DefaultIterartor)
	self.index = -1
	self.length = len(data)
	for k, v := range data {
		p := Pair{
			Key:   []byte(k),
			Value: v,
		}
		//遍历出的数据添加到data
		self.data = append(self.data, p)
	}
	return self
}

func (self *DefaultIterartor) Next() bool {
	//还有值
	if self.index < self.index-1 {
		self.index++
		return true
	}
	return false
}
func (self *DefaultIterartor) Key() []byte {
	if self.index == -1 || self.index >= self.length {
		panic(fmt.Errorf("IndexOutOfBoindError"))
	}
	return self.data[self.index].Key
}

func (self *DefaultIterartor) Value() []byte {
	if self.index >= self.length {
		panic(fmt.Errorf("IndexOutOfBoindError"))
	}
	return self.data[self.index].Value
}
