package main
import "fmt"
func main(){
	fmt.Println("hello golang1")
	fmt.Println("hello golang2")
	fmt.Println("hello golang3")
	if 1 == 1 {
	    goto lable1
	}
	fmt.Println("hello golang4")
	fmt.Println("hello golang5")
	fmt.Println("hello golang6")
	lable1:
	fmt.Println("hello golang7")
	fmt.Println("hello golang8")
	fmt.Println("hello golang9")
    
}