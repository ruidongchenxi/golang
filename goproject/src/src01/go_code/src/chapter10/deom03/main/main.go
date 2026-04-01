package main
import (
	"fmt"
	"src/chapter10/deom03/model"
)
func main(){
	t := model.ChuShi("123456",89,"123456")
	// t.SetZH("678907")
	// t.SetYe(2000)
	// t.SetMa("1239897")
	if t != nil{
		fmt.Printf("账户是%v,余额是%v,密码是%v\n",t.GetZh(),t.GetYe(),t.GetMa())
	}else {
		fmt.Println("失败")
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter10\deom03\main\main.go
// 账户是123456,余额是89,密码是123456