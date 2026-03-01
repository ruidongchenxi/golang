package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var b = false
	fmt.Println("b=",b)
	//1.bool 类型占用存储空间是1个字节
	fmt.Println("b的占用空间 =",unsafe.Sizeof(b))
	//2.bool 类型只能取true或者false
	


}
// 执行结果
// b= false
// b的占用空间 = 1