package main


import (
	"net/http"
	"fmt"
)

func welcome(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","text/html;charset=utf-8")
	fmt.Fprintln(w,"服务器返回的信息<b>加粗</b>")
	
}
func main() {
	http.HandleFunc("/",welcome)
	http.ListenAndServe("localhost:9091",nil)
	
}
