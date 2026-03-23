package main

import (
	//"bufio"
	"fmt"
	//"io"
	//"io/ioutil"

	//"io/ioutil"
	"os"
)
func main(){
	//将d:/adc.txt文件内容导入d:/kkk.txt
	//首先将d:/adc.txt 内容读取到内存
	//2将读取到内容写入kkk.txt
	filepath1:= "D:/abc.txt"
	filepath2:= "D:/kkk.txt"
	data,err :=os.ReadFile(filepath1)
	if err !=nil{
		fmt.Println("错误：",err)
	}
	err =os.WriteFile(filepath2,data,0666)
	if err!=nil{
		fmt.Println("writeFile err=",err)
	}
}