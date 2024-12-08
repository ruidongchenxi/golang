package main
import (
	"fmt"
	//"os"
	//"bufio"
	"io/ioutil"

)
func main(){
	//定义原文件
	filePath := "D:/nfs-deployment.yaml"
	//定义目标文件
	filePath2 := "D:/nfs.yaml"
	//对文件进行读取
	content,err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取有问题")
	}
	//写出文件
	//filePath2 := "D:/nfs.yaml"
	err2 := ioutil.WriteFile(filePath2,content,0666)
	if err2 != nil {
		fmt.Println("写入失败")
	} 
}