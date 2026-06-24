package main
import(
	"fmt"
)
func main(){
	// var sun int
	// for i:=1;i<=100;i++{
	// 	sun+=i
	// 	//fmt.Println(i,sun)
	// }
	// fmt.Println(sun)
	// for i:=1;i<10;i++{
	// 	for x:=1;x<=i;x++{
	// 		fmt.Printf("%d * %d = %d\t",x,i,x*i)
	// 	}
	// 	fmt.Println()
	// }
	//字符串遍历
	// name :="imooc go python 课程"
	// for v,k:=range name{
	// 	fmt.Printf("%d=%c\t",v,k)
	// }
	// r:=10
	// for {
	// 	r--
	// 	if r == 5{
	// 		goto t1//直接跳转到t1标签语句哪里
	// 	}else if r==0{
	// 		fmt.Println(r)
	// 		break //退出循环
	// 	}
	// 	fmt.Println(r)
		
	// }
	// t1://标签
	// fmt.Println("goto 标签")
	// var n int 
	// fmt.Printf("输入选择")
	// fmt.Scanln(&n)
	// switch n{
	// case 1:
	// 	fmt.Println("大拇指")
	// case 2:
	// 	fmt.Println("食指")
	// case 3:
	// 	fmt.Println("中指")
	// case 4:
	// 	fmt.Println("无名指")
	// case 5:
	// 	fmt.Println("小拇指")
	// default:
	// 	fmt.Println("输入不正确")
	// }
	var t int 
	fmt.Printf("输入分数：")
	fmt.Scanln(&t)
	switch {
	case t==100:
		fmt.Println("卓越")
	case t>90:
		fmt.Println("优秀")
	case t>80:
		fmt.Println("好")
	case t>70:
		fmt.Println("良")
	case t>60 && t==60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格，继续努力")
	}
}