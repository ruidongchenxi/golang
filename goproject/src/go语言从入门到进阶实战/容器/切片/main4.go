package main
import (
	"fmt"
)
func main(){
	//声明字符串切片
	var strList []string
	//声明整型切片
	var numList []int
	//声明空切片
	var numListEmpty = []int{}//numListEmpty声明为一个整型切片。本来会在“{}”中填充切片的初始化元素，这里没有填充，所以切片是空的，但此时numListEmpty已经被分配了内存，但没有元素
	//输出3个切片
	fmt.Println(strList, numList, numListEmpty)//切片均没有元素，3个切片输出元素内容均为空
	//输出3个切片大小
	fmt.Println(len(strList),len(numList),len(numListEmpty)) //没有对切片进行任何操作，strList和numList没有指向任何数组或者其他切片
	//切片判定空结果
	fmt.Println(strList == nil)//声明但未使用的切片的默认值是nil.strList和numList也是nil，所以和nil比较的结果是true
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)//numListEmpty 已经被分配到了内存，但没有元素，因此和nil比较是false，切片是动态结构，只能与nil判定，不能互相判等时

}
/*执行后结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\切片\main4.go
[] [] []
0 0 0
true
true
false

*/