package main
import (
	"fmt"
	//cuser "lnhgo/day04/demo01/user"//导包使用别名
	_ "lnhgo/day04/demo01/user"//匿名导入，只是启动时调用包里init函数
)
func main(){
	// t:=Coures{//使用.导入使用方式
	// 	Name: "c",
	// }
	 fmt.Println("t")//调用方式

}