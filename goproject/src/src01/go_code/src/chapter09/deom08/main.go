package main
import (
	"fmt"
)
type Monster int 
func (m Monster)test1(){
	for i:=0; i<10; i++{
		for x:=0;x<8;x++{
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
func (d Monster)test2(m int,n int){
	for i:=0;i<m;i++{
		for x:=0;x<n;x++{
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
func (c Monster)test3(m int,n int) int{
	return m*n
}
func (t Monster)test4(s [3][3]int) {
	for i:=0;i<3;i++{
		for x:=i;x<3;x++{
			s[i][x],s[x][i]=s[x][i],s[i][x]
			//fmt.Println(s[i][x],s[x][i])
		}
	}
	fmt.Println(s)
	
}
func main(){
	var e Monster
	fmt.Printf("长5宽4的面积是=%v\n",e.test3(5,4))
	t := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(t)
	e.test4(t)

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter09\deom07\main.go
// 长5宽4的面积是=20
// [[1 2 3] [4 5 6] [7 8 9]]
// [[1 4 7] [2 5 8] [3 6 9]]