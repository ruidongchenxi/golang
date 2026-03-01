package main

import (
	"fmt"
	//"src01/demo16/test"
)

func main(){
	//声明变量
	var name string
	var age byte
	var sal float32
	var isPass bool
	
	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年零")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水")
	// fmt.Scanln(&sal)
	// fmt.Println("请输入是否通过考试(true/false)")
	// fmt.Scanln(&isPass)
	// fmt.Printf("名字是%v \n年龄是%v \n薪水是%v \n是否通过考试%v \n",name,age,sal,isPass)

	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，输入时用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	 fmt.Printf("名字是%v \n年龄是%v \n薪水是%v \n是否通过考试%v \n",name,age,sal,isPass)


}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter03\demo25\main.go
// 请输入你的姓名，年龄，薪水，是否通过考试，输入时用空格隔开
// 晨曦 19 1900 true
// 名字是晨曦 
// 年龄是19
// 薪水是1900
// 是否通过考试true

