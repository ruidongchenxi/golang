package main
import (
	"fmt"
	
) 
func main(){
    num := 100
    fmt.Printf("num的类型=%T,num的值=%v,num的地址值=%v\n",num,num,&num)
    num1 :=new(int)
    fmt.Printf("num1的类型=%T,num1的值=%v,num1的地址值=%v,num1指针指向的的值是%v\n",num1,num1,&num1,*num1) 
    *num1 = 40//改变值
    fmt.Printf("num1的类型=%T,num1的值=%v,num1的地址值=%v,num1指针指向的的值是%v\n",num1,num1,&num1,*num1) 
 
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter05\demo21\mian.go
// num1的类型=*int,num1的值=0xc00000a0e0,num1的地址值=0xc000058030,num1指针指向的的值是0
// num1的类型=*int,num1的值=0xc00000a0e0,num1的地址值=0xc000058030,num1指针指向的的值是40



