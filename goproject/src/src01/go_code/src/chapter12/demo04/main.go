package main
import (
	"fmt"
	// "math/rand"
	// "sort"
)
type BirdAble interface{
	Flying()
}
type fish interface{
	Fish()
}
type  Monkey struct{
	Name string
}
func (this *Monkey)climbing(){
	fmt.Println("爬树")
}
type  LittleMonkey struct{
	Monkey
}
func (t *LittleMonkey)Flying(){
	fmt.Println(t.Name,"学会鸟的飞行")
}
func (t *LittleMonkey)Fish(){
	fmt.Println(t.Name,"学会了游泳")
}
func main(){
	monkey:=LittleMonkey{Monkey{Name: "悟空"}}
	monkey.climbing()
	var B BirdAble = &monkey
	B.Flying()
	var S fish =&monkey
	S.Fish()

}
// 执行结过
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter11\demo04\main.go
// 爬树
// 悟空 学会鸟的飞行
// 悟空 学会了游泳