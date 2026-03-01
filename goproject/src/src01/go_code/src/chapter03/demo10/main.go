package main

import (
	"fmt"
)

 func main(){
	var t int = 45
	//将t =>float32
	var ti float32 = float32(t)

	fmt.Printf("i=%v n1=%v\n",t,ti)
	var r int8 = int8(t)
	fmt.Printf("r=%v \n",r)
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=",num2)

}
// 执行结果
// i=45 n1=45
// r=45  
// num2= 63