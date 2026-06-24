package main

import (
	"fmt"
	//"net"
)
type Person struct{
	name string
	age int
}
func main(){
	/*
	不同类型的数据零值不一样
	bool false
	数值  0
	string ""
	pointer nil
	slice nil
	map nil
	channel、interface、function nil
	struct 默认值不是nil、默认值是具体字段的默认值
	*/
	p1 := Person{
		name: "美羊羊",
		age: 18,

	}
	p2 := Person{
		name: "美羊羊",
		age: 18,
	}
	if p1==p2{
		fmt.Println("yes")
	}
	//切片默认值
	var ps []Person
	if ps == nil{
		fmt.Println("ps为空")
	}
	var ps1 = make([]Person,0)//初始化默认值就不是nil了
	if ps1 == nil{
		fmt.Println("ps1为空")
	}else {
		fmt.Println("ps1不为空")
	}
	var m map[string]string//没有初始化为空
	if m == nil{
		fmt.Println("m为nil")
	}else {
		fmt.Println("m不为空")
	}
	var m1 = make(map[string]string,0)//初始化就不为空了
	if m1 == nil{
		fmt.Println("m1为nil")
	}else {
		fmt.Println("m1不为空")
	}	

	//有些类型默认值不可以和int类型比较
	// var a int
	// if a != nil{
	// 	fmt.Println()
	// }

}