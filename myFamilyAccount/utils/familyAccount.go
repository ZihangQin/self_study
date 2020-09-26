package utils

import "fmt"

type FamilyAccount struct {
	//添加结构体字段名
	//声明一个变量，保存接受用户输入的选项
	Key string
	//声明一个变量，控制是否退出
	Loop bool
	//定义账户余额
	Balance float64
	//每次收支的金额
	Money float64
	//每次收支的说明
	Note string
	//定义变量记录是否有收支信息
	Flag bool
	//收支的详情（用字符串来接受）
	//当有收支信息时用details进行处理
	Details string
}
//编写一个工厂模式的构造方法，返回一个*FamilyAccount的实例
func NweFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		Key : "",
		Loop : true,
		Balance : 10000.0,
		Money : 0.0,
		Note : "",
		Flag : false,
		Details : "收支\t 账户金额\t 收支金额\t 说明",
	}
}



//收支明细方法
func (this *FamilyAccount) showdetails() {
	fmt.Println("=====当前收支明细=====")
	if this.Flag {
		fmt.Println(this.Details)
	}else {
		fmt.Println("当前还没有交易信息，去进行交易吧。。")
	}
}
//登记收入方法
func (this *FamilyAccount) income(){
	fmt.Println("收入记录。。。")
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.Money)
	this.Balance += this.Money//进行修改账户余额
	fmt.Println("本次收入的说明：")
	fmt.Scanln(&this.Note)
	//将交易信息记录到details
	this.Details += fmt.Sprintf("\n收入\t%v   \t%v    \t%v",this.Balance,this.Money,this.Note)
	this.Flag = true
}
//支出方法
func (this *FamilyAccount) pay(){
	fmt.Println("支出记录。。。")
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.Money)
	//判断
	if this.Money >  this.Balance {
		fmt.Println("余额不足。")
		//return
	}
	this.Balance -=this.Money
	fmt.Println("支出说明：")
	fmt.Scanln(&this.Note)
	this.Details += fmt.Sprintf("\n支出\t%v   \t%v    \t%v",this.Balance,this.Money,this.Note)
	this.Flag = true
}
//退出方法
func (this *FamilyAccount) exit(){
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
		this.Loop = false
	}
}




//给该结构体绑定方法
//显示主菜单
func (this *FamilyAccount) MainMenu()  {

	for {
		fmt.Println("\n====家庭收支记账系统===")
		fmt.Println("	  1收支明细")
		fmt.Println("	  2登记收入")
		fmt.Println("	  3登记支出")
		fmt.Println("	  4退出软件")
		fmt.Print("请选择（1-4）")
		fmt.Scanln(&this.Key)

		switch this.Key {
		case "1":
				this .showdetails()
		case "2":
				this.income()
		case "3":
				this.pay()
		case "4":
				this.exit()
		default:
			fmt.Println("请输入正确选项=====")
		}
		if !this.Loop {
			break
		}
	}
	fmt.Println("已退出家庭记账软件。")
}