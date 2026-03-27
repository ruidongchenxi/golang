package main

import (
	"fmt"
	"time"
)

//write
func writeData(intchan chan int){
	for i :=1; i<=50;i++{
		intchan <-i
		fmt.Println("writeData",i)
		//time.Sleep(time.Second)//休眠1秒看交互效果
	}
	close(intchan)//关闭停止写
}
func reaData(intchan chan int,execchan chan bool){
	for{
		time.Sleep(time.Second)
		 _,ok :=<-intchan
		if !ok {
			break
		}
		fmt.Println("reaData",<-intchan)
		
	}
	execchan <-true
}

func main(){
		//创建
	intChan := make(chan int,10)
	//reaChan := make(chan int,10)
	exitChan := make(chan bool,1)
	go writeData(intChan) //协程
	go reaData(intChan,exitChan) //协程

	for {
		_,t :=<-exitChan
		if t {
			break
		} 
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo07\main.go
// writeData 1
// writeData 2
// writeData 3
// writeData 4
// writeData 5
// writeData 6
// writeData 7
// writeData 8
// writeData 9
// writeData 10
// ...
// reaData 50