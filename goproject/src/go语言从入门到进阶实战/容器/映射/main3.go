package main
import (
	"fmt"
	"sort"
)
func main(){
	scene := make(map[string]int)
	//准备数据
	scene["route"] = 66
	scene["brazi"] = 4
	scene["china"] = 960
	//声明一个切片
	var sceneList []string
	//将map数据遍历复制到切片中
	for k := range scene{
		sceneList = append(sceneList, k)
	} 
	sort.Strings(sceneList)
	fmt.Println(sceneList)
}
/*执行
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\映射\main3.go
[brazi china route]
创建一个map实例，键为字符串，值为整型
将3个键值对写入map中
声明sceneList为字符串切片，以缓冲和排序map中的所有元素
将map中元素的键遍历出来，并放入切片中
对sceneList字符串切片进行排序。排序时，sceneList会被修改
输出排好序的map的键
sort.Strings的作用是对传入的字符串切片进行字符串字符的升序排序
*/