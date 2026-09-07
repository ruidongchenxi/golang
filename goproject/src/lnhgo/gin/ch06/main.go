package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lnhgo/gin/ch06/proto"
)

func main() {
	router := gin.Default()
	//匹配的url格式：/welcome?firstname=Jane&lastname=Doe
	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtoBuf",returnProto)
	router.Run(":8015")
}
func moreJSON(c *gin.Context){
	var msg struct{
		Name string `json:"user"`
		Messages string 
		Number int
	}
	msg.Name="boobby"
	msg.Messages="这是一个测试"
	msg.Number= 20
	c.JSON(http.StatusOK,msg)
}
func returnProto(c *gin.Context){
	coures:=[]string{"python","go","微服务"}
	user:=&proto.Teacher{
		Name: "bobby",
		Course: coures,
	}
	c.ProtoBuf(http.StatusOK,user)
}