package main
import (
	"fmt"
)
//将int声明为ChipType芯片类型
type ChipType int
// 将const 里定义的一句常量值设为ChipType类型，且从0开始，每行值加1
const (
    None ChipType = iota
	CPU
	GPU
)
//定义ChipType类型的方法String(),返回字符串
func (c ChipType) String() string{
	//使用switch语句判断当前ChipType类型，返回对应的字符串
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
		
	}
	return "N/A"
}
func main(){
	//输出CPU的值，并以整形格式显示
	fmt.Printf("%s %d",CPU, CPU)
}
/*
使用String()方法的ChipType在使用上和普通常量没有区别。当这个类型需要显示为字符串时，go语言会自动寻找方法并调用
*/