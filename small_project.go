package main

import (
	"fmt"
)

func main() {

	//声明一个变量，保存接受用户输入的选项
	key := ""
	//声明一个变量，控制是否退出
	loop := true
	//定义账户余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//收支的详情（用字符串来接受）
	//当有收支信息时用details进行处理
	details := "收支\t 账户金额\t 收支金额\t 说明"
	//每次收支的说明
	note := ""
	//定义变量记录是否有收支信息
	flag := false
	//显示主菜单
	for{
		fmt.Println("\n====家庭收支记账系统===")
		fmt.Println("	  1收支明细")
		fmt.Println("	  2登记收入")
		fmt.Println("	  3登记支出")
		fmt.Println("	  4退出软件")
		fmt.Print("请选择（1-4）")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("=====当前收支明细=====")
			if flag == true {
				fmt.Println(details)
			}else {
				fmt.Println("当前还没有交易信息，去进行交易吧。。")
			}
		case "2":
			fmt.Println("收入记录。。。")
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money//进行修改账户余额
			fmt.Println("本次收入的说明：")
			fmt.Scanln(&note)
			//将交易信息记录到details
			details += fmt.Sprintf("\n收入\t%v   \t%v    \t%v",balance,money,note)
			flag = true
		case "3":
			fmt.Println("支出记录。。。")
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			//判断
			if money >  balance {
					fmt.Println("余额不足。")
					break
			}
			balance -= money
			fmt.Println("支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v   \t%v    \t%v",balance,money,note)
			flag = true
		case "4":
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
				loop = false
			}
		default:
			fmt.Println("请输入正确选项=====")
		}
		if !loop {
			break
		}
	}

	fmt.Println("已退出家庭记账软件。")
}
