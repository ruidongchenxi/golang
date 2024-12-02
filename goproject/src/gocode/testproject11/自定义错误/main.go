package main
import (
	"fmt"
    "errors"
)	
func main(){
	err := test()
	if err != nil {
		fmt.Println("自定义错误：",err)
        panic(err)
	}
	fmt.Println("上面除法操作执行成功")
	fmt.Println("正常执行")
}
func test() (err error){

	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("除数不能为0")

	}else{ //除数不为0 正常执行
	result := num1 / num2
	fmt.Println(result)
	//如果没有错误返回nil 值
	return nil 
	}
}