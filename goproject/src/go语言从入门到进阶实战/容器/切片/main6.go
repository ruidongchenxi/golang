package main
import (
	"fmt"
)
func main(){
	//声明一个整型切片
	var numbers []int
	for i := 0; i < 10; i++ {
		//循环向numbers切片添加10个数
		numbers = append(numbers,i)
		//打印输出切片的长度、容量和指针变化。使用len()函数查看切片拥有的元素个数，使用cap()函数查看切片的容量情况
		fmt.Printf("len: %d cap: %d pointer: %p\n",len(numbers),cap(numbers),numbers)
	}
}
/*
执行后结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\切片\main6.go                                  
len: 1 cap: 1 pointer: 0xc00000a0d8                                                                                     
len: 2 cap: 2 pointer: 0xc00000a120                                                                                     
len: 3 cap: 4 pointer: 0xc000014180                                                                                     
len: 4 cap: 4 pointer: 0xc000014180                                                                                     
len: 5 cap: 8 pointer: 0xc000012240                                                                                     
len: 6 cap: 8 pointer: 0xc000012240                                                                                     
len: 7 cap: 8 pointer: 0xc000012240                                                                                     
len: 8 cap: 8 pointer: 0xc000012240                                                                                     
len: 9 cap: 16 pointer: 0xc000074080                                                                                    
len: 10 cap: 16 pointer: 0xc000074080  
声明一个整型切片

*/