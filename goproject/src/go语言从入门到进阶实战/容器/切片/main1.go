package main
import (
	"fmt"
)
func main(){
	var highRiseBuilding [30] int
	for i := 0; i<30; i++{
		highRiseBuilding[i]= i+1
	}
	//区间
	fmt.Println(highRiseBuilding[10:15])
	//中间到尾部
	fmt.Println(highRiseBuilding[20:])
	//开头到中间
	fmt.Println(highRiseBuilding[:3])
}
/*
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\切片\main1.go
[11 12 13 14 15]
[21 22 23 24 25 26 27 28 29 30]
[1 2 3]


代码构建了一个30元素整数数组。值分别是1~30
打印区间数组的值
打印元素20以上的值
打印元素3以下的值
切片有点像C语言里的指针。指针可以做运算，但代价是内存操作越界。切片在指针的基础上增加了大小，约束了切片对应的内存区域，切片使用中无法对切片内部地址和大小进行手动调整，因袭切片币指针更安全强大
*/ 
