package main
import "fmt"
func main(){
	//定义切片
	// 
	var intarr [6]int = [6]int{1,2,3,4,7,9}
		//定义切片
		var sclice []int = intarr[1:4]
		fmt.Println(sclice)
		fmt.Println(sclice[0])
		fmt.Println(sclice[1])
		fmt.Println(sclice[2])
		//fmt.Println(sclice[3])//下标越界*/
		//对切片再切片
		sclice2 := sclice[1:2]
		fmt.Println(sclice2)
		sclice2[0] = 67
		fmt.Println(sclice[1])
		fmt.Println(sclice2[0])
}