package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)
func main(){
	file2,err2 := os.Open("D:/abc.txt")
	if err2 !=nil{
		fmt.Println("open file err =",err2)
	}
	defer file2.Close()
	t := bufio.NewScanner(file2)
	for t.Scan(){
		line := t.Text() // 当前行
		fmt.Println(line)
	}
	//1.打开文件
	file,err:= os.OpenFile("d:/abc.txt", os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		fmt.Printf("open file %v\n",err)
		return
	}
	defer file.Close()
	//准备写入内容
	str := "'hello,北京\r\n'"
	//写入时使用带缓存的*Writer
	Writer:=bufio.NewWriter(file)
	for i:=0;i<5;i++{
		Writer.WriteString(str)
	}
	//因此writer 带缓存到内存的，因此要同步磁盘
	Writer.Flush()//同步到磁盘
	//
}
