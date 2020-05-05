package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/bootstrap"
	"github.com/day-dreams/vshare.zhangnan.xyz/config"
	"github.com/day-dreams/vshare.zhangnan.xyz/handler"
	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func main() {
	bootstrap.Bootstrap()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.GET("/version", func(c *gin.Context) {
		utils.GinJson(c, map[string]interface{}{
			"version": config.GetString("version"),
		}, nil)
	})

	r.GET("/", handler.Index())
	r.GET("/api/video/info/demo", handler.VideoInfoDemo())
	r.GET("/api/video/hls/playlist", handler.M3u8PlayList())
	r.GET("/api/video/hls/segment", handler.M3u8Segment())

	addr := os.Getenv("addr")
	r.Run(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
