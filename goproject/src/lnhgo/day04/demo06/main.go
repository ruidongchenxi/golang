package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
锁 资源竞争
*/
var total int32
var t sync.WaitGroup
var lock sync.Mutex//锁千万不要复制否则就会出现锁失效
func add(){
	defer t.Done()
	for i:=0;i<10000000;i++{
		atomic.AddInt32(&total,+1)
		//lock.Lock() //操作数据前加锁
		//total+=1
		//lock.Unlock()//操作完解锁
	}
}
func sub(){
	defer t.Done()
	for i:=0;i<10000000;i++{
		atomic.AddInt32(&total,-1)
		//lock.Lock()  //枷锁
		//total -=1
		//lock.Unlock() //解锁
	}
}
func main(){
	t.Add(2)
	go add()
	go sub()
	t.Wait()
	fmt.Println(total)
}