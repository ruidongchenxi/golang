package main
import (
	"fmt"
)
func main(){
	var t1 int32 = 32
	var t2 int32 = 23
	
	if ( t1 + t2 ) >= 50{
		fmt.Println("hello world")
	}
	var f1 float32 = 10.5
	var f2 float32 = 8.6

	if f1 > 10.0 && f2 <20 {
		fmt.Println("和=",(f1+f2))
	}
	var c1 int32 = 65
	var c2 int32 = 25

	if ( c1 + c2 ) % 3 == 0 && (c1 + c2) % 5 == 0{
		fmt.Println((c1+c2),"能被整除")
	}
	var s int = 2026

	if (s %4 ==0 &&s % 100 !=0) || s % 400 == 0 {
		fmt.Println(s,"是闰年")
	} else {
		fmt.Println(s,"不是闰年")
	}
	

}
// 执行结果
// hello world
// 和= 19.1
// 90 能被整除
// 2026 不是闰年