package main
import (
	"errors"
	"fmt"
)
//使用一个数据结构
type Queue struct{
	maxSize int 
	array [5]int
	front int //表示指向数组前面
	rear int //表示指向队列的尾部
}
func (this *Queue)Add(val int) (err error){
	//先判断队列是否已满
	if this.rear > this.maxSize-1{ //重要提示；rear是队列尾部含最后一个元素
		return errors.New("队列已满无法添加")
	}
	this.rear++
	this.array[this.rear]=val
	return err 
}
//显示队列
func (this *Queue)Show(){
  //找到队首遍历到
  for i:= this.front+1 ;i <= this.rear;i++{
	fmt.Println(this.array[i])
  }
}
func (this *Queue)GetQueue() (val int,err error){
	if this.rear == this.front{
		return -1 , errors.New("queue ")

	}
	this.front++
	val= this.array[this.front]
	return val ,err
}
func main(){
	queue := &Queue{
		maxSize: 5,
		front: -1,
		rear: -1,
	}
	// var key string
	// var val int
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
			err:=queue.Add(val)
			if err != nil{
				fmt.Println(err.Error())
			}else{
				
				fmt.Println("加入队列ok ")
			}
		case "get":
			val,err:=queue.GetQueue()
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println(val)
			}
		case "show":
			queue.Show()
		}
	}

}

