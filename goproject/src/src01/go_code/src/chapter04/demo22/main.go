package main
import (
	"fmt"
	//"strconv"
	//"math/rand"
	//"time"
)
func main() {
	label1:
	for i :=1;i<100;i++{
		for j :=1;j<100;j++{
			if i + j > 20 {
				fmt.Println(i,"加",j,"和大于20等于",i+j)
				break label1
			} 
		}
	}
	logchance :=3
	for i := 1;i<=3;i++{
		var name string
		var pwd string
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)
		if name == "张无忌" && pwd == "888" {
			fmt.Println("恭喜登录成功")
			break
		}else{
			logchance --
			fmt.Println("你还有",logchance,"次机会，请珍惜")
		}
	} 
	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo22\main.go
// 1 加 20 和大于20等于 21
// 请输入用户名
// 张无忌
// 请输入密码
// 888
// 恭喜登录成功