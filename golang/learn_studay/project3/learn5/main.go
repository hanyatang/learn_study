package main

import (
	"fmt"
	"encoding/json"
)
type Monster struct{
	Name string `json:"name"`//将大写变量编程小写
	Age int	`json:"age"`
	Color string `json:"color"`
}//josn将数据字符串化
func main() {
	monster := Monster{"牛魔王",500,"芭蕉扇"}
	josn,err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误",err)
	}
	fmt.Println("josnstr=",string(josn))
	
}
