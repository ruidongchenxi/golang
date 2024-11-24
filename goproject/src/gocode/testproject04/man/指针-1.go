package main
import (
	"fmt"
	//"strconv"
)
func main(){
    var age int = 18
	//&符号+变量就可以获取这个变量内存地址
	fmt.Println(&age)
	//定义一个指针变量
	//var 代表要声明
	//ptr 指针变量的名字
	//ptr 对应的类型：*int 是一个指针类型可以理解为 指向int 类型的指针
	//&age 就是一个地址，是ptr变量的具体值
	//指针就是内存地址
	//& 取内存地址
	//* 根据地址取值
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr 本身这个存储空间得地址为： ",&ptr)
	fmt.Printf("ptr 指向得数值为：%v \n",*ptr)
	fmt.Println("ptr 指向得数值为: ",*ptr)
 } 