package main
import(
	"fmt"
)
func b (n *[5]int) [5]int{
	for x:= len((*n))-1;x >1;x--{
		for i := 0; i<x ;i++{
			if (*n)[i] > (*n)[i+1]{
				(*n)[i],(*n)[i+1]=(*n)[i+1],(*n)[i]
			}
		}
	}
	return *n

}
func c (n *[5]int) [5]int{
	for i:=0;i<len(*n)-1;i++{
		for x:=0; x <len(*n)-1-i;x++{
			if (*n)[x] > (*n)[x+1]{
				(*n)[x],(*n)[x+1]=(*n)[x+1],(*n)[x]
			}
		}
	}
	return *n
}
func main(){
	a := [5]int{23,56,98,53,24}
	fmt.Println(a)
	for i :=len(a)-1; i > 1 ;i--{
		for x:=0; x < i ;x++{
			if a[x]> a[x+1]{
				a[x],a[x+1]=a[x+1],a[x]
			}
			
		}
		fmt.Println("小圈打印交换后",a)
	}
	fmt.Println("处理",a)
	a1 := [5]int{23,56,98,53,24}
	fmt.Println("a1,操作前",a1)
	//a1 = b(&a1)
	b(&a1)
	fmt.Println("a1=",a1)
	a2 := [5]int{23,56,98,53,24}
	fmt.Println("a2=",a2)
	c(&a2)
	fmt.Println("a2=",a2)

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo01\main.go
// [23 56 98 53 24]
// 小圈打印交换后 [23 56 53 24 98]
// 小圈打印交换后 [23 53 24 56 98]
// 小圈打印交换后 [23 24 53 56 98]
// 处理 [23 24 53 56 98]
// a1,操作前 [23 56 98 53 24]
// a1= [23 24 53 56 98]
// a2= [23 56 98 53 24]
// a2= [23 24 53 56 98]