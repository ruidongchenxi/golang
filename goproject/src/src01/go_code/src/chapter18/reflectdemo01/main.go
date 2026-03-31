package main

import (
	"fmt"
	"reflect"
)

//反射
func receiveTest01(b interface{}){
	//通过反射获取到传入的变量类型，以及类别，值是什么
	//先获取类型
	rType :=reflect.TypeOf(b)//
	fmt.Println(rType)
	fmt.Println(rType.Name())
	//获取vakue
	rvlue := reflect.ValueOf(b)
	fmt.Printf("revlue=%v type=%T\n",rvlue,rvlue)
	//取值并计算
	n2 := 2+rvlue.Int()
	fmt.Println(n2)
	//下面将rvlue转成接口
	c := rvlue.Interface()
	//接口通过断言转回int
	d := c.(int)
	fmt.Printf("c=%v type=%T\n",d,d)
}
type  Student struct{
	Name string
	Age  int
}
func receiveTest02(b interface{}){
	rType :=reflect.TypeOf(b)
	fmt.Println("类型",rType)
	rVal := reflect.ValueOf(b)
	//获取变量对应的kind
	//rVal.Kind()===>
	//rType.Kind()===>
	fmt.Printf("kind1=%v ****\n",rType.Kind())
	fmt.Printf("kind2=%v 55\n",rVal.Kind())

	fmt.Printf("revlue=%v type=%T\n",rVal,rVal)
	//转为
	x := rVal.Interface()
	//将接口转成通过断言需要判断的类型
	//通过Switch 的断言形式做的更加灵活
	d ,ok:= x.(Student)
	if ok{
		fmt.Println("名字",d.Name)
	}

}

func main(){
	//定义int 类型
	// var num int = 100
	// receiveTest01(num)
	var t Student = Student{"小鬼",8}
	receiveTest02(t)
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter18\reflectdemo01\main.go
// 类型 main.Student
// kind=struct ****
// kind=struct 55
// revlue={小鬼 8} type=reflect.Value
// 名字 小鬼