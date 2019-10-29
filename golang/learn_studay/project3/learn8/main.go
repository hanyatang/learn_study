package main

import (
	"fmt"
)

	

func main() {
	name := ""
	age := 0
	sex := ""
	mark := 0
	key := ""
	booa := true
	detalis := "姓名\t年龄\t性别\t成绩"
	flag := false
	
	for {	
		fmt.Println("-----------学生管理系统-----------")
		fmt.Println("           1.查看学生目录")
		fmt.Println("           2.添加学生信息")
		fmt.Println("           3.修改学生信息")
		fmt.Println("           4.退出")
		fmt.Println("请输入1-4：")
		fmt.Scanln(&key)
		switch key {
			case "1":
				if flag {
					fmt.Println("-------学生信息如下----------")
					fmt.Println(detalis)
				}else{
					fmt.Println("暂无信息,请添加一个")
				}
					
			case "2":
				fmt.Println("请输入学生姓名")
				fmt.Scanln(&name)
				fmt.Println("请输入学生年龄")
				fmt.Scanln(&age)
				fmt.Println("请输入学生性别")
				fmt.Scanln(&sex)
				fmt.Println("请输入学生成绩")
				fmt.Scanln(&mark)
				detalis += fmt.Sprintf("\n%v\t%v\t%v\t%v",name,age,sex,mark)
				flag = true
			case "3":
				
			case "4":
				fmt.Println("你确定要退出么 y/n")
				data := ""
				for {
					fmt.Scanln(&data)
					if data =="y" || data == "n"{
						break
					}
					fmt.Println("请从新输入 y/n")
				}
				if data == "y"{
					booa = false
					fmt.Println("退出程序")
				}
			default:
				fmt.Println("请输入正确的选项")
		}
			if !booa {
				break
			}
	}
	
	
	
}