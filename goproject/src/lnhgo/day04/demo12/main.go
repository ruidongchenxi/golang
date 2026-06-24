package main

import (
	"fmt"
	"sync"
	"time"
)
//ar done bool
var lock sync.Mutex
var done = make(chan struct{})//chan 是多协程安全的;管道要初始化
//很多时候不会多个协程写同一个管道
func g1(ch chan struct{}){
	time.Sleep(time.Second)
	ch <- struct{}{}//
}
func g2(ch chan struct{}){
	time.Sleep(2*time.Second)
	//time.Sleep(time.Second)
	ch <- struct{}{}
}
func main(){
	// g1chan:=make(chan struct{})
	// g2chan:=make(chan struct{})
	g1chan:=make(chan struct{},1)
	g2chan:=make(chan struct{},2)
	// g1chan <-struct{}{}
	// g2chan <-struct{}{}
	go g1(g1chan)
	go g2(g2chan)
	// <-g1chan
	// <-g2chan
	//要监控多个channel任何一个channel有值都知道
	//1.某一个分支就绪就执行该分支
	//2.如果两个都就绪那就随机执行一个
	timer:=time.NewTicker(time.Second)
	for{
		select{
		case <-g1chan://
			fmt.Println("g1 done")
		case <-g2chan:
			fmt.Println("g2 done")
		case <- timer.C:
			//time.Sleep(10*time.Second)
			fmt.Println("timeout")
			return 
		}
	}
	
}