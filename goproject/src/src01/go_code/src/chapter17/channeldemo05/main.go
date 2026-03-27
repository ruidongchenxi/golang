package main

import (
	"fmt"
	"math/rand"
	"time"
)
type Cat struct{
	Name string
	Age int
	Address string
}
func main(){
	rand.Seed(time.Now().UnixNano())
	var t chan Cat
	t = make(chan Cat, 10)
	s :=""
	for i :=0;i<10;i++{
		s =fmt.Sprintf("%v",rand.Intn(10)+1)
		i:=Cat{s,i,s}
		t<-i

	}
	for i :=0;i<10;i++{
		fmt.Println(<-t)
	}


}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo02\main.go
// 小花猫