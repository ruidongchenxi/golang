package main
import (
	"fmt"
)
type Person struct{
	Num string
}

//给A类型绑定一份方法
func (p Person)test(){// A 结构体有一个方法
	fmt.Println("test()",p.Num)
}
func (p Person)speak(){
	fmt.Printf("%s你是好人\n",p.Num)
}
func (p Person)jisuan(){
	var t int 
	for i:=1;i<=1000;i++{
		t+=i
	}
	fmt.Println("1+...+1000=",t)

}
func (P Person)getSum(a int,b int) int{
		return a+b
}
func (P Person)jisuan1(n int){
	var t int 
	for i:=1;i<=n;i++{
		t +=i
	}
	fmt.Printf("1+...+%v=%v",n,t)
}
type Circle struct{
	radius float64
}
func (c Circle) ares() float64{
	return c.radius*c.radius*3.14
}
func (c *Circle) ares2() float64{
	//因为c是一个指针结构体实例，因此标准的访问其字段的方式是.(*c).radius
	return (*c).radius*(*c).radius*3.14
}
func main(){
	r := Circle{85}
	fmt.Printf("ares1()方法的输出半径是%v,的圆的面积%v\n",r.radius,r.ares())//	fmt.Printf("半径是%v,的圆的面积%v\n",r.radius,r.ares) 不可以这样写，方法调用必须加(),否则就输出方法的地址值
	t := Circle{76}
	// fmt.Printf("ares2()方法的输出半径是%v,的圆的面积%v\n",t.radius, (&t).ares2())
	//编译器底层做了优化(&t).ares2()等价于t.ares2()
	//因为编译器底层添加了&
	fmt.Printf("ares2()方法的输出半径是%v,的圆的面积%v。。。。。\n",t.radius, t.ares2())

}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter09\deom05\main.go
// ares1()方法的输出半径是85,的圆的面积22686.5
// ares2()方法的输出半径是76,的圆的面积18136.64。。。。。