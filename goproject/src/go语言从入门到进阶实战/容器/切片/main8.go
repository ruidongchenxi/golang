package main
import (
	"fmt"
)
func main(){
	//设置元素个数1000
	const elementCount = 1000
	//预分配足够多的切片元素
	srcData := make([]int, elementCount)
	//将切片赋值
	for i := 0; i < elementCount;i++{
		srcData[i] = i
	}
	//引用切片数据
	refData := srcData
	//预分配足够多的元素切片
	copyData := make([]int,elementCount)
	//将数据复制到新的切片空间中
	copy(copyData,srcData)
	//修改原始数据第一个元素
	srcData[0] = 999
	//打印引用切片的第一个元素
	fmt.Println(refData[0])
	//打印复制切片第一个元素和最后一个元素
	fmt.Println(copyData[0],copyData[elementCount-1])
	//复制原始从4到6
	copy(copyData,srcData[4:6])
	for i :=0; i < 5;i++{
		fmt.Printf("%d",copyData[i])
	}
}
/*
执行结果
C:\Users\尘曦>go run D:\golang\goproject\src\go语言从入门到进阶实战\容器\切片\main8.go
999
0 999
45234
代码说明
定义元素数量1000
预分配拥有1000个元素的整型切片，这个切片将作为原始数据
将srcData填充0~999的整型值
将refData引用srcData，切片不会因为等号操作进行元素的复制
预分配语srcDate等大、同类型的切片copyData
使用copy()函数将原始数据复制到copyData切片空间中
修改原始数据第一个元素为999
引用数据的第一个元素也会发生变化
打印复制数据的首位数据，由于数据是复制的，因此不会发生变化
将srcData的局部数据复制到copyData中
打印复制局部数据后的copyData元素
*/