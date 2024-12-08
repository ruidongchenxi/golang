package main
import (
	"fmt"
)
type Animal struct{
    Age int
	Weight float32
}
//绑定方法

func (an *Animal) Shout(){
	fmt.Println("大声喊叫")
}
func (an *Animal) ShowInfo(){
	fmt.Printf("动物的年龄：%v,动物的体重：%v",an.Age,an.Weight)
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
	cat.Animal.Age = 3
	cat.Animal.Weight = 10.3
	cat.Animal.Shout()
	cat.Animal.ShowInfo()
}