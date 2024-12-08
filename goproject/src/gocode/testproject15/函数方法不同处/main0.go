package main
import "fmt"

//结构体
type Student struct{
	Name
}
//函数
func method01(s Student){
	fmt.Println(s.Name)
}
func method02(s *Student){
	fmt.Println((*s).Name)
}
func main(){
	var s Student = Student{"丽丽"}
	method01(s)
	method01(&s)//错误
	method02(s)//错误
	method02(&s)
}