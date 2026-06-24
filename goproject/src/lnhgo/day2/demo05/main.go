package main

import (
	"fmt"
	//"math"
)
func main(){
	// var (
	// 	a = 8
	// 	b = 5
	// )
	// fmt.Println(a+b)
	// fmt.Println(a-b)
	// fmt.Println(a*b)
	// fmt.Println(a/b)
	// fmt.Println(a%b)
	// // 函数机制go 的幂运算必须是浮点型数字
	// fmt.Println(math.Pow(float64(a),float64(b)))
	// a++//a=a+1
	// b++//b=b+1
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(a==b)
	// fmt.Println(a!=b)
	// fmt.Println(a>=b)
	// fmt.Println(a>b)
	// fmt.Println(a<=b)
	// fmt.Println(a<b)
	// age := 22
	// if age >18 && age< 60{ //两个条件都满足
	// 	fmt.Println("苦逼")
	// }else{
	// 	fmt.Println("不用上班")
	// }
	// if age<18|| age>60{
	// 	fmt.Println("不用上班")
	// }else{
	// 	fmt.Println("苦逼上班的")
	// }
	// //取反
	// t := false
	// fmt.Println(t)//false
	// fmt.Println(!t)//true
	//5 的二进制 101
	//2的二进制10
	//&按位与(两位均为1才为1)

	// fmt.Println(5&2)//000
	// //|按位或(两位有1个为1就为1)
	// fmt.Println(5|2)//1117
	// //^参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。（两位不一样则为1）
	// fmt.Println(5|2)//1117
	// //<<左移n位就是乘以2的n次方。 “a«b"是把a的各二进位全部左移b位，高位丢弃，低位补0
	// fmt.Println(5<<2)//20
	// //>>右移n位就是除以2的n次方。 “a»b"是把a的各二进位全部右移b位。
	// fmt.Println(5>>2)//1
	x :=10
	x +=1 //x=x+1
	x -=1 //x=x-1
	x *=2 //x=x*2
	x /=2 //x=x/2
	x %=2 //x=x%2
	x <<=2 //x =x<<2
	x &=2 //x = x & 2
	x |=3 //x= x| 2
	x ^=4 //x= x^2
	x >>=2 //x=x>>2
	fmt.Println(x)
}
