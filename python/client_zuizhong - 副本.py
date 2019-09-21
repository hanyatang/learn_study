from socket import *
import json

HOST = 'localhost'  #  或 '127.0.0.1'
POST = 3002
BUFSIZ = 1024
ADDR = (HOST, POST)

tcpCliSock = socket(AF_INET, SOCK_STREAM)
tcpCliSock.connect(ADDR)
print('连接成功，请输入您的内容')
while True:
    
    
    data = input('客户端： ')
    if data == 'quit':
        print('退出程序。。。')
        break
    data = json.dumps(data)
    tcpCliSock.send(data.encode('utf-8'))

    
    data = tcpCliSock.recv(BUFSIZ)
    receive_data =json.loads(data) 

    if not data:
       
        break
    receive_data1 =receive_data
    print(receive_data1)
    

tcpCliSock.close()