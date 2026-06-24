package main

import (
//   "context"
//   "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}


func main() {
//   db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//   if err != nil {
//     panic("连接数据库失败")
//   }

//   ctx := context.Background()

//   // 自动建表
//   db.AutoMigrate(&Product{})

//   // 创建
//   err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})

//   // 查询
//   product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx) // 查找对应主键的产品
//   products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx) // 查找 code 为 D42 的所有产品

//   // 更新 - 将产品价格更新为 200
//   err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
//   // 更新 - 更新多个字段
//   err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, Product{Code: "D42", Price: 100})

//   // 删除 - 删除产品
//   err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
/*
1、代码规范
   代码规范不是强制，但是不同语言一些细微规范还是遵守
   代码规范主要是方便团队内部形成一个统一的代码风格，提高代码可读性统一性
   命名规范
      包名：尽量和目录名保持一致；采取有意义的包名，尽量简短，尽量和标准库重名尽量采用小写
	  文件名： 如果对个单词可以采用蛇形命名
	  变量名：
		蛇形：python php
		驼峰：go c java；go首字母小写
		一些专有 URL
	  结构体
	    驼峰： 
	  接口命名
	    接口名和结构体差不多
      常量命令：全部大写

2、注释规范
  go 提供两种注释行注释 、块注释
3、import 规范

*/
}