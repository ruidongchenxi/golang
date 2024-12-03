package main
import "fmt"
func main(){
    //定义切片：make 函数的三个参数1切片类型 2切片长度 3切片容量
	slice := make([]int,4,20)
	fmt.Println(slice)
	fmt.Println("切片长度",len(slice))
	fmt.Println("切片容量",cap(slice))
	slice2 := []int{1,2,3}
	fmt.Println("切片长度",len(slice2))
	fmt.Println("切片容量",cap(slice2))
}