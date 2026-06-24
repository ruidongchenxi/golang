package main
import (
	"fmt"
)

//批量声明常量
const (
	n1 = iota
	n2 
	n3   
)

const (
	b1 = iota
	b2 
	_
	b3
)

func main(){
	fmt.Println("n1:",n1)
	fmt.Println("n2:",n2)
	fmt.Println("n3:",n3)
	fmt.Println("b1:",b1)
	fmt.Println("b2:",b2)
	fmt.Println("b3:",b3)

}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo03\main.go
// n1: 0
// n2: 1
// n3: 2
// b1: 0
// b2: 1
// b3: 3