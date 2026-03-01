package main

import (
	"fmt"
	"strconv"
)

 func main(){
	// var str string = "true"
	// var w1 bool
	// //说明
	// // strconv.ParseBool 会返回2个值，用一个接收是不行的，所以使用'_'空接收不要的值
	// w1,_ = strconv.ParseBool(str)
	// fmt.Printf("b type %T b 的值%v\n",w1,w1)
	// var str2 string = "1234590"
	// var w2 int64
	// w2,_ = strconv.ParseInt(str2,10,64)
	// fmt.Printf("w2 type %T w2=%v\n",w2,w2)
	// var str3 string = "3.14256"
	// var w3 float64
	// w3,_ = strconv.ParseFloat(str3,64)
	// fmt.Printf("w3 type %T w3=%v\n",w3,w3)
	var str4 string = "hello"
	var n3 int64 
	n3,_=strconv.ParseInt(str4,10,64)
	fmt.Printf("n2 type %T n2=%v\n",n3,n3)

}
// 执行结果
// n2 type int64 n2=0