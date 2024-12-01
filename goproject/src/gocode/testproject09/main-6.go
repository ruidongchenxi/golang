package main
import "fmt"
func test (num int){
	fmt.Println(num)
}
func test2 (num1 int ,num2 float32, testFunc func(int)){
	 fmt.Println("------test02")
}
func main(){
	a := test
	fmt.Printf("a的类型是：%T,test函数的类型是：%T \n",a,test)
	a(10) //等价test(10)
    test2(10,3.19,test)
}