package main

import "fmt"

func modify(map1 map[int]int){
	map1[10] =900
}
type sty  struct{
	name string
	age  int 
	add string
}
func main(){
	students := make(map[string]sty,10)
	//创建
	stu1 := sty{"tom", 29, "天通苑"}
	stu2 := sty{"lina", 29, "胡同"}
	stu3 := sty{"小杰", 29, "青年汇"}
	students["cx"] = stu1
	students["cw"] = stu2
	students["no2"] = stu3
	fmt.Println(students)
	for i,v := range students{
		fmt.Printf("学生编号%v\n",i)
		fmt.Printf("学生名字%v\n",v.name)
		fmt.Printf("学生住址%v\n",v.add)
		fmt.Printf("学生年龄%v\n",v.age)
	}
	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter08\deom05\main.go
// map[cw:{lina 29 胡同} cx:{tom 29 天通苑} no2:{小杰 29 青年汇}]
// 学生编号cx
// 学生名字tom
// 学生住址天通苑
// 学生年龄29
// 学生编号cw
// 学生名字lina
// 学生住址胡同
// 学生年龄29
// 学生编号no2
// 学生名字小杰
// 学生住址青年汇
// 学生年龄29