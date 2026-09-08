package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mxshop_srvs/user_srv/model"
	"os"
	//"os/user"
	//"strings"
	"time"

	// "log"
	// "os"
	// "time"

	// "gorm.io/driver/mysql"
	//"google.golang.org/grpc/resolver/dns"
	// "mxshop_srvs/user_srv/model"

	// "gorm.io/gorm"
	// "gorm.io/gorm/logger"
	// "gorm.io/gorm/schema"
	password "github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)
func genMd5(code string)string{
	Md5:=md5.New()
	_,_=io.WriteString(Md5,code)
	return  hex.EncodeToString(Md5.Sum(nil))
}
func main() {
	dsn := "root:123456@tcp(192.168.40.25:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout,"\r\n",log.LstdFlags),//io writer
		logger.Config{
			SlowThreshold: time.Second,//慢SQL阈值
			LogLevel: logger.Info,//Log level
			Colorful: true,//启用彩色打印
		},
	)
	db,err:=gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			
		},
		Logger: newLogger,

	})
	if err!=nil{
		panic(err)
	}
	_=db.AutoMigrate(&model.User{})
	fmt.Println(genMd5("123456"))
		// Using the default options
		//salt 表示生成的盐值
		//encodedPwd 生成的密码
	// salt, encodedPwd := password.Encode("generic password", nil)
	// fmt.Println(salt)//盐值多少
	// fmt.Println(encodedPwd)//加密后是多少

	// check := password.Verify("generic password", salt, encodedPwd, nil)//去验证
	// fmt.Println(check) // true

	// Using custom options；saltLen 随机盐值长度；Iterations：迭代次数；keyLen:key的长度
	options := &password.Options{SaltLen: 10, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newpassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s",salt,encodedPwd)
	fmt.Println(newpassword)
	// passwordInfo:=strings.Split(newpassword,"$")
	// fmt.Println(passwordInfo)
	// check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	//  fmt.Println(check) // true
	//  fmt.Println(len(newpassword))
	// fmt.Println(salt)//盐值多少
	// fmt.Println(encodedPwd)//加密后是多少
	for i :=0 ;i<10;i++{
		user := model.User{
			NickName: fmt.Sprintf("boobby%d",i),
			Mobile: fmt.Sprintf("1878222222%d",i),
			Password: newpassword,
		}
		db.Save(&user)
	}
}