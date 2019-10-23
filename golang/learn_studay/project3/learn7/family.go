package family

import (
	"fmt"
)
type FamilyAccount struct {
	key string
	loop bool
	balance float64
	money float64
	note string
	datails string
	flag bool
	zhanghao string
	mima float64
}
//编写一个工厂模式的构造方法
func NewFamilyAccount() *FamilyAccount {
	return&FamilyAccount{
		key: "",
		loop:true,
		balance:10000.0,
		money:0.0,
		note:"",
		flag:false,
		datails:"收支\t账户余额\t收支金额\t说明",
		zhanghao:"",
		mima:123,
	}
}
//将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	if this.flag {
		fmt.Println("----------收支明细------------")
		fmt.Println(this.datails)
		} else{
			fmt.Println("没有金额，来一笔把")
		}
}
//登记写成方法
func (this *FamilyAccount)income() {
	fmt.Println("本次收入的金额：")
	fmt.Scanln(&this.money)
	//修改用户总金额
	this.balance += this.money
	fmt.Println("本次收入的说明：")
	fmt.Scanln(&this.note)
	//拼接字符串
	this.datails += fmt.Sprintf("\n收入\t%v	\t%v	\t%v",this.balance,this.money,this.note)
	this.flag = true
}
func(this *FamilyAccount)pay(){
	fmt.Println("----------登记支出------------")
	fmt.Println("本次支出的金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("本次出支出大于总金额")
				
	}
	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.datails +=	fmt.Sprintf("\n支出\t%v	\t%v	\t%v",this.balance,this.money,this.note)
	this.flag = true
}
func(this *FamilyAccount)exit(){
	fmt.Println("你确定要退出么 Y/N")
	data := ""
	for {
		fmt.Scanln(&data)
		if data == "Y" || data == "N"{
			break
		}
			fmt.Println("你输入有误请从新输入 Y/N")
		}
	if data == "Y" {
		this.loop = false
	}
}
func (this *FamilyAccount)user() {
	for {
		fmt.Println("请输入账号:")
		fmt.Scanln(&this.zhanghao)
		fmt.Println("请输入密码")
		fmt.Scanln(&this.mima)
		if this.zhanghao == "abc" && this.mima == 123{
			fmt.Println("恭喜你成功进入")
			break
		}else{
			fmt.Println("账号/密码错误,请从新输入")
			continue
		}
	}
}
func(this *FamilyAccount) MainMenu() {
		this.user()
		for {
			fmt.Println("\n--------------家庭收支记账软件----------------")
			fmt.Println("              1.收支明细")
			fmt.Println("              2.登记收入")
			fmt.Println("              3.登记支出")
			fmt.Println("              4.退出")
			fmt.Println("              请选择(1-4)")
			fmt.Println("请输入内容：")
			fmt.Scanln(&this.key)
			switch this.key {
			case  "1":
				this.showDetails()
				
			case "2" :
				this.income()
			case "3" :
				this.pay()
			case "4" :
				this.exit()
				
			default:
				fmt.Println("请从新输入")
			
			}
			if !this.loop {
				break
			}
		}
	fmt.Println("退出家庭记账程序")
}