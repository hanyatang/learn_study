package main

import (
	"fmt"
)
func main() {


	var customers []string =[]string {"网名","男","123","123@qq.com"}
	fmt.Println(customers)
	customers[1] = "你好"
	customers[0] = "哈哈"
	fmt.Println(customers)
}