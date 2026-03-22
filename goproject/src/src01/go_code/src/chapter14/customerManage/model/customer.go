package model

import "fmt"

//声明Customer结构体，表示一个客户信息
type Customer struct{
	Id int
	Nane string
	Gender string //性别
	Age int
	Phone string
	Email string
}
//使用工厂模式返回结构体
func NewCustomer(id int,name string,gender string,age int,phone string,email string) Customer{
	return Customer{
		Id: id,
		Nane: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
} 
func NewCustomer2(name string,gender string,age int,phone string,email string) Customer{
	return Customer{
		//Id: id,
		Nane: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
} 

func (this *Customer)GetInfo() string{
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v \n",this.Id,this.Nane,this.Gender,this.Age,this.Phone,this.Email)
}