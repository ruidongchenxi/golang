package model

import (
	"time"

	"gorm.io/gorm"
)
/*
1.密码 2.密文不可反解
1.对称加密（加解密同一把钥匙）
2.非对称加密
3.md信息摘要算法加密
密码如果不可反解用户忘记密码怎么处理

*/
type BaseModel struct {
	ID        int32 `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:Update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool 

}
type User struct{
	BaseModel
	Mobile string `gorm:"index:idx_mobile;type:varchar(11);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	NickName string `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender string `gorm:"column:gender;default:male;type:varchar(6) comment 'female 表示女，male表示男'"`
	Role int `gorm:"column:role;default:1;type:int comment '1表示普通用户 2表示管理员'"`

}