package main
import (
	"fmt"
    "gocode/testproject04/test"
)
/*
D:\golang\goproject\src\gocode\testproject04\main>go run test.go                                                        
test.go:4:5: package gocode/testproject04/test is not in std (D:\go\src\gocode\testproject04\test)                      
                                                                                                                        
D:\golang\goproject\src\gocode\testproject04\main>go env -w GO111MODULE=off   处理方法          
*/
func main(){
	/*如果util.go中定义的是stuNo的话，那么在test.go中是访问不到，会报错：
	# command-line-arguments
.\test.go:14:22: undefined: test.stuNo
	*/
    fmt.Println(test.StuNo)
}