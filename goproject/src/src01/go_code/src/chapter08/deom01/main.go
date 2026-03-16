package main

import "fmt"

//"fmt"
func main(){
   var a map[string]string//声明
	a = make(map[string]string,10)//初始化
	a["a1"]= "宋江"
	a["a2"]= "吴用"
	a["a3"]= "吴用"
	a["a1"]= "武松"
	a["a4"]= "林冲"//v可以重复；无序的key-v结构go
	fmt.Println(a)
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter08\deom01\main.go
// map[a1:武松 a2:吴用 a3:吴用 a4:林冲]