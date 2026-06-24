package main

import "fmt"

type Preson struct {
	name string
	age  int
}
type Preson1 struct {
	//结构体嵌套
	Preson
	score int
	name string//两个name 调度是默认这个优先级高
}

func main() {
	//嵌套结构体初始化
	s := Preson1{
		Preson{
		"灰太狼",
		20,
		},
		7,
		"小灰灰",
	}
	fmt.Println(s)
	fmt.Println(s.name)//覆盖作用显示小灰灰
	fmt.Println(s.Preson.name)//显示灰太狼

}