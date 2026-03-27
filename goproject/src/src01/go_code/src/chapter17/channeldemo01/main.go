package main

import "fmt"
func main(){
	var intChan chan int //intChan 用于存放int数据类型
	intChan = make(chan int,3)
	// var mapChan chan map[int]string  //(mapChan用于存放map[int]string类型)
	// var perChan chan Person
	// var perChan chan *Person
	fmt.Printf("intChan的值=%v，intChan本身的地址值=%v\n",intChan,&intChan)
	//向管道写入数据
	intChan<-10
	num := 211
	intChan <-num
	//当我们给管道写入数据时，不能超过其容量(这个管道容量是3)
	intChan<-89
	//当超过管道容量时会报错
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan send]:
	//intChan<-67 超出容量无法写入


	//看看管道的长度和cap(容量),
	fmt.Printf("channel len=%v,cap=%v\n",len(intChan),cap(intChan))
	//管道读取数据
	var num2 int
	num2 =<-intChan//取出第一个数据交给
	<-intChan //也可以
	fmt.Println(num2)
	fmt.Printf("channel len=%v,cap=%v\n",len(intChan),cap(intChan))
	//在没有使用协程的情况下 管道数据已经全部取出来，再取就会报告deadlock
	num3 :=<-intChan
	fmt.Println(num3)
	fmt.Printf("channel len=%v,cap=%v\n",len(intChan),cap(intChan))
	num4 :=<-intChan
	fmt.Println(num4)
	fmt.Printf("channel len=%v,cap=%v\n",len(intChan),cap(intChan))
	//管道数据已经全部取出来，再取就会报告deadlock
	num5 :=<-intChan
	fmt.Println(num5)
	fmt.Printf("channel len=%v,cap=%v\n",len(intChan),cap(intChan))
	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter17\channeldemo01\main.go
// intChan的值=0xc00001e100，intChan本身的地址值=0xc000058028
// channel len=3,cap=3
// 10
// channel len=2,cap=3
// 211
// channel len=1,cap=3
// 89
// channel len=0,cap=3
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         D:/golang/goproject/src/src01/go_code/src/chapter17/channeldemo01/main.go:37 +0x593
// exit status 2