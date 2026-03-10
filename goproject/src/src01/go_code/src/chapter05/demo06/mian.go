package main

import (
	"fmt"
)
func sum(n1 int,n2 int) int{
	return n1+n2
}
//函数既然是一种数据类型，因此在go 中，函数可以作为形参，并且调用
 type myFun func (int,int) int   //这是myFun 就是 func (int,int) int
func mytest(a myFun,num1 int ,num2 int ) int{
	return a(num1,num2)
}
func stest (n1 int,n2 int) (sum int, sub int){// 指定返回
	sum =  n1+n2
	sub = n1-n2
	return
}

func sum1(n1 int,age... int) (sum int){
	sum = n1
    for i :=0; i<len(age);i++{
		sum +=age[i]//表示age[0],表示取出age切片第一个值，其他一次类推
	}
	return
}
func sun2 (n1,n2 *int) {
	n1,n2=n2,n1
}
func main(){	
	a := sum
	fmt.Printf("a的类型%T,sum类型%T\n",a,sum)
	res :=a(10,20)
	fmt.Println(res)
	fmt.Println(a(20,89))
	res2 := mytest(sum,50,60)
	fmt.Println("res2",res2)
	//给int 取别名.在go中myInt 和Int 虽然都是int类型，但是go认为myint 和int 是两个不同的类型
	type myInt int 
	var Num1 myInt 
	Num1 = 40
	fmt.Println("Num1 ",Num1)
	var Num2 int 
	Num2 = int(Num1)//这里需要显示强制转换，因为go认为是不同类型
	fmt.Println("Num2=",Num2)
   
	res3 := mytest(sum,50,56)
	fmt.Println("res3=",res3)	
    res4,res5 := stest(9,7)
	fmt.Printf("两数和%v，两数差%v\n",res4,res5)
	res6 := sum1(1,7,90,99)
	fmt.Println("res6=",res6)
	n6 := 12
	n7 := 56
	fmt.Printf("调用函数前n6=%v,n7=%v\n",n6,n7)
	sun2(&n6,&n7)

	fmt.Printf("调用函数后n6=%v,n7=%v\n",n6,n7)


}
//执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo06\mian.go
// a的类型func(int, int) int,sum类型func(int, int) int
// 30
// 109
// res2 110
// Num1  40
// Num2= 40
// res3= 106
// 两数和16，两数差2
// res6= 197
