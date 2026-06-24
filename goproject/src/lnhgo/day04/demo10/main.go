package main

import (
	"fmt"
	"time"
)
func producer(out chan<- int){
	for i:=1;i<10;i++{
		out <- i*i
	}
	close(out)
}
func consumer(in <-chan int){
	for num := range in{
		fmt.Printf("num=%d\r\n",num)
	}
}
func main(){
	// var ch1 chan int//双向
	// var ch2 chan<- float64//单向只能写
	// var ch3 <-chan int  //单向只能取
	// c:=make(chan int,3)
	// var send chan<- int = c//只能写数据
	// var read <-chan int = c //只能取数据
	// send <- 1
	// <-read 
	c:=make(chan int)
	go producer(c)
	go consumer(c)
	//fmt.Println("j")
	time.Sleep(10*time.Second)

}