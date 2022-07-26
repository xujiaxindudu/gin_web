package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// User 1.定义模型
type User struct{
	gorm.Model
	Name string
	Age int64
}

func main(){
	//连接MySQL数据库
	db,err:=gorm.Open("mysql","root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("conn sql failed,err:",err)
		return
	}
	defer db.Close()

	//2.把模型数据库中的表对应起来
	db.AutoMigrate(&User{})

	//3.创建
	//u1:=User{Name: "q1mi",Age:18}
	//u2:=User{Name: "jinzhu",Age:20}
	//db.Create(&u1)
	//db.Create(&u2)

	//4.查询
	//var user User   //声明模型结构体类型变量user  （文件夹A）
	////根据主键查询第一条记录
	//db.First(&user)       //（文件夹B）
	//fmt.Printf("user:%#v\n",user)

	//user:=new(User) //
	////db.Where("name=?","jinzhu").First(&user)
	//db.Where("name=?","q1mi").Find(&user)
	////db.First(&user)
	//fmt.Printf("user:%#v\n",user)
	//
	//
	//var users []User
	//db.Debug().Find(&users)
	//fmt.Printf("user:%#v\n",users)

	//FirstOrInit
	var user User
	//db.Attrs(User{Age: 99}).FirstOrInit(&user,User{Name: "q1mi"})
	db.FirstOrCreate(&user, User{Name: "non_existing"})
	fmt.Printf("user:%#v\n",user)



}