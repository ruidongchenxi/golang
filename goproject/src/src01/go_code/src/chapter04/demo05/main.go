package main
import (
	"fmt"
	"math"
)
func main(){

	var a float64 = 3.0
	var w float64 = 100.0
	var i float64 = 6.0
	m := w * w - 4 * a * i
	if m > 0 {
		x1 := (-w + math.Sqrt(m))/2*a
		x2 := (-w - math.Sqrt(m))/2*a
		fmt.Printf("x1=%v x2=%v\n",x1,x2)
	}else if m == 0{
		x1 := (-w + math.Sqrt(m))/2*a
		fmt.Printf("x1=%v x2=%v\n",x1)
	} else {
		fmt.Println("无解")
	} 
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo05\main.go
// x1=-0.5409755150261972 x2=-299.4590244849738