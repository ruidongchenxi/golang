package main

import (
	"fmt"
	//"testing/quick"
	"math/rand"
	"time"
)
//最左边下标left
//右边下标right
func QuickSort(left int,right int,array *[80000]int){
	l := left
	r := right
	//中轴
	pivot := array[(left+right)/2]
	//temp:=0
	//比pivot 小的数放到左边
	//比pivot 大的数放到右边

	for;l<r;{
		for ; array[l]<pivot;{
			l++
		}
		for ;array[r]>pivot;{
			r--
		}
		//表示
		if l>=r{
			break
		}
		// temp = array[l]
		// array[l]= array[r]
		// array[r]=temp
		array[l],array[r]=array[r],array[l]
		if array[l]==pivot{
			r--
		}
		if array[r]==pivot{
			l++
		}
	}

	if l==r{
		l++
		r--
	}
	//向做
	if left<r{
		QuickSort(left,r,array)
	}
	if right>l{
		QuickSort(l,right,array)
	}


}
func main(){
//   arr:=[6]int{-9,78,0,23,-567,70}
 	var arr [80000]int
	for i:=0;i<80000;i++{
		// arr[i]=rand.Intn(9000000)
		arr[i]=rand.Intn(900000)
	}
	start:=time.Now().Unix()
	QuickSort(0,len(arr)-1,&arr)
	end:= time.Now().Unix()
	fmt.Printf("耗时=%d秒",end-start)
  
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter20\demo10\main.go
// 耗时=0秒