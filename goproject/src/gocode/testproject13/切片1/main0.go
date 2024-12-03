package main
import "fmt"
func main(){
	//定义数组
	var intarr [6]int = [6]int{3,4,5,6,7,8}
	//切片
	//定义一个切片名字为slice.[]动态变化的数组长度不写，int类型，intarr原数组
	//[1:3]切片切出的一个片段---》索引：从1开始，到3结束，不包括3 【包左不包右】
	//var slice [] int = intarr[1:3]
	slice := intarr[1:3]
	fmt.Println("intarr:",intarr)
	//输出切片
	fmt.Println("slice:",slice)
	//切片元素个数
	fmt.Println("切片元素个数：",len(slice))
	//获取切片的容量
	fmt.Println("切片容量：",cap(slice))
	fmt.Printf("数组下标为1的内存地址：%p\n",&intarr[1])
	fmt.Printf("切片下标为0的内存地址：%p\n",&slice[0])
	slice[1] = 16
	fmt.Println(slice[1])
}