package main

import (
	"fmt"

	//"honnef.co/go/tools/simple/s1007"
)
func main(){
	for i:=1;i<=9;i++{
		for x :=1;x<=i;x++{
			fmt.Printf("%d*%d=%d\t",x,i,x*x)
		}
		fmt.Println()
	}
	s1 := "Hello"
	s2 := "沙河"
	for _,v := range s1{
		fmt.Printf("%c\t",v)
	}
	for _,v := range s2{
		fmt.Printf("%c\t",v)
	}
}
/*
PS D:\golang\goproject\src\lnhgo> go run day2\demo03\main.go
1*1=1
1*2=1   2*2=4
1*3=1   2*3=4   3*3=9
1*4=1   2*4=4   3*4=9   4*4=16
1*5=1   2*5=4   3*5=9   4*5=16  5*5=25
1*6=1   2*6=4   3*6=9   4*6=16  5*6=25  6*6=36
1*7=1   2*7=4   3*7=9   4*7=16  5*7=25  6*7=36  7*7=49
1*8=1   2*8=4   3*8=9   4*8=16  5*8=25  6*8=36  7*8=49  8*8=64
1*9=1   2*9=4   3*9=9   4*9=16  5*9=25  6*9=36  7*9=49  8*9=64  9*9=81
H       e       l       l       o       沙      河
*/