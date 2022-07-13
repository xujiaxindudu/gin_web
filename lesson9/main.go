package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)
//静态文件：
//html页面上用到的样式文件：.css js文件 图片
func main() {
	r:=gin.Default()
	//定义模板
	//加载静态文件
	r.Static("/xxx","./statics")
	//gin框架给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	//解析模板
	//r.LoadHTMLFiles("templates/posts/index.tmpl","templates/users/index.tmpl")
	r.LoadHTMLGlob("templates/**/*") //**表示目录；*表示文件

	//渲染模板
	r.GET("/posts/index", func(c *gin.Context) {
		//http请求
		//gin.H就是一个map[string]interface{}
		c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
			"title":"posts/index.tmpl",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
			"title":"<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	//返回从网上下载的模板
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK,"home.html",nil)
	})
	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
