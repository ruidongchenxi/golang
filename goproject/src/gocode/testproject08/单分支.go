package main
import "fmt"
func main(){
	//实现功能： 如果口罩的库存少于30个，就提示库存不足
	var count int
	fmt.Println("请输入现有口罩库存：")
	fmt.Scanln(&count)
	if count < 30 {
		fmt.Println("口罩数量不足")
	} 
}