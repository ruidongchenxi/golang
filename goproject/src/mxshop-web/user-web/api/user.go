package api

import (
	"context"
	"fmt"

	//"go/token"
	"mxshop-web/user-web/forms"
	"mxshop-web/user-web/global"
	"mxshop-web/user-web/global/reponse"
	"mxshop-web/user-web/middlewares"
	"mxshop-web/user-web/models"
	proto "mxshop-web/user-web/prote"
	"net/http"

	//"os/user"
	"strconv"
	"strings"
	"time"

	//proto "mxshop-web/user-web/proter"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)
func removeTopStruct(fileds map[string]string) map[string]string{
	rsp :=map[string]string{}
	for field,err:=range fileds{
		rsp[field[strings.Index(field,".")+1:]]=err
	}
	return rsp
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
func HandleValidatorError(c *gin.Context,err error){
			errs,ok:=err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK,gin.H{
					"mgs":err.Error(),
				})
			}
			c.JSON(http.StatusOK,gin.H{
				"error":removeTopStruct(errs.Translate(global.Trans)),
			})
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

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
	//
	claims,_:=cxt.Get("claims")
	currentUser:=claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d",currentUser.ID)
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
	if err:=c.ShouldBind(&PassWordLoginForm );err!=nil{
			HandleValidatorError(c,err)
			return 
	}
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
	//登录逻辑实现
	if rsp,err:=userSrvClient.GetUserByMobile(context.Background(),&proto.MobileRequest{
		Mobile: PassWordLoginForm.Mobile,

	}); err != nil{
		if e,ok:= status.FromError(err);ok{
			switch e.Code(){
			case codes.NotFound:
				c.JSON(http.StatusBadRequest,map[string]string{
					"mobile": "用户不存在",	
				})
			default:
				c.JSON(http.StatusInternalServerError,map[string]string{
					"mobile": "登录失败",
				})
			}
			return 
		}
	}else {
		if passRsp,pasErr:= userSrvClient.CheckPassWord(context.Background(),&proto.PasswordCheckInfo{
			Password: PassWordLoginForm.PassWord,
			EncryptedPassword: rsp.Password,
		}); pasErr!=nil{
			c.JSON(http.StatusInternalServerError,map[string]string{
				"password":"登录失败",
			})
		}else{
			if passRsp.Success{
				//生成tonken
				j:=middlewares.NewJWT()
				claims :=models.CustomClaims{
					ID: uint(rsp.Id),
					NickName: rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),//签名生效时间
						ExpiresAt: time.Now().Unix()+60*60*24*30, //30天过期
						Issuer: "imooc",
					},

				}
				token,err:=j.CreateToken(claims)
				if err !=err{
					c.JSON(http.StatusInternalServerError,gin.H{
						"msg":"生成tonken失败",
					})
					return
				}
				c.JSON(http.StatusOK,gin.H{
					"id":rsp.Id,
					"nick_name": rsp.NickName,
					"token":token,
					"expired_at": (time.Now().Unix()+60*60*24*30)*1000,
				})
			}else{
				c.JSON(http.StatusBadRequest,map[string]string{
					"msg":"登录失败",
				})
			}
		}
	}
	
}