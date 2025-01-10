package main
import (
	"fmt"
)
func main(){
	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}
}