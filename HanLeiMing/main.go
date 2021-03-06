package main

import (
	"net/http"
	"html/template"
	"HanLeiMing/src/user"
	"HanLeiMing/src/commons"
	"github.com/gorilla/mux"
)

//显示登录页面
func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/login.html")
	t.Execute(w, nil)
}
//restful现实页面
func showPage(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	t,_:=template.ParseFiles("view/" + vars["page"]+".html")
	t.Execute(w,nil)
}
func main() {
	
	commons.Router.PathPrefix("/static").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	commons.Router.HandleFunc("/",welcome)
	//满足/page/{page}格式的处理
	commons.Router.HandleFunc("/page/{page}",showPage)
	user.UserHandler()
	http.ListenAndServe(":80",commons.Router)
}
