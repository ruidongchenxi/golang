package main
import "fmt"
func main(){
	var age int
	var name string
	var score float32
	var isVIP bool
//方式2： Scanf
    fmt.Println("请输入学生年龄，姓名，成绩，是否是VIP，使用空格进行分割")
	fmt.Scanf("%d %s %f %t",&age,&name,&score,&isVIP)
	fmt.Printf("学生年龄：%v,姓名：%v，成绩：%v,是否vip：%v\n",age,name,score,isVIP)
}