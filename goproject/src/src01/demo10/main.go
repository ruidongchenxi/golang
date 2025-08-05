package main
import "fmt"
func main(){
	//类型转换
	var n1 int = 100
	//var n2 float32 = n1//错误写法 
	//报错
	//# command-line-arguments
    //demo10\main.go:6:19: cannot use n1 (variable of type int) as float32 value in variable declarat
	var n2 float32 = float32(n1) //正确写法
	//执行结果
	//PS D:\golang\goproject\src\src01> go run demo10\main.go
//100
	fmt.Println(n2)
	var n3 int64 = 888888
	fmt.Println(n3)
	var n4 int8 = int8(n3)//int64转为int8的时候编译不会出错的，但是会有数据的溢出
	fmt.Println(n4)// 56
	var n5 int32 = 12
	//var n6 int64 = n5+30 //错误写法 = 左右数据类型
	var n6 int64 = int64(n5)+30
	fmt.Println(n6)
	var n7 int64 = 12
	var n8 int8 = int8(n7) + 127 //编译通过，结果可能会溢出；-117
	//var n9 int8 = int8(n7) + 128 //编译不过溢出；demo10\main.go:25:27: 128 (untyped int constant) overflows int8
	
	fmt.Println(n8)
	//fmt.Println(n9)
}