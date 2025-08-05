package main

import (
	"fmt"
)

func main() {
	a := 0.1
	b := 0.2
	c := a + b

	fmt.Printf("0.1 + 0.2 = %.17f\n", c)
	if c == 0.3 {
		fmt.Println("c 等于 0.3")
	} else {
		fmt.Println("c 不等于 0.3")
	}
	var n1 float32 = -123.0000901
	var n2 float64 = -123.0000901
	fmt.Println("n1=", n1, "\nn2=", n2)
	var num5 = 1.1
	fmt.Printf("num5的数据", num5)

}
