package main
import (
	"fmt"
	"strings"
)
func StringProccess (list []string, chain []func(string) string) {
	//遍历每个字符串
	for index, str := range list {
		//第一个需要处理的字符串
		result := str
		//遍历每一个处理链
		for _,proc := range chain {
			//输入一个字符串进行处理返回数据作为下一个链的输入
			result = proc(result)
		}
		//将结果放回切片
		list[index] = result
	}
}
/*传入字符串切片list作为数据源，一系列的处理函数作为chain处理链
遍历字符串切片的每个字符串，依次对每个字符串进行处理
将当前字符串保存到result变量中作为第一个函数处理的参数
遍历每个处理函数，将字符串按顺序经过这些处理函数处理
result遍历即每个处理函数的输入变量，处理后的变量重新保存到result中
将处理完的字符串保存会切片中


*/
//自定义的移除前缀处理函数
func removePrefix(str string) string{
	return strings.TrimPrefix(str,"go")
}
/*
自定义函数处理
处理函数可以是系统提供的处理函数，如将字符串变大写或小写，也可以使用自定义函数，字符串处理逻辑是使用一个自定义函数实现移除指定go 前缀的过程
此函数使用了strings.TrimPrefix()函数实现移除字符串指定的前缀。处理后，移除前缀的字符串结果将通过 removePrefix()函数的返回值返回

*/
//字符串处理主流程
/*
字符串处理的主流程包含
准备处理字符串列表
准备字符串处理链
处理字符串列表
打印输出后的字符串列表
*/
func main() {
	//待处理的字符串列表
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	//处理函数链
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	//处理字符串
	StringProccess(list,chain)
	//输出处理好的字符串
	for _,str := range  list {
		fmt.Println(str)
	}
}
//执行结果
/*
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\函数变量\main1.go
SCANNER
PARSER
COMPILER
PRINTER
FORMATER
定义字符串切片，字符串包含go前缀及空格，处理的顺序与函数在切片中位置一致。removePrefix()为自定义函数，功能是移除go前缀；移除前缀的字符串左边还有一个空格，使用strings.TrimSpace移除，这个函数的定义刚好符合处理函数格式：func(string)string;strings.ToUpper用于将字符串转为大写
传入字符串切片的每一个字符串，打印处理好的字符串结果

*/