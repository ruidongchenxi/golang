package main//package 进行包声明，建议包的声明和所在文件夹同名
//main 包是程序的入口包，一般main函数会放在这个包下
//包名路径是从$GOPATH/src后开始计算使用/进行路径分隔
//import "fmt"
//包名路径是从$GOPATH/src后开始计算使用/进行路径分隔
//import "gocode/testproject10/crm/db"
//如果有多个包，建议一次性导入
import (
	"fmt"
	test "gocode/testproject10/crm/db" //别名
)
func main(){
	fmt.Println("您好 gomain")
	//dbtset.GetConn()// 函数调用的时候前面要定位到所在的包
    test.GetConn()
}