package main
import "fmt"
func main(){
	//变量声明；类型
	// var age int
	//变量赋值
	// age = 18
	//变量使用
	var age int = 19 //声明赋值合成一句
	fmt.Println("小明年龄",age) //逗号做拼接
/* 	var age int = 9
	fmt.Println("小红年龄",age)
	# command-line-arguments
src\src01\demo1\main0.go:11:6: age redeclared in this block
        src\src01\demo1\main0.go:9:6: other declaration of age */
/* 	var num int = 12.7
	fmt.Println(num)
	PS D:\golang\goproject> go run src\src01\demo1\main0.go
# command-line-arguments
src\src01\demo1\main0.go:16:16: cannot use 12.7 (untyped float constant) as int value in variable declaration (truncated) */
}