package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// UP 设置websocket
//CheckOrigin防止跨站点的请求伪造
var UP= websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}


// ping
// @Description: websocket 实现
// @param c
//
func handleFunc(c *gin.Context) {
	// 升级get请求为webSocket协议
	ws, err := UP.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close() //返回前关闭
	for {
		// 读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// 写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func main() {
	r := gin.Default()
	// LoadHTMLGlob()方法可以加载模板文件
	r.LoadHTMLGlob("templates/*")
	// 路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"websocket.html",nil)
	})
	r.GET("/ws", handleFunc)

	r.Run("localhost:8000")
}
