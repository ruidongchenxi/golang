package main
import (
	"fmt"
)
func test (arr [3]int){
	arr[0] = 88
	fmt.Println("test",arr)
}
func test1 (arr *[3]int){
	(*arr)[0]=90 //（*arr[]）
	fmt.Println("test1",*arr)

}
func main(){
	// var arr01 [3] int
	// arr01[0]=1
	// arr01[1]=30
    // arr01[2]=1.1 cannot use 1.1 (untyped float constant) as int value in assignment (truncated)数据类型不同报错
	// arr01[2]=9
	// arr01[3]=8   invalid argument: index 3 out of bounds [0:3]  数组下标超出范围报错；不能动态变化
	
	// fmt.Println(arr01)
	var arr02 [3] int
	// var arr03 [3] string
	// var arr04 [2] bool
	// fmt.Println(arr02)
	// fmt.Println(arr03)
	// fmt.Println(arr04)
	// test(arr02)
	fmt.Println("man",arr02)
	test1(&arr02)
	fmt.Println("man",arr02)

	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter06\demo03\main.go
// man [0 0 0]
// test1 [90 0 0]
// man [90 0 0]