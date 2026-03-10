package main

import (
	"fmt"
	"src01/go_code/src/chapter05/demo08/utils" //导入函数
)
var  age int = test()
func test() int{
	fmt.Println("12343")
	return 90
}
func init(){
	age = 90
	//fmt.Println("age",age)
	fmt.Println("init ()() .....")
}
func main(){
	fmt.Println("main() ......")
	fmt.Println("test age",age)
	fmt.Println("age00",utils.Age,"Name",utils.Name)
}
//执行结果
// utils 赋值变量····
// 12343
// init ()() .....
// main() ......
// test age 90
// age00 54 Name tom~