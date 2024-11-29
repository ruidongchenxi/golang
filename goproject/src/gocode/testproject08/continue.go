package main
import "fmt"
func main(){
	//功能：输出1-100中被6整除的数
	//方法1
	/*
	for i := 1; i < 100; i++ {
	   if i % 6 == 0 {
		fmt.Println(i)
	   }
	} 
	*/
	for i := 1; i < 100; i++ {
		if i % 6 !=0{
			continue //跳出这一次循环;结束本次
		}
		fmt.Println(i)
	}
}