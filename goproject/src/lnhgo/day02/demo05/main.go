package main

import (
	"fmt"
)
func main(){
	var a []int
	for i:=0;i<20;i++{
		a=append(a, i)
		fmt.Printf("len:%d,cap:%d\n",len(a),cap(a))
	}
}