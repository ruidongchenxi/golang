package main

import (
	"errors"
	"fmt"
	//"testing/quick"
)

//最左边下标left
//右边下标right
type Stack struct{
	MaxTop int //表示我们栈最大可以存放个数
	Top int //表示栈顶
	arr [5]int
}
func (this *Stack)push(val int) (err error){
	// 
	if this.Top==this.MaxTop-1{
		fmt.Printf("栈满")
		return errors.New("栈满溢出")
	}
	this.Top++
	this.arr[this.Top]=val
	return

}
func (this *Stack)Pop() (val int,err error){
	if this.Top==-1{
		fmt.Println("空栈")
		//val = -2
		return -2, errors.New("栈空")
	}
	val = this.arr[this.Top]
	this.Top--
	return val,nil
}
//遍历栈时需要从栈顶开始遍历
func (this *Stack)list(){
	//判断栈是否为空
	if this.Top==-1{
		fmt.Println("是一个空栈")
	}
//	curTop = this.Top
	for i:=this.Top;i>=0;i--{
		fmt.Printf("arr[%d]=%d\n",i,this.arr[i])
	}
}
func main(){
	stack:=&Stack{
		MaxTop: 5,
		Top: -1,
	}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)
	stack.list()
	val,_:=stack.Pop()
	fmt.Println("出栈val=",val)
	stack.list()
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter20\demo11\main.go
// arr[4]=5
// arr[3]=4
// arr[2]=3
// arr[1]=2
// arr[0]=1
// 出栈val= 5
// arr[3]=4
// arr[2]=3
// arr[1]=2
// arr[0]=1