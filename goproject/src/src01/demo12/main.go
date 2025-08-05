package main
import (
	"fmt"
	"strconv"
)
func main() {
	var n1 int = 10
	var s1 string = strconv.FormatInt(int64(n1),10)//参数。第一个参数必须转为int64类型，第二个参数指定字面值得进制形式为10 进制
	fmt.Printf("s1对应类型是：%T, s1 = %v\n",s1 ,s1)
	var n2 float32 = 4.78
	var s2 string = strconv.FormatFloat(float64(n2),'f',9,32)//'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。控制精度（排除指数部分）;32表示给的小数类型 
	fmt.Printf("s2对应类型是：%T, s2 = %v\n",s2 ,s2)
	var n3 bool = false
	var s3 string = strconv.FormatBool(n3)
	fmt.Printf("s3对应类型是：%T, s3 = %v\n",s3 ,s3)
//	var n4 byte = 'a'

}