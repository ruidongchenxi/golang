package main
import (
	"fmt"
)
func main(){
	var mosters []map[string]string
	mosters = make([]map[string]string,2)//初始化切片
	if mosters[0]== nil{
		mosters[0]= make(map[string]string)
		mosters[0]["名字"]="牛魔王"
		mosters[0]["年龄"]="1600"	
	}
	if mosters[1]== nil{
		mosters[1]= make(map[string]string)
		mosters[1]["名字"]="牛魔王"
		mosters[1]["年龄"]="1600"	
	}
	a := map[string]string{
		"name": "黑熊怪",
		"age" : "1400",
	}
	mosters= append(mosters, a)
	fmt.Println(mosters)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter08\deom05\main.go
// [map[名字:牛魔王 年龄:1600] map[名字:牛魔王 年龄:1600] map[age:1400 name:黑熊怪]]