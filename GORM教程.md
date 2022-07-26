# GORM教程

## 一、ORM简介

O:Object对象，程序中的对象/实例；例如Go中的结构体实例

R:Relational关系，关系数据库：Mysql

M:映射

<img src="/Users/xujiaxin/Desktop/picture/截屏2022-07-15 18.04.13.png" alt="截屏2022-07-15 18.04.13" style="zoom:50%;" />

<u>ORM优缺点</u>：

优点：

- 提高开发效率

缺点：

- 牺牲执行性能
- 牺牲灵活性
- 弱化SQL能力

## 二、GORM基本实例

连接数据库进行增删改查

```go
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
```

## 三、GORM Model定义

在使用ORM工具时，通常我们需要在代码中定义模型(Models)与数据库中的数据表进行映射，在GORM中模型(Models)通常是正常 定义的结构体、基本的go类型或它们的指针。 同时也支持`sql.Scanner`及`driver.Valuer`接口（interfaces）。

### 1、gorm.Model

为了方便模型定义，GORM内置了一个`gorm.Model`结构体。`gorm.Model`是一个包含了`ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`四个字段的Golang结构体。

```go
// gorm.Model 定义
type Model struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}
```

### 2、模型定义示例

```go
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
```

### 3、主键、表名列名的约定

#### 3.1主键(Primary Key)

GORM默认会使用名为ID的字段作为表的主键。

```go
type User struct {
  ID   string // 名为`ID`的字段会默认作为表的主键
  Name string
}

// 使用`AnimalID`作为主键
type Animal struct {
  AnimalID int64 `gorm:"primary_key"`
  Name     string
  Age      int64
}
```

#### 3.2表名 (Table Name)

表名默认结构就是复数

```go
type User struct {} // 默认表名是 `users`

// 将 User 的表名设置为 `profiles`
func (User) TableName() string {
  return "profiles"
}

func (u User) TableName() string {
  if u.Role == "admin" {
    return "admin_users"
  } else {
    return "users"
  }
}

// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
db.SingularTable(true)
```

使用`Table()`指定表名：

```go
// 使用User结构体创建名为`deleted_users`的表
db.Table("deleted_users").CreateTable(&User{})

var deleted_users []User
db.Table("deleted_users").Find(&deleted_users)
//// SELECT * FROM deleted_users;

db.Table("deleted_users").Where("name = ?", "jinzhu").Delete()
//// DELETE FROM deleted_users WHERE name = 'jinzhu';
```

GORM还支持更改默认表名规则：

```go
gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
  return "prefix_" + defaultTableName;   //数据表前会同意增加“prefix_”
}
```

#### 3.3列名

列名由字段名称进行下划线分割来生成

```go
type User struct {
  ID        uint      // column name is `id`
  Name      string    // column name is `name`
  Birthday  time.Time // column name is `birthday`
  CreatedAt time.Time // column name is `created_at`
}
```

可以使用结构体tag指定列名：

```go
type Animal struct {
  AnimalId    int64     `gorm:"column:beast_id"`         // set column name to `beast_id`
  Birthday    time.Time `gorm:"column:day_of_the_beast"` // set column name to `day_of_the_beast`
  Age         int64     `gorm:"column:age_of_the_beast"` // set column name to `age_of_the_beast`
}
```

## 四、CRUD

CRUD通常指数据库的增删改查

### 1、创建

首先定义模型：

```
type User struct{
   ID int64
   Name *string 
}
```

使用使用`NewRecord()`查询主键是否存在，主键为空使用`Create()`创建记录：

```go
user := User{Name: "q1mi", Age: 18}

db.NewRecord(user) // 主键为空返回`true`
db.Create(&user)   // 创建user
db.NewRecord(user) // 创建`user`后返回`false`
```

2、通过tag定义字段的默认值：

```go
type User struct {
  ID   int64
  Name string `gorm:"default:'小王子'"`
  Age  int64
}
```

**注意：**通过tag定义字段的默认值，在创建记录时候生成的 SQL 语句会排除没有值或值为 零值 的字段。 在将记录插入到数据库后，Gorm会从数据库加载那些字段的默认值。

```go
var user = User{Name: "", Age: 99}
db.Create(&user)
```

上面代码实际执行的SQL语句是`INSERT INTO users("age") values('99');`，排除了零值字段`Name`，而在数据库中这一条数据会使用设置的默认值`小王子`作为Name字段的值。

**注意：**所有字段的零值, 比如`0`, `""`,`false`或者其它`零值`，都不会保存到数据库内，但会使用他们的默认值。 如果你想避免这种情况，可以考虑使用指针，比如：

```go
// 使用指针
type User struct {
  ID   int64
  Name *string `gorm:"default:'小王子'"`
  Age  int64
}
user := User{Name: new(string), Age: 18))}
/*
如果需要传实际的值
name:="dudu"
user := User{Name: &name, Age: 18))}
*/

db.Create(&user)  // 此时数据库中该条记录name字段的值就是''
```

### 2、查询

#### 2.1 一般查询

```go
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	gorm.Model
	Name string
	Age int64
}

func main() {
	db,err:=gorm.Open("mysql","root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("conn sql failed,err:",err)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	//u1:=User{Name: "q1mi",Age:18}
	//u2:=User{Name: "jinzhu",Age:20}
	//db.Create(&u1)
	//db.Create(&u2)

	//一般查询
	var user []User

	//根据主键查询第一条记录
	db.First(&user)
	fmt.Println(user)
	//随机获取一条记录
	db.Take(&user)
	fmt.Println(user)
	//根据主键查询最后一条记录
	db.Last(&user)
	fmt.Println(user)
	//根据主键查询所有的记录
	db.Find(&user)
	fmt.Println(user)
}
```

#### 2.2 Where查询

```go
  var user []User
	//得到第一条匹配记录
	db.Debug().Where("name=?","jinzhu").First(&user)
	fmt.Println(user)
	//得到所有匹配记录
	db.Debug().Where("name=?","jinzhu").Find(&user)
	fmt.Println(user)
```

### 3、更新

```go
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
```

***注意***：如果想要更新数据表中所有的名为“q1mi”字段的，需要这么操作：

```mysql
+----+---------------------+---------------------+------------+-------+------+--------+
| id | created_at          | updated_at          | deleted_at | name  | age  | active |
+----+---------------------+---------------------+------------+-------+------+--------+
|  1 | 2022-07-26 15:43:21 | 2022-07-26 15:43:21 | NULL       | q1mi  |   18 |      1 |
|  2 | 2022-07-26 15:43:21 | 2022-07-26 15:47:07 | NULL       | hello |   20 |      1 |
|  3 | 2022-07-26 15:43:31 | 2022-07-26 15:43:31 | NULL       | q1mi  |   88 |      1 |
|  4 | 2022-07-26 15:43:31 | 2022-07-26 15:47:07 | NULL       | hello |   70 |      1 |
+----+---------------------+---------------------+------------+-------+------+--------+
```

```go
db.Model(&User{}).Where("name = ?", "q1mi").Update("name", "hello")
```



### 4、删除

```go
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
```