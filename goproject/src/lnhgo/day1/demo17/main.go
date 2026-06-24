package main
import (
	"fmt"
	//"math"
)
func main() {
	var house = "Malibu Point 10880,90265"
	//对字符串取地址值,ptr类型为*string
	ptr := &house
	fmt.Printf("ptr type: %T\n",ptr)
	//打印ptr的指针地址
	fmt.Printf("address:%p\n",ptr)
	//对指针进行取值
	value := *ptr
	//查看取值后的类型
	fmt.Printf("value type: %T\n",value)
	//指针取值后就是指向变量的值
	fmt.Printf("value: %s\n",value)


}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo17\main.go
// ptr type: *string
// address:0xc000026070
// value type: string
// value: Malibu Point 10880,90265