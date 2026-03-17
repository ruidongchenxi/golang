package main
import (
	"fmt"
	//"slices"
	"encoding/json"
)
type Person struct{
	x int
	y int 
}
type Rect struct{
	leftUP,rightDown Person
}
type Rect1 struct{
	leftUP,rightDown *Person
}
type C struct{
	Name string
	Age int 
}
type w C
type Monster struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}
func main(){
	monster:=Monster{"牛魔王",500,"水火棍"}
	//fmt.Println(monster)
//将monster序列化成json格式的字串
    jsonmonster,err := json.Marshal(monster)
	if err==nil{
		fmt.Println(string(jsonmonster))
	}else {
		fmt.Println(err)
	}

}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter09\deom04\main.go
// {"name":"牛魔王","age":500,"skill":"水火棍"}