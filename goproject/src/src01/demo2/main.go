package main
import "fmt"
//全局变量 : 定义在函数外
var n8 = 100
var n9 = 123.5
//全局变量一次性申请
var ( 
	n10 = 45
	n11 = "晨曦"
)
func main(){
	//1变量使用方式
	var mun int =18
	fmt.Println(mun)
	//2 变量没有赋初始值，默认整数型默认是0
	var mun2 int 
	fmt.Println(mun2) 
	//类型推导 
	var mun3 = 20
	fmt.Println(mun3)
	// 类型推导 (简写)省略var 。注意:= 不能写成=
	mun4 := 9
	fmt.Println(mun4)
	var n1,n2,n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	var n4,name,n5 = 10,"jak",7.8
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(n5)

    n6,n7 := 18.7,18
	fmt.Println(n6)
	fmt.Println(n7)
	//全局变量的调用
	fmt.Println(n8)
	fmt.Println(n9)
	fmt.Println(n10)
	fmt.Println(n11)


}