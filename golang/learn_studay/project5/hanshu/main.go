package main

import (
	"fmt"
	
)


func main() {
	fmt.Println("计算的结果为：")
	//a,b := 1,4
	///sum :=add(a,b)
	//fmt.Println(sum)
	a,b,c := show()
	fmt.Println(a,b,c)
	demo("看书","萨达","大萨达","大萨达")
	func (name string) {//以下是你命好函数
		fmt.Println("这个是匿名函数叫:",name)
	}("傻逼")
	
	
}
func add(a,b int)  int{//接受单个返回值
	sum := a + b
	return sum
	
}
func show() (string,int,string) {//接受多个返回值
	fmt.Println("执行了show")
	return "小明",19,"海淀"
}

func demo(hover ... string) {//可变参数函数
	for _,n := range hover{
		fmt.Println(n)
	}
}

