package main
import (
	"fmt"
    "sync"
)
//加入互斥锁
var lock sync.Mutex
var t int
var wg sync.WaitGroup
//var lock sync.Mutex
func add(){
	defer wg.Done()
	for i :=0;i < 100000;i++{
		//枷锁
		lock.Lock()
		t = t +1
		//
		lock.Unlock()
	}
}
func sub(){
	defer wg.Done()
	//lock.Lock()
	for i :=0;i < 100000;i++{
        lock.Lock()	
		t = t -1
		lock.Unlock()
	}
	
}
func main(){
	wg.Add(2)
	//启动协程
	go add()
	go sub()
	wg.Wait()
	fmt.Println(t)
}