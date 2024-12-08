package main
import "fmt"
//接口a
type CInterface interface{
	c()
}
//接口b
type BInterface interface{
	b()
}
type AInterface interface{
	CInterface
	BInterface
	a()
}


type E interface{
	b()
}
type Stu struct{

}
func (s Stu) a(){
    fmt.Println("a")
}
func (s Stu)  b(){
	fmt.Println("b")
}
func (s Stu)  c(){
	fmt.Println("c")
}
func main(){
	var s Stu
	var a E = s //空接口可以接任何接口
	a.a()
	a.b()
	a.c()

}
