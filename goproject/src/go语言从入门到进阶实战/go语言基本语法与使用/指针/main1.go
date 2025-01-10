package main
import (
	"fmt"
)
//交换函数
func swap(a, b *int){
	//取a指针的值，赋值给临时变量t
	t := *a
	*a = *b
	*b = t
}
func main(){
	//准备两个变量，赋值1和2
	x,y := 1,2
	fmt.Println(x,y)
	swap(&x, &y)
	//输出变量的值
	fmt.Println(x,y)
}