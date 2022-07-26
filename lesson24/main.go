package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	//gorm.Model
	Id int64
	Name string
	Age int64
	Active bool
}

func main() {
	db,err:=gorm.Open("mysql","root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("conn sql failed,err:",err)
		return
	}
	defer db.Close()

	//3.把模型和数据库中的表对应起来
	db.AutoMigrate(&User{})

	//4.创建
	//u1:=User{Name: "q1mi1",Age:18,Active: true}
	//db.Create(&u1)
	//u2:=User{Name: "jinzhu1",Age:20,Active: true}
	//db.Create(&u2)

	//5.删除
	//var u User
	//u.ID=1
	//u.Name="jinzu1"
	//db.Debug().Delete(&u)

	db.Debug().Where("name=?","jinzhu1").Delete(User{})
	//
	//db.Delete(User{},"age=?",18)
	//var u2 []User
	//db.Debug().Unscoped().Where("name=?","jinzhu1").Find(&u2)
	//fmt.Println(u1)

	//物理删除
	//db.Debug().Unscoped().Where("name=?","jinzhu1").Delete(User{})
}
