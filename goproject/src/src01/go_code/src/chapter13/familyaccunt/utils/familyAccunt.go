package utils

import (
	//"flag"
	"fmt"
)
type FamilyAccount struct {
	//声明字段
	key   string
	//声
	loop bool
	//
	balance float64
	money float64
	note string
	flag bool
	details string
}
//显示明细做成方法
func (n *FamilyAccount)showdetails(){
			if n.flag{
				fmt.Println(n.details)
			}else {
				fmt.Println("还没有存款来一笔吧")
			}
}
func (n *FamilyAccount)income(){
			fmt.Printf("输入收入金额: ")
			fmt.Scanln(&n.money)
			n.balance+=n.money
			fmt.Printf("输入收入说明: ")
			fmt.Scanln(&n.note)
			n.details += fmt.Sprintf("收入\t%v\t\t%v\t\t%v\n",n.balance,n.money,n.note)
			n.flag=true
}
func (n *FamilyAccount)pay(){
		fmt.Printf("输入支出金额: ")
		fmt.Scanln(&n.money)
		if n.money< n.balance{
			n.balance-=n.money
		}else{
			fmt.Println("你发的存款不够你这次消费")
			//return
		}
		if n.balance<=0{
			n.flag=false
		}

		fmt.Printf("输入支出说明: ")
		fmt.Scanln(&n.note)
		n.details += fmt.Sprintf("支出\t%v\t\t%v\t\t%v\n",n.balance,n.money,n.note)
}
func (n *FamilyAccount)exit(){
	var b string
	fmt.Printf("是否真要退出选择y/n:")
	for {
		fmt.Scanln(&b)
		if b=="y"|| b == "n"{
			break
		}
		fmt.Println("输入有误")
	}
	if b == "y"{
		n.loop= false
	}

}

func (n *FamilyAccount)MainMenu (){
			for n.loop{
			 		fmt.Println("-----------------------家庭收支记账----------------")
		     		fmt.Println("                         1.收支明细")
					fmt.Println("                         2.登记收入")
					fmt.Println("                         3.登记支出")
					fmt.Println("                         4.退出    ")
					fmt.Printf("请选择1-4: ")
					//声明变量接收用户
					//var t string
					fmt.Scanln(&n.key)
					switch n.key {
						case "1":
							n.showdetails()
						case "2":
							n.income()

						case "3":
							n.pay()
						case "4":
						n.exit()
				}
	fmt.Println("退出这个记账软件的使用")
}
}
func NewFamilyAccount() *FamilyAccount{
	return &FamilyAccount{
		key: "",
		loop: true,
		balance: 10000,
		money: 0.0,
		note: "",
		flag: false,
		details: "收支\t账户金额\t收支金额\t说明\n",
	}
}