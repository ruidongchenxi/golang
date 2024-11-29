package main
import "fmt"
func main(){
	for i := 1; i <=100; i++{
		fmt.Println(i)
		if i == 14 {
			return //直接结束函数
		}
	}
    fmt.Println("chenxi !")
}