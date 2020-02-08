package main

import (
	"net"
	"fmt"
	"strings"
	
)
type User struct {
	Username string
	OtherUsername string
	Msg string
	ServerMsg string
}
var (
	userMap = make(map[string]net.Conn)
	user = new(User)
)

func main() {
	//创建服务器地址
	addr,_ :=net.ResolveTCPAddr("tcp","localhost:8899")
	//创建监听器
	list,_ :=net.ListenTCP("tcp4",addr)
	fmt.Println("服务器已经启动")
	for {
		//通过监听器获取客户端传递过来的数据，他是堵塞的
		conn,_:=list.Accept()

		go func() {
			for{
				b:=make([]byte,1024)
				count,_:=conn.Read(b)
				//分割
				array:=strings.Split(string(b[:count]),"-")
				
				user.Username = array[0]
				user.OtherUsername = array[1]
				user.Msg = array[2]
				user.ServerMsg = array[3]
				userMap[user.Username]=conn
				if v,ok :=userMap[user.OtherUsername];ok&&v!=nil {
					
					n,err:=v.Write([]byte(fmt.Sprintf("%s-%s-%s-%s",user.Username,user.OtherUsername,user.Msg,user.ServerMsg)))
					if n <=0 || err!=nil {
						delete(userMap,user.OtherUsername)
						conn.Close()
						v.Close()
						fmt.Println("if,,,,,,,,,,")
						break
					}
					fmt.Println("消息发送成功")
				}else {
					user.ServerMsg="对方不在线"
					conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s",user.Username,user.OtherUsername,user.Msg,user.ServerMsg)))
				}
			

			}
		}()
	}
}