package main

import "fmt"
type Cat struct{
	Name string
	Age int
}
func main(){
	allChan := make(chan interface{},3)
	allChan<-1
	allChan<-3
	cat := Cat{"小花猫",4}
	allChan<-cat
	<-allChan
	<-allChan
	t :=(<-allChan).(Cat)//断言注意类型不能跟变量同名。不然报错
	fmt.Println(t.Name)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo02\main.go
// 小花猫