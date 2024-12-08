package main
import (
	"fmt"
	"os"
	"bufio"
	"io"

)
func main(){
    file,err := os.Open("D:/pvc.yaml")
    if err != nil{
		fmt.Println("文件打开失败",err)
	}
	//当函数退出时，让file 关闭防止内存泄漏
	defer file.Close()
	//创建一个流
	reader := bufio.NewReader(file)
    //读取操作
	for {
	str,err1 := reader.ReadString('\n')//字符单引号
	    if err1 == io.EOF {//io.EOF 表示已经读取到文件结尾
			break//终止循环
		}
		fmt.Println(str)
	}
}