package main
import "fmt"
func main(){
	//第一种
	var arr1 [3]int = [3]int {3,6,9}
	fmt.Println(arr1)
	//第二种
	var arr2 = [3]int {1,4,7}
	fmt.Println(arr2)
	//
	var arr3 = [...]int {1,2,3,4,5,6}
	fmt.Println(arr3)
	//
	var arr4 = [...]int{1:45,0:89,3:44}
	fmt.Println(arr4)
}	