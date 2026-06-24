package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
//需求：可以主动监控
func cpuInfo(ctx context.Context) {
	//函数里拿到链路ID
	 fmt.Printf("trancid:%s\r\n",ctx.Value("traceid"))
	//记录日志
	
	defer wg.Done()
	for{
		select {
		case <- ctx.Done():  //自动超时可以省略
			fmt.Println("退出监控")
			return
		default:
			time.Sleep(2*time.Second)
			fmt.Println("cpu信息")
		}
		// if stop {
		// 	break
		// }
		
	}

}
func main() {
   //var stop = make(chan struct{})
   wg.Add(1)
/* context 包提供了3中函数，WithCancel,WithTimeout,WithValue:返回不同类型的结构体
    如果goroutine 函数中，如果希望被控制，超时消息，传值，但是不希望影响原来的接口信息的时候，函数参数中第一个参数要尽量加上一个ctx，
*/
//    ctx1,cancel1:=context.WithCancel(context.Background())
//    ctx2,_:=context.WithCancel(ctx1)
//方式2：超时自动退出
	//ctx,_:=context.WithTimeout(context.Background(),6*time.Second)
 
   //方式3：WithDeadline 指定退出时间
    ctx1:=context.WithValue(context.Background(),"traceid","gjw12")

   //方式4：WithValue
   go cpuInfo(ctx1)
   wg.Wait()
   fmt.Println("监控完成")
}