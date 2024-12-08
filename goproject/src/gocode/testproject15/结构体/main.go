package main
import "fmt"
type Student struct{
	Name string
	Age int
}
func main(){
	//按结构体赋值；必须按照顺序
    var s1 Student = Student{"小李",19}
    fmt.Println(s1)
	//方式2
	var s2 Student = Student{
		Name : "丽丽"
		Age : 16
	}
	fmt.Println(s2)
	//方式3
	var s3 *Student = &Student("明明",26)
	fmt.Println(*s3)
}