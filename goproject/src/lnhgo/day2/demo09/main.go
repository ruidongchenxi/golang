package main
import(
	"fmt"
)
func main(){
	var s1 []int//定义切片
	//var s2 []int
	fmt.Println(s1==nil)
	s1=make([]int,2,3) //初始化
	s1=append(s1,5,6)
	s1=append(s1, 9,7,19)
	fmt.Println(s1[:])
	s1[0]=54
	s1[1]=18
	fmt.Println(s1[:],cap(s1))
	fmt.Println(s1)
	var s2 []string
	s2=[]string{"长沙","张江","平山村"}
	fmt.Println(s2)
}