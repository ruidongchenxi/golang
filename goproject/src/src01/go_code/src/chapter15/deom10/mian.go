package main

import (
	"bufio"
	"fmt"
	"io"
	//"math/rand"
	"os"
	"unicode"
)

//定义一个结构体用于保存结果
type CharCount struct{
	CharCount  int //记录英文个数
	NumCount int //记录数字个数
	SpaceCount int //记录空格个数
	OtherCount int //记录其他字符
	Chr int
}
func main(){
	//思路： 打开一个文件，创一个Reader
	//每读取一行，就去统计该行有多少个英文、数字、空格和其他字符
	//然后将结果保存到一个结构体中
	fileName:="D:/abc.txt"
	file,err := os.Open(fileName)
	if err !=nil{
		fmt.Println("open file err=",err)//
		return
	}
	defer file.Close()
	//定义CharCount 实力
	var count CharCount 
	//创建一个reader
	reader:=bufio.NewReader(file)
	for {
		str,err:=reader.ReadString('\n')

	
		for _,r:=range str{
		 
		switch {
		case unicode.Is(unicode.Han, r): //匹配中文
			count.Chr++
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'):
			count.CharCount++
		case r >= '0' && r <= '9':
			count.NumCount++
		case r == ' ' || r == '\t':
			count.SpaceCount++
		default:
			count.OtherCount++
	    }
		 }
		if err == io.EOF{
			break
		}
		}
	fmt.Printf("字符个数%v，数字个数%v，空格个数%v，其他字符个数%v,中文%v",
			count.CharCount, count.NumCount, count.SpaceCount,count.OtherCount,count.Chr)
}
