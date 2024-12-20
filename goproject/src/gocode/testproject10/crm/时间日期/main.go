package main
import(
	"fmt"
	"time"
)
func main(){
    now := time.Now()//当前时间
	//Now()返回值是一个结构体，类型是：time.Time 
	fmt.Println(now)
	fmt.Printf("%v ~~~对应的类型为：%T \n",now,now)
	// 获取年月日
	fmt.Printf("年：%v \n",now.Year())
	fmt.Printf("月：%v \n",now.Month())
	fmt.Printf("月：%v \n",int(now.Month()))
	fmt.Printf("日：%v \n",now.Day())
	fmt.Printf("时：%v \n",now.Hour())
	fmt.Printf("分：%v \n",now.Minute())
	fmt.Printf("秒：%v \n",now.Second())
	fmt.Println("---------------")
	//将字符串直接输出
	fmt.Printf("当前年月日：%d-%d-%d 时分秒：%d:%d:%d \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	//Sprintf 得到字符串以便后续使用
	datestr := fmt.Sprintf("当前年月日：%d-%d-%d 时分秒：%d:%d:%d \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Println(datestr)
	fmt.Println("------------------------------------------------------------------")
	//参数字符串的个各数字必须是固定的，必须这样写
	datestr2 := now.Format("2006/01/02 15/04/05")
	fmt.Println(datestr2)

//选择任何组合都可以，根据需求自己选择。时间必须是固定的（2006/01/02 15/04/05 任意组合没问题）
	datestr3 := now.Format("2006 15:04")
	fmt.Println(datestr3)
	
}