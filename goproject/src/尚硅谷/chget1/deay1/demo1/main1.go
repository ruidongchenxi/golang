package main

import "fmt"

func main() {
    a := 10
    b := 20
    sum := a + b
    fmt.Println("整数相加：", sum) // 输出 30

    x := 3.14
    y := 2.71
    fmt.Println("浮点数相加：", x + y) // 输出 5.85

    c := complex(1, 2)
    d := complex(3, 4)
    fmt.Println("复数相加：", c + d) // 输出 (4+6i)
}
