package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
type UserInfo struct {
	username string
	password string

}
func main() {
	r:=gin.Default()
	r.GET("/user", func(c *gin.Context) {
		username:=c.Query("username")
		password:=c.Query("password")
		u:=UserInfo{
			username,
			password,
		}
		fmt.Printf("%#v\n",u)
		c.JSON(http.StatusOK,gin.H{
			"msg":"ok",
		})
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}