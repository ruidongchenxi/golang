package router

import (
	"mxshop-web/user-web/api"
	"mxshop-web/user-web/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup) {
	//UserRouter:=Router.Group("/user").Use(middlewares.JWTAuth()) 一组url
	UserRouter:=Router.Group("/user")
	zap.S().Info("配置用户相关url")
	{
	UserRouter.GET("list",middlewares.JWTAuth(),middlewares.IsAdminAuth(),api.GetUserList)//某个url 访问的限制
	UserRouter.POST("pwd_login",api.PassWordLogin)
	}
}