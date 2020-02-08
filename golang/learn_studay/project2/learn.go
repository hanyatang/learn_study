package main 
//这就是个实验
import (
	"fmt"
	"os"
)

type learn struct{
	name string
	yuwen int
	shuxue int
	yingyu int
}
func showmennu() {
	fmt.Println("欢  迎  来  到  学  生  管  理  系  统")
    fmt.Println("1.添加学生数据")
    fmt.Println("2.修改学生数据")
    fmt.Println("3.显示学生数据")
    fmt.Println("4.退出系统")
}
func userInput() *learn {
	var (
		name string
		yuwen int
		shuxue int
		yingyu int
		)
	fmt.Println("请根据提示输入相关内容")
    fmt.Print("请输入姓名:")
    fmt.Scanln(&name)
    fmt.Print("请输入语文成绩:")
    fmt.Scanln(&yuwen)
    fmt.Print("请输入数学成绩:")
    fmt.Scanln(&shuxue)
    fmt.Print("请输入英语成绩:")
	fmt.Scanln(&yingyu)
	fmt.Println(name,yuwen,shuxue,yingyu)
	book := newstuday(name,yuwen,shuxue,yingyu)
	return book

}
//等待用户输入选项
var allstuday = make([]*learn,0,200)
//定义一个创建新成绩的函数
func newstuday(name string,yuwen,shuxue,yingyu int)*learn{
	return &learn{
		name : name,
		yuwen : yuwen,
		shuxue : shuxue,
		yingyu : yingyu,	
	}
}
func addstuday() {
	var (
		name string
		yuwen int
		shuxue int
		yingyu int
	)
	fmt.Println("请根据提示输入相关内容")
    fmt.Print("请输入姓名:")
    fmt.Scanln(&name)
    fmt.Print("请输入语文成绩:")
    fmt.Scanln(&yuwen)
    fmt.Print("请输入数学成绩:")
    fmt.Scanln(&shuxue)
    fmt.Print("请输入英语成绩:")
    fmt.Scanln(&yingyu)
    fmt.Println(name,yuwen,shuxue,yingyu)
    book := newstuday(name,yuwen,shuxue,yingyu)
    for _, b := range allstuday{
        if b.name == book.name{
			fmt.Printf("%s这位同学已存在",book.name)
			fmt.Print("\n")
			fmt.Print("\n")
			fmt.Print("\n")
            return
        }
    }
    allstuday = append(allstuday,book)
	fmt.Println("添加数据成功！")
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")
	
}
//修改学生数据
func updatestuday() {
	book := userInput()
	for index,b := range allstuday{
		if b.name == book.name{
			allstuday[index] = book
			fmt.Printf("学生%s信息已更新完毕",book.name)
			fmt.Print("\n")
			fmt.Print("\n")
			fmt.Print("\n")
			return
		}
	}
	fmt.Printf("同学%s不存在",book.name)
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")
}
//展示学生信息
func showstuday() {
	if len(allstuday) == 0 {
		fmt.Println("没有信息")
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("\n")

	}
	for _,b :=range allstuday{
		fmt.Printf("%s的成绩为：语文%d,数学：%d，英语：%d",b.name,b.yuwen,b.shuxue,b.yingyu)
		fmt.Print("\n")
		fmt.Print("\n")
		
	}
}
func main() {
	for {
		showmennu()
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			addstuday()
		case 2:
			updatestuday()
		
		case 3:
			showstuday()
		
			
		case 4:
			os.Exit(0)
		}
			
			
	}
	
}



