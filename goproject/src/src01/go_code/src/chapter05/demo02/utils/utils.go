package utils
import (
	"fmt"
)
// 
func Cal(s int,x int, c string)  (int){ // 函数首字母必须大写
	var d int 
	if c == "*" {
		// fmt.Println(s,c,x,"=",s*x)
		d =s*x
	} else if c == "/"{
		// fmt.Println(s,c,x,"=",s/x)
		d = s/x
	} else if c == "+"{
		// fmt.Println(s,c,x,"=",s+x)
		d = s+x
	}else if c == "-"{
		// fmt.Println(s,c,x,"=",s-x)
		d = s - x
	}else if c == "%" {
		//fmt.Println(s,c,x,"=",s%x)
		d = s % x
	}else{
		fmt.Println("操作符错误")
	}
	return d
}

