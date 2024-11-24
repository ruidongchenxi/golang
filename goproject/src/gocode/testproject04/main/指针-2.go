package main
import (
	"fmt"
	//"strconv"
)
func main(){
    var num int = 10
    fmt.Println(num)
	//& 符号一定是地址值；指针变量接收的一定是地址值
	//指针变量的地址值不可以不匹配
	var prt *int = &num
	//修改指针得值
	*prt = 20
	fmt.Println(num)

 } 