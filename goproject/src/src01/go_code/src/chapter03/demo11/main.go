package main

import (
	"fmt"
	"strconv"
)

 func main(){
	var a1 int = 76
	var a2 float64 = 3.14
	var a3 bool = true
	var a4 byte = 'h'
	var str string
	//使用第一种方式
	str = fmt.Sprintf("%d",a1)
	fmt.Printf("str type %T str=%q\n",str,str)
	str = fmt.Sprintf("%f",a2)
	fmt.Printf("str type %T str=%q\n",str,str)
	str = fmt.Sprintf("%t",a3)
	fmt.Printf("str type %T str=%v\n",str,str)
	str = fmt.Sprintf("%q",a4)//%q
	fmt.Printf("str type %T str=%v\n",str,str)
	//使用第二种方式
	var a5 int = 76
	var a6 float64 = 3.14
	var a7 bool = true
	var s1 string
	s1 = strconv.FormatInt(int64(a5),10)
	fmt.Printf("s1 type %T s1=%q\n",s1,s1)
	//注解：'f' 格式 10: 表示小数位保留10位，64：表示这个小数是float64位
	s1 = strconv.FormatFloat(a6,'f',10,64)
	fmt.Printf("s1 type %T s1=%q\n",s1,s1)
	s1 = strconv.FormatBool(a7)
	fmt.Printf("s1 type %T s1=%q\n",s1,s1)
	s1 = strconv.Itoa(a5)
	fmt.Printf("s1 type %T s1=%q\n",s1,s1)
}
// 执行结果
// str type string str=76
// str type string str=3.140000
// str type string str=true
// str type string str='h'
// s1 type string s1="76"
// s1 type string s1="3.1400000000"
// s1 type string s1="true"
// s1 type string s1="76"