package main
import (
	"fmt"
)
func main(){
   var  cx float32 = 3.14
   fmt.Println(cx)
   var  cx1 float32 = -3.14
   fmt.Println(cx1)
   var  cx2 float32 = 3.14E-2
   fmt.Println(cx2)
   var  cx3 float32 = 3.14E+2
   fmt.Println(cx3)
   var  cx4 float32 = 3.14e+2
   fmt.Println(cx4)
//浮点数可能会有精度损失
   var  cx5 float32 = 125.000000045
   fmt.Println(cx5)/*
   
D:\golang\goproject\src\gocode\testproject03\man>go run float32.go
3.14
-3.14
0.0314
314
314
125
125.000000000056
   */
   var  cx6 float64 = 125.000000000056
   fmt.Println(cx6)
   var cx7 = 8.0000965499008
   fmt.Printf
}