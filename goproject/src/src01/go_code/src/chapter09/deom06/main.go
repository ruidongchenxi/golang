package main
import (
	"fmt"
)
//
type integer int 
func (i integer) print(){
	fmt.Println("i=",i)
}

//给A类型绑定一份方法

type Monster struct{
	Name string
	age int 
}
func (stu *Monster) String() string{
	str := fmt.Sprintf("Name=[%v],Age=[%v]",(*stu).Name,(*stu).age)
	return str
}
func main(){
	var i integer =120
	i.print()
	var t Monster = Monster{"cx",67}
	//如果实现了*Monster类型的String 方法，就会自动调用
	fmt.Println(&t)//
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter09\deom05\main.go
// i= 120
// Name=[cx],Age=[67]