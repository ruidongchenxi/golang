package main
import "fmt"
func main(){
	b := make(map[int]string)
	b[12] = "cf"
	b[13] = "ds"
	// 增加
	b[20] = "678"
	fmt.Println(b)
	//修改
	b[12] = "cx"
	fmt.Println(b)
	//删除
    delete(b,13)
	fmt.Println(b)
	//查
	value,flag := b[12]
	fmt.Println(value)
	fmt.Println(flag)
	
}