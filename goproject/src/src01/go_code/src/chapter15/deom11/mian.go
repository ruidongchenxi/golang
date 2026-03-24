package main

import (
	//"bufio"
	"bufio"
	"fmt"
	"io"

	//"internal/syscall/windows"
	//"io"
	//"io/ioutil"

	//"io/ioutil"
	"os"
)
func CopyFile(detFileName string,srcFileName string)(written int64,err error){
	srcFile,err:= os.Open(srcFileName)
	if err !=nil{
		fmt.Printf("源文件异常:%v\n",err)
	}
	defer srcFile.Close()
	//通过文件句柄
	reade :=bufio.NewReader(srcFile)
	detFile,err:=os.OpenFile(detFileName,os.O_WRONLY|os.O_CREATE,0666)
	if err !=nil{
		fmt.Printf("open file err=%v\n",err)
	}
	defer detFile.Close()
	// 通过disFile,获取Write
	writer:= bufio.NewWriter(detFile)
	return io.Copy(writer,reade)

}
func main(){
	//将"C:/Users/尘曦/Desktop/2x.jpg"拷贝到D:/abc.jpg
	//1
     src:="C:/Users/尘曦/Desktop/2x.jpg"
	 dst:="d:/1.jpg"
	_,err:=CopyFile(dst,src)
	if err == nil{
		fmt.Println("拷贝完成")
	}else{
		fmt.Println("拷贝失败")
	}
}