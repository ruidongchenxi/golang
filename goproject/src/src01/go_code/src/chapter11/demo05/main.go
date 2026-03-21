package main
import (
	"fmt"
	// "math/rand"
	// "sort"
)
type Englishs interface{
	English()
}
type Athlete struct{
	Name string
}
type Basketball struct{
	Athlete
}
type Booter struct{
	Athlete
}
func (r *Booter)English(){
	fmt.Printf("足球运动员%v会使用英语沟通\n",r.Name)
}
// type Student struct{
type Student struct{
	Name string
	Age int
}
type Middlestudent struct{
	Student
}
type Collegestudents struct{
	Student
}
func (x *Collegestudents)English(){
	fmt.Printf("大学生%v会使用英语沟通\n",x.Name)
}
// type Athlete struct{
// 	Name string
// 	Height int
// }
// func (r *Athlete)Booter(){//足球
// 	if r.Height<170{
// 		fmt.Printf("%v 是足球运动员",r.Name)
// 	}else{
// 		fmt.Printf("%v 不是足球运动员",r.Name)
// 	}
// }
// func (r *Athlete)Basketball(){
// 	if r.Height< 170{
// 		fmt.Printf("%v 是篮球运动员",r.Name)
// 	}else{
// 		fmt.Printf("%v 不是篮球运动员",r.Name)
// 	}
// }
// func (r *Athlete)English(){
// 	fmt.Printf("运动员%v会使用英语沟通\n",r.Name)
// }
// type Student struct{
// 	Name string
// 	Age int
// }
// func (x *Student)English(){
// 	if x.Age>=18{
// 		fmt.Printf("大学生%v会使用英语沟通\n",(*x).Name)
// 	}else{
// 		fmt.Printf("中学生%v不会使用英语沟\n",x.Name)
// 	}
// }

// 执行结过
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter11\demo04\main.go
// 爬树
// 悟空 学会鸟的飞行
// 悟空 学会了游泳