package model
import "fmt"
type person struct {
	Name string
	age  int
	sal   float64
}
func NewPerson(n string) *person { //构造函数
	return &person{
		Name: n,
	}
}
func (d *person)SetAge(n int){
	if n >0 && n<150{
		d.age=n
	}else{
		fmt.Println("输入年龄不正确")
	}
}
func (d *person)SetSal(n float64){
	if n>=3000 && n<30000{
	    d.sal=n
	}else {
		fmt.Println("输入不正确")
	}
}
func (d *person)GetAge(){
	fmt.Println(d.Name,"的年龄是",d.age)
}
func (d *person)GetSal(){
	fmt.Println(d.Name,"的工资是",d.sal)
}
func (d *person) Get() {
	fmt.Printf("名字是%v,年龄是%v,工资是%v\n",d.Name,d.age,d.sal)
}