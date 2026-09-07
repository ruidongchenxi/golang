package main

import (
	// "encoding/json"
	"fmt"
	"lnhgo/protobuf/day1/helloworld/proto"

	//"github.com/golang/protobuf/tree/master/proto"
	"google.golang.org/protobuf/proto"
)
type Hello struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Courses []string `json:"courses"`
}
func main(){
	req:=helloworld.HelloRequest{
		Name: "bobby",
		Age: 18,
		Courses: []string{"go","gin","微服务"},
	}
	// jsonStruct:= Hello{Name: "bobby",
	// 	Age: 18,
	// 	Courses: []string{"go","gin","微服务"},
	// }
	// jsonRsp,_:=json.Marshal(jsonStruct)
	// fmt.Println(string(jsonRsp))
	rsp,_:=proto.Marshal(&req)
	//fmt.Println(rsp)
	newReq:=helloworld.HelloRequest{}
	proto.Unmarshal(rsp,&newReq)
	fmt.Println(newReq.Name,newReq.Age)
}