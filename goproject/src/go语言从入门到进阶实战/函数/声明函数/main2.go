package main
import (
	"fmt"
)
const (
	//定义每分钟的秒数
	SecondsPerMinute = 60
	//定义每小时多少秒
	SecondsPerHour = SecondsPerMinute * 60
	//定义每天多少秒
	SecondsPerDay = SecondsPerHour * 24
)
// 将传入的秒解析为3种时间单位
func resolveTime(seconds int) (day int, hour int, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}
func main() {
	//将返回值作为打印参数
	fmt.Println(resolveTime(1000))
	//只小时和分钟
	_, hour, minute := (resolveTime(18000))
	fmt.Println(hour, minute)
	//只获取天
	day,_,_ := resolveTime(90000)
	fmt.Println(day)
}
/*
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\声明函数\main2.go
0 0 16
5 300
1
常量里定义
定义每分钟的秒数
定义每小时的秒数
定义每天的秒数
定义resolveTime()函数，根据输入的秒数返回3个整型值，含义分别是秒数对应的天数小时数分钟数取整
给定1000秒，对应16（16.6667取整）分钟的秒数，resolveTime()函数返回的3个变量会传递给fmt.Println()函数进行打印，因为fmt.Println()使用了可变参数，可以接收不定量参数
将resolveTime()函数中的3个返回值使用变量接收，但是第一个返回值使用匿名函数接收，表示忽略这个变量
忽略后两个返回值，只使用第一个返回值
*/