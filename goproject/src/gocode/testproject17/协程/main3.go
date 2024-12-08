package main
import (
	"fmt"
    "sync"
	"time"
)
//加入读写锁
var lock sync.RWMutex
//var t int
var wg sync.WaitGroup
//var lock sync.Mutex
func read(){
	defer wg.Done()
	lock.RLock()//如果只是读数据，那么这个锁不产生影响，但是读写同时发生时候，就会有影响
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取成功")
	lock.RUnlock()
}
func write(){
	defer wg.Done()
	lock.Lock()
    fmt.Println("开始写入")
	time.Sleep(time.Second * 10)
	fmt.Println("修改数据成功")
    lock.Unlock()
}
func main(){
	wg.Add(6)
	//启动协程
	for i := 0; i<5 ;i++{
		go read()
	}
	//go read()
	go write()
	wg.Wait()
	//fmt.Println()
}