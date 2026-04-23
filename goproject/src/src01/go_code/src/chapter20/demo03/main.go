package main

import (
	"errors"
	"fmt"
	//"hash"
	"os"
)

//使用一个数据结构
type Queue struct{
	maxSize int 
	array [5]int
	// front int //表示指向数组前面
	// rear int //表示指向队列的尾部
	head int //队首
	tail int //队尾
}
func (t *Queue)Push(val int)(err error){
	if t.IsFull(){
		return errors.New("队列已满")
	}
	//分析出t.tail 在队列尾部，但是不含最后元素
	t.array[t.tail]= val
	t.tail++
	return 

}
func (t *Queue)Pop()(val int,err error){
	if t.IsEmpty(){
		return 0,errors.New("队列已空")
	}
	
	val = t.array[t.head]
	t.head++
	return 

}
//判断
func (t *Queue)IsFull() bool{
	return (t.tail+1)%t.maxSize== t.head

	
}
//判断队列为空
func (t *Queue)IsEmpty()bool{
	return  t.tail == t.head
}
//取出队列多少元素
func (t *Queue)Size()int{
	//关键
	return  (t.tail +t.maxSize-t.head)%t.maxSize
}
func (t *Queue)ListQueue(){
	size :=t.Size()
	if size ==0{
		fmt.Println("队列为空")
		return
	}
	tempHead := t.head
	for i:=0;i<size;i++{
		
		fmt.Printf("arr[%d]=%d\t", tempHead,t.array[tempHead])
		tempHead= (tempHead+1)%t.maxSize
	}
	
	
	fmt.Println()
}
func main(){
	queue := &Queue{
		maxSize: 5,
		head: 0,
		tail: 0, 
	}
		for {
		var key string
		var val int
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")
		fmt.Scanln(&key)
		switch key{
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err:=queue.Push(val)
			if err != nil{
				fmt.Println(err.Error())
			}else{
				
				fmt.Println("加入队列ok ")
			}
		case "get":
			val,err:=queue.Pop()
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println(val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter20\demo03\main.go
// 1. 输入add 表示添加数据到队列
// 2. 输入get 表示从队列获取数据
// 3. 输入show 表示显示队列
// 4. 输入exit 表示显示队列
// add
// 输入你要入队列数
// 1
// 加入队列ok 
// 1. 输入add 表示添加数据到队列
// 2. 输入get 表示从队列获取数据
// 3. 输入show 表示显示队列
// 4. 输入exit 表示显示队列
// add
// 输入你要入队列数
// 2
// 加入队列ok 
// 1. 输入add 表示添加数据到队列
// 2. 输入get 表示从队列获取数据
// 3. 输入show 表示显示队列
// 4. 输入exit 表示显示队列
// add
// 输入你要入队列数
// 3
// 加入队列ok 
// 1. 输入add 表示添加数据到队列
// 2. 输入get 表示从队列获取数据
// 3. 输入show 表示显示队列
// 4. 输入exit 表示显示队列
// add
// 输入你要入队列数
// 5
// 加入队列ok 
// 1. 输入add 表示添加数据到队列
// 2. 输入get 表示从队列获取数据
// 3. 输入show 表示显示队列
// 4. 输入exit 表示显示队列
// show
// arr[0]=1        arr[1]=2        arr[2]=3        arr[3]=5
// 1. 输入add 表示添加数据到队列
// 2. 输入get 表示从队列获取数据
// 3. 输入show 表示显示队列
// 4. 输入exit 表示显示队列
// get 
// 1
// 1. 输入add 表示添加数据到队列
// 2. 输入get 表示从队列获取数据
// 3. 输入show 表示显示队列
// 4. 输入exit 表示显示队列
// add 
// 输入你要入队列数
// 4
// 加入队列ok 
// 1. 输入add 表示添加数据到队列
// 2. 输入get 表示从队列获取数据
// 3. 输入show 表示显示队列
// 4. 输入exit 表示显示队列
// add
// 输入你要入队列数
// 6
// 队列已满
// 1. 输入add 表示添加数据到队列
// 2. 输入get 表示从队列获取数据
// 3. 输入show 表示显示队列
// 4. 输入exit 表示显示队列
// show
// arr[1]=2        arr[2]=3        arr[3]=5        arr[4]=4
// 1. 输入add 表示添加数据到队列
// 2. 输入get 表示从队列获取数据
// 3. 输入show 表示显示队列
// 4. 输入exit 表示显示队列
