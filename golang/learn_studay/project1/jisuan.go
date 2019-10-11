package main

import (
	"fmt"
)
var First float32
var FuHao string
var Two float32
var exit string
func main() {

	
	for {
	fmt.Print("输入1为计算器，输入q为退出:")
	fmt.Scanln(&exit)
	fmt.Print("\n")	
			if exit == "1"{
				fmt.Print("请输入第一个数字:")
				fmt.Scanln(&First)
				
				if First == 0 {
					fmt.Print("输入错误，请从新输入——————————————")
					fmt.Print("\n")
					continue
				}
				fmt.Print("请输入运算符号:")
				fmt.Print("只能输入 “+ - * /”:")

				fmt.Scanln(&FuHao)

			
				fmt.Print("请输入第二个数字:")
				fmt.Scanln(&Two)
				if Two == 0 {
					fmt.Print("输入错误.请从新输入————————————")
					fmt.Print("\n")
					continue
				}
				if Two == 0 {
					fmt.Print("输入错误,请从新输入")
					fmt.Print("\n")
					continue
				}
				if FuHao == "+" {
					Addition()	
				}
				if FuHao == "-" {
					Subtaction()
				}
				if FuHao == "*" {
					Nultip()
				}
				if FuHao == "/" {
					Division()					
					}				
				}
	if exit == "q"{
		fmt.Print("已为您退出")
		break

		} else {
			fmt.Print("欢迎来到go语言计算机")
			fmt.Print("\n")
		} 
	}
}
func Addition() {
	var JieGuo float32 = First + Two
	fmt.Print("计算结果为：",JieGuo)
	fmt.Print("\n")
	
}
func Subtaction() {
	var JieGuo float32 = First - Two
	fmt.Print("计算结果为：",JieGuo)
	fmt.Print("\n")
}
func Nultip() {
	var JieGuo float32 = First * Two
	fmt.Print("计算结果为：",JieGuo)
	fmt.Print("\n")
}
func Division() {
	var JieGuo float32 = First / Two
	fmt.Print("计算结果为：",JieGuo)
	fmt.Print("\n")		
}
