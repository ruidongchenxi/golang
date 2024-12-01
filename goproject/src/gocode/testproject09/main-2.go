package main
import "fmt"
//定义函数
func cal (num1 int,num2 int) { //如果返回值类型就一个的话，那么（）可以省略不写的
    var sum int
	sum += num1
	sum += num2
	fmt.Println (sum)
}
func main(){
	var c int 
	var d int 
	fmt.Printf("请输入要计算的两个数值空格分开")
	fmt.Scanf ("%d %d",&c, &d)
// 调用函数
    cal(c,d)

}