package main

import (
	//"crypto/cipher"
	"math/rand"
	"time"
	//"crypto/rand"
	"fmt"
	//"github.com/redis/go-redis/v9/helper"
)

func SelectSort(arr *[80000]int){
	//arr[1]=600
	for x:=0;x<len(arr)-1;x++{
		max :=arr[x]
		maxIndex :=x
		for i:=x+1;i<len(arr);i++{
			if max < arr[i]{
				max=arr[i]
				maxIndex=i
				// arr[maxIndex]=arr[i]
			}
		}
		if maxIndex != x{
			arr[x],arr[maxIndex] = arr[maxIndex],arr[x]
		}
		//fmt.Printf("第%d次交换结构结果;%d;\n",x,*arr)
	}
}
func main(){
	//arr := [5]int{10,34,19,100,80}
	var arr [80000]int
	for i:=0;i<80000;i++{
		// arr[i]=rand.Intn(9000000)
		arr[i]=rand.Intn(900000)
	}
	start:=time.Now().Unix()
	SelectSort(&arr)
	end:= time.Now().Unix()
	fmt.Printf("耗时=%d秒",end-start)
	//fmt.Println(arr)
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter20\demo08\main.go
// 耗时=3秒