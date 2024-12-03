package main
import "fmt"
func main(){
	var a []int = []int{1,2,3,4,6,5}
	var b []int = make([]int,10)
	copy(b,a)//将a对应数组中元素拷贝b中对应的数组中去
	fmt.Println(b)
}