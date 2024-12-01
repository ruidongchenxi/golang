package main
import "fmt"
//定义函数
func cal (num1 int,num2 int) (int) { //如果返回值类型就一个的话，那么（）可以省略不写的
    var sum int
	sum += num1
	sum += num2
	return sum
}
func cal2 (num1 int,num2 int) (int,int){
	var sum int
	sum += num1
	sum += num2
	var result int = num1 - num2
	return sum,result
}
func main(){
	var c int 
	var d int 
	fmt.Printf("请输入要计算的两个数值空格分开")
	fmt.Scanf ("%d %d",&c, &d)
// 调用函数
    sum := cal(c,d)
	fmt.Println(sum)
	sum1,result1 := cal2(c,d)
    fmt.Println(sum1,result1)
} 