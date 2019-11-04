package model

import (
	"fmt"
)
//声明一个Customer的结构体

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}
//编写一个工程模式，返回一个Customer的实例
func NewCustomer(id int,name string,gender string,
age int,phone string,email string) Customer {
	return Customer {
		Id : id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}
func NewCustomer2(name string,gender string,
	age int,phone string,email string) Customer {
		return Customer {
			Name : name,
			Gender : gender,
			Age : age,
			Phone : phone,
			Email : email,
		}
	}

//返回用户的格式化信息
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,this.Name,
this.Gender,this.Age,this.Phone,this.Email)
	return info
}