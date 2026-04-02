package main
import (
	"fmt"
	//"strconv"
)
type Student struct{
	Name string
	gender string
	age int 
	id int
	score float64
}
func (s Student)say() string{
	t := fmt.Sprintf("Student 的信息 name=[%v],gender=[%v],age=[%v],id=[%v],score=[%v]",
	s.Name,s.gender,s.age,s.id,s.score)
	return t
}

func main(){

	t:=Student{"西钊","男",25,1,98.5}
    fmt.Println(t.say())
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter09\deom07\main.go
// Student 的信息 name=[西钊],gender=[男],age=[25],id=[1],score=[98.5]