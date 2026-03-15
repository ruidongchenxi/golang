package main
import (
	"fmt"
)
func main(){
	var foarr [6]float64
	foarr[0]= 3.0
	foarr[1]= 5.0
	foarr[2]= 1.0
	foarr[3]= 3.4
	foarr[4]= 2.0
	foarr[5]= 50.0
	var t float64
	for i:=0;i<len(foarr);i++{
		t = t+foarr[i]

	}
	//fmt.Sprintf("%.2f",t / 6) 将这个结果四舍五入保留到小数点两位返回值
	a := fmt.Sprintf("%.2f",t / float64(len(foarr))) // 6 具体的数是常量可以类型推导；所以必须强制转换
	fmt.Println("平均体重",a)
}
// 执行结果
// fmt.Sprintf("%.2f",t / 6) 