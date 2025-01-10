package main
import (
	"fmt"
	"reflect"
)
//定义商标结够
type Brand struct {

}
//为商标结够添加方法
func (t Brand) Show(){

}
//为Brand定义一个别名FakeBrand
type FakeBrand = Brand
type Vehicle struct{
	//嵌入两个结构体
	FakeBrand
	Brand
}
func main(){
	//声明变量a,将Vehicle实例化a
	var a Vehicle
	a.FakeBrand.Show()//显示调用Vehicle中FakeBrand的Show()方法
	//取a 值
	ta := reflect.TypeOf(a)//使用反射取变量a的反射类型对象。以查看其他成员
	//遍历a的结构体成员
	for i:= 0;i<ta.NumField(); i++{
		//a的成员
		f := ta.Field(i)
		//打印成员的字段名和类型
		fmt.Printf("fieldName: %v,FieldType: %v\n", f.Name, f.Type.Name())
	}
}
