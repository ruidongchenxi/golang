
package main
import (
	"fmt"
)
func main(){
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazi"] = 4
	scene["china"] = 960
	for k, v := range scene{
		fmt.Println(k,v)
	}
}
/*执行后结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\映射\main2.go
brazi 4
china 960
route 66
遍历对于go 语言的很多对象来说都差不多，直接使用for range语法。遍历时可以同时获取键和值。如只遍历值时可以使用
for _, v := range scene{
将不需要的键改为匿名变量形式
只遍历键时使用
for k := range scene{
无须将值改为匿名形式，忽略即可

*/