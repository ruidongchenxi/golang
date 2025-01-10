package main
import (
	"fmt"
)
func main(){
	var car []string
	//添加1个元素
	car = append(car,"0ldDriver")
	//添加多个元素
	car = append(car,"lce","Sniper","Monk")
	//添加切片
	team := []string{"Pig","Flyingcake","Chicken"}
	car = append(car,team...)
	fmt.Println(car)
}
/*执行后结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\切片\main7.go
[0ldDriver lce Sniper Monk Pig Flyingcake Chicken]
声明一个字符串切片
往切片中添加一个元素
使用append()函数向切片中添加多个元素
声明另外一个字符串切片
在team后面加上“...”,表示将team整个添加到car的后面
*/