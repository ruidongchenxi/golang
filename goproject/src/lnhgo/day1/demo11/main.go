package main

import (
	"fmt"
	"math"
)

//fmt 占位符
func main() {
	fmt.Printf("%f\n",math.Pi)
	fmt.Printf("%.2f\n",math.Pi)
}
// 执行
// PS D:\golang\goproject\src\lnhgo> go run day1\demo11\main.go
// 3.141593
// 3.14