package main
import (
	"fmt"
	"container/list"
)
func main(){
	l := list.New()
	//尾部添加
	l.PushBack("canon")
	//头部添加
	l.PushFront(67)
	//尾部添加后保存元素句柄
	element := l.PushBack("fist")
	//在fist之后添加high
	l.InsertAfter("high",element)
	//在fist之前添加noon
	l.InsertBefore("noon",element)
	for i := l.Front(); i != nil;i=i.Next(){
		fmt.Println(i.Value)
	}
	fmt.Println("************************************************")
	//使用
	l.Remove(element)
	for i := l.Front(); i != nil;i=i.Next(){
		fmt.Println(i.Value)
	}

}
/*执行后结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\列表\main1.go
67
canon
noon
fist
high
************************************************
67
canon
noon
high

*/