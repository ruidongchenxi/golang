package main
import (
	"fmt"
)
func main(){
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	for k,v := range team{
		fmt.Println(k,v)
	}
}
/*
执行结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\数组\main1.go
0 hammer
1 soldier
2 mum
*/