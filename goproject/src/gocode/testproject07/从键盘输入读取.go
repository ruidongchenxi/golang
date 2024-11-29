package main
import "fmt"
func main(){
	//实现功能：从键盘录入学生的年龄，姓名，成绩，是否vip
	//方式1： scanln
	var age int
	fmt.Println("请录入学生年龄：")
    //传入age的地址的目的： 在Scanln函数中，对地址中的值进行改变的时候，实际外面的age被影响了
	fmt.Scanln(&age)//录入数据的时候类型一定要匹配。因为底层会自动判断类型
	var name string
	fmt.Println("请录入学生姓名：")
    //传入age的地址的目的： 在Scanln函数中，对地址中的值进行改变的时候，实际外面的age被影响了
	fmt.Scanln(&name)
	var score float32
	fmt.Println("请录入学生成绩：")
    //传入age的地址的目的： 在Scanln函数中，对地址中的值进行改变的时候，实际外面的age被影响了
	fmt.Scanln(&score)
	var isVIP bool
	fmt.Println("请录入学生是否VIP：")
    //传入age的地址的目的： 在Scanln函数中，对地址中的值进行改变的时候，实际外面的age被影响了
	fmt.Scanln(&isVIP)
	fmt.Printf("学生年龄：%v,姓名：%v，成绩：%v,是否vip：%v\n",age,name,score,isVIP)
}