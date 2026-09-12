package main

import (
	"context"
	"fmt"
	"mxshop_srvs/user_srv/proto"

	//"os/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var userClient proto.UserClient
var conn *grpc.ClientConn
func Init(){
	var err error
	conn, err = grpc.NewClient(
    "127.0.0.1:50051",
    grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err !=nil{
        panic("连接失败")
    }
    
	userClient=proto.NewUserClient(conn)

}

func TestGetUserList() {
	res,err:=userClient.GetUserList(context.Background(),&proto.PageInfo{
		Pn: 1,
		PSize: 2,
	})
	if err != nil{
		panic(err)
	}
	//fmt.Println(res.Data)
	for _,user:=range res.Data{
		fmt.Printf(user.Mobile,user.NickName,user.Password)
		checkRes,err := userClient.CheckPassWord(context.Background(),&proto.PasswordCheckInfo{
			Password: "generic password",
			EncryptedPassword: user.Password,
		})
		if err !=nil{
			panic(err)
		}
		fmt.Println(checkRes.Success)
	}


}
func TestCreateUser(){
	for i := 10;i<20;i++{
		rsp,err:= userClient.CreateUser(context.Background(),&proto.CreateUserInfo{
			NickName: fmt.Sprintf("bobby%d",i),
			Mobile: fmt.Sprintf("187822222%d",i),
			PassWord: "admin123",
		})
		if err != nil{
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}
func main(){
	Init()
	TestGetUserList()
	//TestCreateUser()
	defer conn.Close()
}