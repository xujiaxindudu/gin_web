package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//返回默认的路由引擎
	//r :=gin.Default()   //默认调用了Logger()和Recovery()
	r:=gin.New()
	//GET:请求方式；/hello：请求的路径
	//客户端浏览器等访问/hello路径时，会执行后面的匿名函数
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"msg":"GET",
		})
	})


	//启动http服务，默认端口8080
	err:=r.Run()
	if err != nil {
		log.Println(err)
		return
	}
}