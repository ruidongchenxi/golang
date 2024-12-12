package main

import (
	"fmt"
)
func GetData() (int,int){
	return 100, 200
}
func main()  {
	//标准格式 var 变量名 类型= 表达式
 	var chenxi int = 5
 	// chenxi 为变量名字，类型是int，chenxi的初始值是5
 	//编译器类型推导
 	var cx = 7
 	//等号右边部分在编译器原理里被称为“右值”
	var attack = 32
	var defence = 20
	var damageRate float32 = 3.14//默认情况，如果不指定damageRate类型，编译器默认推倒为float64.由于这个例子中不需要float64的精度，所以强制指定为float32位
	//将attack和defence相减后的数值结果依旧是整型，使用float32()将结果转换为float32类型。再与float32类型的damageRate相乘。damage类型也是float32
	var damage = float32(attack-defence) * damageRate
	fmt.Println("指",damage)
	fmt.Println(cx,chenxi)
	//oon,err := net.Dial("tcp","127.0.0.1:8000")
	//fmt.Println(coon,err)
     a,_ := GetData()
     _,d := GetData()
     fmt.Println(a,d)
}
