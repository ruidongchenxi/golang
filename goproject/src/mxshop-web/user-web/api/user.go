package api

import (
	"context"
	"fmt"
	"mxshop-web/user-web/forms"
	"mxshop-web/user-web/global"
	"mxshop-web/user-web/global/reponse"
	proto "mxshop-web/user-web/prote"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	//proto "mxshop-web/user-web/proter"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)
var trans ut.Translator
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
func HandleGrpcErrorToHttp(err error,c *gin.Context){
	//将grpc 的code转换为http状态码
	if err != nil{
		if e,ok :=status.FromError(err);ok{
			switch e.Code(){
			case codes.NotFound:
				c.JSON(http.StatusNotFound,gin.H{
					"msg":e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError,gin.H{
					"msg":"内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest,gin.H{
					"msg":"参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError,gin.H{
					"msg":"用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError,gin.H{
					"msg":"其他错误",
				})
			}
			return 
			
		}
	}

}

func GetUserList(cxt *gin.Context){
	// ip := "127.0.0.1"
	// port:= 50051
	//拨号连接
	userConn,err :=grpc.NewClient(
		fmt.Sprintf("%s:%d",global.ServerConfig.UserSrvInfo.Host,global.ServerConfig.UserSrvInfo.Prot),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err!=nil{
		zap.S().Errorw("连接用户服务失败","msg",err.Error(),)
	}
	//生成grpc 的client并调用接口
	userSrvClient:=proto.NewUserClient(userConn)
	pn:= cxt.DefaultQuery("pn","0")//设置默认值
	pnInt,_:=strconv.Atoi(pn)
	pSize:=cxt.DefaultQuery("psize","10")
	pSizeInt,_:= strconv.Atoi(pSize)
	rsq,err:=userSrvClient.GetUserList(context.Background(),&proto.PageInfo{
		Pn: uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil{
		zap.S().Errorw("[GetUserList] 查询【用户列表失败】")
		HandleGrpcErrorToHttp(err,cxt)
		return 
	}
	result := make([]interface{},0)
	for _,value := range rsq.Data{
		//data := make(map[string]interface{})
		user := reponse.UserResPonse{
			Id: value.Id,
			NickName: value.NickName,
			BirthDay: reponse.JsonTime(time.Unix(int64(value.BirthDay),0)),
			Gender: value.Gender,
			Mobile: value.Mobile,
		}
		result=append(result, user)
	}
	
	cxt.JSON(http.StatusOK,result)
}
func PassWordLogin(c *gin.Context){
	//表单验证
	PassWordLoginForm := forms.PassWordLoginForm{}
	if err := c.ShouldBindJSON(&PassWordLoginForm);err!=nil{
		//如何返回错误信息
		
	}
}