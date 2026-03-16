package main

import "fmt"

func modify(map1 map[int]int){
	map1[10] =900
}
type sty  struct{
	name string
	age  int 
	add string
}
func test(a map[string]map[string]string,name string){

	if a[name] !=nil { // 判断是map否存在某key
		//有用户
		a[name]["pws"] = "888888"
	} else {
		//fmt.Println("输入用户名")
		a[name] = make(map[string]string)
		a[name]["pwd"]= "88888"
		a[name]["昵称"]= "昵称"+name

	}
}
func main(){
	t := make(map[string]map[string]string)
	t["go"] = make(map[string]string)
	t["go"]["pwd"]="999999"
	t["go"]["昵称"]="Golang"
	test(t,"tom")
	test(t,"dagu")
	test(t,"go")
	fmt.Println(t)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter08\deom05\main.go
// map[dagu:map[pwd:88888 昵称:昵称dagu] go:map[pwd:999999 pws:888888 昵称:Golang] tom:map[pwd:88888 昵称:昵称tom]]