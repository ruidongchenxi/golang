package main
import (
	"fmt"
)
func main(){
	//准备一个字符串变量
	var house = "mali bu polint 10880 90265"
	//对字符串取值。prt类型位*string
	prt := &house
	//打印ptr的类型
	fmt.Printf("ptr type: %T\n",prt)
	//打印ptr的指针地址
	fmt.Printf("address:%p\n",prt)
	//对指针进行取值操作
	value := *prt
	//取值后的类型
	fmt.Printf("value type: %T\n",value)
	//指针取值后就是指向变量的值
	fmt.Printf("value: %s\n",value)
//	fmt.Println(value)
//	fmt.Println(&value) 
}