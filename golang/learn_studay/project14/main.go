package main

import (
	"net/http"
	"fmt"
)

func first(w http.ResponseWriter,r *http.Request) {
	h:=r.Header
	fmt.Fprintln(w,h)
	fmt.Fprintln(w,h["Sec-Fetch-Mode"])	
	
	r.ParseForm()
	fmt.Fprintln(w,r.Form)
	fmt.Fprintln(w,r.Form["name"][0])
}



func main() {
//多处理函数方式
	server:=http.Server{Addr:"localhost:8091"}
	http.HandleFunc("/first",first)

	server.ListenAndServe()
	
}
