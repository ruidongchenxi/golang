package main
import "fmt"

func main(){
	var i int32
    var j int64
	i, j = 1, 2
	if i == j { // 编译错误，类型不同不能比较
	 fmt.Println("i and j are equal.")
	}
	if i == 1 || j == 2 { // 编译通过
 	 fmt.Println("i and j are equal.")
	}
}