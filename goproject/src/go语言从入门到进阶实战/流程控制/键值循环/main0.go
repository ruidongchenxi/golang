package main
import (
	"fmt"
)
func main(){
	for key, value := range [] int{1,2,3,4} {
		fmt.Printf("key:%d value:%d\n",key, value)
	}
}