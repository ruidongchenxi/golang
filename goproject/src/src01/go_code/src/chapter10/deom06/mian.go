package main
import (
	"fmt"
)

type xs struct{
	Name string
	Age int
	Score float64
}
func (s *xs)SetScoer(a  float64){
	if a >0 && a<=100{
		s.Score=a
	}
}
func (s *xs)ShowInfo(){
	fmt.Printf("%v学生的成绩%v,年龄%v\n",s.Name,s.Score,s.Age)
}
// 
type Pupil struct{
	xs
}
func (p *Pupil)Tesing(){
	fmt.Printf("小学生%v正在考试\n",p.Name)
}
type Graduate struct{
	xs
}
func (p *Graduate)Tesing(){
	fmt.Printf("大学生%v正在考试\n",p.Name)
}

//小学生
// type Pupil struct{
// 	Name string
// 	Age int
// 	Score float64
// }

// func (p *Pupil)ShowInfo(){
// 	fmt.Printf("%v的年龄是%v,成绩是%v\n",p.Name,p.Age,p.Score)
// }
// func (p *Pupil)SetScoer(a float64) {
// 	if a >=1 && a <=100{
// 		p.Score=a
// 	}else {
// 		fmt.Println("输入的分数不正确")
// 	}
// }
// func (p *Pupil)Tesing(){
// 	fmt.Printf("小学生%v正在考试\n",p.Name)
// }
// type Graduate struct{
// 	Name string
// 	Age int 
// 	Score float64
// }
// func (p *Graduate)ShowInfo(){
// 	fmt.Printf("%v的年龄是%v,成绩是%v\n",p.Name,p.Age,p.Score)
// }
// func (p *Graduate)SetScoer(a float64) {
// 	if a >=1||a <=100{
// 		p.Score=a
// 	}else {
// 		fmt.Println("输入的分数不正确")
// 	}
// }
// func (p Graduate)Tesing(){
// 	fmt.Printf("大学生%v正在考试\n",p.Name)
// }
func main(){
//    t := &Pupil{
// 	Name: "小希",
// 	Age: 12,
// 	Score: 78.6,
//    }
//    t.Tesing()
//    t.ShowInfo()
//    c :=&Graduate{
// 		Name: "炘南",
// 		Age: 20,
// 		Score: 89,
//    }
//    c.ShowInfo()
//    c.Tesing()
	t :=&Pupil{xs{Name: "晨曦",Age:12,Score: 97 }}
	t.Tesing()
	t.ShowInfo()
	t.Tesing()
	c:=&Graduate{}
	c.xs.Name="炘南"
	c.xs.Age=23
	c.xs.Score=98
	
	//r:=&Graduate{xs{}}
	//r.xs=
	c.ShowInfo()
	c.Tesing()
	c.SetScoer(56)
}
//代码冗余
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter10\deom05\mian.go
// 小学生晨曦正在考试
// 晨曦学生的成绩97,年龄12
// 小学生晨曦正在考试
// 炘南学生的成绩98,年龄23
// 大学生炘南正在考试