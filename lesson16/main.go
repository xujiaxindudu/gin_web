package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r:=gin.Default()

	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK,gin.H{
		//	"status":"ok",
		//})
		//跳转到sogo.com
		c.Redirect(http.StatusMovedPermanently,"http://www.sogo.com")
	})

	r.GET("/a", func(c *gin.Context) {
		//跳转到/b对应的处理路由函数
		c.Request.URL.Path="/b"    //把请求的URL修改
		r.HandleContext(c)     //继续后续的处理

	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"b",
		})
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}
