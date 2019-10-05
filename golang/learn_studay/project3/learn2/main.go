package main

import(
	"fmt"
	"strconv"
)



func main() {
	n,err := strconv.Atoi("123")
	if err == nil {
		fmt.Println(n)
	}else{
		fmt.Println("错误是",err)
	}//返回错误 验证输入是否正确
	
	str :="hell,beiing"

	str = strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T",str,str)
	
}