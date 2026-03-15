package main
import (
	"fmt"
)
func BinaryFind(arr *[10]int,leftIndex int,rightIndex int, find int) {
	if leftIndex > rightIndex {
		fmt.Println("不存在")
		return
	}
	middle:=(leftIndex+rightIndex)/2
	if (*arr)[middle] >find{
		//说明 find的取值范围在左边。left---middel-1
		BinaryFind(arr,leftIndex,middle-1,find)

	}else if (*arr)[middle]<find{
		//说明 find的取值范围在左边
		BinaryFind(arr,middle+1,rightIndex,find)
	} else  {
		fmt.Printf("找到了下标为%v",middle)
		
	}

}
func main(){
	var s [10]int = [10]int {10,15,23,28,38,45,55,67,79,89}
	 var r int 
	fmt.Println("输入一个数字")
	fmt.Scanln(&r)
	BinaryFind(&s,0,len(s)-1,r)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo03\main.go
// 输入一个数字
// 76
// 不存在
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo03\main.go
// 输入一个数字
// 38
// 找到了下标为4