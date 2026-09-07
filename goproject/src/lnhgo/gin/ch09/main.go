package main

import (
	//"fmt"
	"net/http"
	//"os"
	//"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//dir,_ := filepath.Abs(filepath.Dir(os.Args[0]))
	//fmt.Println(dir)
	//LoadHTMLFiles 将指定文件目录加载好，但是目录是相对目录
	//router.LoadHTMLGlob("gin/ch09/templates/*")//加载指定目录下所有文件
	router.LoadHTMLGlob("gin/ch09/templates/**/*")//加载包括子目录下的所有文件
	//router.LoadHTMLFiles("gin/ch09/templates/index.tmpl","gin/ch09/templates/goods.html")//加载指定的两个文件
	router.GET("/index",func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":"慕课网",
		})
	})
	router.GET("/goods/list",func(c *gin.Context) {
		c.HTML(http.StatusOK,"goods/list.html",gin.H{
			"title":"慕课网",
		})
	})
	router.GET("/user/list",func(c *gin.Context) {
		c.HTML(http.StatusOK,"user/list.html",gin.H{
			"title":"慕课网",
		})
	})
	router.GET("/goods",func(c *gin.Context) {
		c.HTML(http.StatusOK,"goods.html",gin.H{
			"name":"bobby",
		})
	})
	router.Run(":8082")
}