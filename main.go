package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"strconv"
	"strings"
)

var count int
var score int
var slice []string

type File struct {
	path string//文件路径
	loop bool//匿名字段，用于控制循环
}

func main() {
	//1、读取答案
	//1.1、打开答案文件
	file, err := excelize.OpenFile("./genesis_file.xlsx")
	if err != nil {
		fmt.Println("文件打开错误，请重试", err.Error())
	}

	//1.2、按列获取文件中的答案列
	//slice := make([]string, 0)  //将答案放在切片中
	for i := 2; i <= 161; i++ { //获取答案文件夹中的答案列中的答案
		cell, err := file.GetCellValue("Sheet1", "G"+strconv.Itoa(i))
		if err != nil {
			fmt.Println("答案获取错误，请重试！", err.Error())
		}
		slice = append(slice, cell) //将答案添加到切片中
	}
	fmt.Println("标准答案是：", slice)
	fmt.Println("================================================================================================")
	//2、查询文件并将每个文件都跟标准答案对比
	queryFile().MainMenu()

}
//实例化一个查询文件
func queryFile() *File {
	return &File{
		path: "",
		loop: true,
	}
}
func (this *File) getclass3()  {
	fmt.Println("三班")
	pathname := "./class3"
	GetFile(pathname)
	fmt.Println()

}
func (this *File) getclass4()  {
	fmt.Println("四班")
	pathname := "./class4"
	GetFile(pathname)
	fmt.Println()
}
func (this *File) exit(){
	fmt.Println("是否要退出 y/n")
	choice :=""
	for   {
		fmt.Scanln(&choice)
		if choice =="y"|| choice == "n" {
			break
		}
		fmt.Println("输入有误，请从新输入")
	}
	if choice == "y" {
		this.loop = false
	}
}
func (this *File) MainMenu()  {
	for  {
		fmt.Println("开始查询，请输入要查询的班级")
		//class3 := "./class3"
		fmt.Println("1、查询三班")
		//class4 := "./class4"
		fmt.Println("2、查询四班")
		fmt.Println("3、退出查询")
		fmt.Println("请选择（1-3）")
		fmt.Scanln(&this.path)
		switch this.path {
		case "1":
				this.getclass3()
		case "2":
				this.getclass4()
		case "3":
				this.exit()
		default:
			fmt.Println("请输入正确选项")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("已退出查询")
}
//使用回调函数去遍历每一个文件
func GetFile(path string) error {

	rd, err := ioutil.ReadDir(path)

	for _, fi := range rd {
		count = 0
		score = 0
		if fi.IsDir() {
			GetFile(path + fi.Name() + "\\")
		}
		//else {
		//	fmt.Println(fi.Name())
		//}
		name := strings.Split(fi.Name(), ".")
		//fmt.Println("名字:",name[0])

		stSlice, err := ReadFile(path + "\\" + fi.Name())
		if err != nil {
			return err
		}

		//将每个文件的答案与标准对比答案
		for i := 0; i < len(slice); i++ {
			if slice[i] == stSlice[i] {
				score++
			}
			if stSlice[i] != "" {
				count++
			}
		}
		fmt.Printf("%s  实际答题数: %d 得分: %d\n", name[0], count, score)
	}
	return err
}
//获取每个文件中的答案列
func ReadFile(path string) ([]string, error) {
	slice := make([]string, 0)
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 获取工作表中指定单元格的值
	for i := 2; i <= 161; i++ {
		cell, err := f.GetCellValue("Sheet1", "G"+strconv.Itoa(i))
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		slice = append(slice, cell)
	}
	return slice, nil
}