package main

import (
	"fmt"
)

// func primeNum(ints chan int,prim chan int,exits chan bool){
// 	//使用
// 	//var num int 
// 	var flag bool
// 	for {
// 		num,ok :=<-ints
// 		if !ok{
// 			break 
// 		}
// 		flag = true
// 		//判断是不是素数
// 		for i:=2;i<num;i++{
// 			// if num ==2  || num ==1{
// 			// 	prim <-num
// 			// }else if num % i==0{

// 			// }
// 			if num % i == 0{//说明不是素数
// 				flag=false
// 				break
// 			}
// 		}
// 		if flag {
// 			prim <-num 
// 		}

// 	}
// 	fmt.Println("有一个协程完成工作")
// 	exits<-true


// }

func primeNum(ints chan int,prim chan int,x *int){
	//使用
	//var num int 
	var flag bool
	for {
		num,ok :=<-ints
		if !ok{
			break 
		}
		flag = true
		//判断是不是素数
		for i:=2;i<num;i++{

			if num % i == 0{//说明不是素数
				flag=false
				break
			}
		}
		if flag {
			prim <-num 
		}

	}
	fmt.Println("有一个协程完成工作")
	*x++


}

func main(){
	intChan:=make(chan int,1000)
	primeChan := make(chan int,2000)//
	//标识退出管道
	//exit := make(chan bool,4)
//开启协程响，放intChan放入1-8000个数
    go  func (intChan chan int){
		for i :=0; i<80;i++{
			intChan<-i
		}
		close(intChan) 
	}(intChan)
//开启4个协程从int 管道取数据，并判断是不是为素数，如果是放入到primechan
    var x int
	for i :=0 ;i <4 ;i++{
		go primeNum(intChan,primeChan,&x)
		//go primeNum(intChan,primeChan,exit)

	}
	// if x == 4 {
	// 	close(primeChan)
	// }
	for {
		if x == 4{
			close(primeChan)
			break
		}
	}
	
	// 	for v := range primeChan{
	// 	fmt.Println(v)
	// 	}
		
	// go func ()  {
	// 	for i:=0;i<4;i++{
	// 		<-exit
	// 	}
	// 	close(primeChan)
	// }()
	for {
		res,ok :=<-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n",res)
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run  chapter17\channeldemo09\main.go
// 有一个协程完成工作
// 有一个协程完成工作
// 有一个协程完成工作
// 有一个协程完成工作
// 素数=0
// 素数=1
// ...
// 素数=73
// 素数=79