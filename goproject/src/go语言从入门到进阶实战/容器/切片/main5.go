package main
import (
	"fmt"
)
func main(){
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a,b)
	fmt.Println(len(a), len(b))
}
/*执行后的结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\切片\main5.go                                  
[0 0] [0 0]                                                                                                             
2 2    
*/