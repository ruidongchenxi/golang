package main
import (
	"fmt"
)
func main(){
	//数组
	// var s1 [3]string=[3]string{"t","你好","小杜"}//数组初始化
	// var s2 [4]string
	// s2[0]="你好"
	// s1[0]="m"
	// for i,v:=range s2{
	// 	fmt.Println(i,v)
	// }
	//数组初始化
	// var R1 = [...]string{"wc","你好","小鹿"}
	// for _,v:=range R1{
	// 	fmt.Println(v)
	// }
	// R2:= [...]string{"wc","你好","小鹿"}
	// for _,v:=range R2{
	// 	fmt.Println(v)
	// }
	// R3:= [3]string{2:"小熊"}
	// for _,v:=range R3{
	// 	fmt.Println(v)
	// }
	// if R1==R2{
	// 	fmt.Println("exit")
	// }
	//初始化
	var str [3][4]string= [3][4]string{{"go","1h","bobby","go体系课程"},{"python","2h","橘子","python体系课程"},{"grpc","2H","小鹿","grpc"}}
	//遍历
	for _,v:=range str{
		for _,a:=range v{
			fmt.Println(a)
		}
	}
	

}