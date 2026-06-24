package main
import "fmt"
var (
	name string
	age int 
	isok bool
)
const (
	n1 = 100
	n2 = 200
)
func main(){
	//fmt.Println("你好")
	//var t int = 5
	// name = "理想"
	// age = 16
	// isok = true
	// fmt.Print(isok) //在终端打印输出的内容
	// fmt.Println(age)// 打印完指定内容添加换行符
	// fmt.Printf("name:%s\n",name) //%s:占位符 使用name这个变量的值去替换占位符
	var s string = "web"
	//类型推导
	var r = "golang"
	//短变量声明
	e := "你好"
	fmt.Println(s)
	fmt.Println(r)
	fmt.Println(e)

}
//执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo01\main.go
// PS D:\golang\goproject\src\lnhgo> go run day1\demo01\main.go
// web
// golang
// 你好