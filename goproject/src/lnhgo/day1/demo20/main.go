package main

import (
	"flag"
	"fmt"
)
//定义命令行参数
var mode = flag.String("mode","","process mode")
func main() {
	//解析命令行参数
	flag.Parse()
	//输出命令行参数
	fmt.Println(*mode)

}
// 执行结果
// PS D:\golang\goproject\src\lnhgo> go run day1\demo19\main.go --mode=fast
// fast