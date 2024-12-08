package main
import (
	"fmt"
	"time"
)
func main(){
	//定义一个int 的管道
	intChan := make(chan int,1)
	//定义一个string管道
	stringChan := make(chan string,1)
	go func(){
		time.Sleep(time.Second * 7)
		intChan<- 10
	}()
	go func(){
		time.Sleep(time.Second * 2)
		stringChan<- "尘曦"
	}()

	//fmt.Println(<-intChan)//本身取数据就是阻塞
	select{
		case v := <-intChan :
		    fmt.Println("intChan",v)
		case v := <-stringChan :
			fmt.Println("stringChan",v)
		//default:
		//	fmt.Println("防止阻塞")
	}
}  