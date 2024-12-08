package main
import (
	"fmt"
	// "strconv"
    "sync"
)
var wg sync.WaitGroup//指定要不赋值
func main(){
	//匿名函数+ 外部函数= 闭包；启动5个协程
	for i:=1 ;i <=5 ;i++{
	wg.Add(1)// 协程开始加1
	go func() { //匿名函数，启动协程
		defer wg.Done()
		fmt.Println(i)
		//wg.Done() //协程结束执行完成减一
	}()
	}
	//阻塞主线程，\什么时候wg 减为0，就停止阻塞
	wg.Wait()
	
}

