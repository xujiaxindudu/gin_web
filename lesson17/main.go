package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r:=gin.Default()

	//路由组的组
	//把公用的前缀提取出来，创建一个路由组
	
	userGroup:=r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {})
		userGroup.GET("/login", func(c *gin.Context) {})
		userGroup.POST("/login", func(c *gin.Context) {})
	}

	shopGroup:=r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {})
		shopGroup.GET("/cart", func(c *gin.Context) {})
		shopGroup.POST("/checkout", func(c *gin.Context) {})
	}
	
	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}