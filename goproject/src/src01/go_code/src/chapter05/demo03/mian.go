package main
import (	
	"fmt"
)
func getsumandsub (n1 int,n2 int) (int,int){
	n3 := n1+n2
	n4 := n1-n2
	return n3 , n4
}
func main() {

	s2,s3 := getsumandsub(9,5)
	fmt.Printf("和:%v, 差:%v\n",s2,s3)
	fmt.Println(getsumandsub(12,67))//直接调用
	s4,_ :=getsumandsub(98,65)//忽略第二个返回值
	fmt.Printf("和:%v\n",s4)
	
}
// 执行结果
// 和:14, 差:4
// 79 -55
// 和:163