package main

import "fmt"

func mPrint(datas ...interface{}) {
	for _, v := range datas {
		fmt.Println(v)
	}

}
func mPrint2(datas interface{}) {
	fmt.Println(datas)

}
type myinfo struct{}
func (mi *myinfo)Error() string{
	return "我不是错误"
}
func main() {
	var data = []string{
		"bobby", "tr", "te",
	}//如果date 是字符串切片类型就不可以
	

	var datai []interface{}
	for _,v:=range data{
		datai=append(datai,v)
	}
	mPrint(datai...)
	mPrint2(data)
}