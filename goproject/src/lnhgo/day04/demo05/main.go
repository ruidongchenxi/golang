package main

import (
	"fmt"
	"sync"
)

//	""
func main(){
	//WaitGroup 主要用于goroutine的执行，ADD方法要与Done方法配套使用
	var r sync.WaitGroup
	r.Add(100)
	for i :=0;i<100;i++{
		go func (i int )  {
			defer r.Done()
			fmt.Println(i)
			//r.Done()
		}(i)
	}
	r.Wait()
}

