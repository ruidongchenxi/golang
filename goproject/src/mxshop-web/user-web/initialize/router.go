package initialize

import (
	"mxshop-web/user-web/middlewares"
	router2 "mxshop-web/user-web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine{
	Router := gin.Default()
	Router.Use(middlewares.Cors())//配置跨域
	ApiGroup := Router.Group("/u/v1")
	router2.InitUserRouter(ApiGroup)
	return  Router
}