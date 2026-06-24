package main
import (
	"fmt"
)
type Preson struct{
	name string
	age int
}
//给结构体绑定方法；(s Preson)叫接收器，有两种形态。有值传递和引用传递
//有可能该方法中在修改结构体的值，或是结构体对象很大，数据较大
func (s Preson)Println(){
	fmt.Printf("name:%s,age:%d\n",s.name,s.age)
}
func main(){
	//初始化结构体，
	a:=Preson{"小灰灰",8}
	//调用结构体方法
	a.Println()
}