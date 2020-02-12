package main

import (
	"database/sql"
	"fmt"
	//不要忘记饮用驱动
	_"github.com/go-sql-driver/mysql"
)


func main() {
	//打开链接
	db,err:=sql.Open("mysql","root:123@tcp(localhost:3306)/first")
	db.Ping()
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	if err !=nil{
		fmt.Println("数据库打开失败")
		return
	}
	stmt,err:=db.Prepare("select * from people")
	if err!=nil {
		fmt.Println("预处理出错")
		return
	}
	defer func() {
		if stmt!=nil {
			stmt.Close()
		}
	}()
	rows,err:=stmt.Query()
	if err != nil {
		fmt.Println("获取结果失败")
		return
	}
	defer func() {
		if rows != nil{
			rows.Close()
		}
	}()
	for rows.Next() {
		var id int
		var name string
		var address string
		rows.Scan(&id,&name,&address)
		fmt.Println(id,name,address)
	}
}