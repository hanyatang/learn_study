package main

import (
	"fmt"
)

func main() {
	var (
		
		name string
		hepen  int
		
	)
		fmt.Println("请输出你的名字")
		fmt.Scanln(&name)
		fmt.Printf("请%s输出你的努力程度,输入1，2，3程度",name)
		fmt.Scanln(&hepen)
		switch hepen {
		case 1:
			fmt.Println("那就加油吧")
		case 2:
			fmt.Println("还可以啊")
		case 3:
			fmt.Println("垃圾！！！！")
		}
		
		
	
}