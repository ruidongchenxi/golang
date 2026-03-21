package main

import (
	"fmt"
	//"src01/demo16/test"
)
type Binterface interface{
	test01()
	test03()
}
type Cinterface interface{
	test02()
	test03()
}
type Ainterface interface{
	Binterface
	Cinterface
	
}

type Stu struct{
}
func (a Stu)test01(){
	fmt.Println("test01")
}
func (a Stu)test02(){
	fmt.Println("test02")
}
func (a Stu)test03(){
	fmt.Println("test03")
}
type T interface{}

//如果要实现A接口需要把B接口方法和C接口方法全部实质再把A接口里特有的方法也一并实现才可以
func main(){
   var A Stu
   var s Ainterface= A //A的结构类型是stu ，这个类型实现了test01()test02()test03,所以可以直接实现Ainterface，
//   // 有因为Ainterface的接口继承Binterface和Cinterface 和自身的test03接口，所以这个Ainterface接口只能且必须匹配满足这3个方法的类型
    s.test03()// test01 是C 接口的实现，因为A继承了C。所以可以调用
//   var t T = A //ok
//   fmt.Println(t)
//   var t2 interface{}
//   fmt.Println(t2)
//   var t3 interface{}= A //ok
//   fmt.Println(t3) //ok
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter11\demo02\main.go
// # command-line-arguments
// chapter11\demo02\main.go:17:2: duplicate method test03
//         chapter11\demo02\main.go:16 of:2: other declaration  test03