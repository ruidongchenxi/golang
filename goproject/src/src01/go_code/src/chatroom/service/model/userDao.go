package model

import (
	//"context"
	"encoding/json"
	"fmt"
	"src/chatroom/common/message"

	"github.com/gomodule/redigo/redis"
	_ "github.com/redis/go-redis/v9"
)

//希望在服务器启动时就初始化一个UserDao;做成全局的；需要和redis 交互时就直接使用即可
var (
	MyUserDao *UserDao
)


//定义一个UserDao 结构体；完成对user结构体各种操作
// type  struct{
// 	//ctx context.Context
// 	Pool *redis.Client
// }
// //思考方法
// //
// func (this *UserDao)GetUserById(conn context.Context,id int) (user *User){
// 	//查询用户
// 	user,err:=this.Pool.HGet(conn,"users",id).Result()
// }
// 使用工厂模式，创建一个USerDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return

}
type UserDao struct{
	pool *redis.Pool
}
func (this *UserDao)GetUserById(conn redis.Conn,id int)(user *User,err error){
	res,err:= redis.String(conn.Do("hget","users",id))
	if err != nil {
		if err == redis.ErrNil{//没有找到对应ID
			err= ERROR_USER_NOTEXISTS
		}
		return
	}
	//这里需要把res反序列成User实例
	user = &User{}
	err =json.Unmarshal([]byte(res),user)
	if err != nil{
		fmt.Println("反序列化出错",err)
		return
	}
	return
} 
//完成登录校验
//1.Login 完成用户的验证
//2.如果用户的id和pwd都正确，则返回一个user实例
//3. 如果用户的id 或ped有错误，则返回错误信息

func (this *UserDao)Login(userId int,userPwd string) (user *User,err error){
	//先从UserDaon的连接池去一个连接
	conn:= this.pool.Get()
	defer conn.Close()
	user,err=this.GetUserById(conn,userId)
	if err !=nil{
		return
	}
	//这时用户证明是存在的；该判断用户的密码
	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		return
	} 
	return
}



func (this *UserDao)Register(User *message.User) (err error){
	//先从UserDaon的连接池去一个连接
	conn:= this.pool.Get()
	defer conn.Close()
	_,err=this.GetUserById(conn,User.UserID)
	if err ==nil{
		err=ERROR_USER_EXISTS
		return
	}
	//这时用户证明还没有注册
	// if user.UserPwd != userPwd{
	// 	err = ERROR_USER_PWD
	// 	return
	// } 
	data,err:=json.Marshal(User)//序列化
	if err !=nil {
		return
	}
	//入库
	_,err =conn.Do("hset","users",User.UserID,string(data))
	if err != nil {
		fmt.Println("保存注册用户出错",err)
		return
	}
	return
}
