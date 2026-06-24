package main
import (
	"fmt"
)
type Person struct{
	name string
}
func changName(a *Person){
	a.name="慕课网"
}
func main(){
	// p:=Person{
	// 	name: "booy",
	// }
	// var pi *Person
	// pi =&p 
	// changName(pi)
	// fmt.Printf("%p\n",pi)
	// fmt.Println(pi.name)
	// po:=&Person{
	// 	"小杜",
	// }
	//po=&p
	//第一个不同的点就出来了,直接使用变量一样使用指针，第二个点go语言限制了指针的运算，在C语言里你可以拿到一个指针进行加1，go语言中不支持，不能参加运算
	
	/*
	go的指针是阉割版，unsafe包里面，所以一般不会使用unsafe包，但是可以使用
	*/
	// fmt.Println(po.name)
	//var a = 10
	//定义指针
	//b:=&a
	//var p1 *Person//结构体类型放入指针没有初始化不可以调用结构体字段panic: runtime error: invalid memory address or nil pointer dereference；需要初始化
	var p1 Person//结构体为赋值可以调用字段
	fmt.Println(p1.name)
	//结构体类型指针初始化方法1 
	ps := &Person{}
	//结构体类型指针初始化方法2
	var empt Person
	ps2 :=&empt
	//结构体类型指针初始化3
	//初始化两个关键字，map、channel、slice 初始化推荐make；指针初始化new函数，指针要初始化否则会出现nil pointer
	//map 必须使用make初始化
	var ps3 = new(Person)
	fmt.Println(ps.name)
	fmt.Println(ps2.name)
	fmt.Println(ps3.name)


	
}