package main
import (
	"fmt"
)
func main(){
	if err := Connect(); err !=nil{
		fmt.Println(err)
	}
}