package main
import (
	"fmt"
	_"unicode/utf8"
)
func main(){
	theme := "狙击 start"
	for _, s := range theme {
		fmt.Printf("Unicode: %c %d\n", s,s)
	}
}