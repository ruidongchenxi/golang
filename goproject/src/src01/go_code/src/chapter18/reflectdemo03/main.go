package main

import (
	//"fmt"
	"fmt"
	"reflect"
)
func reflect01(b interface{}){
	c := reflect.ValueOf(b)
	d := reflect.TypeOf(b)
	fmt.Printf("c的kind=%v\n",c.Kind())
	fmt.Printf("c的type=%v\n",d)
	fmt.Println("c=",c)	
	r := c.Interface()
	a := r.(float64)
	fmt.Println(a)

}
func main(){
	var c float64= 90.6
	reflect01(c)
	str := "cdf"
	fs :=reflect.ValueOf(&str)//要修改string值要传地址值，因为string是值类型
	fs.Elem().SetString("你好")//通过地址值取到原来值后并修改原来值
	fmt.Println("修改后str=",str)//打印str 变量查看修改
}
// 执行结果
// 修改后str= 你好