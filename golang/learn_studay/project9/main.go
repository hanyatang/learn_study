package main

import (
	"fmt"
)
var (
	first int
	symbol string
	two int
	user string
	result int
	
)
func main() {
	
	for {
		fmt.Println("欢迎来到go语言计算机输入1进入，输入q退出")
		fmt.Scanln(&user)
		if user == "1"{
			
		
		
		
			fmt.Println("请输入第一个数字:")
			fmt.Scanln(&first)
			fmt.Println("请输出符号:")
			fmt.Scanln(&symbol)
			fmt.Println("请输出第二个数字:")
			fmt.Scanln(&two)
			switch symbol {
			case "+":
				fmt.Println("计算结果为:",first + two)
				
			case "-":
				result,_=fmt.Println("计算结果为:",first - two)
				
			case "*":
				result,_=fmt.Println("计算结果为:",first * two)
				
			case "/":
				result,_=fmt.Println("计算结果为:",first / two)
			default:
				fmt.Println("输出错误请正确输入")
				}
		}
		if user == "q" {
			fmt.Println("欢迎再次使用")
			break
		}
		
	}
	
}