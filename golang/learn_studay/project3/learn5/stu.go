package stu

import (
	"fmt"
)

type Family struct {
	key string
	loop bool
	data string
	money float64
	note float64
	mote string
	doop bool
	abc  string
}
func NewFamiyl() *Family {
	return&Family {
		key : "",
		loop : true,
		data : "收支\t账户余额\t收支金额\t说明",
		money : 10000,
		note : 0,
		mote : "",
		doop : true,
		abc : "",
	}

}
func (this *Family)show(){
	if this.doop == true {
		fmt.Println("没有金额请来一笔")
	}else{
		fmt.Println("-------------收支明细------------")
		fmt.Println(this.data)
	}
}
func (this *Family) income(){
	this.doop = false
	fmt.Println("请输入金额")
	fmt.Scanln(&this.note)
	fmt.Println("请输入收支明细")
	fmt.Scanln(&this.mote)
	this.money += this.note
	this.data += fmt.Sprintf("\n收支  \t%v         \t%v       \t%v\t",this.money,this.note,this.mote)
}
func (this *Family) pay() {
	this.doop = false
	fmt.Println("请输入金额")
	fmt.Scanln(&this.note)
	fmt.Println("请输入收支明细")
	fmt.Scanln(&this.mote)
	if this.note > this.money {
		fmt.Println("你没有这么多钱")
	}else{
		this.money -= this.note
		this.data += fmt.Sprintf("\n支出  \t%v         \t%v       \t%v\t",this.money,this.note,this.mote)

	}	
}
func (this *Family) exit() {
	fmt.Println("你确定要退出吗 y/n")
			
		
	this.loop = true
	for {
		fmt.Scanln(&this.abc)
		if this.abc == "y" || this.abc == "n" {
			break
		}
			fmt.Println("输入错误请从新输入")
		
	}
}
func (this *Family) MainMenu() {
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
		case "1":
			this.show()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请从新输入")
		}
		if this.abc == "y"{
			fmt.Println("您已退出程序")
			break
		}
	
	}	
}