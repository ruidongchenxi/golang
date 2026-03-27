package main

import (
	"fmt"
	"time"
)

//write
func writeData(intchan chan int){
	for i :=1; i<51;i++{
		intchan <-i
		fmt.Println("writeData",i)
		time.Sleep(time.Second)//休眠1秒看交互效果
	}
	close(intchan)//关闭停止写
}
func reaData(intchan chan int,execchan chan bool){
	for{
		v ,ok :=<-intchan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据了=%v\n",v)
	}
	execchan <-true
}

func main(){
		//创建
	intChan := make(chan int,50)
	exitChan := make(chan bool,1)
	go writeData(intChan) //协程
	go reaData(intChan,exitChan) //协程

	for {
		t :=<-exitChan
		if t {
			break
		} 
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo07\main.go
// writeData 1
// readData 读到数据了=1
// writeData 2
// readData 读到数据了=2
// writeData 3
// ...
// readData 读到数据了=50