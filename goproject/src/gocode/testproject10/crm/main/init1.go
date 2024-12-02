package main
import (
	"fmt"
	"gocode/testproject10/crm/db"
)
var num int = test1()
func test1() int{
	fmt.Println("test 函数被执行")
	return 10
}
func init() {
	fmt.Println("main 包里的，init 函数被执行")
}
func main(){
	fmt.Println("dbtset 函数被执行")
	fmt.Println("Age=",dbtset.Age)
}