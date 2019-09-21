class Student(object):
    def __init__(self,name,enging,chinese):
        self.name = name
        self.enging = enging
        self.chinese = chinese

    def add():
        name = input('请输入姓名')
        enging = input('请输入英语成绩')
        chinese = input('请输入语文成绩')
            
        studinfo = {}
        studinfo['姓名'] = name
        studinfo['英语'] = enging
        studinfo['语文'] = chinese
        return studinfo
        '''def get_gart(self):
            if self.enging >= 90:
                return 'A'
            else:
                return '不合格'''
            

a = Student.add()   
print(a)
import json
a = input('')
class Myclass:
    def __init__(self):
        self.zhiling = a
        self.id = 3
myclass = Myclass()
myclass.c=123
myclass.a=3

myClassDict = myclass.__dict__

print (myClassDict)

myClassJson = json.dumps(myClassDict)

print (myClassJson)
myClassReBuild = json.loads(myClassJson)
#打印重建的字典
print (myClassReBuild)






