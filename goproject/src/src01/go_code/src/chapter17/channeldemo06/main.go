package main

import (
	"fmt"
	// "math/rand"
	// "time"
)
// type Cat struct{
// 	Name string
// 	Age int
// 	Address string
// }
func main(){
		intChan := make(chan int,3)
		intChan <- 100
		intChan <- 200
		close(intChan) //关闭管道；不能写入操作

		//intChan <- 300
				/*写入操作报错
				panic: send on closed channel

goroutine 1 [running]:
main.main()
        D:/golang/goproject/src/src01/go_code/src/chapter17/channeldemo06/main.go:18 +0x65
exit status 2*/
		//fmt.Println(intChan)
		fmt.Println(<-intChan)//还是可以正常读的
		intChan2 := make(chan int,100)
		for i:=0;i<100;i++{
			intChan2 <- i*2
		}
		close(intChan2)//如果没有这一行关闭管道操作则报all goroutines are asleep - deadlock!
		//遍历管道时只能使用for-range
		for v:=range intChan2 {
			fmt.Println("v",v)
		}


}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo06\main.go
// v 0
// ...
// v 198 