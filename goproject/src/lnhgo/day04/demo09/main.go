package main
import (
	"fmt"
	"time"
)
func main(){
	var msg chan int
	msg = make(chan int,2)
	go func (msg chan int)  {
		for data:=range msg{
			fmt.Println(data)
		}
		fmt.Println("all done")

	}(msg)
	msg <-1
	msg <-2
	close(msg)//关闭管道
	d:=<-msg//关闭通道可以取数据;不能放值
	fmt.Println(d)
	//msg <-3 已经关闭的管道无法写入数据
	time.Sleep(time.Second*3)
	
} 