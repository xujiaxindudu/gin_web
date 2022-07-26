package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

// User 定义模型
type User struct {
	gorm.Model   //内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}

// Animal 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// TableName 更改表名，唯一指定表名
func (Animal)TableName()string{
	return "dudu"
}

func main(){
	//修改默认表名规则
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "SMS_" + defaultTableName
	}

	db,err:=gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("connect mysql err:",err)
		return
	}
	defer db.Close()
	//禁止表名使用复数形式
	db.SingularTable(true)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})
	//使用User结构体创建名为'yiyi'的表
	//db.Table("yiyi").CreateTable(&User{})





}