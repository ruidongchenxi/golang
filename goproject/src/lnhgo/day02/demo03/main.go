package main

import (
	"fmt"
	"strconv"
)
func printSlice(d []string){
	d[0]="java"
	for i:=0;i<10;i++{
		d=append(d, strconv.Itoa(i))
	}
	fmt.Println(d)
}
func main(){
	c := []string{"go","grpc","gin"}
	printSlice(c)
	fmt.Println("-------------------------")
	fmt.Println(c)
}