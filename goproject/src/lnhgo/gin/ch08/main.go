package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func MyLogger() gin.HandlerFunc{
	return  func(c *gin.Context) {
		t:=time.Now()
		c.Set("example","123456")
		//return 
		c.Next()//让原来该执行的逻辑继续执行
		end:=time.Since(t)
		fmt.Printf("耗时%v\n",end)
		status :=c.Writer.Status()
		fmt.Println("状态",status)
	}
}
func TokenRequired() gin.HandlerFunc{
	return  func(c *gin.Context) {
		var token string
		for k,v:=range c.Request.Header{
			//if k == "x-token"{注意这里一定要首字母大写
			if k=="X-Token"{
				token=v[0]
			}else{
				fmt.Println(k,v)
			}
		}
		if token != "bobby"{
			c.JSON(http.StatusUnauthorized,gin.H{
				"msg":"未登录",
			})
			return//在这里无法终止后续流程;有层级关系，终止的是当前返回匿名函数执行流程，它只是没有调用c.Next()而不是整条调用链
			//c.Abort()//必须如此才能终止后续流程 ；终止整个调用链
		}
		//c.Next()
	}
}
func main() {
	router := gin.Default()
	router.Use(TokenRequired())//框架调用
	router.GET("/ping",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
	router.Run(":8083")
}