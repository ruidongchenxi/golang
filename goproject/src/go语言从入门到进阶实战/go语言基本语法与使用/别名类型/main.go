package main
import (
	"fmt"
)
//将NewInt 定义为int类型
type NewInt int
//将int取一个别名
type IntAlias = int
func main(){
	//将a声明为NewInt类型
	var a NewInt
	//查看a的类型名
	fmt.Printf("a type: %T\n", a)
	//将a2 声明IntAlias类型
	var a2 IntAlias
	//查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)
}
/*
将NewInt定义为int 类型，这是常见定义类型的方法，通过type关键字的定义，NewInt会形成一种新的类型。NewInt本身依然具备int特性
将IntAlias 设置为int的别名,使用IntAlias与int 等效
使用%T格式化参数，显示a变量本身的类型
将a2声明为IntAlias类型，此时打印a2的值为0
显示a2 变量类型
*/