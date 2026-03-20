package main
import(
	"fmt"
)
//定义接口这个接口的含义是：任何类型，只要实现了 Say() 方法，就属于 Ainterface
type Ainterface interface{
	Say()
}
//定义结构体这一步非常关键：👉 Stu 已经“实现了”接口 Ainterface

type Stu struct{
	Name string
}
//这一步非常关键：👉 Stu 已经“实现了”接口 Ainterface
func (stu Stu) Say(){
	fmt.Println("stu Say(11111)")
}
type integer int
//自定义int 数据类型实现了 已经“实现了”接口 Ainterface
func (i integer) Say(){
	fmt.Println("integer Say()")
}
//自定义切片 数据类型实现了 已经“实现了”接口 Ainterface
type intp []int
func (i intp)Say(){
	fmt.Println("切片实现 Say()")
}
type intspec [3]int
func (i intspec)Say(){
	fmt.Println("数组 Say()")
}
type maps map[int]int
func (i maps)Say(){
	fmt.Println("集合实现Say")
}
type my func()
func (i my)Say(){
	fmt.Println("函数接口的实现")
	i()

}
//使用方法
func main(){
	var stu Stu
	//不是“接口实例化结构体”，而是：👉 把一个“实现了接口的对象”赋值给接口变量
	var a Ainterface =stu
	/*
	可以理解成：接口变量 a = (类型信息 + 数据) = (Stu + stu实例)
	接口底层其实是：type interface {
    				type  // 动态类型（这里是 Stu）
    				value // 实际数据（stu）
					}
	*/
	/*调用时发生了什么？
	Go 会：找到 a 里面的真实类型 👉 Stu 调用 Stu 的 Say()👉 这叫：动态分发（多态）
	*/
	a.Say()
	var t integer
	var T Ainterface = t
	T.Say()
	var c intp 
	var C Ainterface =c
	C.Say()
	var b intspec 
	var B Ainterface =b
	B.Say()
	var d maps 
	var D Ainterface = d
	D.Say()
	var e my = my(func() {fmt.Println("你好e函数")})
	var E Ainterface= e
	E.Say()
}
//执行结构
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter11\damo01\main.go
// stu Say(11111)
// integer Say()
// 切片实现 Say()
// 数组 Say()
// 集合实现Say
// 函数接口的实现
// 你好e函数