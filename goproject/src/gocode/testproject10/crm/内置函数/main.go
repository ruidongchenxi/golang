package main
import "fmt"
func main(){
	//定义一个字符串
	str := "golang"
	fmt.Println(len(str)) 
	// new 分配内存，new函数的实参是一个类型。而不是具体的值，new 函数返回值是对应类型的指针num：*int
	num := new(int)
	fmt.Printf("num的类型：%T,num的值是： %v,num的地址：%v,num 指针指向的值是：%v \n",num,num,&num,*num)
}