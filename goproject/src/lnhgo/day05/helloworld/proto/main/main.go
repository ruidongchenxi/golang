package main

import (
	//"encoding/json"
	"fmt"
	"lnhgo/day05/helloworld/proto"

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
	// jsonStruct:=Hello{Name: "bobby",
	// 	Age: 18,
	// 	Courses: []string{"go","gin","微服务"},
	// }
	//jsonrsp,_:=json.Marshal(jsonStruct)
	//fmt.Println(len(string(jsonrsp)))
	rsp,_:=proto.Marshal(&req)//
	newReq:=helloworld.HelloRequest{}
	proto.Unmarshal(rsp,&newReq)
	fmt.Println(len(string(rsp)))
	fmt.Println(newReq.Age,newReq.Name,newReq.Courses)
}