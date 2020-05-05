package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/config"
	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func main() {
	config.InitConfig()
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.GET("/version", func(c *gin.Context) {
		utils.GinJson(c, map[string]interface{}{
			"version": config.GetString("version"),
		}, nil)
	})

	addr := os.Getenv("addr")
	r.Run(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
