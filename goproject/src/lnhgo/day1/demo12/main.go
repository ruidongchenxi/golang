package main

import (
	"fmt"
	//"string"
	//"strings"
)

//fmt 占位符
func main() {
	// 字符串
	// s1 := "白萝卜"
	// s3 := []rune(s1)
	// s3[0]='红'
	// fmt.Println(string(s3))
	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n",c1,c2)
	c3 :="H" //string
	c4 := byte('H')//uint8
	c5 :='H'//int32
	fmt.Printf("c3:%T c4:%T c5:%T\n",c3,c4,c5)
	fmt.Printf("%d\n",c4)

}
// PS D:\golang\goproject\src\lnhgo> go run day1\demo11\main.go
// c1:string c2:int32
// c3:string c4:uint8 c5:int32
// 72