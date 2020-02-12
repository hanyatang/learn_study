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
	stmt,err:=db.Prepare("delete from people where id =?")
	if err!=nil {
		fmt.Println("预处理出错")
		return
	}
	defer func() {
		if stmt!=nil {
			stmt.Close()
		}
	}()
	r,err :=stmt.Exec(1)
	if err!=nil {
		fmt.Println("执行失败")
		return
	}
	count,err := r.RowsAffected()
	if count > 0 {
		fmt.Println("删除成功")
	}else{
		fmt.Println("删除失败")
	}	
}
