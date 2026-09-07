package main

import (
	"encoding/json"
	"fmt"
	//"strings"

	"github.com/kirinlabs/HttpRequest"
)
type ResponseData struct{
	Data int `json: "data"`
}
func Add(a,b int) int{
	req := HttpRequest.NewRequest()
	res,_:=req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d","add",a,b))
	body,_:=res.Body()
	//fmt.Println(string(body))
	rspdata:=ResponseData{}
	_=json.Unmarshal(body,&rspdata)
	return rspdata.Data

}
func main() {
	fmt.Println(Add(1,2))

}