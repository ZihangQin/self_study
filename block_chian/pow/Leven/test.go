package Leven

import (
	"fmt"
	"testing"
)

//测试默认迭代器
func TestNewDefaultIterator(t *testing.T) {
	data := make(map[string][]byte)
	data["k1"] = []byte("v1")
	data["k2"] = []byte("v2")
	data["k3"] = []byte("v3")

	iter := NewDefauItIterartor(data)

	if iter.length != 3 {
		t.Fatal()
	}

	for iter.Next() {
		fmt.Printf("%s : %s\n", iter.Key(), string(iter.Value()))
	}
}
