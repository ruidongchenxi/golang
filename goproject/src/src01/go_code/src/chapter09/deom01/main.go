package main
import (
	"fmt"
)
type cat struct{   //结构体定义
	name string
	color string
	age int
}
func main() {
	//
	// var c1 cat
	t1 := cat{"小花","黑色",65}//实例化结构体
	t2 := cat{"小红","黄色",67}
	fmt.Println(t1.name,t2.name)//调用实例化的结果体
}
// 执行结果
// 小花 小红