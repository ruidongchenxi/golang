package main
import "fmt"
func main(){
	//定义
	b := make(map[int]string)
	b[12] = "cf"
	b[13] = "ds"
	b[9] = "az"
	b[5] = "do"
	b[3] = "io" 
	//获取 
	fmt.Println(len(b))
	// 遍历
	for k,v := range b{
		fmt.Printf("key为：%v value: %v \n",k,v)
	}
	//加深
	a := make(map[string]map[int]string)
	//赋值
	a["班级1"] = make(map[int]string,3)
	a["班级1"][01] = "cx"
	a["班级1"][02] = "ix"
	a["班级1"][03] = "bx"
	a["班级2"] = make(map[int]string,3)
	a["班级2"][01] = "cxe"
	a["班级2"][02] = "cxo"
	a["班级2"][03] = "ceo"
	for k1,v1 := range a {
		fmt.Println(k1)
		for k,v := range v1{
			fmt.Printf("学号：%v 姓名: %v \n",k,v)
		}
		fmt.Println()
	}
}