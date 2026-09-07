package main
import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()
	//匹配的url格式：/welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome",func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname","Guest")
		//lastname := c.Query("lastname")//是c.Request.URL.Query().Get("lastname")
		lastname:=c.DefaultQuery("lastname","imooc")
		// c.String(http.StatusOK,"Hello %s %s",firstname,lastname)
		c.JSON(http.StatusOK,gin.H{
			"first_name":firstname,
			"last_name":lastname,
		})
	})
	router.POST("/form_post",func(c *gin.Context) {
		message:=c.PostForm("message")
		nick:= c.DefaultPostForm("nike","anonymous")
		c.JSON(http.StatusOK,gin.H{
			"message":message,
			"nick":nick,
		})
	})
	router.POST("/post",func(c *gin.Context) {
		id := c.Query("id")//获取url参数
		page:=c.DefaultQuery("page","0")
		name:=c.PostForm("name")//获取表单里内容
		message:=c.DefaultPostForm("massage","信息")
		c.JSON(http.StatusOK,gin.H{
			"id":id,
			"page":page,
			"name":name,
			"message":message,
		})
	})
	router.Run(":8084")
}