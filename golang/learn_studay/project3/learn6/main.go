package main
//面向对象实例
import (
	"fmt"
	
)

/*type Student struct {
	Name string
	Gender string
	Age int
	Id int
	Score float64
}
func (student *Student) say() string {
	infostr := fmt.Sprintf("student的信息 name = [%v]gender=[%v]age=[%v]id=[%v]score=[%v}",
	student.Name,student.Gender,student.Age,student.Id,student.Score)
	return infostr
}

func main() {
	var stu = Student{
		Name : "tom",
		Gender : "male",
		Age : 20,
		Id : 1000,
		Score : 67.4,
	}
	fmt.Println(stu.say())
}
*/
/*
type Dog struct{
	name string
	age int
	weight string
}
func (Dog *Dog) say() string {
	b :=fmt.Sprintf("dog的信息 name=[%v]agr=[%v]weight=[%v]",
	Dog.name,Dog.age,Dog.weight)
	return b
}
func main() {
	var a = Dog {
		name : "小花",
		age : 19,
		weight : "while",
	}
	fmt.Println(a.say())
}*/

type menp struct {
	name string
	age int
}
func (menp *menp) a()  {
	if menp.age > 18 {
		fmt.Printf("游客名字是%v，年龄为%v收费20\n",menp.name,menp.age)
	}else{
		fmt.Println("免费")
	}
}
func main() {
	var v menp
	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&v.name)
		if v.name == "n"{
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.age)
		v.a()
	}
}