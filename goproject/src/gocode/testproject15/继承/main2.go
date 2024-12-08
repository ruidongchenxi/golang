package main
import (
	"fmt"
)
type Animal struct{
    Age int
	wight float32
}
//绑定方法

func (an *Animal) shout(){
	fmt.Println("大声喊叫")
}
func (an *Animal) showInfo(){
	fmt.Printf("动物的年龄：%v,动物的体重：%v",an.Age,an.wight)
}
//定义结构体
type Cat struct{
	//为了复用性，体现继承思维，加入匿名结构体
	Animal
}
//对cat 绑定特有方法
func (c *Cat) scratch(){
	fmt.Println("我是小猫")
}
func main(){
	cat := &Cat{}
	//直接访问
	cat.Age = 3
	cat.wight = 10.3
	cat.shout()
	cat.showInfo()
}