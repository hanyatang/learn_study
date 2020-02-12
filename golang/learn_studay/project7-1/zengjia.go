package project7

import (
	"database/sql"
	"fmt"
	//不要忘记饮用驱动
	_"github.com/go-sql-driver/mysql"
)


func zengjia() {
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
	//预处理sql
	//?表示占位符
	stmt,err:=db.Prepare("insert into people values(default,?,?)")
	defer func() {
		if stmt !=nil {
			stmt.Close()
		}
	}()
	if err !=nil{
		fmt.Println("预处理失败")
		return
	}
	r,err:=stmt.Exec("张三","海淀")
	if err !=nil{
		fmt.Println("sql执行失败")
		return
	}

	//获取结果	
	count,err := r.RowsAffected()
	if err !=nil{
		fmt.Println("结果获取失败")
		return
	}
	if count > 0 {
		fmt.Println("新增成功")
	}else{
		fmt.Println("新增失败")
	}

	id,_ :=r.LastInsertId()
	fmt.Println(id)
	

}