package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)
var trans ut.Translator
type LoginForm struct {
	User     string `json:"user"  binding:"required,min=3,max=10"` //前端传递的模式；各种类型数据
	Password string `json:"password" binding:"required"`
}
type SignUpForm struct{
	Age uint8 `json:"age" binding:"required,gte=1,lte=130"`//表示大于1；或者小于130
	Name string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`//跨字段验证
}
func removeTopStruct(fileds map[string]string) map[string]string{
	rsp :=map[string]string{}
	for field,err:=range fileds{
		rsp[field[strings.Index(field,".")+1:]]=err
	}
	return rsp
}
func InintTrans(locale string)(err error){
	//修改gin框架中的validator引擎属性，实现定制
	if v,ok :=binding.Validator.Engine().(*validator.Validate);ok{
		//注册一个获取json的tag的自定义来                                                  
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name:= strings.SplitN(fld.Tag.Get("json"),",",2)[0]
			if name=="-"{
				return ""
			}
			return name
		})
		zhT:= zh.New()//中文
		enT:=en.New()//英文翻译
		//第一个参数备用语言环境，后续参数应该支持的语言环境
		uni:=ut.New(enT,zhT,enT)
		if trans,ok =uni.GetTranslator(locale);!ok{
			return fmt.Errorf("GetTranslator(%s)",locale)
		}else{
			switch locale{
			case "en":
				enTranslations.RegisterDefaultTranslations(v,trans)
			case "zh":
				zhTranslations.RegisterDefaultTranslations(v,trans)
			default:
				enTranslations.RegisterDefaultTranslations(v,trans)

			}
			return 
		}
		//return

	}
	return 
}

func main() {
	//解决某个时间问题添加代码有很大的代码侵入性。中间件
	if err := InintTrans("zh");err!=nil{
		fmt.Println("初始化获取翻译器错误")
		return
	}
	routes := gin.Default()
	routes.POST("/loginJSON",func(c *gin.Context) {
		var loginFor LoginForm
		if err:=c.ShouldBind(&loginFor);err!=nil{
			errs,ok:=err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK,gin.H{
					"mgs":err.Error(),
				})
			}
			c.JSON(http.StatusOK,gin.H{
				"error":removeTopStruct(errs.Translate(trans)),
			})
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return 
		}
		c.JSON(http.StatusOK,gin.H{
			"mag":"登录成功",
		})
	})
	routes.POST("/signup",func(c *gin.Context) {
		var SignUpForm SignUpForm
		if err := c.ShouldBind(&SignUpForm);err !=nil{
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"msg":"登录成功",
		})
	})
	routes.Run(":8083")
}