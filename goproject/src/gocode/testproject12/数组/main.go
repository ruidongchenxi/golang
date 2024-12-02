package main
import "fmt"
func main(){
	//给出五个学生的成绩，求成绩总和，平均数
    score1 := 95
	score2 := 91
	score3 := 97
	score4 := 76
	score5 := 65
	//求和
	sum1 := score1 + score2 + score3 + score4 + score5
	//平均数
	avg1 := sum1 / 5 
	fmt.Printf("成绩总和多少: %v,平均总和为：%v \n",sum1,avg1)
	//定义一个数组
	var scores [5]int 
	//将成绩存入数组
	scores[0] = 95
	scores[1] = 83 
	scores[2] = 73
	scores[3] = 72
	scores[4] = 80
	sum := 0
	for i :=0; i <len(scores);i++{ //遍历数组
		sum += scores[i]
	}
	//求和
	avg := sum / len(scores)
	//输出
	fmt.Printf("和：%v，平均：%v",sum,avg)

}