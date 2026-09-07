package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"strings"
	"time"

	//"fmt"
	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/model"
	"mxshop_srvs/user_srv/proto"

	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)
type UserServer struct{
	proto.UnsafeUserServer
}
func Paginate(page ,pageSize int ) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    if page <= 0 {
      page = 1
    }

    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := (page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}


func MobileToResponse(User model.User) *proto.UserInfoResPonse{
	//在grpc的message中字段有默认值；不能随便赋值nil进去，容易出错
	//搞清楚哪些字段有默认有默认值
	userInfoRsp:=&proto.UserInfoResPonse{
		Id: User.ID,
		Password: User.Password,
		NickName:User.NickName,
		Gender: User.Gender,
		Role: int32(User.Role),
	}
	if User.Birthday!=nil{
		userInfoRsp.BirthDay = uint64(User.Birthday.Unix())
		}
	return  userInfoRsp
}
func (s *UserServer)GetUserList(cx context.Context, req *proto.PageInfo) (*proto.UserListResPonse, error){
	//获取用户列表
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error !=nil {
		return  nil,result.Error
	}
	rsq := &proto.UserListResPonse{}
	rsq.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(req.Pn),int(req.PSize))).Find(&users)
	for _,user:=range users{
		userInfoRsp:=MobileToResponse(user)
		rsq.Data= append(rsq.Data,userInfoRsp)
		
	}
	return rsq,nil

}
func (s *UserServer)GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResPonse, error){
	//通过手机号码查询用户
	var user model.User
	result :=global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0{
		return nil,status.Errorf(codes.NotFound,"用户不存在")
	}
	if result.Error !=nil{
		return nil,result.Error
	}
	userInfoResPonse := MobileToResponse(user)
	return userInfoResPonse,nil

}
func (s *UserServer)GetUserByID(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResPonse, error){
	//通过ID查用户
	var user model.User
	result :=global.DB.First(&user,req.Id)
	if result.RowsAffected == 0{
		return nil,status.Errorf(codes.NotFound,"ID不存在")
	}
	if result.Error !=nil{
		return nil,result.Error
	}
	userInfoResPonse := MobileToResponse(user)
	return userInfoResPonse,nil
}
func (s *UserServer)CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResPonse, error){
	//新建用户
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)//将查询到的信息填充到user结构体里
	if result.RowsAffected ==1{
		return nil,status.Errorf(codes.AlreadyExists,"用户已存在")
	}
	user.Mobile = req.Mobile
	user.NickName = req.NickName
	//密码加密
	options := &password.Options{SaltLen: 10, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(req.PassWord, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s",salt,encodedPwd)
	result = global.DB.Create(&user)

	if result.Error !=nil{
		return nil, status.Errorf(codes.Internal,result.Error.Error())
	}

	userInfoRsp:=MobileToResponse(user)
	return  userInfoRsp,nil
}
func (s *UserServer)UpdateUser(ctx context.Context,req *proto.UpdateUserInfo) (*emptypb.Empty, error){
	//个人中信更新用户
	var user model.User
	result := global.DB.First(&user,req.Id)
	if result.RowsAffected ==0{
		return  nil,status.Error(codes.NotFound,"用户不存在")
	}
	birthDay:=time.Unix(int64(req.BirthDay),0)
	user.NickName = req.NickName
	user.Birthday=  &birthDay
	user.Gender=req.Gender
	result= global.DB.Save(user)
	if result.Error !=nil{
		return  nil,status.Error(codes.Internal,result.Error.Error())
	}
	return  &emptypb.Empty{},nil
}
func (s *UserServer)CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckReponse, error){
	//校验密码
	passwordInfo:=strings.Split(req.EncryptedPassword,"$")
	// fmt.Println(passwordInfo)
	options := &password.Options{SaltLen: 10, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	//  fmt.Println(check) // true
	return &proto.CheckReponse{Success: check},nil
}