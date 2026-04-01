package service

import (
	//"fmt"
	//"fmt"
	//"fmt"
	"src/chapter14/customerManage/model"
)

//该结构体完成对customer的操作包括增删改查
type CustomerService struct{
	coustmers []model.Customer
	//声明一个字段表示当前切片含有多少个客户
	//该字段后面，也可以作为新客户id 号
	coustmerNum int
}
func NewCustomerService() *CustomerService{
	customerservice :=&CustomerService{}
	customerservice.coustmerNum=1
	customer:=model.NewCustomer(customerservice.coustmerNum,"张三","男",32,"苹果", "123@qq.com")
	customerservice.coustmers=append(customerservice.coustmers,customer)
	return customerservice
}
func (this *CustomerService)List()[]model.Customer{
	return this.coustmers
}
func (this *CustomerService)Add(Customer model.Customer) bool{
	//确定id 规则
	this.coustmerNum++
	Customer.Id=this.coustmerNum
	this.coustmers=append(this.coustmers, Customer)
	return true
}


// func (this *CustomerService)delete(v string){
// 	var t int
// 	for i,_:=range this.coustmers{

// 		if this.coustmers[i].Nane==v{
// 			t=i
// 		}
// 	}
// 	if t == len(this.coustmers)-1{
// 		this.coustmers = this.coustmers[:len(this.coustmers)-1] //删除最后一个元素
// 	}else if t == 0 {
// 		this.coustmers = this.coustmers[1:] //删除开头1个元素
// 	}else{
// 		this.coustmers= append(this.coustmers[:t],this.coustmers[t+1:]...) //删除中间一个元素
// 	}
// }
func (this *CustomerService)Delete(v int) bool{
	index :=this.FindByld(v)
	r:=false
	if index == -1{
		// fmt.Println("删除ID不存在")
		return r
	}
	if index == len(this.coustmers) {
		this.coustmers = this.coustmers[:len(this.coustmers)-1] //删除最后一个元素
		r = true
	}else if index == 0 {
		this.coustmers = this.coustmers[1:] //删除开头1个元素
		r= true
	}else {
		this.coustmers= append(this.coustmers[:index],this.coustmers[index+1:]...) //删除中间一个元素
		r= true
	}
	return r
}
//根据id查索引
func (this *CustomerService) FindByld(id int) int{
	//var t int 
	var b bool
	for i,v:=range this.coustmers{
		if v.Id==id{
			id = i
			b= true
		}
	}
	if b {
		return id
	}else {
		return -1
	}
}
//判断输入ID号是否存在
func (this *CustomerService) FindID(id int) bool{
	//var t int 
	var b bool
	for i,v:=range this.coustmers{
		if v.Id==id{
			id = i
			b= true
		}
	}
	return b
}
//修改 需要传入的参数 用户名及ID
func (this *CustomerService)Revise(Customer model.Customer,id int){
	//确定id 规则
	Customer.Id=id
	index :=this.FindByld(id)
	this.coustmers[index]=Customer
}
