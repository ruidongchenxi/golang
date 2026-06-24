package main

import (
	"fmt"
	"sync"
	"time"
)

/*
锁本质上是将并行的代码串型化了，使用lock肯定会影响性能
即使设计锁也要尽量保证并行
有两组协程
一组负责写，一组负责读
读协程之间应该允许并发、读写之间应该串行
*/
func main(){
	var num int 
	var rwlock sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func ()  {
		defer wg.Done()
	
			rwlock.Lock()
			defer rwlock.Unlock()
			num =12
		
	}()
	time.Sleep(time.Second)
	go func ()  {
		defer wg.Done()
		// for i:=0;i<1000000;i++{
			rwlock.RLock()
			defer rwlock.RUnlock()
			fmt.Println(num)
		//}
	}()
	wg.Wait()
}