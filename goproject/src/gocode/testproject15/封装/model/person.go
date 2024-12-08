package model
import "fmt"

type person struct{
	Name string
	age int
}
//定义工厂模式的函数，相当于构造器
func NewPerson(name string) *person{
	return &person{
		Name:name,
	}
}
//定义set和age函数，对age字段进行封装，因为在函数中可以加一些列的限制操作，确保被封装字段的安全合理性
func (p *person) SetAge(age int){
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄输入不对")
	}
}
func (p *person) GetAge() int{
	return p.age
}