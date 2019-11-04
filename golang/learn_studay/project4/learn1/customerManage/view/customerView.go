package main

import (
	"fmt"
	"project4/learn1/customerManage/service"
	"project4/learn1/customerManage/model"
	
)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService

}
//显示所有的客户信息
func (this *customerView) list() {
	//获取到当前所有用户的信息
	customers := this.customerService.List()
	fmt.Println("------------------客户列表--------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i :=0;i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------------客户列表完成--------------")
}
//拿到切片的信息
func (this *customerView) add() {
	fmt.Println("------------------客户列表--------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := " "
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	Customer := model.NewCustomer2(name,gender,age,phone,email)
	this.customerService.Add(Customer)
	if this.customerService.Add(Customer) {
		fmt.Println("---------------添加完成------------")
	}else{
		fmt.Println("---------------添加失败-------------")
	}	
}
func (this *customerView) exit() {
	fmt.Println("你确定要退出么 Y/N")
	
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "n"{
			break
		}
		fmt.Println("你输入有误请从新输入 y/n")
			
	}
		
	if this.key == "y" {
		this.loop = false
	}
}

func (this *customerView) update() {
	fmt.Println("----------------客户信息如下---------------")
	id := -1
	fmt.Println("请输入修改的客户id")
	fmt.Scanln(&id)
	fmt.Println(this.customerService.Update(id))

	

	
}
//删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("--------------客户信息管理软件----------------")
	fmt.Println("请选择删除编号（-1放弃）")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return 
	}
	if this.customerService.Delete(id){
		fmt.Println("--------------------删除完成----------------")

	}else {
		fmt.Println("-------------------删除失败------------------")
	}
}
//显示主菜单
func (this *customerView)  mainMenu() {
	for {
		fmt.Println("\n--------------客户信息管理软件----------------")
		fmt.Println("              1.添加客户")
		fmt.Println("              2.修改客户")
		fmt.Println("              3.删除客户")
		fmt.Println("              4.客户列表")
		fmt.Println("              5.退出")
		fmt.Println("              请选择(1-5)")
		fmt.Println("请输入内容：")
		fmt.Scanln(&this.key)
		switch this.key {
		case  "1":
			this.add()
		case "2" :
			this.update()

			
		case "3" :
			this.delete()
		case "4" :
			this.list()
		case "5" :
			this.exit()
			
		default:
			fmt.Println("请从新输入")
		
		}
		if !this.loop {
			break
		}
	}
fmt.Println("退出程序")
}
func main() {
	customerView := customerView{
		key : " ",
		loop : true,
	}
	//完成对customerService初始化
	customerView.customerService = service.NewCustomerSercice()
	customerView.mainMenu()
}