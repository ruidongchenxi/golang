package main

import (
	"fmt"
	//"slices"
)

//如果结构体的字段是类型是：指针、slice、和map的零值都是nil,即没有分配空间
//如果需要使用这样字段需要先make
type Person struct{
	Name string
	Age  int 
	Scores [5]float64
	ptr *int 
	slice []int 
	map1 map[string]string 
}
type Person1 struct{
	Name string
	age int 
}
func main(){

	// var p1 Person
	// fmt.Println(p1)
	// if p1.ptr == nil{
	// 	fmt.Println("为空")
	// }
	// p1.slice= make([]int,10)
	// p1.slice[0]=67
	// fmt.Println(p1)
	// p1.map1 = make(map[string]string)
	// p1.map1["你好"]="ty"
	// fmt.Println(p1)
	// var c1 Person1
	// c1.age= 10
	// c1.Name= "红孩儿"
	// fmt.Println(c1)
	// c2 := c1
	// fmt.Println(c2)
	// c2.Name = "哪吒"
	// fmt.Println(c2)
	// var p2 Person1 = Person1{}
	// fmt.Println(p2)
	var p3 *Person1 = new(Person1)
	p3.Name= "cd"
	(*p3).age= 56
	fmt.Println("p3赋值后",*p3)
	var p4 *Person1= p3
	p4.age= 98 //
	fmt.Println("p4 改值后p3",*p3)
	fmt.Println("p4的值",*p4)
// 	var p4 *Person1 = &Person1{}
// 	fmt.Println(*p4)
}
//执行结果
// p3赋值后 {cd 56}
// p4 改值后p3 {cd 98}
// p4的值 {cd 98}