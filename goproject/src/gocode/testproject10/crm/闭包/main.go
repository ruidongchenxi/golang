package main
import "fmt"
// 函数求和
//函数名字：getSum 参数为空
//函数getSum 函数返回值为一个参数，这个函数的参数是一个int 类型的参数，返回值也是int 类型

func getSum() func (int) int {
    var sum int = 10 
	return  func(num1 int) int{
		sum = sum + num1
		return sum
	}
}
//闭包： 返回的匿名函数+匿名函数以外的变量是num
func main(){
	f := getSum()
	fmt.Println(f(1))//11
	fmt.Println(f(3))//14
	fmt.Println("---------------------------------------")
	fmt.Println(getSum01(1))
	fmt.Println(getSum01(2))
	fmt.Println(getSum01(3))
}
// 不累加
func getSum01(num int) int{
	var sum int = 0
	sum = sum + num
	return sum
}
//不使用闭包的时候： 想保留的值，不可以反复使用
//闭包应用场景：闭包可以保留上次引用的某个值，我们传入一次就可以反复使用