package main

import (
	"errors"
	"fmt"
) 
// func test(){
//     //使用defer + recover来处理和捕获异常
//     defer func() {
//         err := recover() //是内置函数,可以捕获异常
//         if err !=nil{ //说明捕获到错误
//             fmt.Println(err)

//         }
//     }()
//     num1 := 10
//     num2 := 0
//     res:= num1/num2
//     fmt.Println("res=",res)
// }
//函数去读init.conf
func readConf(name string) (err error){
    if name == "init.conf"{
        fmt.Println("启动")
        return nil
    }else{
        // 返回一个自定义错误
        return errors.New("指定文件错误")
    }

}
func test01() {
    err := readConf("init.conf")
    if err != nil {
        //如果读取文件发生错误，就输出这个错误并终止程序
        panic(err)
        
    }
    fmt.Println("test02()继续执行")
}
func main(){

        // test()
        // fmt.Println("下面的代码和逻辑..........")
        test01()
        
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter05\demo22\mian.go
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter05\demo22\mian.go
// panic: 指定文件错误  //文件不一致报错

// goroutine 1 [running]:
// main.test01()
//         D:/golang/goproject/src/src01/go_code/src/chapter05/demo22/mian.go:36 +0x76
// main.main()
//         D:/golang/goproject/src/src01/go_code/src/chapter05/demo22/mian.go:45 +0xf
// exit status 2
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter05\demo22\mian.go
// 启动
// test02()继续执行