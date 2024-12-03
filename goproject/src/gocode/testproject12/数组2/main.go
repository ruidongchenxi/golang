package main
import "fmt"
func main(){
	//声明数组
	//var arr [3]int16
	//获取数组长度
	//fmt.Println(len(arr))
	//输出数组
	//fmt.Println(arr)
	//验证arr中存储的是地址值
	//fmt.Printf("arr的地址为：%p\n",&arr)
	//fmt.Printf("arr的地址为：%p\n",&arr[0])
	//fmt.Printf("arr的地址为：%p\n",&arr[1])
	//fmt.Printf("arr的地址为：%p\n",&arr[2])
	//s1 :=0
	//fmt.Println("请输入学生个数：")
    //fmt.Scanln(&s1)
	var scores [5]int
	sum := 0
	for i :=0; i< len(scores); i++{
		fmt.Printf("请输入第%d学生成绩：" ,i+1)
		fmt.Scanln(&scores[i])
		sum +=scores[i]
	}
    avg := sum / len(scores) 
	fmt.Printf("和：%v,平均数：%v \n",sum,avg)
	for key,value := range scores {
		fmt.Printf("第%d个学生成绩为：%d\n",key+1,value)
	}
}