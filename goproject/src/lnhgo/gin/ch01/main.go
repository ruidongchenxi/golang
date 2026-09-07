package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func goodsList(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"name":"goodsList",
	})
}
func goodsDetail(c *gin.Context){
	id:=c.Param("id")
	action:=c.Param("action")

	c.JSON(http.StatusOK,gin.H{
		"name":id,
		"action":action,
	})

}
func createGoods(c *gin.Context){}

func main() {
	router := gin.Default()
	goodsGroup:=router.Group("/goods")//路由分组
	//路由分组;{}这对括号可有可无
	{
		goodsGroup.GET("/list",goodsList)
		goodsGroup.GET("/:id/*action",goodsDetail)
		goodsGroup.POST("/add",createGoods)
	}
	// router.GET("/goods/list",goodsList)
	// router.GET("/goods/1",goodsDetail)
	// router.POST("/goods/add",createGoods)
	// v1:= router.Group("/v1")
	// {
	// 	v1.POST("/login",loginEndpoint)
	// 	v1.POST("/submit",submitEndpoint)
	// 	v1.POST("/read",readEndpoint)
	// }
	// v2:= router.Group("/v2")
	// {
	// 	v2.POST("/login",loginEndpoint)
	// 	v2.POST("/submit",submitEndpoint)
	// 	v2.POST("/read",readEndpoint)		
	// }
	router.Run(":8081")//执行端口；如果不指定默认是8080
}