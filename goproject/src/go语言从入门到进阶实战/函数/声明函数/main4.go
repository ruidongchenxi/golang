package main
import "fmt"
//用于测试值传递效果的结构体
type Date struct{
	complax []int    //测试切片在参数传递中的效果
	instance lnnerData //实例分配的	instance lnnerData 
	ptr *lnnerData //将ptr声明为lnnerData的指针类型
}
//代表各种结构体字段
type lnnerData struct {
	a int
}
/*
将Data声明为结构体类型，结构体是拥有多个字段的复杂结构
complax 为整型切片，切片是一种动态类型，内部以指针存在
instance 成员以lnnerData的类型作为Data的成员
将ptr声明为lnnerData的指针类型
声明一个内嵌的结构lnnerData
*/
func passByValue(inFunc Date) Date {
	//输出参数的成员情况
	fmt.Printf("inFunc value: %+v\n", inFunc)
	//打印inFunc的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc)
	return inFunc
}
/*使用格式化的“%+v” 动词输出in变量的详细结构，以便观察Data结构在传递前后的内部数值变化
打印传入参数inFunc的指针地址，在计算机中，拥有相同地址且类型相同的变量，表示的是同一块内存区域
将传入的变量作为返回值返回，返回的过程发生值复制
*/
func main(){
	//准备传入函数的结构
	in := Date{
		complax: []int{1,2,3},
		instance: lnnerData{
			5,
		},
		ptr: &lnnerData{1},
	}
	//输入结构成员情况
	fmt.Printf("in value: %+v\n", in)
	//输入结构的指针地址
	fmt.Printf("in prt: %p\n",&in)
	//传入结构体返回同类型的结构体
	out := passByValue(in)
	//输出结构成员情况
	fmt.Printf("out value: %+v\n", out)
	//输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)
}
/*
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\声明函数\main4.go
in value: {complax:[1 2 3] instance:{a:5} ptr:0xc00000a0b8}
in prt: 0xc0000200f0
inFunc value: {complax:[1 2 3] instance:{a:5} ptr:0xc00000a0b8}
inFunc ptr: 0xc000020180
out value: {complax:[1 2 3] instance:{a:5} ptr:0xc00000a0b8}
out ptr: 0xc000020150
*/