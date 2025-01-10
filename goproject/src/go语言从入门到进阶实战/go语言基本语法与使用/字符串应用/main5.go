package main  
import (
	"fmt"
)
func main(){

	/*
	格式化在逻辑中非常常用。使用格式化函数
	fmt.Sprint(格式化样式，参数列表)
	格式化样式： 字符串形式，格式化动词以%开头
	参数列表：多个参数以逗号分隔，个数必须与格式化样式中的的个数一一对应，否则运行报错

	*/
	var progress =2
	var target =8
	title  := fmt.Sprintf("已经采集%d个草药,还需要%d个任务完成", progress, target)
	fmt.Println(title)
	pi := 3.14159
	//按照数值本身的格式化输出
	variant := fmt.Sprintf("%v %v %v","月球基地", pi,true)
	fmt.Println(variant)
	//匿名结构体
	profile := &struct{
		Name string 
		HP int
	}{
		Name: "rat",
		HP: 150,
	}
	fmt.Printf("使用'%%+v' %+v\n",profile)
	fmt.Printf("使用'%%#v' %#v\n",profile)
	fmt.Printf("使用'%%T' %T\n",profile)

}
/*

*/