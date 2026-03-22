package main

import (
	"fmt"
	//"time"
	//"net"
	"src01/go_code/src/chapter14/customerManage/model"
	"src01/go_code/src/chapter14/customerManage/service"
)
type coustmerView struct{
	key string
	loop bool
	coustmerService *service.CustomerService
}
//显示主菜单
func (cv *coustmerView) mainMenu(){
	t1:
	for{
		//
		fmt.Println("------------------------------客户信息管理系统-----------------------------")
		fmt.Println("                              1. 添 加 客 户")
		fmt.Println("                              2. 修 改 客 户")
		fmt.Println("                              3. 删 除 客 户")
		fmt.Println("                              4. 客 户 列 表")
		fmt.Println("                              5. 退       出")
		fmt.Printf("请选择（1-5）：")
		fmt.Scanln(&cv.key)
		switch cv.key{
		
		case "1":
			cv.add()
		case "2":
			cv.revise()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			t2:
			fmt.Println("请选择是否退出")
			v := ""
			fmt.Scanln(&v)
			if v == "y" || v == "Y" {
				break t1
			}else if v == "n"|| v == "N" {
				goto t1
			}else {
				goto t2
			}
		}
	}
	fmt.Println("退出客户关系管理")
}
func (this *coustmerView)list(){
	//获取到当前所有客户信息
	customers := this.coustmerService.List()
	fmt.Println("---------------------------客户列表-------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:=0;i<len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------------------客户信息列表完成--------------------")
}
func (this *coustmerView)add(){
	fmt.Println("----------------------------添加客户------------------")
	fmt.Printf("姓名")
	name :=""
	fmt.Scanln(&name)
	fmt.Printf("性别")
	gender :=""
	fmt.Scanln(&gender)
	fmt.Printf("年龄")
	age :=0
	fmt.Scanln(&age)
	fmt.Printf("电话")
	phone :=""
	fmt.Scanln(&phone)
	fmt.Printf("邮件")
	email :=""
	fmt.Scanln(&email)
	//coustmer:= model.Customer{Nane: name,Gender: gender,Age:age,Phone:phone,Email: email }
	coustmer := model.NewCustomer2(name,gender,age,phone,email)
	//this.coustmerService.Add(coustmer)
	if this.coustmerService.Add(coustmer){
	fmt.Println("-----------------------------------添加完成--------------------------------")
	}else {
		fmt.Println("-----------------------------------添加失败--------------------------------")	
	}
}


func (this *coustmerView)revise(){
	fmt.Println("----------------------------修改客户------------------")
	fmt.Printf("输入要修改的用户ID ")
	var t int
	fmt.Scanln(&t)
	if !this.coustmerService.FindID(t){
		fmt.Println("ID不存在")
		return
	}
		fmt.Printf("姓名")
		name :=""
		fmt.Scanln(&name)
		fmt.Printf("性别")
		gender :=""
		fmt.Scanln(&gender)
		fmt.Printf("年龄")
		age :=0
		fmt.Scanln(&age)
		fmt.Printf("电话")
		phone :=""
		fmt.Scanln(&phone)
		fmt.Printf("邮件")
		email :=""
		fmt.Scanln(&email)
		coustmer := model.NewCustomer2(name,gender,age,phone,email)
		this.coustmerService.Revise(coustmer,t)
	fmt.Println("-----------------------------------修改完成--------------------------------")
	//}else {
	//	fmt.Println("-----------------------------------添加失败--------------------------------")	
//	}
}

func (this *coustmerView)delete(){
	fmt.Println("----------------------------删除客户------------------")
	t :=-1
	fmt.Printf("输入要删除的用户ID(-1退出): ")
	fmt.Scanln(&t)
	if t == -1 {
		return
	}
	fmt.Printf("是否真的要删除（y/n)")
	s := ""
	fmt.Scanln(&s)
	
	if s == "y"{
		if this.coustmerService.Delete(t){
			fmt.Println("-----------------------删除完成-----------------------")
		}else{
			fmt.Println("删除ID不存在")
		}
	}
	
	if this.coustmerService.Delete(t){

	}

}

func main(){
	coustmerView :=coustmerView{
		key: "",
		loop: true,
		//coustmerService: service.NewCustomerService(),
	}
	coustmerView.coustmerService=service.NewCustomerService()
	coustmerView.mainMenu()

}
