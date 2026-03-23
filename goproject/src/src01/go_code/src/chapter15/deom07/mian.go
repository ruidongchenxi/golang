package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	//"strings"
)
func main(){

	//1.打开文件
	file,err:= os.OpenFile("d:/abc.txt", os.O_RDWR|os.O_APPEND,0666)
	if err!=nil{
		fmt.Printf("open file %v\n",err)
		return
	}
	defer file.Close()
	read:=bufio.NewReader(file)
	for {
		str,err:=read.ReadString('\n') //读到换行符结束一行
		if err == io.EOF{ //io.EOF,表示文件结尾
			break
		}
	   fmt.Print(str) 
	}
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
