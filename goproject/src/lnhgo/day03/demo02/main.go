package main

import (
	"fmt"
	//"structs"
	//"go/types"
)
type Preson struct{
	name string
	age int
	address string
	height float32
}
func main(){
	// //结构体初始化1
	// p1 := Preson{
	// 	name: "喜羊羊",
	// 	age: 12,
	// 	address: "羊村",
	// 	height: 3.24,
	// }
	// //结构体初始化2
	// p2:=Preson{ "美羊羊", 12, "羊村美容课",3.14}
	// fmt.Println(p1)
	// fmt.Println(p2)
	// //结构体初始化3
	// var p3 =Preson{name: "懒羊羊",age: 12,address: "睡觉觉",height: 2.65}
	// fmt.Println(p3)
	// var p4 []Preson
	// p4=append(p4, p1,p2,p3)
	// p5:=[]Preson{
	// 	{
	// 		name: "b4",
	// 	},
	// 	{
	// 		age: 19,
	// 	},

	// }
	// fmt.Println(p5)
	// var p6 Preson
	// p6.name="小灰灰"
	// p6.age=5
	// p6.address= "羊村"
	// p6.height=3.14
	// fmt.Println(p6)
	//匿名函数
	adder:=struct{
		province string
		city string
		address string
	}{
		province: "北京市",
		city: "通州",
		address: "XXX",

	}
	fmt.Println(adder.city)
	
}
/*
执行
通州
*/