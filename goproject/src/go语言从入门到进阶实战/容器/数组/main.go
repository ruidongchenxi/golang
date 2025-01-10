package main
import (
	"fmt"
)
func main(){
	//声明team为3个元素的字符串数组
	
	var team [3]string
	//为team的元素赋值
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mun"
	fmt.Println(team)
}
/*
执行后的结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\数组\main.go
[hammer soldier mun]
*/