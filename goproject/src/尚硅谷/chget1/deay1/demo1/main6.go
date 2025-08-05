package main

import (
	"fmt"
)

func main() {
	var a float32 = 3.14159  // 声明一个 float32 类型变量
	var b float64 = 3.141592653589793  // 声明一个 float64 类型变量

	// 通过短变量声明方式让编译器自动推断类型为 float64
	c := 2.71828

	fmt.Printf("a (float32): %f\n", a)
	fmt.Printf("b (float64): %f\n", b)
	fmt.Printf("c (默认 float64): %f\n", c)
}
