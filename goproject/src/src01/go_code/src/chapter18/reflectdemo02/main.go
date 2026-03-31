package main

import (
	//"fmt"
	"fmt"
	"reflect"
)

//通过反射修改
// num int 的值，修改student的值

func reflect01(b interface{}){
	rVal:= reflect.ValueOf(b)
	fmt.Printf("kind=%v\n",rVal.Kind())//地址值
	//rVal.SetInt(20)
	//*&rVal.SetInt(20)
	/*
Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。
	*/
	rVal.Elem().SetInt(20)


}
func main(){
	var num int = 10
	//fmt.Println()
	fmt.Println("num=",num)
	reflect01(&num)
	fmt.Println("num=",num)
}
//执行结构
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter18\reflectdemo02\main.go
// num= 10
// kind=ptr
// num= 20