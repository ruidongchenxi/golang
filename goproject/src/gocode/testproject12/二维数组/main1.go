package main
import "fmt"
func main(){
    var arr1 [2][3]int = [2][3]int{{1,2,4},{7,8,9}}
	fmt.Println(arr1)
	//方式1
	for i := 0;i<len(arr1);i++{
		for x :=0;x<len(arr1[i]);x++{
			fmt.Println(arr1[i][x])
		}
	}
	//方式2
	for key,value := range arr1 {
		for k,v := range value {
			fmt.Printf("arr[%v]%v]=%v\n",key,k,v)
		}
	}
}