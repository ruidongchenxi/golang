package main

import (
	"fmt"
	"demo14/model"
	_ "strconv"

	//
	// 引入包
	//"chapter03/demo14/model"
)

 func main(){
	//golang 中严格区分大小写
	var num int = 20
	var Num int = 45
	fmt.Printf("num=%v Num=%v\n",num,Num)
    fmt.Println(model.HeroName)

}
// 执行结果
// D:\golang\goproject\src\src01\go_code\src\chapter03\demo14>go mod tidy

// D:\golang\goproject\src\src01\go_code\src\chapter03\demo14>go run ./main
// num=20 Num=45
// 松江
