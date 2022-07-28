package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

var (
	DB *gorm.DB
)

// Todo 模型
type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func InitMySQL()(err error){
	dsn:="root:root1234@tcp(127.0.0.1:13306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB,err=gorm.Open("mysql",dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	//创建数据库
	//连接数据库
	err:=InitMySQL()
	if err != nil {
		log.Println("connect mysql failed,err:",err)
		return
	}
	defer DB.Close()
	//模型与数据库中的表对应起来
	DB.AutoMigrate(&Todo{})

	r:=gin.Default()
	r.Static("/static","static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	v1Group:=r.Group("v1")
	{
		//添加待办事项
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面输入待办事项后，请求发送到这里
			//将请求拿出来
			var todo Todo
			c.BindJSON(&todo)
			//存入数据库
			//做出响应
			if err=DB.Create(&todo).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else{
				c.JSON(http.StatusOK,todo)
			}
		})
		//查询待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			//查询这个表里所有的数据
			var todoList []Todo
			if err=DB.Find(&todoList).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else{
				c.JSON(http.StatusOK,todoList)
			}
		})
		//修改待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id,ok:=c.Params.Get("id")
			if !ok{
				c.JSON(http.StatusOK,gin.H{"error":"无效ID"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err!=nil{
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err!= nil{
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			}else{
				c.JSON(http.StatusOK, todo)
			}
		})

		//删除待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id ,ok:=c.Params.Get("id")
			if !ok{
				c.JSON(http.StatusOK,gin.H{"error":"无效id"})
				return
			}
			if err:=DB.Where("id=?",id).Delete(Todo{}).Error;err!=nil{
				c.JSON(http.StatusOK,gin.H{"error":err.Error()})
			}else{
				c.JSON(http.StatusOK,gin.H{id:"delete"})
			}
		})

	}


	err=r.Run()
	if err != nil {
		log.Println("http server start failed,err",err)
		return
	}


}
