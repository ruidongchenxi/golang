package main
import (
	"fmt"
	"strconv"
)
func main(){
	var n1 int = 19
	var s1 string = strconv.FormatInt(int64(n1),10)//第一个参数必须转为int64类型。第二个参数指定字面值的进制形式为十进制
	fmt.Printf("s1 对应的数据类型是：%T,s1=%q \n",s1,s1)
	var n2 float64 = 4.29
	//第二个参数'f'（-ddd.dddd），第三个参数。表示保留小数点后的位数，第四个参数表示float64类型
	var s2 string = strconv.FormatFloat(n2,'f',9,64)
	fmt.Printf("s2 对应的数据类型是：%T,s2=%q \n",s2,s2)
	var n3 bool = true
	var s3 string = strconv.FormatBool(n3)
	fmt.Printf("s3 对应的数据类型是：%T,s3=%q \n",s3,s3)
}