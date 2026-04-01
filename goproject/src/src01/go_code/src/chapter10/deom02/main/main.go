package main

import (
	//"fmt"
	//"fmt"
	"src/chapter10/deom02/model"
)
func main(){
	s:=model.NewPerson("晨曦")
	s.SetAge(56)
	s.SetSal(25000)
	s.Get()
	s.GetAge()
	s.GetSal()
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter10\deom02\main\main.go
// 名字是晨曦,年龄是56,工资是25000
// 晨曦 的年龄是 56
// 晨曦 的工资是 25000