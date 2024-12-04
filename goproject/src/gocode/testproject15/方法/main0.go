package main
import "fmt"
//定义Person结构体
type Person struct {
	Name string
}
//给Person结构体绑定方法test
//func (p Person) test(){//参数名字随便起
func (p *Person) test(){ //指针对象的方法
	p.Name="露露"
	fmt.Println(p.Name)
}
func main(){
	//创建结构体对象
	var p Person
	p.Name = "丽丽"
	fmt.Printf("p的地址值：%p \n",&p)
	fmt.Println(p.Name)
	//p.test() //调用方法
	//(&p).test()//调用方法
	p.test() //指针类型也可以这么写。底层编译器做了优化，底层自动帮我们加上&符号
	fmt.Println(p.Name )
}