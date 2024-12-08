package main
import (
	"fmt"
)
func main(){
	var intChan chan int
	intChan = make(chan int,100)
	for i :=10; i<100;i++{
		//t := i+1
		intChan<- i+1
	}
	//关闭管道
	close(intChan)
	//遍历
	for v := range intChan {
        fmt.Println(v)
	}
}