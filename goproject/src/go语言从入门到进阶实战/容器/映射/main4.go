package main
import (
	"fmt"
)
func main(){
	scene := make(map[string]int)
	//准备
	scene["route"] = 66
	scene["brazil"] = 2
	scene["china"] = 960
	for k,v := range scene{
		fmt.Println(k,v)
	}
	fmt.Println("********************")
	delete(scene,"brazil")
	for k,v := range scene{
		fmt.Println(k,v)
	}
}
/*执行后结果

*/