package main
import (
	"fmt"
)
//定义一个结构体
type Account struct {
	AccountNo string
	pwd string
	Balance float64

}
//方法
func (a *Account) Deposite(money float64,pwd string){
	//
	if pwd != a.pwd {
		fmt.Println("密码错误")
		return 
	}
	if money<=0 {
		fmt.Println("存款金额不正确")
		return
	}

	a.Balance+=money
	fmt.Printf("存款成功，余额是%v\n",a.Balance)
}
func (a *Account) WithDrw(money float64,pwd string){
	if pwd != a.pwd{
		fmt.Println("密码错误")
		return
	}
	if money > a.Balance|| money<=0{
		fmt.Printf("余额不足，无法取%v\n",money)
		return
	}
	a.Balance-=money
	fmt.Printf("成功取钱%v元，账户余额%v\n",money,a.Balance)

}
func (a *Account) Query(pwd string){
	if pwd == a.pwd{
		fmt.Printf("账户%v，余额=%v\n",a.AccountNo,a.Balance)
	}else{
		fmt.Println("密码有误")
	}
}
func main(){
	t := &Account{"456765123","666666",67000}
	look:
	for {
		var R string
		var c float64 
		//var P string
		fmt.Println("请选择操作的选项")
		fmt.Println("1 查账户")
		fmt.Println("2 存款")
		fmt.Println("3 取款")
		fmt.Println("4 退出")
		fmt.Scanln(&R)

		switch R{
		case "1":
			t.Query("666666")
		case "2":
			fmt.Println("输入存款金额")
			fmt.Scanln(&c)
			t.Deposite(c,"666666")
		case "3":
			fmt.Println("输入取款金额")
			fmt.Scanln(&c)
			t.WithDrw(c,"666666")
		case "4":
			break look
		default:
			fmt.Println("输入不正确重新输入")

		}
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter10\deom01\mian.go
// 请选择操作的选项
// 1 查账户
// 2 存款
// 3 取款
// 4 退出
// 4
