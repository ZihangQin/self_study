package Leven

import "fmt"

//用结构体定义个DB
type DB struct {
	path string
	data map[string][]byte
}

//模拟开启连接
func New(path string) (*DB, error) {
	self := DB{
		path: path,
		data: make(map[string][]byte),
	}
	return &self, nil
}

//模拟关闭连接
func (self *DB) Close() error {
	return nil
}

//设置键值对
func (self *DB) Put(key []byte, value []byte) error {
	self.data[string(key)] = value
	return nil
}

//取值
func (self *DB) Get(key []byte) ([]byte, error) {
	if v, ok := self.data[string(key)]; ok {
		return v, nil
	} else {
		return nil, fmt.Errorf("NotFound")
	}
}

//删除
func (self *DB) Del(key []byte) error {
	if _, ok := self.data[string(key)]; ok {
		delete(self.data, string(key))
		return nil
	} else {
		return fmt.Errorf("NotFound")
	}
}

//模拟遍历取值
func (self *DB) Iterator() Iterator { //设置迭代器对象
	return NewDefauItIterartor(self.data)
}
