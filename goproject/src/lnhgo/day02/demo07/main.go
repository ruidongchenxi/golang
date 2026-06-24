package main

import (
	"container/list"
	"fmt"
)
func main(){
	var mylist list.List//初始化1
	//初始化2操作
//mylist := list.New()
	mylist.PushBack("go")
	mylist.PushBack("java")
	mylist.PushBack("PHA")
	//fmt.Println(mylist)
	//正向遍历列表
	for i:=mylist.Front();i!=nil;i=i.Next(){
		fmt.Println(i.Value)
	}
	//反向遍历
	for i:=mylist.Back();i!=nil;i=i.Prev(){
		fmt.Println(i.Value)
	}
	//方式2 初始化列表
	t:=list.New()
	t.PushBack("t")
	t.PushBack("grpc")
	//头部插入
	t.PushFront("JAVA")
	
	fmt.Println("遍历T列表")
	for i:=t.Front();i!=nil;i=i.Next(){
		fmt.Println(i.Value)
	}
	//插入元素
	s := t.Front()
	for; s!=nil;s=s.Next(){
		if s.Value.(string)== "grpc" {
			break
		}
	}
	t.InsertBefore("python",s)
	fmt.Println("-----------------------")
	for i:=t.Front();i!=nil;i=i.Next(){
		fmt.Println(i.Value)
	}
	fmt.Println("-------------------")
	//删除元素
	for; s!=nil;s=s.Next(){
		if s.Value.(string)== "grpc" {
			break
		}
	}
	t.Remove(s)
	for i:=t.Front();i!=nil;i=i.Next(){
		fmt.Println(i.Value)
	}
//数组 切片 map list：用的少

}