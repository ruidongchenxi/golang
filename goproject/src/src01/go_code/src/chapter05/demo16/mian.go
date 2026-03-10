package main

import "fmt"

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

func main() {

	test([][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	})

}