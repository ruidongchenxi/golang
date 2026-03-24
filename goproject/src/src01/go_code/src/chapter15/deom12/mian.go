package main
import (
	"fmt"
	"os"
)
func  main(){
	fmt.Println("命令行的参数个数有",len(os.Args))
	//遍历这个切片
	for i,v:= range os.Args{
		fmt.Printf("agrs[%v]=%v\n",i,v)
	}
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go build -o chapter15\deom12\test.exe chapter15\deom12\mian.go
// PS D:\golang\goproject\src\src01\go_code\src>  chapter15\deom12\test.exe tom  123 d:/1.txt 909
// 命令行的参数个数有 5
// agrs[0]=D:\golang\goproject\src\src01\go_code\src\chapter15\deom12\test.exe
// agrs[1]=tom
// agrs[2]=123
// agrs[3]=d:/1.txt
// agrs[4]=909