package main

import (
	"fmt"
)

 func main(){
	var t int = 97
	r := t / 7
	c := t % 7

	fmt.Printf("还剩%v个星期,零%v天\n",r,c) 
	var h float32 =134.2
	var s float32 = 5.0 / 9*(h-100)
	fmt.Printf("%v 对应的摄氏温度=%v\n",h,s)
	}

// 执行结果
// 还剩13个星期,零6天
// 134.2 对应的摄氏温度=19