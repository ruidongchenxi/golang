package message

//import "go/types"
//
const(
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"

)

type Message struct {
	Type  string 	`json:"typt"`//消息类型
	Data  string  `json:"data"`//内容

}
//先定义消息，后面需要增加
type LoginMes struct{
	UserId int  `json:"userId"`
	UserPwd string  `json:"userPwd"`
	UserName string  `json:"userName"`
}
type LoginResMes struct{
	Code  int  `json:"code"`//返回状态码 500表示用户未注册；200表示登录成功；
	Error string `json:"error"`//
}