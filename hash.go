package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)



/*func HASH( text string, hashType string, isHex bool) string {
	var hashInstance hash.Hash
		switch hashType {
		case "md5":
			hashInstance = md5.New()
		case "sha1":
			hashInstance = sha1.New()
		case "sha256":
			hashInstance = sha256.New()
		case "sha512":
			hashInstance = sha512.New()
		}
	if isHex {
		arr , _ := hex.DecodeString(text)
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}
	bytes := hashInstance.Sum(nil)
	return fmt.Sprintf("%x",bytes)
}
*/
func SHA256Double(text string, isHex bool) []byte {
	hashInstance := sha256.New()
	if isHex {
		arr, _ := hex.DecodeString(text) //匿名变量
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}
	bytes := hashInstance.Sum(nil)
	hashInstance.Reset() //第二次调用sha256.New。节省内存
	hashInstance.Write(bytes)
	bytes = hashInstance.Sum(nil)
	return bytes
}

//通过SHA256Double获得一个字节数组

func SHA256DoubleString(text string, isHex bool) string {
	bytes := SHA256Double(text, isHex)
	return fmt.Sprintf("%x", bytes)
}


	func main() {
	res := SHA256DoubleString("0", true)
	//第一个变量输入需要加密的明文密码；
	//第二个变量可以控制变为：md5、sha1、sha256、sha512
	//第三个变量输入true/false。true就是String hash，false就是Binary hash
	fmt.Println(res)
}