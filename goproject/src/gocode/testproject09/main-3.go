package main
import "fmt"
//定义可变变量
func test (args...int)  { //如果返回值类型就一个的话，那么（）可以省略不写的
//在函数内部处理可变参数的时候，将可变参数当做切片处理
//遍历可变参数
     for i :=0; i <len(args);i++{
		fmt.Println(args[i])
	 }
}
func main(){
    test ()
	fmt.Println("..............")
	test (1)
	fmt.Println("..............")
	test (2,3,4,5)
	fmt.Println("..............")

}