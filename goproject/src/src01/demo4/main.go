package main
import "fmt"
func main(){
	//定义浮点类型
	var num1 float32 = 3.14
	fmt.Println(num1)
	//正数负数都可以
	var num2 float32 = -3.14
	fmt.Println(num2)
	//浮点数可以用十进制表示，也可以用科学计数法表示，E大小写都可以
	var num3 float32 = 314E-2
	fmt.Println(num3)
	var num4 float32 = 314E+2
	fmt.Println(num4)
	var num5 float32 = 314e+2
	fmt.Println(num5)
	var num6 float32 = 314e+2
	fmt.Println(num6)
	//浮点数可能会有精度损失，建议使用float64
	var num7 float32 = 256.000000916
	fmt.Println(num7)
	var num8 float64 = 256.000000916
	fmt.Println(num8)
	//golang里默认的浮点类型为float64
	cx := 3.17
	fmt.Printf("cx 对应的数据类型:%T",cx)
/* 	3.14
	-3.14
	3.14
	31400
	31400
	31400
	256
	256.000000916
	cx 对应的数据类型:float64 */
}