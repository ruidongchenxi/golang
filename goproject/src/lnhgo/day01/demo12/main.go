package main

import (
	"fmt"
	//"path"
	"time"
)
func main(){
	for i:=0;i<3;i++{
		fmt.Println(i)
	}
	var i int
	for  {//死循环

		time.Sleep(2*time.Second)
		fmt.Println(i)
		i++
	}
}