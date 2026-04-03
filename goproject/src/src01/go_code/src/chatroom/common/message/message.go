package message

//import "go/types"

// import(
// 	"src/chatroom/service/model"
// )
const(
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes" 
	RegisterResMesType = "RegisterResMes"

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
type RegisterMes struct{
	User User `json:"user"` 	
}
type RegisterResMes struct{
	Code int `json:"code"`// 返回状态码 400表示用户占用；200 表示注册成功
	Error string `json:"error"`
}