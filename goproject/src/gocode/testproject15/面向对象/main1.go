package main
import "fmt"
//定义一个结构体
type Teacher struct{
	//变量名字
	name string
	Age int
	School string  
}
func main(){
	/*
	//创建结构体
	var t1 Teacher 
	fmt.Println(t1)//在为赋值时默认值：{0}
	t1.name = "小溪"
	t1.Age = 45
	t1.School = "三中"
	fmt.Println(t1)
	
	//创建结构体；*根据地址取值
	
	var t *Teacher = new(Teacher)
    //t 是指针，t其实指向的就是地址，应该给这个地址指向的对象赋值操作
	(*t).name = "小虎"
	(*t).Age = 16
	//(*t).School = "三中"
	//为了符合程序员编程习惯，go提供了简化赋值操作
	t.School = "五中" //go编译器底层对t.School 做转化(*t).School  = "三中"
	fmt.Println(*t)
	*/
	//创建结构体
	var t *Teacher = &Teacher{}
	(*t).name = "小虎"
	(*t).Age = 16
	//(*t).School = "三中"
	//为了符合程序员编程习惯，go提供了简化赋值操作
	t.School = "五中"
	fmt.Println(*t)
}