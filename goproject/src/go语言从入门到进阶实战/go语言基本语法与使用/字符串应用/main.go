package main
import (
	"fmt"
)
func main(){
	tip1 := "genji is a ninja"
	fmt.Println(len(tip1))
	tip2 := "忍者"
	fmt.Println(len(tip2))
	fmt.Println(utf8.RuneCountInString("忍者"))
}