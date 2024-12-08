package main
import (
	"fmt"
	//"strconv"
	//"time"
)
func main(){
	//定义管道
	var intChan chan int
	//通过make 初始化,管道可以存放3个int类型的数据
	intChan = make( chan int,3)
	//证明管道是引用数据类型
	fmt.Printf("intCha的值：%v\n",intChan)
	//向管道存放数据
	intChan<- 10
	num := 20
	intChan<- num
	intChan<- 50
	//不能存放大于容量的数据
	//从管道中读取数据
	num1 := <- intChan
	fmt.Println(num1)
	num2 := <- intChan
	fmt.Println(num2)
	num3 := <- intChan
	fmt.Println(num3)
	//在没有使用协程的情况下，如果管道的数据全部取出，那么再取就会报错
	num4 := <- intChan
	fmt.Println(num4)
	//输出管道长度
	fmt.Printf("管道实际的场地%v,管道容量%v\n",len(intChan),cap(intChan))
}