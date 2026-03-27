// 在主线程(可以理解成进程)中，开启goroutine,该协程每隔1秒输出"hello,world"
// 在主线程中也每隔一秒输出"hello,golang",输出10次后，退出程序
// 要求主线程和goroutine同时执行
package main

import (
	"fmt"
	"strconv"
	"time"
)

//编写一个函数，每隔1秒输出"hello,world"
func test(){
	for i:=1;i<=5;i++{
		fmt.Println("test hello,world"+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main(){
	go test()//开启一个协程，同时与主函数共同执行
	for i:=1;i<=5;i++{
		fmt.Println("main,Golang"+ strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\goroutinedemo\main.go
// main,Golang1
// test hello,world1
// test hello,world2
// main,Golang2
// test hello,world3
// main,Golang3
// main,Golang4
// test hello,world4
// test hello,world5
// main,Golang5