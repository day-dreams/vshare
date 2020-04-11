package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/location"

	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/handler"
	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(location.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", utils.Index())
	})
	r.GET("/api/video/list", handler.VideoList())
	r.GET("/api/room/enter", handler.RoomEnter())
	r.GET("/api/room/status/read", handler.StatusRead())
	r.GET("/api/room/status/write", handler.StatusWrite())
	r.GET("/api/play/m3u8/playlist", handler.M3u8PlayList())
	r.GET("/api/play/m3u8/ts", handler.M3u8TsFile())

	r.Run(os.Getenv("addr")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
