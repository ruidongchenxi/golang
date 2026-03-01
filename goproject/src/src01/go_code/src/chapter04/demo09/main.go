package main

import (
	"fmt"
	//"math"
)

// func test (chra byte) byte{
// 	return chra + 1
// }


func main(){

//    var n1 int32 = 20
//    var n2 int32 = 20
  // var n2 int64 = 20 // 报数据类型不匹配 invalid case n2 in switch on n1 (mismatched types int64 and int32)
//    switch n1 {
//    		case n2, 10, 5:// 
// 			fmt.Println("ok1")
//     	// case 5:
// 		// 	fmt.Println("ok2") //duplicate case 5 (constant of type int32) in expression switch
// 		default :
// 			fmt.Println("没有匹配到")
//    }
//    var age int = 10
//    switch {
//    case  age == 10:
// 		fmt.Println("age=10")
//    case age == 20:
// 		fmt.Println("age== 20")
//    default:
// 		fmt.Println("没有匹配到")
//    }
   var score int = 30
   switch{
   case score >90:
		fmt.Println("优秀")
   case score >=70 && score <=90:
		fmt.Println("优良")
   case score >=60 && score <70 :
		fmt.Println("成绩及格")
	default : 
		fmt.Println("不及格")
   }
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo09\main.go
// b
// 穿衣服d
