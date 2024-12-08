package main
import(
	"fmt"
	"time"
	"sync"
)
var wg sync.WaitGroup
//写
func writData(intChan chan int){
	defer wg.Done()
	for i := 1; i<=50;i++{
		intChan<- i
		fmt.Println("已经写入",i)
	    time.Sleep(time.Second)
	}
	close(intChan)
}
//读
func readData(intChan chan int){
	//遍历
	defer wg.Done()
	for v := range intChan{
		fmt.Println("读取的数据",v)
		time.Sleep(time.Second)
	}
}
func main(){
	wg.Add(2)
	//定义管道
	intChan := make(chan int,50)
	//写协程
    go writData(intChan)
	//读协程
	go readData(intChan)
	wg.Wait()
}
