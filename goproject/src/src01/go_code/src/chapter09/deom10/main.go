package main

import (
	"fmt"
	//"internal/syscall/windows"
	//"strconv"
)
type Dog struct{
	Name string 
	Age int 
	Weighi string
}
func (t Dog) say()  {
//	d := fmt.Sprintf("Dog的信息 Name=[%v],Age=[%v],Weighi=[%v]",t.Name,t.Age,t.Weighi)
	fmt.Printf("Dog的信息 Name=[%v],Age=[%v],Weighi=[%v]\n",t.Name,t.Age,t.Weighi)
}
type Box struct{
	c float64
	k float64
	g float64
}
func (b Box) test(){
	fmt.Println("输入长")
	fmt.Scanln(&b.c)
	fmt.Println("输入宽")
	fmt.Scanln(&b.k)
	fmt.Println("输入高")
	fmt.Scanln(&b.g)
	fmt.Println("体积是",b.c*b.g*b.k)
}
type Visitor struct{
	Name string
	age int
}
func(v *Visitor)test(){
	for {
		fmt.Println("请输入名字")
		fmt.Scanln(&v.Name)
		if v.Name== "n"{
			break
		}
		fmt.Println("请输入年龄")
		fmt.Scanln(&v.age)
		if v.age > 18{
			fmt.Printf("%v的年龄是%v,门票是20元\n",v.Name,v.age)
		}else{
			fmt.Printf("%v的年龄是%v,门票免费\n",v.Name,v.age)
		}
		

	}
} 
func main(){
	t:=Dog{"小黑",5,"骨头"}
	t.say()
	// var F Box
	// F.test()
	var T Visitor
	T.test()

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter09\deom10\main.go
// 请输入名字
// 你好
// 请输入年龄
// 45
// 你好的年龄是45,门票是20元
// 请输入名字
// 小光
// 请输入年龄
// 6
// 小光的年龄是6,门票免费
// 请输入名字
// n