package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//CRUD增删改查

// User 1.定义模型
type User struct{
	ID int64
	Name string  `gorm:"default:'嘟嘟'"`
	Age int64
}

func main() {
	db,err:=gorm.Open("mysql","mysql","root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("conn mysql failed,err:",err)
	}
	defer db.Close()
	//2.把模型与数据库中的表对应起来
	db.AutoMigrate(User{})

	//3.创建

	u:=User{
		Age:12,
	}  //代码层面创建一个User对象,如果设置了默认Name，零值会被忽略



	fmt.Println(db.NewRecord(&u))  //判断主键是否为空 true
	db.Debug().Create(&u)                  //在数据库中创建了一条q1mi 18 的记录
	fmt.Println(db.NewRecord(&u))  //判断主键是否为空 false
}

