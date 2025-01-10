package main
import (
	"fmt"
)
func main(){
	scene := make(map[string]int)
	scene["route"] = 66
	fmt.Println(scene["route"])
	v := scene["route2"]
	fmt.Println(v)
}
/* 执行结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\映射\main0.go
66
0
map 是一个内部实现的类型，使用时，需要动手使用make创建。如果不创建使用map类型，会触发宕机错误
向map中加入映射关系，写法与使用数组一样，key可以使用除函数外任意类型
查找map 中的值
尝试查找一个不存在的键，那么返回的将是valueType的默认值
某些情况下需要明确知道查询中某个键是否在map中，可以使用一种特殊写法来实现
v, ok := scene["route"]
*/