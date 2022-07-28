package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
var up=websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
func ws(c *gin.Context){
	conn,err:=up.Upgrade(c.Writer,c.Request,nil)
	if err != nil {
		log.Println("upgrade websocket failed,err:",err)
		return
	}
	defer conn.Close()
	for {
		_,msg,err:=conn.ReadMessage()
		if err != nil {
			log.Println("read msg failed,err:",err)
			return
		}
		fmt.Println("读到了消息",string(msg))
	}
}
func main() {
	//注册一个默认路由引擎
	r:=gin.Default()

	r.GET("/ws",ws)

	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
