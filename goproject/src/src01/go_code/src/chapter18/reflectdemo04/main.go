package main

import (
	//"fmt"
	"fmt"
	"reflect"
)
//结构体
type Monster struct{
	Name string `json:"name"`
	Age int `json:"int"`
	Score float32
	sex string
}
//绑定方法
func (s Monster)Print(){
	fmt.Println("--------start-------")
	fmt.Println(s)
	fmt.Println("--------end---------")
}
//绑定方法返回两数和
func (s Monster)GgeSum(n1,n2 int) int{
	return n1+n2
}
//接收4个值给结构体变量赋值
func (s Monster)Set(name string,age int,score float32,sex string){
	s.Name=name
	s.Age=age
	s.Score=score
	s.sex=sex
}

func TestStruct(a interface{}){
	//获取reflect.Type 类型
	tye := reflect.TypeOf(a)
	//获取reflect.Value 类型
	val := reflect.ValueOf(a)
	//获取类别
	kb := val.Kind()
	//如果传入的是不是结构体，如果不是结构体退出
	if kb != reflect.Struct{
		fmt.Println("expect struct")
		return
	}
//获取结构体几个字段
	num := val.NumField()//返回v持有的结构体类型值的字段数，如果v的Kind不是Struct会panic
	fmt.Printf("struct has %d fields\n",num)
	//遍历结构体字段
	for i :=0;i<num;i++{
		fmt.Printf("Field %d:值为=%v\n",i, val.Field(i))//func (Value) Field  ;func (v Value) Field(i int) Value;返回结构体的第i个字段（的Value封装）。如果v的Kind不是Struct或i出界会panic
		//获取到结构体标签，注意需要通过reflect.type来获取; Field(i int) StructField
   		 // 返回索引序列指定的嵌套字段的类型，
    	// 等价于用索引中每个值链式调用本方法，如非结构体将会panic;func (StructTag) Get
		//func (tag StructTag) Get(key string) string
		//Get方法返回标签字符串中键key对应的值。如果标签中没有该键，会返回""。如果标签不符合标准格式，Get的返回值是不确定的。//func (tag StructTag) Get(key string) string
		tagVal := tye.Field(i).Tag.Get("json") 
		if tagVal !=""{
			fmt.Printf("Field %d: tag为=%v\n",i,tagVal)
		}
	}

	numOfMethod := val.NumMethod()//获取多少个方法
	fmt.Printf("struct has %d methods\n",numOfMethod)

	val.Method(1).Call(nil)// val.Method(1)从0开始获取第二个.Call(nil)调用方法，排序是按照方法首字母AISSC进行排序的，所以调用的是Print
	//调用结构体的第一个方法Method(0)
	var params []reflect.Value //声明切片
	params= append(params, reflect.ValueOf(10))
	params= append(params, reflect.ValueOf(40))
	res:=val.Method(0).Call(params)//传入参数是[]reflect.Value
	fmt.Println("res=",res[0].Int())//返回结果，返回结果是[]reflect.Value
}
func main(){
	var a Monster = Monster{
		Name: "哪吒",
		Age: 5,
		Score: 65.9,
		sex: "跑",
	}
	TestStruct(a)
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter18\reflectdemo04\main.go
// struct has 4 fields
// Field 0:值为=哪吒
// Field 0: tag为=name
// Field 1:值为=5
// Field 1: tag为=int
// Field 2:值为=65.9
// Field 3:值为=跑
// struct has 3 methods
// --------start-------
// {哪吒 5 65.9 跑}
// --------end---------
// res= 50