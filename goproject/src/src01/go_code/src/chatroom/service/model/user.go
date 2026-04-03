package model
//定义用户结构体
type User struct{
	//为了序列化跟反序列化成功必须保证
	//用户信息的json 字符串的key和结构体的字段对应tag名字一致，否则失败
	UserID int `json:"userID"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}