package main
import (
	"fmt"
	"os"
	"bufio"

)
func main(){
	file , err := os.OpenFile("D:/Demo.txt",os.O_RDWR | os.O_APPEND | os.O_CREATE,0666)
	if err != nil{//文件失败
		fmt.Println("文件打开失败")
	}
	//及时关闭从操作
	defer file.Close()
	//写入文件操作：-----》IO流-----》 输出流
    writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
	writer.WriteString("你好，马士兵\n")
	}
	//流带缓存区
	writer.Flush()
}