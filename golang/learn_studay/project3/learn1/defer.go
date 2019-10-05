package main

import(
	"fmt"
)

func learn(n1,n2 int) int {
	defer fmt.Println(n1)//3
	defer fmt.Println(n2)//2
	res := n1 + n2
	fmt.Println(res)//1
	return res
}
func main(){
	res := learn(10,20)
	fmt.Println(res)//4
	
}