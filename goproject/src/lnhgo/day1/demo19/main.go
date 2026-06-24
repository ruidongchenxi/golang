package main

import (
	//"flag"
	"fmt"
)
//定义命令行参数
func main() {
	age := 19
	if age >18 {
		fmt.Println("你已经成年")
	}else{
		fmt.Println("写作业")
	}

}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo19\main.go            
// 你已经成年