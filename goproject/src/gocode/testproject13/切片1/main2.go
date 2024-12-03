package main
import "fmt"
func main(){
	//
	slice := make([]int,4,20)
	slice[0] = 87
	slice[1] = 45
	slice[2] = 68
	slice[3] = 90
	//方式1
    for i :=0; i<len(slice);i++{
		fmt.Println(slice[i])
	}
	fmt.Println("----------------------------------------------------")
	for k,_ := range slice{
		fmt.Println(slice[k])
	}
}