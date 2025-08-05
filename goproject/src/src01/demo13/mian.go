package main
import (
	"fmt"
	"strconv"
)
func main () {
	//string ----> bool
	var s1 string = "string"
	var b bool
	b,_ = strconv.ParseBool(s1)// 返回两个值
	fmt.Printf("b的类型是：%T,b=%v\n",b,b)
	var s2 string = "98"
	var n1 int64
	n1,_ = strconv.ParseInt(s2,10,64)//参数s2 表示字符串，10表示10进制，64表示转成的int64进制；指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64
	fmt.Printf("n1的类型是：%T,b=%v\n",n1,n1)
	var s3 string = "3.14"
	var n3 float64 
	n3,_=strconv.ParseFloat(s3,64)//参数s3 表示字符串，64表示转为float64 ,指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。
	fmgit t.Printf("n3的类型是：%T,b=%v\n",n3,n3)
}