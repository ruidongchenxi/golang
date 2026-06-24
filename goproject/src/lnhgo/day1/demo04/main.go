package main
import (
	"fmt"
)

// //批量声明常量
// const (
// 	n1 = iota
// 	n2 
// 	n3   
// )

// const (
// 	b1 = iota
// 	b2 
// 	_
// 	b3
// )
const (
		c1 = iota //0
		c2 = 100  //100
		c3 = iota //2
		c4        //3
	)
	const c5 = iota //0

func main(){
	// fmt.Println("n1:",n1)
	// fmt.Println("n2:",n2)
	// fmt.Println("n3:",n3)
	// fmt.Println("b1:",b1)
	// fmt.Println("b2:",b2)
	// fmt.Println("b3:",b3)
	fmt.Println("c1:",c1)
	fmt.Println("c2:",c2)
	fmt.Println("c3:",c3)
	fmt.Println("c4:",c4)
	fmt.Println("c5:",c5)


}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo04\main.go
// c1: 0
// c2: 100
// c3: 2
// c4: 3
// c5: 0