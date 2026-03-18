package main

import (
	"fmt"
	//"internal/syscall/windows"
	//"strconv"
)
type Stu struct{
	Name string
	Age int
}
func main(){
	//在创建结构体变量时，就直接指定字段的值
	var stu = Stu{"晨",8}
	te := Stu{"小明",6}
	//在创建结构体变量时，把字段和字段值写在一起，这种方式更灵活，不依赖字段的定义顺序
	te1 := Stu{Name: "小明",Age: 6}
	fmt.Println(stu,te,te1)
	//方式2 返回指针 stu2---存储地址----》结构体实例化空间的内存地址值
	var str2 = &Stu{"小王",29}
	str6 :=&Stu{Age: 9,Name: "杨逍"}
	fmt.Println(*str2,*str6) 


}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter09\deom11\main.go
// {晨 8} {小明 6} {小明 6}
// {小王 29} {杨逍 9}