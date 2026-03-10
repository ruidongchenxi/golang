package main

import (
	"fmt"
	"strings"
)
func Addupoer () func (int) int{
	var n int = 10
	return func(x int) int {
		n = n+x
		return n
	}
}
func Acc(n int) func ()int{
	return func() int {
		n++
		return n
	}
}
func makeSuffix() func (string) string{
	var f string
	return func(s string) string {
	if strings.HasSuffix(s,".jpg") {
		f = s
		return f
	}else{
		f = s+".jpg"
		return f
	} 
	}

}
func ma(x string) func (string) string{
	return func(s string) string {
		if strings.HasSuffix(s,x){
			return s
		}else {
			return s+x
		}
		
	}
}
func main(){	
	f :=Addupoer()//返回的是函数，f 是变量存储的是函数，他有初始值
	fmt.Println(f(1))//11 初始值是10 f调他加1 也就是11，此时n=11（闭包特性，被引用的自由变量和函数一同存在，及时已经离开了自由变量的环境也不会被释放或删除，在闭包中可以继续使用这个变量）
	fmt.Println(f(2))//13//根据闭包特性f函数只要不消n还是11 f传2 11+2=13；所以返回13
	c := Addupoer()
	fmt.Println("新闭包c(6)返回",c(6))//返回16
	fmt.Println(f(6))
	fmt.Printf("%p\n",f)
	fmt.Printf("%p\n",c)
	// d := Acc(1)
	// c1 := Acc(10)
	// fmt.Printf("%p\n",d)
	// fmt.Printf("%p\n",c1)
	t := makeSuffix()
	c1 := ma(".jpg")
	fmt.Println("文件",t("1.jpg"))
	fmt.Println("文件2",t("2"))
	fmt.Println("文件3",c1("4.jpg"))

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo10\mian.go
// 11
// 13
// 新闭包c(6)返回 16
// 19
// 0xe8b420
// 0xe8b400
// 文件 1.jpg
// 文件2 2.jpg
// 文件3 4.jpg