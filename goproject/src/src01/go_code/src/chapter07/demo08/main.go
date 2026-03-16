package main

import (
	_ "container/list"
	"fmt"
	"math/rand"
)
func BinaryFind(arr *[10]int,leftIndex int,rightIndex int, find int) {
	if leftIndex > rightIndex {
		fmt.Println("不存在")
		return
	}
	middle:=(leftIndex+rightIndex)/2
	if (*arr)[middle] <find{
		//说明 find的取值范围在左边。left---middel-1
		BinaryFind(arr,leftIndex,middle-1,find)

	}else if (*arr)[middle]>find{
		//说明 find的取值范围在左边
		BinaryFind(arr,middle+1,rightIndex,find)
	} else  {
		fmt.Printf("找到了下标为%v",middle)
		
	}

}
func main(){
	var t  [10]int 
	for i,_ := range t{
		t[i]=rand.Intn(100)+1

	}
	for i:=0;i<len(t)-1;i++{
		for x:=0; x <len(t)-1-i;x++{
			if t[x] <t[x+1]{
				t[x],t[x+1]=t[x+1],t[x]
			}
		}
	}
	fmt.Println(t)
	var a int 
	for _,v := range t{
		a+=v
	}

	fmt.Println("平均值",a/len(t))
	fmt.Println("最大值",t[0])
	fmt.Println("最大值的下标",0)
	BinaryFind(&t,0,len(t),55)

}
// 执行结果
// [86 64 59 56 54 47 46 24 12 10]
// 平均值 45
// 最大值 86
// 最大值的下标 0
// 不存在