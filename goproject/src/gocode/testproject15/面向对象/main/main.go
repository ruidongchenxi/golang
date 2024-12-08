package main
import (
	"fmt"
	cx "gocode/testproject15/面向对象/model"
)

func main(){
	//跨包访问
	//var s cx.Student = cx.Student{"丽丽",10}
	//s :=  cx.Student{"丽丽",10}
	//fmt.Println(s)
	s := cx.NewStudent("小龙",17)
	fmt.Println(*s)
}
