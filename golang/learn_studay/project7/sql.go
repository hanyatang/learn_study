package main

import (
	"database/sql"
	"fmt"
)


func main() {
	//打开链接
	DB,err:=sql.Open("mysql","root:smallming@tcp(localhost:3306)/first")
	DB.Ping()
	if err !=nil{
		fmt.Println("数据库打开失败")
		return
	}
	//预处理sql




	//获取结果
}