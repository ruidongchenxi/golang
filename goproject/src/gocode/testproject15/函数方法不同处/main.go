package main
import "fmt"

//结构体
type Student struct{
	Name
}
//方法
func (s Student) test01(){
	fmt.Println(s.Name)
}
//函数
func method01(s Student){
	fmt.Println(s.Name)
}
func main(){
	//函数调用
	var s Student = Student{"丽丽"}
	//函数调用
	method01(s)
	//方法调用
	s.test01

}