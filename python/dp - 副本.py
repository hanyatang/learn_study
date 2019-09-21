import pymysql
import json

con = pymysql.connect(
host='localhost',
port=3306,
user='root',
password='123',
db='fenxiangniu',
charset='utf8'
)
cur = con.cursor()


  
def quit():#退出程序
    mes = '你好，再见'
    return mes


def cheak(mes):#查看数据库
    row = cur.execute('select * from  stu limit '+ mes[1] )
    all = cur.fetchall()
    return  all

def add(mes):#增加数据
    cur.execute('select * from stu')
    all = cur.fetchall()
    sql_insert = 'insert into stu values' + '(' + mes[1] + ',' + mes[2] + ',' + mes[3] + ',' + mes[4] + ')'
    try:
        cur.execute(sql_insert)  
        con.commit()  
    except Exception as e:
        con.rollback()  
    finally:
        cur.close()  
        con.close()  
        print(all)
        return all

def delete(mes):#修改数据
    cur.execute('select * from stu')
    all = cur.fetchall()

    sql_delete = "delete from stu where ID =" + mes[1]  # 删除SQL语句

    try:
        cur.execute(sql_delete)  

                    
        con.commit()  
    except Exception as e:
        con.rollback()  
    finally:
        cur.close()  
        con.close()  
        print(all)
        return all

def revise(mes):#修改数据
    cur.execute('select * from stu')
    all = cur.fetchall()
    sql_updates = 'update stu set name = ' + mes[2] + ',addr =' + mes[3] + ',score =' + mes[4] + ' where id = ' + mes[1]  # 修改SQL语句
    try:
        cur.execute(sql_updates)  
        con.commit()  
    except Exception as e:
        con.rollback()  
    finally:
        cur.close()  
        con.close()  
        print(all)
        return all
    
    