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
}