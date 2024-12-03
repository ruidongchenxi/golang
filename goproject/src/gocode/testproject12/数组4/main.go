package main
import "fmt"
func main(){
	//定义一个数组
	/*
	var arr1 = [3]int{3,6,9}
	fmt.Println(arr1)
	fmt.Printf("数组的类型:%T\n",arr1)
	var arr2 = [5]int{3,4,5,6,8}
	fmt.Printf("数组的类型:%T\n",arr2)
	*/
	var arr3 = [3]int{3,6,9}
	test1(arr3)
    fmt.Println(arr3)
}
func test1 (arr [3]int){
	arr[0] = 7
}