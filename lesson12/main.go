package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//获取form表单提交的参数

func main() {
	r:=gin.Default()
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})
	r.POST("/login", func(c *gin.Context) {
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name":username,
			"Password":password,
		})
	})
	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}