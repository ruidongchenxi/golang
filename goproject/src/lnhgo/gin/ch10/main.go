package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	//"sync"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()
	routes.GET("/",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"pong",
		})
	})
	go func ()  {
		routes.Run(":8083")	
	}()
	//如果想要接收kill-9强杀命令；不会处理后续逻辑的
	quit := make(chan os.Signal,1)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	fmt.Println("关闭中")
	fmt.Println("注销操作")

}