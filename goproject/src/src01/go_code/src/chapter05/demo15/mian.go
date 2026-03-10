package main

import (
	"fmt"
	) 

func test(t [][]int) {
	s := make([][]int, 3)
	for i := 0; i < 3; i++ {
		s[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for x := 0; x < 3; x++ {
			s[i][x] = t[x][i]
		}
	}
	fmt.Println(s)

}
func main(){


	test([][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	})
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter05\demo15\mian.go
// [[1 4 7] [2 5 8] [3 6 9]]