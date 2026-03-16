package main
import (
	"fmt"
)

func main(){
	// var  t [4][6]int
	// // label1:
	// // for i,_:= range t{
	// // 	for s,_:= range t[i]{
	// // 		if i==1 && s==2 {
	// // 			t[i][s]=1
	// // 		}else if i ==2 && s ==1{
	// // 			t[i][s]=2
	// // 		}else if i ==2 && s ==3{
	// // 			t[i][s]=3
	// // 			//y = 1
	// // 			break label1

	// // 		}
	// // 	}
	// // }
	
	// for i,_:= range t{
	// 	for s,v:= range t[i]{
	// 		if i==1 && s==2 {
	// 			v=1
	// 		}else if i ==2 && s ==1{
	// 			v=2
	// 		}else if i ==2 && s ==3{
	// 			v=3

	// 		}
	// 		fmt.Printf("%v ",v)
	// 	}
	// 	fmt.Println()
	// }
	// for i,_:=range t{
	// 	for _,v:=range t[i]{
	// 		fmt.Printf("%v ",v)
	// 	}
	// 	fmt.Println()
	// 
	// var arr2 [2][3]int 
	// fmt.Printf("arr2[0]的地址值%p\n",&arr2[0])
	// fmt.Printf("arr2[1]的地址值%p\n",&arr2[1])
	// fmt.Printf("arr2[0][0]的地址值%p\n",&arr2[0][0])
	// fmt.Printf("arr2[1][0]的地址值%p\n",&arr2[1][0])
	
	arr3 := [2][2]int {{1,3},{6,9}}
	fmt.Println("arr3=",arr3)
	//初始化时就赋值
	var arr4 = [2][2]int {{1,5},{2,7}}
	fmt.Println("arr4=",arr4)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo04\main.go
// arr3= [[1 3] [6 9]]
// arr4= [[1 5] [2 7]]