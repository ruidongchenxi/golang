package main

import "fmt"
	

func main(){
	var n1 int = 19
	var n2 float32 = 4.78
	var n3 bool = false
	var n4 byte = 'a'
	var s1 string = fmt.Sprintf("%d",n1)// 整数：
/* 
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式：U+1234，等价于"U+%04X" */
	fmt.Printf("s1对应类型是：%T, s1 = %v\n",s1 ,s1)//%v	值的默认格式表示;%T	值的类型的Go语法表示
	var s2 string = fmt.Sprintf("%f",n2)//%f	有小数部分但无指数部分，如123.456
	fmt.Printf("s2对应类型是：%T, s2 = %v\n",s2 ,s2)
	var s16 string = fmt.Sprintf("%f",n2)//%f	有小数部分但无指数部分，如123.456
	fmt.Printf("s16对应类型是：%T, s16 = %q\n",s16 ,s16)
	var s3 string = fmt.Sprintf("%t",n3)//布尔值： %t	单词true或false
	fmt.Printf("s3对应类型是：%T, s3 = %v\n",s3 ,s3)
	var s4 string = fmt.Sprintf("%c",n4)//%c	该值对应的unicode码值
	fmt.Printf("s4对应类型是：%T, s4 = %v\n",s4 ,s4)
}
/*
执行结果
PS D:\golang\goproject\src\src01> go run demo11\main.go
s1对应类型是：string, s1 = 19
s2对应类型是：string, s2 = 4.780000
s16对应类型是：string, s16 = "4.780000"
s3对应类型是：string, s3 = false
s4对应类型是：string, s4 = a
*/