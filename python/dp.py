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




def cheak(data):#查看数据库
    a = data['id']
    cur.execute('select * from  stu limit '+ a )
    
    all = cur.fetchall()
    return  all

    

# def add(mes):#增加数据
#     cur.execute('select * from stu')
#     all = cur.fetchall()
#     sql_insert = 'insert into stu values' + '(' + mes[2] + ',' + mes[3] +','+ mes[4] + ',' + mes[6] + ')'

#     try:
#         cur.execute(sql_insert)  
#         con.commit()  
#     except Exception as e:
#         con.rollback()  
#     finally:
#         cur.close()  
#         con.close()  
#         print(all)
#         return all

# def delete(mes):#修改数据
#     cur.execute('select * from stu')
#     all = cur.fetchall()

#     sql_delete = "delete from stu where ID =" + mes[1]  # 删除SQL语句

#     try:
#         cur.execute(sql_delete)  

                    
#         con.commit()  
#     except Exception as e:
#         con.rollback()  
#     finally:
#         cur.close()  
#         con.close()  
#         print(all)
#         return all

# def revise(mes):#修改数据
#     cur.execute('select * from stu')
#     all = cur.fetchall()
#     sql_updates = 'update stu set name = ' + mes[2] + ',addr =' + mes[3] + ',score =' + mes[4] + ' where id = ' + mes[1]  # 修改SQL语句
#     try:
#         cur.execute(sql_updates)  
#         con.commit()  
#     except Exception as e:
#         con.rollback()  
#     finally:
#         cur.close()  
#         con.close()  
#         print(all)
#         return all




def operate(data):
    
    a = data['zhiling']
    
    if a == 'fetch_mysql':#查看数据库
        return cheak(data)

    # elif cmd == 'add_mysql':#增加数据
    #     return add(mes)

    # elif cmd == 'delete_mysql':#删除数据
    #     return delete(mes)
    # elif cmd == 'revise_mysql':#修改数据
    #     return revise(mes)
    
    return None

