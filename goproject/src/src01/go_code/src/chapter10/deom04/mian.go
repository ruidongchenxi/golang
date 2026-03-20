package main
import (
	"fmt"
)
//小学生
type Pupil struct{
	Name string
	Age int
	Score float64
}

func (p *Pupil)ShowInfo(){
	fmt.Printf("%v的年龄是%v,成绩是%v\n",p.Name,p.Age,p.Score)
}
func (p *Pupil)SetScoer(a float64) {
	if a >=1||a <=100{
		p.Score=a
	}else {
		fmt.Println("输入的分数不正确")
	}
}
func (p *Pupil)Tesing(){
	fmt.Printf("小学生%v正在考试\n",p.Name)
}
type Graduate struct{
	Name string
	Age int 
	Score float64
}
func (p *Graduate)ShowInfo(){
	fmt.Printf("%v的年龄是%v,成绩是%v\n",p.Name,p.Age,p.Score)
}
func (p *Graduate)SetScoer(a float64) {
	if a >=1||a <=100{
		p.Score=a
	}else {
		fmt.Println("输入的分数不正确")
	}
}
func (p Graduate)Tesing(){
	fmt.Printf("大学生%v正在考试\n",p.Name)
}
func main(){
   t := &Pupil{
	Name: "小希",
	Age: 12,
	Score: 78.6,
   }
   t.Tesing()
   t.ShowInfo()
   c :=&Graduate{
		Name: "炘南",
		Age: 20,
		Score: 89,
   }
   c.ShowInfo()
   c.Tesing()
}
//代码冗余
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter10\deom04\mian.go
// 小学生小希正在考试
// 小希的年龄是12,成绩是78.6
// 炘南的年龄是20,成绩是89
// 大学生炘南正在考试