package main

import (
	"fmt"
	//"unicode"s
	//"testing"
	//"time"
	"sync"
)

//思路
//1 编写一个函数，计算各个数的阶乘，并放入map
//2 启动的协程多个统一的将结果放入到map
//map 应该是全局的
var (
	myMap=make(map[int]uint64,10)
	//声明全局互斥锁
	//sync是包
	//mutex互斥意思
	lock sync.Mutex
	wg sync.WaitGroup
)
//test 

func tset(n int){
	defer wg.Done()
	res :=uint64(1)
	
	for i:=1;i<=n;i++{
		res *=uint64(i)
	}
   //写操作之前加锁
   lock.Lock()
	myMap[n]=res //fatal error: concurrent map writes
	lock.Unlock() //解锁
}
func main(){
	for i :=1 ;i<=50;i++{
		wg.Add(1)
		go tset(i)
	}
	wg.Wait()
	//time.Sleep(time.Second*10)

	for k,v:=range myMap{
		fmt.Printf("map[%v]=%v\n",k,v)
	}
	
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\goroutinedemo03\main.go
// map[1]=1
// map[5]=120
// map[9]=362880
// map[10]=3628800
// map[8]=40320
// map[4]=24
// map[2]=2
// map[3]=6
// map[6]=720
// map[7]=5040