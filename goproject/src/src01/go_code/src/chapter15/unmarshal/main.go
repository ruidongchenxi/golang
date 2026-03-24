package main
import (
	"fmt"
	"encoding/json"
)
type Monster struct{
	Name string `json:"name"` //tag 指定序列号后的字段格式
	Age int `json:"age"` //tag 指序列后的字段格式
	Birthday string
	Sal float64
	Skill string
}
//演示将JSON字符串，反序列化成struct
func unmarshalStruct(){
// 说明str 在项目开发中，是通过网络传输获取到的。。。或者读文件获取到的
	str:="{\"name\":\"牛魔王\",\"age\":2500,\"Birthday\":\"1年1月16日\",\"Sal\":8900,\"Skill\":\"牛魔拳\"}"
    var monster Monster
	//反序列化操作
	err:=json.Unmarshal([]byte(str),&monster)
	if err !=nil{
		fmt.Printf("Unmarshal err=%v\n",err)
	}
	fmt.Println(monster)

}
//演示将JSON字符串，反序列化成map
func unmarshalmap(){
	var a map[string]interface{}
	str := "{\"name\":\"牛魔王\",\"age\":2500,\"Birthday\":\"1年1月16日\",\"Sal\":8900,\"Skill\":\"牛魔拳\"}"
	//直接反序列化不需要make;因为make操作直接被封装在json.Unmarshal函数里了
	err:=json.Unmarshal([]byte(str),&a)
	if err !=nil{
		fmt.Printf("Unmarshal err=%v\n",err)
	}
	fmt.Println(a)

}
//将json 反序列切片
func unmarshalslice(){
	var slice []map[string]interface{}
	str := "[{\"address\":\"北京\",\"age\":6,\"name\":\"jack\"},{\"Name\":\"小龙\",\"address\":\"上海\",\"age\":7}]"
	// err:=json.Unmarshal([]byte(str),slice)
	// if err!=nil{
	// 	fmt.Printf("Unmarshal err=%v\n",err)
	// }
	err:=json.Unmarshal([]byte(str),&slice)
	if err !=nil{
		fmt.Printf("Unmarshal err=%v\n",err)
	}
	fmt.Println(slice)
}
func main(){
	unmarshalStruct()
	unmarshalmap()
	unmarshalslice()

}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter15\unmarshal\main.go  
// {牛魔王 2500 1年1月16日 8900 牛魔拳}
// map[Birthday:1年1月16日 Sal:8900 Skill:牛魔拳 age:2500 name:牛魔王]
// [map[address:北京 age:6 name:jack] map[Name:小龙 address:上海 age:7]]