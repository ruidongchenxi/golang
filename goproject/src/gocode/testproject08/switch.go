package main
import "fmt"
func main(){
    var score int 
    fmt.Println("输入成绩：")
	fmt.Scanln(&score)
	switch  score/10 {
    case 10 :
		fmt.Println("您的等级为： A")
	case 9 :
		fmt.Println("您的等级为： B")
	case 8 :
		fmt.Println("您的等级为： C")
	case 7 :
		fmt.Println("您的等级为： D")
	case 6 :
		fmt.Println("您的等级为： E")
	case 5 :
		fmt.Println("您的等级为： E")
	case 4 :
		fmt.Println("您的等级为： E")
	case 3 :
		fmt.Println("您的等级为： E")
	case 2 :
		fmt.Println("您的等级为： E")
	case 1 :
		fmt.Println("您的等级为： E")
		fallthrough
	case 0 :
		fmt.Println("您的等级为： E")
	default: //用来兜底的分支
		fmt.Println("您的成绩有误")
    }
}
