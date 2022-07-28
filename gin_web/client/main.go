package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)
func handle(c *gin.Context){

	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/ws", nil)
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	for {
		err=conn.WriteMessage(websocket.TextMessage,[]byte("dudu"))
		if err != nil {
			log.Println("write failed,err",err)
			return
		}
	}
}

func main() {
	r:=gin.Default()
	r.GET("/",handle)
}