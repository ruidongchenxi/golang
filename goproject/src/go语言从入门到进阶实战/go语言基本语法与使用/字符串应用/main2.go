package main
import (
	"fmt"
	"strings"
)
//使用string.Index()函数在字符在搜索另一个子串
func main() {
	//utf8 一个汉字3个字节
	tracer := "死神来了，  死神bye bye"
    //在尝试tracer的字符中搜索中文逗号，返回的位置存在comma变量里，类型是int。表示从tracar变量字符串开始的aSCII 
	comma := strings.Index(tracer, "，")
	//strings.Index()函数并没有像其他语言一样，，提供一个从某偏移开始搜索的功能，不过可以对字符串进行切片操作来实现这个逻辑
	//tracer[comma:]从stracer的comma位置开始到tracer字符串的结尾构造一个子串，返回strings.Index()进行在索引，得到pos是相对于tracer[comma:]的结果
	// 
	pos := strings.Index(tracer[comma:] , "死神")
	//  
	fmt.Println(comma,pos, tracer[comma+pos:])

}
//字符串索引常用方法
/*
strings.Index:正向搜索
strings.LastIndex: 反向搜索子窜
搜索的起始位置可以通过切片偏移制作
*/