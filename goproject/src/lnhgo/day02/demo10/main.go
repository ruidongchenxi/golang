package main
import (
	"fmt"
)
func d()(r int){
	defer func ()  {
		r ++
		
	}()
	return 10
}
func main(){
	fmt.Println(d())

}