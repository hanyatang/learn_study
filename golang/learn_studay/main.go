package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func sayHi(w http.ResponseWriter,r *http.Request)  {
    // 解析指定文件生成模板对象
    tem,err := template.ParseFiles("view/hello.html")
    if err != nil{
        fmt.Println("读取文件失败,err",err)
        return
    }
    // 利用给定数据渲染模板，并将结果写入w
    tem.Execute(w,"候津锋")
}
func main()  {
    http.HandleFunc("/",sayHi)
    err := http.ListenAndServe("127.0.0.1:8888",nil)
    if err != nil{
        fmt.Println("监听失败,err",err)
        return
    }
}