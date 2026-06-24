package main
import(
	"fmt"
)
func main(){
	// var a11 = [...][3]int{{5,7},{7,75,8}}
	// var a12 [5][2]int
	// a12=[5][2]int{
	// 	{2,8},
	// 	{5,8},
	// 	{7,4},
	// 	{9,3},
	// 	{8,12},
	// }
	// for i,v:=range a11{
	// 	for x,j:=range v{
	// 		fmt.Printf("a11[%d][%d]=%d\t",i,x,j)
	// 	}
	// 	fmt.Println()
	// }
	// 	for i,v:=range a12{
	// 	for x,j:=range v{
	// 		fmt.Printf("a12[%d][%d]=%d\t",i,x,j)
	// 	}
	// 	fmt.Println()
	// }
	a1 := [...]int{1,3,5,7,8}
	for i,v:=range a1{
		t := 8 -v
		for x:=i+1; x<len(a1);x++{
			if t ==a1[x]{
				fmt.Println(i)
				fmt.Println(x)
			}
		}
	}
}