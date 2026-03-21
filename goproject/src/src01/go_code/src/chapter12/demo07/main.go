package main

import (
	"fmt"
	//"go/types"
	//"internal/syscall/windows"
)

//定义接口
type Usb interface{
	//声明两个没有实现的方法
	Start()
	Stop()
}
type Phone struct{
	Name string
}
func (p Phone)Start(){
	fmt.Println("手机开始工作")
}
func (p Phone)Stop(){
	fmt.Println("手机停止工作")
}
func (p Phone)Call(){
	fmt.Println("手机打电话")
}
type Camer struct{
	Name string
}
func (p Camer)Start(){
	fmt.Println("相机开始工作")
}
func (p Camer)Stop(){
	fmt.Println("相机停止工作")
}
type Compose struct{
}
//此方法接收一个Usb 接口类型变量；只要实现了Usb方法；就是指实现了USB接口声明的所有方法
func (p Compose)Working(usb Usb){
	//var t Phone
	usb.Start()
	//类型断言应用
	if t,ok:=usb.(Phone);ok{
		t.Call()
	}
	usb.Stop()
}
func TypeJudge(it ...interface{}){
	for i,x :=range it{
		switch x.(type){//这里是判断type是固定写法
		case bool:
			fmt.Printf("param #%d is a bool 值是%v\n",i,x)
		case float64:
			fmt.Printf("param #%d is a float64 值是%v\n",i,x)
		case int, int64:
			fmt.Printf("param #%d is a int 值是%v\n",i,x)
		case nil:
			fmt.Printf("param #%d is a nil 值是%v\n",i,x)
		case string:
			fmt.Printf("param #%d is a string 值是%v\n",i,x)
		case *float64:
			fmt.Printf("param #%d is a *float64 值是%v\n",i,x)
		case *int,*int64:
			fmt.Printf("param #%d is a *int 值是%v\n",i,x)
		case *string:
			fmt.Printf("param #%d is a *string 值是%v\n",i,x)
		default:
			fmt.Printf("param #%d's type is unknown 值是%v\n",i,x)
		}
	}
}

func main(){
	//定义一个Usb接口数组,可以存放Phone和Camera 的结构体变量；体现多态数组
	// var usbArr [3]Usb
	
	// usbArr[0]=Phone{"华为"}
	// usbArr[1]=Phone{"小米"}
	// usbArr[2]=Camer{"佳能"}
	// fmt.Println(usbArr)
	// var c Compose
	// // var t Phone
	// for _,v:=range usbArr{
	// 	c.Working(v)
	// 	fmt.Println()
	// }
	t := 6
	//var y interface{}=t
	TypeJudge(t)
    r := "65"
	TypeJudge(r)
	e :="你好"
	w := &e
	TypeJudge(w)
	
}

// PS D:\golang\goproject\src\src01\go_code\src> go run chapter11\demo07\main.go
// param #0 is a int 值是6
// param #0 is a string 值是65
// param #0 is a *string 值是0xc000026080