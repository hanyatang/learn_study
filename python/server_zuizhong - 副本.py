#-*-conding:utf-8-*-
import pymysql
from socket import *
from time import ctime
from select import select
import json
import dp
HOST = ''
POST = 3002
BUFSIZ = 1024
ADDR = (HOST, POST)

tcpSerSock = socket(AF_INET, SOCK_STREAM)
tcpSerSock.bind(ADDR)
tcpSerSock.listen(5)
tcpSerSock.setblocking(False)  # 将tcpSerSock设置为非租塞模式
inputs = [tcpSerSock]

con = pymysql.connect(
host='localhost',
port=3306,
user='root',
password='123',
db='fenxiangniu',
charset='utf8'
)
cur = con.cursor()

print('开始连接'+'-'*30)

while True:
    rlist,wlist,xlist = select(inputs,[],[])
    for s in rlist:
        if s is tcpSerSock:
            tcpCliSock, addr = s.accept()
            tcpCliSock.setblocking(False)  # 将tcpCliSock设置为非租塞模式
            inputs.append(tcpCliSock)      # 将tcpCliSock插入inputs中

        else:
            data = s.recv(1024)
            data = json.loads(data)
            mes = data.split(',')
            section = mes[0]
            
            if section == "quit":#退出程序
                result = dp.quit()
                
                result = json.dumps('服务端：'+ format(result))
                s.send(result.encode('utf-8'))
                break
                
            elif section == 'fetch_mysql':#查看数据库
                result = dp.cheak(mes)
                result = json.dumps('服务端：'+ format(result))
                s.send(result.encode('utf-8'))

            elif section == 'add_mysql':#增加数据
                result = dp.add(mes)
                result = json.dumps('服务端：'+ format(result))
                s.send(result.encode('utf-8'))

            elif section == 'delete_mysql':#删除数据
                result = dp.delete(mes)
                result = json.dumps('服务端：'+ format(result))
                s.send(result.encode('utf-8'))
            elif section == 'revise_mysql':#修改数据
                result = dp.revise(mes)
                result = json.dumps('服务端：'+ format(result))
                s.send(result.encode('utf-8'))
                


                
            data1 = json.dumps('服务端：'+format(data))

            s.send(data1.encode('utf-8'))

tcpCliSock.close()
