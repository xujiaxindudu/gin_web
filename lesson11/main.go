package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r:=gin.Default()
	//querystring

	//GET请求URL ？后面的是querystring等参数
	//key=value，多个key-value用 & 符号连接
	//eq: /web?query=yiyi&age=3
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的query string 参数
		//name:=c.Query("query")                             //1.通过Query获取请求中携带的querystring参数

		name:=c.DefaultQuery("query","嘟嘟")  //2.取不到就用指定的默认值
		age:=c.DefaultQuery("age","3")

		//name,ok:=c.GetQuery("query")                       //3.取到返回（值，true）取不到第二个参数就返回（""，false）
		//if !ok{
		//	//取不到
		//	name="嘟嘟"
		//}
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
