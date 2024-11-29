package main
import "fmt"
func main(){
	//实现功能： 如果口罩的库存少于30个，就提示库存不足
	var count int
	fmt.Println("请输入现有口罩库存：")
	fmt.Scanln(&count)
	if count <= 30 {
		fmt.Println("口罩数量不足")
	} else if count <= 35 && count > 30{
      fmt.Println("该补充口罩了")
	} else {// 建议你保证else的存在，只有有了else 的存储才真正起到多选一的效果
		fmt.Println("口罩充足")
	}
}