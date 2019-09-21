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
    a = data.split(',')
    b = a[0]
    if b[0] == 'c':

        class Myclass:
            def __init__(self):
                self.caozuo = a[0]
                self.zhiling = a[1]
                self.id = a[2]
        myclass = Myclass()
        myclassDict = myclass.__dict__
        
        data = json.dumps(myclassDict)
        tcpCliSock.send(data.encode('utf-8'))
        #添加了else
    else:
        data = json.dumps(b)
        tcpCliSock.send(data.encode('utf-8'))
        

    
    data = tcpCliSock.recv(BUFSIZ)
    receive_data =json.loads(data) 

    if receive_data[0] == "88":
        tcpCliSock.close()
        break
    receive_data1 =receive_data
    print(receive_data1)
    

exit()

