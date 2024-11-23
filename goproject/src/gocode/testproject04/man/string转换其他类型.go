package main
import (
	"fmt"
	"strconv"
)
func main(){
    var s1 string = "true"
	var b bool 
	//b = strconv.ParseBool(s1)
	//strconv.ParseBool 函数方法返回值有两个：func ParseBool(str string) (value bool, err error)
	//value就是我们返回值的布尔数据。err出现错误
	//我们只关注得到的布尔类型的数据，eer可以用_直接忽略
	b , _= strconv.ParseBool(s1)
	fmt.Printf("b的类型是：%T，b=%v \n",b,b)
	var s2 string = "19"
	var n1 int64 
	n1, _ = strconv.ParseInt(s2,10,64)
	/*
	func ParseInt(s string, base int, bitSize int) (i int64, err error)
返回字符串表示的整数值，接受正负号。
base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
	*/
	fmt.Printf("n1的类型：%T,n1=%v\n",n1,n1)
	var s3 string = "3.14"
	/*
	cannot use strconv.ParseFloat(s3, 32) (value of type float64) as float32 value in assignment  

	*/
	var f1 float64
	f1, _ = strconv.ParseFloat(s3,32)// 总是返回float64
	fmt.Printf("f1的类型：%T,f1=%v \n",f1,f1)
} 