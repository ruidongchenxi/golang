package main

import (
	"fmt"
	"strconv"
	//"main"
)
func main(){
	//浮点数
	a:=5.5
	//转换为int
	b:=int(a)//5
	//Go允许在底层结构相同的两个数据类型之间互转，例如：
	//IT类型的底层是int类型
	type IT int
	var a3 IT = 5
	//将a(IT)转换为int，b现在是int类型
	c := int(a)
	var a1 int32 = 1
	var a2 int64 = 1
	b1 := int64(a1)
	fmt.Println(a,a1,a2,a3,b,b1,c)
	//字符串转数字
	var istr="12"
	t,_:=strconv.Atoi(istr)//有可能错误
	fmt.Println(t)
	var my= 32
	d:=strconv.Itoa(my)
	fmt.Println(d)
	//字符串转flaot32类型。字符串转bool类型
	f1,_:=strconv.ParseFloat("3.14567",32)//参数里填写64或32返回的值都是float64，参数作用是按照64或32为格式转换；返回两个值
	fmt.Println(f1)
	//字符串转数字方式2
	i1,_:=strconv.ParseInt("-45",10,64)//参数字符串:指定转换的进制（这里指定为10进制）；转为多少位
	fmt.Println(i1)
	i2,_:=strconv.ParseInt("54",8,16)
	fmt.Println(i2)
	//字符串转布尔;ParseBool 函数返回字符串所表示的布尔值。它接受的值包括 1、t、T、TRUE、true、True、0、f、F、FALSE、false 和 False。任何其他值都会返回错误。
	b2,err := strconv.ParseBool("t")
	if err !=nil{
		fmt.Println("转换失败")
		return
	}
	fmt.Println(b2)
	//基本数据类型转字符串
	b3:=strconv.FormatBool(true)
	fmt.Println(b3)
	f2:=strconv.FormatFloat(3.1415926,'E',-1,64)
	fmt.Println(f2)
	i3:=strconv.FormatInt(-42,16)//16表示16进制
	fmt.Println(i3)


}