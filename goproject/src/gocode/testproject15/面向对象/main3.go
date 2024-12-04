package main 
import "fmt"
type Student struct {
	Age int
}
type Stu Student
func main(){
	var s1 Student = Student{19}
	var s2 Stu = Stu{15}
	s1 = Student(s2)//强转
	fmt.Println(s1)
	fmt.Println(s2)
}