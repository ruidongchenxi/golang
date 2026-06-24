package main
import (
	"fmt"
)
type Person struct{
	name string
}
//接受者
func (p Person)SayHello(){

}
//通过指针交换2个值
func swap(a,b *int) {
	//a,b=b,a//这样只是改变函数里ab变量的值，而不是变量值(地址值)指向的底层数据空间的值
	*a,*b=*b,*a //交换两个值,*改底层的值
}
func main(){
	a:=10
	b:=20
	fmt.Println(a,b)
	swap(&a,&b)
	fmt.Println(a,b)
}