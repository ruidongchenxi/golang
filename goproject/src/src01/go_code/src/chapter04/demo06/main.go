package main
import (
	"fmt"
	//"math"
)
func main(){

	var t int
	var c int 
	var s bool
	fmt.Println("请输入你的身高、财产、是否帅")
	fmt.Scanf("%d %d %t",&t,&c,&s)
	if t > 180 && c > 1000 && s == true {
		fmt.Println("嫁给他")
	}else if t > 180 || c >1000 || s == true{
		fmt.Println("嫁吧，比上不足，比下有余")
	}else{
		fmt.Println("不嫁")
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo06\main.go
// 请输入你的身高、财产、是否帅
// 185 1200 true
// 嫁给他
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo06\main.go
// 请输入你的身高、财产、是否帅
// 174 800 true
// 嫁吧，比上不足，比下有余
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo06\main.go
// 请输入你的身高、财产、是否帅
// 165 600 false
// 不嫁