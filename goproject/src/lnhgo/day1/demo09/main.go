package main
import (
	"fmt"
)

func main() {
    var r int 
	var t int
	fmt.Printf("输入有一个数：")
	fmt.Scanln(&t)
	for t != 1{
		if t %2 == 0 {
			fmt.Println(t)
			t= t / 2
		    if t > r{
				r =t
			}
			//fmt.Println(r)
		}else {
			fmt.Println(t)
			t=t *3 +1 
			if t > r{
				r =t
			}
			//fmt.Println(r)
			
		}
	}
	fmt.Println(t)
	fmt.Println(r)
}
