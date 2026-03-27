package main
import (
	"fmt"
	//"time"
)

func main(){
	intchan := make(chan int,200)
	numchan := make(chan int,200)
	b := make(chan bool,8)
	go func (intchan chan int){
		for i :=0; i<200;i++{
			intchan <- i
		}
		close(intchan)
	}(intchan)
	r := func (intchan chan int,numchan chan int,b chan bool)   {
		for {
			c := 0
			v,ok := <-intchan
			if !ok{
				b <- true
				break
			}
			for i:=1;i <=v;i++{
				c+=i
			}
			numchan <-c
		}
	}
	for i :=0 ;i<8;i++{
		go r(intchan,numchan,b)
	}
	// for {
	// 	// if t ==8 {
	// 	// 	close(numchan)
	// 	// 	break
	// 	// }

	// }
	for i :=0 ;i<8;i++{
		<-b
	}
	close(numchan)
	for {
		v,ok := <-numchan
		if !ok{
			break
		}
		fmt.Printf("v=%v\n",v)
	}


}

// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo11\main.go
// v=0
// v=1
// v=3
// ...
// v=19306
