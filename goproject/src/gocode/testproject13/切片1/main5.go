package main
import "fmt"
func main(){
	//定义数组
	var intarr [6]int = [6]int{1,2,4,56,17,23}
	//定义切片
	var sclice []int = intarr[1:4]//2,4,56
	fmt.Println(len(sclice))
	sclice2 := append(sclice,89,45)
	sclice = append(sclice,89,45)
	fmt.Println(sclice2)
	fmt.Println(sclice)
	sclice3 := []int{99,44}
	sclice = append(sclice,sclice3...)
	fmt.Println(sclice)
}