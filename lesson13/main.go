package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//获取请求的path(URL)参数，返回的都是字符串类型
//注意URL的匹配不要冲突

func main() {
	r:=gin.Default()

	r.GET("/user/:name/:age/:ds", func(c *gin.Context) {
		//获取参数路径
		name:=c.Param("name")
		age:=c.Param("age")   //string类型
		ds:=c.Param("ds")
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
			"ds":ds,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year:=c.Param("year")
		month:=c.Param("month")
		c.JSON(http.StatusOK,gin.H{
			"year":year,
			"month":month,
		})
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
