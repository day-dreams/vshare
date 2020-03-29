package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/handler"
	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", utils.Index())
	})
	r.GET("/api/video/list", handler.VideoList())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
