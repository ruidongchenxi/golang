package main

import (
	"fmt"
	"time"
)

func asyncPrint() {
	fmt.Println("chen")
}
//主协程
func main(){
	//主死随从
	go asyncPrint()//异步执行；
	for i:=0;i<100;i++{
	// 为什么不按顺序重复；
	// 1.闭包问题，2for循环问题for循环的时候，每个变量会重用，，当进入到第二轮的时候for循环时候，这个i的值就变了；所以进入for循环第一时间将i赋值给一个变量；
	// go协程启动的时候打印新变量;或者值传递
		//tmp:=i
		//go func (){
		go func (i int)  {

			//fmt.Println(tmp)
			fmt.Println(i)
		//}()
		}(i)
	}
	fmt.Println("main goroutine")
	time.Sleep(5*time.Second)

}