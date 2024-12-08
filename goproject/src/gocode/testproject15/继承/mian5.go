package main
import (
	"fmt"
)
type A struct{
	a int
	b string
}
type B struct{
	c int
	d string
	a int 
}
type C struct{
	A
	B
	int
}
type D struct{
	a int 
	b int 
	c B
}
func main(){
	d := D{10,20,B{23,"ppp",99}}
	fmt.Println(d)
	fmt.Println(d.c.d)
	//构造C结构体的实例化
	//
	//c := C{A{10,"aaa"},B{20,"bbb",89},17}
	c := C{
		A{
			a : 10 ,
			b : "aa",},
		B{
			c : 54,
			d : "福贵",
			a : 98, },
		75,
	}
    fmt.Println(c.A.a)
	fmt.Println(c.B.a)
	fmt.Println(c.b)
	fmt.Println(c.int)
}
