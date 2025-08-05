package main
import (
	"fmt"
	"flag"
)
var skillParam = flag.String("skill","","skill to perform")
func main(){
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok{
		f()
	} else{
		fmt.Println("skill not found")
	}
}
/*
执行结果
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\匿名函数\main1.go --skill=run
soldier run
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\匿名函数\main1.go --skill=fly
angel fly
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\匿名函数\main1.go --skill=fire
chicken fire  
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\匿名函数\main1.go --skill=f   
skill not found
PS D:\golang\goproject> go run D:\golang\goproject\src\go语言从入门到进阶实战\函数\匿名函数\main1.go          
skill not found
*/