package main
import (
	"fmt"
	"src/chapter09/deom12/model"//引包
)
func main(){
	//实例化student结构体
	// var stu = model.Student{Name: "晨曦",Score: 90}
	var stu = model.NewStudent("tom~",88.8)
	//fmt.Println(stu.Name,stu.score) 因为结构体变量score首字母小写不可以直接导出
	//所以在model包里定义一个函数将他导出
	fmt.Println(stu.Name,model.Newget(stu))
	//使用方法调
	fmt.Println(stu.Test())

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter09\deom12\main\main.go
// tom~ 88.8
// 88.8