package main

import (
	"fmt"
	"time"
	//"time"
)

//函数
func sayHello(){
	for i :=0; i <10;i ++{
		time.Sleep(time.Second)
		fmt.Println("hello")
	}
}
func test(){
	// 这里使用错误处理机制defer +recover
	defer func ()  {
		//捕获抛出的异常
		if err:= recover();err !=nil{
			fmt.Println("test()协程发生异常了err=",err)
		}
		
	}()
	/* 如果没有这个defer ，就会直接中断程序
	PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo14\main.go
panic: assignment to entry in nil map

goroutine 7 [running]:
main.test()
        D:/golang/goproject/src/src01/go_code/src/chapter17/channeldemo14/main.go:20 +0x25
created by main.main in goroutine 1
        D:/golang/goproject/src/src01/go_code/src/chapter17/channeldemo14/main.go:25 +0x2a
exit status 2
PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo14\main.go
# command-line-arguments
	*/
	//t := make(map[string]string)
	var t map[int]string
	t[1]= "golang"

}
func main(){
	go sayHello()
	go test()
	for i:=0;i <10 ;i++{
		time.Sleep(time.Second)
		fmt.Println("hello main")
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo14\main.go
// test()协程发生异常了err= assignment to entry in nil map
// hello main
// hello
// hello
// hello main
// hello main
// hello
// hello
// hello main
// hello main
// hello
// hello
// hello main
// hello main
// hello
// hello main
// hello
// hello
// hello main
// hello main