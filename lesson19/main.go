package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type UserInfo struct {
	ID int
	Name string
	Gender string
	Hobby string
}

func main() {
	//连接MySQL数据库
	db,err:=gorm.Open("mysql","root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("open mysql failed,err:",err)
	}
	defer db.Close()

	//创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	u1 := UserInfo{1, "七米", "男", "篮球"}
	db.Create(&u1)

	//查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n",u)

	// 更新
	db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)
}
