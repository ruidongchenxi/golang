package main
import "fmt"
func main(){
	var n1 int = 100
    //var n2 float32 = n1 在这里自动转换不好使
	fmt.Println(n1)
	var n2 float32 = float32(n1)
	fmt.Println(n2)
	//注意： n1的类型还是int 只是将n1的值100 转为float32 而已，n1还是int类型
	fmt.Printf("%T",n1)
	fmt.Println()
	fmt.Printf("%T",n2)
	fmt.Println()
	var n3 int64 = 88888
	fmt.Println(n3)
	var n4 int8 = int8(n3)//将int64 转为int8的时候，编译不会出错会有精度损失，但是数据会溢出，结果56
	fmt.Println(n4)
	var n5 int32 = 12
	var n6 int64 = int64(n5) + 30// 一定要匹配等号左右的数据类型否则会出错
	//var n6 int64 = n5 + 30
	// # command-line-arguments
	// .\数据类型转换.go:19:17: cannot use n5 + 30 (value of type int32) as int64 value in variable declaration
	fmt.Println()
	fmt.Println(n6)
	var n7 int64 = 12
	var n8 int8 =int8(n7) +127// 编译不会出错会有精度损失，但是数据会溢出，结果-117
	var n9 int8 = int8(n7) + 128/* 编译报错
	# command-line-arguments                                                                                                
    .\数据类型转换.go:27:6: n9 declared and not used                                                                        
     .\数据类型转换.go:27:27: 128 (untyped int constant) overflows int8   
	*/
	fmt.Println(n8)
}