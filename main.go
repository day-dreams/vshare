package main

import (
	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		utils.GinJson(c, "hello world", nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
