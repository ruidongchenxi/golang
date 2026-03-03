package main

import (
	//"crypto/x509"
	"fmt"
)

func main() {
    fmt.Println("ok1")
	goto abc
    //fmt.Println("ok2")
    fmt.Println("ok3")
    fmt.Println("ok4")
    fmt.Println("ok5")
	abc:
    fmt.Println("ok6")
    fmt.Println("ok7")
    fmt.Println("ok8")
    fmt.Println("ok9")
    fmt.Println("ok10")
    fmt.Println("ok11")

}
// 执行结果

// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo26\main.go
// ok1
// ok6
// ok7
// ok8
// ok9
// ok10
// ok11