package main
import (
	"fmt"
)
func main(){
	//if判断
	age:=22
	if age<18{
		fmt.Println("未成年")
	}else{
		fmt.Println("成年")
	}
	if age<18{
		fmt.Println("未成年")
	}else if age>60{
		fmt.Println("老年人")
	}else{
		fmt.Println("青年")
	}
}