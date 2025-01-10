package main
import (
	"fmt"
)
func main(){
	m := map[string]string{
		"W" : "forward",
		"A" : "left",
		"D" : "right",
		"S" : "backward",
	}
	fmt.Println(m["W"])
}
/*
例子中并没有使用make，而是使用大括号进行内容定义，就像json格式一样，冒号的左边是key，右边是值，键值对之间使用逗号分隔
*/