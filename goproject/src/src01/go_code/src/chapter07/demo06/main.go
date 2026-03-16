package main

import (

	"fmt"
)

func main(){
	// var t = [3][5]int {{67,98,75,86,96},{95,66,97,78,64},{65,78,83,78,96}}
	var t [3][5]float64
	//var f float64
	var s float64
	var z float64
	for i,v := range t{
		for s,_ := range v{
			fmt.Printf("请输入%v班的第%v个学生成绩\n",i+1,s+1)
			fmt.Scanln(&t[i][s])
		}
		//fmt.Println(t[i])
	}

	for i,v := range t{
		//var w int
		var a float64  
		for _,x := range v{
			a += x
			z++
		}
		fmt.Printf("%v班的平均分%v\n",i+1,a/float64(len(v)))
		s +=a
	}
	fmt.Printf("所有班级人的平均分是%v",s/z)
	


}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo06\main.go
// 请输入1班的第1个学生成绩
// 98
// 请输入1班的第2个学生成绩
// 78.5
// 请输入1班的第3个学生成绩
// 45.5
// 请输入1班的第4个学生成绩
// 98
// 请输入1班的第5个学生成绩
// 88
// 请输入2班的第1个学生成绩
// 87
// 请输入2班的第2个学生成绩
// 98
// 请输入2班的第3个学生成绩
// 56
// 请输入2班的第4个学生成绩
// 98
// 请输入2班的第5个学生成绩
// 45
// 请输入3班的第1个学生成绩
// 98
// 请输入3班的第2个学生成绩
// 98.5
// 请输入3班的第3个学生成绩
// 56
// 请输入3班的第4个学生成绩
// 98
// 请输入3班的第5个学生成绩
// 54
// 1班的平均分81.6
// 2班的平均分76.8
// 3班的平均分80.9
// 所有班级人的平均分是79.76666666666667