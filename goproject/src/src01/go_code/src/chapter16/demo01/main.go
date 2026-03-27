package main
import(
	_"fmt"
)
func addUpper(n int)int {
	res:=0
	for i:=1;i<=n;i++{
		res+=i
	}
	return res
}
// func main(){
// 	//传统测试，在main 函数中测试
// 	res:=addUpper(10)
// 	if res!=55{
// 		fmt.Printf("addupper错误 返回值=%v，期望值=%v\n",res,55)
// 	}else{
// 		fmt.Printf("addUpper()正确 返回值=%v，期望值=%v\n",res,55)
// 	}
// }
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter16\demo01\main.go
// addupper错误 返回值=45，期望值=55