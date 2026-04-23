package main

import (
	"fmt"
	"math/rand"
	"time"
)
func InsertSort(arr *[80000]int){
	// insertVal:=arr[1]
	// insertIdex:=1-1//下标
	for i :=1;i<len(arr);i++{
		insertVal:=arr[i]
		insertIdex:=i-1//下标
		for insertIdex>=0&& arr[insertIdex]<insertVal{
		// if arr[insertIdex]< insertVal{
		// 	arr[insertIdex]=arr[insertIdex+1]
		// }
			arr[insertIdex+1]=arr[insertIdex]
			insertIdex--
		}
		if insertIdex+1!=insertIdex{
		arr[insertIdex+1]=insertVal
		}
		//fmt.Printf("第%d次,%d\n",i,*arr)
	}
}
func main(){
  //arr:=[5]int{23,0,12,56,34}
  var arr [80000]int
	for i:=0;i<80000;i++{
		// arr[i]=rand.Intn(9000000)
		arr[i]=rand.Intn(900000)
	}
	start:=time.Now().Unix()
  	InsertSort(&arr)
	end:= time.Now().Unix()
	fmt.Printf("耗时=%d秒",end-start)

}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter20\demo09\main.go
// 耗时=2秒
