package main

import (
	"fmt"
	//"go/scanner"
	//"strconv"
)

type Details struct{
	balance float64
	money float64
	note string
	detail string
}
func (n *Details)accounts(){
	fmt.Println(n.detail)
}
func (n *Details)Increase(f float64,s string){
	n.money=f
	n.balance+=n.money
	n.detail= fmt.Sprintf("支出\t%v\t\t%v\t\t%v\n",n.balance,n.money,s)
}
func (n *Details)Expenditure(f float64,s string){
	n.money=f
	n.balance-=n.money
	n.detail=fmt.Sprintf("收入\t%v\t\t%v\t\t%v\n",n.balance,n.money,s)
}
func main(){
	var Acc Details=Details{balance: 10000}
	t1:
	for {
		fmt.Println("-----------------------家庭收支记账----------------")
		fmt.Println("                         1.收支明细")
		fmt.Println("                         2.登记收入")
		fmt.Println("                         3.登记支出")
		fmt.Println("                         4.退出    ")
		fmt.Printf("请选择1-4: ")
		//声明变量接收用户
		var t string
		var money float64
		var note string
		fmt.Scanln(&t)
		switch t {
			case "1":
				Acc.accounts()
			case "2":
				fmt.Printf("输入收入金额：")
				fmt.Scanln(&money)
				fmt.Printf("输入收入说明：")
				fmt.Scanln(&note)
				Acc.Increase(money,note)
				
			case "3":
				fmt.Printf("输入支出金额：")
				fmt.Scanln(&money)
				if money< Acc.balance{
				Acc.Expenditure(money,note)
				}else{
					fmt.Println("你发的存款不够你这次消费")
					break
				}
				fmt.Printf("输入支出说明：")
				fmt.Scanln(&note)
				
			case "4":
				break t1 
			default:
				fmt.Println("输入正确菜单")
		}
	}
	fmt.Println("退出这个记账软件的使用")
}