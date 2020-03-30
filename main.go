package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/handler"
	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", utils.Index())
	})
	r.GET("/api/video/list", handler.VideoList())
	r.GET("/api/room/enter", handler.RoomEnter())
	r.GET("/api/room/status/read", handler.StatusRead())
	r.GET("/api/room/status/write", handler.StatusWrite())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
