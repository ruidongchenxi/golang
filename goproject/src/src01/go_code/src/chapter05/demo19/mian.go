package main

import (
	"fmt"
	"strconv"
	"strings"
) 

func main(){
    var c string = "chenxi河"//编码是统一utf-8 (字母数字ASCII（占一个字节），一个汉字占用3个字节)；按字节返回
    q := len(c) //统计字符串的长度
    fmt.Println(q)
    s:="hello晨曦"
    r := []rune(s)
    for i :=0;i<len(r);i++{
        fmt.Printf("字符串=%c\n",r[i])
    }
    x := "89"
    n,err := strconv.Atoi(x)
    if err == nil {
        fmt.Println("q=",n)
    }else{
        fmt.Println("err=",err)
    }
    d := "ty"
    z,err := strconv.Atoi(d) //错误转换报错：strconv.Atoi: parsing "ty": invalid syntax
    if err==nil {
        fmt.Println("z=",z)
    }else{
        fmt.Println("err=",err)
    }
    str1 :=strconv.Itoa(789)//数字转字符串
    fmt.Printf("str1的类型%T值是%v\n",str1,str1)
    var botes = []byte("hello go")
    fmt.Printf("botes=%v\n",botes)
    str2 :=string([]byte{97,98,99})
    fmt.Printf("str2=%v\n",str2)
    str3:=strconv.FormatInt(123,2)//2->,8.16//转2进制
    fmt.Printf("str3的2进制=%v\n",str3)
    str3=strconv.FormatInt(123,8)//2->,8.16//转8进制
    fmt.Printf("str3的8进制=%v\n",str3)
    str3=strconv.FormatInt(123,16)//2->,8.16//转8进制
    fmt.Printf("str3的16=%v\n",str3)
    fmt.Println(strings.Contains("chenxi","nx"))//如果字符串里存在指定的内容（子串）返回true，否则返回false
    fmt.Println(strings.Count("ceheese","e"))
    // fmt.Println(fmt.Println(strings.EqualFold("abc","ABC")))
    fmt.Println(strings.EqualFold("abc","ABC"))//true
    fmt.Println(strings.Index("NLT_abc","abc"))
    fmt.Println(strings.LastIndex("go golang","go"))
    fmt.Println(strings.Replace("go go hello","go","go语言",-1))
    fmt.Println(strings.Split("hello,wrold,ok",","))
    fmt.Println(strings.ToLower("Go")) //全部小写
    fmt.Println(strings.ToUpper("Go"))//全部大写
    fmt.Println(strings.TrimSpace(" tn a lone gopher ntrn "))//去掉左右两边空格
    fmt.Println(strings.Trim("! hello!"," !"))
    fmt.Println(strings.TrimLeft("! hello!"," !"))
    fmt.Println(strings.TrimRight("! hello!"," !"))
    fmt.Println(strings.HasPrefix("ftp://192.168.10.1","ftp"))
    fmt.Println(strings.HasSuffix("NLT_abc.jpg","abc"))
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo18\mian.go
// 9
// 字符串=h
// 字符串=e
// 字符串=l
// 字符串=l
// 字符串=o
// 字符串=晨
// 字符串=曦
// q= 89
// err= strconv.Atoi: parsing "ty": invalid syntax
// str1的类型string值是789
// botes=[104 101 108 108 111 32 103 111]
// str2=abc
// str3的2进制=1111011
// str3的8进制=173
// str3的16=7b
// true
// 4
// true
// 4
// 3
// go语言 go语言 hello
// [hello wrold ok]
// go
// GO
// tn a lone gopher ntrn
// hello
// hello!
// ! hello
// true
// false