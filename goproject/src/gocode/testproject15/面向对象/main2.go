package main 
import "fmt"
type Student struct {
	Age int
}
type Person struct {
	Age int
}
func main(){
	var s Student
	var p Person
	s = Student(p)
	fmt.Println(s)
	fmt.Println(p)
}