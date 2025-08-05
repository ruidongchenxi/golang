package main
import "fmt"
// var (//全局变量
// 	cw int
// 	by byte
// )
 //var ce,cd = 12,"小捷"// 全局变量
func main(){
    var  i int = 1
	fmt.Println(i)
	//测试int8;溢出
	var e int8 = -128
	fmt.Println(e)
	//# command-line-arguments
//src\尚硅谷\chget1\deay1\dea1\main.go:7:15: cannot use -129 (untyped int constant) as int8 value in variable declaration (overflows)
    var w byte = 8
	fmt.Println(w)
	// var i int
	// fmt.Println("i=",i)
	// //类型推导
	// var num = 10
	// fmt.Println(num)
	// //短语
	// c := 7
	// fmt.Println(c)
	// x,z := 4,"cx"
	// fmt.Println(x,z)
	// c1,cs,re := 12,true,"晨曦"
	// fmt.Println(c1,cs,re)
	// fmt.Println(ce,cd)
	// fmt.Println(cw,by)

}



