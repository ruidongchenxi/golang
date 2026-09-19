package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc{
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin","*")//允许任意 Origin 访问
		c.Header("Access-Control-Allow-Headers","Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,x-token")//你请求的时候可以带什么 Header
		c.Header("Access-Control-Allow-Methods","POST,GET,OPTIONS,DELETE,PATCH,PUT")//允许跨域使用这些 HTTP 方法
		c.Header("Access-Control-Expose-Headers","Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type,Authorization")//告诉浏览器：允许前端 JavaScript 读取响应中的哪些响应头。
		c.Header("Access-Control-Allow-Credentials","true")
		if method == "OPTIONS"{
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}