package main
import (
	"fmt"
	"reflect"
)
//定义结构
type Student struct{
	Name string
	Age int

}
func (s Student) CPrint()  {
	fmt.Println("调用了Print方法")
	fmt.Println("名字",s.Name)
}
func (s Student) GetSum(n1,n2 int) int{
	return n1 + n2
}
func (s Student) Set(name string,age int){
	s.Name = name
	s.Age = age
}
//定义函数操作结构体反射进行反射操作
func TestStudentStruct(a interface{}){
    va := reflect.ValueOf(a)
	fmt.Println(va)
	n1 := va.NumField()
	fmt.Println(n1)
	//遍历
	for i := 0; 1<n1;i++{
		fmt.Println("第%d个字段的值是:%v",i,va.Field(i))
		n2 := va.NumMethod()
		fmt.Println(n2)
	}
	//方法
	n2 := va.NumMethod()
	fmt.Println(n2)
	va.Method(2).Call(nil)
}
func main(){
	//定义结构体具体的
	s := Student{
		Name : "丽丽",
		Age : 18,
	}
	TestStudentStruct(s)
}