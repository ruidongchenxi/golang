package main

import (
	"fmt"
	//"slices"
	//"C"
)
func fbn (n int) ([]uint64){
	fbnSlice := make([]uint64,n)
	for i,_ := range fbnSlice{
		if i == 0 || i==1{
			fbnSlice[i]=1
		}else {
			fbnSlice[i]= fbnSlice[i-1] +fbnSlice[i-2] 
		}
	}
	return fbnSlice

}


func main(){
	fmt.Println("值=",fbn(30))

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter06\demo10\main.go
// 值= [1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946 17711 28657 46368 75025 121393 196418 317811 514229 832040]