package main
import (
	"fmt"
	"unsafe"
)
func main(){
	w := 8
	fmt.Printf("w的类型%T\n",w)
	//unsafe.Sizeof(w)是unsafe包的一个函数可以返回w变量占用的字节数
	fmt.Printf("w的类型%T w占用字节数%d", w, unsafe.Sizeof(w))
}