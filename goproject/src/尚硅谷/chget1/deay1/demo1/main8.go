package main
import (
    "fmt"
   // "math"
)

func main() {
    // sum := 0.1 + 0.2
    // fmt.Printf("未经舍入: %.17f\n", sum)
    // fmt.Printf("四舍五入后: %.2f\n", math.Round(sum*100)/100)
	a := 0.1 + 0.2
	if a == 0.3 {
    	fmt.Println("相等")
	} else {
    fmt.Println("不相等")  // 通常会进入这个分支
	}
	
}
