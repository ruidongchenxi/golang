package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)
func main(){
	// //字符串比较
	// a:="hello"
	// b:="bello"
	// fmt.Println(a==b)
	// fmt.Println(a!=b)
	// //字符串大小比较
	// fmt.Println(a>b)
	// 
	name:="imooc体系课-go工程师go"
	//是否包含
	fmt.Println(strings.Contains(name,"go"))
	//字符串长度
	//len() 统计的是：字节数（byte）而不是字符数。汉字使用UTF-8编码，UTF-8 占 3 个字节
	fmt.Println(len(name))
	//统计字符个数（包含汉字）使用 utf8.RuneCountInString
	fmt.Println(utf8.RuneCountInString(name))
	//查询子串出现的次数
	fmt.Println(strings.Count(name,"o"))
	//分隔
	fmt.Println(strings.Split(name,"-"))
	//字符串是否指定的前缀开头
	fmt.Println(strings.HasPrefix(name,"im"))
	//字符串是否以指定的内容结尾
	fmt.Println(strings.HasSuffix(name,"师"))
	//查询子串出现的位置
	fmt.Println(strings.Index(name,"go"))
	//查询子串出现的位置
	fmt.Println(strings.IndexRune(name,[]rune(name)[8]))
	//子串替换-1表示全部替换整数表示替换几个
	fmt.Println(strings.Replace(name,"go","python",-1))
	//转换成小写
	fmt.Println(strings.ToLower("GO"))
	//转换成大写
	fmt.Println(strings.ToUpper(name))
	//去掉字符串左右两边特舒字符
	fmt.Println(strings.Trim("*$hello *go*","*$"))


	

}