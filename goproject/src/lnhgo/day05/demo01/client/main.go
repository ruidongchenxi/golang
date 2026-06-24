package main

import (
	"encoding/json"
	"fmt"
	//"image/jpeg"

	"github.com/kirinlabs/HttpRequest"
)
type ResponseData struct{
	Data int `json:"data"`
	
}
func Add(a,b int) int {
	req:=HttpRequest.NewRequest()
	res,_:=req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d","hello",a,b))
	body,_:=res.Body()
	//fmt.Println(string(body))
	Rsp:=ResponseData{}
	json.Unmarshal(body,&Rsp)
	return Rsp.Data
}
func main(){
	fmt.Println(Add(4,9))

}