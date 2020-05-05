package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/bootstrap"
	"github.com/day-dreams/vshare.zhangnan.xyz/config"
	"github.com/day-dreams/vshare.zhangnan.xyz/service"
	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func main() {
	bootstrap.Bootstrap()

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.GET("/version", func(c *gin.Context) {
		utils.GinJson(c, map[string]interface{}{
			"version": config.GetString("version"),
		}, nil)
	})

	r.GET("/api/video/info/demo", func(c *gin.Context) {
		info, err := service.VideoInfoGet(c, "/data/huojianshaonv101.MP4")
		utils.GinJson(c, info, err)
	})

	addr := os.Getenv("addr")
	r.Run(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
