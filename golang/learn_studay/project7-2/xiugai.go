package xiugai

import (
	"database/sql"
	"fmt"
	//不要忘记饮用驱动
	_"github.com/go-sql-driver/mysql"
)

func xiugai() {
	db,_:=sql.Open("mysql","root:123@tcp(localhost:3306)/first")
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	stmt,_:=db.Prepare("update people set name=?,address=? where id=?")
	defer func() {
		if stmt !=nil {
			stmt.Close()
		}
	}()
	r,_:=stmt.Exec("李四","朝阳",2)
	count,_:=r.RowsAffected()
	if count >0{
		fmt.Println("修改成功")
	}else{
		fmt.Println("修改失败")
	}
}