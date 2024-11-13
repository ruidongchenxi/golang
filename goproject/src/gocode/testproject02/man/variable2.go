package main
import "fmt"
//全局变量定义
var n7 = 10
var n8 = 9.7
func main(){
	//定义在{}里的变量叫局部变量
	var w int 
	w = 3
	fmt.Println(w)
	//第二种.指定类型使用默认值
	var nu2 int
	fmt.Println(nu2)
	//第三种类型推导
	var nu1 = 1
	fmt.Println(nu1)
	//第四种省略var 关键子,注意是:=
	set := "男"
	fmt.Println(set) 
    //声明多个变量
	var chx2,cx,re int
	fmt.Println(chx2)
	fmt.Println(cx)
    fmt.Println(re)
	var ca,cs,cd = 1,"吃",7.8
	fmt.Println(ca)
	fmt.Println(cs)
	fmt.Println(cd)
    fmt.Println(n8)
}
