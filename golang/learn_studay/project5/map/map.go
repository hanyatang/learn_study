package main
   
import (
	"fmt"
)


func main(){
	/*name := ""
	shouji := ""
	iphone :=make(map[string]string)
	fmt.Println("请输入你的姓名和电话号码")
	fmt.Scanln(&name)
	fmt.Scanln(&shouji)
	iphone[name] = shouji
	fmt.Println(iphone)
	//delete(iphone,输入删除的key)删除操作
	*/

	m :=map[string]string{"张三":"1331231231","李四":"2131231231"}
	m["王五"] = "123123123"
	fmt.Println(m)//添加操作
	//如果有的话就是修改
	m["张三"] = "11111111"
	fmt.Println(m)//修改操作
	//删除操作
	delete(m,"张三")
	fmt.Println(m)
	//查看操作
	fmt.Println(m["李四"])
	//越界查看返回空值
	//循环例边map里的值
	for key,value := range m {
		fmt.Println(key,value)

		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("\n")
		delete(m,"李四")
		fmt.Println(m)
	}


}