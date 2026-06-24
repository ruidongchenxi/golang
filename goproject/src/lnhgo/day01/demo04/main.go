package main

import "fmt"

func a() (int, bool) {
	return 0, true
}
func main() {
	//
	var _ int
	_, r := a()//匿名变量使用场景；占位符_
	if r {
		fmt.Println("a")
	}

}