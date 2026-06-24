package main
import (
	"fmt"
)
func main(){
	//常量：定义的时候指定值，不能修改，常量全部大写
	const PI float32 = 3.1415926//显示定义
	// 定义一组
	const (
		UNKNOWN = 1
		FMMALE=2
		MALE= 3
	)
	const(
		X int = 6
		C //如果常量没有定义值，那会沿用它场面常量的值
		S  = "sa"
		D 
	)
	fmt.Println(X,C,S,D)
	/*
	常量类型可以定义为bool、数值、和字符串
	不曾使用的常量没有强制使用的要求
	显示指定类型的时候，必须保证常量左右值类型一致
	如果常量没有定义值，那会沿用它场面常量的值

	*/

}