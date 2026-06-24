package main

import "fmt"

func main() {
	/*
		不要通过共享内存来通信，而是要通过通信来实现内存共享
		channel
	*/
	var msg chan string
	/*
	无缓冲channel适用于通知，B要第一时间知道A是否已经完成
	有缓冲channel适用于消费者和生产者之间通信
	使用场景
	1.消息传递
	2.信号广播
	3.事件订阅和广播
	4.任务分发
	5.结果汇总
	6. 并发控制
	7.同步和异步

	*/
	//msg = make(chan string, 1)//channel 的初始化值，如果为0的话，放值进去会阻塞；有缓冲无缓冲；设置为0表示无缓冲管道；只要大于0；就是有缓冲
	msg = make(chan string,0)// 无缓存
	//
	go func (msg chan string)  {// go有一种happen-before机制，可以保障
		data:=<-msg
		fmt.Println(data)
	}(msg)
	msg <- "bobby" //放值到channel里
	// data := <-msg  //从channel里取值
	// fmt.Println(data)

}