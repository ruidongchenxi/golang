package main
import (
	"fmt"
	"os"
)
func main(){
	//打开文件
	file,err := os.Open("D:/pvc.yaml");
    if err != nil {
		fmt.Println("文件打开出错，对应错误为",err)
	}
	fmt.Printf("打开 yaml文件=%v",&file)
	//err := file.Close();
	if file.Close() != nil {
		fmt.Println("关闭失败")
	}
	fmt.Println("失败")
}