package main
import (
	"fmt"

)
func main(){
	a := make(map[string]string)
	a["a1"]= "阿库"
	a["a2"]= "小米"
	a["a3"]= "小鹿"
	a["a4"]= "小鹏"
	// fmt.Println("修改前打印",a)
	// //修改a[a2的值]
	// a["a2"]="小白"
	// fmt.Println("修改后打印",a["a2"])
	// fmt.Println("修改打印",a)
	// //增加一个
	// a["w1"]= "奶茶"
	// fmt.Println("增加新值打印",a)
	// //删除一个值
	// delete(a,"a3")

	// fmt.Println("删除后打印",a)
	// 查找
	val,ok := a["a1"]
	if ok{
		fmt.Println("有这个a1 key；值为",val)
	}else {
		fmt.Println("没有有这个a1 key；值为")
		
	}
	// fmt.Println("删除整个集合",a)
	// a = map[string]string{}
	// fmt.Println("清空后打印",a)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter08\deom03\main.go
// 有这个a1 key；值为 阿库