package main
import (
	"fmt"
	"math"
)
func main() {
	// 输出各数值范围
	fmt.Println("int8 range:",math.MinInt8,math.MaxInt8)
	fmt.Println("int16 range:",math.MinInt16,math.MaxInt16)
	fmt.Println("int8 range:",math.MinInt32,math.MaxInt32)
	fmt.Println("int8 range:",math.MinInt64,math.MaxInt64)

	//初始化一个32位整型值
	var a int32 = 1047483647
	//输出变量的十六进制形式和十进制值
	fmt.Printf("int32:0x%x %d\n",a,a)
	// 将a变量数值转换为十六进制，发生数值截断
	b:=int16(a)
	//输出变量的十六进制形式和十进制值
	fmt.Printf("int16:0x%x %d\n",b,b)
	//将常量保存为float32类型
	var c float32 = math.Pi
	//转换为int类型，浮点发生精度丢失
	fmt.Println(c)
	fmt.Println(int(c))
}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo15\main.go
// int8 range: -128 127
// int16 range: -32768 32767
// int8 range: -2147483648 2147483647
// int8 range: -9223372036854775808 9223372036854775807
// int32:0x3e6f54ff 1047483647
// int16:0x54ff 21759
// 3.1415927
// 3