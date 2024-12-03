package main
import "fmt"
func main(){
	var arr [3][3] int16
	fmt.Println(arr)
	fmt.Printf("arr地址值：%p\n",&arr)
	fmt.Printf("arr[0]的地址值：%p\n",&arr[0])
	fmt.Printf("arr[0][0]的地址值：%p\n",&arr[0][0])
	fmt.Printf("arr[0][2]的地址值：%p\n",&arr[0][2])
	fmt.Printf("arr[1]的地址值：%p\n",&arr[1])
	fmt.Printf("arr[1][0]的地址值：%p\n",&arr[1][0])
	//赋值
	arr[0][2] = 16
	arr[1][2] = 67
	fmt.Println(arr)
	//初始化操作
	var arr1 [2][3]int = [2][3]int{{1,2,4},{7,8,9}}
	fmt.Println(arr1)
	
}
