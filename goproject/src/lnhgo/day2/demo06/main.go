package main

import (
	"fmt"
	//"slices"
)
func main(){
	var a1 [3]bool
	var a2 [4]bool
	//数组的初始化；如果不初始化，默认元素都是0值
	fmt.Println(a1,a2)
	a1[0]=true
	a1[1]=true
	a1[2]=false
	a2[0]=true
	a2[1]=true
	a2[2]=true
	a2[3]=false
	fmt.Println(a2,a1)
	a100:=[5]int{1,3,56,89,7}
	for v,i:=range a100{
		fmt.Printf("a100[%d]=%d \n",v,i)
	}
	fmt.Println()
	a5 :=[...]int{3,5,8,2}
	for v,i:=range a5{
		fmt.Printf("a5[%d]=%d \n",v,i)
	}
	fmt.Println()
	a7 :=[5]int{0:2,4:8}
	for v,i:=range a7{
		if v<len(a7){
		fmt.Printf("a7[%d]=%d \t",v,i)
		}else{
			fmt.Printf("a7[%d]=%d \n",v,i)
		}
	}

	
}