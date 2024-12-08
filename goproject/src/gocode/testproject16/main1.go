package main
import (
	"fmt"
	"io/ioutil"
)
func main() {
	//备注： 在下面的过程中不需要进行open/Close操作，因为文件的打开和关闭操作被封装在readFile函数内部了
	//读取文件
	content,err := ioutil.ReadFile("D:/pvc.yaml")//返回内容为：[]byte，err
	if err != nil {//读取有误
		fmt.Println("读取错误",err)
	}
	//将打开的文件以字符串形式显示在终端
	fmt.Printf("%v",string(content))
}