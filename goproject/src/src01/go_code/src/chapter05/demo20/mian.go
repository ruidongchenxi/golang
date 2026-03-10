package main

import (
	"fmt"
	"strconv"
	"time"
) 
func test03(){
    str := ""
    for i:=0;i < 100000;i++{
        str += "hello"+strconv.Itoa(i)
    }
}


func main(){
 //在执行test03前；先获取当前的unix时间戳
 start := time.Now().Unix()
 test03()
 stop := time.Now().Unix()
 fmt.Printf("执行test时间%v秒\n",stop-start)
 
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter05\demo20\mian.go
// 执行test时间9秒