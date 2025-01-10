package main
import (
	"fmt"
)
func main(){
	str := "hello 您好"
	for key, value := range str {
		fmt.Printf("key: %d value:0x%x\n", key, value)
	}
}