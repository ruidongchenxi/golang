package main
import (
	"fmt"
	"time"
)
func main(){
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	// for i:= 0; i< 3; i++ {

	// 	go func(n int) {
	// 		fmt.Println(n)
	// 	}(i)
	// }
	time.Sleep(time.Second)
}