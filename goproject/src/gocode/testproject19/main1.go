package main
import (
	"fmt"
	"reflect"
)
func testReflect(i interface{}){
	//空接口没有任何方法，所以可以理解为所有类型都实现空接口
	//调用TypeOf函数，返回reflect.Value类型数据
	retype := reflect.TypeOf(i)
	fmt.Println("reType:",retype)
	//调ValueOf函数，返回reflect.Value类型数据
    reValue := reflect.ValueOf(i)
	fmt.Println("reValue",reValue)
	//如果真想获取reValue的数值，要调用Int()方法：返回v持有的有符号整数
	//num2 := 80 + reValue.Int()
	//fmt.Println(num2)
	k1 := retype.Kind()
	fmt.Println(k1)
	k2 := reValue.Kind()
	fmt.Println(k2)
	iw := reValue.Interface()
	//类型断言
	n,flag := iw.(Student)
	
}

type Student struct{
	Name string
	Age int
}
func main(){
	stu := Student{
		Name : "丽丽",
		Age : 18,
	}
	testReflect(stu)

}