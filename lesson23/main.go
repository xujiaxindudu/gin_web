package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// User 1.定义模型
type User struct {
	gorm.Model
	Name string
	Age int64
	Active bool
}

func main() {
	//2.连接mysql数据库
	db,err:=gorm.Open("mysql","root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("conn sql failed,err:",err)
		return
	}
	defer db.Close()

	//3.把模型和数据库中的表对应起来
	db.AutoMigrate(&User{})

	//4.创建
	//u1:=User{Name: "q1mi",Age:18,Active: true}
	//db.Create(&u1)
	//u2:=User{Name: "jinzhu",Age:20,Active: true}
	//db.Create(&u2)

	//5.查询
	var user User
	db.First(&user)

	//6.更新
	user.Name="七米"
	user.Age=29
	db.Debug().Save(&user)

	db.Debug().Model(&user).Update("name","小王子")

	db.Debug().Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 29, "active": false})

	db.Debug().Model(&user).Updates(map[string]interface{}{"name":"liwenzhou","age":28,"active":true}) //列出来的所有字段都会更新

	db.Debug().Model(&user).Select("age").Update(map[string]interface{}{"name":"dsa","age":30,"active":true}) //只更新age字段

	db.Debug().Model(&user).Omit("active").Update(map[string]interface{}{"name":"d3213","age":21,"active":false}) //排除active只更新其他字段

	db.Debug().Model(user).UpdateColumns(User{Name:"hello",Age:18})

	db.Debug().Model(&User{}).Update("age",gorm.Expr("age+?",2))











}
