package main
import "fmt"
func test (num int){
	fmt.Println(num)
}
func test2 (num1 int ,num2 float32, testFunc func(int)){
	 fmt.Println("------test02")
}
type myFunc func(int)
func test3 (num1 int ,num2 float32, testFunc myFunc){
	fmt.Println("------test03")
}
func test04 (num1 int,num2 int)(int,int){
	result1 := num1 + num2
	result2 := num1 - num2
	return result1,result2
}
//
func test05 (num1 int,num2 int)(sub int,sum int){
	sum := num1 + num2
	sub := num1 - num2
	return 
}
func main(){
	a := test
	fmt.Printf("a的类型是：%T,test函数的类型是：%T \n",a,test)
	a(10) //等价test(10)
    test2(10,3.19,test)
	//自定义数据类型
	type myInt int 
	var num1 myInt = 30
	fmt.Println("num1",num1)
	var num2 int = 30
	num2 = int(num1) //虽然起别名，但是编译识别的时候还是识别为不同数据类型的变量；强制数据类型转换，
	fmt.Println("num2",num2)
	test3(3,3.13,a)

}