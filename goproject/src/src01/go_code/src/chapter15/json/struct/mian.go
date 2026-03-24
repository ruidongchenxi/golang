package main

import (
	"encoding/json"
	"fmt"
	//"strings"
)
type Monster struct{
	Name string `json:"name"` //tag 指定序列号后的字段格式
	Age int `json:"age"` //tag 指序列后的字段格式
	Birthday string
	Sal float64
	Skill string
}
func tsetStruct() {
	//
	monster:=Monster{
		Name: "牛魔王" ,
		Age: 2500 ,
		Birthday: "1年1月16日",
		Sal: 8900,
		Skill: "牛魔拳",
	}
	//将moster序列化
	data,err:=json.Marshal(&monster)
	if err !=nil{
		fmt.Println("序列号失败",err)
		return
	}
	//输出系列化后结果
	fmt.Printf("序列化后数据=%v\n",string(data))
	
}
func testMap(){
	//定义一个map
	var a map[string]interface{}
	//
	a= make(map[string]interface{})
	a["name"]="红孩儿"
	a["age"]=30
	a["adderss"]="火云洞"
	//将map 序列号
	data,err:=json.Marshal(a)
	if err !=nil{
		fmt.Println("序列号失败",err)
		return
	}
	//输出系列化后结果
	fmt.Printf("序列化后数据=%v\n",string(data))

}
func testSlice(){
	var slice []map[string]interface{}
	//slice=append(slice, map[string]interface{}{"名字":"雷锋"})
	var m1 map[string]interface{}
	m1= make(map[string]interface{})
	m1["name"]="jack"
	m1["age"]= 6
	m1["address"]="北京"
	slice = append(slice, m1)
	var m2 map[string]interface{}
	m2=make(map[string]interface{})
	m2["Name"]="小龙"
	m2["age"]=7
	m2["address"]="上海"
	slice=append(slice, m2)
	data,err:=json.Marshal(slice)
	if err !=nil{
		fmt.Println("序列号失败",err)
		return
	}
	//输出系列化后结果
	fmt.Printf("序列化后数据=%v\n",string(data))

}
//对基本数据类型序列号没有什么意义
func tsetInt(){
	var num1 float64=123.67
	data,err:=json.Marshal(num1)
	if err !=nil{
		fmt.Println("序列号失败",err)
		return
	}
	//输出系列化后结果
	fmt.Printf("序列化后数据=%v\n",string(data))

}
func main(){
	//演示将结构体、map、切片进行序列化
	tsetStruct()
	testMap()
	testSlice()
	tsetInt()
}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter15\json\struct\mian.go
// 序列化后数据={"name":"牛魔王","age":2500,"Birthday":"1年1月16日","Sal":8900,"Skill":"牛魔拳"}
// 序列化后数据={"adderss":"火云洞","age":30,"name":"红孩儿"}
// 序列化后数据=[{"address":"北京","age":6,"name":"jack"},{"Name":"小龙","address":"上海","age":7}]
// 序列化后数据=123.67