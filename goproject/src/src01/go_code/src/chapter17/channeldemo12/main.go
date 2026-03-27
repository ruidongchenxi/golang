package main
import (
	"fmt"
	//"time"
)

func main(){
	//管道可以声明只读或者只写
	//1.在默认情况下，管道是双向的
	// var chan1 chan int //可读可写
	//2声明为只写
	var chan1 chan<- int
	chan1 = make(chan int,3)
	chan1 <-1
	//chan1<-2
	//chan1<-3
	//fmt.Println(<-chan1)//invalid operation: cannot receive from send-only channel chan1 (variable of type chan<- int) 
/*报错提示不可以读
invalid operation: cannot receive from send-only channel chan1 (variable of type chan<- int) 
*/
    var chan2 <-chan int
	//chan2 <- 67
	 /* 报错提示不可以写
	invalid operation: cannot send to receive-only channel chan2 (variable of type <-chan int)
	*/
	num :=<- chan2
	 
	fmt.Println("num",num)


}

// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo11\main.go
// v=0
// v=1
// v=3
// ...
// v=19306
