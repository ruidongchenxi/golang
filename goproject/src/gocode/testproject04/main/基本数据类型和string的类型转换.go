package main
import "fmt"
//import "strconv"
func main(){
	var n1 int = 19
	var n2 float32 = 4.78
	var n3 bool = false
	var n4 byte = 'a'
	var s1 string = fmt.Sprintf("%d",n1)
	fmt.Printf("s1 对应的数据类型是：%T,s1=%q",s1,s1)
	fmt.Println()
	var s2 string = fmt.Sprintf("%f",n2)
	fmt.Printf("s2 对应的数据类型是：%T,s2=%q",s2,s2)
	fmt.Println()
	var s3 string = fmt.Sprintf("%t",n3)
	fmt.Printf("s3 对应的数据类型是：%T,s3=%q",s3,s3)
	fmt.Println()
	var s4 string = fmt.Sprintf("%c",n4)
	fmt.Printf("s3 对应的数据类型是：%T,s3=%q",s4,s4)


}