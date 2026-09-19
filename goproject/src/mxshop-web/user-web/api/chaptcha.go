package api

import (
	//"database/sql/driver"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)
var store = base64Captcha.DefaultMemStore
func GetCaptcha(cxt *gin.Context){
	driver:=base64Captcha.NewDriverDigit(80,240,5,0.7,80)
	cp :=base64Captcha.NewCaptcha(driver,store)
	id,b64,s,err:=cp.Generate()
	if err !=nil{
		zap.S().Errorf("验证码生成失败：",err.Error())
		cxt.JSON(http.StatusInternalServerError,gin.H{
			"msg":"验证码生成失败",
		})
		return
	}
	cxt.JSON(http.StatusOK,gin.H{
		"captchaID":id,
		"picPath":b64,
		"verifys":s,//验证码
	})


}