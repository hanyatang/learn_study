package service
import (
	"project4/learn1/customerManage/model"
	"fmt"
)

//改CustomerService,完成对Customer的更删改查
type CustomerService struct {
	//声明一个字段，表示当前切片含有多少个客户
	customers []model.Customer
	customerNum int
	
}
//将他变成一个实例
func NewCustomerSercice() *CustomerService{
	CustomerService := &CustomerService{}
	CustomerService.customerNum = 1
	customers := model.NewCustomer(1,"网名","男",20,"123","123@qq.com")
	CustomerService.customers = append(CustomerService.customers,customers)
	return CustomerService
}
func (this *CustomerService) List() []model.Customer{
	return this.customers
}

func (this *CustomerService) Add(customer model.
	Customer) bool{
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers,customer)
	return true
}
//根据id修改客户
func (this *CustomerService) Update(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		fmt.Println("没有改用户")
		return false
	}
	
	
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := " "
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	
	return true
}
//根据id删除客户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
}
//根据id查找客户在切片的下标
func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历this.customers切片
	for i := 0;i <len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}
