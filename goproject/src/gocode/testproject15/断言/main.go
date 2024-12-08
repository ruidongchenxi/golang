package main
import "fmt"
//接口定义
type Sayhello interface{
	//声明没有实现的方法
	Sayhello()
}
//接口的实现,结构体
type Chinese struct{

}
//自定义数据类型
type integer int 
func (i integer) Sayhello(){
	fmt.Println("say hi" ,i)
}
//实现接口的方法----》具体实现
func (person Chinese) Sayhello(){
	fmt.Println("你好")
}
// 中国人特有方法
func (person Chinese) NiuyangGe(){
    fmt.Println("东北人")
}
//接口的实现
//美国人
type American struct{

}
//实现接口的方法-----》具体的实现
func (person American) Sayhello(){
	fmt.Println("hi")
}
//定义函数:专门用来各国人打招呼的函数，接收具备Sayhello接口能力的变量：
func greet (s Sayhello){
	s.Sayhello()
	//判断S是否可以转成Chinese类型并且赋值给ch
	ch,flag := s.(Chinese)// //判断S是否可以转成Chinese类型并且赋值给ch
	if ch,flag := s.(Chinese);flag {
		ch.NiuyangGe()
	}else{
		fmt.Println("buhio")
	}
	//s.NiuyangGe()
	//ch.NiuyangGe()

}
func main(){
	//创建中国人
    //c := Chinese{}
	//
	a := American{}
	//var i integer = 10
	//var s Sayhello = i
	//i.Sayhello()
	greet(a)
	//greet(c)
	//greet(i)
}
