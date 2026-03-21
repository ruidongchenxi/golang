package main

import (
	"fmt"
	//"strconv"
)


func main(){
	//显示主菜单
    balance:=10000.00
	// 每次收支金额
	money :=0.0
	//每次收支说明
	var note string
	//收支使用的字符段
	//当有收支时只需要对你进行拼接处理
	details := "收支\t账户金额\t收支金额\t说明\n" 
	var bool = false  //标志位
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
		fmt.Scanln(&t)
		switch t {
			case "1":
				if bool{
					fmt.Println(details)
				}else {
					fmt.Println("还没有存款来一笔吧")
				}
			case "2":
				fmt.Printf("输入收入金额: ")
				fmt.Scanln(&money)
				balance+=money
				fmt.Printf("输入收入说明: ")
				fmt.Scanln(&note)
				details += fmt.Sprintf("收入\t%v\t\t%v\t\t%v\n",balance,money,note)
				bool= true
			case "3":
				fmt.Printf("输入支出金额: ")
				fmt.Scanln(&money)
				if money< balance{
					balance-=money
				}else{
					fmt.Println("你发的存款不够你这次消费")
					break
				}
				if balance<=0{
					bool=false
				}

				fmt.Printf("输入支出说明: ")
				fmt.Scanln(&note)
				details += fmt.Sprintf("支出\t%v\t\t%v\t\t%v\n",balance,money,note)
			case "4":
				var b string
				fmt.Printf("是否真要退出选择y/n:")
				fmt.Scanln(&b)
				for {
				if b == "y"|| b=="Y"{
					break t1 
				}else if b=="n"|| b=="N"{
					break
				}else{
					fmt.Printf("输入有误请重新选择y/n:")
					fmt.Scanln(&b)
				}
				
				// switch b {
				// case "y","Y": 
				// 	
				// case "n","N":
				// 	fmt.Println("")
				// default:
				// 	fmt.Println("输入不正确")
				// }
			 	}
			default:
				fmt.Println("输入正确菜单")
		}
	}
	fmt.Println("退出这个记账软件的使用")
}